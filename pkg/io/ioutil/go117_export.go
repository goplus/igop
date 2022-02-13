// export by github.com/goplus/gossa/cmd/qexp

//go:build go1.17
// +build go1.17

package ioutil

import (
	q "io/ioutil"

	"reflect"

	"github.com/goplus/gossa"
)

func init() {
	gossa.RegisterPackage(&gossa.Package{
		Name: "ioutil",
		Path: "io/ioutil",
		Deps: map[string]string{
			"io":    "io",
			"io/fs": "fs",
			"os":    "os",
			"sort":  "sort",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]gossa.NamedType{},
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
		TypedConsts:   map[string]gossa.TypedConst{},
		UntypedConsts: map[string]gossa.UntypedConst{},
	})
}
