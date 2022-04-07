//go:build go1.17
// +build go1.17

package gossa_test

import (
	"testing"

	"github.com/goplus/gossa"
)

func TestUnsafeBuiltins(t *testing.T) {
	src := `// run

// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"math"
	"unsafe"
)

const maxUintptr = 1 << (8 * unsafe.Sizeof(uintptr(0)))

func main() {
	var p [10]byte

	// unsafe.Add
	{
		p1 := unsafe.Pointer(&p[1])
		assert(unsafe.Add(p1, 1) == unsafe.Pointer(&p[2]))
		assert(unsafe.Add(p1, -1) == unsafe.Pointer(&p[0]))
	}

	// unsafe.Slice
	{
		s := unsafe.Slice(&p[0], len(p))
		assert(&s[0] == &p[0])
		assert(len(s) == len(p))
		assert(cap(s) == len(p))

		// nil pointer with zero length returns nil
		assert(unsafe.Slice((*int)(nil), 0) == nil)

		// nil pointer with positive length panics
		mustPanic(func() { _ = unsafe.Slice((*int)(nil), 1) })

		// negative length
		var neg int = -1
		mustPanic(func() { _ = unsafe.Slice(new(byte), neg) })

		// length too large
		var tooBig uint64 = math.MaxUint64
		mustPanic(func() { _ = unsafe.Slice(new(byte), tooBig) })

		// size overflows address space
		mustPanic(func() { _ = unsafe.Slice(new(uint64), maxUintptr/8) })
		mustPanic(func() { _ = unsafe.Slice(new(uint64), maxUintptr/8+1) })
	}
}

func assert(ok bool) {
	if !ok {
		panic("FAIL")
	}
}

func mustPanic(f func()) {
	defer func() {
		assert(recover() != nil)
	}()
	f()
}
`
	_, err := gossa.RunFile("main.go", src, nil, 0)
	if err != nil {
		t.Fatal(err)
	}
}

func TestSliceToArrayPointer(t *testing.T) {
	src := `// run

// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Test conversion from slice to array pointer.

package main

func wantPanic(fn func(), s string) {
	defer func() {
		err := recover()
		if err == nil {
			panic("expected panic")
		}
		if got := err.(error).Error(); got != s {
			panic("expected panic " + s + " got " + got)
		}
	}()
	fn()
}

func main() {
	s := make([]byte, 8, 10)
	if p := (*[8]byte)(s); &p[0] != &s[0] {
		panic("*[8]byte conversion failed")
	}
	wantPanic(
		func() {
			_ = (*[9]byte)(s)
		},
		"runtime error: cannot convert slice with length 8 to pointer to array with length 9",
	)

	var n []byte
	if p := (*[0]byte)(n); p != nil {
		panic("nil slice converted to *[0]byte should be nil")
	}

	z := make([]byte, 0)
	if p := (*[0]byte)(z); p == nil {
		panic("empty slice converted to *[0]byte should be non-nil")
	}

	// Test with named types
	type Slice []int
	type Int4 [4]int
	type PInt4 *[4]int
	ii := make(Slice, 4)
	if p := (*Int4)(ii); &p[0] != &ii[0] {
		panic("*Int4 conversion failed")
	}
	if p := PInt4(ii); &p[0] != &ii[0] {
		panic("PInt4 conversion failed")
	}
}

// test static variable conversion

var (
	ss  = make([]string, 10)
	s5  = (*[5]string)(ss)
	s10 = (*[10]string)(ss)

	ns  []string
	ns0 = (*[0]string)(ns)

	zs  = make([]string, 0)
	zs0 = (*[0]string)(zs)
)

func init() {
	if &ss[0] != &s5[0] {
		panic("s5 conversion failed")
	}
	if &ss[0] != &s10[0] {
		panic("s5 conversion failed")
	}
	if ns0 != nil {
		panic("ns0 should be nil")
	}
	if zs0 == nil {
		panic("zs0 should not be nil")
	}
}
`
	_, err := gossa.RunFile("main.go", src, nil, 0)
	if err != nil {
		t.Fatal(err)
	}
}
