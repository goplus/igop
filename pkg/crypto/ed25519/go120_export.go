// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.20 && !go1.21
// +build go1.20,!go1.21

package ed25519

import (
	q "crypto/ed25519"

	"go/constant"
	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage(&ixgo.Package{
		Name: "ed25519",
		Path: "crypto/ed25519",
		Deps: map[string]string{
			"bytes":                        "bytes",
			"crypto":                       "crypto",
			"crypto/internal/edwards25519": "edwards25519",
			"crypto/rand":                  "rand",
			"crypto/sha512":                "sha512",
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
		TypedConsts: map[string]ixgo.TypedConst{},
		UntypedConsts: map[string]ixgo.UntypedConst{
			"PrivateKeySize": {"untyped int", constant.MakeInt64(int64(q.PrivateKeySize))},
			"PublicKeySize":  {"untyped int", constant.MakeInt64(int64(q.PublicKeySize))},
			"SeedSize":       {"untyped int", constant.MakeInt64(int64(q.SeedSize))},
			"SignatureSize":  {"untyped int", constant.MakeInt64(int64(q.SignatureSize))},
		},
	})
}
