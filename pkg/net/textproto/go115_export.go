// export by github.com/goplus/ixgo/cmd/qexp

//+build go1.15,!go1.16

package textproto

import (
	q "net/textproto"

	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage(&ixgo.Package{
		Name: "textproto",
		Path: "net/textproto",
		Deps: map[string]string{
			"bufio":     "bufio",
			"bytes":     "bytes",
			"fmt":       "fmt",
			"io":        "io",
			"io/ioutil": "ioutil",
			"net":       "net",
			"strconv":   "strconv",
			"strings":   "strings",
			"sync":      "sync",
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
		TypedConsts:   map[string]ixgo.TypedConst{},
		UntypedConsts: map[string]ixgo.UntypedConst{},
	})
}
