// export by github.com/goplus/interp/cmd/qexp

package syscall

import (
	"syscall"

	"github.com/goplus/interp"
)

func init() {
	interp.RegisterPackage("syscall", extMap, typList)
}

var extMap = map[string]interface{}{
	"(*syscall.Timespec).Nano":        (*syscall.Timespec).Nano,
	"(*syscall.Timespec).Unix":        (*syscall.Timespec).Unix,
	"(*syscall.Timeval).Nano":         (*syscall.Timeval).Nano,
	"(*syscall.Timeval).Unix":         (*syscall.Timeval).Unix,
	"(syscall.Conn).SyscallConn":      (syscall.Conn).SyscallConn,
	"(syscall.Errno).Error":           (syscall.Errno).Error,
	"(syscall.Errno).Is":              (syscall.Errno).Is,
	"(syscall.Errno).Temporary":       (syscall.Errno).Temporary,
	"(syscall.Errno).Timeout":         (syscall.Errno).Timeout,
	"(syscall.RawConn).Control":       (syscall.RawConn).Control,
	"(syscall.RawConn).Read":          (syscall.RawConn).Read,
	"(syscall.RawConn).Write":         (syscall.RawConn).Write,
	"(syscall.Signal).Signal":         (syscall.Signal).Signal,
	"(syscall.Signal).String":         (syscall.Signal).String,
	"(syscall.WaitStatus).Continued":  (syscall.WaitStatus).Continued,
	"(syscall.WaitStatus).CoreDump":   (syscall.WaitStatus).CoreDump,
	"(syscall.WaitStatus).ExitStatus": (syscall.WaitStatus).ExitStatus,
	"(syscall.WaitStatus).Exited":     (syscall.WaitStatus).Exited,
	"(syscall.WaitStatus).Signal":     (syscall.WaitStatus).Signal,
	"(syscall.WaitStatus).Signaled":   (syscall.WaitStatus).Signaled,
	"(syscall.WaitStatus).StopSignal": (syscall.WaitStatus).StopSignal,
	"(syscall.WaitStatus).Stopped":    (syscall.WaitStatus).Stopped,
	"(syscall.WaitStatus).TrapCause":  (syscall.WaitStatus).TrapCause,
	"syscall.BytePtrFromString":       syscall.BytePtrFromString,
	"syscall.ByteSliceFromString":     syscall.ByteSliceFromString,
	"syscall.Chdir":                   syscall.Chdir,
	"syscall.Chmod":                   syscall.Chmod,
	"syscall.Chown":                   syscall.Chown,
	"syscall.Clearenv":                syscall.Clearenv,
	"syscall.Environ":                 syscall.Environ,
	"syscall.Exec":                    syscall.Exec,
	"syscall.Exit":                    syscall.Exit,
	"syscall.ForkLock":                &syscall.ForkLock,
	"syscall.Getegid":                 syscall.Getegid,
	"syscall.Getenv":                  syscall.Getenv,
	"syscall.Geteuid":                 syscall.Geteuid,
	"syscall.Getgid":                  syscall.Getgid,
	"syscall.Getgroups":               syscall.Getgroups,
	"syscall.Getpagesize":             syscall.Getpagesize,
	"syscall.Getpid":                  syscall.Getpid,
	"syscall.Getppid":                 syscall.Getppid,
	"syscall.Gettimeofday":            syscall.Gettimeofday,
	"syscall.Getuid":                  syscall.Getuid,
	"syscall.Getwd":                   syscall.Getwd,
	"syscall.Lchown":                  syscall.Lchown,
	"syscall.Link":                    syscall.Link,
	"syscall.Mkdir":                   syscall.Mkdir,
	"syscall.NsecToTimeval":           syscall.NsecToTimeval,
	"syscall.Readlink":                syscall.Readlink,
	"syscall.Rename":                  syscall.Rename,
	"syscall.Rmdir":                   syscall.Rmdir,
	"syscall.Setenv":                  syscall.Setenv,
	"syscall.SocketDisableIPv6":       &syscall.SocketDisableIPv6,
	"syscall.StartProcess":            syscall.StartProcess,
	"syscall.StringBytePtr":           syscall.StringBytePtr,
	"syscall.StringByteSlice":         syscall.StringByteSlice,
	"syscall.Symlink":                 syscall.Symlink,
	"syscall.Unlink":                  syscall.Unlink,
	"syscall.Unsetenv":                syscall.Unsetenv,
	"syscall.Utimes":                  syscall.Utimes,
	"syscall.UtimesNano":              syscall.UtimesNano,
}

var typList = []interface{}{
	(*syscall.Conn)(nil),
	(*syscall.Errno)(nil),
	(*syscall.IPMreq)(nil),
	(*syscall.IPv6Mreq)(nil),
	(*syscall.Linger)(nil),
	(*syscall.ProcAttr)(nil),
	(*syscall.RawConn)(nil),
	(*syscall.RawSockaddr)(nil),
	(*syscall.RawSockaddrAny)(nil),
	(*syscall.RawSockaddrInet4)(nil),
	(*syscall.Rusage)(nil),
	(*syscall.Signal)(nil),
	(*syscall.SockaddrInet4)(nil),
	(*syscall.SockaddrInet6)(nil),
	(*syscall.SockaddrUnix)(nil),
	(*syscall.SysProcAttr)(nil),
	(*syscall.Timespec)(nil),
	(*syscall.Timeval)(nil),
}
