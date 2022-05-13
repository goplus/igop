package basic_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/goplus/gossa/internal/basic"
)

func TestInt(t *testing.T) {
	type T int
	i1 := T(100)
	if n := basic.Int(i1); n != 100 {
		t.Fatal(n)
	}
	typ := basic.TypeOf(i1)
	i2 := basic.Make(typ, 200)
	if n := basic.Int(i2); n != 200 {
		t.Fatal(n)
	}
}

func TestFloat32(t *testing.T) {
	type T float32
	i1 := T(100.1)
	if n := basic.Float32(i1); n != 100.1 {
		t.Fatal("check", n)
	}
	typ := basic.TypeOf(i1)
	i2 := basic.Make(typ, float32(200.1))
	if n := basic.Float32(i2); n != 200.1 {
		t.Fatal("new at", n)
	}
}

func TestString(t *testing.T) {
	type T string
	i1 := T("hello")
	if n := basic.String(i1); n != "hello" {
		t.Fatal("check", n)
	}
	typ := basic.TypeOf(i1)
	i2 := basic.Make(typ, "world")
	if n := basic.String(i2); n != "world" {
		t.Fatal("new at", n)
	}
}

func TestNot(t *testing.T) {
	type T bool
	var b T
	if r := basic.Not(b); basic.Bool(r) != true {
		t.Fatal("must true")
	}
	b = true
	if r := basic.Not(b); basic.Bool(r) != false {
		t.Fatal("must false")
	}
}

func TestNegInt(t *testing.T) {
	type T int
	var n T = 100
	if r := basic.NegInt(n); basic.Int(r) != -100 {
		t.Fatal(r)
	}
	n = -200
	if r := basic.NegInt(n); basic.Int(r) != 200 {
		t.Fatal(r)
	}
}

func TestNegUint8(t *testing.T) {
	type T uint8
	var n T = 255
	if r := basic.NegUint8(n); basic.Uint8(r) != 1 {
		t.Fatal(r)
	}
	n = 1
	if r := basic.NegUint8(n); basic.Uint8(r) != 255 {
		t.Fatal(r)
	}
}

func TestNegFloat32(t *testing.T) {
	type T float32
	var n T = 123.456
	if r := basic.NegFloat32(n); basic.Float32(r) != -123.456 {
		t.Fatal(r)
	}
	n = -1
	if r := basic.NegFloat32(n); basic.Float32(r) != 1 {
		t.Fatal(r)
	}
}

func TestNegComplex64(t *testing.T) {
	type T complex64
	var n T = 1 + 2i
	if r := basic.NegComplex64(n); basic.Complex64(r) != -1-2i {
		t.Fatal(r)
	}
	n = -1
	if r := basic.NegComplex64(n); basic.Complex64(r) != 1 {
		t.Fatal()
	}
}

func TestXorInt8(t *testing.T) {
	type T int8
	var n T = 0b11
	if r := basic.XorInt8(n); basic.Int8(r) != -0b100 {
		t.Fatal(r)
	}
}

func TestXorUint8(t *testing.T) {
	type T uint
	var n T = 0b11
	if r := basic.XorUint8(n); basic.Uint8(r) != 0b11111100 {
		t.Fatal(r)
	}
}

func TestConvertInt(t *testing.T) {
	type T int
	var n int = 100
	typ := basic.TypeOf(T(0))
	if r := basic.ConvertInt(typ, n); basic.Int(r) != n {
		t.Fatal(r)
	}
}

func TestConvertFloat32(t *testing.T) {
	type T float32
	var n float32 = 100.1
	typ := basic.TypeOf(T(0))
	if r := basic.ConvertFloat32(typ, n); basic.Float32(r) != n {
		t.Fatal(r)
	}
}

func TestConvertComplex64(t *testing.T) {
	type T complex64
	var n complex64 = 1 + 2i
	typ := basic.TypeOf(T(0))
	if r := basic.ConvertComplex64(typ, n); basic.Complex64(r) != n {
		t.Fatal(r)
	}
}

func TestConvertString(t *testing.T) {
	type T string
	var n string = "hello"
	typ := basic.TypeOf(T(""))
	if r := basic.ConvertString(typ, n); basic.String(r) != n {
		t.Fatal(r)
	}
}

func TestMakeInt(t *testing.T) {
	type T int
	typ := basic.TypeOf(T(0))
	if r := basic.MakeInt(typ, 100); basic.Int(r) != 100 {
		t.Fatal(r)
	}
}

func TestMakeFloat32(t *testing.T) {
	type T float32
	typ := basic.TypeOf(T(0))
	if r := basic.MakeFloat32(typ, 100.1); basic.Float32(r) != 100.1 {
		t.Fatal(r)
	}
}

func TestAllocInt(t *testing.T) {
	type T int
	typ := basic.TypeOf(T(0))
	r := basic.Alloc(typ)
	if s := fmt.Sprintf("%T %#v", r, r); s != "basic_test.T 0" {
		panic(s)
	}
}

func TestAllocStruct(t *testing.T) {
	type T struct {
		X int
		Y int
	}
	typ := basic.TypeOf(T{})
	r := basic.Alloc(typ)
	if s := fmt.Sprintf("%#v", r); s != "basic_test.T{X:0, Y:0}" {
		panic(s)
	}
}

func TestAllocPtr(t *testing.T) {
	type T struct {
		X int
		Y int
	}
	typ := basic.TypeOf(&T{})
	r := basic.Alloc(typ)
	if s := fmt.Sprintf("%#v", r); s != "&basic_test.T{X:0, Y:0}" {
		panic(s)
	}
}

func TestAllocInterfce(t *testing.T) {
	type T = fmt.Stringer
	typ := basic.TypeOfType(reflect.TypeOf((*T)(nil)).Elem())
	r := basic.Alloc(typ)
	if s := fmt.Sprintf("%#v", r); s != "fmt.Stringer(nil)" {
		panic(s)
	}
}

func TestNewInt(t *testing.T) {
	type T int
	v := T(0)
	typ := basic.TypeOf(v)
	ptrto := basic.TypeOf(&v)
	r := basic.New(typ, ptrto)
	if s := fmt.Sprintf("%T %#v", r, *(r.(*T))); s != "*basic_test.T 0" {
		panic(s)
	}
}

func TestNewStruct(t *testing.T) {
	type T struct {
		X int
		Y int
	}
	var v T
	typ := basic.TypeOf(v)
	ptrto := basic.TypeOf(&v)
	r := basic.New(typ, ptrto)
	if s := fmt.Sprintf("%#v", r); s != "&basic_test.T{X:0, Y:0}" {
		panic(s)
	}
}

func TestNewInterfce(t *testing.T) {
	type T = fmt.Stringer
	rt := reflect.TypeOf((*T)(nil))
	typ := basic.TypeOfType(rt.Elem())
	ptrto := basic.TypeOfType(rt)
	r := basic.New(typ, ptrto)
	if s := fmt.Sprintf("%T %v", r, *(r.(*T))); s != "*fmt.Stringer <nil>" {
		panic(s)
	}
}
