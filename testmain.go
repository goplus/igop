package igop

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/parser"
	"go/types"
	"log"
	"os"
	"sort"
	"strings"
	"text/template"

	"golang.org/x/tools/go/ssa"
)

// CreateTestMainPackage creates and returns a synthetic "testmain"
// package for the specified package if it defines tests, benchmarks or
// executable examples, or nil otherwise.  The new package is named
// "main" and provides a function named "main" that runs the tests,
// similar to the one that would be created by the 'go test' tool.
//
// Subsequent calls to prog.AllPackages include the new package.
// The package pkg must belong to the program prog.
//
// Deprecated: Use golang.org/x/tools/go/packages to access synthetic
// testmain packages.
func CreateTestMainPackage(ctx *Context, pkg *ssa.Package, xtestpkg *ssa.Package) (*ssa.Package, error) {
	prog := pkg.Prog

	// Template data
	var data struct {
		TestPkg                                    *ssa.Package
		XTestPkg                                   *ssa.Package
		Tests, Benchmarks, Examples                []*ssa.Function
		FuzzTargets                                []*ssa.Function
		XTestTests, XTestBenchmarks, XTestExamples []*ssa.Function
		XTestFuzzTargets                           []*ssa.Function
		Main                                       *ssa.Function
		Go118                                      bool
	}
	sortFunc := func(funcs []*ssa.Function) {
		sort.Slice(funcs, func(i, j int) bool {
			return funcs[i].Pos() < funcs[j].Pos()
		})
	}
	// Enumerate tests.
	data.Tests, data.Benchmarks, data.Examples, data.Main = FindTests(pkg)
	sortFunc(data.Tests)
	sortFunc(data.Benchmarks)
	sortFunc(data.Examples)
	if len(data.Tests) != 0 || len(data.Benchmarks) != 0 || len(data.Examples) != 0 {
		data.TestPkg = pkg
	}
	if xtestpkg != nil {
		data.XTestTests, data.XTestBenchmarks, data.XTestExamples, _ = FindTests(xtestpkg)
		sortFunc(data.XTestTests)
		sortFunc(data.XTestBenchmarks)
		sortFunc(data.XTestExamples)
		if len(data.XTestTests) != 0 || len(data.XTestBenchmarks) != 0 || len(data.XTestExamples) != 0 {
			data.XTestPkg = xtestpkg
		}
	}

	// Synthesize source for testmain package.
	path := pkg.Pkg.Path() + "$testmain"
	tmpl := testmainTmpl
	if testingPkg := prog.ImportedPackage("testing"); testingPkg != nil {
		if testingPkg.Type("InternalFuzzTarget") != nil {
			data.Go118 = true
		}
		// In Go 1.8, testing.MainStart's first argument is an interface, not a func.
		// data.Go18 = types.IsInterface(testingPkg.Func("MainStart").Signature.Params().At(0).Type())
	} else {
		// The program does not import "testing", but FindTests
		// returned non-nil, which must mean there were Examples
		// but no Test, Benchmark, or TestMain functions.

		// We'll simply call them from testmain.main; this will
		// ensure they don't panic, but will not check any
		// "Output:" comments.
		// (We should not execute an Example that has no
		// "Output:" comment, but it's impossible to tell here.)
		tmpl = examplesOnlyTmpl
	}
	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		log.Fatalf("internal error expanding template for %s: %v", path, err)
	}

	if false { // debugging
		fmt.Fprintln(os.Stderr, buf.String())
	}

	// check import regexp
	if prog.ImportedPackage("regexp") == nil {
		rpkg, err := ctx.Loader.Import("regexp")
		if err != nil {
			return nil, err
		}
		prog.CreatePackage(rpkg, nil, nil, true)
	}

	// Parse and type-check the testmain package.
	f, err := parser.ParseFile(prog.Fset, path+".go", &buf, parser.Mode(0))
	if err != nil {
		return nil, err
	}

	files := []*ast.File{f}
	info := &types.Info{
		Types:      make(map[ast.Expr]types.TypeAndValue),
		Defs:       make(map[*ast.Ident]types.Object),
		Uses:       make(map[*ast.Ident]types.Object),
		Implicits:  make(map[ast.Node]types.Object),
		Scopes:     make(map[ast.Node]*types.Scope),
		Selections: make(map[*ast.SelectorExpr]*types.Selection),
	}
	conf := types.Config{
		Importer: testImporter{ctx, pkg.Prog},
	}
	testmainPkg, err := conf.Check(path, prog.Fset, files, info)
	if err != nil {
		return nil, err
	}

	// Create and build SSA code.
	testmain := prog.CreatePackage(testmainPkg, files, info, false)
	testmain.SetDebugMode(false)
	testmain.Build()
	testmain.Func("main").Synthetic = "test main function"
	testmain.Func("init").Synthetic = "package initializer"
	return testmain, nil
}

// An implementation of types.Importer for an already loaded SSA program.
type testImporter struct {
	ctx  *Context
	prog *ssa.Program
}

func (imp testImporter) Import(path string) (*types.Package, error) {
	if p := imp.prog.ImportedPackage(path); p != nil {
		return p.Pkg, nil
	}
	return imp.ctx.Loader.Import(path)
}

var testmainTmpl = template.Must(template.New("testmain").Parse(`
package main

import (
	"io"
	"os"
	"regexp"
	"testing"
	{{if .Go118}}
	"time"
	"reflect"
	{{end}}
	{{if .TestPkg}}
	_test {{printf "%q" .TestPkg.Pkg.Path}}
	{{end}}
	{{if .XTestPkg}}
	_xtest {{printf "%q" .XTestPkg.Pkg.Path}}
	{{end}}
)

type deps struct{}

func (deps) ImportPath() string { return "" }
func (deps) SetPanicOnExit0(bool) {}
func (deps) StartCPUProfile(io.Writer) error { return nil }
func (deps) StartTestLog(io.Writer) {}
func (deps) StopCPUProfile() {}
func (deps) StopTestLog() error { return nil }
func (deps) WriteHeapProfile(io.Writer) error { return nil }
func (deps) WriteProfileTo(string, io.Writer, int) error { return nil }

{{if .Go118}}
type corpusEntry = struct {
	Parent     string
	Path       string
	Data       []byte
	Values     []any
	Generation int
	IsSeed     bool
}

func (deps) CoordinateFuzzing(time.Duration, int64, time.Duration, int64, int, []corpusEntry, []reflect.Type, string, string) error {
	return nil
}
func (deps) RunFuzzWorker(func(corpusEntry) error) error { return nil }
func (deps) ReadCorpus(string, []reflect.Type) ([]corpusEntry, error) {
	return nil, nil
}
func (f deps) CheckCorpus([]any, []reflect.Type) error { return nil }
func (f deps) ResetCoverage() {}
func (f deps) SnapshotCoverage() {}
{{end}}

var matchPat string
var matchRe *regexp.Regexp

func (deps) MatchString(pat, str string) (result bool, err error) {
	if matchRe == nil || matchPat != pat {
		matchPat = pat
		matchRe, err = regexp.Compile(matchPat)
		if err != nil {
			return
		}
	}
	return matchRe.MatchString(str), nil
}

var tests = []testing.InternalTest{
{{range .Tests}}
	{ {{printf "%q" .Name}}, _test.{{.Name}} },
{{end}}
{{if .XTestPkg}}
{{range .XTestTests}}
	{ {{printf "%q" .Name}}, _xtest.{{.Name}} },
{{end}}
{{end}}
}

var benchmarks = []testing.InternalBenchmark{
{{range .Benchmarks}}
	{ {{printf "%q" .Name}}, _test.{{.Name}} },
{{end}}
{{if .XTestPkg}}
{{range .XTestBenchmarks}}
	{ {{printf "%q" .Name}}, _xtest.{{.Name}} },
{{end}}
{{end}}
}

var examples = []testing.InternalExample{
{{range .Examples}}
	{Name: {{printf "%q" .Name}}, F: _test.{{.Name}}},
{{end}}
{{if .XTestPkg}}
{{range .XTestExamples}}
	{Name: {{printf "%q" .Name}}, F: _xtest.{{.Name}}},
{{end}}
{{end}}
}

{{if .Go118}}
var fuzzTargets = []testing.InternalFuzzTarget{
{{range .FuzzTargets}}
	{ {{printf "%q" .Name}}, _test.{{.Name}} },
{{end}}
{{if .XTestPkg}}
{{range .XTestFuzzTargets}}
	{ {{printf "%q" .Name}}, _xtest.{{.Name}} },
{{end}}
{{end}}
}
{{end}}

func main() {
{{if .Go118}}
	m := testing.MainStart(deps{}, tests, benchmarks, fuzzTargets, examples)
{{else}}
	m := testing.MainStart(deps{}, tests, benchmarks, examples)
{{end}}
{{with .Main}}
	_test.{{.Name}}(m)
{{else}}
	os.Exit(m.Run())
{{end}}
}

`))

var examplesOnlyTmpl = template.Must(template.New("examples").Parse(`
package main

import p {{printf "%q" .Pkg.Pkg.Path}}

func main() {
{{range .Examples}}
	p.{{.Name}}()
{{end}}
}
`))

// FindTests returns the Test, Benchmark, and Example functions
// (as defined by "go test") defined in the specified package,
// and its TestMain function, if any.
//
// Deprecated: Use golang.org/x/tools/go/packages to access synthetic
// testmain packages.
func FindTests(pkg *ssa.Package) (tests, benchmarks, examples []*ssa.Function, main *ssa.Function) {
	prog := pkg.Prog

	// The first two of these may be nil: if the program doesn't import "testing",
	// it can't contain any tests, but it may yet contain Examples.
	var testSig *types.Signature                              // func(*testing.T)
	var benchmarkSig *types.Signature                         // func(*testing.B)
	var exampleSig = types.NewSignature(nil, nil, nil, false) // func()

	// Obtain the types from the parameters of testing.MainStart.
	if testingPkg := prog.ImportedPackage("testing"); testingPkg != nil {
		mainStart := testingPkg.Func("MainStart")
		params := mainStart.Signature.Params()
		testSig = funcField(params.At(1).Type())
		benchmarkSig = funcField(params.At(2).Type())

		// Does the package define this function?
		//   func TestMain(*testing.M)
		if f := pkg.Func("TestMain"); f != nil {
			sig := f.Type().(*types.Signature)
			starM := mainStart.Signature.Results().At(0).Type() // *testing.M
			if sig.Results().Len() == 0 &&
				sig.Params().Len() == 1 &&
				types.Identical(sig.Params().At(0).Type(), starM) {
				main = f
			}
		}
	}

	// TODO(adonovan): use a stable order, e.g. lexical.
	for _, mem := range pkg.Members {
		if f, ok := mem.(*ssa.Function); ok &&
			ast.IsExported(f.Name()) &&
			strings.HasSuffix(prog.Fset.Position(f.Pos()).Filename, "_test.go") {

			switch {
			case testSig != nil && isTestSig(f, "Test", testSig):
				tests = append(tests, f)
			case benchmarkSig != nil && isTestSig(f, "Benchmark", benchmarkSig):
				benchmarks = append(benchmarks, f)
			case isTestSig(f, "Example", exampleSig):
				examples = append(examples, f)
			default:
				continue
			}
		}
	}
	return
}

// Like isTest, but checks the signature too.
func isTestSig(f *ssa.Function, prefix string, sig *types.Signature) bool {
	return isTest(f.Name(), prefix) && types.Identical(f.Signature, sig)
}

// Given the type of one of the three slice parameters of testing.Main,
// returns the function type.
func funcField(slice types.Type) *types.Signature {
	return slice.(*types.Slice).Elem().Underlying().(*types.Struct).Field(1).Type().(*types.Signature)
}

// isTest tells whether name looks like a test (or benchmark, according to prefix).
// It is a Test (say) if there is a character after Test that is not a lower-case letter.
// We don't want TesticularCancer.
// Plundered from $GOROOT/src/cmd/go/test.go
func isTest(name, prefix string) bool {
	if !strings.HasPrefix(name, prefix) {
		return false
	}
	if len(name) == len(prefix) { // "Test" is ok
		return true
	}
	return ast.IsExported(name[len(prefix):])
}
