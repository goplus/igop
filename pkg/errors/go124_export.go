// export by github.com/goplus/igop/cmd/qexp

//go:build go1.24
// +build go1.24

package errors

import (
	q "errors"

	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "errors",
		Path: "errors",
		Deps: map[string]string{
			"internal/reflectlite": "reflectlite",
			"unsafe":               "unsafe",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{},
		AliasTypes: map[string]reflect.Type{},
		Vars: map[string]reflect.Value{
			"ErrUnsupported": reflect.ValueOf(&q.ErrUnsupported),
		},
		Funcs: map[string]reflect.Value{
			"As":     reflect.ValueOf(q.As),
			"Is":     reflect.ValueOf(q.Is),
			"Join":   reflect.ValueOf(q.Join),
			"New":    reflect.ValueOf(q.New),
			"Unwrap": reflect.ValueOf(q.Unwrap),
		},
		TypedConsts:   map[string]igop.TypedConst{},
		UntypedConsts: map[string]igop.UntypedConst{},
	})
}
