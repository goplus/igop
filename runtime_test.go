package igop_test

import (
	"testing"

	"github.com/goplus/igop"
	_ "github.com/goplus/igop/pkg/log"
)

func TestFuncForPC(t *testing.T) {
	src := `package main

import (
	"fmt"
	"reflect"
	"runtime"
)

func test() {
	println("hello")
}

func main() {
	pc := reflect.ValueOf(test).Pointer()
	fn := runtime.FuncForPC(pc)
	if fn == nil {
		panic("error runtime.FuncForPC")
	}
	if fn.Name() != "main.test" {
		panic("error name: " + fn.Name())
	}
	file, line := fn.FileLine(fn.Entry())
	if file != "main.go" {
		panic("error file:" + file)
	}
	if line != 9 {
		panic(fmt.Errorf("error line: %v", line))
	}
}
`
	_, err := igop.RunFile("main.go", src, nil, 0)
	if err != nil {
		t.Fatal(err)
	}
}

func TestRuntimeCaller(t *testing.T) {
	src := `package main

import (
	"fmt"
	"path/filepath"
	"runtime"
	"strings"
)

var t *struct {
	c chan int
}

var c chan int

func f() {
	select {
	case <-t.c: // THIS IS LINE 18
		break
	case <-c:
		break
	}
}

func g() {
	defer func() {
		panic("error g1")
	}()
	defer func() {
		panic("error g2")
	}()
	f()
	panic("unreachable")
}

func main() {
	defer func() {
		recover()
		var list []string
		for i := 0; ; i++ {
			pc, file, line, ok := runtime.Caller(i)
			if !ok {
				break
			}
			fn := runtime.FuncForPC(pc)
			fnName := fn.Name()
			if fnName == "runtime.gopanic" {
				list = append(list, "runtime.gopanic")
			} else if strings.HasPrefix(fnName, "main.") {
				_, fname := filepath.Split(file)
				list = append(list, fmt.Sprintf("%v %v:%v", fnName, fname, line))
			}
			fmt.Println(pc, fnName, file, line)
		}
		if v := fmt.Sprint(list); v != infos {
			panic(v)
		}
	}()
	g()
}

var infos = "[main.main.func1 main.go:41 runtime.gopanic main.g.func1 main.go:27 runtime.gopanic main.g.func2 main.go:30 runtime.gopanic main.f main.go:18 main.g main.go:32 main.main main.go:59]"
`
	_, err := igop.RunFile("main.go", src, nil, 0)
	if err != nil {
		t.Fatal(err)
	}
}

func TestRuntimeCallers1(t *testing.T) {
	src := `package main

import (
	"fmt"
	"path/filepath"
	"runtime"
	"strings"
)

var t *struct {
	c chan int
}

var c chan int

func f() {
	select {
	case <-t.c: // THIS IS LINE 18
		break
	case <-c:
		break
	}
}

func g() {
	// defer func() {
	// 	panic("error g1")
	// }()
	// defer func() {
	// 	panic("error g2")
	// }()
	f()
	panic("unreachable")
}

func main() {
	defer func() {
		recover()
		rpc := make([]uintptr, 64)
		runtime.Callers(0, rpc)
		fs := runtime.CallersFrames(rpc)
		var list []string
		for {
			f, more := fs.Next()
			if f.Function == "runtime.gopanic" {
				list = append(list, "runtime.gopanic")
			} else if strings.HasPrefix(f.Function, "main.") {
				_, fname := filepath.Split(f.File)
				list = append(list, fmt.Sprintf("%v %v:%v", f.Function, fname, f.Line))
			}
			fmt.Println(f)
			if !more {
				break
			}
		}
		if v := fmt.Sprint(list); v != infos {
			panic(v)
		}
	}()
	g()
}

var infos = "[main.main.func1 main.go:40 runtime.gopanic main.f main.go:18 main.g main.go:32 main.main main.go:60]"
`
	_, err := igop.RunFile("main.go", src, nil, 0)
	if err != nil {
		t.Fatal(err)
	}
}

func TestRuntimeCallers2(t *testing.T) {
	src := `package main

import (
	"fmt"
	"path/filepath"
	"runtime"
	"strings"
)

var t *struct {
	c chan int
}

var c chan int

func f() {
	select {
	case <-t.c: // THIS IS LINE 18
		break
	case <-c:
		break
	}
}

func g() {
	defer func() {
		panic("error g1")
	}()
	// defer func() {
	// 	panic("error g2")
	// }()
	f()
	panic("unreachable")
}

func main() {
	defer func() {
		recover()
		rpc := make([]uintptr, 64)
		runtime.Callers(0, rpc)
		fs := runtime.CallersFrames(rpc)
		var list []string
		for {
			f, more := fs.Next()
			if f.Function == "runtime.gopanic" {
				list = append(list, "runtime.gopanic")
			} else if strings.HasPrefix(f.Function, "main.") {
				_, fname := filepath.Split(f.File)
				list = append(list, fmt.Sprintf("%v %v:%v", f.Function, fname, f.Line))
			}
			fmt.Println(f)
			if !more {
				break
			}
		}
		if v := fmt.Sprint(list); v != infos {
			panic(v)
		}
	}()
	g()
}

var infos = "[main.main.func1 main.go:40 runtime.gopanic main.g.func1 main.go:27 runtime.gopanic main.f main.go:18 main.g main.go:32 main.main main.go:60]"
`
	_, err := igop.RunFile("main.go", src, nil, 0)
	if err != nil {
		t.Fatal(err)
	}
}

func TestRuntimeCallers3(t *testing.T) {
	src := `package main

import (
	"fmt"
	"path/filepath"
	"runtime"
	"strings"
)

var t *struct {
	c chan int
}

var c chan int

func f() {
	select {
	case <-t.c: // THIS IS LINE 18
		break
	case <-c:
		break
	}
}

func g() {
	defer func() {
		panic("error g1")
	}()
	defer func() {
		panic("error g2")
	}()
	f()
	panic("unreachable")
}

func main() {
	defer func() {
		recover()
		rpc := make([]uintptr, 64)
		runtime.Callers(0, rpc)
		fs := runtime.CallersFrames(rpc)
		var list []string
		for {
			f, more := fs.Next()
			if f.Function == "runtime.gopanic" {
				list = append(list, "runtime.gopanic")
			} else if strings.HasPrefix(f.Function, "main.") {
				_, fname := filepath.Split(f.File)
				list = append(list, fmt.Sprintf("%v %v:%v", f.Function, fname, f.Line))
			}
			fmt.Println(f)
			if !more {
				break
			}
		}
		if v := fmt.Sprint(list); v != infos {
			panic(v)
		}
	}()
	g()
}

var infos = "[main.main.func1 main.go:40 runtime.gopanic main.g.func1 main.go:27 runtime.gopanic main.g.func2 main.go:30 runtime.gopanic main.f main.go:18 main.g main.go:32 main.main main.go:60]"
`
	_, err := igop.RunFile("main.go", src, nil, 0)
	if err != nil {
		t.Fatal(err)
	}
}

func TestRuntimeCallers4(t *testing.T) {
	src := `package main

import (
	"fmt"
	"path/filepath"
	"runtime"
	"strings"
)

var t *struct {
	c chan int
}

var c chan int

func f() {
	select {
	case <-t.c: // THIS IS LINE 18
		break
	case <-c:
		break
	}
}

func g() {
	defer func() {
		// panic("error g1")
	}()
	defer func() {
		// panic("error g2")
		f()
	}()
	// f()
}

func main() {
	defer func() {
		recover()
		rpc := make([]uintptr, 64)
		runtime.Callers(0, rpc)
		fs := runtime.CallersFrames(rpc)
		var list []string
		for {
			f, more := fs.Next()
			if f.Function == "runtime.gopanic" {
				list = append(list, "runtime.gopanic")
			} else if strings.HasPrefix(f.Function, "main.") {
				_, fname := filepath.Split(f.File)
				list = append(list, fmt.Sprintf("%v %v:%v", f.Function, fname, f.Line))
			}
			fmt.Println(f)
			if !more {
				break
			}
		}
		if v := fmt.Sprint(list); v != infos {
			panic(v)
		}
	}()
	g()
}

var infos = "[main.main.func1 main.go:40 runtime.gopanic main.f main.go:18 main.g.func2 main.go:31 main.g main.go:34 main.main main.go:60]"
`
	_, err := igop.RunFile("main.go", src, nil, 0)
	if err != nil {
		t.Fatal(err)
	}
}

func TestRuntimeCallers5(t *testing.T) {
	src := `package main

import (
	"fmt"
	"path/filepath"
	"runtime"
	"strings"
)

var t *struct {
	c chan int
}

var c chan int

func f() {
	select {
	case <-t.c: // THIS IS LINE 18
		break
	case <-c:
		break
	}
}

func g() {
	defer func() {
		// panic("error g1")
	}()
	defer func() {
		// panic("error g2")
		f()
	}()
	f()
}

func main() {
	defer func() {
		recover()
		rpc := make([]uintptr, 64)
		runtime.Callers(0, rpc)
		fs := runtime.CallersFrames(rpc)
		var list []string
		for {
			f, more := fs.Next()
			if f.Function == "runtime.gopanic" {
				list = append(list, "runtime.gopanic")
			} else if strings.HasPrefix(f.Function, "main.") {
				_, fname := filepath.Split(f.File)
				list = append(list, fmt.Sprintf("%v %v:%v", f.Function, fname, f.Line))
			}
			fmt.Println(f)
			if !more {
				break
			}
		}
		if v := fmt.Sprint(list); v != infos {
			panic(v)
		}
	}()
	g()
}

var infos = "[main.main.func1 main.go:40 runtime.gopanic main.f main.go:18 main.g.func2 main.go:31 runtime.gopanic main.f main.go:18 main.g main.go:33 main.main main.go:60]"
`
	_, err := igop.RunFile("main.go", src, nil, 0)
	if err != nil {
		t.Fatal(err)
	}
}

func TestRuntimeCallers6(t *testing.T) {
	src := `package main

import (
	"fmt"
	"path/filepath"
	"runtime"
	"strings"
)

var t *struct {
	c chan int
}

var c chan int

func f() {
	select {
	case <-t.c: // THIS IS LINE 18
		break
	case <-c:
		break
	}
}

func g() {
	defer func() {
		panic("error g1")
	}()
	defer func() {
		// panic("error g2")
		f()
	}()
	// f()
}

func main() {
	defer func() {
		recover()
		rpc := make([]uintptr, 64)
		runtime.Callers(0, rpc)
		fs := runtime.CallersFrames(rpc)
		var list []string
		for {
			f, more := fs.Next()
			if f.Function == "runtime.gopanic" {
				list = append(list, "runtime.gopanic")
			} else if strings.HasPrefix(f.Function, "main.") {
				_, fname := filepath.Split(f.File)
				list = append(list, fmt.Sprintf("%v %v:%v", f.Function, fname, f.Line))
			}
			fmt.Println(f)
			if !more {
				break
			}
		}
		if v := fmt.Sprint(list); v != infos {
			panic(v)
		}
	}()
	g()
}

var infos = "[main.main.func1 main.go:40 runtime.gopanic main.g.func1 main.go:27 runtime.gopanic main.f main.go:18 main.g.func2 main.go:31 main.g main.go:34 main.main main.go:60]"
`
	_, err := igop.RunFile("main.go", src, nil, 0)
	if err != nil {
		t.Fatal(err)
	}
}

func TestRuntimeCallerMethodWrapper(t *testing.T) {
	src := `package main

import (
	"fmt"
	"reflect"
	"runtime"
	"strings"
)

type T int

func (T) f() int { return 0 }

func (*T) g() int { return 0 }

type I interface {
	f() int
}

var (
	index int
)

func check(v interface{}, name string, autogen bool) {
	index++
	pc := reflect.ValueOf(v).Pointer()
	fn := runtime.FuncForPC(pc)
	typ := reflect.ValueOf(v).Type()
	file, line := fn.FileLine(pc)
	fmt.Println(typ, fn.Name(), file, line)
	if name != "" && name != fn.Name() {
		panic(fmt.Errorf("%v: name got %v, want %v", index, fn.Name(), name))
	}
	if autogen {
		if file != "<autogenerated>" || line != 1 {
			panic(fmt.Errorf("%v: file must autogen %v %v", index, file, line))
		}
	} else {
		if !strings.Contains(file, "main.go") {
			panic(fmt.Errorf("%v: file must main.go %v %v", index, file, line))
		}
	}
}

func main() {
	ver := runtime.Version()[:6]
	if ver < "go1.18" {
		// $thunk
		check(T.f, "main.T.f", false)
		check((*T).g, "main.(*T).g", false)
		check((struct{ T }).f, "", true)
		check((struct{ *T }).g, "", true)
		// $bound
		check(T(0).f, "main.T.f-fm", false)
		check(new(T).f, "main.T.f-fm", false)
		check(new(T).g, "main.(*T).g-fm", false)
		// wrapper
		check(I.f, "main.I.f", true)
	} else {
		// $thunk
		check(T.f, "main.T.f", false)
		check((*T).g, "main.(*T).g", false)
		check((struct{ T }).f, "", true)
		check((struct{ *T }).g, "", true)
		// $bound
		check(T(0).f, "main.T.f-fm", true)
		check(new(T).f, "main.T.f-fm", true)
		check(new(T).g, "main.(*T).g-fm", true)
		// wrapper
		check(I.f, "main.I.f", true)
	}
}
`
	_, err := igop.RunFile("main.go", src, nil, 0)
	if err != nil {
		t.Fatal(err)
	}
}

func TestGC15281(t *testing.T) {
	// $GOROOT/test/fixedbugs/issue15281.go
	src := `// run

// Copyright 2016 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import "runtime"

func main() {
	{
		x := inuse()
		c := make(chan []byte, 10)
		c <- make([]byte, 10<<20)
		close(c)
		f1(c, x)
	}
	{
		x := inuse()
		c := make(chan []byte, 10)
		c <- make([]byte, 10<<20)
		close(c)
		f2(c, x)
	}
}

func f1(c chan []byte, start int64) {
	for x := range c {
		if delta := inuse() - start; delta < 9<<20 {
			println("BUG: f1: after alloc: expected delta at least 9MB, got: ", delta)
			println(x)
		}
		x = nil
		if delta := inuse() - start; delta > 1<<20 {
			println("BUG: f1: after alloc: expected delta below 1MB, got: ", delta)
			println(x)
		}
	}
}

func f2(c chan []byte, start int64) {
	for {
		x, ok := <-c
		if !ok {
			break
		}
		if delta := inuse() - start; delta < 9<<20 {
			println("BUG: f2: after alloc: expected delta at least 9MB, got: ", delta)
			println(x)
		}
		x = nil
		if delta := inuse() - start; delta > 1<<20 {
			println("BUG: f2: after alloc: expected delta below 1MB, got: ", delta)
			println(x)
		}
	}
}

func inuse() int64 {
	runtime.GC()
	var st runtime.MemStats
	runtime.ReadMemStats(&st)
	return int64(st.Alloc)
}
`
	_, err := igop.RunFile("main.go", src, nil, igop.ExperimentalSupportGC)
	if err != nil {
		t.Fatal(err)
	}
}

func TestLinknameSource(t *testing.T) {
	pkg := `package pkg
type point struct {
	x int
	y int
}
func (t point) scale(n int) point {
	return point{t.x*n,t.y*n}
}
func (t *point) setScale(n int) {
	t.x *= n
	t.y *= n
}
func (t *point) info() (int,int) {
	return t.x,t.y
}
func add(v ...int) (sum int) {
	for _, n := range v {
		sum += n
	}
	return
}
func New(x int, y int) *point {
	return &point{x,y}
}
`
	src := `package main
import (
	_ "unsafe"
	_ "pkg"
	"fmt"
)

//go:linkname add pkg.add
func add(v ...int) int

type point struct {
	x int
	y int
}

//go:linkname newPoint pkg.New
func newPoint(x int, y int) *point

//go:linkname scalePoint pkg.point.scale
func scalePoint(point,int) point

//go:linkname setScale pkg.(*point).setScale
func setScale(*point,int)

func main() {
	v := add(100,200,300)
	if v != 600 {
		panic(fmt.Errorf("add got %v, must 600",v))
	}
	pt := newPoint(100,200)
	if pt == nil || pt.x != 100 || pt.y != 200 {
		panic(fmt.Errorf("newPoint got %v, must &{100 200}",pt))
	}
	pt2 := scalePoint(*pt,2)
	if pt2.x != 200 || pt2.y != 400 {
		panic(fmt.Errorf("scalePoint got %v, must {200 400}",pt2))
	}
	setScale(pt,3)
	if pt.x != 300 || pt.y != 600 {
		panic(fmt.Errorf("setScale got %v, must &{300 600}",pt))
	}
}
`
	ctx := igop.NewContext(0)
	ctx.AddImportFile("pkg", "pkg.go", pkg)
	_, err := ctx.RunFile("main.go", src, nil)
	if err != nil {
		t.Fatal(err)
	}
}

func TestLinknameExtern(t *testing.T) {
	pkg := `package pkg
func add(v ...int) (sum int) {
	for _, n := range v {
		sum += n
	}
	return
}
`
	src := `package main
import (
	_ "unsafe"
	_ "pkg"
	"fmt"
)

//go:linkname add pkg.add
func add(v ...int) int

type point struct {
	x int
	y int
}

//go:linkname newPoint pkg.New
func newPoint(x int, y int) *point

//go:linkname scalePoint pkg.point.scale
func scalePoint(point,int) point

//go:linkname setScale pkg.(*point).setScale
func setScale(*point,int)

func main() {
	v := add(100,200,300)
	if v != 600 {
		panic(fmt.Errorf("add got %v, must 600",v))
	}
	pt := newPoint(100,200)
	if pt == nil || pt.x != 100 || pt.y != 200 {
		panic(fmt.Errorf("newPoint got %v, must &{100 200}",pt))
	}
	pt2 := scalePoint(*pt,2)
	if pt2.x != 200 || pt2.y != 400 {
		panic(fmt.Errorf("scalePoint got %v, must {200 400}",pt2))
	}
	setScale(pt,3)
	if pt.x != 300 || pt.y != 600 {
		panic(fmt.Errorf("setScale got %v, must &{300 600}",pt))
	}
}
`
	type point struct {
		x int
		y int
	}
	ctx := igop.NewContext(0)
	ctx.AddImportFile("pkg", "pkg.go", pkg)
	ctx.RegisterExternal("pkg.New", func(x int, y int) *point {
		return &point{x, y}
	})
	ctx.RegisterExternal("pkg.point.scale", func(pt point, n int) point {
		return point{pt.x * n, pt.y * n}
	})
	ctx.RegisterExternal("pkg.(*point).setScale", func(pt *point, n int) {
		pt.x *= n
		pt.y *= n
	})
	_, err := ctx.RunFile("main.go", src, nil)
	if err != nil {
		t.Fatal(err)
	}
}

func TestLinknameVar(t *testing.T) {
	pkg := `package pkg
var n int = 100
func N() int {
	return n
}
`
	src := `package main
import (
	_ "unsafe"
	"pkg"
	"fmt"
)

//go:linkname v pkg.n
var v int

//go:linkname e pkg.e
var e int

func main() {
	if v != 100 {
		panic(fmt.Errorf("v got %v, must 100",v))
	}
	v = 200
	if pkg.N() != 200 {
		panic(fmt.Errorf("pkg.N got %v, must 200",pkg.N()))
	}
	if e != 100 {
		panic(fmt.Errorf("3 got %v, must 100",e))
	}
}
`
	ctx := igop.NewContext(0)
	ctx.AddImportFile("pkg", "pkg.go", pkg)
	var e int = 100
	ctx.RegisterExternal("pkg.e", &e)
	_, err := ctx.RunFile("main.go", src, nil)
	if err != nil {
		t.Fatal(err)
	}
}

func TestLinknamePkg(t *testing.T) {
	src := `package main
import (
	_ "unsafe"
	_ "strings"
	"os"
	"bytes"
)

//go:linkname stdout os.Stdout
var stdout *os.File

//go:linkname join strings.Join
func join(elems []string, sep string) string

//go:linkname writeBuffer bytes.(*Buffer).Write
func writeBuffer(buf *bytes.Buffer, data []byte) (n int, err error)

func main() {
	if stdout != os.Stdout {
		panic("error stdout")
	}
	if v := join([]string{"hello","world"},","); v != "hello,world" {
		panic("error join "+v)
	}
	var buf bytes.Buffer
	writeBuffer(&buf,[]byte("hello"))
	if v := buf.String(); v != "hello" {
		panic("error write buffer "+v)
	}
}
`
	ctx := igop.NewContext(0)
	_, err := ctx.RunFile("main.go", src, nil)
	if err != nil {
		t.Fatal(err)
	}
}

func TestLinknameError1(t *testing.T) {
	pkg := `package pkg
func add(v ...int) (sum int) {
	for _, n := range v {
		sum += n
	}
	return
}
`
	src := `package main
import (
	_ "pkg"
)

//go:linkname add pkg.add
func add(v ...int) int

func main() {
}
`
	ctx := igop.NewContext(0)
	ctx.AddImportFile("pkg", "pkg.go", pkg)
	_, err := ctx.RunFile("main.go", src, nil)
	if err == nil {
		t.Fatal("must error")
	}
	if err.Error() != `main.go:6:1: //go:linkname only allowed in Go files that import "unsafe"` {
		t.Fatal(err)
	}
}

func TestLinknameError2(t *testing.T) {
	pkg := `package pkg
func add(v ...int) (sum int) {
	for _, n := range v {
		sum += n
	}
	return
}
`
	src := `package main
import (
	_ "unsafe"
	_ "pkg"
)

//go:linkname add2 pkg.add
func add(v ...int) int

func main() {
}
`
	ctx := igop.NewContext(0)
	ctx.AddImportFile("pkg", "pkg.go", pkg)
	_, err := ctx.RunFile("main.go", src, nil)
	if err == nil {
		t.Fatal("must error")
	}
	if err.Error() != `main.go:7:1: //go:linkname must refer to declared function or variable` {
		t.Fatal(err)
	}
}

func TestLinknameError3(t *testing.T) {
	pkg := `package pkg
var v int = 100
`
	src := `package main
import (
	_ "unsafe"
	_ "pkg"
)

//go:linkname a pkg.v
var a int = 200

func main() {
}
`
	ctx := igop.NewContext(0)
	ctx.AddImportFile("pkg", "pkg.go", pkg)
	_, err := ctx.RunFile("main.go", src, nil)
	if err == nil {
		t.Fatal("must error")
	}
	if err.Error() != `duplicated definition of symbol pkg.v, from main and pkg` {
		t.Fatal(err)
	}
}
