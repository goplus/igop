// export by github.com/goplus/gossa/cmd/qexp

package runtime

import (
	"runtime"

	"github.com/goplus/gossa"
)

func init() {
	gossa.RegisterPackage("runtime", extMap, typList)
}

var extMap = map[string]interface{}{
	"(*runtime.BlockProfileRecord).Stack":        (*runtime.BlockProfileRecord).Stack,
	"(*runtime.Frames).Next":                     (*runtime.Frames).Next,
	"(*runtime.Func).Entry":                      (*runtime.Func).Entry,
	"(*runtime.Func).FileLine":                   (*runtime.Func).FileLine,
	"(*runtime.Func).Name":                       (*runtime.Func).Name,
	"(*runtime.MemProfileRecord).InUseBytes":     (*runtime.MemProfileRecord).InUseBytes,
	"(*runtime.MemProfileRecord).InUseObjects":   (*runtime.MemProfileRecord).InUseObjects,
	"(*runtime.MemProfileRecord).Stack":          (*runtime.MemProfileRecord).Stack,
	"(*runtime.StackRecord).Stack":               (*runtime.StackRecord).Stack,
	"(*runtime.TypeAssertionError).Error":        (*runtime.TypeAssertionError).Error,
	"(*runtime.TypeAssertionError).RuntimeError": (*runtime.TypeAssertionError).RuntimeError,
	"(runtime.Error).Error":                      (runtime.Error).Error,
	"(runtime.Error).RuntimeError":               (runtime.Error).RuntimeError,
	"runtime.BlockProfile":                       runtime.BlockProfile,
	"runtime.Breakpoint":                         runtime.Breakpoint,
	"runtime.CPUProfile":                         runtime.CPUProfile,
	"runtime.Caller":                             runtime.Caller,
	"runtime.Callers":                            runtime.Callers,
	"runtime.CallersFrames":                      runtime.CallersFrames,
	"runtime.FuncForPC":                          runtime.FuncForPC,
	"runtime.GC":                                 runtime.GC,
	"runtime.GOMAXPROCS":                         runtime.GOMAXPROCS,
	"runtime.GOROOT":                             runtime.GOROOT,
	"runtime.Goexit":                             runtime.Goexit,
	"runtime.GoroutineProfile":                   runtime.GoroutineProfile,
	"runtime.Gosched":                            runtime.Gosched,
	"runtime.KeepAlive":                          runtime.KeepAlive,
	"runtime.LockOSThread":                       runtime.LockOSThread,
	"runtime.MemProfile":                         runtime.MemProfile,
	"runtime.MemProfileRate":                     &runtime.MemProfileRate,
	"runtime.MutexProfile":                       runtime.MutexProfile,
	"runtime.NumCPU":                             runtime.NumCPU,
	"runtime.NumCgoCall":                         runtime.NumCgoCall,
	"runtime.NumGoroutine":                       runtime.NumGoroutine,
	"runtime.ReadMemStats":                       runtime.ReadMemStats,
	"runtime.ReadTrace":                          runtime.ReadTrace,
	"runtime.SetBlockProfileRate":                runtime.SetBlockProfileRate,
	"runtime.SetCPUProfileRate":                  runtime.SetCPUProfileRate,
	"runtime.SetCgoTraceback":                    runtime.SetCgoTraceback,
	"runtime.SetFinalizer":                       runtime.SetFinalizer,
	"runtime.SetMutexProfileFraction":            runtime.SetMutexProfileFraction,
	"runtime.Stack":                              runtime.Stack,
	"runtime.StartTrace":                         runtime.StartTrace,
	"runtime.StopTrace":                          runtime.StopTrace,
	"runtime.ThreadCreateProfile":                runtime.ThreadCreateProfile,
	"runtime.UnlockOSThread":                     runtime.UnlockOSThread,
	"runtime.Version":                            runtime.Version,
}

var typList = []interface{}{
	(*runtime.BlockProfileRecord)(nil),
	(*runtime.Error)(nil),
	(*runtime.Frame)(nil),
	(*runtime.Frames)(nil),
	(*runtime.Func)(nil),
	(*runtime.MemProfileRecord)(nil),
	(*runtime.MemStats)(nil),
	(*runtime.StackRecord)(nil),
	(*runtime.TypeAssertionError)(nil),
}
