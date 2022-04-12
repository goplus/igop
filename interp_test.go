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
	"testing"
	"time"

	"github.com/goplus/gossa"
	_ "github.com/goplus/gossa/pkg/bytes"
	_ "github.com/goplus/gossa/pkg/context"
	_ "github.com/goplus/gossa/pkg/crypto/md5"
	_ "github.com/goplus/gossa/pkg/encoding/binary"
	_ "github.com/goplus/gossa/pkg/errors"
	_ "github.com/goplus/gossa/pkg/flag"
	_ "github.com/goplus/gossa/pkg/fmt"
	_ "github.com/goplus/gossa/pkg/go/ast"
	_ "github.com/goplus/gossa/pkg/io"
	_ "github.com/goplus/gossa/pkg/io/ioutil"
	_ "github.com/goplus/gossa/pkg/log"
	_ "github.com/goplus/gossa/pkg/math"
	_ "github.com/goplus/gossa/pkg/math/rand"
	_ "github.com/goplus/gossa/pkg/os"
	_ "github.com/goplus/gossa/pkg/reflect"
	_ "github.com/goplus/gossa/pkg/runtime"
	_ "github.com/goplus/gossa/pkg/runtime/debug"
	_ "github.com/goplus/gossa/pkg/strconv"
	_ "github.com/goplus/gossa/pkg/strings"
	_ "github.com/goplus/gossa/pkg/sync"
	_ "github.com/goplus/gossa/pkg/sync/atomic"
	_ "github.com/goplus/gossa/pkg/syscall"
	_ "github.com/goplus/gossa/pkg/testing"
	_ "github.com/goplus/gossa/pkg/time"
)

// These are files in go.tools/go/ssa/interp/testdata/.
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
