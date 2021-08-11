// export by github.com/goplus/interp/cmd/qexp

package html

import (
	"html"

	"github.com/goplus/interp"
)

func init() {
	interp.RegisterPackage("html", extMap, typList)
}

var extMap = map[string]interface{}{
	"html.EscapeString":   html.EscapeString,
	"html.UnescapeString": html.UnescapeString,
}

var typList = []interface{}{}
