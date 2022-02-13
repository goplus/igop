// export by github.com/goplus/gossa/cmd/qexp

//+build go1.14,!go1.15

package atomic

import (
	q "sync/atomic"

	"reflect"

	"github.com/goplus/gossa"
)

func init() {
	gossa.RegisterPackage(&gossa.Package{
		Name: "atomic",
		Path: "sync/atomic",
		Deps: map[string]string{
			"unsafe": "unsafe",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]gossa.NamedType{
			"Value": {reflect.TypeOf((*q.Value)(nil)).Elem(), "", "Load,Store"},
		},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"AddInt32":              reflect.ValueOf(q.AddInt32),
			"AddInt64":              reflect.ValueOf(q.AddInt64),
			"AddUint32":             reflect.ValueOf(q.AddUint32),
			"AddUint64":             reflect.ValueOf(q.AddUint64),
			"AddUintptr":            reflect.ValueOf(q.AddUintptr),
			"CompareAndSwapInt32":   reflect.ValueOf(q.CompareAndSwapInt32),
			"CompareAndSwapInt64":   reflect.ValueOf(q.CompareAndSwapInt64),
			"CompareAndSwapPointer": reflect.ValueOf(q.CompareAndSwapPointer),
			"CompareAndSwapUint32":  reflect.ValueOf(q.CompareAndSwapUint32),
			"CompareAndSwapUint64":  reflect.ValueOf(q.CompareAndSwapUint64),
			"CompareAndSwapUintptr": reflect.ValueOf(q.CompareAndSwapUintptr),
			"LoadInt32":             reflect.ValueOf(q.LoadInt32),
			"LoadInt64":             reflect.ValueOf(q.LoadInt64),
			"LoadPointer":           reflect.ValueOf(q.LoadPointer),
			"LoadUint32":            reflect.ValueOf(q.LoadUint32),
			"LoadUint64":            reflect.ValueOf(q.LoadUint64),
			"LoadUintptr":           reflect.ValueOf(q.LoadUintptr),
			"StoreInt32":            reflect.ValueOf(q.StoreInt32),
			"StoreInt64":            reflect.ValueOf(q.StoreInt64),
			"StorePointer":          reflect.ValueOf(q.StorePointer),
			"StoreUint32":           reflect.ValueOf(q.StoreUint32),
			"StoreUint64":           reflect.ValueOf(q.StoreUint64),
			"StoreUintptr":          reflect.ValueOf(q.StoreUintptr),
			"SwapInt32":             reflect.ValueOf(q.SwapInt32),
			"SwapInt64":             reflect.ValueOf(q.SwapInt64),
			"SwapPointer":           reflect.ValueOf(q.SwapPointer),
			"SwapUint32":            reflect.ValueOf(q.SwapUint32),
			"SwapUint64":            reflect.ValueOf(q.SwapUint64),
			"SwapUintptr":           reflect.ValueOf(q.SwapUintptr),
		},
		TypedConsts:   map[string]gossa.TypedConst{},
		UntypedConsts: map[string]gossa.UntypedConst{},
	})
}
