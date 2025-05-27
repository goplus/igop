// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.21 && !go1.22
// +build go1.21,!go1.22

package embed

import (
	q "embed"

	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage(&ixgo.Package{
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
		TypedConsts:   map[string]ixgo.TypedConst{},
		UntypedConsts: map[string]ixgo.UntypedConst{},
	})
}
