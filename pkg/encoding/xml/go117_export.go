// export by github.com/goplus/gossa/cmd/qexp

//go:build go1.17 && !go1.18
// +build go1.17,!go1.18

package xml

import (
	q "encoding/xml"

	"go/constant"
	"reflect"

	"github.com/goplus/gossa"
)

func init() {
	gossa.RegisterPackage(&gossa.Package{
		Name: "xml",
		Path: "encoding/xml",
		Deps: map[string]string{
			"bufio":        "bufio",
			"bytes":        "bytes",
			"encoding":     "encoding",
			"errors":       "errors",
			"fmt":          "fmt",
			"io":           "io",
			"reflect":      "reflect",
			"strconv":      "strconv",
			"strings":      "strings",
			"sync":         "sync",
			"unicode":      "unicode",
			"unicode/utf8": "utf8",
		},
		Interfaces: map[string]reflect.Type{
			"Marshaler":       reflect.TypeOf((*q.Marshaler)(nil)).Elem(),
			"MarshalerAttr":   reflect.TypeOf((*q.MarshalerAttr)(nil)).Elem(),
			"Token":           reflect.TypeOf((*q.Token)(nil)).Elem(),
			"TokenReader":     reflect.TypeOf((*q.TokenReader)(nil)).Elem(),
			"Unmarshaler":     reflect.TypeOf((*q.Unmarshaler)(nil)).Elem(),
			"UnmarshalerAttr": reflect.TypeOf((*q.UnmarshalerAttr)(nil)).Elem(),
		},
		NamedTypes: map[string]gossa.NamedType{
			"Attr":                 {reflect.TypeOf((*q.Attr)(nil)).Elem(), "", ""},
			"CharData":             {reflect.TypeOf((*q.CharData)(nil)).Elem(), "Copy", ""},
			"Comment":              {reflect.TypeOf((*q.Comment)(nil)).Elem(), "Copy", ""},
			"Decoder":              {reflect.TypeOf((*q.Decoder)(nil)).Elem(), "", "Decode,DecodeElement,InputOffset,RawToken,Skip,Token,attrval,autoClose,getc,mustgetc,name,nsname,pop,popEOF,popElement,push,pushEOF,pushElement,pushNs,rawToken,readName,savedOffset,space,switchToReader,syntaxError,text,translate,ungetc,unmarshal,unmarshalAttr,unmarshalInterface,unmarshalPath,unmarshalTextInterface"},
			"Directive":            {reflect.TypeOf((*q.Directive)(nil)).Elem(), "Copy", ""},
			"Encoder":              {reflect.TypeOf((*q.Encoder)(nil)).Elem(), "", "Encode,EncodeElement,EncodeToken,Flush,Indent"},
			"EndElement":           {reflect.TypeOf((*q.EndElement)(nil)).Elem(), "", ""},
			"Name":                 {reflect.TypeOf((*q.Name)(nil)).Elem(), "", ""},
			"ProcInst":             {reflect.TypeOf((*q.ProcInst)(nil)).Elem(), "Copy", ""},
			"StartElement":         {reflect.TypeOf((*q.StartElement)(nil)).Elem(), "Copy,End", ""},
			"SyntaxError":          {reflect.TypeOf((*q.SyntaxError)(nil)).Elem(), "", "Error"},
			"TagPathError":         {reflect.TypeOf((*q.TagPathError)(nil)).Elem(), "", "Error"},
			"UnmarshalError":       {reflect.TypeOf((*q.UnmarshalError)(nil)).Elem(), "Error", ""},
			"UnsupportedTypeError": {reflect.TypeOf((*q.UnsupportedTypeError)(nil)).Elem(), "", "Error"},
		},
		AliasTypes: map[string]reflect.Type{},
		Vars: map[string]reflect.Value{
			"HTMLAutoClose": reflect.ValueOf(&q.HTMLAutoClose),
			"HTMLEntity":    reflect.ValueOf(&q.HTMLEntity),
		},
		Funcs: map[string]reflect.Value{
			"CopyToken":       reflect.ValueOf(q.CopyToken),
			"Escape":          reflect.ValueOf(q.Escape),
			"EscapeText":      reflect.ValueOf(q.EscapeText),
			"Marshal":         reflect.ValueOf(q.Marshal),
			"MarshalIndent":   reflect.ValueOf(q.MarshalIndent),
			"NewDecoder":      reflect.ValueOf(q.NewDecoder),
			"NewEncoder":      reflect.ValueOf(q.NewEncoder),
			"NewTokenDecoder": reflect.ValueOf(q.NewTokenDecoder),
			"Unmarshal":       reflect.ValueOf(q.Unmarshal),
		},
		TypedConsts: map[string]gossa.TypedConst{},
		UntypedConsts: map[string]gossa.UntypedConst{
			"Header": {"untyped string", constant.MakeString(string(q.Header))},
		},
	})
}
