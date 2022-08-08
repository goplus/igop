//go:build !go1.16
// +build !go1.16

package load

import (
	"github.com/goplus/igop"
)

const (
	SupportGop = false
)

func BuildGopDir(ctx *igop.Context, path string) error {
	return nil
}

func IsGopProject(dir string) bool {
	return false
}
