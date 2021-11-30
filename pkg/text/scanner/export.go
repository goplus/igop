// export by github.com/goplus/gossa/cmd/qexp

package scanner

import (
	q "text/scanner"

	"go/constant"
	"reflect"

	"github.com/goplus/gossa"
)

func init() {
	gossa.RegisterPackage(&gossa.Package{
		Name: "scanner",
		Path: "text/scanner",
		Deps: map[string]string{
			"bytes":        "bytes",
			"fmt":          "fmt",
			"io":           "io",
			"os":           "os",
			"unicode":      "unicode",
			"unicode/utf8": "utf8",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]gossa.NamedType{
			"Position": {reflect.TypeOf((*q.Position)(nil)).Elem(), "String", "IsValid"},
			"Scanner":  {reflect.TypeOf((*q.Scanner)(nil)).Elem(), "", "Init,Next,Peek,Pos,Scan,TokenText,digits,error,errorf,isIdentRune,next,scanChar,scanComment,scanDigits,scanEscape,scanIdentifier,scanNumber,scanRawString,scanString"},
		},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"TokenString": reflect.ValueOf(q.TokenString),
		},
		TypedConsts: map[string]gossa.TypedConst{},
		UntypedConsts: map[string]gossa.UntypedConst{
			"Char":           {"untyped int", constant.MakeInt64(-5)},
			"Comment":        {"untyped int", constant.MakeInt64(-8)},
			"EOF":            {"untyped int", constant.MakeInt64(-1)},
			"Float":          {"untyped int", constant.MakeInt64(-4)},
			"GoTokens":       {"untyped int", constant.MakeInt64(1012)},
			"GoWhitespace":   {"untyped int", constant.MakeInt64(4294977024)},
			"Ident":          {"untyped int", constant.MakeInt64(-2)},
			"Int":            {"untyped int", constant.MakeInt64(-3)},
			"RawString":      {"untyped int", constant.MakeInt64(-7)},
			"ScanChars":      {"untyped int", constant.MakeInt64(32)},
			"ScanComments":   {"untyped int", constant.MakeInt64(256)},
			"ScanFloats":     {"untyped int", constant.MakeInt64(16)},
			"ScanIdents":     {"untyped int", constant.MakeInt64(4)},
			"ScanInts":       {"untyped int", constant.MakeInt64(8)},
			"ScanRawStrings": {"untyped int", constant.MakeInt64(128)},
			"ScanStrings":    {"untyped int", constant.MakeInt64(64)},
			"SkipComments":   {"untyped int", constant.MakeInt64(512)},
			"String":         {"untyped int", constant.MakeInt64(-6)},
		},
	})
}
