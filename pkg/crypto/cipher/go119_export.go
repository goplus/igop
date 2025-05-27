// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.19 && !go1.20
// +build go1.19,!go1.20

package cipher

import (
	q "crypto/cipher"

	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage(&ixgo.Package{
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
		TypedConsts:   map[string]ixgo.TypedConst{},
		UntypedConsts: map[string]ixgo.UntypedConst{},
	})
}
