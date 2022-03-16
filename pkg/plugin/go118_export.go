// export by github.com/goplus/gossa/cmd/qexp

//go:build go1.18
// +build go1.18

package plugin

import (
	q "plugin"

	"reflect"

	"github.com/goplus/gossa"
)

func init() {
	gossa.RegisterPackage(&gossa.Package{
		Name: "plugin",
		Path: "plugin",
		Deps: map[string]string{
			"errors":      "errors",
			"runtime/cgo": "cgo",
			"sync":        "sync",
			"syscall":     "syscall",
			"unsafe":      "unsafe",
		},
		Interfaces: map[string]reflect.Type{
			"Symbol": reflect.TypeOf((*q.Symbol)(nil)).Elem(),
		},
		NamedTypes: map[string]gossa.NamedType{
			"Plugin": {reflect.TypeOf((*q.Plugin)(nil)).Elem(), "", "Lookup"},
		},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"Open": reflect.ValueOf(q.Open),
		},
		TypedConsts:   map[string]gossa.TypedConst{},
		UntypedConsts: map[string]gossa.UntypedConst{},
	})
}
