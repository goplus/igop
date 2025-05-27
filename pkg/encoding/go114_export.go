// export by github.com/goplus/ixgo/cmd/qexp

//+build go1.14,!go1.15

package encoding

import (
	q "encoding"

	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage(&ixgo.Package{
		Name: "encoding",
		Path: "encoding",
		Deps: map[string]string{},
		Interfaces: map[string]reflect.Type{
			"BinaryMarshaler":   reflect.TypeOf((*q.BinaryMarshaler)(nil)).Elem(),
			"BinaryUnmarshaler": reflect.TypeOf((*q.BinaryUnmarshaler)(nil)).Elem(),
			"TextMarshaler":     reflect.TypeOf((*q.TextMarshaler)(nil)).Elem(),
			"TextUnmarshaler":   reflect.TypeOf((*q.TextUnmarshaler)(nil)).Elem(),
		},
		NamedTypes:    map[string]reflect.Type{},
		AliasTypes:    map[string]reflect.Type{},
		Vars:          map[string]reflect.Value{},
		Funcs:         map[string]reflect.Value{},
		TypedConsts:   map[string]ixgo.TypedConst{},
		UntypedConsts: map[string]ixgo.UntypedConst{},
	})
}
