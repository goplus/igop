package gossa

import (
	"fmt"
	"go/ast"
	"go/format"
	"go/scanner"
	"go/token"
	"go/types"
	"strings"

	"golang.org/x/tools/go/ssa"
)

type Repl struct {
	ctx      *Context
	fset     *token.FileSet
	pkg      *ssa.Package
	frame    *frame // last frame
	pc       int    // last pc
	imports  []string
	lastExpr string
	lastSrc  string
	lastUsed []interface{}
}

func NewRepl(ctx *Context) *Repl {
	r := &Repl{
		ctx:  ctx,
		fset: token.NewFileSet(),
	}
	ctx.evalMode = true
	ctx.SetOverrideFunction("main.__gossa_repl_used__", func(v interface{}) {
		r.lastUsed = append(r.lastUsed, v)
	})
	return r
}

func (r *Repl) Eval(expr string) (v interface{}, err error) {
	r.lastUsed = nil
	err = r.eval(expr)
	return r.lastUsed, err
}

func (r *Repl) Source() string {
	return r.lastSrc
}

func (r *Repl) eval(expr string) (err error) {
	org := expr
	tok := r.firstToken(expr)
	var src string
	switch tok {
	case token.PACKAGE:
		// source
	case token.IMPORT:
		src = "package main;" + expr
		errors, err := r.check("main.gop", src)
		if err != nil {
			return err
		}
		if errors != nil {
			return errors[0]
		}
		r.imports = append(r.imports, expr)
		return nil
	case token.CONST, token.FUNC, token.TYPE, token.VAR:
		src = "package main;" + expr
	default:
		if r.lastExpr != "" {
			expr = r.lastExpr + "\n" + expr
		}
		src = r.wrapInMain(expr)
		errors, err := r.check("main.gop", src)
		if err != nil {
			return err
		}
		var fixed []string
		for _, err := range errors {
			e, ok := err.(types.Error)
			if !ok {
				return e
			}
			if strings.HasSuffix(e.Msg, errDeclNotUsed) {
				v := e.Msg[0 : len(e.Msg)-len(errDeclNotUsed)-1]
				fixed = append(fixed, "__gossa_repl_used__("+v+")")
			} else if strings.HasSuffix(e.Msg, errIsNotUsed) {
				org = "__gossa_repl_used__(" + org + ")"
			} else {
				return e
			}
		}
		expr = r.lastExpr + ";" + org + ";" + strings.Join(fixed, ";")
		src = r.wrapInMain(expr)
	}
	dst, err := format.Source([]byte(src))
	if err != nil {
		return err
	}
	src = string(dst)
	r.pkg, err = r.ctx.LoadFile(r.fset, "main.gop", src)
	if err != nil {
		return err
	}
	err = r.run()
	if err == nil {
		r.lastSrc = src
		r.lastExpr = expr
	}
	return err
}

const (
	errDeclNotUsed = "declared but not used"
	errIsNotUsed   = "is not used"
)

func (r *Repl) check(filename string, src interface{}) (errors []error, err error) {
	file, err := r.ctx.ParseFile(r.fset, "main.gop", src)
	if err != nil {
		return nil, err
	}
	tc := &types.Config{
		Importer:                 NewImporter(r.ctx.Loader, r.ctx.External),
		Sizes:                    r.ctx.Sizes,
		DisableUnusedImportCheck: true,
	}
	tc.Error = func(err error) {
		errors = append(errors, err)
	}
	pkg := types.NewPackage(file.Name.Name, "")
	types.NewChecker(tc, r.fset, pkg, &types.Info{}).Files([]*ast.File{file})
	return
}

func (r *Repl) wrapInMain(src string) string {
	return fmt.Sprintf(`package main
%v
func __gossa_repl_used__(v interface{}){}
func main() {
	%v
}
`, strings.Join(r.imports, "\n"), src)
}

func (r *Repl) firstToken(src string) token.Token {
	var s scanner.Scanner
	file := r.fset.AddFile("", r.fset.Base(), len(src))
	s.Init(file, []byte(src), nil, 0)
	_, tok, _ := s.Scan()
	return tok
}

func (repl *Repl) run() error {
	i, err := NewInterp(repl.ctx, repl.pkg)
	if err != nil {
		return err
	}
	err = i.RunInit()
	if err != nil {
		return err
	}
	fn := repl.pkg.Func("main")
	pfn := i.funcs[fn]
	fr := pfn.allocFrame(nil)
	if repl.frame != nil {
		copy(fr.stack, repl.frame.stack)
		fr.pc = repl.pc
	}
	for fr.pc != -1 {
		fn := fr.pfn.Instrs[fr.pc]
		repl.pc = fr.pc
		fr.pc++
		fn(fr)
	}
	repl.frame = fr
	return nil
}
