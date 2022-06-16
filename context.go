package igop

import (
	"bytes"
	"flag"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"go/types"
	"io"
	"os"
	"path/filepath"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/goplus/reflectx"

	"golang.org/x/tools/go/ssa"
	"golang.org/x/tools/go/ssa/ssautil"
)

// Mode is a bitmask of options affecting the interpreter.
type Mode uint

const (
	DisableRecover         Mode = 1 << iota // Disable recover() in target programs; show interpreter crash instead.
	DisableCustomBuiltin                    // Disable load custom builtin func
	DisableUnexportMethods                  // Disable unexport methods
	EnableTracing                           // Print a trace of all instructions as they are interpreted.
	EnableDumpInstr                         // Print packages & SSA instruction code
	EnablePrintAny                          // Enable builtin print for any type ( struct/array )
)

// Loader types loader interface
type Loader interface {
	Import(path string) (*types.Package, error)
	Installed(path string) (*Package, bool)
	Packages() []*types.Package
	LookupReflect(typ types.Type) (reflect.Type, bool)
	LookupTypes(typ reflect.Type) (types.Type, bool)
}

// Context ssa context
type Context struct {
	Loader      Loader                   // types loader
	Mode        Mode                     // mode
	ParserMode  parser.Mode              // parser mode
	BuilderMode ssa.BuilderMode          // ssa builder mode
	External    types.Importer           // external import
	Sizes       types.Sizes              // types size for package unsafe
	debugFunc   func(*DebugInfo)         // debug func
	override    map[string]reflect.Value // override function
	output      io.Writer                // capture print/println output
	callForPool int                      // least call count for enable function pool
	evalMode    bool                     // eval mode
	evalInit    map[string]bool          // eval init check
	evalCallFn  func(call *ssa.Call, res ...interface{})
}

func (c *Context) IsEvalMode() bool {
	return c.evalMode
}

// NewContext create a new Context
func NewContext(mode Mode) *Context {
	ctx := &Context{
		Loader:      NewTypesLoader(mode),
		Mode:        mode,
		ParserMode:  parser.AllErrors,
		BuilderMode: 0, //ssa.SanityCheckFunctions,
		override:    make(map[string]reflect.Value),
		callForPool: 64,
	}
	if mode&EnableDumpInstr != 0 {
		ctx.BuilderMode |= ssa.PrintFunctions
	}
	return ctx
}

// SetLeastCallForEnablePool set least call count for enable function pool, default 64
func (c *Context) SetLeastCallForEnablePool(count int) {
	c.callForPool = count
}

func (c *Context) SetDebug(fn func(*DebugInfo)) {
	c.BuilderMode |= ssa.GlobalDebug
	c.debugFunc = fn
}

// SetOverrideFunction register external function to override function.
// match func fullname and signature
func (c *Context) SetOverrideFunction(key string, fn interface{}) {
	c.override[key] = reflect.ValueOf(fn)
}

// ClearOverrideFunction reset override function
func (c *Context) ClearOverrideFunction(key string) {
	delete(c.override, key)
}

// set builtin print/println captured output
func (c *Context) SetPrintOutput(output *bytes.Buffer) {
	c.output = output
}

func (c *Context) writeOutput(data []byte) (n int, err error) {
	if c.output != nil {
		return c.output.Write(data)
	}
	return os.Stdout.Write(data)
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
	sourceProcessor[ext] = fn
}

type SourceProcessFunc func(ctx *Context, filename string, src interface{}) ([]byte, error)

var (
	sourceProcessor = make(map[string]SourceProcessFunc)
)

func (c *Context) LoadFile(fset *token.FileSet, filename string, src interface{}) (*ssa.Package, error) {
	file, err := c.ParseFile(fset, filename, src)
	if err != nil {
		return nil, err
	}
	return c.LoadAstFile(fset, file)
}

func (c *Context) ParseFile(fset *token.FileSet, filename string, src interface{}) (*ast.File, error) {
	if ext := filepath.Ext(filename); ext != "" {
		if fn, ok := sourceProcessor[ext]; ok {
			data, err := fn(c, filename, src)
			if err != nil {
				return nil, err
			}
			src = data
		}
	}
	return parser.ParseFile(fset, filename, src, c.ParserMode)
}

func (c *Context) LoadAstFile(fset *token.FileSet, file *ast.File) (*ssa.Package, error) {
	pkg := types.NewPackage(file.Name.Name, "")
	files := []*ast.File{file}
	if c.Mode&DisableCustomBuiltin == 0 {
		if f, err := ParseBuiltin(fset, file.Name.Name); err == nil {
			files = []*ast.File{f, file}
		}
	}
	ssapkg, _, err := c.BuildPackage(fset, pkg, files)
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
	if c.Mode&DisableCustomBuiltin == 0 {
		if f, err := ParseBuiltin(fset, apkg.Name); err == nil {
			files = append([]*ast.File{f}, files...)
		}
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
	if err = interp.RunInit(); err != nil {
		return 2, err
	}
	return interp.RunMain()
}

func (c *Context) RunFunc(mainPkg *ssa.Package, fnname string, args ...Value) (ret Value, err error) {
	interp, err := c.NewInterp(mainPkg)
	if err != nil {
		return nil, err
	}
	return interp.RunFunc(fnname, args...)
}

func (c *Context) NewInterp(mainPkg *ssa.Package) (*Interp, error) {
	return NewInterp(c, mainPkg)
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
		interp, err := NewInterp(c, pkg)
		if err != nil {
			failed = true
			fmt.Printf("create interp failed: %v\n", err)
			continue
		}
		if err = interp.RunInit(); err != nil {
			failed = true
			fmt.Printf("init error: %v\n", err)
			continue
		}
		exitCode, _ := interp.RunMain()
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
	if ctx.evalMode {
		tc.DisableUnusedImportCheck = true
	}
	if err := types.NewChecker(tc, fset, pkg, info).Files(files); err != nil {
		return nil, nil, err
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
					if ctx.Mode&EnableDumpInstr != 0 {
						fmt.Println("# indirect", p)
					}
					p.MarkComplete()
				} else {
					if ctx.Mode&EnableDumpInstr != 0 {
						fmt.Println("# imported", p)
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

var (
	builtinPkg = &Package{
		Name:          "builtin",
		Path:          "github.com/goplus/igop/builtin",
		Deps:          make(map[string]string),
		Interfaces:    map[string]reflect.Type{},
		NamedTypes:    map[string]NamedType{},
		AliasTypes:    map[string]reflect.Type{},
		Vars:          map[string]reflect.Value{},
		Funcs:         map[string]reflect.Value{},
		TypedConsts:   map[string]TypedConst{},
		UntypedConsts: map[string]UntypedConst{},
	}
	builtinPrefix = "Builtin_"
)

func init() {
	RegisterPackage(builtinPkg)
}

func RegisterCustomBuiltin(key string, fn interface{}) error {
	v := reflect.ValueOf(fn)
	switch v.Kind() {
	case reflect.Func:
		if !strings.HasPrefix(key, builtinPrefix) {
			key = builtinPrefix + key
		}
		builtinPkg.Funcs[key] = v
		typ := v.Type()
		for i := 0; i < typ.NumIn(); i++ {
			checkBuiltinDeps(typ.In(i))
		}
		for i := 0; i < typ.NumOut(); i++ {
			checkBuiltinDeps(typ.Out(i))
		}
		return nil
	}
	return ErrNoFunction
}

func checkBuiltinDeps(typ reflect.Type) {
	if path := typ.PkgPath(); path != "" {
		v := strings.Split(path, "/")
		builtinPkg.Deps[path] = v[len(v)-1]
	}
}

func builtinFuncList() []string {
	var list []string
	for k, v := range builtinPkg.Funcs {
		if strings.HasPrefix(k, builtinPrefix) {
			name := k[len(builtinPrefix):]
			typ := v.Type()
			var ins []string
			var outs []string
			var call []string
			numIn := typ.NumIn()
			numOut := typ.NumOut()
			for i := 0; i < numIn; i++ {
				ins = append(ins, fmt.Sprintf("p%v %v", i, typ.In(i).String()))
				call = append(call, fmt.Sprintf("p%v", i))
			}
			for i := 0; i < numOut; i++ {
				outs = append(outs, typ.Out(i).String())
			}
			var str string
			if numOut > 0 {
				str = fmt.Sprintf("func %v(%v)(%v) { return builtin.%v(%v) }",
					name, strings.Join(ins, ","), strings.Join(outs, ","), k, strings.Join(call, ","))
			} else {
				str = fmt.Sprintf("func %v(%v) { builtin.%v(%v) }",
					name, strings.Join(ins, ","), k, strings.Join(call, ","))
			}
			list = append(list, str)
		}
	}
	return list
}

func ParseBuiltin(fset *token.FileSet, pkg string) (*ast.File, error) {
	list := builtinFuncList()
	if len(list) == 0 {
		return nil, os.ErrInvalid
	}
	var deps []string
	for k := range builtinPkg.Deps {
		deps = append(deps, strconv.Quote(k))
	}
	sort.Strings(deps)
	src := fmt.Sprintf(`package %v
import (
	"github.com/goplus/igop/builtin"
	%v
)
%v
`, pkg, strings.Join(deps, "\n"), strings.Join(list, "\n"))
	return parser.ParseFile(fset, "gossa_builtin.go", src, 0)
}
