// export by github.com/goplus/gossa/cmd/qexp

package cgo

import (
	q "runtime/cgo"

	"reflect"

	"github.com/goplus/gossa"
)

func init() {
	gossa.RegisterPackage(&gossa.Package{
		Name: "cgo",
		Path: "runtime/cgo",
		Deps: map[string]string{
			"unsafe": "unsafe",
		},
		Interfaces:    map[string]reflect.Type{},
		NamedTypes:    map[string]gossa.NamedType{},
		AliasTypes:    map[string]reflect.Type{},
		Vars:          map[string]reflect.Value{},
		Funcs:         map[string]reflect.Value{},
		TypedConsts:   map[string]gossa.TypedConst{},
		UntypedConsts: map[string]gossa.UntypedConst{},
	})
}
