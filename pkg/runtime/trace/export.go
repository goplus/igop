// export by github.com/goplus/gossa/cmd/qexp

package trace

import (
	"runtime/trace"

	"github.com/goplus/gossa"
)

func init() {
	gossa.RegisterPackage("runtime/trace", extMap, typList)
}

var extMap = map[string]interface{}{
	"(*runtime/trace.Region).End": (*trace.Region).End,
	"(*runtime/trace.Task).End":   (*trace.Task).End,
	"runtime/trace.IsEnabled":     trace.IsEnabled,
	"runtime/trace.Log":           trace.Log,
	"runtime/trace.Logf":          trace.Logf,
	"runtime/trace.NewTask":       trace.NewTask,
	"runtime/trace.Start":         trace.Start,
	"runtime/trace.StartRegion":   trace.StartRegion,
	"runtime/trace.Stop":          trace.Stop,
	"runtime/trace.WithRegion":    trace.WithRegion,
}

var typList = []interface{}{
	(*trace.Region)(nil),
	(*trace.Task)(nil),
}
