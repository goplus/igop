// export by github.com/goplus/igop/cmd/qexp

package tpl

import (
	q "github.com/goplus/gop/tpl"

	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "tpl",
		Path: "github.com/goplus/gop/tpl",
		Deps: map[string]string{
			"fmt":                               "fmt",
			"github.com/goplus/gop/parser/iox":  "iox",
			"github.com/goplus/gop/tpl/ast":     "ast",
			"github.com/goplus/gop/tpl/cl":      "cl",
			"github.com/goplus/gop/tpl/matcher": "matcher",
			"github.com/goplus/gop/tpl/parser":  "parser",
			"github.com/goplus/gop/tpl/scanner": "scanner",
			"github.com/goplus/gop/tpl/token":   "token",
			"github.com/goplus/gop/tpl/types":   "types",
			"github.com/qiniu/x/errors":         "errors",
			"io":                                "io",
			"os":                                "os",
			"reflect":                           "reflect",
		},
		Interfaces: map[string]reflect.Type{
			"Scanner": reflect.TypeOf((*q.Scanner)(nil)).Elem(),
		},
		NamedTypes: map[string]reflect.Type{
			"Compiler":   reflect.TypeOf((*q.Compiler)(nil)).Elem(),
			"Config":     reflect.TypeOf((*q.Config)(nil)).Elem(),
			"MatchState": reflect.TypeOf((*q.MatchState)(nil)).Elem(),
		},
		AliasTypes: map[string]reflect.Type{
			"Error": reflect.TypeOf((*q.Error)(nil)).Elem(),
			"Token": reflect.TypeOf((*q.Token)(nil)).Elem(),
		},
		Vars: map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"BasicLit":     reflect.ValueOf(q.BasicLit),
			"BinaryExpr":   reflect.ValueOf(q.BinaryExpr),
			"BinaryExprNR": reflect.ValueOf(q.BinaryExprNR),
			"BinaryExprR":  reflect.ValueOf(q.BinaryExprR),
			"BinaryOp":     reflect.ValueOf(q.BinaryOp),
			"BinaryOpNR":   reflect.ValueOf(q.BinaryOpNR),
			"BinaryOpR":    reflect.ValueOf(q.BinaryOpR),
			"Dump":         reflect.ValueOf(q.Dump),
			"Fdump":        reflect.ValueOf(q.Fdump),
			"FromFile":     reflect.ValueOf(q.FromFile),
			"Ident":        reflect.ValueOf(q.Ident),
			"List":         reflect.ValueOf(q.List),
			"New":          reflect.ValueOf(q.New),
			"NewEx":        reflect.ValueOf(q.NewEx),
			"Panic":        reflect.ValueOf(q.Panic),
			"RangeOp":      reflect.ValueOf(q.RangeOp),
			"Relocate":     reflect.ValueOf(q.Relocate),
			"ShowConflict": reflect.ValueOf(q.ShowConflict),
			"UnaryExpr":    reflect.ValueOf(q.UnaryExpr),
		},
		TypedConsts:   map[string]igop.TypedConst{},
		UntypedConsts: map[string]igop.UntypedConst{},
	})
}
