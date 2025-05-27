// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.23 && !go1.24
// +build go1.23,!go1.24

package unique

import (
	_ "unique"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage(&ixgo.Package{
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
