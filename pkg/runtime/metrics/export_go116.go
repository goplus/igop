// export by github.com/goplus/interp/cmd/qexp

// +build go1.16

package metrics

import (
	"runtime/metrics"

	"github.com/goplus/interp"
)

func init() {
	interp.RegisterPackage("runtime/metrics", extMap_go116, typList_go116)
}

var extMap_go116 = map[string]interface{}{
	"(runtime/metrics.Value).Float64":          (metrics.Value).Float64,
	"(runtime/metrics.Value).Float64Histogram": (metrics.Value).Float64Histogram,
	"(runtime/metrics.Value).Kind":             (metrics.Value).Kind,
	"(runtime/metrics.Value).Uint64":           (metrics.Value).Uint64,
	"runtime/metrics.All":                      metrics.All,
	"runtime/metrics.Read":                     metrics.Read,
}

var typList_go116 = []interface{}{
	(*metrics.Description)(nil),
	(*metrics.Float64Histogram)(nil),
	(*metrics.Sample)(nil),
	(*metrics.Value)(nil),
	(*metrics.ValueKind)(nil),
}
