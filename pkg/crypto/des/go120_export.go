// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.20 && !go1.21
// +build go1.20,!go1.21

package des

import (
	q "crypto/des"

	"go/constant"
	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage(&ixgo.Package{
		Name: "des",
		Path: "crypto/des",
		Deps: map[string]string{
			"crypto/cipher":         "cipher",
			"crypto/internal/alias": "alias",
			"encoding/binary":       "binary",
			"strconv":               "strconv",
			"sync":                  "sync",
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
		TypedConsts: map[string]ixgo.TypedConst{},
		UntypedConsts: map[string]ixgo.UntypedConst{
			"BlockSize": {"untyped int", constant.MakeInt64(int64(q.BlockSize))},
		},
	})
}
