// export by github.com/goplus/igop/cmd/qexp

//go:build go1.20
// +build go1.20

package ecdh

import (
	q "crypto/ecdh"

	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "ecdh",
		Path: "crypto/ecdh",
		Deps: map[string]string{
			"crypto":                             "crypto",
			"crypto/internal/boring":             "boring",
			"crypto/internal/edwards25519/field": "field",
			"crypto/internal/nistec":             "nistec",
			"crypto/internal/randutil":           "randutil",
			"crypto/subtle":                      "subtle",
			"encoding/binary":                    "binary",
			"errors":                             "errors",
			"io":                                 "io",
			"math/bits":                          "bits",
			"sync":                               "sync",
		},
		Interfaces: map[string]reflect.Type{
			"Curve": reflect.TypeOf((*q.Curve)(nil)).Elem(),
		},
		NamedTypes: map[string]reflect.Type{
			"PrivateKey": reflect.TypeOf((*q.PrivateKey)(nil)).Elem(),
			"PublicKey":  reflect.TypeOf((*q.PublicKey)(nil)).Elem(),
		},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"P256":   reflect.ValueOf(q.P256),
			"P384":   reflect.ValueOf(q.P384),
			"P521":   reflect.ValueOf(q.P521),
			"X25519": reflect.ValueOf(q.X25519),
		},
		TypedConsts:   map[string]igop.TypedConst{},
		UntypedConsts: map[string]igop.UntypedConst{},
	})
}
