// export by github.com/goplus/igop/cmd/qexp

//go:build go1.20 && !go1.20
// +build go1.20,!go1.20

package mime

import (
	q "mime"

	"go/constant"
	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
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
		NamedTypes: map[string]reflect.Type{
			"WordDecoder": reflect.TypeOf((*q.WordDecoder)(nil)).Elem(),
			"WordEncoder": reflect.TypeOf((*q.WordEncoder)(nil)).Elem(),
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
		TypedConsts: map[string]igop.TypedConst{
			"BEncoding": {reflect.TypeOf(q.BEncoding), constant.MakeInt64(int64(q.BEncoding))},
			"QEncoding": {reflect.TypeOf(q.QEncoding), constant.MakeInt64(int64(q.QEncoding))},
		},
		UntypedConsts: map[string]igop.UntypedConst{},
	})
}
