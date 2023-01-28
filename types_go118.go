//go:build go1.18
// +build go1.18

package igop

import (
	"go/ast"
	"go/types"
	"reflect"
	"strings"

	"github.com/goplus/reflectx"
	"golang.org/x/tools/go/ssa"
)

const (
	enabledTypeParam = true
)

func hasTypeParam(t types.Type) bool {
	switch t := t.(type) {
	case *types.TypeParam:
		return true
	case *types.Named:
		return t.TypeParams() != nil
	case *types.Signature:
		return t.TypeParams() != nil
	}
	return false
}

func typeId(t reflect.Type) string {
	path := t.PkgPath()
	if path == "" {
		return t.Name()
	}
	return path + "." + t.Name()
}

func (r *TypesRecord) SetFunction(fn *ssa.Function) {
	r.fntargs = r.parseFuncTypeArgs(fn)
}

func (r *TypesRecord) parseFuncTypeArgs(fn *ssa.Function) (targs string) {
	name := fn.Name()
	pos := strings.Index(name, "[")
	if pos < 0 {
		return ""
	}
	v := reflectx.FieldByNameX(reflect.ValueOf(fn).Elem(), "typeargs")
	var args []string
	for i := 0; i < v.Len(); i++ {
		t := r.ToType(v.Index(i).Interface().(types.Type))
		args = append(args, typeId(t))
	}
	return strings.Join(args, ",")
}

func (r *TypesRecord) extractNamed(named *types.Named, totype bool) (pkgpath string, name string, typeargs bool) {
	obj := named.Obj()
	if pkg := obj.Pkg(); pkg != nil {
		if pkg.Name() == "main" {
			pkgpath = "main"
		} else {
			pkgpath = pkg.Path()
		}
	}
	name = obj.Name()
	var ids string = r.fntargs
	if args := named.TypeArgs(); args != nil {
		typeargs = true
		var targs []string
		for i := 0; i < args.Len(); i++ {
			if totype {
				t := r.ToType(args.At(i))
				targs = append(targs, typeId(t))
			} else {
				targs = append(targs, args.At(i).String())
			}
		}
		if ids != "" {
			ids += ";"
		}
		ids += strings.Join(targs, ",")
	}
	if ids != "" {
		name += "[" + ids + "]"
	}
	return
}

func (r *TypesRecord) extractNamed2(named *types.Named) (pkgpath string, name string) {
	obj := named.Obj()
	if pkg := obj.Pkg(); pkg != nil {
		pkgpath = pkg.Path()
	}
	name = obj.Name()
	var ids string = r.fntargs
	if args := named.TypeArgs(); args != nil {
		var targs []string
		for i := 0; i < args.Len(); i++ {
			t := r.ToType(args.At(i))
			targs = append(targs, typeId(t))
		}
		if ids != "" {
			ids += ";"
		}
		ids += strings.Join(targs, ",")
	}
	if ids != "" {
		name += "[" + ids + "]"
	}
	return
}

func (sp *sourcePackage) Load() (err error) {
	if sp.Info == nil {
		sp.Info = &types.Info{
			Types:      make(map[ast.Expr]types.TypeAndValue),
			Defs:       make(map[*ast.Ident]types.Object),
			Uses:       make(map[*ast.Ident]types.Object),
			Implicits:  make(map[ast.Node]types.Object),
			Scopes:     make(map[ast.Node]*types.Scope),
			Selections: make(map[*ast.SelectorExpr]*types.Selection),
			Instances:  make(map[*ast.Ident]types.Instance),
		}
		if err := types.NewChecker(sp.Context.conf, sp.Context.FileSet, sp.Package, sp.Info).Files(sp.Files); err != nil {
			return err
		}
	}
	return
}
