// export by github.com/goplus/gossa/cmd/qexp

package plugin

import (
	"plugin"

	"github.com/goplus/gossa"
)

func init() {
	gossa.RegisterPackage("plugin", extMap, typList)
}

var extMap = map[string]interface{}{
	"(*plugin.Plugin).Lookup": (*plugin.Plugin).Lookup,
	"plugin.Open":             plugin.Open,
}

var typList = []interface{}{
	(*plugin.Plugin)(nil),
	(*plugin.Symbol)(nil),
}
