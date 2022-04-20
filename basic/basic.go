package basic

import (
	"reflect"
	"unsafe"
)

type eface struct {
	typ  unsafe.Pointer
	word unsafe.Pointer
}

type Type unsafe.Pointer

func TypeOf(i interface{}) Type {
	p := (*eface)(unsafe.Pointer(&i))
	return Type(p.typ)
}

func TypeOfType(typ reflect.Type) Type {
	e := (*eface)(unsafe.Pointer(&typ))
	return Type(e.word)
}

func Bool(i interface{}) bool {
	return *(*bool)((*eface)(unsafe.Pointer(&i)).word)
}

func Int(i interface{}) int {
	return *(*int)((*eface)(unsafe.Pointer(&i)).word)
}

func Int8(i interface{}) int8 {
	return *(*int8)((*eface)(unsafe.Pointer(&i)).word)
}

func Int16(i interface{}) int16 {
	return *(*int16)((*eface)(unsafe.Pointer(&i)).word)
}

func Int32(i interface{}) int32 {
	return *(*int32)((*eface)(unsafe.Pointer(&i)).word)
}

func Int64(i interface{}) int64 {
	return *(*int64)((*eface)(unsafe.Pointer(&i)).word)
}

func Uint(i interface{}) uint {
	return *(*uint)((*eface)(unsafe.Pointer(&i)).word)
}

func Uint8(i interface{}) uint8 {
	return *(*uint8)((*eface)(unsafe.Pointer(&i)).word)
}

func Uint16(i interface{}) uint16 {
	return *(*uint16)((*eface)(unsafe.Pointer(&i)).word)
}

func Uint32(i interface{}) uint32 {
	return *(*uint32)((*eface)(unsafe.Pointer(&i)).word)
}

func Uint64(i interface{}) uint64 {
	return *(*uint64)((*eface)(unsafe.Pointer(&i)).word)
}

func Uintptr(i interface{}) uintptr {
	return *(*uintptr)((*eface)(unsafe.Pointer(&i)).word)
}

func Float32(i interface{}) float32 {
	return *(*float32)((*eface)(unsafe.Pointer(&i)).word)
}

func Float64(i interface{}) float64 {
	return *(*float64)((*eface)(unsafe.Pointer(&i)).word)
}

func Complex64(i interface{}) complex64 {
	return *(*complex64)((*eface)(unsafe.Pointer(&i)).word)
}

func Complex128(i interface{}) complex128 {
	return *(*complex128)((*eface)(unsafe.Pointer(&i)).word)
}

func String(i interface{}) string {
	return *(*string)((*eface)(unsafe.Pointer(&i)).word)
}

func Make(typ Type, i interface{}) interface{} {
	p := (*eface)(unsafe.Pointer(&i))
	p.typ = unsafe.Pointer(typ)
	return i
}
