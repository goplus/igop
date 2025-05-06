// export by github.com/goplus/igop/cmd/qexp

//go:build go1.24
// +build go1.24

package ioutil

import (
	q "io/ioutil"

	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "ioutil",
		Path: "io/ioutil",
		Deps: map[string]string{
			"io":      "io",
			"io/fs":   "fs",
			"os":      "os",
			"slices":  "slices",
			"strings": "strings",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{},
		AliasTypes: map[string]reflect.Type{},
		Vars: map[string]reflect.Value{
			"Discard": reflect.ValueOf(&q.Discard),
		},
		Funcs: map[string]reflect.Value{
			"NopCloser": reflect.ValueOf(q.NopCloser),
			"ReadAll":   reflect.ValueOf(q.ReadAll),
			"ReadDir":   reflect.ValueOf(q.ReadDir),
			"ReadFile":  reflect.ValueOf(q.ReadFile),
			"TempDir":   reflect.ValueOf(q.TempDir),
			"TempFile":  reflect.ValueOf(q.TempFile),
			"WriteFile": reflect.ValueOf(q.WriteFile),
		},
		TypedConsts:   map[string]igop.TypedConst{},
		UntypedConsts: map[string]igop.UntypedConst{},
	})
}
