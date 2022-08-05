package igop

import (
	"bytes"
	"flag"
	"fmt"
	"go/ast"
	"go/build"
	"go/parser"
	"go/token"
	"go/types"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/goplus/igop/testmain"
	"github.com/visualfc/gomod"
	"golang.org/x/tools/go/ssa"
)

// Mode is a bitmask of options affecting the interpreter.
type Mode uint

const (
	DisableRecover       Mode = 1 << iota // Disable recover() in target programs; show interpreter crash instead.
	DisableCustomBuiltin                  // Disable load custom builtin func
	EnableDumpImports                     // print typesimports
	EnableDumpInstr                       // Print packages & SSA instruction code
	EnableTracing                         // Print a trace of all instructions as they are interpreted.
	EnablePrintAny                        // Enable builtin print for any type ( struct/array )
)

// Loader types loader interface
type Loader interface {
	Import(path string) (*types.Package, error)
	Installed(path string) (*Package, bool)
	Packages() []*types.Package
	LookupReflect(typ types.Type) (reflect.Type, bool)
	LookupTypes(typ reflect.Type) (types.Type, bool)
	SetImport(path string, pkg *types.Package, load func() error) error
}

// Context ssa context
type Context struct {
	Loader       Loader                                           // types loader
	FileSet      *token.FileSet                                   // file set
	Mode         Mode                                             // mode
	ParserMode   parser.Mode                                      // parser mode
	BuilderMode  ssa.BuilderMode                                  // ssa builder mode
	BuildContext build.Context                                    // build context
	Lookup       func(root, path string) (dir string, found bool) // lookup external import
	pkgs         map[string]*sourcePackage                        // imports
	override     map[string]reflect.Value                         // override function
	output       io.Writer                                        // capture print/println output
	callForPool  int                                              // least call count for enable function pool
	conf         *types.Config                                    // types check config
	evalMode     bool                                             // eval mode
	evalInit     map[string]bool                                  // eval init check
	evalCallFn   func(interp *Interp, call *ssa.Call, res ...interface{})
	debugFunc    func(*DebugInfo) // debug func
	root         string           // project root
	mod          *gomod.Package   // lookup path for go.mod
}

func (c *Context) setRoot(root string) {
	c.root = root
	c.mod = nil
}

func (c *Context) lookupPath(path string) (dir string, found bool) {
	if c.Lookup != nil {
		dir, found = c.Lookup(c.root, path)
		if found {
			return
		}
	}
	if c.mod == nil {
		var err error
		c.mod, err = gomod.Load(c.root, &c.BuildContext)
		if err != nil {
			panic(err)
		}
	}
	_, dir, found = c.mod.Lookup(path)
	if !found {
		bp, err := build.Import(path, c.root, build.FindOnly)
		if err == nil && bp.ImportPath == path {
			return bp.Dir, true
		}
	}
	return
}

type sourcePackage struct {
	Context *Context
	Package *types.Package
	Info    *types.Info
	Files   []*ast.File
	Dir     string
}

func (sp *sourcePackage) Load() (err error) {
	if sp.Info == nil {
		sp.Info = &types.Info{
			Types:      make(map[ast.Expr]types.TypeAndValue),
			Defs:       make(map[*ast.Ident]types.Object),
			Uses:       make(map[*ast.Ident]types.Object),
			Implicits:  make(map[ast.Node]types.Object),
			Scopes:     make(map[ast.Node]*types.Scope),
			Selections: make(map[*ast.SelectorExpr]*types.Selection),
		}
		if err := types.NewChecker(sp.Context.conf, sp.Context.FileSet, sp.Package, sp.Info).Files(sp.Files); err != nil {
			return err
		}
	}
	return
}

// NewContext create a new Context
func NewContext(mode Mode) *Context {
	ctx := &Context{
		Loader:       NewTypesLoader(mode),
		FileSet:      token.NewFileSet(),
		Mode:         mode,
		ParserMode:   parser.AllErrors,
		BuilderMode:  0, //ssa.SanityCheckFunctions,
		BuildContext: build.Default,
		pkgs:         make(map[string]*sourcePackage),
		override:     make(map[string]reflect.Value),
		callForPool:  64,
	}
	if mode&EnableDumpInstr != 0 {
		ctx.BuilderMode |= ssa.PrintFunctions
	}
	ctx.conf = &types.Config{
		Importer: NewImporter(ctx),
	}
	return ctx
}

func (c *Context) IsEvalMode() bool {
	return c.evalMode
}

func (c *Context) SetEvalMode(b bool) {
	c.evalMode = b
	c.conf.DisableUnusedImportCheck = b
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

func importPathForDir(dir string) (string, error) {
	cmd := exec.Command("go", "list")
	cmd.Dir = dir
	data, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(data)), nil
}

func (c *Context) LoadDir(dir string, test bool) (pkg *ssa.Package, err error) {
	var sp *sourcePackage
	if test {
		sp, err = c.loadTestPackage(dir)
	} else {
		sp, err = c.loadPackage("main", dir)
	}
	if err != nil {
		return nil, err
	}
	if c.Mode&DisableCustomBuiltin == 0 {
		if f, err := ParseBuiltin(c.FileSet, sp.Package.Name()); err == nil {
			sp.Files = append([]*ast.File{f}, sp.Files...)
		}
	}
	c.setRoot(dir)
	if dir != "." {
		if wd, err := os.Getwd(); err == nil {
			os.Chdir(dir)
			defer os.Chdir(wd)
		}
	}
	err = sp.Load()
	if err != nil {
		return nil, err
	}
	return c.buildPackage(sp)
}

func RegisterFileProcess(ext string, fn SourceProcessFunc) {
	sourceProcessor[ext] = fn
}

type SourceProcessFunc func(ctx *Context, filename string, src interface{}) ([]byte, error)

var (
	sourceProcessor = make(map[string]SourceProcessFunc)
)

func (c *Context) AddImport(path string, filename string, src interface{}) (err error) {
	_, err = c.addImport(path, filename, src)
	return
}

func (c *Context) AddImportDir(path string, dir string) (err error) {
	_, err = c.addImportDir(path, dir)
	return
}

func (c *Context) addImport(path string, filename string, src interface{}) (*sourcePackage, error) {
	file, err := c.ParseFile(filename, src)
	if err != nil {
		return nil, err
	}
	pkg := types.NewPackage(path, file.Name.Name)
	tp := &sourcePackage{
		Context: c,
		Package: pkg,
		Files:   []*ast.File{file},
	}
	c.pkgs[path] = tp
	c.Loader.SetImport(path, pkg, tp.Load)
	return tp, nil
}

func (c *Context) addImportDir(path string, dir string) (*sourcePackage, error) {
	tp, err := c.loadPackage(path, dir)
	if err != nil {
		return nil, err
	}
	c.pkgs[path] = tp
	c.Loader.SetImport(path, tp.Package, tp.Load)
	return tp, nil
}

func (c *Context) loadPackage(path string, dir string) (*sourcePackage, error) {
	bp, err := c.BuildContext.ImportDir(dir, 0)
	if err != nil {
		return nil, err
	}
	files, err := c.parseGoFiles(dir, append(bp.GoFiles, bp.CgoFiles...))
	if err != nil {
		return nil, err
	}
	tp := &sourcePackage{
		Package: types.NewPackage(path, bp.Name),
		Files:   files,
		Dir:     dir,
		Context: c,
	}
	c.pkgs[path] = tp
	return tp, nil
}

func (c *Context) loadTestPackage(dir string) (*sourcePackage, error) {
	importPath, err := importPathForDir(dir)
	if err != nil {
		return nil, err
	}
	bp, err := c.BuildContext.ImportDir(dir, 0)
	if err != nil {
		return nil, err
	}
	if len(bp.TestGoFiles) == 0 && len(bp.XTestGoFiles) == 0 {
		return nil, ErrNoTestFiles
	}
	bp.ImportPath = importPath
	files, err := c.parseGoFiles(dir, append(append(bp.GoFiles, bp.CgoFiles...), bp.TestGoFiles...))
	if err != nil {
		return nil, err
	}
	tp := &sourcePackage{
		Package: types.NewPackage(importPath, bp.Name),
		Files:   files,
		Dir:     dir,
		Context: c,
	}
	c.pkgs[importPath] = tp
	if len(bp.XTestGoFiles) > 0 {
		files, err := c.parseGoFiles(dir, bp.XTestGoFiles)
		if err != nil {
			return nil, err
		}
		tp := &sourcePackage{
			Package: types.NewPackage(importPath+"_test", bp.Name+"_test"),
			Files:   files,
			Dir:     dir,
			Context: c,
		}
		c.pkgs[importPath+"_test"] = tp
	}
	data, err := testmain.Load(bp)
	if err != nil {
		return nil, err
	}
	f, err := parser.ParseFile(c.FileSet, "_testmain.go", data, 0)
	if err != nil {
		return nil, err
	}
	return &sourcePackage{
		Package: types.NewPackage(importPath+"$testmain", "main"),
		Files:   []*ast.File{f},
		Dir:     dir,
		Context: c,
	}, nil
}

func (c *Context) parseGoFiles(dir string, filenames []string) ([]*ast.File, error) {
	files := make([]*ast.File, len(filenames))
	errors := make([]error, len(filenames))

	var wg sync.WaitGroup
	wg.Add(len(filenames))
	for i, filename := range filenames {
		go func(i int, filepath string) {
			defer wg.Done()
			files[i], errors[i] = parser.ParseFile(c.FileSet, filepath, nil, 0)
		}(i, filepath.Join(dir, filename))
	}
	wg.Wait()

	for _, err := range errors {
		if err != nil {
			return nil, err
		}
	}
	return files, nil
}

func (c *Context) LoadFile(filename string, src interface{}) (*ssa.Package, error) {
	file, err := c.ParseFile(filename, src)
	if err != nil {
		return nil, err
	}
	root, _ := filepath.Split(filename)
	c.setRoot(root)
	return c.LoadAstFile("main", file)
}

func (c *Context) ParseFile(filename string, src interface{}) (*ast.File, error) {
	return c.parseFile(filename, src, c.ParserMode)
}

func (c *Context) parseFile(filename string, src interface{}, mode parser.Mode) (*ast.File, error) {
	if ext := filepath.Ext(filename); ext != "" {
		if fn, ok := sourceProcessor[ext]; ok {
			data, err := fn(c, filename, src)
			if err != nil {
				return nil, err
			}
			src = data
		}
	}
	return parser.ParseFile(c.FileSet, filename, src, mode)
}

func (c *Context) LoadAstFile(path string, file *ast.File) (*ssa.Package, error) {
	files := []*ast.File{file}
	if c.Mode&DisableCustomBuiltin == 0 {
		if f, err := ParseBuiltin(c.FileSet, file.Name.Name); err == nil {
			files = []*ast.File{f, file}
		}
	}
	sp := &sourcePackage{
		Context: c,
		Package: types.NewPackage(path, file.Name.Name),
		Files:   files,
	}
	err := sp.Load()
	if err != nil {
		return nil, err
	}
	return c.buildPackage(sp)
}

func (c *Context) LoadAstPackage(path string, apkg *ast.Package) (*ssa.Package, error) {
	var files []*ast.File
	for _, f := range apkg.Files {
		files = append(files, f)
	}
	if c.Mode&DisableCustomBuiltin == 0 {
		if f, err := ParseBuiltin(c.FileSet, apkg.Name); err == nil {
			files = append([]*ast.File{f}, files...)
		}
	}
	sp := &sourcePackage{
		Context: c,
		Package: types.NewPackage(path, apkg.Name),
		Files:   files,
	}
	err := sp.Load()
	if err != nil {
		return nil, err
	}
	return c.buildPackage(sp)
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

func (c *Context) TestPkg(pkg *ssa.Package, input string, args []string) error {
	var failed bool
	start := time.Now()
	defer func() {
		sec := time.Since(start).Seconds()
		if failed {
			fmt.Fprintf(os.Stdout, "FAIL\t%s %0.3fs\n", input, sec)
		} else {
			fmt.Fprintf(os.Stdout, "ok\t%s %0.3fs\n", input, sec)
		}
	}()
	os.Args = []string{input}
	if args != nil {
		os.Args = append(os.Args, args...)
	}
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	interp, err := NewInterp(c, pkg)
	if err != nil {
		failed = true
		fmt.Printf("create interp failed: %v\n", err)
	}
	if err = interp.RunInit(); err != nil {
		failed = true
		fmt.Printf("init error: %v\n", err)
	}
	exitCode, _ := interp.RunMain()
	if exitCode != 0 {
		failed = true
	}
	if failed {
		return ErrTestFailed
	}
	return nil
}

func (c *Context) RunFile(filename string, src interface{}, args []string) (exitCode int, err error) {
	pkg, err := c.LoadFile(filename, src)
	if err != nil {
		return 2, err
	}
	return c.RunPkg(pkg, filename, args)
}

func (c *Context) Run(path string, args []string) (exitCode int, err error) {
	if strings.HasSuffix(path, ".go") {
		return c.RunFile(path, nil, args)
	}
	pkg, err := c.LoadDir(path, false)
	if err != nil {
		return 2, err
	}
	if !isMainPkg(pkg) {
		return 2, ErrNotFoundMain
	}
	return c.RunPkg(pkg, path, args)
}

func isMainPkg(pkg *ssa.Package) bool {
	return pkg.Pkg.Name() == "main" && pkg.Func("main") != nil
}

func (c *Context) RunTest(dir string, args []string) error {
	pkg, err := c.LoadDir(dir, true)
	if err != nil {
		if err == ErrNoTestFiles {
			fmt.Println("?", err)
			return nil
		}
		return err
	}
	if filepath.IsAbs(dir) {
		os.Chdir(dir)
	}
	return c.TestPkg(pkg, dir, args)
}

func (ctx *Context) checkTypesInfo(pkg *types.Package, files []*ast.File) (*types.Info, error) {
	info := &types.Info{
		Types:      make(map[ast.Expr]types.TypeAndValue),
		Defs:       make(map[*ast.Ident]types.Object),
		Uses:       make(map[*ast.Ident]types.Object),
		Implicits:  make(map[ast.Node]types.Object),
		Scopes:     make(map[ast.Node]*types.Scope),
		Selections: make(map[*ast.SelectorExpr]*types.Selection),
	}
	if err := types.NewChecker(ctx.conf, ctx.FileSet, pkg, info).Files(files); err != nil {
		return nil, err
	}
	return info, nil
}

func (ctx *Context) buildPackage(sp *sourcePackage) (*ssa.Package, error) {
	prog := ssa.NewProgram(ctx.FileSet, ctx.BuilderMode)
	// Create SSA packages for all imports.
	// Order is not significant.
	created := make(map[*types.Package]bool)
	var createAll func(pkgs []*types.Package)
	createAll = func(pkgs []*types.Package) {
		for _, p := range pkgs {
			if !created[p] {
				created[p] = true
				createAll(p.Imports())
				if pkg, ok := ctx.pkgs[p.Path()]; ok {
					if ctx.Mode&EnableDumpImports != 0 {
						if pkg.Dir != "" {
							fmt.Println("# imported", p.Path(), pkg.Dir)
						} else {
							fmt.Println("# imported", p.Path(), "source")
						}
					}
					prog.CreatePackage(p, pkg.Files, pkg.Info, true).Build()
				} else {
					var indirect bool
					if !p.Complete() {
						indirect = true
						p.MarkComplete()
					}
					if ctx.Mode&EnableDumpImports != 0 {
						if indirect {
							fmt.Println("# indirect", p.Path())
						} else {
							fmt.Println("# imported", p.Path())
						}
					}
					prog.CreatePackage(p, nil, nil, true).Build()
				}
			}
		}
	}
	createAll(sp.Package.Imports())
	// Create and build the primary package.
	pkg := prog.CreatePackage(sp.Package, sp.Files, sp.Info, false)
	pkg.Build()
	return pkg, nil
}

func RunFile(filename string, src interface{}, args []string, mode Mode) (exitCode int, err error) {
	ctx := NewContext(mode)
	return ctx.RunFile(filename, src, args)
}

func Run(path string, args []string, mode Mode) (exitCode int, err error) {
	ctx := NewContext(mode)
	return ctx.Run(path, args)
}

func RunTest(path string, args []string, mode Mode) error {
	ctx := NewContext(mode)
	return ctx.RunTest(path, args)
}

var (
	builtinPkg = &Package{
		Name:          "builtin",
		Path:          "github.com/goplus/igop/builtin",
		Deps:          make(map[string]string),
		Interfaces:    map[string]reflect.Type{},
		NamedTypes:    map[string]reflect.Type{},
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
			if typ.IsVariadic() {
				for i := 0; i < numIn-1; i++ {
					ins = append(ins, fmt.Sprintf("p%v %v", i, typ.In(i).String()))
					call = append(call, fmt.Sprintf("p%v", i))
				}
				ins = append(ins, fmt.Sprintf("p%v ...%v", numIn-1, typ.In(numIn-1).Elem().String()))
				call = append(call, fmt.Sprintf("p%v...", numIn-1))
			} else {
				for i := 0; i < numIn; i++ {
					ins = append(ins, fmt.Sprintf("p%v %v", i, typ.In(i).String()))
					call = append(call, fmt.Sprintf("p%v", i))
				}
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
