// export by github.com/goplus/igop/cmd/qexp

//go:build go1.23 && !go1.24
// +build go1.23,!go1.24

package elliptic

import (
	q "crypto/elliptic"

	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "elliptic",
		Path: "crypto/elliptic",
		Deps: map[string]string{
			"crypto/internal/nistec": "nistec",
			"errors":                 "errors",
			"io":                     "io",
			"math/big":               "big",
			"sync":                   "sync",
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
			"GenerateKey":         reflect.ValueOf(q.GenerateKey),
			"Marshal":             reflect.ValueOf(q.Marshal),
			"MarshalCompressed":   reflect.ValueOf(q.MarshalCompressed),
			"P224":                reflect.ValueOf(q.P224),
			"P256":                reflect.ValueOf(q.P256),
			"P384":                reflect.ValueOf(q.P384),
			"P521":                reflect.ValueOf(q.P521),
			"Unmarshal":           reflect.ValueOf(q.Unmarshal),
			"UnmarshalCompressed": reflect.ValueOf(q.UnmarshalCompressed),
		},
		TypedConsts:   map[string]igop.TypedConst{},
		UntypedConsts: map[string]igop.UntypedConst{},
	})
}
