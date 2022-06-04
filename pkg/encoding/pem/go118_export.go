// export by github.com/goplus/igop/cmd/qexp

//go:build go1.18
// +build go1.18

package pem

import (
	q "encoding/pem"

	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "pem",
		Path: "encoding/pem",
		Deps: map[string]string{
			"bytes":           "bytes",
			"encoding/base64": "base64",
			"errors":          "errors",
			"io":              "io",
			"sort":            "sort",
			"strings":         "strings",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]igop.NamedType{
			"Block": {reflect.TypeOf((*q.Block)(nil)).Elem(), "", ""},
		},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"Decode":         reflect.ValueOf(q.Decode),
			"Encode":         reflect.ValueOf(q.Encode),
			"EncodeToMemory": reflect.ValueOf(q.EncodeToMemory),
		},
		TypedConsts:   map[string]igop.TypedConst{},
		UntypedConsts: map[string]igop.UntypedConst{},
	})
}
