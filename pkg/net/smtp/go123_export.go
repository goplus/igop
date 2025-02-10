// export by github.com/goplus/igop/cmd/qexp

//go:build go1.23
// +build go1.23

package smtp

import (
	q "net/smtp"

	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "smtp",
		Path: "net/smtp",
		Deps: map[string]string{
			"crypto/hmac":     "hmac",
			"crypto/md5":      "md5",
			"crypto/tls":      "tls",
			"encoding/base64": "base64",
			"errors":          "errors",
			"fmt":             "fmt",
			"io":              "io",
			"net":             "net",
			"net/textproto":   "textproto",
			"strings":         "strings",
		},
		Interfaces: map[string]reflect.Type{
			"Auth": reflect.TypeOf((*q.Auth)(nil)).Elem(),
		},
		NamedTypes: map[string]reflect.Type{
			"Client":     reflect.TypeOf((*q.Client)(nil)).Elem(),
			"ServerInfo": reflect.TypeOf((*q.ServerInfo)(nil)).Elem(),
		},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"CRAMMD5Auth": reflect.ValueOf(q.CRAMMD5Auth),
			"Dial":        reflect.ValueOf(q.Dial),
			"NewClient":   reflect.ValueOf(q.NewClient),
			"PlainAuth":   reflect.ValueOf(q.PlainAuth),
			"SendMail":    reflect.ValueOf(q.SendMail),
		},
		TypedConsts:   map[string]igop.TypedConst{},
		UntypedConsts: map[string]igop.UntypedConst{},
	})
}
