// export by github.com/goplus/ixgo/cmd/qexp

//+build go1.14,!go1.15

package fnv

import (
	q "hash/fnv"

	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage(&ixgo.Package{
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
		TypedConsts:   map[string]ixgo.TypedConst{},
		UntypedConsts: map[string]ixgo.UntypedConst{},
	})
}
