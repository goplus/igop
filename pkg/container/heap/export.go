// export by github.com/goplus/gossa/cmd/qexp

package heap

import (
	"container/heap"

	"github.com/goplus/gossa"
)

func init() {
	gossa.RegisterPackage("container/heap", extMap, typList)
}

var extMap = map[string]interface{}{
	"(container/heap.Interface).Len":  (heap.Interface).Len,
	"(container/heap.Interface).Less": (heap.Interface).Less,
	"(container/heap.Interface).Pop":  (heap.Interface).Pop,
	"(container/heap.Interface).Push": (heap.Interface).Push,
	"(container/heap.Interface).Swap": (heap.Interface).Swap,
	"container/heap.Fix":              heap.Fix,
	"container/heap.Init":             heap.Init,
	"container/heap.Pop":              heap.Pop,
	"container/heap.Push":             heap.Push,
	"container/heap.Remove":           heap.Remove,
}

var typList = []interface{}{
	(*heap.Interface)(nil),
}
