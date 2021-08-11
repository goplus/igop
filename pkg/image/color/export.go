// export by github.com/goplus/interp/cmd/qexp

package color

import (
	"image/color"

	"github.com/goplus/interp"
)

func init() {
	interp.RegisterPackage("image/color", extMap, typList)
}

var extMap = map[string]interface{}{
	"(image/color.Alpha).RGBA":      (color.Alpha).RGBA,
	"(image/color.Alpha16).RGBA":    (color.Alpha16).RGBA,
	"(image/color.CMYK).RGBA":       (color.CMYK).RGBA,
	"(image/color.Color).RGBA":      (color.Color).RGBA,
	"(image/color.Gray).RGBA":       (color.Gray).RGBA,
	"(image/color.Gray16).RGBA":     (color.Gray16).RGBA,
	"(image/color.Model).Convert":   (color.Model).Convert,
	"(image/color.NRGBA).RGBA":      (color.NRGBA).RGBA,
	"(image/color.NRGBA64).RGBA":    (color.NRGBA64).RGBA,
	"(image/color.NYCbCrA).RGBA":    (color.NYCbCrA).RGBA,
	"(image/color.Palette).Convert": (color.Palette).Convert,
	"(image/color.Palette).Index":   (color.Palette).Index,
	"(image/color.RGBA).RGBA":       (color.RGBA).RGBA,
	"(image/color.RGBA64).RGBA":     (color.RGBA64).RGBA,
	"(image/color.YCbCr).RGBA":      (color.YCbCr).RGBA,
	"image/color.Alpha16Model":      &color.Alpha16Model,
	"image/color.AlphaModel":        &color.AlphaModel,
	"image/color.Black":             &color.Black,
	"image/color.CMYKModel":         &color.CMYKModel,
	"image/color.CMYKToRGB":         color.CMYKToRGB,
	"image/color.Gray16Model":       &color.Gray16Model,
	"image/color.GrayModel":         &color.GrayModel,
	"image/color.ModelFunc":         color.ModelFunc,
	"image/color.NRGBA64Model":      &color.NRGBA64Model,
	"image/color.NRGBAModel":        &color.NRGBAModel,
	"image/color.NYCbCrAModel":      &color.NYCbCrAModel,
	"image/color.Opaque":            &color.Opaque,
	"image/color.RGBA64Model":       &color.RGBA64Model,
	"image/color.RGBAModel":         &color.RGBAModel,
	"image/color.RGBToCMYK":         color.RGBToCMYK,
	"image/color.RGBToYCbCr":        color.RGBToYCbCr,
	"image/color.Transparent":       &color.Transparent,
	"image/color.White":             &color.White,
	"image/color.YCbCrModel":        &color.YCbCrModel,
	"image/color.YCbCrToRGB":        color.YCbCrToRGB,
}

var typList = []interface{}{
	(*color.Alpha)(nil),
	(*color.Alpha16)(nil),
	(*color.CMYK)(nil),
	(*color.Color)(nil),
	(*color.Gray)(nil),
	(*color.Gray16)(nil),
	(*color.Model)(nil),
	(*color.NRGBA)(nil),
	(*color.NRGBA64)(nil),
	(*color.NYCbCrA)(nil),
	(*color.Palette)(nil),
	(*color.RGBA)(nil),
	(*color.RGBA64)(nil),
	(*color.YCbCr)(nil),
}
