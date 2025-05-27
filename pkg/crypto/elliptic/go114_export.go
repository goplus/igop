// export by github.com/goplus/ixgo/cmd/qexp

//+build go1.14,!go1.15

package elliptic

import (
	q "crypto/elliptic"

	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage(&ixgo.Package{
		Name: "elliptic",
		Path: "crypto/elliptic",
		Deps: map[string]string{
			"io":       "io",
			"math/big": "big",
			"sync":     "sync",
		},
		Interfaces: map[string]reflect.Type{
			"Curve": reflect.TypeOf((*q.Curve)(nil)).Elem(),
		},
		NamedTypes: map[string]reflect.Type{
			"CurveParams": reflect.TypeOf((*q.CurveParams)(nil)).Elem(),
		},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"GenerateKey": reflect.ValueOf(q.GenerateKey),
			"Marshal":     reflect.ValueOf(q.Marshal),
			"P224":        reflect.ValueOf(q.P224),
			"P256":        reflect.ValueOf(q.P256),
			"P384":        reflect.ValueOf(q.P384),
			"P521":        reflect.ValueOf(q.P521),
			"Unmarshal":   reflect.ValueOf(q.Unmarshal),
		},
		TypedConsts:   map[string]ixgo.TypedConst{},
		UntypedConsts: map[string]ixgo.UntypedConst{},
	})
}
