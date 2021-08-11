// export by github.com/goplus/interp/cmd/qexp

// +build darwin freebsd linux netbsd openbsd windows

package syscall

import (
	"syscall"

	"github.com/goplus/interp"
)

func init() {
	interp.RegisterPackage("syscall", extMap_525812225, typList_525812225)
}

var extMap_525812225 = map[string]interface{}{
	"syscall.Accept":              syscall.Accept,
	"syscall.Bind":                syscall.Bind,
	"syscall.Close":               syscall.Close,
	"syscall.CloseOnExec":         syscall.CloseOnExec,
	"syscall.Connect":             syscall.Connect,
	"syscall.Fchdir":              syscall.Fchdir,
	"syscall.Fchmod":              syscall.Fchmod,
	"syscall.Fchown":              syscall.Fchown,
	"syscall.Fsync":               syscall.Fsync,
	"syscall.Ftruncate":           syscall.Ftruncate,
	"syscall.Getpeername":         syscall.Getpeername,
	"syscall.Getsockname":         syscall.Getsockname,
	"syscall.GetsockoptInt":       syscall.GetsockoptInt,
	"syscall.Listen":              syscall.Listen,
	"syscall.Open":                syscall.Open,
	"syscall.Pipe":                syscall.Pipe,
	"syscall.Read":                syscall.Read,
	"syscall.Recvfrom":            syscall.Recvfrom,
	"syscall.Seek":                syscall.Seek,
	"syscall.Sendto":              syscall.Sendto,
	"syscall.SetNonblock":         syscall.SetNonblock,
	"syscall.SetsockoptIPMreq":    syscall.SetsockoptIPMreq,
	"syscall.SetsockoptIPv6Mreq":  syscall.SetsockoptIPv6Mreq,
	"syscall.SetsockoptInet4Addr": syscall.SetsockoptInet4Addr,
	"syscall.SetsockoptInt":       syscall.SetsockoptInt,
	"syscall.SetsockoptLinger":    syscall.SetsockoptLinger,
	"syscall.SetsockoptTimeval":   syscall.SetsockoptTimeval,
	"syscall.Shutdown":            syscall.Shutdown,
	"syscall.Socket":              syscall.Socket,
	"syscall.Stderr":              &syscall.Stderr,
	"syscall.Stdin":               &syscall.Stdin,
	"syscall.Stdout":              &syscall.Stdout,
	"syscall.Syscall":             syscall.Syscall,
	"syscall.Syscall6":            syscall.Syscall6,
	"syscall.Write":               syscall.Write,
}

var typList_525812225 = []interface{}{
	(*syscall.WaitStatus)(nil),
}
