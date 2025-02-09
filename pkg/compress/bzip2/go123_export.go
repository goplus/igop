// export by github.com/goplus/igop/cmd/qexp

//go:build go1.23
// +build go1.23

package bzip2

import (
	q "compress/bzip2"

	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "bzip2",
		Path: "compress/bzip2",
		Deps: map[string]string{
			"bufio":  "bufio",
			"cmp":    "cmp",
			"io":     "io",
			"slices": "slices",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{
			"StructuralError": reflect.TypeOf((*q.StructuralError)(nil)).Elem(),
		},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"NewReader": reflect.ValueOf(q.NewReader),
		},
		TypedConsts:   map[string]igop.TypedConst{},
		UntypedConsts: map[string]igop.UntypedConst{},
	})
}
