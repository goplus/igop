package igop_test

import (
	"fmt"
	"go/token"
	"testing"

	"github.com/goplus/igop"
)

func TestReplExpr(t *testing.T) {
	ctx := igop.NewContext(0)
	repl := igop.NewRepl(ctx)
	list := []string{
		`a := 1`,
		`b := 2`,
		`a+b`,
	}
	result := []string{
		`[1]`,
		`[2]`,
		`[3]`,
	}
	for i, expr := range list {
		_, v, err := repl.Eval(expr)
		if err != nil {
			t.Fatal(err)
		}
		if fmt.Sprint(v) != result[i] {
			t.Fatalf("expr:%v dump:%v src:%v", expr, v, repl.Source())
		}
	}
}

func TestReplImports(t *testing.T) {
	ctx := igop.NewContext(0)
	repl := igop.NewRepl(ctx)
	list := []string{
		`a := 1`,
		`b := 2`,
		`import "fmt"`,
		`c := fmt.Sprintf("%v-%v",a,b)`,
		`c`,
	}
	result := []string{
		`[1]`,
		`[2]`,
		`[]`,
		`[1-2]`,
		`[1-2]`,
	}
	for i, expr := range list {
		_, v, err := repl.Eval(expr)
		if err != nil {
			t.Fatal(err)
		}
		if fmt.Sprint(v) != result[i] {
			t.Fatalf("expr:%v dump:%v src:%v", expr, v, repl.Source())
		}
	}
}

func TestReplClosure(t *testing.T) {
	ctx := igop.NewContext(0)
	repl := igop.NewRepl(ctx)
	list := []string{
		`a := 1`,
		`b := 2`,
		`fn := func() int { return a }`,
		`d := fn()+b`,
		`d`,
	}
	result := []string{
		`[1]`,
		`[2]`,
		`-`,
		`[3]`,
		`[3]`,
	}
	for i, expr := range list {
		_, v, err := repl.Eval(expr)
		if err != nil {
			t.Fatal(err)
		}
		if result[i] == "-" {
			continue
		}
		if fmt.Sprint(v) != result[i] {
			t.Fatalf("expr:%v dump:%v src:%v", expr, v, repl.Source())
		}
	}
}

func TestReplVar(t *testing.T) {
	ctx := igop.NewContext(0)
	repl := igop.NewRepl(ctx)
	list := []string{
		`a := 1`,
		`b := 2`,
		`var c = 100`,
		`fn := func() int { return a+c }`,
		`d := fn()+b`,
		`d`,
		`c`,
	}
	result := []string{
		`[1]`,
		`[2]`,
		`-`,
		`-`,
		`[103]`,
		`[103]`,
		`[100]`,
	}
	for i, expr := range list {
		_, v, err := repl.Eval(expr)
		if err != nil {
			t.Fatal(err)
		}
		if result[i] == "-" {
			continue
		}
		if fmt.Sprint(v) != result[i] {
			t.Fatalf("expr:%v dump:%v src:%v", expr, v, repl.Source())
		}
	}
}

func TestReplType(t *testing.T) {
	ctx := igop.NewContext(0)
	repl := igop.NewRepl(ctx)
	list := []string{
		`type T struct {
	X int
	Y int
}`,
		`v1 := &T{10,20}`,
		`import "fmt"`,
		`r1 := fmt.Sprint(v1)`,
		`func (t *T) String() string {
	return fmt.Sprintf("%v-%v",t.X,t.Y)
}`,
		`v2 := &T{10,20}`,
		`r2 := fmt.Sprint(v2)`,
	}
	result := []string{
		`-`,
		`[&{10 20}]`,
		`-`,
		`[&{10 20}]`,
		`-`,
		`[10-20]`,
		`[10-20]`,
	}
	for i, expr := range list {
		_, v, err := repl.Eval(expr)
		if err != nil {
			t.Fatal(err)
		}
		if result[i] == "-" {
			continue
		}
		if fmt.Sprint(v) != result[i] {
			t.Fatalf("expr:%v dump:%v src:%v", expr, v, repl.Source())
		}
	}
}

func TestReplFunc(t *testing.T) {
	ctx := igop.NewContext(0)
	repl := igop.NewRepl(ctx)
	list := []string{
		`a := "hello"`,
		`import "fmt"`,
		`a`,
		`fmt.Println(a)`,
		`s := fmt.Sprint(a)`,
		`fmt.Println(s)`,
	}
	result := []string{
		`[hello]`,
		`-`,
		`[hello]`,
		`[6 <nil>]`,
		`[hello]`,
		`[6 <nil>]`,
	}
	for i, expr := range list {
		_, v, err := repl.Eval(expr)
		if err != nil {
			t.Fatal(err)
		}
		if result[i] == "-" {
			continue
		}
		if fmt.Sprint(v) != result[i] {
			t.Fatalf("expr:%v dump:%v src:%v", expr, v, repl.Source())
		}
	}
}

func TestReplTok(t *testing.T) {
	ctx := igop.NewContext(0)
	repl := igop.NewRepl(ctx)
	list := []string{
		`a := "hello"`,
		`import "fmt"`,
		`var c int`,
		`const d = 100`,
		`func test(v int) {}`,
		`fmt.Println(a)`,
		`type T int`,
	}
	toks := []token.Token{
		token.IDENT,
		token.IMPORT,
		token.VAR,
		token.CONST,
		token.FUNC,
		token.IDENT,
		token.TYPE,
	}
	for i, expr := range list {
		tok, _, err := repl.Eval(expr)
		if err != nil {
			t.Fatal(err)
		}
		if tok != toks[i] {
			t.Fatalf("expr:%v token:%v", expr, tok)
		}
	}
}

func TestReplInit(t *testing.T) {
	ctx := igop.NewContext(0)
	repl := igop.NewRepl(ctx)
	list := []string{
		`var i int = 10`,
		`var j int = 20`,
		`func init() { i++ }`,
		`func init() { j++ }`,
		`var c int`,
		`i`,
		`j`,
	}
	result := []string{
		`-`,
		`-`,
		`-`,
		`-`,
		`-`,
		`[11]`,
		`[21]`,
	}
	for i, expr := range list {
		_, v, err := repl.Eval(expr)
		if err != nil {
			t.Fatal(err)
		}
		if result[i] == "-" {
			continue
		}
		if fmt.Sprint(v) != result[i] {
			t.Fatalf("expr:%v dump:%v src:%v", expr, v, repl.Source())
		}
	}
}
