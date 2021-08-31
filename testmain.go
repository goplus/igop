package interp

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/parser"
	"go/types"
	"log"
	"os"
	"sort"
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
func CreateTestMainPackage(pkg *ssa.Package) *ssa.Package {
	prog := pkg.Prog

	// Template data
	var data struct {
		Pkg                         *ssa.Package
		Tests, Benchmarks, Examples []*ssa.Function
		Main                        *ssa.Function
		Go18                        bool
	}
	data.Pkg = pkg

	// Enumerate tests.
	data.Tests, data.Benchmarks, data.Examples, data.Main = ssa.FindTests(pkg)
	if data.Main == nil &&
		data.Tests == nil && data.Benchmarks == nil && data.Examples == nil {
		return nil
	}
	sort.Slice(data.Tests, func(i, j int) bool {
		return data.Tests[i].Pos() < data.Tests[j].Pos()
	})
	sort.Slice(data.Benchmarks, func(i, j int) bool {
		return data.Benchmarks[i].Pos() < data.Benchmarks[j].Pos()
	})
	sort.Slice(data.Examples, func(i, j int) bool {
		return data.Examples[i].Pos() < data.Examples[j].Pos()
	})

	// Synthesize source for testmain package.
	path := pkg.Pkg.Path() + "$testmain"
	tmpl := testmainTmpl
	if testingPkg := prog.ImportedPackage("testing"); testingPkg != nil {
		// In Go 1.8, testing.MainStart's first argument is an interface, not a func.
		data.Go18 = types.IsInterface(testingPkg.Func("MainStart").Signature.Params().At(0).Type())
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

	// Parse and type-check the testmain package.
	f, err := parser.ParseFile(prog.Fset, path+".go", &buf, parser.Mode(0))
	if err != nil {
		log.Fatalf("internal error parsing %s: %v", path, err)
	}
	conf := types.Config{
		DisableUnusedImportCheck: true,
		Importer:                 testImporter{pkg},
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
	testmainPkg, err := conf.Check(path, prog.Fset, files, info)
	if err != nil {
		log.Fatalf("internal error type-checking %s: %v", path, err)
	}

	// Create and build SSA code.
	testmain := prog.CreatePackage(testmainPkg, files, info, false)
	testmain.SetDebugMode(false)
	testmain.Build()
	testmain.Func("main").Synthetic = "test main function"
	testmain.Func("init").Synthetic = "package initializer"
	return testmain
}

// An implementation of types.Importer for an already loaded SSA program.
type testImporter struct {
	pkg *ssa.Package // package under test; may be non-importable
}

func (imp testImporter) Import(path string) (*types.Package, error) {
	if p := imp.pkg.Prog.ImportedPackage(path); p != nil {
		return p.Pkg, nil
	}
	if path == imp.pkg.Pkg.Path() {
		return imp.pkg.Pkg, nil
	}
	return nil, fmt.Errorf("not found") // can't happen
}

var testmainTmpl = template.Must(template.New("testmain").Parse(`
package main

import "io"
import "os"
import "testing"
import p {{printf "%q" .Pkg.Pkg.Path}}

{{if .Go18}}
type deps struct{}

func (deps) ImportPath() string { return "" }
func (deps) MatchString(pat, str string) (bool, error) { return true, nil }
func (deps) SetPanicOnExit0(bool) {}
func (deps) StartCPUProfile(io.Writer) error { return nil }
func (deps) StartTestLog(io.Writer) {}
func (deps) StopCPUProfile() {}
func (deps) StopTestLog() error { return nil }
func (deps) WriteHeapProfile(io.Writer) error { return nil }
func (deps) WriteProfileTo(string, io.Writer, int) error { return nil }

var match deps
{{else}}
func match(_, _ string) (bool, error) { return true, nil }
{{end}}

func main() {
	tests := []testing.InternalTest{
{{range .Tests}}
		{ {{printf "%q" .Name}}, p.{{.Name}} },
{{end}}
	}
	benchmarks := []testing.InternalBenchmark{
{{range .Benchmarks}}
		{ {{printf "%q" .Name}}, p.{{.Name}} },
{{end}}
	}
	examples := []testing.InternalExample{
{{range .Examples}}
		{Name: {{printf "%q" .Name}}, F: p.{{.Name}}},
{{end}}
	}
	m := testing.MainStart(match, tests, benchmarks, examples)
{{with .Main}}
	p.{{.Name}}(m)
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
