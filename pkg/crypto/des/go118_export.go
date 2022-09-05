// export by github.com/goplus/igop/cmd/qexp

//go:build go1.18 && !go1.19
// +build go1.18,!go1.19

package des

import (
	q "crypto/des"

	"go/constant"
	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "des",
		Path: "crypto/des",
		Deps: map[string]string{
			"crypto/cipher":          "cipher",
			"crypto/internal/subtle": "subtle",
			"encoding/binary":        "binary",
			"strconv":                "strconv",
			"sync":                   "sync",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{
			"KeySizeError": reflect.TypeOf((*q.KeySizeError)(nil)).Elem(),
		},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"NewCipher":          reflect.ValueOf(q.NewCipher),
			"NewTripleDESCipher": reflect.ValueOf(q.NewTripleDESCipher),
		},
		TypedConsts: map[string]igop.TypedConst{},
		UntypedConsts: map[string]igop.UntypedConst{
			"BlockSize": {"untyped int", constant.MakeInt64(int64(q.BlockSize))},
		},
	})
}
