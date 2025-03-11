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
			"github.com/goplus/gop/tpl/cl":      "cl",
			"github.com/goplus/gop/tpl/matcher": "matcher",
			"github.com/goplus/gop/tpl/parser":  "parser",
			"github.com/goplus/gop/tpl/scanner": "scanner",
			"github.com/goplus/gop/tpl/token":   "token",
			"github.com/goplus/gop/tpl/types":   "types",
			"io":                                "io",
			"os":                                "os",
		},
		Interfaces: map[string]reflect.Type{
			"Scanner": reflect.TypeOf((*q.Scanner)(nil)).Elem(),
		},
		NamedTypes: map[string]reflect.Type{
			"Compiler": reflect.TypeOf((*q.Compiler)(nil)).Elem(),
			"Config":   reflect.TypeOf((*q.Config)(nil)).Elem(),
		},
		AliasTypes: map[string]reflect.Type{
			"Token": reflect.TypeOf((*q.Token)(nil)).Elem(),
		},
		Vars: map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"Dump":     reflect.ValueOf(q.Dump),
			"Fdump":    reflect.ValueOf(q.Fdump),
			"FromFile": reflect.ValueOf(q.FromFile),
			"New":      reflect.ValueOf(q.New),
		},
		TypedConsts:   map[string]igop.TypedConst{},
		UntypedConsts: map[string]igop.UntypedConst{},
	})
}
