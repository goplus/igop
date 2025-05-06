// export by github.com/goplus/igop/cmd/qexp

//go:build go1.24
// +build go1.24

package signal

import (
	q "os/signal"

	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "signal",
		Path: "os/signal",
		Deps: map[string]string{
			"context": "context",
			"os":      "os",
			"slices":  "slices",
			"sync":    "sync",
			"syscall": "syscall",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"Ignore":        reflect.ValueOf(q.Ignore),
			"Ignored":       reflect.ValueOf(q.Ignored),
			"Notify":        reflect.ValueOf(q.Notify),
			"NotifyContext": reflect.ValueOf(q.NotifyContext),
			"Reset":         reflect.ValueOf(q.Reset),
			"Stop":          reflect.ValueOf(q.Stop),
		},
		TypedConsts:   map[string]igop.TypedConst{},
		UntypedConsts: map[string]igop.UntypedConst{},
	})
}
