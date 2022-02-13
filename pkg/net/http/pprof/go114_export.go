// export by github.com/goplus/gossa/cmd/qexp

//+build go1.14,!go1.15

package pprof

import (
	q "net/http/pprof"

	"reflect"

	"github.com/goplus/gossa"
)

func init() {
	gossa.RegisterPackage(&gossa.Package{
		Name: "pprof",
		Path: "net/http/pprof",
		Deps: map[string]string{
			"bufio":         "bufio",
			"bytes":         "bytes",
			"fmt":           "fmt",
			"html/template": "template",
			"io":            "io",
			"log":           "log",
			"net/http":      "http",
			"os":            "os",
			"runtime":       "runtime",
			"runtime/pprof": "pprof",
			"runtime/trace": "trace",
			"sort":          "sort",
			"strconv":       "strconv",
			"strings":       "strings",
			"time":          "time",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]gossa.NamedType{},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"Cmdline": reflect.ValueOf(q.Cmdline),
			"Handler": reflect.ValueOf(q.Handler),
			"Index":   reflect.ValueOf(q.Index),
			"Profile": reflect.ValueOf(q.Profile),
			"Symbol":  reflect.ValueOf(q.Symbol),
			"Trace":   reflect.ValueOf(q.Trace),
		},
		TypedConsts:   map[string]gossa.TypedConst{},
		UntypedConsts: map[string]gossa.UntypedConst{},
	})
}
