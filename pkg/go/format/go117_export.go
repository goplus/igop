// export by github.com/goplus/gossa/cmd/qexp

//go:build go1.17
// +build go1.17

package format

import (
	q "go/format"

	"reflect"

	"github.com/goplus/gossa"
)

func init() {
	gossa.RegisterPackage(&gossa.Package{
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
		NamedTypes: map[string]gossa.NamedType{},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"Node":   reflect.ValueOf(q.Node),
			"Source": reflect.ValueOf(q.Source),
		},
		TypedConsts:   map[string]gossa.TypedConst{},
		UntypedConsts: map[string]gossa.UntypedConst{},
	})
}
