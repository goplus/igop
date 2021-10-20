// export by github.com/goplus/gossa/cmd/qexp

//go:build netbsd || openbsd
// +build netbsd openbsd

package syscall

import (
	"syscall"

	"github.com/goplus/gossa"
)

func init() {
	gossa.RegisterPackage("syscall", extMap_909332034, typList_909332034)
}

var extMap_909332034 = map[string]interface{}{}

var typList_909332034 = []interface{}{
	(*syscall.BpfTimeval)(nil),
	(*syscall.Mclpool)(nil),
}
