// export by github.com/goplus/interp/cmd/qexp

package draw

import (
	"image/draw"

	"github.com/goplus/interp"
)

func init() {
	interp.RegisterPackage("image/draw", extMap, typList)
}

var extMap = map[string]interface{}{
	"(image/draw.Drawer).Draw":        (draw.Drawer).Draw,
	"(image/draw.Image).At":           (draw.Image).At,
	"(image/draw.Image).Bounds":       (draw.Image).Bounds,
	"(image/draw.Image).ColorModel":   (draw.Image).ColorModel,
	"(image/draw.Image).Set":          (draw.Image).Set,
	"(image/draw.Op).Draw":            (draw.Op).Draw,
	"(image/draw.Quantizer).Quantize": (draw.Quantizer).Quantize,
	"image/draw.Draw":                 draw.Draw,
	"image/draw.DrawMask":             draw.DrawMask,
	"image/draw.FloydSteinberg":       &draw.FloydSteinberg,
}

var typList = []interface{}{
	(*draw.Drawer)(nil),
	(*draw.Image)(nil),
	(*draw.Op)(nil),
	(*draw.Quantizer)(nil),
}
