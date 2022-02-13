// export by github.com/goplus/gossa/cmd/qexp

//+build go1.15,!go1.16

package testing

import (
	q "testing"

	"reflect"

	"github.com/goplus/gossa"
)

func init() {
	gossa.RegisterPackage(&gossa.Package{
		Name: "testing",
		Path: "testing",
		Deps: map[string]string{
			"bytes":         "bytes",
			"errors":        "errors",
			"flag":          "flag",
			"fmt":           "fmt",
			"internal/race": "race",
			"io":            "io",
			"io/ioutil":     "ioutil",
			"math":          "math",
			"os":            "os",
			"runtime":       "runtime",
			"runtime/debug": "debug",
			"runtime/trace": "trace",
			"sort":          "sort",
			"strconv":       "strconv",
			"strings":       "strings",
			"sync":          "sync",
			"sync/atomic":   "atomic",
			"time":          "time",
			"unicode":       "unicode",
		},
		Interfaces: map[string]reflect.Type{
			"TB": reflect.TypeOf((*q.TB)(nil)).Elem(),
		},
		NamedTypes: map[string]gossa.NamedType{
			"B":                 {reflect.TypeOf((*q.B)(nil)).Elem(), "", "ReportAllocs,ReportMetric,ResetTimer,Run,RunParallel,SetBytes,SetParallelism,StartTimer,StopTimer,add,doBench,launch,run,run1,runN,trimOutput"},
			"BenchmarkResult":   {reflect.TypeOf((*q.BenchmarkResult)(nil)).Elem(), "AllocedBytesPerOp,AllocsPerOp,MemString,NsPerOp,String,mbPerSec", ""},
			"Cover":             {reflect.TypeOf((*q.Cover)(nil)).Elem(), "", ""},
			"CoverBlock":        {reflect.TypeOf((*q.CoverBlock)(nil)).Elem(), "", ""},
			"InternalBenchmark": {reflect.TypeOf((*q.InternalBenchmark)(nil)).Elem(), "", ""},
			"InternalExample":   {reflect.TypeOf((*q.InternalExample)(nil)).Elem(), "", "processRunResult"},
			"InternalTest":      {reflect.TypeOf((*q.InternalTest)(nil)).Elem(), "", ""},
			"M":                 {reflect.TypeOf((*q.M)(nil)).Elem(), "", "Run,after,before,startAlarm,stopAlarm,writeProfiles"},
			"PB":                {reflect.TypeOf((*q.PB)(nil)).Elem(), "", "Next"},
			"T":                 {reflect.TypeOf((*q.T)(nil)).Elem(), "", "Deadline,Parallel,Run,report"},
		},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"AllocsPerRun":  reflect.ValueOf(q.AllocsPerRun),
			"Benchmark":     reflect.ValueOf(q.Benchmark),
			"CoverMode":     reflect.ValueOf(q.CoverMode),
			"Coverage":      reflect.ValueOf(q.Coverage),
			"Init":          reflect.ValueOf(q.Init),
			"Main":          reflect.ValueOf(q.Main),
			"MainStart":     reflect.ValueOf(q.MainStart),
			"RegisterCover": reflect.ValueOf(q.RegisterCover),
			"RunBenchmarks": reflect.ValueOf(q.RunBenchmarks),
			"RunExamples":   reflect.ValueOf(q.RunExamples),
			"RunTests":      reflect.ValueOf(q.RunTests),
			"Short":         reflect.ValueOf(q.Short),
			"Verbose":       reflect.ValueOf(q.Verbose),
		},
		TypedConsts:   map[string]gossa.TypedConst{},
		UntypedConsts: map[string]gossa.UntypedConst{},
	})
}
