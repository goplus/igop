// export by github.com/goplus/igop/cmd/qexp

//go:build go1.24
// +build go1.24

package weak

import (
	_ "weak"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "weak",
		Path: "weak",
		Deps: map[string]string{
			"internal/abi": "abi",
			"runtime":      "runtime",
			"unsafe":       "unsafe",
		},
	})
}
