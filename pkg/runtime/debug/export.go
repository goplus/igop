// export by github.com/goplus/gossa/cmd/qexp

package debug

import (
	"runtime/debug"

	"github.com/goplus/gossa"
)

func init() {
	gossa.RegisterPackage("runtime/debug", extMap, typList)
}

var extMap = map[string]interface{}{
	"runtime/debug.FreeOSMemory":    debug.FreeOSMemory,
	"runtime/debug.PrintStack":      debug.PrintStack,
	"runtime/debug.ReadBuildInfo":   debug.ReadBuildInfo,
	"runtime/debug.ReadGCStats":     debug.ReadGCStats,
	"runtime/debug.SetGCPercent":    debug.SetGCPercent,
	"runtime/debug.SetMaxStack":     debug.SetMaxStack,
	"runtime/debug.SetMaxThreads":   debug.SetMaxThreads,
	"runtime/debug.SetPanicOnFault": debug.SetPanicOnFault,
	"runtime/debug.SetTraceback":    debug.SetTraceback,
	"runtime/debug.Stack":           debug.Stack,
	"runtime/debug.WriteHeapDump":   debug.WriteHeapDump,
}

var typList = []interface{}{
	(*debug.BuildInfo)(nil),
	(*debug.GCStats)(nil),
	(*debug.Module)(nil),
}
