// export by github.com/goplus/gossa/cmd/qexp

//+build go1.15,!go1.16

package bzip2

import (
	q "compress/bzip2"

	"reflect"

	"github.com/goplus/gossa"
)

func init() {
	gossa.RegisterPackage(&gossa.Package{
		Name: "bzip2",
		Path: "compress/bzip2",
		Deps: map[string]string{
			"bufio": "bufio",
			"io":    "io",
			"sort":  "sort",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]gossa.NamedType{
			"StructuralError": {reflect.TypeOf((*q.StructuralError)(nil)).Elem(), "Error", ""},
		},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"NewReader": reflect.ValueOf(q.NewReader),
		},
		TypedConsts:   map[string]gossa.TypedConst{},
		UntypedConsts: map[string]gossa.UntypedConst{},
	})
}
