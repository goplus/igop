// export by github.com/goplus/igop/cmd/qexp

package ast

import (
	q "github.com/goplus/gop/tpl/ast"

	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "ast",
		Path: "github.com/goplus/gop/tpl/ast",
		Deps: map[string]string{
			"github.com/goplus/gop/tpl/token": "token",
		},
		Interfaces: map[string]reflect.Type{
			"Decl": reflect.TypeOf((*q.Decl)(nil)).Elem(),
			"Expr": reflect.TypeOf((*q.Expr)(nil)).Elem(),
			"Node": reflect.TypeOf((*q.Node)(nil)).Elem(),
		},
		NamedTypes: map[string]reflect.Type{
			"BasicLit":   reflect.TypeOf((*q.BasicLit)(nil)).Elem(),
			"BinaryExpr": reflect.TypeOf((*q.BinaryExpr)(nil)).Elem(),
			"Choice":     reflect.TypeOf((*q.Choice)(nil)).Elem(),
			"File":       reflect.TypeOf((*q.File)(nil)).Elem(),
			"Ident":      reflect.TypeOf((*q.Ident)(nil)).Elem(),
			"Rule":       reflect.TypeOf((*q.Rule)(nil)).Elem(),
			"Sequence":   reflect.TypeOf((*q.Sequence)(nil)).Elem(),
			"UnaryExpr":  reflect.TypeOf((*q.UnaryExpr)(nil)).Elem(),
		},
		AliasTypes:    map[string]reflect.Type{},
		Vars:          map[string]reflect.Value{},
		Funcs:         map[string]reflect.Value{},
		TypedConsts:   map[string]igop.TypedConst{},
		UntypedConsts: map[string]igop.UntypedConst{},
	})
}
