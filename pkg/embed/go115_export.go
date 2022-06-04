// export by github.com/goplus/igop/cmd/qexp

//+build go1.16,!go1.17

package embed

import (
	q "embed"

	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "embed",
		Path: "embed",
		Deps: map[string]string{
			"errors": "errors",
			"io":     "io",
			"io/fs":  "fs",
			"time":   "time",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]igop.NamedType{
			"FS": {reflect.TypeOf((*q.FS)(nil)).Elem(), "Open,ReadDir,ReadFile,lookup,readDir", ""},
		},
		AliasTypes:    map[string]reflect.Type{},
		Vars:          map[string]reflect.Value{},
		Funcs:         map[string]reflect.Value{},
		TypedConsts:   map[string]igop.TypedConst{},
		UntypedConsts: map[string]igop.UntypedConst{},
	})
}
