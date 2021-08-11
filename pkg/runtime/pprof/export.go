// export by github.com/goplus/interp/cmd/qexp

package pprof

import (
	"runtime/pprof"

	"github.com/goplus/interp"
)

func init() {
	interp.RegisterPackage("runtime/pprof", extMap, typList)
}

var extMap = map[string]interface{}{
	"(*runtime/pprof.Profile).Add":     (*pprof.Profile).Add,
	"(*runtime/pprof.Profile).Count":   (*pprof.Profile).Count,
	"(*runtime/pprof.Profile).Name":    (*pprof.Profile).Name,
	"(*runtime/pprof.Profile).Remove":  (*pprof.Profile).Remove,
	"(*runtime/pprof.Profile).WriteTo": (*pprof.Profile).WriteTo,
	"runtime/pprof.Do":                 pprof.Do,
	"runtime/pprof.ForLabels":          pprof.ForLabels,
	"runtime/pprof.Label":              pprof.Label,
	"runtime/pprof.Labels":             pprof.Labels,
	"runtime/pprof.Lookup":             pprof.Lookup,
	"runtime/pprof.NewProfile":         pprof.NewProfile,
	"runtime/pprof.Profiles":           pprof.Profiles,
	"runtime/pprof.SetGoroutineLabels": pprof.SetGoroutineLabels,
	"runtime/pprof.StartCPUProfile":    pprof.StartCPUProfile,
	"runtime/pprof.StopCPUProfile":     pprof.StopCPUProfile,
	"runtime/pprof.WithLabels":         pprof.WithLabels,
	"runtime/pprof.WriteHeapProfile":   pprof.WriteHeapProfile,
}

var typList = []interface{}{
	(*pprof.LabelSet)(nil),
	(*pprof.Profile)(nil),
}
