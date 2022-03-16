// export by github.com/goplus/gossa/cmd/qexp

//go:build go1.18
// +build go1.18

package embed

import (
	q "embed"

	"reflect"

	"github.com/goplus/gossa"
)

func init() {
	gossa.RegisterPackage(&gossa.Package{
		Name: "embed",
		Path: "embed",
		Deps: map[string]string{
			"errors": "errors",
			"io":     "io",
			"io/fs":  "fs",
			"time":   "time",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]gossa.NamedType{
			"FS": {reflect.TypeOf((*q.FS)(nil)).Elem(), "Open,ReadDir,ReadFile,lookup,readDir", ""},
		},
		AliasTypes:    map[string]reflect.Type{},
		Vars:          map[string]reflect.Value{},
		Funcs:         map[string]reflect.Value{},
		TypedConsts:   map[string]gossa.TypedConst{},
		UntypedConsts: map[string]gossa.UntypedConst{},
	})
}
