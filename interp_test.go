// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gossa_test

// This test runs the SSA interpreter over sample Go programs.
// Because the interpreter requires intrinsics for assembly
// functions and many low-level runtime routines, it is inherently
// not robust to evolutionary change in the standard library.
// Therefore the test cases are restricted to programs that
// use a fake standard library in testdata/src containing a tiny
// subset of simple functions useful for writing assertions.
//
// We no longer attempt to interpret any real standard packages such as
// fmt or testing, as it proved too fragile.

import (
	"bytes"
	"fmt"
	"go/build"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
	"time"

	"github.com/goplus/gossa"
	//_ "github.com/goplus/gossa/pkg"
	_ "github.com/goplus/gossa/pkg/bytes"
	_ "github.com/goplus/gossa/pkg/context"
	_ "github.com/goplus/gossa/pkg/crypto/md5"
	_ "github.com/goplus/gossa/pkg/encoding/binary"
	_ "github.com/goplus/gossa/pkg/errors"
	_ "github.com/goplus/gossa/pkg/flag"
	_ "github.com/goplus/gossa/pkg/fmt"
	_ "github.com/goplus/gossa/pkg/go/ast"
	_ "github.com/goplus/gossa/pkg/io"
	_ "github.com/goplus/gossa/pkg/io/ioutil"
	_ "github.com/goplus/gossa/pkg/log"
	_ "github.com/goplus/gossa/pkg/math"
	_ "github.com/goplus/gossa/pkg/math/rand"
	_ "github.com/goplus/gossa/pkg/os"
	_ "github.com/goplus/gossa/pkg/reflect"
	_ "github.com/goplus/gossa/pkg/runtime"
	_ "github.com/goplus/gossa/pkg/runtime/debug"
	_ "github.com/goplus/gossa/pkg/strconv"
	_ "github.com/goplus/gossa/pkg/strings"
	_ "github.com/goplus/gossa/pkg/sync"
	_ "github.com/goplus/gossa/pkg/sync/atomic"
	_ "github.com/goplus/gossa/pkg/syscall"
	_ "github.com/goplus/gossa/pkg/testing"
	_ "github.com/goplus/gossa/pkg/time"
)

// Each line contains a space-separated list of $GOROOT/test/
// filenames comprising the main package of a program.
// They are ordered quickest-first, roughly.
//
// If a test in this list fails spuriously, remove it.
var gorootTestTests = []string{
	"235.go",
	"alias1.go",
	"func5.go",
	"func6.go",
	"func7.go",
	"func8.go",
	"helloworld.go",
	"varinit.go",
	"escape3.go",
	"initcomma.go",
	"cmp.go", // import OS
	"compos.go",
	"turing.go",
	"indirect.go",
	"complit.go",
	"for.go",
	"struct0.go",
	"intcvt.go",
	"printbig.go",
	"deferprint.go",
	"escape.go",
	"range.go",
	"const4.go",
	"float_lit.go",
	"bigalg.go",
	"decl.go",
	"if.go",
	"named.go",
	"bigmap.go",
	"func.go",
	"reorder2.go",
	"gc.go", // import runtime
	"simassign.go",
	"iota.go",
	"nilptr2.go",
	"utf.go", // import unicode/utf8
	"method.go",
	"char_lit.go", // import os
	//"env.go",         // import runtime;os
	"int_lit.go",     //import os
	"string_lit.go",  //import os
	"defer.go",       //import fmt
	"typeswitch.go",  //import os
	"stringrange.go", //import os fmt unicode/utf8
	"reorder.go",     //import fmt
	"method3.go",
	"literal.go",
	"nul1.go", // doesn't actually assert anything (errorcheckoutput)
	"zerodivide.go",
	"convert.go",
	"convT2X.go",
	"switch.go",
	"ddd.go",
	"blank.go",      // partly disabled //import os
	"closedchan.go", //import os
	"divide.go",     //import fmt
	"rename.go",     //import runtime fmt
	"nil.go",
	"recover1.go",
	"recover2.go",
	//"recover3.go", //TODO fix panic info
	"typeswitch1.go",
	"floatcmp.go",
	"crlf.go", // doesn't actually assert anything (runoutput)
	"append.go",
	"chancap.go",
	"const.go",
	"deferfin.go",
}

// These are files in go.tools/go/ssa/interp/testdata/.
var testdataTests = []string{
	"boundmeth.go",
	"complit.go",
	"coverage.go",
	"defer.go",
	"fieldprom.go",
	"ifaceconv.go",
	"ifaceprom.go",
	"initorder.go",
	"methprom.go",
	"mrvchain.go",
	"range.go",
	"recover.go",
	"reflect.go",
	"static.go",
	"recover2.go",
	"static.go",
	"issue23536.go",
}

var (
	gorootTestSkips = make(map[string]string)
)

func init() {
	if runtime.GOARCH == "386" {
		//		gossa.UnsafeSizes = &types.StdSizes{WordSize: 4, MaxAlign: 4}
		gorootTestSkips["printbig.go"] = "load failed"
		gorootTestSkips["peano.go"] = "stack overflow"
	}
	gorootTestSkips["closure.go"] = "runtime.ReadMemStats"
	gorootTestSkips["divmod.go"] = "timeout"
	gorootTestSkips["copy.go"] = "slow"
	//gorootTestSkips["gcstring.go"] = "timeout"
	gorootTestSkips["finprofiled.go"] = "slow"
	gorootTestSkips["gcgort.go"] = "slow"
	gorootTestSkips["inline_literal.go"] = "TODO, runtime.FuncForPC"
	gorootTestSkips["nilptr.go"] = "skip drawin"
	//gorootTestSkips["recover.go"] = "TODO, fix test16"
	gorootTestSkips["heapsampling.go"] = "runtime.MemProfileRecord"
	gorootTestSkips["makeslice.go"] = "TODO, panic info, allocation size out of range"
	gorootTestSkips["stackobj.go"] = "skip gc"
	gorootTestSkips["stackobj3.go"] = "skip gc"
	gorootTestSkips["nilptr_aix.go"] = "skip"
	gorootTestSkips["init1.go"] = "skip gc"
	gorootTestSkips["string_lit.go"] = "call to os.Exit(0) during test"
	gorootTestSkips["switch.go"] = "call to os.Exit(0) during test"
	gorootTestSkips["ken/divconst.go"] = "slow"
	gorootTestSkips["ken/modconst.go"] = "slow"
	gorootTestSkips["fixedbugs/bug347.go"] = "runtime.Caller"
	gorootTestSkips["fixedbugs/bug348.go"] = "runtime.Caller"
	gorootTestSkips["fixedbugs/issue24491b.go"] = "timeout"
	gorootTestSkips["fixedbugs/issue22781.go"] = "slow"
	gorootTestSkips["fixedbugs/issue16249.go"] = "slow"
	gorootTestSkips["fixedbugs/issue13169.go"] = "slow"
	gorootTestSkips["fixedbugs/bug261.go"] = "BUG, ssa slice[low|high] https://github.com/golang/tools/pull/341"
	gorootTestSkips["fixedbugs/issue11656.go"] = "runtime.Caller"
	gorootTestSkips["fixedbugs/issue15281.go"] = "runtime.ReadMemStats"
	gorootTestSkips["fixedbugs/issue17381.go"] = "runtime.FuncForPC"
	gorootTestSkips["fixedbugs/issue18149.go"] = "runtime.Caller"
	gorootTestSkips["fixedbugs/issue22083.go"] = "debug.Stack"
	gorootTestSkips["fixedbugs/issue22662.go"] = "runtime.Caller"
	gorootTestSkips["fixedbugs/issue27201.go"] = "runtime.Caller"
	gorootTestSkips["fixedbugs/issue27201.go"] = "runtime.Caller"
	gorootTestSkips["fixedbugs/issue27518b.go"] = "BUG, runtime.SetFinalizer"
	gorootTestSkips["fixedbugs/issue29504.go"] = "runtime.Caller"
	gorootTestSkips["fixedbugs/issue32477.go"] = "BUG, runtime.SetFinalizer"
	gorootTestSkips["fixedbugs/issue41239.go"] = "BUG, reflect.Append: different capacity on append"
	gorootTestSkips["fixedbugs/issue32477.go"] = "BUG, runtime.SetFinalizer"
	gorootTestSkips["fixedbugs/issue45175.go"] = "BUG, ssa.Phi call order"
	gorootTestSkips["fixedbugs/issue4562.go"] = "runtime.Caller"
	gorootTestSkips["fixedbugs/issue4618.go"] = "testing.AllocsPerRun"
	gorootTestSkips["fixedbugs/issue4667.go"] = "testing.AllocsPerRun"
	gorootTestSkips["fixedbugs/issue5856.go"] = "runtime.Caller"
	gorootTestSkips["fixedbugs/issue5963.go"] = "BUG, recover"
	gorootTestSkips["fixedbugs/issue7740.go"] = "BUG, const float"
	gorootTestSkips["fixedbugs/issue7690.go"] = "runtime.Stack"
	gorootTestSkips["fixedbugs/issue8606b.go"] = "BUG, optimization check"
	gorootTestSkips["fixedbugs/issue30116u.go"] = "BUG, slice bound check"
	gorootTestSkips["fixedbugs/bug295.go"] = "skip, gossa not import testing"
	gorootTestSkips["fixedbugs/issue27695.go"] = "runtime/debug.SetGCPercent"
	gorootTestSkips["atomicload.go"] = "slow"

	ver := runtime.Version()
	if strings.HasPrefix(ver, "go1.17") || strings.HasPrefix(ver, "go1.18") {
		gorootTestSkips["fixedbugs/issue45045.go"] = "runtime.SetFinalizer"
		gorootTestSkips["fixedbugs/issue46725.go"] = "runtime.SetFinalizer"
		gorootTestSkips["abi/fibish.go"] = "very slow"
		gorootTestSkips["abi/fibish_closure.go"] = "very slow"
		gorootTestSkips["abi/uglyfib.go"] = "very slow"
		gorootTestSkips["fixedbugs/issue23017.go"] = "BUG"
	} else if strings.HasPrefix(ver, "go1.15") {
		gorootTestSkips["fixedbugs/issue15039.go"] = "BUG, uint64 -> string"
		gorootTestSkips["fixedbugs/issue9355.go"] = "TODO, chdir"
	} else if strings.HasPrefix(ver, "go1.14") {
		gorootTestSkips["fixedbugs/issue9355.go"] = "TODO, chdir"
	}

	if runtime.GOOS == "windows" {
		gorootTestSkips["env.go"] = "skip GOARCH"
		gorootTestSkips["fixedbugs/issue15002.go"] = "skip windows"
		gorootTestSkips["fixedbugs/issue5493.go"] = "skip windows"

		skips := make(map[string]string)
		for k, v := range gorootTestSkips {
			skips[filepath.FromSlash(k)] = v
		}
		gorootTestSkips = skips
	} else if runtime.GOOS == "darwin" {
		gorootTestSkips["locklinear.go"] = "skip github"
	}
}

var (
	gossaCmd string
)

func init() {
	var err error
	gossaCmd, err = exec.LookPath("gossa")
	if err != nil {
		panic(fmt.Sprintf("not found gossa: %v", err))
	}
}

func runInput(t *testing.T, input string) bool {
	fmt.Println("Input:", input)
	start := time.Now()
	_, err := gossa.Run(input, nil, 0)
	sec := time.Since(start).Seconds()
	if err != nil {
		t.Error(err)
		fmt.Printf("FAIL %0.3fs\n", sec)
		return false
	}
	fmt.Printf("PASS %0.3fs\n", sec)
	return true
}

func runCommand(t *testing.T, input string) bool {
	fmt.Println("Input:", input)
	start := time.Now()
	cmd := exec.Command(gossaCmd, "run", input)
	data, err := cmd.CombinedOutput()
	if len(data) > 0 {
		fmt.Println(string(data))
	}
	sec := time.Since(start).Seconds()
	if err != nil || bytes.Contains(data, []byte("BUG")) {
		t.Error(err)
		fmt.Printf("FAIL %0.3fs\n", sec)
		return false
	}
	fmt.Printf("PASS %0.3fs\n", sec)
	return true
}

func printFailures(failures []string) {
	if failures != nil {
		fmt.Println("The following tests failed:")
		for _, f := range failures {
			fmt.Printf("\t%s\n", f)
		}
	}
}

// TestTestdataFiles runs the interpreter on testdata/*.go.
func TestTestdataFiles(t *testing.T) {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	var failures []string
	for _, input := range testdataTests {
		if !runInput(t, filepath.Join(cwd, "testdata", input)) {
			failures = append(failures, input)
		}
	}
	printFailures(failures)
}

// $GOROOT/test/*.go runs
func getGorootTestRuns(t *testing.T) (dir string, files []string) {
	dir = filepath.Join(build.Default.GOROOT, "test")
	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			if path == dir {
				return nil
			}
			_, n := filepath.Split(path)
			switch n {
			case "bench", "dwarf", "codegen":
				return filepath.SkipDir
			default:
				if strings.Contains(n, ".dir") {
					return filepath.SkipDir
				}
			}
			return nil
		}
		data, err := ioutil.ReadFile(path)
		if err != nil {
			t.Errorf("read %v error: %v", path, err)
			return nil
		}
		lines := strings.Split(string(data), "\n")
		if len(lines) > 0 {
			line := strings.TrimSpace(lines[0])
			if line == "// run" {
				files = append(files, path)
			}
		}
		return nil
	})
	return
}

// TestGorootTest runs the interpreter on $GOROOT/test/*.go.
func TestGorootTest(t *testing.T) {
	dir, files := getGorootTestRuns(t)
	var failures []string

	for _, input := range files {
		f := input[len(dir)+1:]
		if info, ok := gorootTestSkips[f]; ok {
			fmt.Println("Skip:", input, info)
			continue
		}
		if !runCommand(t, input) {
			failures = append(failures, input)
		}
	}
	printFailures(failures)
}
