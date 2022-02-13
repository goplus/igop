// export by github.com/goplus/gossa/cmd/qexp

//+build go1.16,!go1.17

package ed25519

import (
	q "crypto/ed25519"

	"go/constant"
	"reflect"

	"github.com/goplus/gossa"
)

func init() {
	gossa.RegisterPackage(&gossa.Package{
		Name: "ed25519",
		Path: "crypto/ed25519",
		Deps: map[string]string{
			"bytes":                                "bytes",
			"crypto":                               "crypto",
			"crypto/ed25519/internal/edwards25519": "edwards25519",
			"crypto/rand":                          "rand",
			"crypto/sha512":                        "sha512",
			"errors":                               "errors",
			"io":                                   "io",
			"strconv":                              "strconv",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]gossa.NamedType{
			"PrivateKey": {reflect.TypeOf((*q.PrivateKey)(nil)).Elem(), "Equal,Public,Seed,Sign", ""},
			"PublicKey":  {reflect.TypeOf((*q.PublicKey)(nil)).Elem(), "Equal", ""},
		},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"GenerateKey":    reflect.ValueOf(q.GenerateKey),
			"NewKeyFromSeed": reflect.ValueOf(q.NewKeyFromSeed),
			"Sign":           reflect.ValueOf(q.Sign),
			"Verify":         reflect.ValueOf(q.Verify),
		},
		TypedConsts: map[string]gossa.TypedConst{},
		UntypedConsts: map[string]gossa.UntypedConst{
			"PrivateKeySize": {"untyped int", constant.MakeInt64(int64(q.PrivateKeySize))},
			"PublicKeySize":  {"untyped int", constant.MakeInt64(int64(q.PublicKeySize))},
			"SeedSize":       {"untyped int", constant.MakeInt64(int64(q.SeedSize))},
			"SignatureSize":  {"untyped int", constant.MakeInt64(int64(q.SignatureSize))},
		},
	})
}
