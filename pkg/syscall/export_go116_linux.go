// export by github.com/goplus/gossa/cmd/qexp

//go:build go1.16
// +build go1.16

package syscall

import (
	"syscall"

	"github.com/goplus/gossa"
)

func init() {
	gossa.RegisterPackage("syscall", extMap_go116_linux, typList_go116_linux)
}

var extMap_go116_linux = map[string]interface{}{
	"syscall.AllThreadsSyscall":  syscall.AllThreadsSyscall,
	"syscall.AllThreadsSyscall6": syscall.AllThreadsSyscall6,
	"syscall.Setegid":            syscall.Setegid,
	"syscall.Seteuid":            syscall.Seteuid,
}

var typList_go116_linux = []interface{}{}
