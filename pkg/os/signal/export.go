// export by github.com/goplus/gossa/cmd/qexp

package signal

import (
	q "os/signal"

	"reflect"

	"github.com/goplus/gossa"
)

func init() {
	gossa.RegisterPackage(&gossa.Package{
		Name: "signal",
		Path: "os/signal",
		Deps: map[string]string{
			"context": "context",
			"os":      "os",
			"sync":    "sync",
			"syscall": "syscall",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]gossa.NamedType{},
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
		TypedConsts:   map[string]gossa.TypedConst{},
		UntypedConsts: map[string]gossa.UntypedConst{},
	})
}
