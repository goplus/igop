// export by github.com/goplus/gossa/cmd/qexp

package palette

import (
	"image/color/palette"

	"github.com/goplus/gossa"
)

func init() {
	gossa.RegisterPackage("image/color/palette", extMap, typList)
}

var extMap = map[string]interface{}{
	"image/color/palette.Plan9":   &palette.Plan9,
	"image/color/palette.WebSafe": &palette.WebSafe,
}

var typList = []interface{}{}
