package gopbuild

import (
	"fmt"
	"go/token"
	"testing"

	"github.com/goplus/gossa"
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
	ctx := gossa.NewContext(0)
	fset := token.NewFileSet()
	data, err := BuildFile(ctx, fset, "main.gop", test_gop)
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
	builtin "github.com/goplus/gop/builtin"
	big "math/big"
)

func main() {
//line main.gop:1
	a := builtin.Gop_bigrat_Init__2(big.NewRat(1, 2))
//line main.gop:2
	fmt.Println(a.Gop_Add(builtin.Gop_bigrat_Init__2(big.NewRat(1, 2))))
}
`

func TestBig(t *testing.T) {
	ctx := gossa.NewContext(0)
	fset := token.NewFileSet()
	data, err := BuildFile(ctx, fset, "main.gop", test_big)
	if err != nil {
		t.Fatalf("build gop error: %v", err)
	}
	if string(data) != test_big_go {
		fmt.Println("build gop error:")
		fmt.Println(string(data))
		t.Fail()
	}
}
