// export by github.com/goplus/igop/cmd/qexp

//go:build go1.19
// +build go1.19

package format

import (
	q "go/format"

	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "format",
		Path: "go/format",
		Deps: map[string]string{
			"bytes":      "bytes",
			"fmt":        "fmt",
			"go/ast":     "ast",
			"go/parser":  "parser",
			"go/printer": "printer",
			"go/token":   "token",
			"io":         "io",
			"strings":    "strings",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"Node":   reflect.ValueOf(q.Node),
			"Source": reflect.ValueOf(q.Source),
		},
		TypedConsts:   map[string]igop.TypedConst{},
		UntypedConsts: map[string]igop.UntypedConst{},
	})
}
