// export by github.com/goplus/interp/cmd/qexp

// +build darwin freebsd netbsd openbsd

package syscall

import (
	"syscall"

	"github.com/goplus/interp"
)

func init() {
	interp.RegisterPackage("syscall", extMap_2256505852, typList_2256505852)
}

var extMap_2256505852 = map[string]interface{}{
	"syscall.Adjtime":              syscall.Adjtime,
	"syscall.BpfBuflen":            syscall.BpfBuflen,
	"syscall.BpfDatalink":          syscall.BpfDatalink,
	"syscall.BpfHeadercmpl":        syscall.BpfHeadercmpl,
	"syscall.BpfInterface":         syscall.BpfInterface,
	"syscall.BpfJump":              syscall.BpfJump,
	"syscall.BpfStats":             syscall.BpfStats,
	"syscall.BpfStmt":              syscall.BpfStmt,
	"syscall.BpfTimeout":           syscall.BpfTimeout,
	"syscall.CheckBpfVersion":      syscall.CheckBpfVersion,
	"syscall.Chflags":              syscall.Chflags,
	"syscall.Fchflags":             syscall.Fchflags,
	"syscall.FlushBpf":             syscall.FlushBpf,
	"syscall.Fpathconf":            syscall.Fpathconf,
	"syscall.Getdirentries":        syscall.Getdirentries,
	"syscall.Getsid":               syscall.Getsid,
	"syscall.GetsockoptByte":       syscall.GetsockoptByte,
	"syscall.Issetugid":            syscall.Issetugid,
	"syscall.Kevent":               syscall.Kevent,
	"syscall.Kqueue":               syscall.Kqueue,
	"syscall.ParseRoutingMessage":  syscall.ParseRoutingMessage,
	"syscall.ParseRoutingSockaddr": syscall.ParseRoutingSockaddr,
	"syscall.Pathconf":             syscall.Pathconf,
	"syscall.Revoke":               syscall.Revoke,
	"syscall.RouteRIB":             syscall.RouteRIB,
	"syscall.SetBpf":               syscall.SetBpf,
	"syscall.SetBpfBuflen":         syscall.SetBpfBuflen,
	"syscall.SetBpfDatalink":       syscall.SetBpfDatalink,
	"syscall.SetBpfHeadercmpl":     syscall.SetBpfHeadercmpl,
	"syscall.SetBpfImmediate":      syscall.SetBpfImmediate,
	"syscall.SetBpfInterface":      syscall.SetBpfInterface,
	"syscall.SetBpfPromisc":        syscall.SetBpfPromisc,
	"syscall.SetBpfTimeout":        syscall.SetBpfTimeout,
	"syscall.SetKevent":            syscall.SetKevent,
	"syscall.Sysctl":               syscall.Sysctl,
	"syscall.SysctlUint32":         syscall.SysctlUint32,
}

var typList_2256505852 = []interface{}{
	(*syscall.BpfHdr)(nil),
	(*syscall.BpfInsn)(nil),
	(*syscall.BpfProgram)(nil),
	(*syscall.BpfStat)(nil),
	(*syscall.BpfVersion)(nil),
	(*syscall.IfData)(nil),
	(*syscall.IfMsghdr)(nil),
	(*syscall.IfaMsghdr)(nil),
	(*syscall.InterfaceAddrMessage)(nil),
	(*syscall.InterfaceMessage)(nil),
	(*syscall.Kevent_t)(nil),
	(*syscall.RawSockaddrDatalink)(nil),
	(*syscall.RouteMessage)(nil),
	(*syscall.RtMetrics)(nil),
	(*syscall.RtMsghdr)(nil),
	(*syscall.SockaddrDatalink)(nil),
}
