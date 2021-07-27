package interp

import (
	"fmt"
	"go/build"
	"go/parser"
	"go/token"
	"go/types"

	"golang.org/x/tools/go/loader"
	"golang.org/x/tools/go/ssa"
	"golang.org/x/tools/go/ssa/ssautil"
)

type Config struct {
	Build    *build.Context
	Mode     Mode
	Input    []string
	Source   interface{}
	WithTest bool
	Sizes    *types.StdSizes
	Conf     *loader.Config
}

func RunWith(mode Mode, input ...string) error {
	return Run(&Config{Mode: mode, Input: input})
}

func RunSource(mode Mode, source interface{}) error {
	return Run(&Config{Mode: mode, Source: source})
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

	if cfg.Sizes == nil {
		cfg.Sizes = &types.StdSizes{WordSize: 8, MaxAlign: 8}
	}

	iprog, err := cfg.Conf.Load()
	if err != nil {
		return fmt.Errorf("conf.Load failed: %s", err)
	}

	prog := ssautil.CreateProgram(iprog, ssa.SanityCheckFunctions)
	prog.Build()

	mainPkg := prog.Package(iprog.Created[0].Pkg)
	if mainPkg == nil {
		return fmt.Errorf("not a main package: %s", cfg.Input)
	}

	exitCode := Interpret(mainPkg, cfg.Mode, cfg.Sizes, "", []string{})
	if exitCode != 0 {
		return fmt.Errorf("interpreting %v: exit code was %d", cfg.Input, exitCode)
	}
	return nil
}
