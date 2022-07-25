// export by github.com/goplus/igop/cmd/qexp

//+build go1.14,!go1.15

package maphash

import (
	q "hash/maphash"

	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "maphash",
		Path: "hash/maphash",
		Deps: map[string]string{
			"unsafe": "unsafe",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{
			"Hash": reflect.TypeOf((*q.Hash)(nil)).Elem(),
			"Seed": reflect.TypeOf((*q.Seed)(nil)).Elem(),
		},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"MakeSeed": reflect.ValueOf(q.MakeSeed),
		},
		TypedConsts:   map[string]igop.TypedConst{},
		UntypedConsts: map[string]igop.UntypedConst{},
	})
}
