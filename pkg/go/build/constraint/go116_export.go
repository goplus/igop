// export by github.com/goplus/igop/cmd/qexp

//+build go1.16,!go1.17

package constraint

import (
	q "go/build/constraint"

	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
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
		NamedTypes: map[string]igop.NamedType{
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
		TypedConsts:   map[string]igop.TypedConst{},
		UntypedConsts: map[string]igop.UntypedConst{},
	})
}
