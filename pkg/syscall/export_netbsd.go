// export by github.com/goplus/interp/cmd/qexp

package syscall

import (
	"syscall"

	"github.com/goplus/interp"
)

func init() {
	interp.RegisterPackage("syscall", extMap_netbsd, typList_netbsd)
}

var extMap_netbsd = map[string]interface{}{}

var typList_netbsd = []interface{}{
	(*syscall.Sysctlnode)(nil),
}
