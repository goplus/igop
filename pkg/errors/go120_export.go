// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.20 && !go1.21
// +build go1.20,!go1.21

package errors

import (
	q "errors"

	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage(&ixgo.Package{
		Name: "errors",
		Path: "errors",
		Deps: map[string]string{
			"internal/reflectlite": "reflectlite",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"As":     reflect.ValueOf(q.As),
			"Is":     reflect.ValueOf(q.Is),
			"Join":   reflect.ValueOf(q.Join),
			"New":    reflect.ValueOf(q.New),
			"Unwrap": reflect.ValueOf(q.Unwrap),
		},
		TypedConsts:   map[string]ixgo.TypedConst{},
		UntypedConsts: map[string]ixgo.UntypedConst{},
	})
}
