// export by github.com/goplus/interp/cmd/qexp

package sort

import (
	"sort"

	"github.com/goplus/interp"
)

func init() {
	interp.RegisterPackage("sort", extMap, typList)
}

var extMap = map[string]interface{}{
	"(sort.Float64Slice).Len":    (sort.Float64Slice).Len,
	"(sort.Float64Slice).Less":   (sort.Float64Slice).Less,
	"(sort.Float64Slice).Search": (sort.Float64Slice).Search,
	"(sort.Float64Slice).Sort":   (sort.Float64Slice).Sort,
	"(sort.Float64Slice).Swap":   (sort.Float64Slice).Swap,
	"(sort.IntSlice).Len":        (sort.IntSlice).Len,
	"(sort.IntSlice).Less":       (sort.IntSlice).Less,
	"(sort.IntSlice).Search":     (sort.IntSlice).Search,
	"(sort.IntSlice).Sort":       (sort.IntSlice).Sort,
	"(sort.IntSlice).Swap":       (sort.IntSlice).Swap,
	"(sort.Interface).Len":       (sort.Interface).Len,
	"(sort.Interface).Less":      (sort.Interface).Less,
	"(sort.Interface).Swap":      (sort.Interface).Swap,
	"(sort.StringSlice).Len":     (sort.StringSlice).Len,
	"(sort.StringSlice).Less":    (sort.StringSlice).Less,
	"(sort.StringSlice).Search":  (sort.StringSlice).Search,
	"(sort.StringSlice).Sort":    (sort.StringSlice).Sort,
	"(sort.StringSlice).Swap":    (sort.StringSlice).Swap,
	"sort.Float64s":              sort.Float64s,
	"sort.Float64sAreSorted":     sort.Float64sAreSorted,
	"sort.Ints":                  sort.Ints,
	"sort.IntsAreSorted":         sort.IntsAreSorted,
	"sort.IsSorted":              sort.IsSorted,
	"sort.Reverse":               sort.Reverse,
	"sort.Search":                sort.Search,
	"sort.SearchFloat64s":        sort.SearchFloat64s,
	"sort.SearchInts":            sort.SearchInts,
	"sort.SearchStrings":         sort.SearchStrings,
	"sort.Slice":                 sort.Slice,
	"sort.SliceIsSorted":         sort.SliceIsSorted,
	"sort.SliceStable":           sort.SliceStable,
	"sort.Sort":                  sort.Sort,
	"sort.Stable":                sort.Stable,
	"sort.Strings":               sort.Strings,
	"sort.StringsAreSorted":      sort.StringsAreSorted,
}

var typList = []interface{}{
	(*sort.Float64Slice)(nil),
	(*sort.IntSlice)(nil),
	(*sort.Interface)(nil),
	(*sort.StringSlice)(nil),
}
