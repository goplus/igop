// export by github.com/goplus/igop/cmd/qexp

//+build go1.14,!go1.15

package ascii85

import (
	q "encoding/ascii85"

	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "ascii85",
		Path: "encoding/ascii85",
		Deps: map[string]string{
			"io":      "io",
			"strconv": "strconv",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{
			"CorruptInputError": reflect.TypeOf((*q.CorruptInputError)(nil)).Elem(),
		},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"Decode":        reflect.ValueOf(q.Decode),
			"Encode":        reflect.ValueOf(q.Encode),
			"MaxEncodedLen": reflect.ValueOf(q.MaxEncodedLen),
			"NewDecoder":    reflect.ValueOf(q.NewDecoder),
			"NewEncoder":    reflect.ValueOf(q.NewEncoder),
		},
		TypedConsts:   map[string]igop.TypedConst{},
		UntypedConsts: map[string]igop.UntypedConst{},
	})
}
