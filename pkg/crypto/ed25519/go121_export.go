// export by github.com/goplus/igop/cmd/qexp

//go:build go1.21 && !go1.22
// +build go1.21,!go1.22

package ed25519

import (
	q "crypto/ed25519"

	"go/constant"
	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "ed25519",
		Path: "crypto/ed25519",
		Deps: map[string]string{
			"bytes":                        "bytes",
			"crypto":                       "crypto",
			"crypto/internal/edwards25519": "edwards25519",
			"crypto/rand":                  "rand",
			"crypto/sha512":                "sha512",
			"crypto/subtle":                "subtle",
			"errors":                       "errors",
			"io":                           "io",
			"strconv":                      "strconv",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{
			"Options":    reflect.TypeOf((*q.Options)(nil)).Elem(),
			"PrivateKey": reflect.TypeOf((*q.PrivateKey)(nil)).Elem(),
			"PublicKey":  reflect.TypeOf((*q.PublicKey)(nil)).Elem(),
		},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"GenerateKey":       reflect.ValueOf(q.GenerateKey),
			"NewKeyFromSeed":    reflect.ValueOf(q.NewKeyFromSeed),
			"Sign":              reflect.ValueOf(q.Sign),
			"Verify":            reflect.ValueOf(q.Verify),
			"VerifyWithOptions": reflect.ValueOf(q.VerifyWithOptions),
		},
		TypedConsts: map[string]igop.TypedConst{},
		UntypedConsts: map[string]igop.UntypedConst{
			"PrivateKeySize": {"untyped int", constant.MakeInt64(int64(q.PrivateKeySize))},
			"PublicKeySize":  {"untyped int", constant.MakeInt64(int64(q.PublicKeySize))},
			"SeedSize":       {"untyped int", constant.MakeInt64(int64(q.SeedSize))},
			"SignatureSize":  {"untyped int", constant.MakeInt64(int64(q.SignatureSize))},
		},
	})
}
