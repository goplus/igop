// export by github.com/goplus/igop/cmd/qexp

//+build go1.16,!go1.17

package utf16

import (
	q "unicode/utf16"

	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name:       "utf16",
		Path:       "unicode/utf16",
		Deps:       map[string]string{},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"Decode":      reflect.ValueOf(q.Decode),
			"DecodeRune":  reflect.ValueOf(q.DecodeRune),
			"Encode":      reflect.ValueOf(q.Encode),
			"EncodeRune":  reflect.ValueOf(q.EncodeRune),
			"IsSurrogate": reflect.ValueOf(q.IsSurrogate),
		},
		TypedConsts:   map[string]igop.TypedConst{},
		UntypedConsts: map[string]igop.UntypedConst{},
	})
}
