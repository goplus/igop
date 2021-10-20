// export by github.com/goplus/gossa/cmd/qexp

package syscall

import (
	"syscall"

	"github.com/goplus/gossa"
)

func init() {
	gossa.RegisterPackage("syscall", extMap_freebsd, typList_freebsd)
}

var extMap_freebsd = map[string]interface{}{
	"syscall.Fstatat": syscall.Fstatat,
}

var typList_freebsd = []interface{}{
	(*syscall.BpfZbuf)(nil),
	(*syscall.BpfZbufHeader)(nil),
}
