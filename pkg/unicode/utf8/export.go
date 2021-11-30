// export by github.com/goplus/gossa/cmd/qexp

package utf8

import (
	q "unicode/utf8"

	"go/constant"
	"reflect"

	"github.com/goplus/gossa"
)

func init() {
	gossa.RegisterPackage(&gossa.Package{
		Name:       "utf8",
		Path:       "unicode/utf8",
		Deps:       map[string]string{},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]gossa.NamedType{},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"DecodeLastRune":         reflect.ValueOf(q.DecodeLastRune),
			"DecodeLastRuneInString": reflect.ValueOf(q.DecodeLastRuneInString),
			"DecodeRune":             reflect.ValueOf(q.DecodeRune),
			"DecodeRuneInString":     reflect.ValueOf(q.DecodeRuneInString),
			"EncodeRune":             reflect.ValueOf(q.EncodeRune),
			"FullRune":               reflect.ValueOf(q.FullRune),
			"FullRuneInString":       reflect.ValueOf(q.FullRuneInString),
			"RuneCount":              reflect.ValueOf(q.RuneCount),
			"RuneCountInString":      reflect.ValueOf(q.RuneCountInString),
			"RuneLen":                reflect.ValueOf(q.RuneLen),
			"RuneStart":              reflect.ValueOf(q.RuneStart),
			"Valid":                  reflect.ValueOf(q.Valid),
			"ValidRune":              reflect.ValueOf(q.ValidRune),
			"ValidString":            reflect.ValueOf(q.ValidString),
		},
		TypedConsts: map[string]gossa.TypedConst{},
		UntypedConsts: map[string]gossa.UntypedConst{
			"MaxRune":   {"untyped rune", constant.MakeInt64(1114111)},
			"RuneError": {"untyped rune", constant.MakeInt64(65533)},
			"RuneSelf":  {"untyped int", constant.MakeInt64(128)},
			"UTFMax":    {"untyped int", constant.MakeInt64(4)},
		},
	})
}
