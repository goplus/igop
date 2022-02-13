// export by github.com/goplus/gossa/cmd/qexp

//+build go1.15,!go1.16

package pem

import (
	q "encoding/pem"

	"reflect"

	"github.com/goplus/gossa"
)

func init() {
	gossa.RegisterPackage(&gossa.Package{
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
		NamedTypes: map[string]gossa.NamedType{
			"Block": {reflect.TypeOf((*q.Block)(nil)).Elem(), "", ""},
		},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"Decode":         reflect.ValueOf(q.Decode),
			"Encode":         reflect.ValueOf(q.Encode),
			"EncodeToMemory": reflect.ValueOf(q.EncodeToMemory),
		},
		TypedConsts:   map[string]gossa.TypedConst{},
		UntypedConsts: map[string]gossa.UntypedConst{},
	})
}
