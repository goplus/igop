// export by github.com/goplus/interp/cmd/qexp

// +build go1.16

package fstest

import (
	"testing/fstest"

	"github.com/goplus/interp"
)

func init() {
	interp.RegisterPackage("testing/fstest", extMap_go116, typList_go116)
}

var extMap_go116 = map[string]interface{}{
	"(testing/fstest.MapFS).Glob":     (fstest.MapFS).Glob,
	"(testing/fstest.MapFS).Open":     (fstest.MapFS).Open,
	"(testing/fstest.MapFS).ReadDir":  (fstest.MapFS).ReadDir,
	"(testing/fstest.MapFS).ReadFile": (fstest.MapFS).ReadFile,
	"(testing/fstest.MapFS).Stat":     (fstest.MapFS).Stat,
	"(testing/fstest.MapFS).Sub":      (fstest.MapFS).Sub,
	"testing/fstest.TestFS":           fstest.TestFS,
}

var typList_go116 = []interface{}{
	(*fstest.MapFS)(nil),
	(*fstest.MapFile)(nil),
}
