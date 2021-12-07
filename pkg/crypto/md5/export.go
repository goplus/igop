// export by github.com/goplus/gossa/cmd/qexp

package md5

import (
	q "crypto/md5"

	"go/constant"
	"reflect"

	"github.com/goplus/gossa"
)

func init() {
	gossa.RegisterPackage(&gossa.Package{
		Name: "md5",
		Path: "crypto/md5",
		Deps: map[string]string{
			"crypto":          "crypto",
			"encoding/binary": "binary",
			"errors":          "errors",
			"hash":            "hash",
			"math/bits":       "bits",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]gossa.NamedType{},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"New": reflect.ValueOf(q.New),
			"Sum": reflect.ValueOf(q.Sum),
		},
		TypedConsts: map[string]gossa.TypedConst{},
		UntypedConsts: map[string]gossa.UntypedConst{
			"BlockSize": {"untyped int", constant.MakeInt64(int64(q.BlockSize))},
			"Size":      {"untyped int", constant.MakeInt64(int64(q.Size))},
		},
	})
}
