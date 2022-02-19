package gossa

import (
	"flag"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"go/types"
	"log"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"time"

	"github.com/goplus/reflectx"

	"golang.org/x/tools/go/ssa"
	"golang.org/x/tools/go/ssa/ssautil"
)

func init() {
	RegisterExternal("os.Exit", func(code int) {
		panic(exitPanic(code))
	})
}

// Mode is a bitmask of options affecting the interpreter.
type Mode uint

const (
	DisableRecover         Mode = 1 << iota // Disable recover() in target programs; show interpreter crash instead.
	DisableUnexportMethods                  // Disable unexport methods
	EnableTracing                           // Print a trace of all instructions as they are interpreted.
	EnableDumpPackage                       // Print package
	EnableDumpInstr                         // Print instr type & value
)

// types loader interface
type Loader interface {
	Import(path string) (*types.Package, error)
	Installed(path string) (*Package, bool)
	Packages() []*types.Package
	LookupReflect(typ types.Type) (reflect.Type, bool)
	LookupTypes(typ reflect.Type) (types.Type, bool)
}

type Context struct {
	Loader      Loader          // types loader
	Mode        Mode            // mode
	ParserMode  parser.Mode     // parser mode
	BuilderMode ssa.BuilderMode // ssa builder mode
	External    types.Importer  // external import
	Sizes       types.Sizes
	DebugFunc   func(*DebugInfo)
	Types       map[types.Type]bool
}

func NewContext(mode Mode) *Context {
	ctx := &Context{
		Loader:      NewTypesLoader(mode),
		Mode:        mode,
		ParserMode:  parser.AllErrors,
		BuilderMode: 0, //ssa.SanityCheckFunctions,
		Types:       make(map[types.Type]bool),
	}
	ctx.Types[typesEmptyInterface] = true
	ctx.Types[typesError] = true
	return ctx
}

func (c *Context) SetDebug(fn func(*DebugInfo)) {
	c.BuilderMode |= ssa.GlobalDebug
	c.DebugFunc = fn
}

func (c *Context) LoadDir(fset *token.FileSet, path string) (pkgs []*ssa.Package, first error) {
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

func RegisterFileProcess(ext string, fn SourceProcessFunc) {
	sourceProcesser[ext] = fn
}

type SourceProcessFunc func(ctx *Context, filename string, src interface{}) ([]byte, error)

var (
	sourceProcesser = make(map[string]SourceProcessFunc)
)

func (c *Context) LoadFile(fset *token.FileSet, filename string, src interface{}) (*ssa.Package, error) {
	if ext := filepath.Ext(filename); ext != "" {
		if fn, ok := sourceProcesser[ext]; ok {
			data, err := fn(c, filename, src)
			if err != nil {
				return nil, err
			}
			src = data
		}
	}
	file, err := parser.ParseFile(fset, filename, src, c.ParserMode)
	if err != nil {
		return nil, err
	}
	return c.LoadAstFile(fset, file)
}

func (c *Context) LoadAstFile(fset *token.FileSet, file *ast.File) (*ssa.Package, error) {
	pkg := types.NewPackage(file.Name.Name, "")
	ssapkg, _, err := c.BuildPackage(fset, pkg, []*ast.File{file})
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
	ssapkg, _, err := c.BuildPackage(fset, pkg, files)
	if err != nil {
		return nil, err
	}
	ssapkg.Build()
	return ssapkg, nil
}

func (c *Context) RunPkg(mainPkg *ssa.Package, input string, args []string) (exitCode int, err error) {
	// reset os args and flag
	os.Args = []string{input}
	if args != nil {
		os.Args = append(os.Args, args...)
	}
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)

	interp, err := c.NewInterp(mainPkg)
	if err != nil {
		return 2, err
	}
	return interp.Run("main")
}

func (c *Context) RunFunc(mainPkg *ssa.Package, fnname string, args ...Value) (ret Value, err error) {
	interp, err := c.NewInterp(mainPkg)
	if err != nil {
		return nil, err
	}
	return interp.RunFunc(fnname, args...)
}

func (c *Context) NewInterp(mainPkg *ssa.Package) (*Interp, error) {
	r, err := NewInterp(c.Loader, mainPkg, c.Mode)
	if err == nil {
		r.PreloadTypes(c.Types)
		r.setDebug(c.DebugFunc)
	}
	return r, err
}

func (c *Context) TestPkg(pkgs []*ssa.Package, input string, args []string) error {
	var failed bool
	start := time.Now()
	var testPkgs []*ssa.Package
	for _, pkg := range pkgs {
		p, err := CreateTestMainPackage(pkg)
		if err != nil {
			return err
		}
		if p != nil {
			testPkgs = append(testPkgs, p)
		}
	}
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
	os.Args = []string{input}
	if args != nil {
		os.Args = append(os.Args, args...)
	}
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	for _, pkg := range testPkgs {
		interp, _ := NewInterp(c.Loader, pkg, c.Mode)
		exitCode, _ := interp.Run("main")
		if exitCode != 0 {
			failed = true
		}
	}
	if failed {
		return ErrTestFailed
	}
	return nil
}

func (c *Context) RunFile(filename string, src interface{}, args []string) (exitCode int, err error) {
	fset := token.NewFileSet()
	pkg, err := c.LoadFile(fset, filename, src)
	if err != nil {
		return 2, err
	}
	return c.RunPkg(pkg, filename, args)
}

func (c *Context) Run(path string, args []string) (exitCode int, err error) {
	if strings.HasSuffix(path, ".go") {
		return c.RunFile(path, nil, args)
	}
	fset := token.NewFileSet()
	pkgs, err := c.LoadDir(fset, path)
	if err != nil {
		return 2, err
	}
	mainPkgs := ssautil.MainPackages(pkgs)
	if len(mainPkgs) == 0 {
		return 2, ErrNotFoundMain
	}
	return c.RunPkg(mainPkgs[0], path, args)
}

func (c *Context) RunTest(path string, args []string) error {
	fset := token.NewFileSet()
	// preload regexp for create testing
	c.Loader.Import("regexp")
	pkgs, err := c.LoadDir(fset, path)
	if err != nil {
		return err
	}
	return c.TestPkg(pkgs, path, args)
}

func (ctx *Context) BuildPackage(fset *token.FileSet, pkg *types.Package, files []*ast.File) (*ssa.Package, *types.Info, error) {
	if fset == nil {
		panic("no token.FileSet")
	}
	if pkg.Path() == "" {
		panic("package has no import path")
	}

	info := &types.Info{
		Types:      make(map[ast.Expr]types.TypeAndValue),
		Defs:       make(map[*ast.Ident]types.Object),
		Uses:       make(map[*ast.Ident]types.Object),
		Implicits:  make(map[ast.Node]types.Object),
		Scopes:     make(map[ast.Node]*types.Scope),
		Selections: make(map[*ast.SelectorExpr]*types.Selection),
	}

	tc := &types.Config{
		Importer: NewImporter(ctx.Loader, ctx.External),
		Sizes:    ctx.Sizes,
	}
	if err := types.NewChecker(tc, fset, pkg, info).Files(files); err != nil {
		return nil, nil, err
	}
	for _, v := range info.Types {
		if v.IsType() {
			ctx.Types[v.Type] = true
		}
	}
	prog := ssa.NewProgram(fset, ctx.BuilderMode)

	// Create SSA packages for all imports.
	// Order is not significant.
	created := make(map[*types.Package]bool)
	var createAll func(pkgs []*types.Package)
	createAll = func(pkgs []*types.Package) {
		for _, p := range pkgs {
			if !created[p] {
				created[p] = true
				if !p.Complete() {
					if ctx.Mode&EnableDumpPackage != 0 {
						log.Println("indirect", p)
					}
					p.MarkComplete()
				} else {
					if ctx.Mode&EnableDumpPackage != 0 {
						log.Println("imported", p)
					}
				}
				prog.CreatePackage(p, nil, nil, true)
				createAll(p.Imports())
			}
		}
	}
	// create imports
	createAll(pkg.Imports())
	// create indirect depends
	createAll(ctx.Loader.Packages())

	// Create and build the primary package.
	ssapkg := prog.CreatePackage(pkg, files, info, false)
	ssapkg.Build()
	return ssapkg, info, nil
}

func RunFile(filename string, src interface{}, args []string, mode Mode) (exitCode int, err error) {
	reflectx.Reset()
	ctx := NewContext(mode)
	return ctx.RunFile(filename, src, args)
}

func Run(path string, args []string, mode Mode) (exitCode int, err error) {
	reflectx.Reset()
	ctx := NewContext(mode)
	return ctx.Run(path, args)
}

func RunTest(path string, args []string, mode Mode) error {
	reflectx.Reset()
	ctx := NewContext(mode)
	return ctx.RunTest(path, args)
}
