// export by github.com/goplus/gossa/cmd/qexp

//+build go1.16,!go1.17

package json

import (
	q "encoding/json"

	"reflect"

	"github.com/goplus/gossa"
)

func init() {
	gossa.RegisterPackage(&gossa.Package{
		Name: "json",
		Path: "encoding/json",
		Deps: map[string]string{
			"bytes":           "bytes",
			"encoding":        "encoding",
			"encoding/base64": "base64",
			"errors":          "errors",
			"fmt":             "fmt",
			"io":              "io",
			"math":            "math",
			"reflect":         "reflect",
			"sort":            "sort",
			"strconv":         "strconv",
			"strings":         "strings",
			"sync":            "sync",
			"unicode":         "unicode",
			"unicode/utf16":   "utf16",
			"unicode/utf8":    "utf8",
		},
		Interfaces: map[string]reflect.Type{
			"Marshaler":   reflect.TypeOf((*q.Marshaler)(nil)).Elem(),
			"Token":       reflect.TypeOf((*q.Token)(nil)).Elem(),
			"Unmarshaler": reflect.TypeOf((*q.Unmarshaler)(nil)).Elem(),
		},
		NamedTypes: map[string]gossa.NamedType{
			"Decoder":               {reflect.TypeOf((*q.Decoder)(nil)).Elem(), "", "Buffered,Decode,DisallowUnknownFields,InputOffset,More,Token,UseNumber,peek,readValue,refill,tokenError,tokenPrepareForDecode,tokenValueAllowed,tokenValueEnd"},
			"Delim":                 {reflect.TypeOf((*q.Delim)(nil)).Elem(), "String", ""},
			"Encoder":               {reflect.TypeOf((*q.Encoder)(nil)).Elem(), "", "Encode,SetEscapeHTML,SetIndent"},
			"InvalidUTF8Error":      {reflect.TypeOf((*q.InvalidUTF8Error)(nil)).Elem(), "", "Error"},
			"InvalidUnmarshalError": {reflect.TypeOf((*q.InvalidUnmarshalError)(nil)).Elem(), "", "Error"},
			"MarshalerError":        {reflect.TypeOf((*q.MarshalerError)(nil)).Elem(), "", "Error,Unwrap"},
			"Number":                {reflect.TypeOf((*q.Number)(nil)).Elem(), "Float64,Int64,String", ""},
			"RawMessage":            {reflect.TypeOf((*q.RawMessage)(nil)).Elem(), "MarshalJSON", "UnmarshalJSON"},
			"SyntaxError":           {reflect.TypeOf((*q.SyntaxError)(nil)).Elem(), "", "Error"},
			"UnmarshalFieldError":   {reflect.TypeOf((*q.UnmarshalFieldError)(nil)).Elem(), "", "Error"},
			"UnmarshalTypeError":    {reflect.TypeOf((*q.UnmarshalTypeError)(nil)).Elem(), "", "Error"},
			"UnsupportedTypeError":  {reflect.TypeOf((*q.UnsupportedTypeError)(nil)).Elem(), "", "Error"},
			"UnsupportedValueError": {reflect.TypeOf((*q.UnsupportedValueError)(nil)).Elem(), "", "Error"},
		},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"Compact":       reflect.ValueOf(q.Compact),
			"HTMLEscape":    reflect.ValueOf(q.HTMLEscape),
			"Indent":        reflect.ValueOf(q.Indent),
			"Marshal":       reflect.ValueOf(q.Marshal),
			"MarshalIndent": reflect.ValueOf(q.MarshalIndent),
			"NewDecoder":    reflect.ValueOf(q.NewDecoder),
			"NewEncoder":    reflect.ValueOf(q.NewEncoder),
			"Unmarshal":     reflect.ValueOf(q.Unmarshal),
			"Valid":         reflect.ValueOf(q.Valid),
		},
		TypedConsts:   map[string]gossa.TypedConst{},
		UntypedConsts: map[string]gossa.UntypedConst{},
	})
}
