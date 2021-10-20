// export by github.com/goplus/gossa/cmd/qexp

package html

import (
	"html"

	"github.com/goplus/gossa"
)

func init() {
	gossa.RegisterPackage("html", extMap, typList)
}

var extMap = map[string]interface{}{
	"html.EscapeString":   html.EscapeString,
	"html.UnescapeString": html.UnescapeString,
}

var typList = []interface{}{}
