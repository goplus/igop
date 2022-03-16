// export by github.com/goplus/gossa/cmd/qexp

//go:build go1.18
// +build go1.18

package jpeg

import (
	q "image/jpeg"

	"go/constant"
	"reflect"

	"github.com/goplus/gossa"
)

func init() {
	gossa.RegisterPackage(&gossa.Package{
		Name: "jpeg",
		Path: "image/jpeg",
		Deps: map[string]string{
			"bufio":                    "bufio",
			"errors":                   "errors",
			"image":                    "image",
			"image/color":              "color",
			"image/internal/imageutil": "imageutil",
			"io":                       "io",
		},
		Interfaces: map[string]reflect.Type{
			"Reader": reflect.TypeOf((*q.Reader)(nil)).Elem(),
		},
		NamedTypes: map[string]gossa.NamedType{
			"FormatError":      {reflect.TypeOf((*q.FormatError)(nil)).Elem(), "Error", ""},
			"Options":          {reflect.TypeOf((*q.Options)(nil)).Elem(), "", ""},
			"UnsupportedError": {reflect.TypeOf((*q.UnsupportedError)(nil)).Elem(), "Error", ""},
		},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"Decode":       reflect.ValueOf(q.Decode),
			"DecodeConfig": reflect.ValueOf(q.DecodeConfig),
			"Encode":       reflect.ValueOf(q.Encode),
		},
		TypedConsts: map[string]gossa.TypedConst{},
		UntypedConsts: map[string]gossa.UntypedConst{
			"DefaultQuality": {"untyped int", constant.MakeInt64(int64(q.DefaultQuality))},
		},
	})
}
