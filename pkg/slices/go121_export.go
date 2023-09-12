// export by github.com/goplus/igop/cmd/qexp

//go:build go1.21
// +build go1.21

package slices

import (
	_ "slices"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "slices",
		Path: "slices",
		Deps: map[string]string{
			"cmp":       "cmp",
			"math/bits": "bits",
			"unsafe":    "unsafe",
		},
	})
}
