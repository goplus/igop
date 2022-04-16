package main

import (
	"bytes"
	"fmt"
	"go/build"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

var (
	gorootTestSkips = make(map[string]string)
)

func init() {
	if runtime.GOARCH == "386" {
		gorootTestSkips["printbig.go"] = "load failed"
		gorootTestSkips["peano.go"] = "stack overflow"
	}
	gorootTestSkips["closure.go"] = "runtime.ReadMemStats"
	gorootTestSkips["divmod.go"] = "slow, 1m49s"
	gorootTestSkips["copy.go"] = "slow, 13s"
	gorootTestSkips["finprofiled.go"] = "slow, 21s"
	gorootTestSkips["gcgort.go"] = "slow, 2s"
	gorootTestSkips["inline_literal.go"] = "TODO, runtime.FuncForPC"
	gorootTestSkips["nilptr.go"] = "skip drawin"
	gorootTestSkips["heapsampling.go"] = "runtime.MemProfileRecord"
	gorootTestSkips["makeslice.go"] = "TODO, panic info, allocation size out of range"
	gorootTestSkips["stackobj.go"] = "skip gc"
	gorootTestSkips["stackobj3.go"] = "skip gc"
	gorootTestSkips["nilptr_aix.go"] = "skip"
	gorootTestSkips["init1.go"] = "skip gc"
	gorootTestSkips["ken/divconst.go"] = "slow, 4s"
	gorootTestSkips["ken/modconst.go"] = "slow, 4s"
	gorootTestSkips["fixedbugs/bug347.go"] = "runtime.Caller"
	gorootTestSkips["fixedbugs/bug348.go"] = "runtime.Caller"
	gorootTestSkips["fixedbugs/issue24491b.go"] = "timeout"
	//gorootTestSkips["fixedbugs/issue22781.go"] = "slow, 1.4s"
	gorootTestSkips["fixedbugs/issue16249.go"] = "slow, 5.8s"
	gorootTestSkips["fixedbugs/issue13169.go"] = "slow, 7.8s"
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
	//gorootTestSkips["fixedbugs/issue5963.go"] = "BUG, recover"
	gorootTestSkips["fixedbugs/issue7690.go"] = "runtime.Stack"
	gorootTestSkips["fixedbugs/issue8606b.go"] = "BUG, optimization check"
	gorootTestSkips["fixedbugs/issue30116u.go"] = "BUG, slice bound check"
	//gorootTestSkips["fixedbugs/bug295.go"] = "skip, gossa not import testing"
	//gorootTestSkips["fixedbugs/issue27695.go"] = "runtime/debug.SetGCPercent"
	//gorootTestSkips["atomicload.go"] = "slow, 0.5"
	gorootTestSkips["chan/select5.go"] = "bug, select case expr call order"

	// fixedbugs/issue7740.go
	// const ulp = (1.0 + (2.0 / 3.0)) - (5.0 / 3.0)
	// Go 1.14 1.15 1.16 ulp = 1.4916681462400413e-154
	// Go 1.17 1.18 ulp = 0

	ver := runtime.Version()[:6]
	switch ver {
	case "go1.17", "go1.18":
		gorootTestSkips["fixedbugs/issue45045.go"] = "runtime.SetFinalizer"
		gorootTestSkips["fixedbugs/issue46725.go"] = "runtime.SetFinalizer"
		gorootTestSkips["abi/fibish.go"] = "slow, 1m4s"
		gorootTestSkips["abi/fibish_closure.go"] = "slow, 1m9s"
		gorootTestSkips["abi/uglyfib.go"] = "timeout"
		gorootTestSkips["fixedbugs/issue23017.go"] = "BUG"
	case "go1.16":
		gorootTestSkips["fixedbugs/issue7740.go"] = "BUG, const float"
	case "go1.15":
		gorootTestSkips["fixedbugs/issue15039.go"] = "BUG, uint64 -> string"
		gorootTestSkips["fixedbugs/issue9355.go"] = "TODO, chdir"
		gorootTestSkips["fixedbugs/issue7740.go"] = "BUG, const float"
	case "go1.14":
		gorootTestSkips["fixedbugs/issue9355.go"] = "TODO, chdir"
		gorootTestSkips["fixedbugs/issue7740.go"] = "BUG, const float"
	}

	if runtime.GOOS == "windows" {
		gorootTestSkips["env.go"] = "skip GOARCH"
		gorootTestSkips["fixedbugs/issue15002.go"] = "skip windows"
		gorootTestSkips["fixedbugs/issue5493.go"] = "skip windows"
		gorootTestSkips["fixedbugs/issue5963.go"] = "skip windows"

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
	goCmd    string
	gossaCmd string
)

func init() {
	var err error
	gossaCmd, err = exec.LookPath("gossa")
	if err != nil {
		panic(fmt.Sprintf("not found gossa: %v", err))
	}
	goCmd, err = exec.LookPath("go")
	if err != nil {
		panic(fmt.Sprintf("not found go: %v", err))
	}
}

func runCommand(input string) bool {
	fmt.Println("Input:", input)
	start := time.Now()
	cmd := exec.Command(gossaCmd, "run", input)
	data, err := cmd.CombinedOutput()
	if len(data) > 0 {
		fmt.Println(string(data))
	}
	sec := time.Since(start).Seconds()
	if err != nil || bytes.Contains(data, []byte("BUG")) {
		fmt.Println(err)
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

// $GOROOT/test/*.go runs
func getGorootTestRuns() (dir string, run []string, runoutput []string) {
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
			fmt.Printf("read %v error: %v\n", path, err)
			return nil
		}
		lines := strings.Split(string(data), "\n")
		if len(lines) > 0 {
			line := strings.TrimSpace(lines[0])
			if line == "// run" {
				run = append(run, path)
			} else if line == "// runoutput" {
				runoutput = append(runoutput, path)
			}
		}
		return nil
	})
	return
}

func execRunoutput(input string) (string, error) {
	cmd := exec.Command(goCmd, "run", input)
	data, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}
	fileName := filepath.Join(os.TempDir(), "tmp.go")
	err = ioutil.WriteFile(fileName, data, 0666)
	if err != nil {
		return "", err
	}
	return fileName, nil
}

func main() {
	dir, runs, runoutput := getGorootTestRuns()
	var failures []string
	start := time.Now()
	for _, input := range runs {
		f := input[len(dir)+1:]
		if info, ok := gorootTestSkips[f]; ok {
			fmt.Println("Skip:", input, info)
			continue
		}
		if !runCommand(input) {
			failures = append(failures, input)
		}
	}
	for _, input := range runoutput {
		f := input[len(dir)+1:]
		if info, ok := gorootTestSkips[f]; ok {
			fmt.Println("Skip:", input, info)
			continue
		}
		fmt.Println("runoutput:", input)
		file, err := execRunoutput(input)
		if err != nil {
			fmt.Printf("go run %v error, %v\n", input, err)
			continue
		}
		if !runCommand(file) {
			failures = append(failures, input)
		}
	}
	printFailures(failures)
	fmt.Printf("time: %.3fs\n", time.Since(start).Seconds())
}
