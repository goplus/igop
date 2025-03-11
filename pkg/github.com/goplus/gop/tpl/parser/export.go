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
			"Mode": reflect.TypeOf((*q.Mode)(nil)).Elem(),
		},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"ParseFile": reflect.ValueOf(q.ParseFile),
		},
		TypedConsts:   map[string]igop.TypedConst{},
		UntypedConsts: map[string]igop.UntypedConst{},
	})
}
