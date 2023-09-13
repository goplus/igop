// export by github.com/goplus/igop/cmd/qexp

//go:build go1.20 && !go1.21
// +build go1.20,!go1.21

package ecdsa

import (
	q "crypto/ecdsa"

	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "ecdsa",
		Path: "crypto/ecdsa",
		Deps: map[string]string{
			"bytes":                                 "bytes",
			"crypto":                                "crypto",
			"crypto/aes":                            "aes",
			"crypto/cipher":                         "cipher",
			"crypto/ecdh":                           "ecdh",
			"crypto/elliptic":                       "elliptic",
			"crypto/internal/bigmod":                "bigmod",
			"crypto/internal/boring":                "boring",
			"crypto/internal/boring/bbig":           "bbig",
			"crypto/internal/nistec":                "nistec",
			"crypto/internal/randutil":              "randutil",
			"crypto/sha512":                         "sha512",
			"crypto/subtle":                         "subtle",
			"errors":                                "errors",
			"io":                                    "io",
			"math/big":                              "big",
			"sync":                                  "sync",
			"vendor/golang.org/x/crypto/cryptobyte": "cryptobyte",
			"vendor/golang.org/x/crypto/cryptobyte/asn1": "asn1",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{
			"PrivateKey": reflect.TypeOf((*q.PrivateKey)(nil)).Elem(),
			"PublicKey":  reflect.TypeOf((*q.PublicKey)(nil)).Elem(),
		},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"GenerateKey": reflect.ValueOf(q.GenerateKey),
			"Sign":        reflect.ValueOf(q.Sign),
			"SignASN1":    reflect.ValueOf(q.SignASN1),
			"Verify":      reflect.ValueOf(q.Verify),
			"VerifyASN1":  reflect.ValueOf(q.VerifyASN1),
		},
		TypedConsts:   map[string]igop.TypedConst{},
		UntypedConsts: map[string]igop.UntypedConst{},
	})
}
