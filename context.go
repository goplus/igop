package gossa

import (
	"flag"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"go/types"
	"os"
	"time"

	"golang.org/x/tools/go/ssa"
)

var (
	Default = NewContext()
)

type Context struct {
	Loader      Loader
	Mode        Mode
	ParserMode  parser.Mode
	BuilderMode ssa.BuilderMode
}

func NewContext() *Context {
	ctx := &Context{
		Loader:      NewTypesLoader(),
		ParserMode:  parser.AllErrors,
		BuilderMode: ssa.SanityCheckFunctions,
	}
	return ctx
}

func (c *Context) LoadDir(path string) (pkgs []*ssa.Package, first error) {
	fset := token.NewFileSet()
	apkgs, err := parser.ParseDir(fset, path, nil, c.ParserMode)
	if err != nil {
		return nil, err
	}
	for _, apkg := range apkgs {
		if pkg, err := c.LoadAstPackage(fset, apkg); err == nil {
			pkgs = append(pkgs, pkg)
		} else if first == nil {
			first = err
		}
	}
	return
}

func (c *Context) LoadFile(filename string, src interface{}) (*ssa.Package, error) {
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, filename, src, c.ParserMode)
	if err != nil {
		return nil, err
	}
	return c.LoadAstFile(fset, file)
}

func (c *Context) LoadAstFile(fset *token.FileSet, file *ast.File) (*ssa.Package, error) {
	pkg := types.NewPackage(file.Name.Name, "")
	ssapkg, _, err := BuildPackage(c.Loader, fset, pkg, []*ast.File{file}, c.BuilderMode)
	if err != nil {
		return nil, err
	}
	ssapkg.Build()
	return ssapkg, nil
}

func (c *Context) LoadAstPackage(fset *token.FileSet, apkg *ast.Package) (*ssa.Package, error) {
	pkg := types.NewPackage(apkg.Name, "")
	var files []*ast.File
	for _, f := range apkg.Files {
		files = append(files, f)
	}
	ssapkg, _, err := BuildPackage(c.Loader, fset, pkg, files, c.BuilderMode)
	if err != nil {
		return nil, err
	}
	ssapkg.Build()
	return ssapkg, nil
}

func (c *Context) RunPkg(mainPkg *ssa.Package, input string, entry string, args []string) error {
	// reset os args and flag
	os.Args = []string{input}
	if args != nil {
		os.Args = append(os.Args, args...)
	}
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)

	interp := NewInterp(c.Loader, mainPkg, c.Mode)
	interp.Run("init")
	_, exitCode := interp.Run(entry)
	if exitCode != 0 {
		return fmt.Errorf("interpreting %v: exit code was %d", input, exitCode)
	}
	return nil
}

func (c *Context) TestPkg(pkgs []*ssa.Package, input string, args []string) {
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
	var testPkgs []*ssa.Package
	for _, pkg := range pkgs {
		if p := CreateTestMainPackage(pkg); p != nil {
			testPkgs = append(testPkgs, p)
		}
	}
	os.Args = []string{input}
	if args != nil {
		os.Args = append(os.Args, args...)
	}
	if len(testPkgs) == 0 {
		fmt.Println("testing: warning: no tests to run")
	}
	for _, pkg := range testPkgs {
		flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
		interp := NewInterp(c.Loader, pkg, c.Mode)
		interp.Run("init")
		_, exitCode := interp.Run("main")
		if exitCode != 0 {
			failed = true
		}
	}
}
