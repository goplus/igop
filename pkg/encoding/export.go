// export by github.com/goplus/gossa/cmd/qexp

package encoding

import (
	q "encoding"

	"reflect"

	"github.com/goplus/gossa"
)

func init() {
	gossa.RegisterPackage(&gossa.Package{
		Name: "encoding",
		Path: "encoding",
		Deps: map[string]string{},
		Interfaces: map[string]reflect.Type{
			"BinaryMarshaler":   reflect.TypeOf((*q.BinaryMarshaler)(nil)).Elem(),
			"BinaryUnmarshaler": reflect.TypeOf((*q.BinaryUnmarshaler)(nil)).Elem(),
			"TextMarshaler":     reflect.TypeOf((*q.TextMarshaler)(nil)).Elem(),
			"TextUnmarshaler":   reflect.TypeOf((*q.TextUnmarshaler)(nil)).Elem(),
		},
		NamedTypes:    map[string]gossa.NamedType{},
		AliasTypes:    map[string]reflect.Type{},
		Vars:          map[string]reflect.Value{},
		Funcs:         map[string]reflect.Value{},
		TypedConsts:   map[string]gossa.TypedConst{},
		UntypedConsts: map[string]gossa.UntypedConst{},
	})
}
