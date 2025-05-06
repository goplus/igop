// export by github.com/goplus/igop/cmd/qexp

//go:build go1.23 && !go1.24
// +build go1.23,!go1.24

package cipher

import (
	q "crypto/cipher"

	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "cipher",
		Path: "crypto/cipher",
		Deps: map[string]string{
			"bytes":                 "bytes",
			"crypto/internal/alias": "alias",
			"crypto/subtle":         "subtle",
			"errors":                "errors",
			"internal/byteorder":    "byteorder",
			"io":                    "io",
		},
		Interfaces: map[string]reflect.Type{
			"AEAD":      reflect.TypeOf((*q.AEAD)(nil)).Elem(),
			"Block":     reflect.TypeOf((*q.Block)(nil)).Elem(),
			"BlockMode": reflect.TypeOf((*q.BlockMode)(nil)).Elem(),
			"Stream":    reflect.TypeOf((*q.Stream)(nil)).Elem(),
		},
		NamedTypes: map[string]reflect.Type{
			"StreamReader": reflect.TypeOf((*q.StreamReader)(nil)).Elem(),
			"StreamWriter": reflect.TypeOf((*q.StreamWriter)(nil)).Elem(),
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
		TypedConsts:   map[string]igop.TypedConst{},
		UntypedConsts: map[string]igop.UntypedConst{},
	})
}
