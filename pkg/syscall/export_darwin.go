// export by github.com/goplus/gossa/cmd/qexp

package syscall

import (
	"syscall"

	"github.com/goplus/gossa"
)

func init() {
	gossa.RegisterPackage("syscall", extMap_darwin, typList_darwin)
}

var extMap_darwin = map[string]interface{}{
	"syscall.Exchangedata": syscall.Exchangedata,
	"syscall.Setprivexec":  syscall.Setprivexec,
}

var typList_darwin = []interface{}{
	(*syscall.Fbootstraptransfer_t)(nil),
	(*syscall.Fstore_t)(nil),
	(*syscall.IfmaMsghdr2)(nil),
	(*syscall.Log2phys_t)(nil),
	(*syscall.Radvisory_t)(nil),
	(*syscall.Timeval32)(nil),
}
