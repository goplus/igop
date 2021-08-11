// export by github.com/goplus/interp/cmd/qexp

package format

import (
	"go/format"

	"github.com/goplus/interp"
)

func init() {
	interp.RegisterPackage("go/format", extMap, typList)
}

var extMap = map[string]interface{}{
	"go/format.Node":   format.Node,
	"go/format.Source": format.Source,
}

var typList = []interface{}{}
