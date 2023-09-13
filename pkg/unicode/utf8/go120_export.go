// export by github.com/goplus/igop/cmd/qexp

//go:build go1.20 && !go1.21
// +build go1.20,!go1.21

package utf8

import (
	q "unicode/utf8"

	"go/constant"
	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name:       "utf8",
		Path:       "unicode/utf8",
		Deps:       map[string]string{},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"AppendRune":             reflect.ValueOf(q.AppendRune),
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
		TypedConsts: map[string]igop.TypedConst{},
		UntypedConsts: map[string]igop.UntypedConst{
			"MaxRune":   {"untyped rune", constant.MakeInt64(int64(q.MaxRune))},
			"RuneError": {"untyped rune", constant.MakeInt64(int64(q.RuneError))},
			"RuneSelf":  {"untyped int", constant.MakeInt64(int64(q.RuneSelf))},
			"UTFMax":    {"untyped int", constant.MakeInt64(int64(q.UTFMax))},
		},
	})
}
