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

var (
	classfile = make(map[string]*cl.Class)
)

func RegisterClassFileType(extGmx, extSpx string, pkgPaths ...string) {
	//cl.RegisterClassFileType(extGmx, extSpx, pkgPaths...)
	cls := &cl.Class{
		ProjExt:  extGmx,
		WorkExt:  extSpx,
		PkgPaths: pkgPaths,
	}
	classfile[extGmx] = cls
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
	ctx  *gossa.Context
	conf *parser.Config
}

func NewContext(ctx *gossa.Context) *Context {
	return &Context{ctx: ctx, conf: &parser.Config{
		IsClass: func(ext string) (isProj bool, ok bool) {
			if cls, ok := classfile[ext]; ok {
				return ext == cls.ProjExt, true
			}
			return
		},
	}}
}

func (c *Context) ParseDir(fset *token.FileSet, dir string) (*Package, error) {
	pkgs, err := parser.ParseDirEx(fset, dir, *c.conf)
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
		WorkingDir: srcDir, TargetDir: srcDir, Fset: fset}
	conf.Importer = c.ctx.Loader
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
