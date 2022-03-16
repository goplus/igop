// export by github.com/goplus/gossa/cmd/qexp

//go:build go1.18
// +build go1.18

package trace

import (
	q "runtime/trace"

	"reflect"

	"github.com/goplus/gossa"
)

func init() {
	gossa.RegisterPackage(&gossa.Package{
		Name: "trace",
		Path: "runtime/trace",
		Deps: map[string]string{
			"context":     "context",
			"fmt":         "fmt",
			"io":          "io",
			"runtime":     "runtime",
			"sync":        "sync",
			"sync/atomic": "atomic",
			"unsafe":      "unsafe",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]gossa.NamedType{
			"Region": {reflect.TypeOf((*q.Region)(nil)).Elem(), "", "End"},
			"Task":   {reflect.TypeOf((*q.Task)(nil)).Elem(), "", "End"},
		},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"IsEnabled":   reflect.ValueOf(q.IsEnabled),
			"Log":         reflect.ValueOf(q.Log),
			"Logf":        reflect.ValueOf(q.Logf),
			"NewTask":     reflect.ValueOf(q.NewTask),
			"Start":       reflect.ValueOf(q.Start),
			"StartRegion": reflect.ValueOf(q.StartRegion),
			"Stop":        reflect.ValueOf(q.Stop),
			"WithRegion":  reflect.ValueOf(q.WithRegion),
		},
		TypedConsts:   map[string]gossa.TypedConst{},
		UntypedConsts: map[string]gossa.UntypedConst{},
	})
}
