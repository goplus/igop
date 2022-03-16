// export by github.com/goplus/gossa/cmd/qexp

//go:build go1.18
// +build go1.18

package textproto

import (
	q "net/textproto"

	"reflect"

	"github.com/goplus/gossa"
)

func init() {
	gossa.RegisterPackage(&gossa.Package{
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
		NamedTypes: map[string]gossa.NamedType{
			"Conn":          {reflect.TypeOf((*q.Conn)(nil)).Elem(), "", "Close,Cmd"},
			"Error":         {reflect.TypeOf((*q.Error)(nil)).Elem(), "", "Error"},
			"MIMEHeader":    {reflect.TypeOf((*q.MIMEHeader)(nil)).Elem(), "Add,Del,Get,Set,Values", ""},
			"Pipeline":      {reflect.TypeOf((*q.Pipeline)(nil)).Elem(), "", "EndRequest,EndResponse,Next,StartRequest,StartResponse"},
			"ProtocolError": {reflect.TypeOf((*q.ProtocolError)(nil)).Elem(), "Error", ""},
			"Reader":        {reflect.TypeOf((*q.Reader)(nil)).Elem(), "", "DotReader,ReadCodeLine,ReadContinuedLine,ReadContinuedLineBytes,ReadDotBytes,ReadDotLines,ReadLine,ReadLineBytes,ReadMIMEHeader,ReadResponse,closeDot,readCodeLine,readContinuedLineSlice,readLineSlice,skipSpace,upcomingHeaderNewlines"},
			"Writer":        {reflect.TypeOf((*q.Writer)(nil)).Elem(), "", "DotWriter,PrintfLine,closeDot"},
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
		TypedConsts:   map[string]gossa.TypedConst{},
		UntypedConsts: map[string]gossa.UntypedConst{},
	})
}
