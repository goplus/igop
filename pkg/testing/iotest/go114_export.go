// export by github.com/goplus/igop/cmd/qexp

//+build go1.14,!go1.15

package iotest

import (
	q "testing/iotest"

	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "iotest",
		Path: "testing/iotest",
		Deps: map[string]string{
			"errors": "errors",
			"io":     "io",
			"log":    "log",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{},
		AliasTypes: map[string]reflect.Type{},
		Vars: map[string]reflect.Value{
			"ErrTimeout": reflect.ValueOf(&q.ErrTimeout),
		},
		Funcs: map[string]reflect.Value{
			"DataErrReader":  reflect.ValueOf(q.DataErrReader),
			"HalfReader":     reflect.ValueOf(q.HalfReader),
			"NewReadLogger":  reflect.ValueOf(q.NewReadLogger),
			"NewWriteLogger": reflect.ValueOf(q.NewWriteLogger),
			"OneByteReader":  reflect.ValueOf(q.OneByteReader),
			"TimeoutReader":  reflect.ValueOf(q.TimeoutReader),
			"TruncateWriter": reflect.ValueOf(q.TruncateWriter),
		},
		TypedConsts:   map[string]igop.TypedConst{},
		UntypedConsts: map[string]igop.UntypedConst{},
	})
}
