// export by github.com/goplus/gossa/cmd/qexp

package csv

import (
	q "encoding/csv"

	"reflect"

	"github.com/goplus/gossa"
)

func init() {
	gossa.RegisterPackage(&gossa.Package{
		Name: "csv",
		Path: "encoding/csv",
		Deps: map[string]string{
			"bufio":        "bufio",
			"bytes":        "bytes",
			"errors":       "errors",
			"fmt":          "fmt",
			"io":           "io",
			"strings":      "strings",
			"unicode":      "unicode",
			"unicode/utf8": "utf8",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]gossa.NamedType{
			"ParseError": {reflect.TypeOf((*q.ParseError)(nil)).Elem(), "", "Error,Unwrap"},
			"Reader":     {reflect.TypeOf((*q.Reader)(nil)).Elem(), "", "Read,ReadAll,readLine,readRecord"},
			"Writer":     {reflect.TypeOf((*q.Writer)(nil)).Elem(), "", "Error,Flush,Write,WriteAll,fieldNeedsQuotes"},
		},
		AliasTypes: map[string]reflect.Type{},
		Vars: map[string]reflect.Value{
			"ErrBareQuote":     reflect.ValueOf(&q.ErrBareQuote),
			"ErrFieldCount":    reflect.ValueOf(&q.ErrFieldCount),
			"ErrQuote":         reflect.ValueOf(&q.ErrQuote),
			"ErrTrailingComma": reflect.ValueOf(&q.ErrTrailingComma),
		},
		Funcs: map[string]reflect.Value{
			"NewReader": reflect.ValueOf(q.NewReader),
			"NewWriter": reflect.ValueOf(q.NewWriter),
		},
		TypedConsts:   map[string]gossa.TypedConst{},
		UntypedConsts: map[string]gossa.UntypedConst{},
	})
}
