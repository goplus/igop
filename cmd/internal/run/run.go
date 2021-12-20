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
	"log"
	"os"
	"path/filepath"

	"github.com/goplus/gossa"
	"github.com/goplus/gossa/cmd/internal/base"
	"github.com/goplus/gossa/gopbuild"
)

// -----------------------------------------------------------------------------

// Cmd - gop run
var Cmd = &base.Command{
	UsageLine: "gossa run <gopSrcDir|gopSrcFile> [arguments]",
	Short:     "Run a Go+ program",
}

var (
	flag = &Cmd.Flag
)

func init() {
	Cmd.Run = runCmd
}

func runCmd(cmd *base.Command, args []string) {
	flag.Parse(args)
	if flag.NArg() < 1 {
		cmd.Usage(os.Stderr)
	}
	args = flag.Args()[1:]
	path, _ := filepath.Abs(flag.Arg(0))
	isDir, err := IsDir(path)
	if err != nil {
		log.Fatalln("input arg check failed:", err)
	}
	ctx := gossa.NewContext(0)
	if isDir {
		if containsExt(path, ".gop") {
			data, err := gopbuild.BuildDir(ctx, path)
			if err != nil {
				log.Fatalln(err)
			}
			err = os.WriteFile(filepath.Join(path, "gop_autogen.go"), data, 0666)
			if err != nil {
				log.Fatalln(err)
			}
		}
		runDir(ctx, path, args)
	} else {
		runFile(ctx, path, args)
	}
}

// IsDir checks a target path is dir or not.
func IsDir(target string) (bool, error) {
	fi, err := os.Stat(target)
	if err != nil {
		return false, err
	}
	return fi.IsDir(), nil
}

func runFile(ctx *gossa.Context, target string, args []string) {
	dir, file := filepath.Split(target)
	os.Chdir(dir)
	exitCode, err := ctx.RunFile(file, nil, args)
	if err != nil {
		log.Println(err)
	}
	os.Exit(exitCode)
}

func runDir(ctx *gossa.Context, dir string, args []string) {
	os.Chdir(dir)
	exitCode, err := ctx.Run(dir, args)
	if err != nil {
		log.Println(err)
	}
	os.Exit(exitCode)
}

func containsExt(srcDir string, ext string) bool {
	if f, err := os.Open(srcDir); err == nil {
		defer f.Close()
		fis, _ := f.Readdir(-1)
		for _, fi := range fis {
			if !fi.IsDir() && filepath.Ext(fi.Name()) == ext {
				return true
			}
		}
	}
	return false
}

// -----------------------------------------------------------------------------
