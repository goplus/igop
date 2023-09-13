// export by github.com/goplus/igop/cmd/qexp

//go:build go1.21
// +build go1.21

package maps

import (
	_ "maps"
	"reflect"
	_ "unsafe"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name:       "maps",
		Path:       "maps",
		Deps:       map[string]string{},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"clone": reflect.ValueOf(_clone),
		},
		TypedConsts:   map[string]igop.TypedConst{},
		UntypedConsts: map[string]igop.UntypedConst{},
		Source:        source,
	})
}

//go:linkname _clone maps.clone
func _clone(m any) any

var source = "package maps\n\nfunc Equal[M1, M2 ~map[K]V, K, V comparable](m1 M1, m2 M2) bool {\n\tif len(m1) != len(m2) {\n\t\treturn false\n\t}\n\tfor k, v1 := range m1 {\n\t\tif v2, ok := m2[k]; !ok || v1 != v2 {\n\t\t\treturn false\n\t\t}\n\t}\n\treturn true\n}\n\nfunc EqualFunc[M1 ~map[K]V1, M2 ~map[K]V2, K comparable, V1, V2 any](m1 M1, m2 M2, eq func(V1, V2) bool) bool {\n\tif len(m1) != len(m2) {\n\t\treturn false\n\t}\n\tfor k, v1 := range m1 {\n\t\tif v2, ok := m2[k]; !ok || !eq(v1, v2) {\n\t\t\treturn false\n\t\t}\n\t}\n\treturn true\n}\n\nfunc clone(m any) any\n\nfunc Clone[M ~map[K]V, K comparable, V any](m M) M {\n\n\tif m == nil {\n\t\treturn nil\n\t}\n\treturn clone(m).(M)\n}\n\nfunc Copy[M1 ~map[K]V, M2 ~map[K]V, K comparable, V any](dst M1, src M2) {\n\tfor k, v := range src {\n\t\tdst[k] = v\n\t}\n}\n\nfunc DeleteFunc[M ~map[K]V, K comparable, V any](m M, del func(K, V) bool) {\n\tfor k, v := range m {\n\t\tif del(k, v) {\n\t\t\tdelete(m, k)\n\t\t}\n\t}\n}\n"
