// export by github.com/goplus/interp/cmd/qexp

// +build go1.16

package os

import (
	"os"

	"github.com/goplus/interp"
)

func init() {
	interp.RegisterPackage("os", extMap_go116, typList_go116)
}

var extMap_go116 = map[string]interface{}{
	"(*os.File).ReadDir": (*os.File).ReadDir,
	"os.CreateTemp":      os.CreateTemp,
	"os.DirFS":           os.DirFS,
	"os.ErrProcessDone":  &os.ErrProcessDone,
	"os.MkdirTemp":       os.MkdirTemp,
	"os.ReadDir":         os.ReadDir,
	"os.ReadFile":        os.ReadFile,
	"os.WriteFile":       os.WriteFile,
}

var typList_go116 = []interface{}{
	(*os.DirEntry)(nil),
}
