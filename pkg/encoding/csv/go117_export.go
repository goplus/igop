// export by github.com/goplus/igop/cmd/qexp

//go:build go1.17 && !go1.18
// +build go1.17,!go1.18

package csv

import (
	q "encoding/csv"

	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
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
		NamedTypes: map[string]igop.NamedType{
			"ParseError": {reflect.TypeOf((*q.ParseError)(nil)).Elem(), "", "Error,Unwrap"},
			"Reader":     {reflect.TypeOf((*q.Reader)(nil)).Elem(), "", "FieldPos,Read,ReadAll,readLine,readRecord"},
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
		TypedConsts:   map[string]igop.TypedConst{},
		UntypedConsts: map[string]igop.UntypedConst{},
	})
}
