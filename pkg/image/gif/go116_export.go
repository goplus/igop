// export by github.com/goplus/igop/cmd/qexp

//+build go1.16,!go1.17

package gif

import (
	q "image/gif"

	"go/constant"
	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "gif",
		Path: "image/gif",
		Deps: map[string]string{
			"bufio":               "bufio",
			"bytes":               "bytes",
			"compress/lzw":        "lzw",
			"errors":              "errors",
			"fmt":                 "fmt",
			"image":               "image",
			"image/color":         "color",
			"image/color/palette": "palette",
			"image/draw":          "draw",
			"io":                  "io",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]igop.NamedType{
			"GIF":     {reflect.TypeOf((*q.GIF)(nil)).Elem(), "", ""},
			"Options": {reflect.TypeOf((*q.Options)(nil)).Elem(), "", ""},
		},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"Decode":       reflect.ValueOf(q.Decode),
			"DecodeAll":    reflect.ValueOf(q.DecodeAll),
			"DecodeConfig": reflect.ValueOf(q.DecodeConfig),
			"Encode":       reflect.ValueOf(q.Encode),
			"EncodeAll":    reflect.ValueOf(q.EncodeAll),
		},
		TypedConsts: map[string]igop.TypedConst{},
		UntypedConsts: map[string]igop.UntypedConst{
			"DisposalBackground": {"untyped int", constant.MakeInt64(int64(q.DisposalBackground))},
			"DisposalNone":       {"untyped int", constant.MakeInt64(int64(q.DisposalNone))},
			"DisposalPrevious":   {"untyped int", constant.MakeInt64(int64(q.DisposalPrevious))},
		},
	})
}
