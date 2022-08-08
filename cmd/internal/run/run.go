/*
 Copyright 2021 The GoPlus Authors (goplus.org)

 Licensed under the Apache License, Version 2.0 (the "License");
 you may not use this file except in compliance with the License.
 You may obtain a copy of the License at

     http://www.apache.org/licenses/LICENSE-2.0

 Unless required by applicable law or agreed to in writing, software
 distributed under the License is distributed on an "AS IS" BASIS,
 WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 See the License for the specific language governing permissions and
 limitations under the License.
*/

// Package run implements the ``gop run'' command.
package run

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/goplus/igop"
	"github.com/goplus/igop/cmd/internal/base"
	"github.com/goplus/igop/cmd/internal/load"
)

// -----------------------------------------------------------------------------

// Cmd - igop run
var Cmd = &base.Command{
	UsageLine: "igop run <gopSrcDir|gopSrcFile> [arguments]",
	Short:     "Run a Go/Go+ program",
}

var (
	flag          = &Cmd.Flag
	flagDumpInstr bool
	flagDumpPkg   bool
	flagTrace     bool
	flagTags      string
)

func init() {
	Cmd.Run = runCmd
	flag.BoolVar(&flagDumpInstr, "dumpssa", false, "print SSA instruction code")
	flag.BoolVar(&flagDumpPkg, "dumppkg", false, "print load import packages")
	flag.BoolVar(&flagTrace, "trace", false, "trace interpreter code")
	flag.StringVar(&flagTags, "tags", "", "a comma-separated list of build tags to consider satisfied during the build.")
}

func runCmd(cmd *base.Command, args []string) {
	err := flag.Parse(args)
	if err != nil {
		os.Exit(2)
	}
	if flag.NArg() < 1 {
		cmd.Usage(os.Stderr)
	}
	args = flag.Args()[1:]
	path, _ := filepath.Abs(flag.Arg(0))
	isDir, err := load.IsDir(path)
	if err != nil {
		log.Fatalln("input arg check failed:", err)
	}
	var mode igop.Mode
	if flagDumpInstr {
		mode |= igop.EnableDumpInstr
	}
	if flagTrace {
		mode |= igop.EnableTracing
	}
	if flagDumpPkg {
		mode |= igop.EnableDumpImports
	}
	ctx := igop.NewContext(mode)
	if flagTags != "" {
		ctx.BuildContext.BuildTags = strings.Split(flagTags, ",")
	}
	if isDir {
		if load.SupportGop && load.IsGopProject(path) {
			err := load.BuildGopDir(ctx, path)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(2)
			}
		}
		runDir(ctx, path, args)
	} else {
		runFile(ctx, path, args)
	}
}

func runFile(ctx *igop.Context, target string, args []string) {
	exitCode, err := ctx.RunFile(target, nil, args)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	os.Exit(exitCode)
}

func runDir(ctx *igop.Context, dir string, args []string) {
	exitCode, err := ctx.Run(dir, args)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	os.Exit(exitCode)
}
