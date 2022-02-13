// export by github.com/goplus/gossa/cmd/qexp

//go:build go1.17
// +build go1.17

package constraint

import (
	q "go/build/constraint"

	"reflect"

	"github.com/goplus/gossa"
)

func init() {
	gossa.RegisterPackage(&gossa.Package{
		Name: "constraint",
		Path: "go/build/constraint",
		Deps: map[string]string{
			"errors":       "errors",
			"strings":      "strings",
			"unicode":      "unicode",
			"unicode/utf8": "utf8",
		},
		Interfaces: map[string]reflect.Type{
			"Expr": reflect.TypeOf((*q.Expr)(nil)).Elem(),
		},
		NamedTypes: map[string]gossa.NamedType{
			"AndExpr":     {reflect.TypeOf((*q.AndExpr)(nil)).Elem(), "", "Eval,String,isExpr"},
			"NotExpr":     {reflect.TypeOf((*q.NotExpr)(nil)).Elem(), "", "Eval,String,isExpr"},
			"OrExpr":      {reflect.TypeOf((*q.OrExpr)(nil)).Elem(), "", "Eval,String,isExpr"},
			"SyntaxError": {reflect.TypeOf((*q.SyntaxError)(nil)).Elem(), "", "Error"},
			"TagExpr":     {reflect.TypeOf((*q.TagExpr)(nil)).Elem(), "", "Eval,String,isExpr"},
		},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"IsGoBuild":      reflect.ValueOf(q.IsGoBuild),
			"IsPlusBuild":    reflect.ValueOf(q.IsPlusBuild),
			"Parse":          reflect.ValueOf(q.Parse),
			"PlusBuildLines": reflect.ValueOf(q.PlusBuildLines),
		},
		TypedConsts:   map[string]gossa.TypedConst{},
		UntypedConsts: map[string]gossa.UntypedConst{},
	})
}
