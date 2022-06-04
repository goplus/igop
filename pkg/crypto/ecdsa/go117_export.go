// export by github.com/goplus/igop/cmd/qexp

//go:build go1.17 && !go1.18
// +build go1.17,!go1.18

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
			"crypto":                                "crypto",
			"crypto/aes":                            "aes",
			"crypto/cipher":                         "cipher",
			"crypto/elliptic":                       "elliptic",
			"crypto/internal/randutil":              "randutil",
			"crypto/sha512":                         "sha512",
			"errors":                                "errors",
			"io":                                    "io",
			"math/big":                              "big",
			"vendor/golang.org/x/crypto/cryptobyte": "cryptobyte",
			"vendor/golang.org/x/crypto/cryptobyte/asn1": "asn1",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]igop.NamedType{
			"PrivateKey": {reflect.TypeOf((*q.PrivateKey)(nil)).Elem(), "", "Equal,Public,Sign"},
			"PublicKey":  {reflect.TypeOf((*q.PublicKey)(nil)).Elem(), "", "Equal"},
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
