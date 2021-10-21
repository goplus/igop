package reflect

import (
	"reflect"
	"unsafe"
)

func fromReflectType(typ reflect.Type) *rtype {
	e := (*emptyInterface)(unsafe.Pointer(&typ))
	return (*rtype)(unsafe.Pointer(e.word))
}

//go:linkname toReflectType reflect.toType
func toReflectType(t *rtype) reflect.Type

func toReflectValue(v Value) reflect.Value {
	return *(*reflect.Value)(unsafe.Pointer(&v))
}

func AllMethod(typ reflect.Type) (ms []reflect.Method) {
	if typ.Kind() == reflect.Interface {
		n := typ.NumMethod()
		ms = make([]reflect.Method, n, n)
		for i := 0; i < n; i++ {
			ms[i] = typ.Method(i)
		}
		return ms
	}
	rt := fromReflectType(typ)
	n := rt.AllNumMethod()
	ms = make([]reflect.Method, n, n)
	for i := 0; i < n; i++ {
		ms[i] = rt.AllMethod(i)
	}
	return ms
}

func (t *rtype) AllMethod(i int) (m reflect.Method) {
	methods := t.allMethods()
	if i < 0 || i >= len(methods) {
		panic("reflect: Method index out of range")
	}
	p := methods[i]
	pname := t.nameOff(p.name)
	m.Name = pname.name()
	m.Index = i
	fl := flag(Func)
	mtyp := t.typeOff(p.mtyp)
	if mtyp == nil {
		return
	}
	ft := (*funcType)(unsafe.Pointer(mtyp))
	in := make([]reflect.Type, 0, 1+len(ft.in()))
	in = append(in, toReflectType(t))
	for _, arg := range ft.in() {
		in = append(in, toReflectType(arg))
	}
	out := make([]reflect.Type, 0, len(ft.out()))
	for _, ret := range ft.out() {
		out = append(out, toReflectType(ret))
	}
	mt := reflect.FuncOf(in, out, ft.IsVariadic())
	m.Type = mt
	tfn := t.textOff(p.tfn)
	fn := unsafe.Pointer(&tfn)
	m.Func = toReflectValue(Value{fromReflectType(mt), fn, fl})
	return m
}
