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

var (
	flag          = &Cmd.Flag
	flagDumpInstr bool
	flagGoplus    bool
)

func init() {
	Cmd.Run = runCmd
	flag.BoolVar(&flagGoplus, "gop", true, "support goplus")
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
	welcomeGo  string = "Welcome to Go REPL!"
	welcomeGop string = "Welcome to Go+ REPL!"
)

var (
	supportGoplus bool
)

func runCmd(cmd *base.Command, args []string) {
	flag.Parse(args)

	if supportGoplus && flagGoplus {
		fmt.Println(welcomeGop)
	} else {
		fmt.Println(welcomeGo)
	}

	state := liner.NewLiner()
	defer state.Close()

	// state.SetCtrlCAborts(true)
	state.SetMultiLineMode(true)
	ui := &LinerUI{state: state}
	r := repl.NewREPL()
	r.SetUI(ui)
	if supportGoplus && flagGoplus {
		r.SetFileName("main.gop")
	}
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
