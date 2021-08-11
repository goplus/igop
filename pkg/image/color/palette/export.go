// export by github.com/goplus/interp/cmd/qexp

package palette

import (
	"image/color/palette"

	"github.com/goplus/interp"
)

func init() {
	interp.RegisterPackage("image/color/palette", extMap, typList)
}

var extMap = map[string]interface{}{
	"image/color/palette.Plan9":   &palette.Plan9,
	"image/color/palette.WebSafe": &palette.WebSafe,
}

var typList = []interface{}{}
