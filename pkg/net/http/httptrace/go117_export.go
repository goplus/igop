// export by github.com/goplus/gossa/cmd/qexp

//go:build go1.17
// +build go1.17

package httptrace

import (
	q "net/http/httptrace"

	"reflect"

	"github.com/goplus/gossa"
)

func init() {
	gossa.RegisterPackage(&gossa.Package{
		Name: "httptrace",
		Path: "net/http/httptrace",
		Deps: map[string]string{
			"context":           "context",
			"crypto/tls":        "tls",
			"internal/nettrace": "nettrace",
			"net":               "net",
			"net/textproto":     "textproto",
			"reflect":           "reflect",
			"time":              "time",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]gossa.NamedType{
			"ClientTrace":      {reflect.TypeOf((*q.ClientTrace)(nil)).Elem(), "", "compose,hasNetHooks"},
			"DNSDoneInfo":      {reflect.TypeOf((*q.DNSDoneInfo)(nil)).Elem(), "", ""},
			"DNSStartInfo":     {reflect.TypeOf((*q.DNSStartInfo)(nil)).Elem(), "", ""},
			"GotConnInfo":      {reflect.TypeOf((*q.GotConnInfo)(nil)).Elem(), "", ""},
			"WroteRequestInfo": {reflect.TypeOf((*q.WroteRequestInfo)(nil)).Elem(), "", ""},
		},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"ContextClientTrace": reflect.ValueOf(q.ContextClientTrace),
			"WithClientTrace":    reflect.ValueOf(q.WithClientTrace),
		},
		TypedConsts:   map[string]gossa.TypedConst{},
		UntypedConsts: map[string]gossa.UntypedConst{},
	})
}
