//go:build go1.21
// +build go1.21

package test

import (
	_ "runtime"
	"unsafe"

	"github.com/goplus/igop"
)

//go:linkname setUpdate internal/godebug.setUpdate
func setUpdate(update func(string, string))

//go:linkname registerMetric internal/godebug.registerMetric
func registerMetric(name string, read func() uint64)

//go:linkname setNewIncNonDefault internal/godebug.setNewIncNonDefault
func setNewIncNonDefault(newIncNonDefault func(string) func())

//go:linkname write runtime.write
func write(fd uintptr, p unsafe.Pointer, n int32) int32

func init() {
	igop.RegisterExternal("internal/godebug.setUpdate", setUpdate)
	igop.RegisterExternal("internal/godebug.registerMetric", registerMetric)
	igop.RegisterExternal("internal/godebug.setNewIncNonDefault", setNewIncNonDefault)
	igop.RegisterExternal("internal/godebug.write", write)
}
