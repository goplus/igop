// export by github.com/goplus/igop/cmd/qexp

//go:build go1.17 && !go1.18
// +build go1.17,!go1.18

package draw

import (
	q "image/draw"

	"go/constant"
	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "draw",
		Path: "image/draw",
		Deps: map[string]string{
			"image":                    "image",
			"image/color":              "color",
			"image/internal/imageutil": "imageutil",
		},
		Interfaces: map[string]reflect.Type{
			"Drawer":      reflect.TypeOf((*q.Drawer)(nil)).Elem(),
			"Image":       reflect.TypeOf((*q.Image)(nil)).Elem(),
			"Quantizer":   reflect.TypeOf((*q.Quantizer)(nil)).Elem(),
			"RGBA64Image": reflect.TypeOf((*q.RGBA64Image)(nil)).Elem(),
		},
		NamedTypes: map[string]igop.NamedType{
			"Op": {reflect.TypeOf((*q.Op)(nil)).Elem(), "Draw", ""},
		},
		AliasTypes: map[string]reflect.Type{},
		Vars: map[string]reflect.Value{
			"FloydSteinberg": reflect.ValueOf(&q.FloydSteinberg),
		},
		Funcs: map[string]reflect.Value{
			"Draw":     reflect.ValueOf(q.Draw),
			"DrawMask": reflect.ValueOf(q.DrawMask),
		},
		TypedConsts: map[string]igop.TypedConst{
			"Over": {reflect.TypeOf(q.Over), constant.MakeInt64(int64(q.Over))},
			"Src":  {reflect.TypeOf(q.Src), constant.MakeInt64(int64(q.Src))},
		},
		UntypedConsts: map[string]igop.UntypedConst{},
	})
}
