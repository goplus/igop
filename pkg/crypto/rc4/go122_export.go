// export by github.com/goplus/igop/cmd/qexp

//go:build go1.22 && !go1.23
// +build go1.22,!go1.23

package rc4

import (
	q "crypto/rc4"

	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "rc4",
		Path: "crypto/rc4",
		Deps: map[string]string{
			"crypto/internal/alias": "alias",
			"strconv":               "strconv",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{
			"Cipher":       reflect.TypeOf((*q.Cipher)(nil)).Elem(),
			"KeySizeError": reflect.TypeOf((*q.KeySizeError)(nil)).Elem(),
		},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"NewCipher": reflect.ValueOf(q.NewCipher),
		},
		TypedConsts:   map[string]igop.TypedConst{},
		UntypedConsts: map[string]igop.UntypedConst{},
	})
}
