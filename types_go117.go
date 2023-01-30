//go:build !go1.18
// +build !go1.18

package igop

import (
	"go/ast"
	"go/types"
	"reflect"

	"golang.org/x/tools/go/ssa"
)

const (
	enabledTypeParam = false
)

func hasTypeParam(t types.Type) bool {
	return false
}

func (r *TypesRecord) SetFunction(fn *ssa.Function) {
}

func (r *TypesRecord) extractNamed(named *types.Named, totype bool) (pkgpath string, name string, typeargs bool, nested bool) {
	obj := named.Obj()
	if pkg := obj.Pkg(); pkg != nil {
		if pkg.Name() == "main" {
			pkgpath = "main"
		} else {
			pkgpath = pkg.Path()
		}
	}
	name = obj.Name()
	return
}

func (r *TypesRecord) LookupReflect(typ types.Type) (rt reflect.Type, ok bool, nested bool) {
	rt, ok = r.loader.LookupReflect(typ)
	if !ok {
		if rt := r.tcache.At(typ); rt != nil {
			return rt.(reflect.Type), true, false
		}
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
		}
		if err := types.NewChecker(sp.Context.conf, sp.Context.FileSet, sp.Package, sp.Info).Files(sp.Files); err != nil {
			return err
		}
	}
	return
}
