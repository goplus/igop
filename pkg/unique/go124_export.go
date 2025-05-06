// export by github.com/goplus/igop/cmd/qexp

//go:build go1.24
// +build go1.24

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
			"internal/stringslite": "stringslite",
			"internal/sync":        "sync",
			"runtime":              "runtime",
			"sync":                 "sync",
			"unsafe":               "unsafe",
			"weak":                 "weak",
		},
	})
}
