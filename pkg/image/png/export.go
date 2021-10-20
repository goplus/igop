// export by github.com/goplus/gossa/cmd/qexp

package png

import (
	"image/png"

	"github.com/goplus/gossa"
)

func init() {
	gossa.RegisterPackage("image/png", extMap, typList)
}

var extMap = map[string]interface{}{
	"(*image/png.Encoder).Encode":        (*png.Encoder).Encode,
	"(image/png.EncoderBufferPool).Get":  (png.EncoderBufferPool).Get,
	"(image/png.EncoderBufferPool).Put":  (png.EncoderBufferPool).Put,
	"(image/png.FormatError).Error":      (png.FormatError).Error,
	"(image/png.UnsupportedError).Error": (png.UnsupportedError).Error,
	"image/png.Decode":                   png.Decode,
	"image/png.DecodeConfig":             png.DecodeConfig,
	"image/png.Encode":                   png.Encode,
}

var typList = []interface{}{
	(*png.CompressionLevel)(nil),
	(*png.Encoder)(nil),
	(*png.EncoderBuffer)(nil),
	(*png.EncoderBufferPool)(nil),
	(*png.FormatError)(nil),
	(*png.UnsupportedError)(nil),
}
