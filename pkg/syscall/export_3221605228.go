// export by github.com/goplus/interp/cmd/qexp

//go:build freebsd || linux
// +build freebsd linux

package syscall

import (
	"syscall"

	"github.com/goplus/interp"
)

func init() {
	interp.RegisterPackage("syscall", extMap_3221605228, typList_3221605228)
}

var extMap_3221605228 = map[string]interface{}{
	"syscall.GetsockoptIPMreqn": syscall.GetsockoptIPMreqn,
	"syscall.SetsockoptIPMreqn": syscall.SetsockoptIPMreqn,
}

var typList_3221605228 = []interface{}{
	(*syscall.IPMreqn)(nil),
}
