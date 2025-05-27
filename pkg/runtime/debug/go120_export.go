// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.20 && !go1.21
// +build go1.20,!go1.21

package debug

import (
	q "runtime/debug"

	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage(&ixgo.Package{
		Name: "debug",
		Path: "runtime/debug",
		Deps: map[string]string{
			"fmt":     "fmt",
			"os":      "os",
			"runtime": "runtime",
			"sort":    "sort",
			"strconv": "strconv",
			"strings": "strings",
			"time":    "time",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{
			"BuildInfo":    reflect.TypeOf((*q.BuildInfo)(nil)).Elem(),
			"BuildSetting": reflect.TypeOf((*q.BuildSetting)(nil)).Elem(),
			"GCStats":      reflect.TypeOf((*q.GCStats)(nil)).Elem(),
			"Module":       reflect.TypeOf((*q.Module)(nil)).Elem(),
		},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"FreeOSMemory":    reflect.ValueOf(q.FreeOSMemory),
			"ParseBuildInfo":  reflect.ValueOf(q.ParseBuildInfo),
			"PrintStack":      reflect.ValueOf(q.PrintStack),
			"ReadBuildInfo":   reflect.ValueOf(q.ReadBuildInfo),
			"ReadGCStats":     reflect.ValueOf(q.ReadGCStats),
			"SetGCPercent":    reflect.ValueOf(q.SetGCPercent),
			"SetMaxStack":     reflect.ValueOf(q.SetMaxStack),
			"SetMaxThreads":   reflect.ValueOf(q.SetMaxThreads),
			"SetMemoryLimit":  reflect.ValueOf(q.SetMemoryLimit),
			"SetPanicOnFault": reflect.ValueOf(q.SetPanicOnFault),
			"SetTraceback":    reflect.ValueOf(q.SetTraceback),
			"Stack":           reflect.ValueOf(q.Stack),
			"WriteHeapDump":   reflect.ValueOf(q.WriteHeapDump),
		},
		TypedConsts:   map[string]ixgo.TypedConst{},
		UntypedConsts: map[string]ixgo.UntypedConst{},
	})
}
