// export by github.com/goplus/interp/cmd/qexp

package gif

import (
	"image/gif"

	"github.com/goplus/interp"
)

func init() {
	interp.RegisterPackage("image/gif", extMap, typList)
}

var extMap = map[string]interface{}{
	"image/gif.Decode":       gif.Decode,
	"image/gif.DecodeAll":    gif.DecodeAll,
	"image/gif.DecodeConfig": gif.DecodeConfig,
	"image/gif.Encode":       gif.Encode,
	"image/gif.EncodeAll":    gif.EncodeAll,
}

var typList = []interface{}{
	(*gif.GIF)(nil),
	(*gif.Options)(nil),
}
