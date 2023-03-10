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

// Package test implements the “igop test” command.
package test

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/goplus/igop"
	"github.com/goplus/igop/cmd/internal/base"
)

// Cmd - igop test
var Cmd = &base.Command{
	UsageLine: "igop test [build/test flags] [package]",
	Short:     "test a package",
}

func init() {
	Cmd.Run = runCmd
	base.AddBuildFlags(Cmd, base.OmitModFlag|base.OmitSSAFlag|base.OmitSSATraceFlag|
		base.OmitVFlag|base.OmitExperimentalGCFlag)
}

func runCmd(cmd *base.Command, args []string) {
	cf := &Cmd.Flag
	cf.String("bench", "", "")
	cf.Bool("benchmem", false, "")
	cf.Duration("benchtime", 1*time.Second, "")
	cf.String("blockprofile", "", "")
	cf.Int("blockprofilerate", 1, "")
	cf.Int("count", 1, "")
	cf.String("coverprofile", "", "")
	cf.String("cpu", "", "")
	cf.String("cpuprofile", "", "")
	cf.Bool("failfast", false, "")
	cf.String("list", "", "")
	cf.String("memprofile", "", "")
	cf.Int("memprofilerate", 0, "")
	cf.String("mutexprofile", "", "")
	cf.Int("mutexprofilefraction", 1, "")
	cf.String("outputdir", "", "")
	cf.Int("parallel", 4, "")
	cf.String("run", "", "")
	cf.Bool("short", false, "")
	cf.Duration("timeout", 10*time.Minute, "")
	cf.String("trace", "", "")
	for name, _ := range passFlagToTest {
		cf.Var(cf.Lookup(name).Value, "test."+name, "")
	}

	err := cf.Parse(args)
	if err != nil {
		os.Exit(2)
		return
	}
	paths := cf.Args()
	if len(paths) == 0 {
		paths = []string{"."}
	}
	path := paths[0]
	var testArgs []string
	cf.VisitAll(func(f *flag.Flag) {
		if strings.HasPrefix(f.Name, "test.") && f.Value.String() != f.DefValue {
			if typ, ok := passFlagToTest[f.Name[5:]]; ok {
				switch typ {
				case String:
					testArgs = append(testArgs, fmt.Sprintf("-%v=%v", f.Name, f.Value))
				default:
					testArgs = append(testArgs, fmt.Sprintf("-%v=%v", f.Name, f.Value))
				}
			}
		}
	})
	var mode igop.Mode
	if base.BuildSSA {
		mode |= igop.EnableDumpInstr
	}
	if base.DebugSSATrace {
		mode |= igop.EnableTracing
	}
	if base.BuildX {
		mode |= igop.EnableDumpImports
	}
	if base.ExperimentalGC {
		mode |= igop.ExperimentalSupportGC
	}
	ctx := igop.NewContext(mode)
	ctx.BuildContext = base.BuildContext

	pkg, err := ctx.LoadDir(path, true)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(2)
	}

	err = ctx.TestPkg(pkg, path, testArgs)
	if err != nil {
		os.Exit(2)
	}
}
