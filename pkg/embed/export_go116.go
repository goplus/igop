// export by github.com/goplus/gossa/cmd/qexp

//go:build go1.16
// +build go1.16

package embed

import (
	"embed"

	"github.com/goplus/gossa"
)

func init() {
	gossa.RegisterPackage("embed", extMap_go116, typList_go116)
}

var extMap_go116 = map[string]interface{}{
	"(embed.FS).Open":     (embed.FS).Open,
	"(embed.FS).ReadDir":  (embed.FS).ReadDir,
	"(embed.FS).ReadFile": (embed.FS).ReadFile,
}

var typList_go116 = []interface{}{
	(*embed.FS)(nil),
}
