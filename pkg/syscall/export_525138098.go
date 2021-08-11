// export by github.com/goplus/interp/cmd/qexp

// +build darwin linux

package syscall

import (
	"syscall"

	"github.com/goplus/interp"
)

func init() {
	interp.RegisterPackage("syscall", extMap_525138098, typList_525138098)
}

var extMap_525138098 = map[string]interface{}{
	"syscall.Mlock":        syscall.Mlock,
	"syscall.Mlockall":     syscall.Mlockall,
	"syscall.Mprotect":     syscall.Mprotect,
	"syscall.Munlock":      syscall.Munlock,
	"syscall.Munlockall":   syscall.Munlockall,
	"syscall.PtraceAttach": syscall.PtraceAttach,
	"syscall.PtraceDetach": syscall.PtraceDetach,
}

var typList_525138098 = []interface{}{
	(*syscall.Inet4Pktinfo)(nil),
}
