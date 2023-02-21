// export by github.com/goplus/igop/cmd/qexp

//go:build go1.20
// +build go1.20

package atomic

import (
	q "sync/atomic"

	"reflect"
	_ "unsafe"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "atomic",
		Path: "sync/atomic",
		Deps: map[string]string{
			"unsafe": "unsafe",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{
			"Bool":    reflect.TypeOf((*q.Bool)(nil)).Elem(),
			"Int32":   reflect.TypeOf((*q.Int32)(nil)).Elem(),
			"Int64":   reflect.TypeOf((*q.Int64)(nil)).Elem(),
			"Uint32":  reflect.TypeOf((*q.Uint32)(nil)).Elem(),
			"Uint64":  reflect.TypeOf((*q.Uint64)(nil)).Elem(),
			"Uintptr": reflect.TypeOf((*q.Uintptr)(nil)).Elem(),
			"Value":   reflect.TypeOf((*q.Value)(nil)).Elem(),
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
			"runtime_procPin":       reflect.ValueOf(_runtime_procPin),
			"runtime_procUnpin":     reflect.ValueOf(_runtime_procUnpin),
		},
		TypedConsts:   map[string]igop.TypedConst{},
		UntypedConsts: map[string]igop.UntypedConst{},
		Source:        source,
	})
}

//go:linkname _runtime_procPin sync/atomic.runtime_procPin
func _runtime_procPin()

//go:linkname _runtime_procUnpin sync/atomic.runtime_procUnpin
func _runtime_procUnpin()

var source = `package atomic
import "unsafe"
func SwapInt32(addr *int32, new int32) (old int32)
func SwapInt64(addr *int64, new int64) (old int64)
func SwapUint32(addr *uint32, new uint32) (old uint32)
func SwapUint64(addr *uint64, new uint64) (old uint64)
func SwapUintptr(addr *uintptr, new uintptr) (old uintptr)
func SwapPointer(addr *unsafe.Pointer, new unsafe.Pointer) (old unsafe.Pointer)
func CompareAndSwapInt32(addr *int32, old, new int32) (swapped bool)
func CompareAndSwapInt64(addr *int64, old, new int64) (swapped bool)
func CompareAndSwapUint32(addr *uint32, old, new uint32) (swapped bool)
func CompareAndSwapUint64(addr *uint64, old, new uint64) (swapped bool)
func CompareAndSwapUintptr(addr *uintptr, old, new uintptr) (swapped bool)
func CompareAndSwapPointer(addr *unsafe.Pointer, old, new unsafe.Pointer) (swapped bool)
func AddInt32(addr *int32, delta int32) (new int32)
func AddUint32(addr *uint32, delta uint32) (new uint32)
func AddInt64(addr *int64, delta int64) (new int64)
func AddUint64(addr *uint64, delta uint64) (new uint64)
func AddUintptr(addr *uintptr, delta uintptr) (new uintptr)
func LoadInt32(addr *int32) (val int32)
func LoadInt64(addr *int64) (val int64)
func LoadUint32(addr *uint32) (val uint32)
func LoadUint64(addr *uint64) (val uint64)
func LoadUintptr(addr *uintptr) (val uintptr)
func LoadPointer(addr *unsafe.Pointer) (val unsafe.Pointer)
func StoreInt32(addr *int32, val int32)
func StoreInt64(addr *int64, val int64)
func StoreUint32(addr *uint32, val uint32)
func StoreUint64(addr *uint64, val uint64)
func StoreUintptr(addr *uintptr, val uintptr)
func StorePointer(addr *unsafe.Pointer, val unsafe.Pointer)
type Bool struct {
	_	noCopy
	v	uint32
}
func (x *Bool) Load() bool
func (x *Bool) Store(val bool)
func (x *Bool) Swap(new bool) (old bool)
func (x *Bool) CompareAndSwap(old, new bool) (swapped bool)
func b32(b bool) uint32 {
	if b {
		return 1
	}
	return 0
}
var _ = &Pointer[int]{}
type Pointer[T any] struct {
	_	[0]*T
	_	noCopy
	v	unsafe.Pointer
}
func (x *Pointer[T]) Load() *T	{ return (*T)(LoadPointer(&x.v)) }
func (x *Pointer[T]) Store(val *T)	{ StorePointer(&x.v, unsafe.Pointer(val)) }
func (x *Pointer[T]) Swap(new *T) (old *T)	{ return (*T)(SwapPointer(&x.v, unsafe.Pointer(new))) }
func (x *Pointer[T]) CompareAndSwap(old, new *T) (swapped bool) {
	return CompareAndSwapPointer(&x.v, unsafe.Pointer(old), unsafe.Pointer(new))
}
type Int32 struct {
	_	noCopy
	v	int32
}
func (x *Int32) Load() int32
func (x *Int32) Store(val int32)
func (x *Int32) Swap(new int32) (old int32)
func (x *Int32) CompareAndSwap(old, new int32) (swapped bool)
func (x *Int32) Add(delta int32) (new int32)
type Int64 struct {
	_	noCopy
	_	align64
	v	int64
}
func (x *Int64) Load() int64
func (x *Int64) Store(val int64)
func (x *Int64) Swap(new int64) (old int64)
func (x *Int64) CompareAndSwap(old, new int64) (swapped bool)
func (x *Int64) Add(delta int64) (new int64)
type Uint32 struct {
	_	noCopy
	v	uint32
}
func (x *Uint32) Load() uint32
func (x *Uint32) Store(val uint32)
func (x *Uint32) Swap(new uint32) (old uint32)
func (x *Uint32) CompareAndSwap(old, new uint32) (swapped bool)
func (x *Uint32) Add(delta uint32) (new uint32)
type Uint64 struct {
	_	noCopy
	_	align64
	v	uint64
}
func (x *Uint64) Load() uint64
func (x *Uint64) Store(val uint64)
func (x *Uint64) Swap(new uint64) (old uint64)
func (x *Uint64) CompareAndSwap(old, new uint64) (swapped bool)
func (x *Uint64) Add(delta uint64) (new uint64)
type Uintptr struct {
	_	noCopy
	v	uintptr
}
func (x *Uintptr) Load() uintptr
func (x *Uintptr) Store(val uintptr)
func (x *Uintptr) Swap(new uintptr) (old uintptr)
func (x *Uintptr) CompareAndSwap(old, new uintptr) (swapped bool)
func (x *Uintptr) Add(delta uintptr) (new uintptr)
type noCopy struct{}
func (*noCopy) Lock()	{}
func (*noCopy) Unlock()	{}
type align64 struct{}
type Value struct {
	v any
}
type efaceWords struct {
	typ	unsafe.Pointer
	data	unsafe.Pointer
}
func (v *Value) Load() (val any)
var firstStoreInProgress byte
func (v *Value) Store(val any)
func (v *Value) Swap(new any) (old any)
func (v *Value) CompareAndSwap(old, new any) (swapped bool)
func runtime_procPin()
func runtime_procUnpin()
`
