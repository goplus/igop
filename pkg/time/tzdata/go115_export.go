// export by github.com/goplus/igop/cmd/qexp

//+build go1.15,!go1.16

package tzdata

import (
	_ "time/tzdata"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "tzdata",
		Path: "time/tzdata",
		Deps: map[string]string{
			"errors":  "errors",
			"syscall": "syscall",
			"unsafe":  "unsafe",
		},
	})
}
