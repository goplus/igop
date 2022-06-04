// export by github.com/goplus/igop/cmd/qexp

//+build go1.15,!go1.16

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
			"crypto/internal/subtle": "subtle",
			"strconv":                "strconv",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]igop.NamedType{
			"Cipher":       {reflect.TypeOf((*q.Cipher)(nil)).Elem(), "", "Reset,XORKeyStream"},
			"KeySizeError": {reflect.TypeOf((*q.KeySizeError)(nil)).Elem(), "Error", ""},
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
