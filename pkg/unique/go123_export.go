// export by github.com/goplus/igop/cmd/qexp

//go:build go1.23
// +build go1.23

package unique

import (
	_ "unique"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "unique",
		Path: "unique",
		Deps: map[string]string{
			"internal/abi":         "abi",
			"internal/concurrent":  "concurrent",
			"internal/stringslite": "stringslite",
			"internal/weak":        "weak",
			"runtime":              "runtime",
			"sync":                 "sync",
			"unsafe":               "unsafe",
		},
	})
}
