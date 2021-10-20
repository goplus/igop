// export by github.com/goplus/gossa/cmd/qexp

//go:build go1.16
// +build go1.16

package filepath

import (
	"path/filepath"

	"github.com/goplus/gossa"
)

func init() {
	gossa.RegisterPackage("path/filepath", extMap_go116, typList_go116)
}

var extMap_go116 = map[string]interface{}{
	"path/filepath.WalkDir": filepath.WalkDir,
}

var typList_go116 = []interface{}{}
