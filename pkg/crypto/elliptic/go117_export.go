// export by github.com/goplus/gossa/cmd/qexp

//go:build go1.17 && !go1.18
// +build go1.17,!go1.18

package elliptic

import (
	q "crypto/elliptic"

	"reflect"

	"github.com/goplus/gossa"
)

func init() {
	gossa.RegisterPackage(&gossa.Package{
		Name: "elliptic",
		Path: "crypto/elliptic",
		Deps: map[string]string{
			"crypto/elliptic/internal/fiat": "fiat",
			"io":                            "io",
			"math/big":                      "big",
			"sync":                          "sync",
		},
		Interfaces: map[string]reflect.Type{
			"Curve": reflect.TypeOf((*q.Curve)(nil)).Elem(),
		},
		NamedTypes: map[string]gossa.NamedType{
			"CurveParams": {reflect.TypeOf((*q.CurveParams)(nil)).Elem(), "", "Add,Double,IsOnCurve,Params,ScalarBaseMult,ScalarMult,addJacobian,affineFromJacobian,doubleJacobian,polynomial"},
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
		TypedConsts:   map[string]gossa.TypedConst{},
		UntypedConsts: map[string]gossa.UntypedConst{},
	})
}
