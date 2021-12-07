// export by github.com/goplus/gossa/cmd/qexp

package des

import (
	q "crypto/des"

	"go/constant"
	"reflect"

	"github.com/goplus/gossa"
)

func init() {
	gossa.RegisterPackage(&gossa.Package{
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
		NamedTypes: map[string]gossa.NamedType{
			"KeySizeError": {reflect.TypeOf((*q.KeySizeError)(nil)).Elem(), "Error", ""},
		},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"NewCipher":          reflect.ValueOf(q.NewCipher),
			"NewTripleDESCipher": reflect.ValueOf(q.NewTripleDESCipher),
		},
		TypedConsts: map[string]gossa.TypedConst{},
		UntypedConsts: map[string]gossa.UntypedConst{
			"BlockSize": {"untyped int", constant.MakeInt64(int64(q.BlockSize))},
		},
	})
}
