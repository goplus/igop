package export

import (
	"fmt"
	"go/build"
	"go/constant"
	"go/token"
	"go/types"
	"log"
	"strings"

	"golang.org/x/tools/go/loader"
)

type Program struct {
	prog *loader.Program
	ctx  *build.Context
}

func NewProgram(ctx *build.Context) *Program {
	return &Program{ctx: ctx}
}

func (p *Program) Load(pkgs []string) error {
	var cfg loader.Config
	cfg.Build = p.ctx
	for _, pkg := range pkgs {
		cfg.Import(pkg)
	}

	iprog, err := cfg.Load()
	if err != nil {
		return fmt.Errorf("conf.Load failed: %s", err)
	}
	p.prog = iprog
	return nil
}

func loadProgram(path string, ctx *build.Context) (*Program, error) {
	var cfg loader.Config
	cfg.Build = ctx
	cfg.Import(path)

	iprog, err := cfg.Load()
	if err != nil {
		return nil, fmt.Errorf("conf.Load failed: %s", err)
	}
	return &Program{prog: iprog, ctx: ctx}, nil
}

func (p *Program) DumpDeps(path string) {
	pkg := p.prog.Package(path)
	for _, im := range pkg.Pkg.Imports() {
		fmt.Println(im.Path())
	}
}

func (p *Program) dumpDeps(path string, sep string) {
	pkg := p.prog.Package(path)
	for _, im := range pkg.Pkg.Imports() {
		fmt.Println(sep, im.Path())
		p.dumpDeps(im.Path(), sep+"  ")
	}
}

func (p *Program) DumpExport(path string) {
	pkg := p.prog.Package(path)
	for _, v := range pkg.Pkg.Scope().Names() {
		if token.IsExported(v) {
			fmt.Println(v)
		}
	}
}

/*
type ConstValue struct {
	Typ   string
	Value constant.Value
}

type Package struct {
	Name    string
	Path    string
	Types   []reflect.Type
	Vars    map[string]reflect.Value
	Funcs   map[string]reflect.Value
	Consts  map[string]ConstValue
	Deps    map[string]string
}

*/

type Package struct {
	Name          string
	Path          string
	Deps          []string
	NamedTypes    []string
	Interfaces    []string
	AliasTypes    []string
	Vars          []string
	Funcs         []string
	Consts        []string
	TypedConsts   []string
	UntypedConsts []string
}

func (p *Package) IsEmpty() bool {
	return len(p.NamedTypes) == 0 && len(p.Interfaces) == 0 &&
		len(p.AliasTypes) == 0 && len(p.Vars) == 0 &&
		len(p.Funcs) == 0 && len(p.Consts) == 0 &&
		len(p.TypedConsts) == 0 && len(p.UntypedConsts) == 0
}

/*
func unmarshalFloat(str string) constant.Value {
	if sep := strings.IndexByte(str, '/'); sep >= 0 {
		x := constant.MakeFromLiteral(str[:sep], token.FLOAT, 0)
		y := constant.MakeFromLiteral(str[sep+1:], token.FLOAT, 0)
		return constant.BinaryOp(x, token.QUO, y)
	}
	return constant.MakeFromLiteral(str, token.FLOAT, 0)
}
*/

func (p *Program) constToLit(named string, c constant.Value) string {
	switch c.Kind() {
	case constant.Bool:
		if named != "" {
			return fmt.Sprintf("constant.MakeBool(bool(%v))", named)
		}
		return fmt.Sprintf("constant.MakeBool(%v)", constant.BoolVal(c))
	case constant.String:
		if named != "" {
			return fmt.Sprintf("constant.MakeString(string(%v))", named)
		}
		return fmt.Sprintf("constant.MakeString(%q)", constant.StringVal(c))
	case constant.Int:
		if v, ok := constant.Int64Val(c); ok {
			if named != "" {
				return fmt.Sprintf("constant.MakeInt64(int64(%v))", named)
			}
			return fmt.Sprintf("constant.MakeInt64(%v)", v)
		} else if v, ok := constant.Uint64Val(c); ok {
			if named != "" {
				return fmt.Sprintf("constant.MakeUint64(uint64(%v))", named)
			}
			return fmt.Sprintf("constant.MakeUint64(%v)", v)
		}
		return fmt.Sprintf("constant.MakeFromLiteral(%q, token.INT, 0)", c.ExactString())
	case constant.Float:
		s := c.ExactString()
		if pos := strings.IndexByte(s, '/'); pos >= 0 {
			sx := s[:pos]
			sy := s[pos+1:]
			// simplify 314/100 => 3.14
			// 80901699437494742410229341718281905886015458990288143106772431
			// 50000000000000000000000000000000000000000000000000000000000000
			if strings.HasPrefix(sy, "1") && strings.Count(sy, "0") == len(sy)-1 {
				if len(sx) == len(sy) {
					return fmt.Sprintf("constant.MakeFromLiteral(\"%v.%v\", token.FLOAT, 0)", sx[:1], sx[1:])
				} else if len(sx) == len(sy)-1 {
					return fmt.Sprintf("constant.MakeFromLiteral(\"0.%v\", token.FLOAT, 0)", sx)
				} else if len(sx) < len(sy) {
					return fmt.Sprintf("constant.MakeFromLiteral(\"%v.%ve-%v\", token.FLOAT, 0)", sx[:1], sx[1:], len(sy)-len(sx))
				}
			} else if strings.HasPrefix(sy, "5") && strings.Count(sy, "0") == len(sy)-1 {
				if len(sx) == len(sy) {
					c := constant.BinaryOp(constant.MakeFromLiteral(sx, token.INT, 0), token.MUL, constant.MakeInt64(2))
					sx = c.ExactString()
					return fmt.Sprintf("constant.MakeFromLiteral(\"%v.%v\", token.FLOAT, 0)", sx[:1], sx[1:])
				}
			} else if strings.HasPrefix(sx, "1") && strings.Count(sx, "0") == len(sx)-1 {
				// skip
			}
			x := fmt.Sprintf("constant.MakeFromLiteral(%q, token.INT, 0)", sx)
			y := fmt.Sprintf("constant.MakeFromLiteral(%q, token.INT, 0)", sy)
			return fmt.Sprintf("constant.BinaryOp(%v, token.QUO, %v)", x, y)
		}
		if pos := strings.LastIndexAny(s, "123456789"); pos != -1 {
			sx := s[:pos+1]
			return fmt.Sprintf("constant.MakeFromLiteral(\"%v.%ve+%v\", token.FLOAT, 0)", sx[:1], sx[1:], len(s)-1)
		}
		return fmt.Sprintf("constant.MakeFromLiteral(%q, token.FLOAT, 0)", s)
	case constant.Complex:
		re := p.constToLit("", constant.Real(c))
		im := p.constToLit("", constant.Imag(c))
		return fmt.Sprintf("constant.BinaryOp(%v, token.ADD, constan.MakeImag(%v))", re, im)
	default:
		panic("unreachable")
	}
}

func (p *Program) ExportPkg(path string, sname string) (*Package, error) {
	info := p.prog.Package(path)
	if info == nil {
		return nil, fmt.Errorf("not found path %v", path)
	}
	pkg := info.Pkg
	pkgPath := pkg.Path()
	pkgName := pkg.Name()
	e := &Package{Name: pkgName, Path: pkgPath}
	pkgName = sname
	for _, v := range pkg.Imports() {
		e.Deps = append(e.Deps, fmt.Sprintf("%q: %q", v.Path(), v.Name()))
	}
	for _, name := range pkg.Scope().Names() {
		if !token.IsExported(name) {
			continue
		}
		obj := pkg.Scope().Lookup(name)
		switch t := obj.(type) {
		case *types.Const:
			named := pkgName + "." + t.Name()
			if typ := t.Type().String(); strings.HasPrefix(typ, "untyped ") {
				e.UntypedConsts = append(e.UntypedConsts, fmt.Sprintf("%q: {%q, %v}", t.Name(), t.Type().String(), p.constToLit(named, t.Val())))
			} else {
				e.TypedConsts = append(e.TypedConsts, fmt.Sprintf("%q : {reflect.TypeOf(%v), %v}", t.Name(), pkgName+"."+t.Name(), p.constToLit(named, t.Val())))
			}
		case *types.Var:
			e.Vars = append(e.Vars, fmt.Sprintf("%q : reflect.ValueOf(&%v)", t.Name(), pkgName+"."+t.Name()))
		case *types.Func:
			e.Funcs = append(e.Funcs, fmt.Sprintf("%q : reflect.ValueOf(%v)", t.Name(), pkgName+"."+t.Name()))
		case *types.TypeName:
			if IsTypeParam(t.Type()) {
				log.Println("skip typeparam", t)
				continue
			}
			if t.IsAlias() {
				name := obj.Name()
				switch typ := obj.Type().(type) {
				case *types.Basic:
					e.AliasTypes = append(e.AliasTypes, fmt.Sprintf("%q: reflect.TypeOf((*%v)(nil)).Elem()", name, typ.Name()))
				// case *types.Named:
				// 	e.AliasTypes = append(e.AliasTypes, fmt.Sprintf("%q: reflect.TypeOf((*%v.%v)(nil)).Elem()", name, sname, name))
				default:
					e.AliasTypes = append(e.AliasTypes, fmt.Sprintf("%q: reflect.TypeOf((*%v.%v)(nil)).Elem()", name, sname, name))
				}
				continue
			}
			typeName := t.Name()
			if types.IsInterface(t.Type()) {
				e.Interfaces = append(e.Interfaces, fmt.Sprintf("%q : reflect.TypeOf((*%v.%v)(nil)).Elem()", typeName, pkgName, typeName))
			} else {
				e.NamedTypes = append(e.NamedTypes, fmt.Sprintf("%q : reflect.TypeOf((*%v.%v)(nil)).Elem()", typeName, pkgName, typeName))
			}
		default:
			log.Panicf("unreachable %v %T\n", name, t)
		}
	}
	return e, nil
}
