// export by github.com/goplus/gossa/cmd/qexp

package ascii85

import (
	q "encoding/ascii85"

	"reflect"

	"github.com/goplus/gossa"
)

func init() {
	gossa.RegisterPackage(&gossa.Package{
		Name: "ascii85",
		Path: "encoding/ascii85",
		Deps: map[string]string{
			"io":      "io",
			"strconv": "strconv",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]gossa.NamedType{
			"CorruptInputError": {reflect.TypeOf((*q.CorruptInputError)(nil)).Elem(), "Error", ""},
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
		TypedConsts:   map[string]gossa.TypedConst{},
		UntypedConsts: map[string]gossa.UntypedConst{},
	})
}
