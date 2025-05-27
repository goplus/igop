// export by github.com/goplus/ixgo/cmd/qexp

//+build go1.15,!go1.16

package hmac

import (
	q "crypto/hmac"

	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage(&ixgo.Package{
		Name: "hmac",
		Path: "crypto/hmac",
		Deps: map[string]string{
			"crypto/subtle": "subtle",
			"hash":          "hash",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"Equal": reflect.ValueOf(q.Equal),
			"New":   reflect.ValueOf(q.New),
		},
		TypedConsts:   map[string]ixgo.TypedConst{},
		UntypedConsts: map[string]ixgo.UntypedConst{},
	})
}
