// export by github.com/goplus/igop/cmd/qexp

//+build go1.14,!go1.15

package image

import (
	q "image"

	"go/constant"
	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "image",
		Path: "image",
		Deps: map[string]string{
			"bufio":       "bufio",
			"errors":      "errors",
			"image/color": "color",
			"io":          "io",
			"strconv":     "strconv",
			"sync":        "sync",
			"sync/atomic": "atomic",
		},
		Interfaces: map[string]reflect.Type{
			"Image":         reflect.TypeOf((*q.Image)(nil)).Elem(),
			"PalettedImage": reflect.TypeOf((*q.PalettedImage)(nil)).Elem(),
		},
		NamedTypes: map[string]igop.NamedType{
			"Alpha":               {reflect.TypeOf((*q.Alpha)(nil)).Elem(), "", "AlphaAt,At,Bounds,ColorModel,Opaque,PixOffset,Set,SetAlpha,SubImage"},
			"Alpha16":             {reflect.TypeOf((*q.Alpha16)(nil)).Elem(), "", "Alpha16At,At,Bounds,ColorModel,Opaque,PixOffset,Set,SetAlpha16,SubImage"},
			"CMYK":                {reflect.TypeOf((*q.CMYK)(nil)).Elem(), "", "At,Bounds,CMYKAt,ColorModel,Opaque,PixOffset,Set,SetCMYK,SubImage"},
			"Config":              {reflect.TypeOf((*q.Config)(nil)).Elem(), "", ""},
			"Gray":                {reflect.TypeOf((*q.Gray)(nil)).Elem(), "", "At,Bounds,ColorModel,GrayAt,Opaque,PixOffset,Set,SetGray,SubImage"},
			"Gray16":              {reflect.TypeOf((*q.Gray16)(nil)).Elem(), "", "At,Bounds,ColorModel,Gray16At,Opaque,PixOffset,Set,SetGray16,SubImage"},
			"NRGBA":               {reflect.TypeOf((*q.NRGBA)(nil)).Elem(), "", "At,Bounds,ColorModel,NRGBAAt,Opaque,PixOffset,Set,SetNRGBA,SubImage"},
			"NRGBA64":             {reflect.TypeOf((*q.NRGBA64)(nil)).Elem(), "", "At,Bounds,ColorModel,NRGBA64At,Opaque,PixOffset,Set,SetNRGBA64,SubImage"},
			"NYCbCrA":             {reflect.TypeOf((*q.NYCbCrA)(nil)).Elem(), "", "AOffset,At,ColorModel,NYCbCrAAt,Opaque,SubImage"},
			"Paletted":            {reflect.TypeOf((*q.Paletted)(nil)).Elem(), "", "At,Bounds,ColorIndexAt,ColorModel,Opaque,PixOffset,Set,SetColorIndex,SubImage"},
			"Point":               {reflect.TypeOf((*q.Point)(nil)).Elem(), "Add,Div,Eq,In,Mod,Mul,String,Sub", ""},
			"RGBA":                {reflect.TypeOf((*q.RGBA)(nil)).Elem(), "", "At,Bounds,ColorModel,Opaque,PixOffset,RGBAAt,Set,SetRGBA,SubImage"},
			"RGBA64":              {reflect.TypeOf((*q.RGBA64)(nil)).Elem(), "", "At,Bounds,ColorModel,Opaque,PixOffset,RGBA64At,Set,SetRGBA64,SubImage"},
			"Rectangle":           {reflect.TypeOf((*q.Rectangle)(nil)).Elem(), "Add,At,Bounds,Canon,ColorModel,Dx,Dy,Empty,Eq,In,Inset,Intersect,Overlaps,Size,String,Sub,Union", ""},
			"Uniform":             {reflect.TypeOf((*q.Uniform)(nil)).Elem(), "", "At,Bounds,ColorModel,Convert,Opaque,RGBA"},
			"YCbCr":               {reflect.TypeOf((*q.YCbCr)(nil)).Elem(), "", "At,Bounds,COffset,ColorModel,Opaque,SubImage,YCbCrAt,YOffset"},
			"YCbCrSubsampleRatio": {reflect.TypeOf((*q.YCbCrSubsampleRatio)(nil)).Elem(), "String", ""},
		},
		AliasTypes: map[string]reflect.Type{},
		Vars: map[string]reflect.Value{
			"Black":       reflect.ValueOf(&q.Black),
			"ErrFormat":   reflect.ValueOf(&q.ErrFormat),
			"Opaque":      reflect.ValueOf(&q.Opaque),
			"Transparent": reflect.ValueOf(&q.Transparent),
			"White":       reflect.ValueOf(&q.White),
			"ZP":          reflect.ValueOf(&q.ZP),
			"ZR":          reflect.ValueOf(&q.ZR),
		},
		Funcs: map[string]reflect.Value{
			"Decode":         reflect.ValueOf(q.Decode),
			"DecodeConfig":   reflect.ValueOf(q.DecodeConfig),
			"NewAlpha":       reflect.ValueOf(q.NewAlpha),
			"NewAlpha16":     reflect.ValueOf(q.NewAlpha16),
			"NewCMYK":        reflect.ValueOf(q.NewCMYK),
			"NewGray":        reflect.ValueOf(q.NewGray),
			"NewGray16":      reflect.ValueOf(q.NewGray16),
			"NewNRGBA":       reflect.ValueOf(q.NewNRGBA),
			"NewNRGBA64":     reflect.ValueOf(q.NewNRGBA64),
			"NewNYCbCrA":     reflect.ValueOf(q.NewNYCbCrA),
			"NewPaletted":    reflect.ValueOf(q.NewPaletted),
			"NewRGBA":        reflect.ValueOf(q.NewRGBA),
			"NewRGBA64":      reflect.ValueOf(q.NewRGBA64),
			"NewUniform":     reflect.ValueOf(q.NewUniform),
			"NewYCbCr":       reflect.ValueOf(q.NewYCbCr),
			"Pt":             reflect.ValueOf(q.Pt),
			"Rect":           reflect.ValueOf(q.Rect),
			"RegisterFormat": reflect.ValueOf(q.RegisterFormat),
		},
		TypedConsts: map[string]igop.TypedConst{
			"YCbCrSubsampleRatio410": {reflect.TypeOf(q.YCbCrSubsampleRatio410), constant.MakeInt64(int64(q.YCbCrSubsampleRatio410))},
			"YCbCrSubsampleRatio411": {reflect.TypeOf(q.YCbCrSubsampleRatio411), constant.MakeInt64(int64(q.YCbCrSubsampleRatio411))},
			"YCbCrSubsampleRatio420": {reflect.TypeOf(q.YCbCrSubsampleRatio420), constant.MakeInt64(int64(q.YCbCrSubsampleRatio420))},
			"YCbCrSubsampleRatio422": {reflect.TypeOf(q.YCbCrSubsampleRatio422), constant.MakeInt64(int64(q.YCbCrSubsampleRatio422))},
			"YCbCrSubsampleRatio440": {reflect.TypeOf(q.YCbCrSubsampleRatio440), constant.MakeInt64(int64(q.YCbCrSubsampleRatio440))},
			"YCbCrSubsampleRatio444": {reflect.TypeOf(q.YCbCrSubsampleRatio444), constant.MakeInt64(int64(q.YCbCrSubsampleRatio444))},
		},
		UntypedConsts: map[string]igop.UntypedConst{},
	})
}
