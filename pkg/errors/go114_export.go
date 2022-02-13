// export by github.com/goplus/gossa/cmd/qexp

//+build go1.14,!go1.15

package errors

import (
	q "errors"

	"reflect"

	"github.com/goplus/gossa"
)

func init() {
	gossa.RegisterPackage(&gossa.Package{
		Name: "errors",
		Path: "errors",
		Deps: map[string]string{
			"internal/reflectlite": "reflectlite",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]gossa.NamedType{},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"As":     reflect.ValueOf(q.As),
			"Is":     reflect.ValueOf(q.Is),
			"New":    reflect.ValueOf(q.New),
			"Unwrap": reflect.ValueOf(q.Unwrap),
		},
		TypedConsts:   map[string]gossa.TypedConst{},
		UntypedConsts: map[string]gossa.UntypedConst{},
	})
}
