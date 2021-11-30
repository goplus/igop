// export by github.com/goplus/gossa/cmd/qexp

package sort

import (
	q "sort"

	"reflect"

	"github.com/goplus/gossa"
)

func init() {
	gossa.RegisterPackage(&gossa.Package{
		Name: "sort",
		Path: "sort",
		Deps: map[string]string{
			"internal/reflectlite": "reflectlite",
		},
		Interfaces: map[string]reflect.Type{
			"Interface": reflect.TypeOf((*q.Interface)(nil)).Elem(),
		},
		NamedTypes: map[string]gossa.NamedType{
			"Float64Slice": {reflect.TypeOf((*q.Float64Slice)(nil)).Elem(), "Len,Less,Search,Sort,Swap", ""},
			"IntSlice":     {reflect.TypeOf((*q.IntSlice)(nil)).Elem(), "Len,Less,Search,Sort,Swap", ""},
			"StringSlice":  {reflect.TypeOf((*q.StringSlice)(nil)).Elem(), "Len,Less,Search,Sort,Swap", ""},
		},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"Float64s":          reflect.ValueOf(q.Float64s),
			"Float64sAreSorted": reflect.ValueOf(q.Float64sAreSorted),
			"Ints":              reflect.ValueOf(q.Ints),
			"IntsAreSorted":     reflect.ValueOf(q.IntsAreSorted),
			"IsSorted":          reflect.ValueOf(q.IsSorted),
			"Reverse":           reflect.ValueOf(q.Reverse),
			"Search":            reflect.ValueOf(q.Search),
			"SearchFloat64s":    reflect.ValueOf(q.SearchFloat64s),
			"SearchInts":        reflect.ValueOf(q.SearchInts),
			"SearchStrings":     reflect.ValueOf(q.SearchStrings),
			"Slice":             reflect.ValueOf(q.Slice),
			"SliceIsSorted":     reflect.ValueOf(q.SliceIsSorted),
			"SliceStable":       reflect.ValueOf(q.SliceStable),
			"Sort":              reflect.ValueOf(q.Sort),
			"Stable":            reflect.ValueOf(q.Stable),
			"Strings":           reflect.ValueOf(q.Strings),
			"StringsAreSorted":  reflect.ValueOf(q.StringsAreSorted),
		},
		TypedConsts:   map[string]gossa.TypedConst{},
		UntypedConsts: map[string]gossa.UntypedConst{},
	})
}
