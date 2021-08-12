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

// Package test implements the ``igop test'' command.
package test

import (
	"log"
	"strings"

	"github.com/goplus/interp"
	"github.com/goplus/interp/cmd/internal/base"
	_ "github.com/goplus/interp/pkg"
)

// Cmd - igop test
var Cmd = &base.Command{
	UsageLine: "igop test [-v] package",
	Short:     "test package",
}

var (
	flag = &Cmd.Flag
	_    = flag.Bool("v", false, "print the names of packages as they are compiled.")
)

func init() {
	Cmd.Run = runCmd
}

func skipSwitches(args []string) []string {
	out := make([]string, 0, len(args))
	for _, arg := range args {
		if strings.HasPrefix(arg, "-") {
			continue
		}
		out = append(out, arg)
	}
	return out
}

func runCmd(cmd *base.Command, args []string) {
	ssargs := skipSwitches(args)
	if len(ssargs) == 0 {
		ssargs = []string{"."}
	}
	for _, pkg := range ssargs {
		err := interp.RunWithTest(0, pkg, nil)
		if err != nil {
			log.Fatalln("interpret test package failed:", err)
		}
	}
}
