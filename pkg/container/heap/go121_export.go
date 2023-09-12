// export by github.com/goplus/igop/cmd/qexp

//go:build go1.21
// +build go1.21

package heap

import (
	q "container/heap"

	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "heap",
		Path: "container/heap",
		Deps: map[string]string{
			"sort": "sort",
		},
		Interfaces: map[string]reflect.Type{
			"Interface": reflect.TypeOf((*q.Interface)(nil)).Elem(),
		},
		NamedTypes: map[string]reflect.Type{},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"Fix":    reflect.ValueOf(q.Fix),
			"Init":   reflect.ValueOf(q.Init),
			"Pop":    reflect.ValueOf(q.Pop),
			"Push":   reflect.ValueOf(q.Push),
			"Remove": reflect.ValueOf(q.Remove),
		},
		TypedConsts:   map[string]igop.TypedConst{},
		UntypedConsts: map[string]igop.UntypedConst{},
	})
}
