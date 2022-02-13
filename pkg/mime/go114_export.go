// export by github.com/goplus/gossa/cmd/qexp

//+build go1.14,!go1.15

package mime

import (
	q "mime"

	"go/constant"
	"reflect"

	"github.com/goplus/gossa"
)

func init() {
	gossa.RegisterPackage(&gossa.Package{
		Name: "mime",
		Path: "mime",
		Deps: map[string]string{
			"bufio":           "bufio",
			"bytes":           "bytes",
			"encoding/base64": "base64",
			"errors":          "errors",
			"fmt":             "fmt",
			"io":              "io",
			"os":              "os",
			"sort":            "sort",
			"strings":         "strings",
			"sync":            "sync",
			"unicode":         "unicode",
			"unicode/utf8":    "utf8",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]gossa.NamedType{
			"WordDecoder": {reflect.TypeOf((*q.WordDecoder)(nil)).Elem(), "", "Decode,DecodeHeader,convert"},
			"WordEncoder": {reflect.TypeOf((*q.WordEncoder)(nil)).Elem(), "Encode,bEncode,encodeWord,openWord,qEncode,splitWord", ""},
		},
		AliasTypes: map[string]reflect.Type{},
		Vars: map[string]reflect.Value{
			"ErrInvalidMediaParameter": reflect.ValueOf(&q.ErrInvalidMediaParameter),
		},
		Funcs: map[string]reflect.Value{
			"AddExtensionType": reflect.ValueOf(q.AddExtensionType),
			"ExtensionsByType": reflect.ValueOf(q.ExtensionsByType),
			"FormatMediaType":  reflect.ValueOf(q.FormatMediaType),
			"ParseMediaType":   reflect.ValueOf(q.ParseMediaType),
			"TypeByExtension":  reflect.ValueOf(q.TypeByExtension),
		},
		TypedConsts: map[string]gossa.TypedConst{
			"BEncoding": {reflect.TypeOf(q.BEncoding), constant.MakeInt64(int64(q.BEncoding))},
			"QEncoding": {reflect.TypeOf(q.QEncoding), constant.MakeInt64(int64(q.QEncoding))},
		},
		UntypedConsts: map[string]gossa.UntypedConst{},
	})
}
