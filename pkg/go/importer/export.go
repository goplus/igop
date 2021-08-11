// export by github.com/goplus/interp/cmd/qexp

package importer

import (
	"go/importer"

	"github.com/goplus/interp"
)

func init() {
	interp.RegisterPackage("go/importer", extMap, typList)
}

var extMap = map[string]interface{}{
	"go/importer.Default":     importer.Default,
	"go/importer.For":         importer.For,
	"go/importer.ForCompiler": importer.ForCompiler,
}

var typList = []interface{}{
	(*importer.Lookup)(nil),
}
