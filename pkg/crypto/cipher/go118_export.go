// export by github.com/goplus/gossa/cmd/qexp

//go:build go1.18
// +build go1.18

package cipher

import (
	q "crypto/cipher"

	"reflect"

	"github.com/goplus/gossa"
)

func init() {
	gossa.RegisterPackage(&gossa.Package{
		Name: "cipher",
		Path: "crypto/cipher",
		Deps: map[string]string{
			"crypto/internal/subtle": "subtle",
			"crypto/subtle":          "subtle",
			"encoding/binary":        "binary",
			"errors":                 "errors",
			"io":                     "io",
		},
		Interfaces: map[string]reflect.Type{
			"AEAD":      reflect.TypeOf((*q.AEAD)(nil)).Elem(),
			"Block":     reflect.TypeOf((*q.Block)(nil)).Elem(),
			"BlockMode": reflect.TypeOf((*q.BlockMode)(nil)).Elem(),
			"Stream":    reflect.TypeOf((*q.Stream)(nil)).Elem(),
		},
		NamedTypes: map[string]gossa.NamedType{
			"StreamReader": {reflect.TypeOf((*q.StreamReader)(nil)).Elem(), "Read", ""},
			"StreamWriter": {reflect.TypeOf((*q.StreamWriter)(nil)).Elem(), "Close,Write", ""},
		},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"NewCBCDecrypter":     reflect.ValueOf(q.NewCBCDecrypter),
			"NewCBCEncrypter":     reflect.ValueOf(q.NewCBCEncrypter),
			"NewCFBDecrypter":     reflect.ValueOf(q.NewCFBDecrypter),
			"NewCFBEncrypter":     reflect.ValueOf(q.NewCFBEncrypter),
			"NewCTR":              reflect.ValueOf(q.NewCTR),
			"NewGCM":              reflect.ValueOf(q.NewGCM),
			"NewGCMWithNonceSize": reflect.ValueOf(q.NewGCMWithNonceSize),
			"NewGCMWithTagSize":   reflect.ValueOf(q.NewGCMWithTagSize),
			"NewOFB":              reflect.ValueOf(q.NewOFB),
		},
		TypedConsts:   map[string]gossa.TypedConst{},
		UntypedConsts: map[string]gossa.UntypedConst{},
	})
}
