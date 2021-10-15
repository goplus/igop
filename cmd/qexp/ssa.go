package main

import (
	"fmt"
	"go/token"
	"go/types"

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

	prog := ssautil.CreateProgram(iprog, ssa.SanityCheckFunctions)
	prog.Build()
	return &Program{prog}, nil
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
			case *ssa.Const:
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
