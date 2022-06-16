package repl

import (
	"fmt"
	"io"
	"runtime"
	"strings"

	"github.com/goplus/igop"
	"github.com/goplus/igop/cmd/internal/base"
	_ "github.com/goplus/igop/pkg"
	"github.com/goplus/igop/repl"
	"github.com/peterh/liner"
)

// Cmd - igop test
var Cmd = &base.Command{
	UsageLine: "igop repl",
	Short:     "igop repl mode",
}

var (
	flag          = &Cmd.Flag
	flagGoplus    bool
	flagDumpInstr bool
	flagTrace     bool
)

func init() {
	Cmd.Run = runCmd
	flag.BoolVar(&flagGoplus, "gop", true, "support goplus")
	flag.BoolVar(&flagDumpInstr, "dump", false, "dump SSA instruction code")
	flag.BoolVar(&flagTrace, "trace", false, "trace interpreter code")
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
	welcomeGo  string = fmt.Sprintf("Welcome to Go (version %v) REPL!", runtime.Version())
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
	state.SetCompleter(func(line string) []string {
		if strings.TrimSpace(line) == "" {
			return []string{line + "    "}
		}
		return nil
	})
	ui := &LinerUI{state: state}
	var mode igop.Mode
	if flagDumpInstr {
		mode |= igop.EnableDumpInstr
	}
	if flagTrace {
		mode |= igop.EnableTracing
	}
	r := repl.NewREPL(mode)
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
		err = r.Run(line)
		switch e := err.(type) {
		case nil:
			//
		case igop.ExitError:
			fmt.Printf("exit %v\n", int(e))
			return
		default:
			fmt.Printf("error: %v\n", err)
		}
	}
}
