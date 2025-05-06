// export by github.com/goplus/igop/cmd/qexp

//go:build go1.23 && !go1.24
// +build go1.23,!go1.24

package importer

import (
	q "go/importer"

	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "importer",
		Path: "go/importer",
		Deps: map[string]string{
			"go/build":                  "build",
			"go/internal/gccgoimporter": "gccgoimporter",
			"go/internal/gcimporter":    "gcimporter",
			"go/internal/srcimporter":   "srcimporter",
			"go/token":                  "token",
			"go/types":                  "types",
			"io":                        "io",
			"runtime":                   "runtime",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{
			"Lookup": reflect.TypeOf((*q.Lookup)(nil)).Elem(),
		},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"Default":     reflect.ValueOf(q.Default),
			"For":         reflect.ValueOf(q.For),
			"ForCompiler": reflect.ValueOf(q.ForCompiler),
		},
		TypedConsts:   map[string]igop.TypedConst{},
		UntypedConsts: map[string]igop.UntypedConst{},
	})
}
