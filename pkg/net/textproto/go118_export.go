// export by github.com/goplus/igop/cmd/qexp

//go:build go1.18
// +build go1.18

package textproto

import (
	q "net/textproto"

	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "textproto",
		Path: "net/textproto",
		Deps: map[string]string{
			"bufio":   "bufio",
			"bytes":   "bytes",
			"fmt":     "fmt",
			"io":      "io",
			"net":     "net",
			"strconv": "strconv",
			"strings": "strings",
			"sync":    "sync",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{
			"Conn":          reflect.TypeOf((*q.Conn)(nil)).Elem(),
			"Error":         reflect.TypeOf((*q.Error)(nil)).Elem(),
			"MIMEHeader":    reflect.TypeOf((*q.MIMEHeader)(nil)).Elem(),
			"Pipeline":      reflect.TypeOf((*q.Pipeline)(nil)).Elem(),
			"ProtocolError": reflect.TypeOf((*q.ProtocolError)(nil)).Elem(),
			"Reader":        reflect.TypeOf((*q.Reader)(nil)).Elem(),
			"Writer":        reflect.TypeOf((*q.Writer)(nil)).Elem(),
		},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"CanonicalMIMEHeaderKey": reflect.ValueOf(q.CanonicalMIMEHeaderKey),
			"Dial":                   reflect.ValueOf(q.Dial),
			"NewConn":                reflect.ValueOf(q.NewConn),
			"NewReader":              reflect.ValueOf(q.NewReader),
			"NewWriter":              reflect.ValueOf(q.NewWriter),
			"TrimBytes":              reflect.ValueOf(q.TrimBytes),
			"TrimString":             reflect.ValueOf(q.TrimString),
		},
		TypedConsts:   map[string]igop.TypedConst{},
		UntypedConsts: map[string]igop.UntypedConst{},
	})
}
