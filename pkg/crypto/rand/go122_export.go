// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.22 && !go1.23
// +build go1.22,!go1.23

package rand

import (
	q "crypto/rand"

	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage(&ixgo.Package{
		Name: "rand",
		Path: "crypto/rand",
		Deps: map[string]string{
			"crypto/internal/boring":   "boring",
			"crypto/internal/randutil": "randutil",
			"errors":                   "errors",
			"internal/syscall/unix":    "unix",
			"io":                       "io",
			"math/big":                 "big",
			"os":                       "os",
			"sync":                     "sync",
			"sync/atomic":              "atomic",
			"syscall":                  "syscall",
			"time":                     "time",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{},
		AliasTypes: map[string]reflect.Type{},
		Vars: map[string]reflect.Value{
			"Reader": reflect.ValueOf(&q.Reader),
		},
		Funcs: map[string]reflect.Value{
			"Int":   reflect.ValueOf(q.Int),
			"Prime": reflect.ValueOf(q.Prime),
			"Read":  reflect.ValueOf(q.Read),
		},
		TypedConsts:   map[string]ixgo.TypedConst{},
		UntypedConsts: map[string]ixgo.UntypedConst{},
	})
}
