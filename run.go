package interp

import (
	"errors"
	"flag"
	"fmt"
	"go/build"
	"go/parser"
	"go/token"
	"os"
	"time"

	"golang.org/x/tools/go/loader"
	"golang.org/x/tools/go/ssa"
	"golang.org/x/tools/go/ssa/ssautil"
)

func Run(mode Mode, input string, args []string) error {
	pkg, err := Load(input)
	if err != nil {
		return err
	}
	return RunPkg(pkg, mode, input, "main", args)
}

func RunTest(mode Mode, input string, args []string) error {
	pkgpath, pkgs, err := LoadTest(input)
	if err != nil {
		return err
	}
	if len(pkgs) == 0 {
		fmt.Printf("?\t%s [no test files]\n", pkgpath)
		return nil
	}
	RunTestPkg(pkgs, mode, pkgpath, args)
	return nil
}

func RunFile(mode Mode, filename string, src interface{}, args []string) error {
	pkg, err := LoadFile(filename, src)
	if err != nil {
		return err
	}
	return RunPkg(pkg, mode, filename, "main", args)
}

func Load(pkg string) (*ssa.Package, error) {
	cfg := &loader.Config{}
	_, err := cfg.FromArgs([]string{pkg}, false)
	if err != nil {
		return nil, fmt.Errorf("FromArgs(%v) failed: %s", pkg, err)
	}
	return LoadMainPkg(cfg)
}

func LoadTest(input string) (string, []*ssa.Package, error) {
	wd, _ := os.Getwd()
	p, err := build.Import(input, wd, 0)
	if err != nil {
		return "", nil, err
	}
	var paths []string
	if len(p.TestImports) > 0 {
		paths = append(paths, p.ImportPath)
	}
	if len(p.XTestGoFiles) > 0 {
		paths = append(paths, p.ImportPath+"_test")
	}
	cfg := &loader.Config{}
	_, err = cfg.FromArgs([]string{input}, true)
	if err != nil {
		return "", nil, fmt.Errorf("FromArgs(%v) failed: %s", input, err)
	}
	pkgs, err := LoadPkg(cfg, paths)
	return p.ImportPath, pkgs, err
}

func LoadFile(filename string, src interface{}) (*ssa.Package, error) {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, filename, src, parser.AllErrors)
	if err != nil {
		return nil, fmt.Errorf("Parser file error: %v", err)
	}
	cfg := &loader.Config{Fset: fset}
	cfg.CreateFromFiles("", f)
	return LoadMainPkg(cfg)
}

func LoadMainPkg(cfg *loader.Config) (*ssa.Package, error) {
	cfg.Load()
	iprog, err := cfg.Load()
	if err != nil {
		return nil, fmt.Errorf("loader failed: %s", err)
	}

	prog := ssautil.CreateProgram(iprog, ssa.SanityCheckFunctions)
	prog.Build()

	var pkg *ssa.Package
	if len(iprog.Created) > 0 {
		pkg = prog.Package(iprog.Created[0].Pkg)
	} else {
		for _, pkg := range prog.AllPackages() {
			if pkg.Pkg.Name() == "main" {
				pkg = pkg
				break
			}
		}
	}
	if pkg == nil {
		return nil, errors.New("not found main package")
	}
	return pkg, nil
}

func LoadPkg(cfg *loader.Config, paths []string) ([]*ssa.Package, error) {
	cfg.Load()
	iprog, err := cfg.Load()
	if err != nil {
		return nil, fmt.Errorf("loader failed: %s", err)
	}
	prog := ssautil.CreateProgram(iprog, ssa.SanityCheckFunctions)
	prog.Build()
	m := make(map[string]*ssa.Package)
	for _, pkg := range prog.AllPackages() {
		m[pkg.Pkg.Path()] = pkg
	}
	var pkgs []*ssa.Package
	for _, path := range paths {
		if pkg, ok := m[path]; ok {
			pkgs = append(pkgs, pkg)
		}
	}
	return pkgs, nil
}

func RunTestPkg(pkgs []*ssa.Package, mode Mode, input string, args []string) {
	var testPkgs []*ssa.Package
	for _, pkg := range pkgs {
		if p := pkg.Prog.CreateTestMainPackage(pkg); p != nil {
			testPkgs = append(testPkgs, p)
		}
	}
	os.Args = []string{input}
	if args != nil {
		os.Args = append(os.Args, args...)
	}
	var failed bool
	start := time.Now()
	defer func() {
		sec := time.Since(start).Seconds()
		if failed {
			fmt.Printf("FAIL\t%s %0.3fs\n", input, sec)
		} else {
			fmt.Printf("ok\t%s %0.3fs\n", input, sec)
		}
	}()
	if len(testPkgs) == 0 {
		fmt.Println("testing: warning: no tests to run")
	}
	for _, pkg := range testPkgs {
		flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
		if code := Interpret(pkg, mode, "main"); code != 0 {
			failed = true
		}
	}
}

func RunPkg(mainPkg *ssa.Package, mode Mode, input string, entry string, args []string) error {
	// reset os args and flag
	os.Args = []string{input}
	if args != nil {
		os.Args = append(os.Args, args...)
	}
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)

	exitCode := Interpret(mainPkg, mode, entry)
	if exitCode != 0 {
		return fmt.Errorf("interpreting %v: exit code was %d", input, exitCode)
	}
	return nil
}
