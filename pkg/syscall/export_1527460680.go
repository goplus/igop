// export by github.com/goplus/interp/cmd/qexp

// +build darwin freebsd linux openbsd

package syscall

import (
	"syscall"

	"github.com/goplus/interp"
)

func init() {
	interp.RegisterPackage("syscall", extMap_1527460680, typList_1527460680)
}

var extMap_1527460680 = map[string]interface{}{
	"syscall.Fstatfs": syscall.Fstatfs,
	"syscall.Statfs":  syscall.Statfs,
}

var typList_1527460680 = []interface{}{}
