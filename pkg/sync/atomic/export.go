// export by github.com/goplus/interp/cmd/qexp

package atomic

import (
	"sync/atomic"

	"github.com/goplus/interp"
)

func init() {
	interp.RegisterPackage("sync/atomic", extMap, typList)
}

var extMap = map[string]interface{}{
	"(*sync/atomic.Value).Load":         (*atomic.Value).Load,
	"(*sync/atomic.Value).Store":        (*atomic.Value).Store,
	"sync/atomic.AddInt32":              atomic.AddInt32,
	"sync/atomic.AddInt64":              atomic.AddInt64,
	"sync/atomic.AddUint32":             atomic.AddUint32,
	"sync/atomic.AddUint64":             atomic.AddUint64,
	"sync/atomic.AddUintptr":            atomic.AddUintptr,
	"sync/atomic.CompareAndSwapInt32":   atomic.CompareAndSwapInt32,
	"sync/atomic.CompareAndSwapInt64":   atomic.CompareAndSwapInt64,
	"sync/atomic.CompareAndSwapPointer": atomic.CompareAndSwapPointer,
	"sync/atomic.CompareAndSwapUint32":  atomic.CompareAndSwapUint32,
	"sync/atomic.CompareAndSwapUint64":  atomic.CompareAndSwapUint64,
	"sync/atomic.CompareAndSwapUintptr": atomic.CompareAndSwapUintptr,
	"sync/atomic.LoadInt32":             atomic.LoadInt32,
	"sync/atomic.LoadInt64":             atomic.LoadInt64,
	"sync/atomic.LoadPointer":           atomic.LoadPointer,
	"sync/atomic.LoadUint32":            atomic.LoadUint32,
	"sync/atomic.LoadUint64":            atomic.LoadUint64,
	"sync/atomic.LoadUintptr":           atomic.LoadUintptr,
	"sync/atomic.StoreInt32":            atomic.StoreInt32,
	"sync/atomic.StoreInt64":            atomic.StoreInt64,
	"sync/atomic.StorePointer":          atomic.StorePointer,
	"sync/atomic.StoreUint32":           atomic.StoreUint32,
	"sync/atomic.StoreUint64":           atomic.StoreUint64,
	"sync/atomic.StoreUintptr":          atomic.StoreUintptr,
	"sync/atomic.SwapInt32":             atomic.SwapInt32,
	"sync/atomic.SwapInt64":             atomic.SwapInt64,
	"sync/atomic.SwapPointer":           atomic.SwapPointer,
	"sync/atomic.SwapUint32":            atomic.SwapUint32,
	"sync/atomic.SwapUint64":            atomic.SwapUint64,
	"sync/atomic.SwapUintptr":           atomic.SwapUintptr,
}

var typList = []interface{}{
	(*atomic.Value)(nil),
}
