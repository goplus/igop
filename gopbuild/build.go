package gopbuild

//go:generate go run ../cmd/qexp -outdir ../pkg github.com/goplus/gop/builtin
//go:generate go run ../cmd/qexp -outdir ../pkg github.com/goplus/gop/builtin/ng

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

	_ "github.com/goplus/gossa/pkg/fmt"
	_ "github.com/goplus/gossa/pkg/github.com/goplus/gop/builtin"
	_ "github.com/goplus/gossa/pkg/github.com/goplus/gop/builtin/ng"
	_ "github.com/goplus/gossa/pkg/math/big"
	_ "github.com/goplus/gossa/pkg/strconv"
	_ "github.com/goplus/gossa/pkg/strings"
)

var (
	classfile = make(map[string]*cl.Class)
)

func RegisterClassFileType(projExt, workExt string, pkgPaths ...string) {
	cls := &cl.Class{
		ProjExt:  projExt,
		WorkExt:  workExt,
		PkgPaths: pkgPaths,
	}
	classfile[projExt] = cls
	if workExt != "" {
		classfile[workExt] = cls
	}
}

func init() {
	gossa.RegisterFileProcess(".gop", BuildFile)
	RegisterClassFileType(".gmx", ".spx", "github.com/goplus/spx", "math")
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

func BuildFSDir(ctx *gossa.Context, fs parser.FileSystem, dir string) (data []byte, err error) {
	defer func() {
		r := recover()
		if r != nil {
			err = fmt.Errorf("compile %v failed. %v", dir, err)
		}
	}()
	c := NewContext(ctx)
	fset := token.NewFileSet()
	pkg, err := c.ParseFSDir(fset, fs, dir)
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
	gop gossa.Loader
}

func IsClass(ext string) (isProj bool, ok bool) {
	if cls, ok := classfile[ext]; ok {
		return ext == cls.ProjExt, true
	}
	return
}

func NewContext(ctx *gossa.Context) *Context {
	return &Context{ctx: ctx, gop: gossa.NewTypesLoader(0)}
}

func isGopPackage(path string) bool {
	if pkg, ok := gossa.LookupPackage(path); ok {
		if _, ok := pkg.UntypedConsts["GopPackage"]; ok {
			return true
		}
	}
	return false
}

func (c *Context) Import(path string) (*types.Package, error) {
	if isGopPackage(path) {
		return c.gop.Import(path)
	}
	return c.ctx.Loader.Import(path)
}

func (c *Context) ParseDir(fset *token.FileSet, dir string) (*Package, error) {
	pkgs, err := parser.ParseDirEx(fset, dir, parser.Config{
		IsClass: IsClass,
	})
	if err != nil {
		return nil, err
	}
	return c.loadPackage(dir, fset, pkgs)
}

func (c *Context) ParseFSDir(fset *token.FileSet, fs parser.FileSystem, dir string) (*Package, error) {
	pkgs, err := parser.ParseFSDir(fset, fs, dir, parser.Config{
		IsClass: IsClass,
	})
	if err != nil {
		return nil, err
	}
	return c.loadPackage(dir, fset, pkgs)
}

func (c *Context) ParseFile(fset *token.FileSet, filename string, src interface{}) (*Package, error) {
	srcDir, _ := filepath.Split(filename)
	f, err := parser.ParseFile(fset, filename, src, 0)
	if err != nil {
		return nil, err
	}
	f.IsProj, f.IsClass = IsClass(filepath.Ext(filename))
	name := f.Name.Name
	pkgs := map[string]*ast.Package{
		name: &ast.Package{
			Name: name,
			Files: map[string]*ast.File{
				filename: f,
			},
		},
	}
	return c.loadPackage(srcDir, fset, pkgs)
}

func (c *Context) loadPackage(srcDir string, fset *token.FileSet, pkgs map[string]*ast.Package) (*Package, error) {
	mainPkg, ok := pkgs["main"]
	if !ok {
		return nil, fmt.Errorf("not a main package")
	}
	conf := &cl.Config{
		WorkingDir: srcDir, TargetDir: srcDir, Fset: fset}
	conf.Importer = c
	conf.LookupClass = func(ext string) (c *cl.Class, ok bool) {
		c, ok = classfile[ext]
		return
	}
	out, err := cl.NewPackage("", mainPkg, conf)
	if err != nil {
		return nil, err
	}
	return &Package{fset, out}, nil
}
