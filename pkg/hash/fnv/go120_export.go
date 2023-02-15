// export by github.com/goplus/igop/cmd/qexp

//go:build go1.20
// +build go1.20

package fnv

import (
	q "hash/fnv"

	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "fnv",
		Path: "hash/fnv",
		Deps: map[string]string{
			"errors":    "errors",
			"hash":      "hash",
			"math/bits": "bits",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"New128":  reflect.ValueOf(q.New128),
			"New128a": reflect.ValueOf(q.New128a),
			"New32":   reflect.ValueOf(q.New32),
			"New32a":  reflect.ValueOf(q.New32a),
			"New64":   reflect.ValueOf(q.New64),
			"New64a":  reflect.ValueOf(q.New64a),
		},
		TypedConsts:   map[string]igop.TypedConst{},
		UntypedConsts: map[string]igop.UntypedConst{},
	})
}
