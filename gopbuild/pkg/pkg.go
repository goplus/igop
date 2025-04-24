package pkg

//go:generate go run ../../cmd/qexp -outdir ../../pkg github.com/qiniu/x/gsh
//go:generate go run ../../cmd/qexp -outdir ../../pkg github.com/goplus/gop/ast
//go:generate go run ../../cmd/qexp -outdir ../../pkg github.com/goplus/gop/parser
//go:generate go run ../../cmd/qexp -outdir ../../pkg github.com/goplus/gop/printer
//go:generate go run ../../cmd/qexp -outdir ../../pkg github.com/goplus/gop/scanner
//go:generate go run ../../cmd/qexp -outdir ../../pkg github.com/goplus/gop/token
//go:generate go run ../../cmd/qexp -outdir ../../pkg github.com/goplus/gop/tpl/...
//go:generate go run ../../cmd/qexp -outdir ../../pkg -code github.com/goplus/gop/test

import (
	_ "github.com/goplus/igop/gopbuild/pkg/gop"
	_ "github.com/goplus/igop/gopbuild/pkg/gsh"
	_ "github.com/goplus/igop/gopbuild/pkg/test"
	_ "github.com/goplus/igop/gopbuild/pkg/tpl"
)
