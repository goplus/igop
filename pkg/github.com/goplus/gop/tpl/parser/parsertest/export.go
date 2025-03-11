// export by github.com/goplus/igop/cmd/qexp

package parsertest

import (
	q "github.com/goplus/gop/tpl/parser/parsertest"

	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "parsertest",
		Path: "github.com/goplus/gop/tpl/parser/parsertest",
		Deps: map[string]string{
			"bytes":                           "bytes",
			"fmt":                             "fmt",
			"github.com/goplus/gop/tpl/ast":   "ast",
			"github.com/goplus/gop/tpl/token": "token",
			"io":                              "io",
			"log":                             "log",
			"os":                              "os",
			"reflect":                         "reflect",
			"testing":                         "testing",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"Expect":     reflect.ValueOf(q.Expect),
			"Fprint":     reflect.ValueOf(q.Fprint),
			"FprintNode": reflect.ValueOf(q.FprintNode),
		},
		TypedConsts:   map[string]igop.TypedConst{},
		UntypedConsts: map[string]igop.UntypedConst{},
	})
}
