// export by github.com/goplus/igop/cmd/qexp

//+build go1.16,!go1.17

package hash

import (
	q "hash"

	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "hash",
		Path: "hash",
		Deps: map[string]string{
			"io": "io",
		},
		Interfaces: map[string]reflect.Type{
			"Hash":   reflect.TypeOf((*q.Hash)(nil)).Elem(),
			"Hash32": reflect.TypeOf((*q.Hash32)(nil)).Elem(),
			"Hash64": reflect.TypeOf((*q.Hash64)(nil)).Elem(),
		},
		NamedTypes:    map[string]igop.NamedType{},
		AliasTypes:    map[string]reflect.Type{},
		Vars:          map[string]reflect.Value{},
		Funcs:         map[string]reflect.Value{},
		TypedConsts:   map[string]igop.TypedConst{},
		UntypedConsts: map[string]igop.UntypedConst{},
	})
}
