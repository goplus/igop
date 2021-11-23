package gossa

import (
	"errors"
	"flag"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"go/types"
	"os"
	"reflect"
	"strings"
	"time"

	"github.com/goplus/reflectx"

	"golang.org/x/tools/go/ssa"
	"golang.org/x/tools/go/ssa/ssautil"
)

var (
	ErrNoPackage        = errors.New("no package")
	ErrPackage          = errors.New("package contain errors")
	ErrNotFoundMain     = errors.New("not found main package")
	ErrTestFailed       = errors.New("test failed")
	ErrNotFoundImporter = errors.New("not found provider for types.Importer")
)

// types loader interface
type Loader interface {
	InstallPackage(path string) (*types.Package, error)
	Packages() []*types.Package
	LookupPackage(pkgpath string) (*types.Package, bool)
	LookupReflect(typ types.Type) (reflect.Type, bool)
	LookupTypes(typ reflect.Type) (types.Type, bool)
}

type Context struct {
	Loader      Loader
	Mode        Mode
	ParserMode  parser.Mode
	BuilderMode ssa.BuilderMode
	Importer    types.Importer
}

func NewContext(mode Mode) *Context {
	ctx := &Context{
		Loader:      NewTypesLoader(),
		Mode:        mode,
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
	ssapkg, _, err := BuildPackage(c.Loader, c.Importer, fset, pkg, []*ast.File{file}, c.BuilderMode)
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
	ssapkg, _, err := BuildPackage(c.Loader, c.Importer, fset, pkg, files, c.BuilderMode)
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

func (c *Context) TestPkg(pkgs []*ssa.Package, input string, args []string) error {
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
	if failed {
		return ErrTestFailed
	}
	return nil
}

func (c *Context) RunFile(filename string, src interface{}, args []string) error {
	pkg, err := c.LoadFile(filename, src)
	if err != nil {
		return err
	}
	return c.RunPkg(pkg, filename, "main", args)
}

func (c *Context) Run(path string, args []string, mode Mode) error {
	if strings.HasSuffix(path, ".go") {
		return c.RunFile(path, nil, args)
	}
	pkgs, err := c.LoadDir(path)
	if err != nil {
		return err
	}
	mainPkgs := ssautil.MainPackages(pkgs)
	if len(mainPkgs) == 0 {
		return ErrNotFoundMain
	}
	return c.RunPkg(mainPkgs[0], path, "main", args)
}

func (c *Context) RunTest(path string, args []string) error {
	pkgs, err := c.LoadDir(path)
	if err != nil {
		return err
	}
	return c.TestPkg(pkgs, path, args)
}

func RunFile(filename string, src interface{}, args []string, mode Mode) error {
	reflectx.Reset()
	ctx := NewContext(mode)
	return ctx.RunFile(filename, src, args)
}

func Run(path string, args []string, mode Mode) error {
	reflectx.Reset()
	ctx := NewContext(mode)
	return ctx.Run(path, args, mode)
}

func RunTest(path string, args []string, mode Mode) error {
	reflectx.Reset()
	ctx := NewContext(mode)
	return ctx.RunTest(path, args)
}
