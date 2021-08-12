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

// Package run implements the ``igop run'' command.
package run

import (
	"log"
	"os"

	"github.com/goplus/interp"
	"github.com/goplus/interp/cmd/internal/base"
	_ "github.com/goplus/interp/pkg"
)

// -----------------------------------------------------------------------------

// Cmd - igop run
var Cmd = &base.Command{
	UsageLine: "igop run [-v] package",
	Short:     "interpret package",
}

var (
	flag        = &Cmd.Flag
	flagVerbose = flag.Bool("v", false, "print verbose information")
)

func init() {
	Cmd.Run = runCmd
}

func runCmd(cmd *base.Command, args []string) {
	flag.Parse(args)
	if flag.NArg() < 1 {
		cmd.Usage(os.Stderr)
	}
	var mode interp.Mode
	if *flagVerbose {
		mode = interp.EnableTracing
	}
	iargs := flag.Args()
	err := interp.RunWith(mode, iargs[0], iargs[1:])
	if err != nil {
		log.Fatalln("interpret package failed:", err)
	}
}
