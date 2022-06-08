//go:build go1.16
// +build go1.16

package repl

import (
	"fmt"

	"github.com/goplus/gop/env"
	_ "github.com/goplus/igop/gopbuild"
)

func init() {
	supportGoplus = true
	welcomeGop = fmt.Sprintf("Welcome to Go+ (version %v) REPL!", env.Version())
}
