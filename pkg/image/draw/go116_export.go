// export by github.com/goplus/gossa/cmd/qexp

//+build go1.16,!go1.17

package draw

import (
	q "image/draw"

	"go/constant"
	"reflect"

	"github.com/goplus/gossa"
)

func init() {
	gossa.RegisterPackage(&gossa.Package{
		Name: "draw",
		Path: "image/draw",
		Deps: map[string]string{
			"image":                    "image",
			"image/color":              "color",
			"image/internal/imageutil": "imageutil",
		},
		Interfaces: map[string]reflect.Type{
			"Drawer":    reflect.TypeOf((*q.Drawer)(nil)).Elem(),
			"Image":     reflect.TypeOf((*q.Image)(nil)).Elem(),
			"Quantizer": reflect.TypeOf((*q.Quantizer)(nil)).Elem(),
		},
		NamedTypes: map[string]gossa.NamedType{
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
		TypedConsts: map[string]gossa.TypedConst{
			"Over": {reflect.TypeOf(q.Over), constant.MakeInt64(int64(q.Over))},
			"Src":  {reflect.TypeOf(q.Src), constant.MakeInt64(int64(q.Src))},
		},
		UntypedConsts: map[string]gossa.UntypedConst{},
	})
}
