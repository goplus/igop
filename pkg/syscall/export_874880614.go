// export by github.com/goplus/interp/cmd/qexp

// +build darwin freebsd openbsd

package syscall

import (
	"syscall"

	"github.com/goplus/interp"
)

func init() {
	interp.RegisterPackage("syscall", extMap_874880614, typList_874880614)
}

var extMap_874880614 = map[string]interface{}{
	"syscall.Getfsstat": syscall.Getfsstat,
	"syscall.Setlogin":  syscall.Setlogin,
}

var typList_874880614 = []interface{}{}
