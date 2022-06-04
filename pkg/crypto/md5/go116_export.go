// export by github.com/goplus/igop/cmd/qexp

//+build go1.16,!go1.17

package md5

import (
	q "crypto/md5"

	"go/constant"
	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
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
		NamedTypes: map[string]igop.NamedType{},
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
