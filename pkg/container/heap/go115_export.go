// export by github.com/goplus/gossa/cmd/qexp

//+build go1.15,!go1.16

package heap

import (
	q "container/heap"

	"reflect"

	"github.com/goplus/gossa"
)

func init() {
	gossa.RegisterPackage(&gossa.Package{
		Name: "heap",
		Path: "container/heap",
		Deps: map[string]string{
			"sort": "sort",
		},
		Interfaces: map[string]reflect.Type{
			"Interface": reflect.TypeOf((*q.Interface)(nil)).Elem(),
		},
		NamedTypes: map[string]gossa.NamedType{},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"Fix":    reflect.ValueOf(q.Fix),
			"Init":   reflect.ValueOf(q.Init),
			"Pop":    reflect.ValueOf(q.Pop),
			"Push":   reflect.ValueOf(q.Push),
			"Remove": reflect.ValueOf(q.Remove),
		},
		TypedConsts:   map[string]gossa.TypedConst{},
		UntypedConsts: map[string]gossa.UntypedConst{},
	})
}
