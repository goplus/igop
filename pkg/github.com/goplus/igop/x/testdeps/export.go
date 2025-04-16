// export by github.com/goplus/igop/cmd/qexp

package testdeps

import (
	q "github.com/goplus/igop/x/testdeps"

	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "testdeps",
		Path: "github.com/goplus/igop/x/testdeps",
		Deps: map[string]string{
			"bufio":                            "bufio",
			"github.com/goplus/igop/x/testlog": "testlog",
			"io":                               "io",
			"os":                               "os",
			"reflect":                          "reflect",
			"regexp":                           "regexp",
			"runtime/pprof":                    "pprof",
			"strings":                          "strings",
			"sync":                             "sync",
			"time":                             "time",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{
			"TestDeps": reflect.TypeOf((*q.TestDeps)(nil)).Elem(),
		},
		AliasTypes: map[string]reflect.Type{
			"CorpusEntry": reflect.TypeOf((*q.CorpusEntry)(nil)).Elem(),
		},
		Vars: map[string]reflect.Value{
			"CoverMarkProfileEmittedFunc": reflect.ValueOf(&q.CoverMarkProfileEmittedFunc),
			"CoverMode":                   reflect.ValueOf(&q.CoverMode),
			"CoverProcessTestDirFunc":     reflect.ValueOf(&q.CoverProcessTestDirFunc),
			"CoverSelectedPackages":       reflect.ValueOf(&q.CoverSelectedPackages),
			"CoverSnapshotFunc":           reflect.ValueOf(&q.CoverSnapshotFunc),
			"Covered":                     reflect.ValueOf(&q.Covered),
			"ImportPath":                  reflect.ValueOf(&q.ImportPath),
		},
		Funcs:         map[string]reflect.Value{},
		TypedConsts:   map[string]igop.TypedConst{},
		UntypedConsts: map[string]igop.UntypedConst{},
	})
}
