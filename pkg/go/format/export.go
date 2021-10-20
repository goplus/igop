// export by github.com/goplus/gossa/cmd/qexp

package format

import (
	"go/format"

	"github.com/goplus/gossa"
)

func init() {
	gossa.RegisterPackage("go/format", extMap, typList)
}

var extMap = map[string]interface{}{
	"go/format.Node":   format.Node,
	"go/format.Source": format.Source,
}

var typList = []interface{}{}
