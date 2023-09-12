// export by github.com/goplus/igop/cmd/qexp

//go:build go1.20 && !go1.20
// +build go1.20,!go1.20

package syscall

import (
	q "syscall"

	"go/constant"
	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "syscall",
		Path: "syscall",
		Deps: map[string]string{
			"errors":           "errors",
			"internal/bytealg": "bytealg",
			"internal/itoa":    "itoa",
			"internal/oserror": "oserror",
			"runtime":          "runtime",
			"sync":             "sync",
			"syscall/js":       "js",
			"unsafe":           "unsafe",
		},
		Interfaces: map[string]reflect.Type{
			"Conn":     reflect.TypeOf((*q.Conn)(nil)).Elem(),
			"RawConn":  reflect.TypeOf((*q.RawConn)(nil)).Elem(),
			"Sockaddr": reflect.TypeOf((*q.Sockaddr)(nil)).Elem(),
		},
		NamedTypes: map[string]reflect.Type{
			"Dirent":        reflect.TypeOf((*q.Dirent)(nil)).Elem(),
			"Errno":         reflect.TypeOf((*q.Errno)(nil)).Elem(),
			"Iovec":         reflect.TypeOf((*q.Iovec)(nil)).Elem(),
			"ProcAttr":      reflect.TypeOf((*q.ProcAttr)(nil)).Elem(),
			"Rusage":        reflect.TypeOf((*q.Rusage)(nil)).Elem(),
			"Signal":        reflect.TypeOf((*q.Signal)(nil)).Elem(),
			"SockaddrInet4": reflect.TypeOf((*q.SockaddrInet4)(nil)).Elem(),
			"SockaddrInet6": reflect.TypeOf((*q.SockaddrInet6)(nil)).Elem(),
			"SockaddrUnix":  reflect.TypeOf((*q.SockaddrUnix)(nil)).Elem(),
			"Stat_t":        reflect.TypeOf((*q.Stat_t)(nil)).Elem(),
			"SysProcAttr":   reflect.TypeOf((*q.SysProcAttr)(nil)).Elem(),
			"Timespec":      reflect.TypeOf((*q.Timespec)(nil)).Elem(),
			"Timeval":       reflect.TypeOf((*q.Timeval)(nil)).Elem(),
			"WaitStatus":    reflect.TypeOf((*q.WaitStatus)(nil)).Elem(),
		},
		AliasTypes: map[string]reflect.Type{},
		Vars: map[string]reflect.Value{
			"ForkLock": reflect.ValueOf(&q.ForkLock),
		},
		Funcs: map[string]reflect.Value{
			"Accept":              reflect.ValueOf(q.Accept),
			"Bind":                reflect.ValueOf(q.Bind),
			"BytePtrFromString":   reflect.ValueOf(q.BytePtrFromString),
			"ByteSliceFromString": reflect.ValueOf(q.ByteSliceFromString),
			"Chdir":               reflect.ValueOf(q.Chdir),
			"Chmod":               reflect.ValueOf(q.Chmod),
			"Chown":               reflect.ValueOf(q.Chown),
			"Clearenv":            reflect.ValueOf(q.Clearenv),
			"Close":               reflect.ValueOf(q.Close),
			"CloseOnExec":         reflect.ValueOf(q.CloseOnExec),
			"Connect":             reflect.ValueOf(q.Connect),
			"Dup":                 reflect.ValueOf(q.Dup),
			"Dup2":                reflect.ValueOf(q.Dup2),
			"Environ":             reflect.ValueOf(q.Environ),
			"Exit":                reflect.ValueOf(q.Exit),
			"Fchdir":              reflect.ValueOf(q.Fchdir),
			"Fchmod":              reflect.ValueOf(q.Fchmod),
			"Fchown":              reflect.ValueOf(q.Fchown),
			"Fstat":               reflect.ValueOf(q.Fstat),
			"Fsync":               reflect.ValueOf(q.Fsync),
			"Ftruncate":           reflect.ValueOf(q.Ftruncate),
			"Getcwd":              reflect.ValueOf(q.Getcwd),
			"Getegid":             reflect.ValueOf(q.Getegid),
			"Getenv":              reflect.ValueOf(q.Getenv),
			"Geteuid":             reflect.ValueOf(q.Geteuid),
			"Getgid":              reflect.ValueOf(q.Getgid),
			"Getgroups":           reflect.ValueOf(q.Getgroups),
			"Getpagesize":         reflect.ValueOf(q.Getpagesize),
			"Getpid":              reflect.ValueOf(q.Getpid),
			"Getppid":             reflect.ValueOf(q.Getppid),
			"GetsockoptInt":       reflect.ValueOf(q.GetsockoptInt),
			"Gettimeofday":        reflect.ValueOf(q.Gettimeofday),
			"Getuid":              reflect.ValueOf(q.Getuid),
			"Getwd":               reflect.ValueOf(q.Getwd),
			"Kill":                reflect.ValueOf(q.Kill),
			"Lchown":              reflect.ValueOf(q.Lchown),
			"Link":                reflect.ValueOf(q.Link),
			"Listen":              reflect.ValueOf(q.Listen),
			"Lstat":               reflect.ValueOf(q.Lstat),
			"Mkdir":               reflect.ValueOf(q.Mkdir),
			"NsecToTimespec":      reflect.ValueOf(q.NsecToTimespec),
			"NsecToTimeval":       reflect.ValueOf(q.NsecToTimeval),
			"Open":                reflect.ValueOf(q.Open),
			"ParseDirent":         reflect.ValueOf(q.ParseDirent),
			"Pipe":                reflect.ValueOf(q.Pipe),
			"Pread":               reflect.ValueOf(q.Pread),
			"Pwrite":              reflect.ValueOf(q.Pwrite),
			"RawSyscall":          reflect.ValueOf(q.RawSyscall),
			"RawSyscall6":         reflect.ValueOf(q.RawSyscall6),
			"Read":                reflect.ValueOf(q.Read),
			"ReadDirent":          reflect.ValueOf(q.ReadDirent),
			"Readlink":            reflect.ValueOf(q.Readlink),
			"Recvfrom":            reflect.ValueOf(q.Recvfrom),
			"Recvmsg":             reflect.ValueOf(q.Recvmsg),
			"Rename":              reflect.ValueOf(q.Rename),
			"Rmdir":               reflect.ValueOf(q.Rmdir),
			"Seek":                reflect.ValueOf(q.Seek),
			"Sendfile":            reflect.ValueOf(q.Sendfile),
			"SendmsgN":            reflect.ValueOf(q.SendmsgN),
			"Sendto":              reflect.ValueOf(q.Sendto),
			"SetNonblock":         reflect.ValueOf(q.SetNonblock),
			"SetReadDeadline":     reflect.ValueOf(q.SetReadDeadline),
			"SetWriteDeadline":    reflect.ValueOf(q.SetWriteDeadline),
			"Setenv":              reflect.ValueOf(q.Setenv),
			"SetsockoptInt":       reflect.ValueOf(q.SetsockoptInt),
			"Shutdown":            reflect.ValueOf(q.Shutdown),
			"Socket":              reflect.ValueOf(q.Socket),
			"StartProcess":        reflect.ValueOf(q.StartProcess),
			"Stat":                reflect.ValueOf(q.Stat),
			"StopIO":              reflect.ValueOf(q.StopIO),
			"StringBytePtr":       reflect.ValueOf(q.StringBytePtr),
			"StringByteSlice":     reflect.ValueOf(q.StringByteSlice),
			"Symlink":             reflect.ValueOf(q.Symlink),
			"Syscall":             reflect.ValueOf(q.Syscall),
			"Syscall6":            reflect.ValueOf(q.Syscall6),
			"Sysctl":              reflect.ValueOf(q.Sysctl),
			"TimespecToNsec":      reflect.ValueOf(q.TimespecToNsec),
			"TimevalToNsec":       reflect.ValueOf(q.TimevalToNsec),
			"Truncate":            reflect.ValueOf(q.Truncate),
			"Umask":               reflect.ValueOf(q.Umask),
			"Unlink":              reflect.ValueOf(q.Unlink),
			"Unsetenv":            reflect.ValueOf(q.Unsetenv),
			"UtimesNano":          reflect.ValueOf(q.UtimesNano),
			"Wait4":               reflect.ValueOf(q.Wait4),
			"Write":               reflect.ValueOf(q.Write),
		},
		TypedConsts: map[string]igop.TypedConst{
			"E2BIG":           {reflect.TypeOf(q.E2BIG), constant.MakeInt64(int64(q.E2BIG))},
			"EACCES":          {reflect.TypeOf(q.EACCES), constant.MakeInt64(int64(q.EACCES))},
			"EADDRINUSE":      {reflect.TypeOf(q.EADDRINUSE), constant.MakeInt64(int64(q.EADDRINUSE))},
			"EADDRNOTAVAIL":   {reflect.TypeOf(q.EADDRNOTAVAIL), constant.MakeInt64(int64(q.EADDRNOTAVAIL))},
			"EADV":            {reflect.TypeOf(q.EADV), constant.MakeInt64(int64(q.EADV))},
			"EAFNOSUPPORT":    {reflect.TypeOf(q.EAFNOSUPPORT), constant.MakeInt64(int64(q.EAFNOSUPPORT))},
			"EAGAIN":          {reflect.TypeOf(q.EAGAIN), constant.MakeInt64(int64(q.EAGAIN))},
			"EALREADY":        {reflect.TypeOf(q.EALREADY), constant.MakeInt64(int64(q.EALREADY))},
			"EBADE":           {reflect.TypeOf(q.EBADE), constant.MakeInt64(int64(q.EBADE))},
			"EBADF":           {reflect.TypeOf(q.EBADF), constant.MakeInt64(int64(q.EBADF))},
			"EBADFD":          {reflect.TypeOf(q.EBADFD), constant.MakeInt64(int64(q.EBADFD))},
			"EBADMSG":         {reflect.TypeOf(q.EBADMSG), constant.MakeInt64(int64(q.EBADMSG))},
			"EBADR":           {reflect.TypeOf(q.EBADR), constant.MakeInt64(int64(q.EBADR))},
			"EBADRQC":         {reflect.TypeOf(q.EBADRQC), constant.MakeInt64(int64(q.EBADRQC))},
			"EBADSLT":         {reflect.TypeOf(q.EBADSLT), constant.MakeInt64(int64(q.EBADSLT))},
			"EBFONT":          {reflect.TypeOf(q.EBFONT), constant.MakeInt64(int64(q.EBFONT))},
			"EBUSY":           {reflect.TypeOf(q.EBUSY), constant.MakeInt64(int64(q.EBUSY))},
			"ECANCELED":       {reflect.TypeOf(q.ECANCELED), constant.MakeInt64(int64(q.ECANCELED))},
			"ECASECLASH":      {reflect.TypeOf(q.ECASECLASH), constant.MakeInt64(int64(q.ECASECLASH))},
			"ECHILD":          {reflect.TypeOf(q.ECHILD), constant.MakeInt64(int64(q.ECHILD))},
			"ECHRNG":          {reflect.TypeOf(q.ECHRNG), constant.MakeInt64(int64(q.ECHRNG))},
			"ECOMM":           {reflect.TypeOf(q.ECOMM), constant.MakeInt64(int64(q.ECOMM))},
			"ECONNABORTED":    {reflect.TypeOf(q.ECONNABORTED), constant.MakeInt64(int64(q.ECONNABORTED))},
			"ECONNREFUSED":    {reflect.TypeOf(q.ECONNREFUSED), constant.MakeInt64(int64(q.ECONNREFUSED))},
			"ECONNRESET":      {reflect.TypeOf(q.ECONNRESET), constant.MakeInt64(int64(q.ECONNRESET))},
			"EDEADLK":         {reflect.TypeOf(q.EDEADLK), constant.MakeInt64(int64(q.EDEADLK))},
			"EDEADLOCK":       {reflect.TypeOf(q.EDEADLOCK), constant.MakeInt64(int64(q.EDEADLOCK))},
			"EDESTADDRREQ":    {reflect.TypeOf(q.EDESTADDRREQ), constant.MakeInt64(int64(q.EDESTADDRREQ))},
			"EDOM":            {reflect.TypeOf(q.EDOM), constant.MakeInt64(int64(q.EDOM))},
			"EDOTDOT":         {reflect.TypeOf(q.EDOTDOT), constant.MakeInt64(int64(q.EDOTDOT))},
			"EDQUOT":          {reflect.TypeOf(q.EDQUOT), constant.MakeInt64(int64(q.EDQUOT))},
			"EEXIST":          {reflect.TypeOf(q.EEXIST), constant.MakeInt64(int64(q.EEXIST))},
			"EFAULT":          {reflect.TypeOf(q.EFAULT), constant.MakeInt64(int64(q.EFAULT))},
			"EFBIG":           {reflect.TypeOf(q.EFBIG), constant.MakeInt64(int64(q.EFBIG))},
			"EFTYPE":          {reflect.TypeOf(q.EFTYPE), constant.MakeInt64(int64(q.EFTYPE))},
			"EHOSTDOWN":       {reflect.TypeOf(q.EHOSTDOWN), constant.MakeInt64(int64(q.EHOSTDOWN))},
			"EHOSTUNREACH":    {reflect.TypeOf(q.EHOSTUNREACH), constant.MakeInt64(int64(q.EHOSTUNREACH))},
			"EIDRM":           {reflect.TypeOf(q.EIDRM), constant.MakeInt64(int64(q.EIDRM))},
			"EILSEQ":          {reflect.TypeOf(q.EILSEQ), constant.MakeInt64(int64(q.EILSEQ))},
			"EINPROGRESS":     {reflect.TypeOf(q.EINPROGRESS), constant.MakeInt64(int64(q.EINPROGRESS))},
			"EINTR":           {reflect.TypeOf(q.EINTR), constant.MakeInt64(int64(q.EINTR))},
			"EINVAL":          {reflect.TypeOf(q.EINVAL), constant.MakeInt64(int64(q.EINVAL))},
			"EIO":             {reflect.TypeOf(q.EIO), constant.MakeInt64(int64(q.EIO))},
			"EISCONN":         {reflect.TypeOf(q.EISCONN), constant.MakeInt64(int64(q.EISCONN))},
			"EISDIR":          {reflect.TypeOf(q.EISDIR), constant.MakeInt64(int64(q.EISDIR))},
			"EL2HLT":          {reflect.TypeOf(q.EL2HLT), constant.MakeInt64(int64(q.EL2HLT))},
			"EL2NSYNC":        {reflect.TypeOf(q.EL2NSYNC), constant.MakeInt64(int64(q.EL2NSYNC))},
			"EL3HLT":          {reflect.TypeOf(q.EL3HLT), constant.MakeInt64(int64(q.EL3HLT))},
			"EL3RST":          {reflect.TypeOf(q.EL3RST), constant.MakeInt64(int64(q.EL3RST))},
			"ELBIN":           {reflect.TypeOf(q.ELBIN), constant.MakeInt64(int64(q.ELBIN))},
			"ELIBACC":         {reflect.TypeOf(q.ELIBACC), constant.MakeInt64(int64(q.ELIBACC))},
			"ELIBBAD":         {reflect.TypeOf(q.ELIBBAD), constant.MakeInt64(int64(q.ELIBBAD))},
			"ELIBEXEC":        {reflect.TypeOf(q.ELIBEXEC), constant.MakeInt64(int64(q.ELIBEXEC))},
			"ELIBMAX":         {reflect.TypeOf(q.ELIBMAX), constant.MakeInt64(int64(q.ELIBMAX))},
			"ELIBSCN":         {reflect.TypeOf(q.ELIBSCN), constant.MakeInt64(int64(q.ELIBSCN))},
			"ELNRNG":          {reflect.TypeOf(q.ELNRNG), constant.MakeInt64(int64(q.ELNRNG))},
			"ELOOP":           {reflect.TypeOf(q.ELOOP), constant.MakeInt64(int64(q.ELOOP))},
			"EMFILE":          {reflect.TypeOf(q.EMFILE), constant.MakeInt64(int64(q.EMFILE))},
			"EMLINK":          {reflect.TypeOf(q.EMLINK), constant.MakeInt64(int64(q.EMLINK))},
			"EMSGSIZE":        {reflect.TypeOf(q.EMSGSIZE), constant.MakeInt64(int64(q.EMSGSIZE))},
			"EMULTIHOP":       {reflect.TypeOf(q.EMULTIHOP), constant.MakeInt64(int64(q.EMULTIHOP))},
			"ENAMETOOLONG":    {reflect.TypeOf(q.ENAMETOOLONG), constant.MakeInt64(int64(q.ENAMETOOLONG))},
			"ENETDOWN":        {reflect.TypeOf(q.ENETDOWN), constant.MakeInt64(int64(q.ENETDOWN))},
			"ENETRESET":       {reflect.TypeOf(q.ENETRESET), constant.MakeInt64(int64(q.ENETRESET))},
			"ENETUNREACH":     {reflect.TypeOf(q.ENETUNREACH), constant.MakeInt64(int64(q.ENETUNREACH))},
			"ENFILE":          {reflect.TypeOf(q.ENFILE), constant.MakeInt64(int64(q.ENFILE))},
			"ENMFILE":         {reflect.TypeOf(q.ENMFILE), constant.MakeInt64(int64(q.ENMFILE))},
			"ENOANO":          {reflect.TypeOf(q.ENOANO), constant.MakeInt64(int64(q.ENOANO))},
			"ENOBUFS":         {reflect.TypeOf(q.ENOBUFS), constant.MakeInt64(int64(q.ENOBUFS))},
			"ENOCSI":          {reflect.TypeOf(q.ENOCSI), constant.MakeInt64(int64(q.ENOCSI))},
			"ENODATA":         {reflect.TypeOf(q.ENODATA), constant.MakeInt64(int64(q.ENODATA))},
			"ENODEV":          {reflect.TypeOf(q.ENODEV), constant.MakeInt64(int64(q.ENODEV))},
			"ENOENT":          {reflect.TypeOf(q.ENOENT), constant.MakeInt64(int64(q.ENOENT))},
			"ENOEXEC":         {reflect.TypeOf(q.ENOEXEC), constant.MakeInt64(int64(q.ENOEXEC))},
			"ENOLCK":          {reflect.TypeOf(q.ENOLCK), constant.MakeInt64(int64(q.ENOLCK))},
			"ENOLINK":         {reflect.TypeOf(q.ENOLINK), constant.MakeInt64(int64(q.ENOLINK))},
			"ENOMEDIUM":       {reflect.TypeOf(q.ENOMEDIUM), constant.MakeInt64(int64(q.ENOMEDIUM))},
			"ENOMEM":          {reflect.TypeOf(q.ENOMEM), constant.MakeInt64(int64(q.ENOMEM))},
			"ENOMSG":          {reflect.TypeOf(q.ENOMSG), constant.MakeInt64(int64(q.ENOMSG))},
			"ENONET":          {reflect.TypeOf(q.ENONET), constant.MakeInt64(int64(q.ENONET))},
			"ENOPKG":          {reflect.TypeOf(q.ENOPKG), constant.MakeInt64(int64(q.ENOPKG))},
			"ENOPROTOOPT":     {reflect.TypeOf(q.ENOPROTOOPT), constant.MakeInt64(int64(q.ENOPROTOOPT))},
			"ENOSHARE":        {reflect.TypeOf(q.ENOSHARE), constant.MakeInt64(int64(q.ENOSHARE))},
			"ENOSPC":          {reflect.TypeOf(q.ENOSPC), constant.MakeInt64(int64(q.ENOSPC))},
			"ENOSR":           {reflect.TypeOf(q.ENOSR), constant.MakeInt64(int64(q.ENOSR))},
			"ENOSTR":          {reflect.TypeOf(q.ENOSTR), constant.MakeInt64(int64(q.ENOSTR))},
			"ENOSYS":          {reflect.TypeOf(q.ENOSYS), constant.MakeInt64(int64(q.ENOSYS))},
			"ENOTCONN":        {reflect.TypeOf(q.ENOTCONN), constant.MakeInt64(int64(q.ENOTCONN))},
			"ENOTDIR":         {reflect.TypeOf(q.ENOTDIR), constant.MakeInt64(int64(q.ENOTDIR))},
			"ENOTEMPTY":       {reflect.TypeOf(q.ENOTEMPTY), constant.MakeInt64(int64(q.ENOTEMPTY))},
			"ENOTSOCK":        {reflect.TypeOf(q.ENOTSOCK), constant.MakeInt64(int64(q.ENOTSOCK))},
			"ENOTSUP":         {reflect.TypeOf(q.ENOTSUP), constant.MakeInt64(int64(q.ENOTSUP))},
			"ENOTTY":          {reflect.TypeOf(q.ENOTTY), constant.MakeInt64(int64(q.ENOTTY))},
			"ENOTUNIQ":        {reflect.TypeOf(q.ENOTUNIQ), constant.MakeInt64(int64(q.ENOTUNIQ))},
			"ENXIO":           {reflect.TypeOf(q.ENXIO), constant.MakeInt64(int64(q.ENXIO))},
			"EOPNOTSUPP":      {reflect.TypeOf(q.EOPNOTSUPP), constant.MakeInt64(int64(q.EOPNOTSUPP))},
			"EOVERFLOW":       {reflect.TypeOf(q.EOVERFLOW), constant.MakeInt64(int64(q.EOVERFLOW))},
			"EPERM":           {reflect.TypeOf(q.EPERM), constant.MakeInt64(int64(q.EPERM))},
			"EPFNOSUPPORT":    {reflect.TypeOf(q.EPFNOSUPPORT), constant.MakeInt64(int64(q.EPFNOSUPPORT))},
			"EPIPE":           {reflect.TypeOf(q.EPIPE), constant.MakeInt64(int64(q.EPIPE))},
			"EPROCLIM":        {reflect.TypeOf(q.EPROCLIM), constant.MakeInt64(int64(q.EPROCLIM))},
			"EPROTO":          {reflect.TypeOf(q.EPROTO), constant.MakeInt64(int64(q.EPROTO))},
			"EPROTONOSUPPORT": {reflect.TypeOf(q.EPROTONOSUPPORT), constant.MakeInt64(int64(q.EPROTONOSUPPORT))},
			"EPROTOTYPE":      {reflect.TypeOf(q.EPROTOTYPE), constant.MakeInt64(int64(q.EPROTOTYPE))},
			"ERANGE":          {reflect.TypeOf(q.ERANGE), constant.MakeInt64(int64(q.ERANGE))},
			"EREMCHG":         {reflect.TypeOf(q.EREMCHG), constant.MakeInt64(int64(q.EREMCHG))},
			"EREMOTE":         {reflect.TypeOf(q.EREMOTE), constant.MakeInt64(int64(q.EREMOTE))},
			"EROFS":           {reflect.TypeOf(q.EROFS), constant.MakeInt64(int64(q.EROFS))},
			"ESHUTDOWN":       {reflect.TypeOf(q.ESHUTDOWN), constant.MakeInt64(int64(q.ESHUTDOWN))},
			"ESOCKTNOSUPPORT": {reflect.TypeOf(q.ESOCKTNOSUPPORT), constant.MakeInt64(int64(q.ESOCKTNOSUPPORT))},
			"ESPIPE":          {reflect.TypeOf(q.ESPIPE), constant.MakeInt64(int64(q.ESPIPE))},
			"ESRCH":           {reflect.TypeOf(q.ESRCH), constant.MakeInt64(int64(q.ESRCH))},
			"ESRMNT":          {reflect.TypeOf(q.ESRMNT), constant.MakeInt64(int64(q.ESRMNT))},
			"ESTALE":          {reflect.TypeOf(q.ESTALE), constant.MakeInt64(int64(q.ESTALE))},
			"ETIME":           {reflect.TypeOf(q.ETIME), constant.MakeInt64(int64(q.ETIME))},
			"ETIMEDOUT":       {reflect.TypeOf(q.ETIMEDOUT), constant.MakeInt64(int64(q.ETIMEDOUT))},
			"ETOOMANYREFS":    {reflect.TypeOf(q.ETOOMANYREFS), constant.MakeInt64(int64(q.ETOOMANYREFS))},
			"EUNATCH":         {reflect.TypeOf(q.EUNATCH), constant.MakeInt64(int64(q.EUNATCH))},
			"EUSERS":          {reflect.TypeOf(q.EUSERS), constant.MakeInt64(int64(q.EUSERS))},
			"EWOULDBLOCK":     {reflect.TypeOf(q.EWOULDBLOCK), constant.MakeInt64(int64(q.EWOULDBLOCK))},
			"EXDEV":           {reflect.TypeOf(q.EXDEV), constant.MakeInt64(int64(q.EXDEV))},
			"EXFULL":          {reflect.TypeOf(q.EXFULL), constant.MakeInt64(int64(q.EXFULL))},
			"SIGCHLD":         {reflect.TypeOf(q.SIGCHLD), constant.MakeInt64(int64(q.SIGCHLD))},
			"SIGINT":          {reflect.TypeOf(q.SIGINT), constant.MakeInt64(int64(q.SIGINT))},
			"SIGKILL":         {reflect.TypeOf(q.SIGKILL), constant.MakeInt64(int64(q.SIGKILL))},
			"SIGQUIT":         {reflect.TypeOf(q.SIGQUIT), constant.MakeInt64(int64(q.SIGQUIT))},
			"SIGTERM":         {reflect.TypeOf(q.SIGTERM), constant.MakeInt64(int64(q.SIGTERM))},
			"SIGTRAP":         {reflect.TypeOf(q.SIGTRAP), constant.MakeInt64(int64(q.SIGTRAP))},
		},
		UntypedConsts: map[string]igop.UntypedConst{
			"AF_INET":         {"untyped int", constant.MakeInt64(int64(q.AF_INET))},
			"AF_INET6":        {"untyped int", constant.MakeInt64(int64(q.AF_INET6))},
			"AF_UNIX":         {"untyped int", constant.MakeInt64(int64(q.AF_UNIX))},
			"AF_UNSPEC":       {"untyped int", constant.MakeInt64(int64(q.AF_UNSPEC))},
			"F_CNVT":          {"untyped int", constant.MakeInt64(int64(q.F_CNVT))},
			"F_DUPFD":         {"untyped int", constant.MakeInt64(int64(q.F_DUPFD))},
			"F_DUPFD_CLOEXEC": {"untyped int", constant.MakeInt64(int64(q.F_DUPFD_CLOEXEC))},
			"F_GETFD":         {"untyped int", constant.MakeInt64(int64(q.F_GETFD))},
			"F_GETFL":         {"untyped int", constant.MakeInt64(int64(q.F_GETFL))},
			"F_GETLK":         {"untyped int", constant.MakeInt64(int64(q.F_GETLK))},
			"F_GETOWN":        {"untyped int", constant.MakeInt64(int64(q.F_GETOWN))},
			"F_RDLCK":         {"untyped int", constant.MakeInt64(int64(q.F_RDLCK))},
			"F_RGETLK":        {"untyped int", constant.MakeInt64(int64(q.F_RGETLK))},
			"F_RSETLK":        {"untyped int", constant.MakeInt64(int64(q.F_RSETLK))},
			"F_RSETLKW":       {"untyped int", constant.MakeInt64(int64(q.F_RSETLKW))},
			"F_SETFD":         {"untyped int", constant.MakeInt64(int64(q.F_SETFD))},
			"F_SETFL":         {"untyped int", constant.MakeInt64(int64(q.F_SETFL))},
			"F_SETLK":         {"untyped int", constant.MakeInt64(int64(q.F_SETLK))},
			"F_SETLKW":        {"untyped int", constant.MakeInt64(int64(q.F_SETLKW))},
			"F_SETOWN":        {"untyped int", constant.MakeInt64(int64(q.F_SETOWN))},
			"F_UNLCK":         {"untyped int", constant.MakeInt64(int64(q.F_UNLCK))},
			"F_UNLKSYS":       {"untyped int", constant.MakeInt64(int64(q.F_UNLKSYS))},
			"F_WRLCK":         {"untyped int", constant.MakeInt64(int64(q.F_WRLCK))},
			"IPPROTO_IP":      {"untyped int", constant.MakeInt64(int64(q.IPPROTO_IP))},
			"IPPROTO_IPV4":    {"untyped int", constant.MakeInt64(int64(q.IPPROTO_IPV4))},
			"IPPROTO_IPV6":    {"untyped int", constant.MakeInt64(int64(q.IPPROTO_IPV6))},
			"IPPROTO_TCP":     {"untyped int", constant.MakeInt64(int64(q.IPPROTO_TCP))},
			"IPPROTO_UDP":     {"untyped int", constant.MakeInt64(int64(q.IPPROTO_UDP))},
			"IPV6_V6ONLY":     {"untyped int", constant.MakeInt64(int64(q.IPV6_V6ONLY))},
			"ImplementsGetwd": {"untyped bool", constant.MakeBool(bool(q.ImplementsGetwd))},
			"O_APPEND":        {"untyped int", constant.MakeInt64(int64(q.O_APPEND))},
			"O_CLOEXEC":       {"untyped int", constant.MakeInt64(int64(q.O_CLOEXEC))},
			"O_CREAT":         {"untyped int", constant.MakeInt64(int64(q.O_CREAT))},
			"O_CREATE":        {"untyped int", constant.MakeInt64(int64(q.O_CREATE))},
			"O_EXCL":          {"untyped int", constant.MakeInt64(int64(q.O_EXCL))},
			"O_RDONLY":        {"untyped int", constant.MakeInt64(int64(q.O_RDONLY))},
			"O_RDWR":          {"untyped int", constant.MakeInt64(int64(q.O_RDWR))},
			"O_SYNC":          {"untyped int", constant.MakeInt64(int64(q.O_SYNC))},
			"O_TRUNC":         {"untyped int", constant.MakeInt64(int64(q.O_TRUNC))},
			"O_WRONLY":        {"untyped int", constant.MakeInt64(int64(q.O_WRONLY))},
			"PathMax":         {"untyped int", constant.MakeInt64(int64(q.PathMax))},
			"SOCK_DGRAM":      {"untyped int", constant.MakeInt64(int64(q.SOCK_DGRAM))},
			"SOCK_RAW":        {"untyped int", constant.MakeInt64(int64(q.SOCK_RAW))},
			"SOCK_SEQPACKET":  {"untyped int", constant.MakeInt64(int64(q.SOCK_SEQPACKET))},
			"SOCK_STREAM":     {"untyped int", constant.MakeInt64(int64(q.SOCK_STREAM))},
			"SOMAXCONN":       {"untyped int", constant.MakeInt64(int64(q.SOMAXCONN))},
			"SO_ERROR":        {"untyped int", constant.MakeInt64(int64(q.SO_ERROR))},
			"SYS_FCNTL":       {"untyped int", constant.MakeInt64(int64(q.SYS_FCNTL))},
			"S_IEXEC":         {"untyped int", constant.MakeInt64(int64(q.S_IEXEC))},
			"S_IFBLK":         {"untyped int", constant.MakeInt64(int64(q.S_IFBLK))},
			"S_IFBOUNDSOCK":   {"untyped int", constant.MakeInt64(int64(q.S_IFBOUNDSOCK))},
			"S_IFCHR":         {"untyped int", constant.MakeInt64(int64(q.S_IFCHR))},
			"S_IFCOND":        {"untyped int", constant.MakeInt64(int64(q.S_IFCOND))},
			"S_IFDIR":         {"untyped int", constant.MakeInt64(int64(q.S_IFDIR))},
			"S_IFDSOCK":       {"untyped int", constant.MakeInt64(int64(q.S_IFDSOCK))},
			"S_IFIFO":         {"untyped int", constant.MakeInt64(int64(q.S_IFIFO))},
			"S_IFLNK":         {"untyped int", constant.MakeInt64(int64(q.S_IFLNK))},
			"S_IFMT":          {"untyped int", constant.MakeInt64(int64(q.S_IFMT))},
			"S_IFMUTEX":       {"untyped int", constant.MakeInt64(int64(q.S_IFMUTEX))},
			"S_IFREG":         {"untyped int", constant.MakeInt64(int64(q.S_IFREG))},
			"S_IFSEMA":        {"untyped int", constant.MakeInt64(int64(q.S_IFSEMA))},
			"S_IFSHM":         {"untyped int", constant.MakeInt64(int64(q.S_IFSHM))},
			"S_IFSHM_SYSV":    {"untyped int", constant.MakeInt64(int64(q.S_IFSHM_SYSV))},
			"S_IFSOCK":        {"untyped int", constant.MakeInt64(int64(q.S_IFSOCK))},
			"S_IFSOCKADDR":    {"untyped int", constant.MakeInt64(int64(q.S_IFSOCKADDR))},
			"S_IREAD":         {"untyped int", constant.MakeInt64(int64(q.S_IREAD))},
			"S_IRGRP":         {"untyped int", constant.MakeInt64(int64(q.S_IRGRP))},
			"S_IROTH":         {"untyped int", constant.MakeInt64(int64(q.S_IROTH))},
			"S_IRUSR":         {"untyped int", constant.MakeInt64(int64(q.S_IRUSR))},
			"S_IRWXG":         {"untyped int", constant.MakeInt64(int64(q.S_IRWXG))},
			"S_IRWXO":         {"untyped int", constant.MakeInt64(int64(q.S_IRWXO))},
			"S_IRWXU":         {"untyped int", constant.MakeInt64(int64(q.S_IRWXU))},
			"S_ISGID":         {"untyped int", constant.MakeInt64(int64(q.S_ISGID))},
			"S_ISUID":         {"untyped int", constant.MakeInt64(int64(q.S_ISUID))},
			"S_ISVTX":         {"untyped int", constant.MakeInt64(int64(q.S_ISVTX))},
			"S_IWGRP":         {"untyped int", constant.MakeInt64(int64(q.S_IWGRP))},
			"S_IWOTH":         {"untyped int", constant.MakeInt64(int64(q.S_IWOTH))},
			"S_IWRITE":        {"untyped int", constant.MakeInt64(int64(q.S_IWRITE))},
			"S_IWUSR":         {"untyped int", constant.MakeInt64(int64(q.S_IWUSR))},
			"S_IXGRP":         {"untyped int", constant.MakeInt64(int64(q.S_IXGRP))},
			"S_IXOTH":         {"untyped int", constant.MakeInt64(int64(q.S_IXOTH))},
			"S_IXUSR":         {"untyped int", constant.MakeInt64(int64(q.S_IXUSR))},
			"S_UNSUP":         {"untyped int", constant.MakeInt64(int64(q.S_UNSUP))},
			"Stderr":          {"untyped int", constant.MakeInt64(int64(q.Stderr))},
			"Stdin":           {"untyped int", constant.MakeInt64(int64(q.Stdin))},
			"Stdout":          {"untyped int", constant.MakeInt64(int64(q.Stdout))},
		},
	})
}
