package gopbuild

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/goplus/igop"
	_ "github.com/goplus/igop/pkg/reflect"
)

var test_gop = `println "Go+"`
var test_gop_go = `package main

import fmt "fmt"

func main() {
//line main.gop:1
	fmt.Println("Go+")
}
`

func TestGop(t *testing.T) {
	ctx := igop.NewContext(0)
	data, err := BuildFile(ctx, "main.gop", test_gop)
	if err != nil {
		t.Fatalf("build gop error: %v", err)
	}
	if string(data) != test_gop_go {
		fmt.Println("build gop error:")
		fmt.Println(string(data))
		t.Fail()
	}
}

var test_big = `a := 1/2r
println a+1/2r
`
var test_big_go = `package main

import (
	fmt "fmt"
	ng "github.com/goplus/gop/builtin/ng"
	big "math/big"
)

func main() {
//line main.gop:1
	a := ng.Bigrat_Init__2(big.NewRat(1, 2))
//line main.gop:2
	fmt.Println(a.Gop_Add(ng.Bigrat_Init__2(big.NewRat(1, 2))))
}
`

func TestBig(t *testing.T) {
	ctx := igop.NewContext(0)
	data, err := BuildFile(ctx, "main.gop", test_big)
	if err != nil {
		t.Fatalf("build gop error: %v", err)
	}
	if string(data) != test_big_go {
		fmt.Println("build gop error:")
		fmt.Println(string(data))
		t.Fail()
	}
}

var test_builtin = `
v := typeof(100)
println(v)
`
var test_builtin_go = `package main

import fmt "fmt"

func main() {
//line main.gop:2
	v := typeof(100)
//line main.gop:3
	fmt.Println(v)
}
`

func TestBuiltin(t *testing.T) {
	ctx := igop.NewContext(0)
	igop.RegisterCustomBuiltin("typeof", reflect.TypeOf)
	data, err := BuildFile(ctx, "main.gop", test_builtin)
	if err != nil {
		t.Fatalf("build gop error: %v", err)
	}
	if string(data) != test_builtin_go {
		fmt.Println("build gop error:")
		fmt.Println(string(data))
		t.Fail()
	}
}

var test_iox = `
import "io"

var r io.Reader

for line <- lines(r) {
	println line
}
`
var test_iox_go = `package main

import (
	fmt "fmt"
	iox "github.com/goplus/gop/builtin/iox"
	io "io"
)

var r io.Reader

func main() {
//line main.gop:6
	for _gop_it := iox.Lines(r).Gop_Enum(); ; {
		var _gop_ok bool
		line, _gop_ok := _gop_it.Next()
		if !_gop_ok {
			break
		}
//line main.gop:7
		fmt.Println(line)
	}
}
`

func TestIox(t *testing.T) {
	ctx := igop.NewContext(0)
	data, err := BuildFile(ctx, "main.gop", test_iox)
	if err != nil {
		t.Fatalf("build gop error: %v", err)
	}
	if string(data) != test_iox_go {
		fmt.Println("build gop error:")
		fmt.Println(string(data))
		t.Fail()
	}
}
