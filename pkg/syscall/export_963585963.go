// export by github.com/goplus/gossa/cmd/qexp

//go:build darwin || freebsd || linux
// +build darwin freebsd linux

package syscall

import (
	"syscall"

	"github.com/goplus/gossa"
)

func init() {
	gossa.RegisterPackage("syscall", extMap_963585963, typList_963585963)
}

var extMap_963585963 = map[string]interface{}{
	"syscall.NsecToTimespec": syscall.NsecToTimespec,
	"syscall.TimespecToNsec": syscall.TimespecToNsec,
}

var typList_963585963 = []interface{}{
	(*syscall.RawSockaddrInet6)(nil),
}
