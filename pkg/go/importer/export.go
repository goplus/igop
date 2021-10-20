// export by github.com/goplus/gossa/cmd/qexp

package importer

import (
	"go/importer"

	"github.com/goplus/gossa"
)

func init() {
	gossa.RegisterPackage("go/importer", extMap, typList)
}

var extMap = map[string]interface{}{
	"go/importer.Default":     importer.Default,
	"go/importer.For":         importer.For,
	"go/importer.ForCompiler": importer.ForCompiler,
}

var typList = []interface{}{
	(*importer.Lookup)(nil),
}
