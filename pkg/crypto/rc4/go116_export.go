// export by github.com/goplus/gossa/cmd/qexp

//+build go1.16,!go1.17

package rc4

import (
	q "crypto/rc4"

	"reflect"

	"github.com/goplus/gossa"
)

func init() {
	gossa.RegisterPackage(&gossa.Package{
		Name: "rc4",
		Path: "crypto/rc4",
		Deps: map[string]string{
			"crypto/internal/subtle": "subtle",
			"strconv":                "strconv",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]gossa.NamedType{
			"Cipher":       {reflect.TypeOf((*q.Cipher)(nil)).Elem(), "", "Reset,XORKeyStream"},
			"KeySizeError": {reflect.TypeOf((*q.KeySizeError)(nil)).Elem(), "Error", ""},
		},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"NewCipher": reflect.ValueOf(q.NewCipher),
		},
		TypedConsts:   map[string]gossa.TypedConst{},
		UntypedConsts: map[string]gossa.UntypedConst{},
	})
}
