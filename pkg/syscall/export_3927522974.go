// export by github.com/goplus/interp/cmd/qexp

//go:build darwin || freebsd || linux || netbsd || openbsd
// +build darwin freebsd linux netbsd openbsd

package syscall

import (
	"syscall"

	"github.com/goplus/interp"
)

func init() {
	interp.RegisterPackage("syscall", extMap_3927522974, typList_3927522974)
}

var extMap_3927522974 = map[string]interface{}{
	"(*syscall.Cmsghdr).SetLen":         (*syscall.Cmsghdr).SetLen,
	"(*syscall.Iovec).SetLen":           (*syscall.Iovec).SetLen,
	"(*syscall.Msghdr).SetControllen":   (*syscall.Msghdr).SetControllen,
	"syscall.Access":                    syscall.Access,
	"syscall.Chroot":                    syscall.Chroot,
	"syscall.CmsgLen":                   syscall.CmsgLen,
	"syscall.CmsgSpace":                 syscall.CmsgSpace,
	"syscall.Dup":                       syscall.Dup,
	"syscall.Dup2":                      syscall.Dup2,
	"syscall.FcntlFlock":                syscall.FcntlFlock,
	"syscall.Flock":                     syscall.Flock,
	"syscall.ForkExec":                  syscall.ForkExec,
	"syscall.Fstat":                     syscall.Fstat,
	"syscall.Futimes":                   syscall.Futimes,
	"syscall.Getpgid":                   syscall.Getpgid,
	"syscall.Getpgrp":                   syscall.Getpgrp,
	"syscall.Getpriority":               syscall.Getpriority,
	"syscall.Getrlimit":                 syscall.Getrlimit,
	"syscall.Getrusage":                 syscall.Getrusage,
	"syscall.GetsockoptICMPv6Filter":    syscall.GetsockoptICMPv6Filter,
	"syscall.GetsockoptIPMreq":          syscall.GetsockoptIPMreq,
	"syscall.GetsockoptIPv6MTUInfo":     syscall.GetsockoptIPv6MTUInfo,
	"syscall.GetsockoptIPv6Mreq":        syscall.GetsockoptIPv6Mreq,
	"syscall.GetsockoptInet4Addr":       syscall.GetsockoptInet4Addr,
	"syscall.Kill":                      syscall.Kill,
	"syscall.Lstat":                     syscall.Lstat,
	"syscall.Mkfifo":                    syscall.Mkfifo,
	"syscall.Mknod":                     syscall.Mknod,
	"syscall.Mmap":                      syscall.Mmap,
	"syscall.Munmap":                    syscall.Munmap,
	"syscall.ParseDirent":               syscall.ParseDirent,
	"syscall.ParseSocketControlMessage": syscall.ParseSocketControlMessage,
	"syscall.ParseUnixRights":           syscall.ParseUnixRights,
	"syscall.Pread":                     syscall.Pread,
	"syscall.Pwrite":                    syscall.Pwrite,
	"syscall.RawSyscall":                syscall.RawSyscall,
	"syscall.RawSyscall6":               syscall.RawSyscall6,
	"syscall.ReadDirent":                syscall.ReadDirent,
	"syscall.Recvmsg":                   syscall.Recvmsg,
	"syscall.Select":                    syscall.Select,
	"syscall.Sendfile":                  syscall.Sendfile,
	"syscall.Sendmsg":                   syscall.Sendmsg,
	"syscall.SendmsgN":                  syscall.SendmsgN,
	"syscall.Setgid":                    syscall.Setgid,
	"syscall.Setgroups":                 syscall.Setgroups,
	"syscall.Setpgid":                   syscall.Setpgid,
	"syscall.Setpriority":               syscall.Setpriority,
	"syscall.Setregid":                  syscall.Setregid,
	"syscall.Setreuid":                  syscall.Setreuid,
	"syscall.Setrlimit":                 syscall.Setrlimit,
	"syscall.Setsid":                    syscall.Setsid,
	"syscall.SetsockoptByte":            syscall.SetsockoptByte,
	"syscall.SetsockoptICMPv6Filter":    syscall.SetsockoptICMPv6Filter,
	"syscall.SetsockoptString":          syscall.SetsockoptString,
	"syscall.Settimeofday":              syscall.Settimeofday,
	"syscall.Setuid":                    syscall.Setuid,
	"syscall.SlicePtrFromStrings":       syscall.SlicePtrFromStrings,
	"syscall.Socketpair":                syscall.Socketpair,
	"syscall.Stat":                      syscall.Stat,
	"syscall.StringSlicePtr":            syscall.StringSlicePtr,
	"syscall.Sync":                      syscall.Sync,
	"syscall.TimevalToNsec":             syscall.TimevalToNsec,
	"syscall.Truncate":                  syscall.Truncate,
	"syscall.Umask":                     syscall.Umask,
	"syscall.UnixRights":                syscall.UnixRights,
	"syscall.Unmount":                   syscall.Unmount,
	"syscall.Wait4":                     syscall.Wait4,
}

var typList_3927522974 = []interface{}{
	(*syscall.Cmsghdr)(nil),
	(*syscall.Credential)(nil),
	(*syscall.Dirent)(nil),
	(*syscall.FdSet)(nil),
	(*syscall.Flock_t)(nil),
	(*syscall.Fsid)(nil),
	(*syscall.ICMPv6Filter)(nil),
	(*syscall.IPv6MTUInfo)(nil),
	(*syscall.Inet6Pktinfo)(nil),
	(*syscall.Iovec)(nil),
	(*syscall.Msghdr)(nil),
	(*syscall.RawSockaddrUnix)(nil),
	(*syscall.Rlimit)(nil),
	(*syscall.SocketControlMessage)(nil),
	(*syscall.Stat_t)(nil),
	(*syscall.Statfs_t)(nil),
	(*syscall.Termios)(nil),
}
