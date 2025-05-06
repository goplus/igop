// export by github.com/goplus/igop/cmd/qexp

//go:build go1.24
// +build go1.24

package sha1

import (
	q "crypto/sha1"

	"go/constant"
	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "sha1",
		Path: "crypto/sha1",
		Deps: map[string]string{
			"crypto":                      "crypto",
			"crypto/internal/boring":      "boring",
			"crypto/internal/fips140only": "fips140only",
			"errors":                      "errors",
			"hash":                        "hash",
			"internal/byteorder":          "byteorder",
			"internal/cpu":                "cpu",
			"math/bits":                   "bits",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"New": reflect.ValueOf(q.New),
			"Sum": reflect.ValueOf(q.Sum),
		},
		TypedConsts: map[string]igop.TypedConst{},
		UntypedConsts: map[string]igop.UntypedConst{
			"BlockSize": {"untyped int", constant.MakeInt64(int64(q.BlockSize))},
			"Size":      {"untyped int", constant.MakeInt64(int64(q.Size))},
		},
	})
}
