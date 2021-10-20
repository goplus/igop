// export by github.com/goplus/gossa/cmd/qexp

package syscall

import (
	"syscall"

	"github.com/goplus/gossa"
)

func init() {
	gossa.RegisterPackage("syscall", extMap_netbsd, typList_netbsd)
}

var extMap_netbsd = map[string]interface{}{}

var typList_netbsd = []interface{}{
	(*syscall.Sysctlnode)(nil),
}
