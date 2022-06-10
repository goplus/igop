package repl

import (
	"fmt"
	"go/token"
	"go/types"
	"os/exec"
	"reflect"
	"strings"

	"golang.org/x/tools/go/ssa"

	"github.com/goplus/igop"
)

const (
	// ContinuePrompt - the current code statement is not completed.
	ContinuePrompt string = "... "
	// NormalPrompt - start of a code statement.
	NormalPrompt string = ">>> "
)

type UI interface {
	SetPrompt(prompt string)
	Printf(format string, a ...interface{})
}

type REPL struct {
	*igop.Repl
	term UI
	more string
}

func NewREPL(mode igop.Mode) *REPL {
	ctx := igop.NewContext(mode)
	repl := igop.NewRepl(ctx)
	r := &REPL{
		Repl: repl,
	}

	ctx.SetOverrideFunction("main.__igop_repl_info__", func(name string, v interface{}) {
		r.Printf("(main) %v %v %v\n", name, reflect.TypeOf(v), v)
	})

	return r
}

func (r *REPL) SetUI(term UI) {
	r.term = term
	term.SetPrompt(NormalPrompt)
}

func (r *REPL) SetNormal() {
	r.more = ""
	r.SetPrompt(NormalPrompt)
}

func (r *REPL) SetPrompt(prompt string) {
	if r.term != nil {
		r.term.SetPrompt(prompt)
	}
}

func (r *REPL) IsNormal() bool {
	return r.more == ""
}

func (r *REPL) Dump(expr string) {
	i := r.Interp()
	if i == nil {
		r.godoc(expr)
		return
	}
	pkg := i.MainPkg()
	if m, ok := pkg.Members[expr]; ok {
		switch v := m.(type) {
		case *ssa.NamedConst:
			r.Printf("(global) const %v %v = %v\n", v.Name(), v.Type(), v.Value.Value)
		case *ssa.Global:
			r.Printf("(global) var %v %v\n", v.Name(), v.Type().(*types.Pointer).Elem())
		case *ssa.Type:
			r.Printf("(global) type %v %v\n", v.Name(), v.Type())
		case *ssa.Function:
			r.Printf("(global) func %v %v\n", v.Name(), v.Type())
		}
		return
	}
	_, _, err := r.Eval(fmt.Sprintf("__igop_repl_info__(%q,%v)", expr, expr))
	if err != nil {
		if err := r.godoc(expr); err != nil {
			r.Printf("not found %v\n", expr)
		}
	}
}

func (r *REPL) godoc(expr string) error {
	gobin, err := exec.LookPath("go")
	if err != nil {
		return err
	}
	cmd := exec.Command(gobin, "doc", expr)
	data, err := cmd.CombinedOutput()
	if err != nil {
		return err
	}
	r.Printf("%v\n", string(data))
	return nil
}

func (r *REPL) Printf(_fmt string, a ...interface{}) {
	if r.term != nil {
		r.term.Printf(_fmt, a...)
	} else {
		fmt.Printf(_fmt, a...)
	}
}

func (r *REPL) Run(line string) {
	var expr string
	if r.more != "" {
		if line == "" {
			r.SetNormal()
			return
		}
		expr = r.more + "\n" + line
	} else {
		if strings.HasPrefix(line, "?") {
			r.Dump(strings.TrimSpace(line[1:]))
			return
		}
		expr = line
	}
	tok, dump, err := r.Eval(expr)
	if err != nil {
		if checkMore(tok, err) {
			r.more += "\n" + line
			r.SetPrompt(ContinuePrompt)
		} else {
			r.SetNormal()
			r.Printf("error: %v\n", err)
		}
		return
	}
	if len(dump) > 0 {
		r.Printf(": %v\n", strings.Join(dump, " "))
	}
	r.SetNormal()
}

func checkMore(tok token.Token, err error) bool {
	s := err.Error()
	if strings.Contains(s, `expected `) {
		return true
	}
	return false
}
