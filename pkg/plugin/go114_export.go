// export by github.com/goplus/igop/cmd/qexp

//+build go1.14,!go1.15

package plugin

import (
	q "plugin"

	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
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
		NamedTypes: map[string]igop.NamedType{
			"Plugin": {reflect.TypeOf((*q.Plugin)(nil)).Elem(), "", "Lookup"},
		},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"Open": reflect.ValueOf(q.Open),
		},
		TypedConsts:   map[string]igop.TypedConst{},
		UntypedConsts: map[string]igop.UntypedConst{},
	})
}
