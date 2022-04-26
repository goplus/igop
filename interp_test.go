// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gossa_test

// This test runs the SSA interpreter over sample Go programs.
// Because the interpreter requires intrinsics for assembly
// functions and many low-level runtime routines, it is inherently
// not robust to evolutionary change in the standard library.
// Therefore the test cases are restricted to programs that
// use a fake standard library in testdata/src containing a tiny
// subset of simple functions useful for writing assertions.
//
// We no longer attempt to interpret any real standard packages such as
// fmt or testing, as it proved too fragile.

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/goplus/gossa"
	_ "github.com/goplus/gossa/pkg/bytes"
	_ "github.com/goplus/gossa/pkg/errors"
	_ "github.com/goplus/gossa/pkg/fmt"
	_ "github.com/goplus/gossa/pkg/math"
	_ "github.com/goplus/gossa/pkg/os"
	_ "github.com/goplus/gossa/pkg/reflect"
	_ "github.com/goplus/gossa/pkg/runtime"
	_ "github.com/goplus/gossa/pkg/strings"
	_ "github.com/goplus/gossa/pkg/sync"
	_ "github.com/goplus/gossa/pkg/time"
)

// These are files in github.com/goplus/gossa/testdata/.
var testdataTests = []string{
	"boundmeth.go",
	"complit.go",
	"coverage.go",
	"defer.go",
	"fieldprom.go",
	"ifaceconv.go",
	"ifaceprom.go",
	"initorder.go",
	"methprom.go",
	"mrvchain.go",
	"range.go",
	"recover.go",
	"reflect.go",
	"static.go",
	"recover2.go",
	"static.go",
	"issue23536.go",
	"tinyfin.go",
}

func runInput(t *testing.T, input string) bool {
	fmt.Println("Input:", input)
	start := time.Now()
	_, err := gossa.Run(input, nil, 0)
	sec := time.Since(start).Seconds()
	if err != nil {
		t.Error(err)
		fmt.Printf("FAIL %0.3fs\n", sec)
		return false
	}
	fmt.Printf("PASS %0.3fs\n", sec)
	return true
}

// TestTestdataFiles runs the interpreter on testdata/*.go.
func TestTestdataFiles(t *testing.T) {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	var failures []string
	for _, input := range testdataTests {
		if !runInput(t, filepath.Join(cwd, "testdata", input)) {
			failures = append(failures, input)
		}
	}
	printFailures(failures)
}

func printFailures(failures []string) {
	if failures != nil {
		fmt.Println("The following tests failed:")
		for _, f := range failures {
			fmt.Printf("\t%s\n", f)
		}
	}
}

func TestUntypedNil(t *testing.T) {
	src := `package main

type T func()

func main() {
	if T(nil) != nil {
		panic("error")
	}
}
`
	_, err := gossa.RunFile("main.go", src, nil, 0)
	if err != nil {
		t.Fatal(err)
	}
}

func TestOverrideFunction(t *testing.T) {
	ctx := gossa.NewContext(0)
	ctx.SetOverrideFunction("main.call", func(i, j int) int {
		return i * j
	})
	src := `package main

func call(i, j int) int {
	return i+j
}

func main() {
	if n := call(10,20); n != 200 {
		panic(n)
	}
}
`
	_, err := ctx.RunFile("main.go", src, nil)
	if err != nil {
		t.Fatal(err)
	}

	// reset override func
	ctx.ClearOverrideFunction("main.call")
	_, err = ctx.RunFile("main.go", src, nil)
	if err == nil || err.Error() != "30" {
		t.Fatal("must panic 30")
	}
}

func TestOsExit(t *testing.T) {
	src := `package main

import "os"

func main() {
	os.Exit(-2)
}
`
	code, err := gossa.RunFile("main.go", src, nil, 0)
	if err != nil {
		t.Fatal(err)
	}
	if code != -2 {
		t.Fatalf("exit code %v, must -2", code)
	}
}

func TestOpAlloc(t *testing.T) {
	src := `package main

type T struct {
	n1 int
	n2 int
}

func (t T) call() int {
	return t.n1 * t.n2
}

func main() {
	var n int
	for i := 0; i < 3; i++ {
		n += T{i,3}.call()
	}
	if n != 9 {
		panic(n)
	}
}
`
	_, err := gossa.RunFile("main.go", src, nil, 0)
	if err != nil {
		t.Fatal(err)
	}
}

func TestOpBin(t *testing.T) {
	src := `package main

func main() {
	a := 10 // 1010
	b := 12 // 1100
	n1 := a + b
	n2 := a - b
	n3 := a * b
	n4 := a / b
	n5 := a % b
	n6 := a & b
	n7 := a | b
	n8 := a ^ b
	n9 := a &^ b
	n10 := a << 3
	n11 := a >> 3
	v1 := a > b
	v2 := a < b
	v3 := a >= b
	v4 := a <= b
	v5 := a == b
	v6 := a != b
	if n1 != 22 || n2 != -2 || n3 != 120 || n4 != 0 || n5 != 10 || n6 != 8 ||
		n7 != 14 || n8 != 6 || n9 != 2 || n10 != 80 || n11 != 1 ||
		v1 != false || v2 != true || v3 != false || v4 != true || v5 != false || v6 != true {
		println(n1, n2, n3, n4, n5, n6, n7, n8, n9, n10, n11)
		println(v1, v2, v3, v4, v5, v6)
		panic("error")
	}
}
`
	_, err := gossa.RunFile("main.go", src, nil, 0)
	if err != nil {
		t.Fatal(err)
	}
}

func TestOpChangeType(t *testing.T) {
	src := `package main

type T func(int, int) int

func add(i, j int) int {
	return i + j
}

type I interface{}

func main() {
	// static change type
	fn := T(add)
	if n := fn(10, 20); n != 30 {
		panic(n)
	}
	var i interface{} = I(nil)
	if i != nil {
		panic("must nil")
	}
	// dynamic change type
	var k interface{}
	i = I(k)
	if i != nil {
		panic("must nil")
	}
	i = I(100)
	if i == nil {
		panic("must not nil")
	}
}
`
	_, err := gossa.RunFile("main.go", src, nil, 0)
	if err != nil {
		t.Fatal(err)
	}
}

func TestOpIndex(t *testing.T) {
	src := `package main

func main() {
	var n int
	for _, v := range [3]int{1, 2, 4} {
		n += v
	}
	if n != 7 {
		panic(n)
	}
}
`
	_, err := gossa.RunFile("main.go", src, nil, 0)
	if err != nil {
		t.Fatal(err)
	}
}

func TestUnderscopeMap(t *testing.T) {
	src := `package main

type key struct {
	_ int
	x int
}

var (
	s1 = "s1"
	s2 = "s2"
	s3 = "s3"
	s4 = "s4"
)

func main() {
	m := make(map[key]string)
	m[key{0, 1}] = s1
	m[key{0, 2}] = s2
	m[key{0, 3}] = s3
	m[key{1, 1}] = s4
	m[key{1, 2}] = "s22"
	if n := len(m); n != 3 {
		panic(n)
	}
	if v := m[key{2, 2}]; v != "s22" {
		panic(v)
	}
	if v, ok := m[key{1,4}]; ok {
		panic(v)
	}
}
`
	_, err := gossa.RunFile("main.go", src, nil, 0)
	if err != nil {
		t.Fatal(err)
	}
}

func TestOpSend(t *testing.T) {
	src := `package main

import (
	"fmt"
)

type T struct{}

func main() {
	ch := make(chan *T)
	go func() {
		ch <- nil
	}()
	v := <-ch
	if v != nil {
		panic("must nil")
	}
	if s := fmt.Sprintf("%T", v); s != "*main.T" {
		panic(s)
	}
	go func() {
		ch <- &T{}
	}()
	v = <-ch
	if v == nil {
		panic("must not nil")
	}
}
`
	_, err := gossa.RunFile("main.go", src, nil, 0)
	if err != nil {
		t.Fatal(err)
	}
}

func TestShadowedMethod(t *testing.T) {
	src := `// run

// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// When computing method sets with shadowed methods, make sure we
// compute whether a method promotion involved a pointer traversal
// based on the promoted method, not the shadowed method.

package main

import (
	"bytes"
	"fmt"
)

type mystruct struct {
	f int
}

func (t mystruct) String() string {
	return "FAIL"
}

func main() {
	type deep struct {
		mystruct
	}
	s := struct {
		deep
		*bytes.Buffer
	}{
		deep{},
		bytes.NewBufferString("ok"),
	}

	if got := s.String(); got != "ok" {
		panic(got)
	}

	var i fmt.Stringer = s
	if got := i.String(); got != "ok" {
		panic(got)
	}
}
`
	_, err := gossa.RunFile("main.go", src, nil, 0)
	if err != nil {
		t.Fatal(err)
	}
}

func TestConvertUnsafePointer(t *testing.T) {
	src := `package main

import (
	"unsafe"
)

func main() {
	a := [4]int{0, 1, 2, 3}
	p := unsafe.Pointer(&a)
	p2 := unsafe.Pointer(uintptr(p) + 2*unsafe.Sizeof(a[0]))
	*(*int)(p2) = 4
	if a != [4]int{0, 1, 4, 3} {
		panic("error")
	}
}
`
	_, err := gossa.RunFile("main.go", src, nil, 0)
	if err != nil {
		t.Fatal(err)
	}
}

func TestBuiltinPrintln(t *testing.T) {
	src := `// run

// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Test internal print routines that are generated
// by the print builtin.  This test is not exhaustive,
// we're just checking that the formatting is correct.

package main

func main() {
	println((interface{})(nil)) // printeface
	println((interface {        // printiface
		f()
	})(nil))
	println((map[int]int)(nil)) // printpointer
	println(([]int)(nil))       // printslice
	println(int64(-7))          // printint
	println(uint64(7))          // printuint
	println(uint32(7))          // printuint
	println(uint16(7))          // printuint
	println(uint8(7))           // printuint
	println(uint(7))            // printuint
	println(uintptr(7))         // printuint
	println(8.0)                // printfloat
	println(complex(9.0, 10.0)) // printcomplex
	println(true)               // printbool
	println(false)              // printbool
	println("hello")            // printstring
	println("one", "two")       // printsp

	// test goprintf
	defer println((interface{})(nil))
	defer println((interface {
		f()
	})(nil))
	defer println((map[int]int)(nil))
	defer println(([]int)(nil))
	defer println(int64(-11))
	defer println(uint64(12))
	defer println(uint32(12))
	defer println(uint16(12))
	defer println(uint8(12))
	defer println(uint(12))
	defer println(uintptr(12))
	defer println(13.0)
	defer println(complex(14.0, 15.0))
	defer println(true)
	defer println(false)
	defer println("hello")
	defer println("one", "two")
}
`
	out := `(0x0,0x0)
(0x0,0x0)
0x0
[0/0]0x0
-7
7
7
7
7
7
7
+8.000000e+000
(+9.000000e+000+1.000000e+001i)
true
false
hello
one two
one two
hello
false
true
(+1.400000e+001+1.500000e+001i)
+1.300000e+001
12
12
12
12
12
12
-11
[0/0]0x0
0x0
(0x0,0x0)
(0x0,0x0)
`
	ctx := gossa.NewContext(0)
	var buf bytes.Buffer
	ctx.SetPrintOutput(&buf)
	_, err := ctx.RunFile("main.go", src, nil)
	if err != nil {
		t.Fatal(err)
	}
	if buf.String() != out {
		t.Fatal("error")
	}
}

type panicinfo struct {
	src string
	err string
}

func TestPanicInfo(t *testing.T) {
	infos := []panicinfo{
		{`panic(100)`, `100`},
		{`panic(100.0)`, `+1.000000e+002`},
		{`panic("hello")`, `hello`},
		{`panic((*interface{})(nil))`, `(*interface {}) 0x0`},
		{`type T int; panic(T(100))`, `main.T(100)`},
		{`type T float64; panic(T(100.0))`, `main.T(+1.000000e+002)`},
		{`type T string; panic(T("hello"))`, `main.T("hello")`},
		{`type T struct{}; panic((*T)(nil))`, `(*main.T) 0x0`},
	}
	ctx := gossa.NewContext(0)
	for _, info := range infos {
		src := fmt.Sprintf("package main;func main(){%v}", info.src)
		code, err := ctx.RunFile("main.go", src, nil)
		if code != 2 || err == nil {
			t.Fatalf("%v must panic", info.src)
		}
		if s := err.Error(); s != info.err {
			t.Fatalf("%v err is %v, want %v", info.src, s, info.err)
		}
	}
}

func TestPanicError(t *testing.T) {
	src := `package main
import "errors"
func main() {
	panic(errors.New("error info"))
}
`
	_, err := gossa.RunFile("main.go", src, nil, 0)
	if err == nil {
		t.Fatal("must panic")
	}
	if s := err.Error(); s != "error info" {
		t.Fatalf("error %q", s)
	}
}

func TestPanicErrorRecover(t *testing.T) {
	src := `package main
import "errors"
func main() {
	defer func() {
		err := recover()
		if err == nil {
			panic("must error")
		}
		if s := err.(error).Error(); s != "error info" {
			panic(s)
		}
	}()
	panic(errors.New("error info"))
}
`
	_, err := gossa.RunFile("main.go", src, nil, 0)
	if err != nil {
		t.Fatal(err)
	}
}

func TestEnablePrintAny(t *testing.T) {
	src := `package main

type T struct {
	X int
	Y int
}
func main() {
	println([2]int{100,200})
	println(T{100,200})
}
`
	ctx := gossa.NewContext(0)
	var buf bytes.Buffer
	ctx.SetPrintOutput(&buf)
	_, err := ctx.RunFile("main.go", src, nil)
	if err == nil {
		t.Fatal("must panic")
	}
	if s := err.Error(); s != "illegal types for operand: print\n\t[2]int" {
		t.Fatal(s)
	}
	ctx.Mode |= gossa.EnablePrintAny
	_, err = ctx.RunFile("main.go", src, nil)
	if err != nil {
		t.Fatal(err)
	}
	if s := buf.String(); s != "[100 200]\n{100 200}\n" {
		t.Fatal(s)
	}
}

func TestFib(t *testing.T) {
	src := `package main

import (
	"sync"
)

func fib(n int) int {
	if n < 2 {
		return n
	}
	return fib(n-2) + fib(n-1)
}

func main() {
	var wg sync.WaitGroup
	const N = 10
	for i := 0; i < N; i++ {
		wg.Add(1)
		go func() {
			if fib(20) != 6765 {
				panic("error")
			}
			wg.Done()
		}()
	}
	wg.Wait()
}
`
	_, err := gossa.RunFile("main.go", src, nil, 0)
	if err != nil {
		t.Fatal(err)
	}
}

func TestFibv(t *testing.T) {
	src := `package main

import "sync"

func main() {
	var fib func(n int) int
	fib = func(n int) int {
		if n < 2 {
			return n
		}
		return fib(n-2) + fib(n-1)
	}

	var wg sync.WaitGroup
	const N = 10
	for i := 0; i < N; i++ {
		wg.Add(1)
		go func() {
			if fib(20) != 6765 {
				panic("error")
			}
			wg.Done()
		}()
	}
	wg.Wait()
}
`
	_, err := gossa.RunFile("main.go", src, nil, 0)
	if err != nil {
		t.Fatal(err)
	}
}

func TestFibi(t *testing.T) {
	src := `package main

import "sync"

type I interface {
	fib(i I, n int) int
}

type T struct {
}

func (t *T) fib(i I, n int) int {
	if n < 2 {
		return n
	}
	return i.fib(i, n-2) + i.fib(i, n-1)
}

func fib(n int) int {
	t := &T{}
	return t.fib(t, n)
}

func main() {
	var wg sync.WaitGroup
	const N = 10
	for i := 0; i < N; i++ {
		wg.Add(1)
		go func() {
			if fib(20) != 6765 {
				panic("error")
			}
			wg.Done()
		}()
	}
	wg.Wait()
}
`
	_, err := gossa.RunFile("main.go", src, nil, 0)
	if err != nil {
		t.Fatal(err)
	}
}

func TestBinOpInt(t *testing.T) {
	// + - * / % & ^ &^ < <= > >=
	tsrc := `package main

import "fmt"

type T $int

func main() {
	// 0101 0110
	test(5, 6)
	testConst()
	testDivideZero()
}

func testConst() {
	var a T = 4
	var b T = 2
	check(a+b,6)
	check(a-b,2)
	check(a*b,8)
	check(a/b,2)
	check(a%b,0)
	check(a&b,0)
	check(a|b,6)
	check(a^b,6)
	check(a&^b,4)
	assert(a>b)
	assert(a>=b)
	assert(b<a)
	assert(a<=a)
}

func testDivideZero() {
	defer func() {
		if r := recover(); r == nil {
			panic("must panic")
		}
	}()
	var a T
	var b T
	_ = a/b
}

func test(a, b T) {
	check(a+1, 6)
	check(1+a, 6)
	check(a+b, 11)
	check(a-1, 4)
	check(6-a, 1)
	check(b-a, 1)
	check(a*2, 10)
	check(2*a, 10)
	check(a*b, 30)
	check(a/2, 2)
	check(T(10)/a, 2)
	check(a/b, 0)
	check(a%2, 1)
	check(10%a, 0)
	check(b%a, 1)
	check(a&4, 4)
	check(1&a, 1)
	check(a&b, 4)
	check(a|4, 5)
	check(2|a, 7)
	check(a|b, 7)
	check(a^2, 7)
	check(7^a, 2)
	check(a^b, 3)
	check(a&^3, 4)
	check(3&^a, 2)
	check(a&^b, 1)
	assert(a < 6)
	assert(4 < a)
	assert(a < b)
	assert(a <= 5)
	assert(4 <= a)
	assert(a <= b)
	assert(a > 4)
	assert(6 > a)
	assert(b > a)
	assert(a >= 5)
	assert(6 >= a)
	assert(b >= a)
}

func assert(t bool) {
	if !t {
		panic("error")
	}
}

func check(a, b T) {
	if a != b {
		panic(fmt.Errorf("error %v != %v", a, b))
	}
}
`
	types := []string{
		"int", "int8", "int16", "int32", "int64",
		"uint", "uint8", "uint16", "uint32", "uint64", "uintptr",
	}
	for _, s := range types {
		t.Log("test binop basic", s)
		src := strings.Replace(tsrc, "$int", "="+s, 1)
		_, err := gossa.RunFile("main.go", src, nil, 0)
		if err != nil {
			t.Fatal(err)
		}
	}
	for _, s := range types {
		t.Log("test binop named", s)
		src := strings.Replace(tsrc, "$int", s, 1)
		_, err := gossa.RunFile("main.go", src, nil, 0)
		if err != nil {
			t.Fatal(err)
		}
	}
}

func TestBinOpFloat(t *testing.T) {
	tsrc := `package main

import "fmt"

type T $float

func main() {
	test(5.0, 6.5)
	testConst()
}

func testConst() {
	var a T = 4.0
	var b T = 2.0
	check(a+b,6)
	check(a-b,2)
	check(a*b,8)
	check(a/b,2)
	assert(a>b)
	assert(a>=b)
	assert(b<a)
	assert(a<=a)
}

func test(a, b T) {
	check(a+1, 6)
	check(1+a, 6)
	check(a+b, 11.5)
	check(a-1, 4)
	check(6-a, 1)
	check(b-a, 1.5)
	check(a*2, 10)
	check(2*a, 10)
	check(a*b, 32.5)
	check(a/2, 2.5)
	check(10/a, 2)
	check(a/b, 5.0/6.5)
	assert(a < 6)
	assert(4 < a)
	assert(a < b)
	assert(a <= 5)
	assert(4 <= a)
	assert(a <= b)
	assert(a > 4)
	assert(6 > a)
	assert(b > a)
	assert(a >= 5)
	assert(6 >= a)
	assert(b >= a)
}

func assert(t bool) {
	if !t {
		panic("error")
	}
}

func check(a, b T) {
	if a != b {
		panic(fmt.Errorf("error %v != %v", a, b))
	}
}
`
	types := []string{
		"float32",
		"float64",
	}
	for _, s := range types {
		t.Log("test binop basic", s)
		src := strings.Replace(tsrc, "$float", "="+s, 1)
		_, err := gossa.RunFile("main.go", src, nil, 0)
		if err != nil {
			t.Fatal(err)
		}
	}
	for _, s := range types {
		t.Log("test binop named", s)
		src := strings.Replace(tsrc, "$float", s, 1)
		_, err := gossa.RunFile("main.go", src, nil, 0)
		if err != nil {
			t.Fatal(err)
		}
	}
}

func TestBinOpComplex(t *testing.T) {
	tsrc := `package main

import "fmt"

type T $complex

func main() {
	test(1+2i, 3+4i)
	testConst()
}

func testConst() {
	var a T = 4i
	var b T = 2i
	check(a+b,6i)
	check(a-b,2i)
	check(a*b,-8)
	check(a/b,2)
}

func test(a, b T) {
	check(a+1+2i, 2+4i)
	check(1+a, 2+2i)
	check(a+b, 4+6i)
	check(a-1, 2i)
	check(6-a, 5-2i)
	check(b-a, 2+2i)
	check(a*(2+3i), -4+7i)
	check((2-3i)*a, 8+1i)
	check(a*b, -5+10i)
	check(a/2i, 1-0.5i)
	check(10/a, 2-4i)
	check(a/b, 0.44+0.08i)
}

func assert(t bool) {
	if !t {
		panic("error")
	}
}

func check(a, b T) {
	if a != b {
		panic(fmt.Errorf("error %v != %v", a, b))
	}
}

`
	types := []string{
		"complex64",
		"complex128",
	}
	for _, s := range types {
		t.Log("test binop basic", s)
		src := strings.Replace(tsrc, "$complex", "="+s, 1)
		_, err := gossa.RunFile("main.go", src, nil, 0)
		if err != nil {
			t.Fatal(err)
		}
	}
	for _, s := range types {
		t.Log("test binop named", s)
		src := strings.Replace(tsrc, "$complex", s, 1)
		_, err := gossa.RunFile("main.go", src, nil, 0)
		if err != nil {
			t.Fatal(err)
		}
	}
}

func TestBinOpString(t *testing.T) {
	tsrc := `package main

import "fmt"

type T = string

func main() {
	test("go", "ssa")
	testConst()
}

func testConst() {
	var a T = "hello"
	var b T = "world"
	check(a+b,"helloworld")
	assert(a < b)
	assert(a <= b)
	assert(b > a)
	assert(b >= a)
}

func test(a, b T) {
	check(a+"run", "gorun")
	check("run"+a, "rungo")
	check(a+b, "gossa")
	assert(a < "go1")
	assert("ao" < a)
	assert(a < b)
	assert(a <= "go")
	assert("ao" <= a)
	assert(a <= b)
	assert(a > "ao")
	assert("go1" > a)
	assert(b > a)
	assert(a >= "go")
	assert("go1" >= a)
	assert(b >= a)
}

func assert(t bool) {
	if !t {
		panic("error")
	}
}

func check(a, b T) {
	if a != b {
		panic(fmt.Errorf("error %v != %v", a, b))
	}
}
`
	t.Log("test binop string")
	_, err := gossa.RunFile("main.go", tsrc, nil, 0)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("test binop named string")
	src := strings.Replace(tsrc, "= string", "string", 1)
	_, err = gossa.RunFile("main.go", src, nil, 0)
	if err != nil {
		t.Fatal(err)
	}
}

func TestBinOpShift(t *testing.T) {
	tsrc := `package main
type T1 $T1
type T2 $T2

func main() {
	test(4, 2)
}
func test(a T1, b T2) {
	assert(a << b == 16)
	assert(a >> b == 1)
	assert(a << T2(2) == 16)
	assert(a >> T2(2) == 1)
	assert(T1(4) << b == 16)
	assert(T1(4) >> b == 1)
	var c T1 = 4
	var d T2 = 2
	assert(c << d == 16)
	assert(c >> d == 1)
}
func assert(t bool) {
	if !t {
		panic("error")
	}
}
`
	types := []string{
		"int", "int8", "int16", "int32", "int64",
		"uint", "uint8", "uint16", "uint32", "uint64", "uintptr",
	}
	for _, t1 := range types {
		t.Log("test binop shift", t1)
		for _, t2 := range types {
			r := strings.NewReplacer("$T1", "="+t1, "$T2", "="+t2)
			src := r.Replace(tsrc)
			_, err := gossa.RunFile("main.go", src, nil, 0)
			if err != nil {
				t.Fatal(err)
			}
		}
		for _, t2 := range types {
			r := strings.NewReplacer("$T1", "="+t1, "$T2", t2)
			src := r.Replace(tsrc)
			_, err := gossa.RunFile("main.go", src, nil, 0)
			if err != nil {
				t.Fatal(err)
			}
		}
		for _, t2 := range types {
			r := strings.NewReplacer("$T1", t1, "$T2", "="+t2)
			src := r.Replace(tsrc)
			_, err := gossa.RunFile("main.go", src, nil, 0)
			if err != nil {
				t.Fatal(err)
			}
		}
		for _, t2 := range types {
			r := strings.NewReplacer("$T1", t1, "$T2", t2)
			src := r.Replace(tsrc)
			_, err := gossa.RunFile("main.go", src, nil, 0)
			if err != nil {
				t.Fatal(err)
			}
		}
	}
}

func TestUnOpNot(t *testing.T) {
	src := `package main
type T bool
func main() {
	test(false,false)
	testConst()
}
func test(b1 bool, b2 T) {
	if v := !b1; v != true {
		panic("must true")
	}
	if v := !b2; v != true {
		panic("must true")
	}
}
func testConst() {
	var b1 bool
	var b2 T
	if v := !b1; v != true {
		panic("must true")
	}
	if v := !b2; v != true {
		panic("must true")
	}
}
`
	_, err := gossa.RunFile("main.go", src, nil, 0)
	if err != nil {
		t.Fatal(err)
	}
}

func TestUnOpSubInt(t *testing.T) {
	tsrc := `package main
type T $int

func main() {
	test(10,10)
	testConst()
}

func test(a $int, b T) {
	if r := -a; r != -10 {
		panic("must -10")
	}
	if r := -b; r != -10 {
		panic("must -10")
	}
}
func testConst() {
	var a $int = 10
	var b T = 10
	if r := -a; r != -10 {
		panic("must -10")
	}
	if r := -b; r != -10 {
		panic("must -10")
	}
}
`
	types := []string{
		"int", "int8", "int16", "int32", "int64",
	}

	for _, s := range types {
		t.Log("test unop sub", s)
		src := strings.Replace(tsrc, "$int", s, -1)
		_, err := gossa.RunFile("main.go", src, nil, 0)
		if err != nil {
			t.Fatal(err)
		}
	}
}

func TestUnOpSubUint(t *testing.T) {
	tsrc := `package main

import (
	"math"
)

var (
	_ = math.Pi
)

type T $uint

func main() {
	test(10, 10)
	testConst()
}

const (
	r1 $uint = math.MaxUint - 9
	r2 T    = T(math.MaxUint - 9)
)

func test(a $uint, b T) {
	if r := -a; r != r1 {
		panic(r)
	}
	if r := -b; r != r2 {
		panic(r)
	}
}
func testConst() {
	var a $uint = 10
	var b T = 10
	if r := -a; r != r1 {
		panic(r)
	}
	if r := -b; r != r2 {
		panic(r)
	}
}
`
	types := []string{
		"uint", "uint8", "uint16", "uint32", "uint64", "uintptr",
	}
	intSize := 32 << (^uint(0) >> 63)
	for _, s := range types {
		t.Log("test unop sub", s)
		var r *strings.Replacer
		if s == "uint" || s == "uintptr" {
			r = strings.NewReplacer("$uint", s, "math.MaxUint", fmt.Sprintf("%v(%v)", s, uint(1<<intSize-1)))
		} else {
			r = strings.NewReplacer("$uint", s, "Uint", strings.Title(s))
		}
		src := r.Replace(tsrc)
		_, err := gossa.RunFile("main.go", src, nil, 0)
		if err != nil {
			t.Fatal(err)
		}
	}
}

func TestUnOpSubFloat(t *testing.T) {
	tsrc := `package main

type T $float

func main() {
	test(5.0, 6.5)
	testConst()
}

func test(a $float, b T) {
	if r := -a; r != -5.0 {
		panic(r)
	}
	if r := -b; r != -6.5 {
		panic(r)
	}
}
func testConst() {
	var a $float = 5.0
	var b T = 6.5
	if r := -a; r != -5.0 {
		panic(r)
	}
	if r := -b; r != -6.5 {
		panic(r)
	}
}
`
	types := []string{
		"float32", "float64",
	}
	for _, s := range types {
		t.Log("test unop sub", s)
		src := strings.Replace(tsrc, "$float", s, -1)
		_, err := gossa.RunFile("main.go", src, nil, 0)
		if err != nil {
			t.Fatal(err)
		}
	}
}

func TestUnOpSubComplex(t *testing.T) {
	tsrc := `package main

type T $complex

func main() {
	test(1+2i, 3+4i)
	testConst()
}

func test(a $complex, b T) {
	if r := -a; r != -1-2i {
		panic(r)
	}
	if r := -b; r != -3-4i {
		panic(r)
	}
}
func testConst() {
	var a $complex = 1+2i
	var b T = 3+4i
	if r := -a; r != -1-2i {
		panic(r)
	}
	if r := -b; r != -3-4i {
		panic(r)
	}
}
`
	types := []string{
		"complex64", "complex128",
	}
	for _, s := range types {
		t.Log("test unop sub", s)
		src := strings.Replace(tsrc, "$complex", s, -1)
		_, err := gossa.RunFile("main.go", src, nil, 0)
		if err != nil {
			t.Fatal(err)
		}
	}
}

func TestUnOpXorInt(t *testing.T) {
	tsrc := `package main

type T $int

func main() {
	test(10, 11)
	testConst()
}

func test(a $int, b T) {
	if r := ^a; r != -11 {
		panic(r)
	}
	if r := ^b; r != -12 {
		panic(r)
	}
}
func testConst() {
	var a $int = 10
	var b T = 11
	if r := ^a; r != -11 {
		panic(r)
	}
	if r := ^b; r != -12 {
		panic(r)
	}
}
`
	types := []string{
		"int", "int8", "int16", "int32", "int64",
	}

	for _, s := range types {
		t.Log("test unop xor", s)
		src := strings.Replace(tsrc, "$int", s, -1)
		_, err := gossa.RunFile("main.go", src, nil, 0)
		if err != nil {
			t.Fatal(err)
		}
	}
}

func TestUnOpXorUint(t *testing.T) {
	tsrc := `package main

type T $uint

func main() {
	test(0b1010, 0b1011)
	testConst()
}

func test(a $uint, b T) {
	if r := ^a; r&0xff != 0b11110101 {
		panic(r)
	}
	if r := ^b; r&0xff != 0b11110100 {
		panic(r)
	}
}
func testConst() {
	var a $uint = 0b1010
	var b T = 0b1011
	if r := ^a; r&0xff != 0b11110101 {
		panic(r)
	}
	if r := ^b; r&0xff != 0b11110100 {
		panic(r)
	}
}
`
	types := []string{
		"uint", "uint8", "uint16", "uint32", "uint64", "uintptr",
	}

	for _, s := range types {
		t.Log("test unop xor", s)
		src := strings.Replace(tsrc, "$uint", s, -1)
		_, err := gossa.RunFile("main.go", src, nil, 0)
		if err != nil {
			t.Fatal(err)
		}
	}
}
