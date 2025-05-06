// export by github.com/goplus/igop/cmd/qexp

//go:build go1.23 && !go1.24
// +build go1.23,!go1.24

package coverage

import (
	q "runtime/coverage"

	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "coverage",
		Path: "runtime/coverage",
		Deps: map[string]string{
			"internal/coverage/cfile": "cfile",
			"io":                      "io",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"ClearCounters":    reflect.ValueOf(q.ClearCounters),
			"WriteCounters":    reflect.ValueOf(q.WriteCounters),
			"WriteCountersDir": reflect.ValueOf(q.WriteCountersDir),
			"WriteMeta":        reflect.ValueOf(q.WriteMeta),
			"WriteMetaDir":     reflect.ValueOf(q.WriteMetaDir),
		},
		TypedConsts:   map[string]igop.TypedConst{},
		UntypedConsts: map[string]igop.UntypedConst{},
	})
}
