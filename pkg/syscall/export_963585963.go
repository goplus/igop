// export by github.com/goplus/interp/cmd/qexp

//go:build darwin || freebsd || linux
// +build darwin freebsd linux

package syscall

import (
	"syscall"

	"github.com/goplus/interp"
)

func init() {
	interp.RegisterPackage("syscall", extMap_963585963, typList_963585963)
}

var extMap_963585963 = map[string]interface{}{
	"syscall.NsecToTimespec": syscall.NsecToTimespec,
	"syscall.TimespecToNsec": syscall.TimespecToNsec,
}

var typList_963585963 = []interface{}{
	(*syscall.RawSockaddrInet6)(nil),
}
