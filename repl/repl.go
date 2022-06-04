package repl

import (
	"go/token"
	"strings"

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
	repl *igop.Repl
	term UI
	more string
}

func NewREPL(mode igop.Mode) *REPL {
	ctx := igop.NewContext(mode)
	repl := igop.NewRepl(ctx)
	return &REPL{
		repl: repl,
	}
}

func (r *REPL) SetUI(term UI) {
	r.term = term
	term.SetPrompt(NormalPrompt)
}

func (r *REPL) SetNormal() {
	r.more = ""
	r.term.SetPrompt(NormalPrompt)
}

func (r *REPL) IsNormal() bool {
	return r.more == ""
}

func (r *REPL) SetFileName(filename string) {
	r.repl.SetFileName(filename)
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
		expr = line
	}
	tok, dump, err := r.repl.Eval(expr)
	if err != nil {
		if checkMore(tok, err) {
			r.more += "\n" + line
			r.term.SetPrompt(ContinuePrompt)
		} else {
			r.SetNormal()
			r.term.Printf("error: %v\n", err)
		}
		return
	}
	if len(dump) > 0 {
		r.term.Printf(": %v\n", strings.Join(dump, " "))
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
