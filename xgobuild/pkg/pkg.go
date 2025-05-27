package pkg

//go:generate go run ../../cmd/qexp -outdir ../../pkg github.com/qiniu/x/gsh
//go:generate go run ../../cmd/qexp -outdir ../../pkg github.com/goplus/xgo/ast
//go:generate go run ../../cmd/qexp -outdir ../../pkg github.com/goplus/xgo/parser
//go:generate go run ../../cmd/qexp -outdir ../../pkg github.com/goplus/xgo/printer
//go:generate go run ../../cmd/qexp -outdir ../../pkg github.com/goplus/xgo/scanner
//go:generate go run ../../cmd/qexp -outdir ../../pkg github.com/goplus/xgo/token
//go:generate go run ../../cmd/qexp -outdir ../../pkg github.com/goplus/xgo/tpl/...
//go:generate go run ../../cmd/qexp -outdir ../../pkg -code github.com/goplus/xgo/test

import (
	_ "github.com/goplus/ixgo/xgobuild/pkg/gsh"
	_ "github.com/goplus/ixgo/xgobuild/pkg/test"
	_ "github.com/goplus/ixgo/xgobuild/pkg/tpl"
	_ "github.com/goplus/ixgo/xgobuild/pkg/xgo"
)
