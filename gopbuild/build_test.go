package gopbuild

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/goplus/igop"
)

func gopClTest(t *testing.T, gopcode, expected string) {
	ctx := igop.NewContext(0)
	data, err := BuildFile(ctx, "main.gop", gopcode)
	if err != nil {
		t.Fatalf("build gop error: %v", err)
	}
	if string(data) != expected {
		fmt.Println("build gop error:")
		fmt.Println(string(data))
		t.Fail()
	}
}

func TestGop(t *testing.T) {
	gopClTest(t, `
println "Go+"
`, `package main

import fmt "fmt"

func main() {
//line main.gop:2
	fmt.Println("Go+")
}
`)
}

func TestBig(t *testing.T) {
	gopClTest(t, `
a := 1/2r
println a+1/2r
`, `package main

import (
	fmt "fmt"
	ng "github.com/goplus/gop/builtin/ng"
	big "math/big"
)

func main() {
//line main.gop:2
	a := ng.Bigrat_Init__2(big.NewRat(1, 2))
//line main.gop:3
	fmt.Println(a.Gop_Add(ng.Bigrat_Init__2(big.NewRat(1, 2))))
}
`)
}

func TestBuiltin(t *testing.T) {
	igop.RegisterCustomBuiltin("typeof", reflect.TypeOf)
	gopClTest(t, `
v := typeof(100)
println(v)
`, `package main

import fmt "fmt"

func main() {
//line main.gop:2
	v := typeof(100)
//line main.gop:3
	fmt.Println(v)
}
`)
}

func TestIoxLines(t *testing.T) {
	gopClTest(t, `
import "io"

var r io.Reader

for line <- lines(r) {
	println line
}
`, `package main

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
`)
}
