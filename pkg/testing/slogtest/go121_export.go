// export by github.com/goplus/igop/cmd/qexp

//go:build go1.21
// +build go1.21

package slogtest

import (
	q "testing/slogtest"

	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "slogtest",
		Path: "testing/slogtest",
		Deps: map[string]string{
			"context":  "context",
			"errors":   "errors",
			"fmt":      "fmt",
			"log/slog": "slog",
			"reflect":  "reflect",
			"runtime":  "runtime",
			"time":     "time",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"TestHandler": reflect.ValueOf(q.TestHandler),
		},
		TypedConsts:   map[string]igop.TypedConst{},
		UntypedConsts: map[string]igop.UntypedConst{},
	})
}
