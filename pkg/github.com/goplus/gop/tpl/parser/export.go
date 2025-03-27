// export by github.com/goplus/igop/cmd/qexp

package parser

import (
	q "github.com/goplus/gop/tpl/parser"

	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "parser",
		Path: "github.com/goplus/gop/tpl/parser",
		Deps: map[string]string{
			"github.com/goplus/gop/parser/iox":  "iox",
			"github.com/goplus/gop/tpl/ast":     "ast",
			"github.com/goplus/gop/tpl/scanner": "scanner",
			"github.com/goplus/gop/tpl/token":   "token",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{
			"Config": reflect.TypeOf((*q.Config)(nil)).Elem(),
		},
		AliasTypes: map[string]reflect.Type{
			"RetProcParser": reflect.TypeOf((*q.RetProcParser)(nil)).Elem(),
		},
		Vars: map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"ParseEx":   reflect.ValueOf(q.ParseEx),
			"ParseFile": reflect.ValueOf(q.ParseFile),
		},
		TypedConsts:   map[string]igop.TypedConst{},
		UntypedConsts: map[string]igop.UntypedConst{},
	})
}
