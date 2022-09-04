package main

import (
	"os"

	"github.com/goplus/igop/cmd/internal/base"
	"github.com/goplus/igop/cmd/internal/export"
)

func main() {
	base.Main(export.Cmd, "qexp", os.Args[1:])
}
