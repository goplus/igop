// export by github.com/goplus/igop/cmd/qexp

//go:build go1.20 && !go1.20
// +build go1.20,!go1.20

package context

import (
	q "context"

	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
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
			"CancelCauseFunc": reflect.TypeOf((*q.CancelCauseFunc)(nil)).Elem(),
			"CancelFunc":      reflect.TypeOf((*q.CancelFunc)(nil)).Elem(),
		},
		AliasTypes: map[string]reflect.Type{},
		Vars: map[string]reflect.Value{
			"Canceled":         reflect.ValueOf(&q.Canceled),
			"DeadlineExceeded": reflect.ValueOf(&q.DeadlineExceeded),
		},
		Funcs: map[string]reflect.Value{
			"Background":      reflect.ValueOf(q.Background),
			"Cause":           reflect.ValueOf(q.Cause),
			"TODO":            reflect.ValueOf(q.TODO),
			"WithCancel":      reflect.ValueOf(q.WithCancel),
			"WithCancelCause": reflect.ValueOf(q.WithCancelCause),
			"WithDeadline":    reflect.ValueOf(q.WithDeadline),
			"WithTimeout":     reflect.ValueOf(q.WithTimeout),
			"WithValue":       reflect.ValueOf(q.WithValue),
		},
		TypedConsts:   map[string]igop.TypedConst{},
		UntypedConsts: map[string]igop.UntypedConst{},
	})
}
