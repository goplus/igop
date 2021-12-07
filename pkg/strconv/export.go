// export by github.com/goplus/gossa/cmd/qexp

package strconv

import (
	q "strconv"

	"go/constant"
	"reflect"

	"github.com/goplus/gossa"
)

func init() {
	gossa.RegisterPackage(&gossa.Package{
		Name: "strconv",
		Path: "strconv",
		Deps: map[string]string{
			"errors":           "errors",
			"internal/bytealg": "bytealg",
			"math":             "math",
			"math/bits":        "bits",
			"unicode/utf8":     "utf8",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]gossa.NamedType{
			"NumError": {reflect.TypeOf((*q.NumError)(nil)).Elem(), "", "Error,Unwrap"},
		},
		AliasTypes: map[string]reflect.Type{},
		Vars: map[string]reflect.Value{
			"ErrRange":  reflect.ValueOf(&q.ErrRange),
			"ErrSyntax": reflect.ValueOf(&q.ErrSyntax),
		},
		Funcs: map[string]reflect.Value{
			"AppendBool":               reflect.ValueOf(q.AppendBool),
			"AppendFloat":              reflect.ValueOf(q.AppendFloat),
			"AppendInt":                reflect.ValueOf(q.AppendInt),
			"AppendQuote":              reflect.ValueOf(q.AppendQuote),
			"AppendQuoteRune":          reflect.ValueOf(q.AppendQuoteRune),
			"AppendQuoteRuneToASCII":   reflect.ValueOf(q.AppendQuoteRuneToASCII),
			"AppendQuoteRuneToGraphic": reflect.ValueOf(q.AppendQuoteRuneToGraphic),
			"AppendQuoteToASCII":       reflect.ValueOf(q.AppendQuoteToASCII),
			"AppendQuoteToGraphic":     reflect.ValueOf(q.AppendQuoteToGraphic),
			"AppendUint":               reflect.ValueOf(q.AppendUint),
			"Atoi":                     reflect.ValueOf(q.Atoi),
			"CanBackquote":             reflect.ValueOf(q.CanBackquote),
			"FormatBool":               reflect.ValueOf(q.FormatBool),
			"FormatComplex":            reflect.ValueOf(q.FormatComplex),
			"FormatFloat":              reflect.ValueOf(q.FormatFloat),
			"FormatInt":                reflect.ValueOf(q.FormatInt),
			"FormatUint":               reflect.ValueOf(q.FormatUint),
			"IsGraphic":                reflect.ValueOf(q.IsGraphic),
			"IsPrint":                  reflect.ValueOf(q.IsPrint),
			"Itoa":                     reflect.ValueOf(q.Itoa),
			"ParseBool":                reflect.ValueOf(q.ParseBool),
			"ParseComplex":             reflect.ValueOf(q.ParseComplex),
			"ParseFloat":               reflect.ValueOf(q.ParseFloat),
			"ParseInt":                 reflect.ValueOf(q.ParseInt),
			"ParseUint":                reflect.ValueOf(q.ParseUint),
			"Quote":                    reflect.ValueOf(q.Quote),
			"QuoteRune":                reflect.ValueOf(q.QuoteRune),
			"QuoteRuneToASCII":         reflect.ValueOf(q.QuoteRuneToASCII),
			"QuoteRuneToGraphic":       reflect.ValueOf(q.QuoteRuneToGraphic),
			"QuoteToASCII":             reflect.ValueOf(q.QuoteToASCII),
			"QuoteToGraphic":           reflect.ValueOf(q.QuoteToGraphic),
			"Unquote":                  reflect.ValueOf(q.Unquote),
			"UnquoteChar":              reflect.ValueOf(q.UnquoteChar),
		},
		TypedConsts: map[string]gossa.TypedConst{},
		UntypedConsts: map[string]gossa.UntypedConst{
			"IntSize": {"untyped int", constant.MakeInt64(int64(q.IntSize))},
		},
	})
}
