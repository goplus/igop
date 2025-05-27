//go:build go1.20 && !go1.21
// +build go1.20,!go1.21

package test

import (
	_ "runtime"
	_ "unsafe"

	"github.com/goplus/ixgo"
)

//go:linkname setUpdate internal/godebug.setUpdate
func setUpdate(update func(string, string))

func init() {
	ixgo.RegisterExternal("internal/godebug.setUpdate", setUpdate)
}
