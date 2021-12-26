package gopbuild

import (
	"bytes"
	"fmt"
	goast "go/ast"
	"go/types"
	"path/filepath"

	"github.com/goplus/gop/ast"
	"github.com/goplus/gop/cl"
	"github.com/goplus/gop/parser"
	"github.com/goplus/gop/token"
	"github.com/goplus/gossa"
	"github.com/goplus/gox"
	"golang.org/x/tools/go/packages"

	_ "github.com/goplus/gossa/pkg/fmt"
	_ "github.com/goplus/gossa/pkg/github.com/goplus/gop/builtin"
	_ "github.com/goplus/gossa/pkg/math/big"
	_ "github.com/goplus/gossa/pkg/strconv"
	_ "github.com/goplus/gossa/pkg/strings"
)

func RegisterClassFileType(extGmx, extSpx string, pkgPaths ...string) {
	cl.RegisterClassFileType(extGmx, extSpx, pkgPaths...)
}

func init() {
	gossa.RegisterFileProcess(".gop", BuildFile)
}

func BuildFile(ctx *gossa.Context, filename string, src interface{}) (data []byte, err error) {
	defer func() {
		r := recover()
		if r != nil {
			err = fmt.Errorf("compile %v failed. %v", filename, r)
		}
	}()
	c := NewContext(ctx)
	fset := token.NewFileSet()
	pkg, err := c.ParseFile(fset, filename, src)
	if err != nil {
		return nil, err
	}
	return pkg.ToSource()
}

func BuildDir(ctx *gossa.Context, dir string) (data []byte, err error) {
	defer func() {
		r := recover()
		if r != nil {
			err = fmt.Errorf("compile %v failed. %v", dir, err)
		}
	}()
	c := NewContext(ctx)
	fset := token.NewFileSet()
	pkg, err := c.ParseDir(fset, dir)
	if err != nil {
		return nil, err
	}
	return pkg.ToSource()
}

type Package struct {
	Fset *token.FileSet
	Pkg  *gox.Package
}

func (p *Package) ToSource() ([]byte, error) {
	var buf bytes.Buffer
	if err := gox.WriteTo(&buf, p.Pkg, false); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (p *Package) ToAst() *goast.File {
	return gox.ASTFile(p.Pkg, false)
}

type Context struct {
	ctx *gossa.Context
}

func NewContext(ctx *gossa.Context) *Context {
	return &Context{ctx}
}

func (c *Context) ParseDir(fset *token.FileSet, dir string) (*Package, error) {
	pkgs, err := parser.ParseDir(fset, dir, nil, 0)
	if err != nil {
		return nil, err
	}
	return c.loadPackage(dir, fset, pkgs)
}

func (c *Context) ParseFile(fset *token.FileSet, filename string, src interface{}) (*Package, error) {
	srcDir, _ := filepath.Split(filename)
	pkgs, err := parser.Parse(fset, filename, src, 0)
	if err != nil {
		return nil, err
	}
	return c.loadPackage(srcDir, fset, pkgs)
}

func typesToPackage(p *types.Package) *packages.Package {
	imports := make(map[string]*packages.Package)
	for _, dep := range p.Imports() {
		imports[dep.Path()] = typesToPackage(dep)
	}
	pkg := &packages.Package{
		ID:      p.Path(),
		Name:    p.Name(),
		PkgPath: p.Path(),
		Types:   p,
		Imports: imports,
	}
	return pkg
}

func (c *Context) loadPackage(srcDir string, fset *token.FileSet, pkgs map[string]*ast.Package) (*Package, error) {
	mainPkg, ok := pkgs["main"]
	if !ok {
		return nil, fmt.Errorf("not a main package")
	}
	conf := &cl.Config{
		Dir: srcDir, TargetDir: srcDir, Fset: fset, CacheLoadPkgs: false, PersistLoadPkgs: false}
	var loaderror error
	conf.PkgsLoader = &cl.PkgsLoader{}
	loaded := make(map[string]bool)
	var load func(at *gox.Package, imports map[string]*gox.PkgRef, pkg *packages.Package)
	load = func(at *gox.Package, imports map[string]*gox.PkgRef, pkg *packages.Package) {
		if loaded[pkg.PkgPath] {
			return
		}
		for _, dep := range pkg.Imports {
			load(at, imports, dep)
		}
		if !pkg.Types.Complete() {
			//return
		}
		loaded[pkg.PkgPath] = true
		gox.LoadGoPkg(at, imports, pkg)
	}
	conf.PkgsLoader.LoadPkgs = func(at *gox.Package, imports map[string]*gox.PkgRef, pkgPaths ...string) int {
		var pkgs []*packages.Package
		for _, path := range pkgPaths {
			p, err := c.ctx.Loader.Import(path)
			if err != nil {
				loaderror = err
				continue
			}
			pkgs = append(pkgs, typesToPackage(p))
		}
		for _, pkg := range pkgs {
			load(at, imports, pkg)
		}
		return 0
	}
	out, err := cl.NewPackage("", mainPkg, conf)
	if loaderror != nil {
		return nil, loaderror
	}
	if err != nil {
		return nil, err
	}
	return &Package{fset, out}, nil
}
