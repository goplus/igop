// export by github.com/goplus/igop/cmd/qexp

//+build go1.15,!go1.16

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
		NamedTypes: map[string]igop.NamedType{
			"Hash": {reflect.TypeOf((*q.Hash)(nil)).Elem(), "", "BlockSize,Reset,Seed,SetSeed,Size,Sum,Sum64,Write,WriteByte,WriteString,flush,initSeed,setSeed"},
			"Seed": {reflect.TypeOf((*q.Seed)(nil)).Elem(), "", ""},
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
