// export by github.com/goplus/ixgo/cmd/qexp

//+build go1.16,!go1.17

package context

import (
	q "context"

	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage(&ixgo.Package{
		Name: "context",
		Path: "context",
		Deps: map[string]string{
			"errors":               "errors",
			"internal/reflectlite": "reflectlite",
			"sync":                 "sync",
			"sync/atomic":          "atomic",
			"time":                 "time",
		},
		Interfaces: map[string]reflect.Type{
			"Context": reflect.TypeOf((*q.Context)(nil)).Elem(),
		},
		NamedTypes: map[string]reflect.Type{
			"CancelFunc": reflect.TypeOf((*q.CancelFunc)(nil)).Elem(),
		},
		AliasTypes: map[string]reflect.Type{},
		Vars: map[string]reflect.Value{
			"Canceled":         reflect.ValueOf(&q.Canceled),
			"DeadlineExceeded": reflect.ValueOf(&q.DeadlineExceeded),
		},
		Funcs: map[string]reflect.Value{
			"Background":   reflect.ValueOf(q.Background),
			"TODO":         reflect.ValueOf(q.TODO),
			"WithCancel":   reflect.ValueOf(q.WithCancel),
			"WithDeadline": reflect.ValueOf(q.WithDeadline),
			"WithTimeout":  reflect.ValueOf(q.WithTimeout),
			"WithValue":    reflect.ValueOf(q.WithValue),
		},
		TypedConsts:   map[string]ixgo.TypedConst{},
		UntypedConsts: map[string]ixgo.UntypedConst{},
	})
}
