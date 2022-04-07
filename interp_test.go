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
	ctx.SetOverrideFunction("main.call", nil)
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
