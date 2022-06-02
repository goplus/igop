package repl

import (
	"fmt"
	"io"

	"github.com/goplus/gossa/cmd/internal/base"
	_ "github.com/goplus/gossa/pkg"
	"github.com/goplus/gossa/repl"
	"github.com/peterh/liner"
)

// Cmd - gossa test
var Cmd = &base.Command{
	UsageLine: "gossa repl",
	Short:     "gossa repl mode",
}

func init() {
	Cmd.Run = runCmd
}

// LinerUI implements repl.UI interface.
type LinerUI struct {
	state  *liner.State
	prompt string
}

// SetPrompt is required by repl.UI interface.
func (u *LinerUI) SetPrompt(prompt string) {
	u.prompt = prompt
}

// Printf is required by repl.UI interface.
func (u *LinerUI) Printf(format string, a ...interface{}) {
	fmt.Printf(format, a...)
}

var (
	welcome string = "Welcome to Go REPL!"
)

func runCmd(cmd *base.Command, args []string) {
	fmt.Println(welcome)

	state := liner.NewLiner()
	defer state.Close()

	// state.SetCtrlCAborts(true)
	state.SetMultiLineMode(true)
	ui := &LinerUI{state: state}
	r := repl.NewREPL()
	r.SetUI(ui)
	for {
		line, err := ui.state.Prompt(ui.prompt)
		if err != nil {
			if err == liner.ErrPromptAborted || err == io.EOF {
				fmt.Printf("exit\n")
				break
			}
			fmt.Printf("Problem reading line: %v\n", err)
			continue
		}
		if line != "" {
			state.AppendHistory(line)
		}
		r.Run(line)
	}
}
