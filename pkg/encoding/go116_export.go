// export by github.com/goplus/igop/cmd/qexp

//+build go1.16,!go1.17

package encoding

import (
	q "encoding"

	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "encoding",
		Path: "encoding",
		Deps: map[string]string{},
		Interfaces: map[string]reflect.Type{
			"BinaryMarshaler":   reflect.TypeOf((*q.BinaryMarshaler)(nil)).Elem(),
			"BinaryUnmarshaler": reflect.TypeOf((*q.BinaryUnmarshaler)(nil)).Elem(),
			"TextMarshaler":     reflect.TypeOf((*q.TextMarshaler)(nil)).Elem(),
			"TextUnmarshaler":   reflect.TypeOf((*q.TextUnmarshaler)(nil)).Elem(),
		},
		NamedTypes:    map[string]igop.NamedType{},
		AliasTypes:    map[string]reflect.Type{},
		Vars:          map[string]reflect.Value{},
		Funcs:         map[string]reflect.Value{},
		TypedConsts:   map[string]igop.TypedConst{},
		UntypedConsts: map[string]igop.UntypedConst{},
	})
}
