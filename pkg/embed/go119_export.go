// export by github.com/goplus/igop/cmd/qexp

//go:build go1.19 && !go1.20
// +build go1.19,!go1.20

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
		NamedTypes: map[string]reflect.Type{
			"FS": reflect.TypeOf((*q.FS)(nil)).Elem(),
		},
		AliasTypes:    map[string]reflect.Type{},
		Vars:          map[string]reflect.Value{},
		Funcs:         map[string]reflect.Value{},
		TypedConsts:   map[string]igop.TypedConst{},
		UntypedConsts: map[string]igop.UntypedConst{},
	})
}
