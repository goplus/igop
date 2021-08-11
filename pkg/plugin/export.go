// export by github.com/goplus/interp/cmd/qexp

package plugin

import (
	"plugin"

	"github.com/goplus/interp"
)

func init() {
	interp.RegisterPackage("plugin", extMap, typList)
}

var extMap = map[string]interface{}{
	"(*plugin.Plugin).Lookup": (*plugin.Plugin).Lookup,
	"plugin.Open":             plugin.Open,
}

var typList = []interface{}{
	(*plugin.Plugin)(nil),
	(*plugin.Symbol)(nil),
}
