// export by github.com/goplus/igop/cmd/qexp

//+build go1.15,!go1.16

package rand

import (
	q "crypto/rand"

	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "rand",
		Path: "crypto/rand",
		Deps: map[string]string{
			"bufio":           "bufio",
			"crypto/aes":      "aes",
			"crypto/cipher":   "cipher",
			"encoding/binary": "binary",
			"errors":          "errors",
			"io":              "io",
			"math/big":        "big",
			"os":              "os",
			"runtime":         "runtime",
			"sync":            "sync",
			"sync/atomic":     "atomic",
			"syscall":         "syscall",
			"time":            "time",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]igop.NamedType{},
		AliasTypes: map[string]reflect.Type{},
		Vars: map[string]reflect.Value{
			"Reader": reflect.ValueOf(&q.Reader),
		},
		Funcs: map[string]reflect.Value{
			"Int":   reflect.ValueOf(q.Int),
			"Prime": reflect.ValueOf(q.Prime),
			"Read":  reflect.ValueOf(q.Read),
		},
		TypedConsts:   map[string]igop.TypedConst{},
		UntypedConsts: map[string]igop.UntypedConst{},
	})
}
