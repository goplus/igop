package igop

import (
	"fmt"
	"go/ast"
	"go/constant"
	"go/scanner"
	"go/token"
	"go/types"
	"strings"

	xconst "github.com/goplus/igop/constant"
	"golang.org/x/tools/go/ssa"
)

/*
package main

import "fmt"

const/var/type/func

func main(){}
*/

type Repl struct {
	ctx       *Context
	pkg       *ssa.Package
	builtin   *ast.File
	interp    *Interp  // last interp
	fsInit    *fnState // func init
	fsMain    *fnState // func main
	globalMap map[string]interface{}
	fileName  string
	source    string   // all source
	imports   []string // import lines
	globals   []string // global var/func/type
	infuncs   []string // in main func
	lastDump  []string // last __igop_repl_dump__
}

func toDump(i []interface{}) (dump []string) {
	for _, v := range i {
		dump = append(dump, fmt.Sprintf("%v %T", v, v))
	}
	return
}

func toResultDump(interp *Interp, vs []interface{}, rs *types.Tuple) (dump []string) {
	for i, v := range vs {
		dump = append(dump, fmt.Sprintf("%v %v", v, interp.toType(rs.At(i).Type())))
	}
	return
}

func NewRepl(ctx *Context) *Repl {
	r := &Repl{
		ctx:       ctx,
		globalMap: make(map[string]interface{}),
	}
	ctx.SetEvalMode(true)
	RegisterCustomBuiltin("__igop_repl_used__", func(v interface{}) {})
	RegisterCustomBuiltin("__igop_repl_dump__", func(v ...interface{}) {
		r.lastDump = toDump(v)
	})
	ctx.evalCallFn = func(interp *Interp, call *ssa.Call, res ...interface{}) {
		if len(*call.Referrers()) != 0 {
			return
		}
		v := call.Call.Value
		if un, ok := v.(*ssa.UnOp); ok {
			v = un.X
		}
		if strings.Contains(v.Name(), "__igop_repl_") {
			return
		}
		r.lastDump = toResultDump(interp, res, call.Call.Signature().Results())
	}
	if f, err := ParseBuiltin(r.ctx.FileSet, "main"); err == nil {
		r.builtin = f
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

func (r *Repl) Interp() *Interp {
	return r.interp
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
%v
func main() {
	%v
}
`, v0, v1, v2)
}

func (r *Repl) eval(tok token.Token, expr string) (err error) {
	var src string
	var inMain bool
	var evalConst bool
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
	case token.FUNC:
		src = r.buildSource(expr, tok)
		errors, err := r.check(r.fileName, src)
		if err != nil {
			// check funlit
			msg := err.Error()
			if strings.Contains(msg, errMaybeGoFunLit) || strings.Contains(msg, errMaybeGopFunLit) {
				inMain = true
				src = r.buildSource(expr, token.ILLEGAL)
				errors, err = r.check(r.fileName, src)
			}
			if err != nil {
				return err
			}
		}
		if errors != nil {
			return errors[0]
		}
	case token.CONST, token.TYPE, token.VAR:
		src = r.buildSource(expr, tok)
		errors, err := r.check(r.fileName, src)
		if err != nil {
			return err
		}
		if errors != nil {
			return errors[0]
		}
	default:
		switch tok {
		case token.FOR, token.IF, token.SWITCH, token.SELECT:
			expr = "func(){\n" + expr + "\n}()"
		}
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
			if strings.HasSuffix(e.Msg, errDeclaredNotUsed) {
				v := e.Msg[0 : len(e.Msg)-len(errDeclaredNotUsed)-1]
				fixed = append(fixed, "__igop_repl_used__(&"+v+")")
				// fixed = append(fixed, "__igop_repl_dump__("+v+")")
			} else if strings.HasSuffix(e.Msg, errIsNotUsed) {
				if _, ok := extractConstant([]byte(e.Msg[:len(e.Msg)-len(errIsNotUsed)])); ok {
					// r.lastDump = []string{c.Lit}
					// return nil
					expr = "const __igop_repl_const__ = " + expr
					tok = token.CONST
					evalConst = true
				} else {
					expr = "__igop_repl_dump__(" + expr + ")"
				}
			} else {
				return e
			}
		}
		if len(fixed) != 0 {
			expr += "\n" + strings.Join(fixed, "\n")
		}
		src = r.buildSource(expr, tok)
	}
	r.pkg, err = r.ctx.LoadFile(r.fileName, src)
	if err != nil {
		return err
	}
	i, err := newInterp(r.ctx, r.pkg, r.globalMap)
	if err != nil {
		return err
	}
	rinit, err := r.runFunc(i, "init", r.fsInit)
	if err == nil {
		rinit.pc--
	}
	rmain, err := r.runFunc(i, "main", r.fsMain)
	if err != nil {
		return err
	}
	if evalConst {
		m := i.mainpkg.Members["__igop_repl_const__"]
		c, _ := m.(*ssa.NamedConst)
		s, _ := xconst.ExactConstantEx(c.Value.Value, true)
		kind := c.Value.Value.Kind()
		if kind == constant.Float {
			es := c.Value.Value.ExactString()
			r.lastDump = []string{fmt.Sprintf("%v %v\n(%v)", s, c.Type(), es)}
		} else {
			r.lastDump = []string{fmt.Sprintf("%v %v", s, c.Type())}
		}
		return nil
	}
	r.interp = i
	r.fsInit = rinit
	r.fsMain = rmain
	for k, v := range i.globals {
		r.globalMap[k.String()] = v
	}
	if inMain {
		r.infuncs = append(r.infuncs, expr)
	} else {
		r.globals = append(r.globals, expr)
	}
	r.source = src
	return nil
}

const (
	errIsNotUsed      = "is not used"
	errMaybeGoFunLit  = `expected 'IDENT', found '{'`
	errMaybeGopFunLit = `expected '(',`
)

func (r *Repl) check(filename string, src interface{}) (errors []error, err error) {
	file, err := r.ctx.ParseFile(filename, src)
	if err != nil {
		return nil, err
	}
	tc := &types.Config{
		Importer:                 NewImporter(r.ctx),
		DisableUnusedImportCheck: true,
	}
	tc.Error = func(err error) {
		errors = append(errors, err)
	}
	pkg := types.NewPackage(file.Name.Name, "")
	var files []*ast.File
	if r.builtin != nil {
		files = []*ast.File{r.builtin, file}
	} else {
		files = []*ast.File{file}
	}
	types.NewChecker(tc, r.ctx.FileSet, pkg, &types.Info{}).Files(files)
	return
}

func (r *Repl) firstToken(src string) token.Token {
	var s scanner.Scanner
	file := r.ctx.FileSet.AddFile("", -1, len(src))
	s.Init(file, []byte(src), nil, 0)
	_, tok, _ := s.Scan()
	return tok
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
			err = ExitError(int(p))
		default:
			err = fmt.Errorf("%v", p)
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
		fr.ipc = fs.pc
	}
	var pc int
	for fr.ipc != -1 {
		fn := fr.pfn.Instrs[fr.ipc]
		pc = fr.ipc
		fr.ipc++
		fn(fr)
	}
	return &fnState{fr: fr, pc: pc}, nil
}

type tokenLit struct {
	Lit string
	Pos token.Pos
	Tok token.Token
}

// extractConstant extract constant int and overflows float/complex
func extractConstant(src []byte) (constant *tokenLit, ok bool) {
	var s scanner.Scanner
	fset := token.NewFileSet()
	file := fset.AddFile("", fset.Base(), len(src))
	s.Init(file, src, nil, 0)
	pos, tok, lit := s.Scan()
	if tok == token.EOF {
		return
	}
	first := &tokenLit{lit, pos, tok}
	var check int
	for {
		_, tok, lit := s.Scan()
		if tok == token.EOF {
			break
		}
		switch tok {
		case token.LPAREN:
			check++
		case token.RPAREN:
			check--
		case token.IDENT:
			if check == 1 && lit == "constant" {
				var nodes []*tokenLit
			loop:
				for {
					pos, tok, lit := s.Scan()
					if tok == token.EOF {
						break
					}
					switch tok {
					case token.LPAREN:
						check++
					case token.RPAREN:
						check--
						if check == 0 {
							break loop
						}
					default:
						nodes = append(nodes, &tokenLit{lit, pos, tok})
					}
				}
				switch len(nodes) {
				case 0:
					return first, true
				case 1:
					switch nodes[0].Tok {
					case token.INT:
						return nodes[0], true
					case token.FLOAT:
						return nodes[0], true
						// extract if not parse float64
						// if _, err := strconv.ParseFloat(nodes[0].Lit, 128); err != nil {
						// 	return nodes[0], true
						// }
					}
				default:
					last := nodes[len(nodes)-1]
					return &tokenLit{string(src[int(nodes[0].Pos)-1 : int(last.Pos)+len(last.Lit)]), nodes[0].Pos, token.IMAG}, true
				}
			}
		}
	}
	return
}
