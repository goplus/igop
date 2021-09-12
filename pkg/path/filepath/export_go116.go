// export by github.com/goplus/interp/cmd/qexp

//go:build go1.16
// +build go1.16

package filepath

import (
	"path/filepath"

	"github.com/goplus/interp"
)

func init() {
	interp.RegisterPackage("path/filepath", extMap_go116, typList_go116)
}

var extMap_go116 = map[string]interface{}{
	"path/filepath.WalkDir": filepath.WalkDir,
}

var typList_go116 = []interface{}{}
