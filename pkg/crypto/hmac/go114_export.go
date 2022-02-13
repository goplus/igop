// export by github.com/goplus/gossa/cmd/qexp

//+build go1.14,!go1.15

package hmac

import (
	q "crypto/hmac"

	"reflect"

	"github.com/goplus/gossa"
)

func init() {
	gossa.RegisterPackage(&gossa.Package{
		Name: "hmac",
		Path: "crypto/hmac",
		Deps: map[string]string{
			"crypto/subtle": "subtle",
			"hash":          "hash",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]gossa.NamedType{},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"Equal": reflect.ValueOf(q.Equal),
			"New":   reflect.ValueOf(q.New),
		},
		TypedConsts:   map[string]gossa.TypedConst{},
		UntypedConsts: map[string]gossa.UntypedConst{},
	})
}
