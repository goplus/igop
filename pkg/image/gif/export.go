// export by github.com/goplus/gossa/cmd/qexp

package gif

import (
	"image/gif"

	"github.com/goplus/gossa"
)

func init() {
	gossa.RegisterPackage("image/gif", extMap, typList)
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
