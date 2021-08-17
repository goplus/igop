// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package interp_test

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
	"errors"
	"fmt"
	"go/build"
	"go/types"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"strings"
	"testing"
	"time"

	"github.com/goplus/interp"
	_ "github.com/goplus/interp/pkg"
	// _ "github.com/goplus/interp/pkg/errors"
	// _ "github.com/goplus/interp/pkg/fmt"
	// _ "github.com/goplus/interp/pkg/math"
	// _ "github.com/goplus/interp/pkg/os"
	// _ "github.com/goplus/interp/pkg/reflect"
	// _ "github.com/goplus/interp/pkg/runtime"
	// _ "github.com/goplus/interp/pkg/strings"
	// _ "github.com/goplus/interp/pkg/sync"
	// _ "github.com/goplus/interp/pkg/time"
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
	"coverage.go", //TODO fix panic info
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
}

var (
	gorootTestSkips = make(map[string]bool)
)

func init() {
	if runtime.GOARCH == "386" {
		interp.UnsafeSizes = &types.StdSizes{WordSize: 4, MaxAlign: 4}
		gorootTestSkips["printbig.go"] = true // load failed
	}
	gorootTestSkips["closure.go"] = true        // runtime.ReadMemStats
	gorootTestSkips["divmod.go"] = true         // timeout
	gorootTestSkips["copy.go"] = true           // slow
	gorootTestSkips["gcstring.go"] = true       // timeout
	gorootTestSkips["finprofiled.go"] = true    // slow
	gorootTestSkips["gcgort.go"] = true         // slow
	gorootTestSkips["inline_literal.go"] = true // bug, runtime.FuncForPC
	gorootTestSkips["nilptr.go"] = true         // skip drawin
}

func _init() {
	interp.RegisterExternal("runtime.init", func() {})
	interp.RegisterExternal("runtime.GC", runtime.GC)
	interp.RegisterExternal("os.init", func() {})
	interp.RegisterExternal("os.Getenv", os.Getenv)
	interp.RegisterExternal("fmt.init", func() {})
	interp.RegisterExternal("fmt.Print", fmt.Print)
	interp.RegisterExternal("fmt.Printf", fmt.Printf)
	interp.RegisterExternal("fmt.Println", fmt.Println)
	interp.RegisterExternal("fmt.Sprint", fmt.Sprint)
	interp.RegisterExternal("math.init", func() {})
	interp.RegisterExternal("strings.init", func() {})
	interp.RegisterExternal("strings.IndexByte", strings.IndexByte)
	interp.RegisterExternal("strings.Contains", strings.Contains)
	interp.RegisterExternal("reflect.init", func() {})
	interp.RegisterExternal("reflect.TypeOf", reflect.TypeOf)
	interp.RegisterExternal("reflect.SliceOf", reflect.SliceOf)
	interp.RegisterExternal("time.init", func() {})
	interp.RegisterExternal("time.Sleep", time.Sleep)
	interp.RegisterTypeOf((*time.Duration)(nil))
	interp.RegisterExternal("errors.init", func() {})
	interp.RegisterExternal("errors.New", errors.New)
}

func run(t *testing.T, input string) bool {
	fmt.Println("Input:", input)
	err := interp.Run(0, input, nil)
	if err != nil {
		t.Error(err)
		fmt.Println("FAIL")
		return false
	}
	fmt.Println("PASS")
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
		if !run(t, filepath.Join(cwd, "testdata", input)) {
			failures = append(failures, input)
		}
	}
	printFailures(failures)
}

// TestGorootTest runs the interpreter on $GOROOT/test/*.go.
func TestGorootTest(t *testing.T) {
	var failures []string

	for _, input := range gorootTestTests {
		if gorootTestSkips[input] {
			continue
		}
		if !run(t, filepath.Join(build.Default.GOROOT, "test", input)) {
			failures = append(failures, input)
		}
	}
	printFailures(failures)
}

// $GOROOT/test/*.go runs
func getGorootTestRuns(t *testing.T) (files []string) {
	dir := filepath.Join(build.Default.GOROOT, "test")
	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			if path == dir {
				return nil
			}
			return filepath.SkipDir
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
func _TestGorootTest(t *testing.T) {
	files := getGorootTestRuns(t)
	var failures []string

	for _, input := range files {
		if _, f := filepath.Split(input); gorootTestSkips[f] {
			continue
		}
		if !run(t, input) {
			failures = append(failures, input)
		}
	}
	printFailures(failures)
}
