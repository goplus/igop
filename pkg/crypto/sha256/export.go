// export by github.com/goplus/gossa/cmd/qexp

package sha256

import (
	q "crypto/sha256"

	"go/constant"
	"reflect"

	"github.com/goplus/gossa"
)

func init() {
	gossa.RegisterPackage(&gossa.Package{
		Name: "sha256",
		Path: "crypto/sha256",
		Deps: map[string]string{
			"crypto":          "crypto",
			"encoding/binary": "binary",
			"errors":          "errors",
			"hash":            "hash",
			"internal/cpu":    "cpu",
			"math/bits":       "bits",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]gossa.NamedType{},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"New":    reflect.ValueOf(q.New),
			"New224": reflect.ValueOf(q.New224),
			"Sum224": reflect.ValueOf(q.Sum224),
			"Sum256": reflect.ValueOf(q.Sum256),
		},
		TypedConsts: map[string]gossa.TypedConst{},
		UntypedConsts: map[string]gossa.UntypedConst{
			"BlockSize": {"untyped int", constant.MakeInt64(int64(q.BlockSize))},
			"Size":      {"untyped int", constant.MakeInt64(int64(q.Size))},
			"Size224":   {"untyped int", constant.MakeInt64(int64(q.Size224))},
		},
	})
}
