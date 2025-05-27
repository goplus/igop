// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.17 && !go1.18
// +build go1.17,!go1.18

package md5

import (
	q "crypto/md5"

	"go/constant"
	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage(&ixgo.Package{
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
		NamedTypes: map[string]reflect.Type{},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"New": reflect.ValueOf(q.New),
			"Sum": reflect.ValueOf(q.Sum),
		},
		TypedConsts: map[string]ixgo.TypedConst{},
		UntypedConsts: map[string]ixgo.UntypedConst{
			"BlockSize": {"untyped int", constant.MakeInt64(int64(q.BlockSize))},
			"Size":      {"untyped int", constant.MakeInt64(int64(q.Size))},
		},
	})
}
