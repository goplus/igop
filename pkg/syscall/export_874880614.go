// export by github.com/goplus/gossa/cmd/qexp

//go:build darwin || freebsd || openbsd
// +build darwin freebsd openbsd

package syscall

import (
	"syscall"

	"github.com/goplus/gossa"
)

func init() {
	gossa.RegisterPackage("syscall", extMap_874880614, typList_874880614)
}

var extMap_874880614 = map[string]interface{}{
	"syscall.Getfsstat": syscall.Getfsstat,
	"syscall.Setlogin":  syscall.Setlogin,
}

var typList_874880614 = []interface{}{}
