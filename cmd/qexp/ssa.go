package main

import (
	"fmt"
	"go/constant"
	"go/token"
	"go/types"
	"log"
	"strings"

	"golang.org/x/tools/go/loader"
	"golang.org/x/tools/go/ssa"
	"golang.org/x/tools/go/ssa/ssautil"
)

type Program struct {
	prog *ssa.Program
}

func loadProgram(path string) (*Program, error) {
	var cfg loader.Config
	cfg.Import(path)

	iprog, err := cfg.Load()
	if err != nil {
		return nil, fmt.Errorf("conf.Load failed: %s", err)
	}

	prog := ssautil.CreateProgram(iprog, ssa.SanityCheckFunctions|ssa.NaiveForm)
	prog.Build()
	return &Program{prog: prog}, nil
}

func (p *Program) DumpDeps(path string) {
	pkg := p.prog.ImportedPackage(path)
	for _, im := range pkg.Pkg.Imports() {
		fmt.Println(im.Path())
	}
}

func (p *Program) dumpDeps(path string, sep string) {
	pkg := p.prog.ImportedPackage(path)
	for _, im := range pkg.Pkg.Imports() {
		fmt.Println(sep, im.Path())
		p.dumpDeps(im.Path(), sep+"  ")
	}
}

func (p *Program) DumpExport(path string) {
	pkg := p.prog.ImportedPackage(path)
	for k, v := range pkg.Members {
		if token.IsExported(k) {
			fmt.Printf("%v %v %T\n", k, v, v)
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
	Methods map[string]reflect.Value
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
	Methods       []string
	Consts        []string
	TypedConsts   []string
	UntypedConsts []string
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

func (p *Program) constToLit(c constant.Value) string {
	switch c.Kind() {
	case constant.Bool:
		return fmt.Sprintf("constant.MakeBool(%v)", constant.BoolVal(c))
	case constant.String:
		return fmt.Sprintf("constant.MakeString(%q)", constant.StringVal(c))
	case constant.Int:
		if v, ok := constant.Int64Val(c); ok {
			return fmt.Sprintf("constant.MakeInt64(%v)", v)
		} else if v, ok := constant.Uint64Val(c); ok {
			return fmt.Sprintf("constant.MakeUint64(%v)", v)
		}
		return fmt.Sprintf("constant.MakeFromLiteral(%q, token.INT, 0)", c.ExactString())
	case constant.Float:
		s := c.ExactString()
		if pos := strings.IndexByte(s, '/'); pos >= 0 {
			x := fmt.Sprintf("constant.MakeFromLiteral(%q, token.FLOAT, 0)", s[:pos])
			y := fmt.Sprintf("constant.MakeFromLiteral(%q, token.FLOAT, 0)", s[pos+1:])
			return fmt.Sprintf("constant.BinaryOp(%v, token.QUO, %v)", x, y)
		}
		return fmt.Sprintf("constant.MakeFromLiteral(%q, token.FLOAT, 0)", c.ExactString())
	case constant.Complex:
		re := p.constToLit(constant.Real(c))
		im := p.constToLit(constant.Imag(c))
		return fmt.Sprintf("constant.BinaryOp(%v, token.ADD, constan.MakeImag(%v))", re, im)
	default:
		panic("unreachable")
	}
}

func (p *Program) ExportPkg(path string, sname string) *Package {
	pkg := p.prog.ImportedPackage(path)
	pkgPath := pkg.Pkg.Path()
	pkgName := pkg.Pkg.Name()
	e := &Package{Name: pkgName, Path: pkgPath}
	pkgName = sname
	for _, v := range pkg.Pkg.Imports() {
		e.Deps = append(e.Deps, fmt.Sprintf("%q: %q", v.Path(), v.Name()))
	}
	checked := make(map[ssa.Member]bool)
	for k, v := range pkg.Members {
		if token.IsExported(k) {
			if checked[v] {
				continue
			}
			checked[v] = true
			switch t := v.(type) {
			case *ssa.NamedConst:
				if typ := t.Type().String(); strings.HasPrefix(typ, "untyped ") {
					e.UntypedConsts = append(e.UntypedConsts, fmt.Sprintf("%q: {%q, %v}", t.Name(), t.Type().String(), p.constToLit(t.Value.Value)))
				} else {
					e.TypedConsts = append(e.TypedConsts, fmt.Sprintf("%q : {reflect.TypeOf(%v), %v}", t.Name(), pkgName+"."+t.Name(), p.constToLit(t.Value.Value)))
				}
			case *ssa.Global:
				e.Vars = append(e.Vars, fmt.Sprintf("%q : reflect.ValueOf(&%v)", t.Name(), pkgName+"."+t.Name()))
			case *ssa.Function:
				e.Funcs = append(e.Funcs, fmt.Sprintf("%q : reflect.ValueOf(%v)", t.Name(), pkgName+"."+t.Name()))
			case *ssa.Type:
				typ := t.Type()
				if obj, ok := t.Object().(*types.TypeName); ok && obj.IsAlias() {
					name := obj.Name()
					switch typ := obj.Type().(type) {
					case *types.Named:
						e.AliasTypes = append(e.AliasTypes, fmt.Sprintf("%q: reflect.TypeOf((*%v.%v)(nil)).Elem()", name, sname, name))
					case *types.Basic:
						e.AliasTypes = append(e.AliasTypes, fmt.Sprintf("%q: reflect.TypeOf((*%v)(nil)).Elem()", name, typ.Name()))
					default:
						log.Panicln("error parser", typ)
					}
					continue
				}
				var ms, pms []string
				recvId := typ.String()
				typeName := typ.(*types.Named).Obj().Name()
				methods := IntuitiveMethodSet(typ)
				for _, method := range methods {
					name := method.Obj().Name()
					mid := method.Obj().Type().(*types.Signature).Recv().Type().String()
					if mid[0] == '*' {
						if mid[1:] == recvId {
							pms = append(pms, name)
						}
					} else if mid == recvId {
						ms = append(ms, name)
					}
					if token.IsExported(name) {
						info := fmt.Sprintf("(%v).%v", method.Recv(), name)
						if pkgPath == pkgName {
							e.Methods = append(e.Methods, fmt.Sprintf("%q : reflect.ValueOf(%v)", info, info))
						} else {
							var fn string
							if isPointer(method.Recv()) {
								fn = fmt.Sprintf("(*%v.%v).%v", pkgName, typeName, name)
							} else {
								fn = fmt.Sprintf("(%v.%v).%v", pkgName, typeName, name)
							}
							e.Methods = append(e.Methods, fmt.Sprintf("%q: reflect.ValueOf(%v)", info, fn))
						}
					}
				}
				if types.IsInterface(typ) {
					e.Interfaces = append(e.Interfaces, fmt.Sprintf("%q : reflect.TypeOf((*%v.%v)(nil)).Elem()", typeName, pkgName, typeName))
				} else {
					e.NamedTypes = append(e.NamedTypes, fmt.Sprintf("%q : {reflect.TypeOf((*%v.%v)(nil)).Elem(),\"%v\",\"%v\"}", typeName, pkgName, typeName,
						strings.Join(ms, ","), strings.Join(pms, ",")))
				}
			default:
				panic("unreachable")
			}
		}
	}
	return e
}

func (p *Program) Export(path string) (extList []string, typList []string) {
	pkg := p.prog.ImportedPackage(path)
	pkgPath := pkg.Pkg.Path()
	pkgName := pkg.Pkg.Name()
	for k, v := range pkg.Members {
		if token.IsExported(k) {
			switch t := v.(type) {
			case *ssa.NamedConst:
			case *ssa.Global:
				extList = append(extList, fmt.Sprintf("%q : &%v", pkgPath+"."+t.Name(), pkgName+"."+t.Name()))
			case *ssa.Function:
				extList = append(extList, fmt.Sprintf("%q : %v", pkgPath+"."+t.Name(), pkgName+"."+t.Name()))
			case *ssa.Type:
				named := t.Type().(*types.Named)
				typeName := named.Obj().Name()

				typList = append(typList, fmt.Sprintf("(*%v.%v)(nil)", pkgName, typeName))
				if named.Obj().Pkg() != pkg.Pkg {
					continue
				}
				methods := IntuitiveMethodSet(t.Type())
				for _, method := range methods {
					name := method.Obj().Name()
					if token.IsExported(name) {
						info := fmt.Sprintf("(%v).%v", method.Recv(), name)
						if pkgPath == pkgName {
							extList = append(extList, fmt.Sprintf("%q : %v", info, info))
						} else {
							var fn string
							if isPointer(method.Recv()) {
								fn = fmt.Sprintf("(*%v.%v).%v", pkgName, typeName, name)
							} else {
								fn = fmt.Sprintf("(%v.%v).%v", pkgName, typeName, name)
							}
							extList = append(extList, fmt.Sprintf("%q : %v", info, fn))
						}
					}
				}
			}
		}
	}
	return
}

func isPointer(T types.Type) bool {
	_, ok := T.(*types.Pointer)
	return ok
}

// golang.org/x/tools/go/types/typeutil.IntuitiveMethodSet
func IntuitiveMethodSet(T types.Type) []*types.Selection {
	isPointerToConcrete := func(T types.Type) bool {
		ptr, ok := T.(*types.Pointer)
		return ok && !types.IsInterface(ptr.Elem())
	}

	var result []*types.Selection
	mset := types.NewMethodSet(T)
	if types.IsInterface(T) || isPointerToConcrete(T) {
		for i, n := 0, mset.Len(); i < n; i++ {
			result = append(result, mset.At(i))
		}
	} else {
		// T is some other concrete type.
		// Report methods of T and *T, preferring those of T.
		pmset := types.NewMethodSet(types.NewPointer(T))
		for i, n := 0, pmset.Len(); i < n; i++ {
			meth := pmset.At(i)
			if m := mset.Lookup(meth.Obj().Pkg(), meth.Obj().Name()); m != nil {
				meth = m
			}
			result = append(result, meth)
		}
	}
	return result
}
