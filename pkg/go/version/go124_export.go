// export by github.com/goplus/igop/cmd/qexp

//go:build go1.24
// +build go1.24

package version

import (
	q "go/version"

	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "version",
		Path: "go/version",
		Deps: map[string]string{
			"internal/gover": "gover",
			"strings":        "strings",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"Compare": reflect.ValueOf(q.Compare),
			"IsValid": reflect.ValueOf(q.IsValid),
			"Lang":    reflect.ValueOf(q.Lang),
		},
		TypedConsts:   map[string]igop.TypedConst{},
		UntypedConsts: map[string]igop.UntypedConst{},
	})
}
