// export by github.com/goplus/gossa/cmd/qexp

//+build go1.16,!go1.17

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
