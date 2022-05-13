package xtype_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/goplus/gossa/internal/xtype"
)

func TestInt(t *testing.T) {
	type T int
	i1 := T(100)
	if n := xtype.Int(i1); n != 100 {
		t.Fatal(n)
	}
	typ := xtype.TypeOf(i1)
	i2 := xtype.Make(typ, 200)
	if n := xtype.Int(i2); n != 200 {
		t.Fatal(n)
	}
}

func TestFloat32(t *testing.T) {
	type T float32
	i1 := T(100.1)
	if n := xtype.Float32(i1); n != 100.1 {
		t.Fatal("check", n)
	}
	typ := xtype.TypeOf(i1)
	i2 := xtype.Make(typ, float32(200.1))
	if n := xtype.Float32(i2); n != 200.1 {
		t.Fatal("new at", n)
	}
}

func TestString(t *testing.T) {
	type T string
	i1 := T("hello")
	if n := xtype.String(i1); n != "hello" {
		t.Fatal("check", n)
	}
	typ := xtype.TypeOf(i1)
	i2 := xtype.Make(typ, "world")
	if n := xtype.String(i2); n != "world" {
		t.Fatal("new at", n)
	}
}

func TestNot(t *testing.T) {
	type T bool
	var b T
	if r := xtype.Not(b); xtype.Bool(r) != true {
		t.Fatal("must true")
	}
	b = true
	if r := xtype.Not(b); xtype.Bool(r) != false {
		t.Fatal("must false")
	}
}

func TestNegInt(t *testing.T) {
	type T int
	var n T = 100
	if r := xtype.NegInt(n); xtype.Int(r) != -100 {
		t.Fatal(r)
	}
	n = -200
	if r := xtype.NegInt(n); xtype.Int(r) != 200 {
		t.Fatal(r)
	}
}

func TestNegUint8(t *testing.T) {
	type T uint8
	var n T = 255
	if r := xtype.NegUint8(n); xtype.Uint8(r) != 1 {
		t.Fatal(r)
	}
	n = 1
	if r := xtype.NegUint8(n); xtype.Uint8(r) != 255 {
		t.Fatal(r)
	}
}

func TestNegFloat32(t *testing.T) {
	type T float32
	var n T = 123.456
	if r := xtype.NegFloat32(n); xtype.Float32(r) != -123.456 {
		t.Fatal(r)
	}
	n = -1
	if r := xtype.NegFloat32(n); xtype.Float32(r) != 1 {
		t.Fatal(r)
	}
}

func TestNegComplex64(t *testing.T) {
	type T complex64
	var n T = 1 + 2i
	if r := xtype.NegComplex64(n); xtype.Complex64(r) != -1-2i {
		t.Fatal(r)
	}
	n = -1
	if r := xtype.NegComplex64(n); xtype.Complex64(r) != 1 {
		t.Fatal()
	}
}

func TestXorInt8(t *testing.T) {
	type T int8
	var n T = 0b11
	if r := xtype.XorInt8(n); xtype.Int8(r) != -0b100 {
		t.Fatal(r)
	}
}

func TestXorUint8(t *testing.T) {
	type T uint
	var n T = 0b11
	if r := xtype.XorUint8(n); xtype.Uint8(r) != 0b11111100 {
		t.Fatal(r)
	}
}

func TestConvertInt(t *testing.T) {
	type T int
	var n int = 100
	typ := xtype.TypeOf(T(0))
	if r := xtype.ConvertInt(typ, n); xtype.Int(r) != n {
		t.Fatal(r)
	}
}

func TestConvertFloat32(t *testing.T) {
	type T float32
	var n float32 = 100.1
	typ := xtype.TypeOf(T(0))
	if r := xtype.ConvertFloat32(typ, n); xtype.Float32(r) != n {
		t.Fatal(r)
	}
}

func TestConvertComplex64(t *testing.T) {
	type T complex64
	var n complex64 = 1 + 2i
	typ := xtype.TypeOf(T(0))
	if r := xtype.ConvertComplex64(typ, n); xtype.Complex64(r) != n {
		t.Fatal(r)
	}
}

func TestConvertString(t *testing.T) {
	type T string
	var n string = "hello"
	typ := xtype.TypeOf(T(""))
	if r := xtype.ConvertString(typ, n); xtype.String(r) != n {
		t.Fatal(r)
	}
}

func TestMakeInt(t *testing.T) {
	type T int
	typ := xtype.TypeOf(T(0))
	if r := xtype.MakeInt(typ, 100); xtype.Int(r) != 100 {
		t.Fatal(r)
	}
}

func TestMakeFloat32(t *testing.T) {
	type T float32
	typ := xtype.TypeOf(T(0))
	if r := xtype.MakeFloat32(typ, 100.1); xtype.Float32(r) != 100.1 {
		t.Fatal(r)
	}
}

func TestAllocInt(t *testing.T) {
	type T int
	typ := xtype.TypeOf(T(0))
	r := xtype.Alloc(typ)
	if s := fmt.Sprintf("%T %#v", r, r); s != "xtype_test.T 0" {
		panic(s)
	}
}

func TestAllocStruct(t *testing.T) {
	type T struct {
		X int
		Y int
	}
	typ := xtype.TypeOf(T{})
	r := xtype.Alloc(typ)
	if s := fmt.Sprintf("%#v", r); s != "xtype_test.T{X:0, Y:0}" {
		panic(s)
	}
}

func TestAllocPtr(t *testing.T) {
	type T struct {
		X int
		Y int
	}
	typ := xtype.TypeOf(&T{})
	r := xtype.Alloc(typ)
	if s := fmt.Sprintf("%#v", r); s != "&xtype_test.T{X:0, Y:0}" {
		panic(s)
	}
}

func TestAllocInterfce(t *testing.T) {
	type T = fmt.Stringer
	typ := xtype.TypeOfType(reflect.TypeOf((*T)(nil)).Elem())
	r := xtype.Alloc(typ)
	if s := fmt.Sprintf("%#v", r); s != "fmt.Stringer(nil)" {
		panic(s)
	}
}

func TestNewInt(t *testing.T) {
	type T int
	v := T(0)
	typ := xtype.TypeOf(v)
	ptrto := xtype.TypeOf(&v)
	r := xtype.New(typ, ptrto)
	if s := fmt.Sprintf("%T %#v", r, *(r.(*T))); s != "*xtype_test.T 0" {
		panic(s)
	}
}

func TestNewStruct(t *testing.T) {
	type T struct {
		X int
		Y int
	}
	var v T
	typ := xtype.TypeOf(v)
	ptrto := xtype.TypeOf(&v)
	r := xtype.New(typ, ptrto)
	if s := fmt.Sprintf("%#v", r); s != "&xtype_test.T{X:0, Y:0}" {
		panic(s)
	}
}

func TestNewInterfce(t *testing.T) {
	type T = fmt.Stringer
	rt := reflect.TypeOf((*T)(nil))
	typ := xtype.TypeOfType(rt.Elem())
	ptrto := xtype.TypeOfType(rt)
	r := xtype.New(typ, ptrto)
	if s := fmt.Sprintf("%T %v", r, *(r.(*T))); s != "*fmt.Stringer <nil>" {
		panic(s)
	}
}
