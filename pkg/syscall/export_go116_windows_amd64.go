// export by github.com/goplus/interp/cmd/qexp

//go:build go1.16
// +build go1.16

package syscall

import (
	"syscall"

	"github.com/goplus/interp"
)

func init() {
	interp.RegisterPackage("syscall", extMap_go116_windows_amd64, typList_go116_windows_amd64)
}

var extMap_go116_windows_amd64 = map[string]interface{}{
	"(*syscall.DLLError).Unwrap": (*syscall.DLLError).Unwrap,
}

var typList_go116_windows_amd64 = []interface{}{}
