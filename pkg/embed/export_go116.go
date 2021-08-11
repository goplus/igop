// export by github.com/goplus/interp/cmd/qexp

// +build go1.16

package embed

import (
	"embed"

	"github.com/goplus/interp"
)

func init() {
	interp.RegisterPackage("embed", extMap_go116, typList_go116)
}

var extMap_go116 = map[string]interface{}{
	"(embed.FS).Open":     (embed.FS).Open,
	"(embed.FS).ReadDir":  (embed.FS).ReadDir,
	"(embed.FS).ReadFile": (embed.FS).ReadFile,
}

var typList_go116 = []interface{}{
	(*embed.FS)(nil),
}
