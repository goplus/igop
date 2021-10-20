// export by github.com/goplus/gossa/cmd/qexp

package jpeg

import (
	"image/jpeg"

	"github.com/goplus/gossa"
)

func init() {
	gossa.RegisterPackage("image/jpeg", extMap, typList)
}

var extMap = map[string]interface{}{
	"(image/jpeg.FormatError).Error":      (jpeg.FormatError).Error,
	"(image/jpeg.Reader).Read":            (jpeg.Reader).Read,
	"(image/jpeg.Reader).ReadByte":        (jpeg.Reader).ReadByte,
	"(image/jpeg.UnsupportedError).Error": (jpeg.UnsupportedError).Error,
	"image/jpeg.Decode":                   jpeg.Decode,
	"image/jpeg.DecodeConfig":             jpeg.DecodeConfig,
	"image/jpeg.Encode":                   jpeg.Encode,
}

var typList = []interface{}{
	(*jpeg.FormatError)(nil),
	(*jpeg.Options)(nil),
	(*jpeg.Reader)(nil),
	(*jpeg.UnsupportedError)(nil),
}
