// export by github.com/goplus/igop/cmd/qexp

//go:build go1.21
// +build go1.21

package maps

import (
	_ "maps"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "maps",
		Path: "maps",
		Deps: map[string]string{},
	})
}
