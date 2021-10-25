// export by github.com/goplus/gossa/cmd/qexp

//go:build darwin || freebsd || linux || openbsd
// +build darwin freebsd linux openbsd

package syscall

import (
	"syscall"

	"github.com/goplus/gossa"
)

func init() {
	gossa.RegisterPackage("syscall", extMap_1527460680, typList_1527460680)
}

var extMap_1527460680 = map[string]interface{}{
	"syscall.Fstatfs": syscall.Fstatfs,
	"syscall.Statfs":  syscall.Statfs,
}

var typList_1527460680 = []interface{}{}
