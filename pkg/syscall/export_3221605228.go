// export by github.com/goplus/gossa/cmd/qexp

//go:build freebsd || linux
// +build freebsd linux

package syscall

import (
	"syscall"

	"github.com/goplus/gossa"
)

func init() {
	gossa.RegisterPackage("syscall", extMap_3221605228, typList_3221605228)
}

var extMap_3221605228 = map[string]interface{}{
	"syscall.GetsockoptIPMreqn": syscall.GetsockoptIPMreqn,
	"syscall.SetsockoptIPMreqn": syscall.SetsockoptIPMreqn,
}

var typList_3221605228 = []interface{}{
	(*syscall.IPMreqn)(nil),
}
