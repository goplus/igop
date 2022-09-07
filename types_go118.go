//go:build go1.18
// +build go1.18

package igop

import (
	"go/ast"
	"go/types"
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

func extractNamed(named *types.Named) (pkgpath string, name string) {
	obj := named.Obj()
	if pkg := obj.Pkg(); pkg != nil {
		pkgpath = pkg.Path()
	}
	name = obj.Name()
	if args := named.TypeArgs(); args != nil {
		name += "["
		for i := 0; i < args.Len(); i++ {
			if i != 0 {
				name += ","
			}
			name += args.At(i).String()
		}
		name += "]"
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
