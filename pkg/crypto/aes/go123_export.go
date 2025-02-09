// export by github.com/goplus/igop/cmd/qexp

//go:build go1.23
// +build go1.23

package aes

import (
	q "crypto/aes"

	"go/constant"
	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "aes",
		Path: "crypto/aes",
		Deps: map[string]string{
			"crypto/cipher":          "cipher",
			"crypto/internal/alias":  "alias",
			"crypto/internal/boring": "boring",
			"crypto/subtle":          "subtle",
			"errors":                 "errors",
			"internal/byteorder":     "byteorder",
			"internal/cpu":           "cpu",
			"internal/goarch":        "goarch",
			"strconv":                "strconv",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{
			"KeySizeError": reflect.TypeOf((*q.KeySizeError)(nil)).Elem(),
		},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"NewCipher": reflect.ValueOf(q.NewCipher),
		},
		TypedConsts: map[string]igop.TypedConst{},
		UntypedConsts: map[string]igop.UntypedConst{
			"BlockSize": {"untyped int", constant.MakeInt64(int64(q.BlockSize))},
		},
	})
}
