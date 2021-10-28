package gossa

import (
	"errors"
	"flag"
	"fmt"
	"go/ast"
	"go/build"
	"go/importer"
	"go/parser"
	"go/token"
	"go/types"
	"os"
	"path/filepath"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/goplus/gossa/internal/gopfile"
	"github.com/goplus/reflectx"
	"golang.org/x/tools/go/packages"
	"golang.org/x/tools/go/ssa"
	"golang.org/x/tools/go/ssa/ssautil"
)

var (
	ErrNoPackage    = errors.New("no package")
	ErrPackage      = errors.New("package contain errors")
	ErrNotFoundMain = errors.New("not found main package")
)

var (
	UnsafeSizes types.Sizes
)

// func loadFile2(input string, src interface{}) (*ssa.Package, error) {
// 	if !filepath.IsAbs(input) {
// 		wd, _ := os.Getwd()
// 		input = filepath.Join(wd, input)
// 	}
// 	const mode = parser.AllErrors | parser.ParseComments
// 	fset := token.NewFileSet()
// 	f, err := parser.ParseFile(fset, input, src, mode)
// 	if err != nil {
// 		return nil, err
// 	}
// 	cfg := &loader.Config{}
// 	cfg.Fset = fset
// 	cfg.CreateFromFiles(input, f)
// 	iprog, err := cfg.Load()
// 	if err != nil {
// 		return nil, err
// 	}
// 	prog := ssautil.CreateProgram(iprog, ssa.SanityCheckFunctions)
// 	prog.Build()
// 	var mainPkg *ssa.Package
// 	if len(iprog.Created) > 0 {
// 		mainPkg = prog.Package(iprog.Created[0].Pkg)
// 	} else {
// 		if pkgs := ssautil.MainPackages(prog.AllPackages()); len(pkgs) > 0 {
// 			mainPkg = pkgs[0]
// 		}
// 	}
// 	if mainPkg == nil {
// 		return nil, ErrNotFoundMain
// 	}
// 	return mainPkg, nil
// }

type Importer struct {
	Default types.Importer
}

func NewImporter() types.Importer {
	return &Importer{importer.Default()}
}

func (i *Importer) Import(path string) (*types.Package, error) {
	if pkg, ok := loadTypesPackage(path); ok {
		return pkg, nil
	}
	pkg, err := i.Default.Import(path)
	return pkg, err
}

func LoadFile(input string, src interface{}) (*ssa.Package, error) {
	if !filepath.IsAbs(input) {
		wd, _ := os.Getwd()
		input = filepath.Join(wd, input)
	}
	const mode = parser.AllErrors | parser.ParseComments
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, input, src, mode)
	if err != nil {
		return nil, err
	}
	// if f.Name.Name != "main" {
	// 	return nil, ErrNotFoundMain
	// }
	var hasOtherPkgs bool
	for _, im := range f.Imports {
		v, _ := strconv.Unquote(im.Path.Value)
		if !externPackages[v] {
			hasOtherPkgs = true
			break
		}
	}
	if !hasOtherPkgs {
		impl := NewImporter()
		pkg := types.NewPackage(f.Name.Name, "")
		var chkerr error
		ssapkg, _, err := BuildPackage(&types.Config{
			Importer: impl,
			Error: func(err error) {
				fmt.Fprintln(os.Stderr, err)
				chkerr = ErrPackage
			},
		}, fset, pkg, []*ast.File{f}, ssa.NaiveForm) //ssa.SanityCheckFunctions)
		if chkerr != nil {
			return nil, chkerr
		}
		if err != nil {
			return nil, err
		}
		ssapkg.Build()
		return ssapkg, nil
	}
	cfg := &packages.Config{
		Fset: fset,
		Mode: packages.NeedName | packages.NeedDeps | packages.LoadTypes | packages.NeedSyntax | packages.NeedTypesInfo | packages.NeedTypesSizes,
	}
	cfg.ParseFile = func(fset *token.FileSet, filename string, src []byte) (*ast.File, error) {
		if filename == input {
			return f, nil
		}
		return parser.ParseFile(fset, filename, src, mode)
	}
	list, err := packages.Load(cfg, input)
	if err != nil {
		return nil, err
	}
	if len(list) == 0 {
		return nil, ErrNoPackage
	}
	if packages.PrintErrors(list) > 0 {
		return nil, ErrPackage
	}
	list[0].ID = "main"
	list[0].PkgPath = "main"
	// hack fix types.Types.Path() command-line-arguments
	v := reflect.ValueOf(list[0].Types).Elem()
	reflectx.Field(v, 0).SetString("main")
	prog, pkgs := ssautil.AllPackages(list, ssa.SanityCheckFunctions)
	prog.Build()
	mainPkgs := ssautil.MainPackages(pkgs)
	if len(mainPkgs) == 0 {
		return nil, ErrNotFoundMain
	}
	return mainPkgs[0], nil
}

func loadPkg(input string) (*ssa.Package, error) {
	wd, _ := os.Getwd()
	p, err := build.Import(input, wd, 0)
	if err != nil {
		return nil, err
	}
	if p.Name != "main" {
		return nil, ErrNotFoundMain
	}
	var hasOtherPkgs bool
	for _, im := range p.Imports {
		if !externPackages[im] {
			hasOtherPkgs = true
			break
		}
	}
	if !hasOtherPkgs {
		const mode = parser.AllErrors | parser.ParseComments
		fset := token.NewFileSet()
		var files []*ast.File
		for _, fname := range p.GoFiles {
			filename := filepath.Join(p.Dir, fname)
			f, err := parser.ParseFile(fset, filename, nil, mode)
			if err != nil {
				return nil, err
			}
			files = append(files, f)
		}
		pkg := types.NewPackage("main", "")
		var chkerr error
		ssapkg, _, err := ssautil.BuildPackage(&types.Config{
			Importer: importer.Default(),
			Error: func(err error) {
				fmt.Fprintln(os.Stderr, err)
				chkerr = ErrPackage
			},
		}, fset, pkg, files, ssa.SanityCheckFunctions)
		if chkerr != nil {
			return nil, chkerr
		}
		if err != nil {
			return nil, err
		}
		ssapkg.Build()
		return ssapkg, nil
	}
	cfg := &packages.Config{
		Mode: packages.NeedName | packages.NeedDeps | packages.LoadTypes | packages.NeedSyntax | packages.NeedTypesInfo | packages.NeedTypesSizes,
	}
	list, err := packages.Load(cfg, input)
	if err != nil {
		return nil, err
	}
	if len(list) == 0 {
		return nil, ErrNoPackage
	}
	if packages.PrintErrors(list) > 0 {
		return nil, ErrPackage
	}
	prog, pkgs := ssautil.AllPackages(list, ssa.SanityCheckFunctions)
	prog.Build()
	mainPkgs := ssautil.MainPackages(pkgs)
	if len(mainPkgs) == 0 {
		return nil, ErrNotFoundMain
	}
	return mainPkgs[0], nil
}

func Run(mode Mode, input string, args []string) error {
	if strings.HasSuffix(input, ".go") {
		return RunFile(mode, input, nil, args)
	}
	pkg, err := loadPkg(input)
	if err != nil {
		return err
	}
	return RunPkg(pkg, mode, input, "main", args)
}

func foundPkg(pkg string) (*build.Package, error) {
	if filepath.IsAbs(pkg) {
		return build.ImportDir(pkg, build.FindOnly)
	} else {
		return build.Import(pkg, ".", build.FindOnly)
	}
}

func RunTest(mode Mode, input string, args []string) error {
	p, err := foundPkg(input)
	if err != nil {
		return fmt.Errorf("not found pkg: %v", err)
	}
	if p.Dir != "." {
		os.Chdir(p.Dir)
	}
	pkgpath, pkgs, err := LoadTest(".")
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
	ext := filepath.Ext(filename)
	switch ext {
	case ".gop":
		data, err := gopfile.Build(filename, src)
		if err != nil {
			return err
		}
		src = data
	}
	pkg, err := LoadFile(filename, src)
	if err != nil {
		return err
	}
	return RunPkg(pkg, mode, filename, "main", args)
}

func LoadTest(input string) (string, []*ssa.Package, error) {
	cfg := &packages.Config{
		Mode:  packages.NeedName | packages.NeedDeps | packages.LoadTypes | packages.NeedSyntax | packages.NeedTypesInfo | packages.NeedTypesSizes,
		Tests: true,
	}
	list, err := packages.Load(cfg, input)
	if err != nil {
		return "", nil, err
	}
	var foundError bool
	for _, v := range list {
		if len(v.Errors) > 0 {
			for _, err := range v.Errors {
				fmt.Fprintln(os.Stderr, err)
			}
			foundError = true
		}
	}
	if foundError {
		return "", nil, errors.New("build failed")
	}
	prog, ppkgs := ssautil.AllPackages(list, ssa.SanityCheckFunctions)
	prog.Build()
	var pkgs []*ssa.Package
	for _, p := range ppkgs {
		if p == nil {
			continue
		}
		pkgs = append(pkgs, p)
	}
	if len(pkgs) == 0 {
		return "", nil, errors.New("not found package")
	}
	return pkgs[0].Pkg.Path(), pkgs, nil
}

func RunTestPkg(pkgs []*ssa.Package, mode Mode, input string, args []string) {
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

	interp := NewInterp(mainPkg, mode)
	interp.Run("init")
	_, exitCode := interp.Run(entry)
	if exitCode != 0 {
		return fmt.Errorf("interpreting %v: exit code was %d", input, exitCode)
	}
	return nil
}
