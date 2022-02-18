// export by github.com/goplus/gossa/cmd/qexp

//+build go1.16,!go1.17

package color

import (
	q "image/color"

	"reflect"

	"github.com/goplus/gossa"
)

func init() {
	gossa.RegisterPackage(&gossa.Package{
		Name: "color",
		Path: "image/color",
		Deps: map[string]string{},
		Interfaces: map[string]reflect.Type{
			"Color": reflect.TypeOf((*q.Color)(nil)).Elem(),
			"Model": reflect.TypeOf((*q.Model)(nil)).Elem(),
		},
		NamedTypes: map[string]gossa.NamedType{
			"Alpha":   {reflect.TypeOf((*q.Alpha)(nil)).Elem(), "RGBA", ""},
			"Alpha16": {reflect.TypeOf((*q.Alpha16)(nil)).Elem(), "RGBA", ""},
			"CMYK":    {reflect.TypeOf((*q.CMYK)(nil)).Elem(), "RGBA", ""},
			"Gray":    {reflect.TypeOf((*q.Gray)(nil)).Elem(), "RGBA", ""},
			"Gray16":  {reflect.TypeOf((*q.Gray16)(nil)).Elem(), "RGBA", ""},
			"NRGBA":   {reflect.TypeOf((*q.NRGBA)(nil)).Elem(), "RGBA", ""},
			"NRGBA64": {reflect.TypeOf((*q.NRGBA64)(nil)).Elem(), "RGBA", ""},
			"NYCbCrA": {reflect.TypeOf((*q.NYCbCrA)(nil)).Elem(), "RGBA", ""},
			"Palette": {reflect.TypeOf((*q.Palette)(nil)).Elem(), "Convert,Index", ""},
			"RGBA":    {reflect.TypeOf((*q.RGBA)(nil)).Elem(), "RGBA", ""},
			"RGBA64":  {reflect.TypeOf((*q.RGBA64)(nil)).Elem(), "RGBA", ""},
			"YCbCr":   {reflect.TypeOf((*q.YCbCr)(nil)).Elem(), "RGBA", ""},
		},
		AliasTypes: map[string]reflect.Type{},
		Vars: map[string]reflect.Value{
			"Alpha16Model": reflect.ValueOf(&q.Alpha16Model),
			"AlphaModel":   reflect.ValueOf(&q.AlphaModel),
			"Black":        reflect.ValueOf(&q.Black),
			"CMYKModel":    reflect.ValueOf(&q.CMYKModel),
			"Gray16Model":  reflect.ValueOf(&q.Gray16Model),
			"GrayModel":    reflect.ValueOf(&q.GrayModel),
			"NRGBA64Model": reflect.ValueOf(&q.NRGBA64Model),
			"NRGBAModel":   reflect.ValueOf(&q.NRGBAModel),
			"NYCbCrAModel": reflect.ValueOf(&q.NYCbCrAModel),
			"Opaque":       reflect.ValueOf(&q.Opaque),
			"RGBA64Model":  reflect.ValueOf(&q.RGBA64Model),
			"RGBAModel":    reflect.ValueOf(&q.RGBAModel),
			"Transparent":  reflect.ValueOf(&q.Transparent),
			"White":        reflect.ValueOf(&q.White),
			"YCbCrModel":   reflect.ValueOf(&q.YCbCrModel),
		},
		Funcs: map[string]reflect.Value{
			"CMYKToRGB":  reflect.ValueOf(q.CMYKToRGB),
			"ModelFunc":  reflect.ValueOf(q.ModelFunc),
			"RGBToCMYK":  reflect.ValueOf(q.RGBToCMYK),
			"RGBToYCbCr": reflect.ValueOf(q.RGBToYCbCr),
			"YCbCrToRGB": reflect.ValueOf(q.YCbCrToRGB),
		},
		TypedConsts:   map[string]gossa.TypedConst{},
		UntypedConsts: map[string]gossa.UntypedConst{},
	})
}