package goxtest

import (
	_ "embed"

	"github.com/goplus/igop"
	_ "github.com/goplus/igop/pkg/github.com/goplus/igop/x/testdeps"
	_ "github.com/goplus/igop/pkg/testing"
)

//go:embed _test/classfile.go
var src string

func Register(ctx *igop.Context) {
	ctx.AddImportFile("github.com/goplus/gop/test", "classfile.go", src)
}
