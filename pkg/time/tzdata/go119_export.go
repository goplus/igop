// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.19 && !go1.20
// +build go1.19,!go1.20

package tzdata

import (
	_ "time/tzdata"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage(&ixgo.Package{
		Name: "tzdata",
		Path: "time/tzdata",
		Deps: map[string]string{
			"errors":  "errors",
			"syscall": "syscall",
			"unsafe":  "unsafe",
		},
	})
}
