package igop

import (
	"fmt"
	"go/ast"
	"go/scanner"
	"go/token"
	"go/types"
	"os"
	"strings"

	"golang.org/x/tools/go/ssa"
)

/*
package main

import "fmt"

const/var/type/func

func main(){}
*/

type Repl struct {
	ctx        *Context
	fset       *token.FileSet
	pkg        *ssa.Package
	frame      *frame   // last frame
	pc         int      // last pc
	fsInit     *fnState // func init
	fsMain     *fnState // func main
	imports    []string // import lines
	globals    []string // global var/func/type
	infuncs    []string // in main func
	source     string   // all source
	lastDump   []string // last __igop_repl_dump__
	globalMap  map[string]interface{}
	lastInterp *Interp
	fileName   string
}

func toDump(i []interface{}) (dump []string) {
	for _, v := range i {
		dump = append(dump, fmt.Sprintf("%v", v))
	}
	return
}

func NewRepl(ctx *Context) *Repl {
	r := &Repl{
		ctx:       ctx,
		fset:      token.NewFileSet(),
		globalMap: make(map[string]interface{}),
	}
	ctx.evalMode = true
	ctx.evalInit = make(map[string]bool)

	ctx.SetOverrideFunction("main.__igop_repl_dump__", func(v ...interface{}) {
		r.lastDump = toDump(v)
	})
	ctx.evalCallFn = func(call *ssa.Call, res ...interface{}) {
		if strings.HasPrefix(call.Call.Value.Name(), "__igop_repl_") {
			return
		}
		r.lastDump = toDump(res)
	}
	return r
}

func (r *Repl) SetFileName(filename string) {
	r.fileName = filename
}

func (r *Repl) Eval(expr string) (tok token.Token, dump []string, err error) {
	r.lastDump = nil
	tok = r.firstToken(expr)
	err = r.eval(tok, expr)
	return tok, r.lastDump, err
}

func (r *Repl) Source() string {
	return r.source
}

func (r *Repl) buildSource(expr string, tok token.Token) string {
	v0 := strings.Join(r.imports, "\n")
	v1 := strings.Join(r.globals, "\n")
	v2 := strings.Join(r.infuncs, "\n")
	switch tok {
	case token.IMPORT:
		v0 += "\n" + expr
	case token.CONST, token.VAR, token.TYPE, token.FUNC:
		v1 += "\n" + expr
	default:
		v2 += "\n" + expr
	}
	return fmt.Sprintf(`package main
%v
func __igop_repl_used__(v interface{}){}
func __igop_repl_dump__(v ...interface{}){}
%v
func main() {
	%v
}
`, v0, v1, v2)
}

func (r *Repl) eval(tok token.Token, expr string) (err error) {
	var src string
	var inMain bool
	switch tok {
	case token.PACKAGE:
		// skip package
		return nil
	case token.IMPORT:
		src = r.buildSource(expr, tok)
		errors, err := r.check(r.fileName, src)
		if err != nil {
			return err
		}
		if errors != nil {
			return errors[0]
		}
		r.imports = append(r.imports, expr)
		return nil
	case token.CONST, token.FUNC, token.TYPE, token.VAR:
		src = r.buildSource(expr, tok)
		errors, err := r.check(r.fileName, src)
		if err != nil {
			return err
		}
		if errors != nil {
			return errors[0]
		}
	default:
		inMain = true
		src = r.buildSource(expr, tok)
		errors, err := r.check(r.fileName, src)
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
				fixed = append(fixed, "__igop_repl_used__(&"+v+")")
				fixed = append(fixed, "__igop_repl_dump__("+v+")")
			} else if strings.HasSuffix(e.Msg, errIsNotUsed) {
				expr = "__igop_repl_dump__(" + expr + ")"
			} else {
				return e
			}
		}
		if len(fixed) != 0 {
			expr += "\n" + strings.Join(fixed, "\n")
		}
		src = r.buildSource(expr, tok)
	}
	r.pkg, err = r.ctx.LoadFile(r.fset, r.fileName, src)
	if err != nil {
		return err
	}
	err = r.run()
	if err == nil {
		if inMain {
			r.infuncs = append(r.infuncs, expr)
		} else {
			r.globals = append(r.globals, expr)
		}
		r.source = src
	}
	return err
}

const (
	errDeclNotUsed = "declared but not used"
	errIsNotUsed   = "is not used"
)

func (r *Repl) check(filename string, src interface{}) (errors []error, err error) {
	file, err := r.ctx.ParseFile(r.fset, filename, src)
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

func (r *Repl) firstToken(src string) token.Token {
	var s scanner.Scanner
	file := r.fset.AddFile("", r.fset.Base(), len(src))
	s.Init(file, []byte(src), nil, 0)
	_, tok, _ := s.Scan()
	return tok
}

func (repl *Repl) run() error {
	i, err := newInterp(repl.ctx, repl.pkg, repl.globalMap)
	if err != nil {
		return err
	}
	rinit, err := repl.runFunc(i, "init", repl.fsInit)
	if err == nil {
		repl.fsInit = rinit
		repl.fsInit.pc--
	}
	rmain, err := repl.runFunc(i, "main", repl.fsMain)
	if err != nil {
		return err
	}
	repl.fsMain = rmain
	for k, v := range i.globals {
		repl.globalMap[k.String()] = v
	}
	return nil
}

type fnState struct {
	fr *frame
	pc int
}

func (r *Repl) runFunc(i *Interp, fnname string, fs *fnState) (rfs *fnState, err error) {
	defer func() {
		switch p := recover().(type) {
		case nil:
		case exitPanic:
			os.Exit(int(p))
		default:
			err = fmt.Errorf("%v", r)
		}
	}()
	fn := r.pkg.Func(fnname)
	if fn == nil {
		return nil, fmt.Errorf("no function %v", fnname)
	}
	pfn := i.funcs[fn]
	fr := pfn.allocFrame(nil)
	if fs != nil {
		copy(fr.stack, fs.fr.stack)
		fr.pc = fs.pc
	}
	var pc int
	for fr.pc != -1 {
		fn := fr.pfn.Instrs[fr.pc]
		pc = fr.pc
		fr.pc++
		fn(fr)
	}
	return &fnState{fr: fr, pc: pc}, nil
}
