package interp

import (
	"fmt"
	"go/build"
	"go/parser"
	"go/token"
	"os"

	"golang.org/x/tools/go/loader"
	"golang.org/x/tools/go/ssa"
	"golang.org/x/tools/go/ssa/ssautil"
)

type Config struct {
	Build    *build.Context
	Mode     Mode
	Entry    string
	Input    []string
	Args     []string
	Source   interface{}
	WithTest bool
	Conf     *loader.Config
}

func RunWith(mode Mode, input string, args []string) error {
	return Run(&Config{Mode: mode, Input: []string{input}, Args: args})
}

func RunWithTest(mode Mode, input string, args []string) error {
	return Run(&Config{Mode: mode, Input: []string{input}, Args: args, WithTest: true})
}

func RunSource(mode Mode, source interface{}, args []string) error {
	return Run(&Config{Mode: mode, Source: source, Input: []string{"source"}, Args: args})
}

func Run(cfg *Config) error {
	if cfg.Conf == nil {
		if cfg.Build == nil {
			cfg.Build = &build.Default
		}
		cfg.Conf = &loader.Config{Build: cfg.Build}
	}
	if len(cfg.Conf.CreatePkgs) == 0 {
		if cfg.Source != nil {
			if cfg.Conf.Fset == nil {
				cfg.Conf.Fset = token.NewFileSet()
			}
			f, err := parser.ParseFile(cfg.Conf.Fset, "main.go", cfg.Source, parser.AllErrors)
			if err != nil {
				return fmt.Errorf("Parser File error: %v", err)
			}
			cfg.Conf.CreateFromFiles("", f)
		} else if _, err := cfg.Conf.FromArgs(cfg.Input, cfg.WithTest); err != nil {
			return fmt.Errorf("FromArgs(%v) failed: %s", cfg.Input, err)
		}
	}

	if cfg.Entry == "" {
		cfg.Entry = "main"
	}

	iprog, err := cfg.Conf.Load()
	if err != nil {
		return fmt.Errorf("conf.Load failed: %s", err)
	}

	prog := ssautil.CreateProgram(iprog, ssa.SanityCheckFunctions)
	prog.Build()

	var mainPkg *ssa.Package
	if len(iprog.Created) > 0 {
		mainPkg = prog.Package(iprog.Created[0].Pkg)
	} else {
		for _, pkg := range prog.AllPackages() {
			if pkg.Pkg.Name() == "main" {
				mainPkg = pkg
				break
			}
		}
	}
	if mainPkg == nil {
		return fmt.Errorf("not a main package: %s", cfg.Input)
	}

	if cfg.WithTest {
		mainPkg = prog.CreateTestMainPackage(mainPkg)
		if mainPkg == nil {
			return fmt.Errorf("%s [no test files]", cfg.Input)
		}
	}
	os.Args = []string{cfg.Input[0]}
	if cfg.Args != nil {
		os.Args = append(os.Args, cfg.Args...)
	}

	exitCode := Interpret(mainPkg, cfg.Mode, cfg.Entry)
	if exitCode != 0 {
		return fmt.Errorf("interpreting %v: exit code was %d", cfg.Input, exitCode)
	}
	return nil
}
