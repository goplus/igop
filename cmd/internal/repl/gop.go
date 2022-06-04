//go:build go1.16
// +build go1.16

package repl

import (
	_ "github.com/goplus/gossa/gopbuild"
)

func init() {
	supportGoplus = true
}
