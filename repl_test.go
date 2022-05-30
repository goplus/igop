package gossa_test

import (
	"fmt"
	"testing"

	"github.com/goplus/gossa"
)

func TestReplExpr(t *testing.T) {
	ctx := gossa.NewContext(0)
	repl := gossa.NewRepl(ctx)
	list := []string{
		`a := 1`,
		`b := 2`,
		`a+b`,
	}
	result := []string{
		`1`,
		`2`,
		`3`,
	}
	for i, expr := range list {
		v, err := repl.Eval(expr)
		if err != nil {
			t.Fatal(err)
		}
		if fmt.Sprint(v) != result[i] {
			t.Fatalf("%v", v)
		}
	}
}

func TestReplImports(t *testing.T) {
	ctx := gossa.NewContext(0)
	repl := gossa.NewRepl(ctx)
	list := []string{
		`a := 1`,
		`b := 2`,
		`import "fmt"`,
		`c := fmt.Sprintf("%v-%v",a,b)`,
		`c`,
	}
	result := []string{
		`1`,
		`2`,
		`<nil>`,
		`1-2`,
		`1-2`,
	}
	for i, expr := range list {
		v, err := repl.Eval(expr)
		if err != nil {
			t.Fatal(err)
		}
		if fmt.Sprint(v) != result[i] {
			t.Fatalf("%v", v)
		}
	}
}
