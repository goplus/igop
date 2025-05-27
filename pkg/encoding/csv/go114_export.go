// export by github.com/goplus/ixgo/cmd/qexp

//+build go1.14,!go1.15

package csv

import (
	q "encoding/csv"

	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage(&ixgo.Package{
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
		NamedTypes: map[string]reflect.Type{
			"ParseError": reflect.TypeOf((*q.ParseError)(nil)).Elem(),
			"Reader":     reflect.TypeOf((*q.Reader)(nil)).Elem(),
			"Writer":     reflect.TypeOf((*q.Writer)(nil)).Elem(),
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
		TypedConsts:   map[string]ixgo.TypedConst{},
		UntypedConsts: map[string]ixgo.UntypedConst{},
	})
}
