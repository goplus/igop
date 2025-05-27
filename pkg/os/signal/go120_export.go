// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.20 && !go1.21
// +build go1.20,!go1.21

package signal

import (
	q "os/signal"

	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage(&ixgo.Package{
		Name: "signal",
		Path: "os/signal",
		Deps: map[string]string{
			"context": "context",
			"os":      "os",
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
		TypedConsts:   map[string]ixgo.TypedConst{},
		UntypedConsts: map[string]ixgo.UntypedConst{},
	})
}
