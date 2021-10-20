// export by github.com/goplus/gossa/cmd/qexp

//go:build freebsd || linux || netbsd || openbsd
// +build freebsd linux netbsd openbsd

package syscall

import (
	"syscall"

	"github.com/goplus/gossa"
)

func init() {
	gossa.RegisterPackage("syscall", extMap_3829281923, typList_3829281923)
}

var extMap_3829281923 = map[string]interface{}{
	"syscall.Accept4":   syscall.Accept4,
	"syscall.Nanosleep": syscall.Nanosleep,
	"syscall.Pipe2":     syscall.Pipe2,
}

var typList_3829281923 = []interface{}{}
