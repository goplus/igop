// export by github.com/goplus/gossa/cmd/qexp

//go:build darwin || freebsd || netbsd || openbsd || windows
// +build darwin freebsd netbsd openbsd windows

package syscall

import (
	"syscall"

	"github.com/goplus/gossa"
)

func init() {
	gossa.RegisterPackage("syscall", extMap_2528687091, typList_2528687091)
}

var extMap_2528687091 = map[string]interface{}{
	"syscall.Syscall9": syscall.Syscall9,
}

var typList_2528687091 = []interface{}{}
