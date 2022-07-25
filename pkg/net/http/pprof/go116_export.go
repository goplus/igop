// export by github.com/goplus/igop/cmd/qexp

//+build go1.16,!go1.17

package pprof

import (
	q "net/http/pprof"

	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "pprof",
		Path: "net/http/pprof",
		Deps: map[string]string{
			"bufio":            "bufio",
			"bytes":            "bytes",
			"context":          "context",
			"fmt":              "fmt",
			"html":             "html",
			"internal/profile": "profile",
			"io":               "io",
			"log":              "log",
			"net/http":         "http",
			"net/url":          "url",
			"os":               "os",
			"runtime":          "runtime",
			"runtime/pprof":    "pprof",
			"runtime/trace":    "trace",
			"sort":             "sort",
			"strconv":          "strconv",
			"strings":          "strings",
			"time":             "time",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{},
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
		TypedConsts:   map[string]igop.TypedConst{},
		UntypedConsts: map[string]igop.UntypedConst{},
	})
}
