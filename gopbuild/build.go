/*
 * Copyright (c) 2022 The GoPlus Authors (goplus.org). All rights reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package gopbuild

import (
	"bytes"
	"fmt"
	goast "go/ast"
	"go/types"
	"os"
	"path/filepath"

	"github.com/goplus/gogen"
	"github.com/goplus/gop/ast"
	"github.com/goplus/gop/cl"
	"github.com/goplus/gop/parser"
	"github.com/goplus/gop/token"
	"github.com/goplus/igop"
	"github.com/goplus/mod/modfile"
)

type Class = cl.Class

var (
	projects = make(map[string]*cl.Project)
)

func RegisterClassFileType(ext string, class string, works []*Class, pkgPaths ...string) {
	cls := &cl.Project{
		Ext:      ext,
		Class:    class,
		Works:    works,
		PkgPaths: pkgPaths,
	}
	if ext != "" {
		projects[ext] = cls
	}
	for _, w := range works {
		projects[w.Ext] = cls
	}
}

func init() {
	cl.SetDebug(cl.FlagNoMarkAutogen)
	igop.RegisterFileProcess(".gop", BuildFile)
	igop.RegisterFileProcess(".gox", BuildFile)
	igop.RegisterFileProcess(".gsh", BuildFile)
	RegisterClassFileType(".gmx", "Game", []*Class{{Ext: ".spx", Class: "Sprite"}}, "github.com/goplus/spx", "math")
	RegisterClassFileType(".gsh", "App", nil, "github.com/qiniu/x/gsh", "math")
	RegisterClassFileType("_test.gox", "App", []*Class{{Ext: "_test.gox", Class: "Case"}}, "github.com/goplus/gop/test")
}

func BuildFile(ctx *igop.Context, filename string, src interface{}) (data []byte, err error) {
	defer func() {
		r := recover()
		if r != nil {
			err = fmt.Errorf("compile %v failed. %v", filename, r)
		}
	}()
	c := NewContext(ctx)
	pkg, err := c.ParseFile(filename, src)
	if err != nil {
		return nil, err
	}
	return pkg.ToSource()
}

func BuildFSDir(ctx *igop.Context, fs parser.FileSystem, dir string) (data []byte, err error) {
	defer func() {
		r := recover()
		if r != nil {
			err = fmt.Errorf("compile %v failed. %v", dir, r)
		}
	}()
	c := NewContext(ctx)
	pkg, err := c.ParseFSDir(fs, dir)
	if err != nil {
		return nil, err
	}
	return pkg.ToSource()
}

func BuildDir(ctx *igop.Context, dir string) (data []byte, err error) {
	defer func() {
		r := recover()
		if r != nil {
			err = fmt.Errorf("compile %v failed. %v", dir, r)
		}
	}()
	c := NewContext(ctx)
	pkg, err := c.ParseDir(dir)
	if err != nil {
		return nil, err
	}
	return pkg.ToSource()
}

type Package struct {
	Fset *token.FileSet
	Pkgs []*gogen.Package
}

func (p *Package) MainPkg() *gogen.Package {
	if len(p.Pkgs) == 0 {
		return nil
	}
	for _, pkg := range p.Pkgs {
		if pkg.Types.Name() == "main" {
			return pkg
		}
	}
	return p.Pkgs[0]
}

func (p *Package) ToSource() ([]byte, error) {
	pkg := p.MainPkg()
	if pkg == nil {
		return nil, os.ErrNotExist
	}
	var buf bytes.Buffer
	if err := pkg.WriteTo(&buf); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (p *Package) ToAst() (*goast.File, error) {
	pkg := p.MainPkg()
	if pkg == nil {
		return nil, os.ErrNotExist
	}
	return pkg.ASTFile(), nil
}

func (p *Package) ForEachFile(fn func(pkg *gogen.Package, fname string)) {
	for _, pkg := range p.Pkgs {
		pkg.ForEachFile(func(fname string, file *gogen.File) {
			fn(pkg, fname)
		})
	}
}

type Context struct {
	Context  *igop.Context
	FileSet  *token.FileSet
	Importer *igop.Importer
	Loader   igop.Loader
	pkgs     map[string]*types.Package
}

func ClassKind(fname string) (isProj, ok bool) {
	ext := modfile.ClassExt(fname)
	switch ext {
	case ".gmx", ".gsh":
		return true, true
	case ".spx":
		return fname == "main.spx", true
	default:
		if c, ok := projects[ext]; ok {
			for _, w := range c.Works {
				if w.Ext == ext {
					if ext != c.Ext || fname != "main"+ext {
						return false, true
					}
					break
				}
			}
			return true, true
		}
	}
	return
}

func NewContext(ctx *igop.Context) *Context {
	if ctx.IsEvalMode() {
		ctx = igop.NewContext(0)
	}
	ctx.Mode |= igop.CheckGopOverloadFunc
	c := &Context{Context: ctx, Importer: igop.NewImporter(ctx), FileSet: token.NewFileSet(),
		Loader: igop.NewTypesLoader(ctx, 0), pkgs: make(map[string]*types.Package)}
	return c
}

func RegisterPackagePatch(ctx *igop.Context, path string, src interface{}) error {
	err := ctx.AddImportFile(path+"@patch", "src.go", src)
	if err != nil {
		return err
	}
	return ctx.AddImportFile(path+"@patch.gop", "src.go", src)
}

func isGopPackage(path string) bool {
	if pkg, ok := igop.LookupPackage(path); ok {
		if _, ok := pkg.UntypedConsts["GopPackage"]; ok {
			return true
		}
	}
	return false
}

func (c *Context) importPath(path string) (gop bool, pkg *types.Package, err error) {
	if isGopPackage(path) {
		gop = true
		pkg, err = c.Loader.Import(path)
	} else {
		pkg, err = c.Importer.Import(path)
	}
	return
}

func (c *Context) Import(path string) (*types.Package, error) {
	if pkg, ok := c.pkgs[path]; ok {
		return pkg, nil
	}
	gop, pkg, err := c.importPath(path)
	if err != nil {
		return pkg, err
	}
	c.pkgs[path] = pkg
	if gop {
		if sp := c.Context.SourcePackage(path + "@patch.gop"); sp != nil {
			sp.Importer = c
			err := sp.Load()
			if err != nil {
				return nil, err
			}
			patch := types.NewPackage(path+"@patch", pkg.Name())
			for _, name := range sp.Package.Scope().Names() {
				obj := sp.Package.Scope().Lookup(name)
				switch obj.(type) {
				case *types.Func:
					obj = types.NewFunc(obj.Pos(), patch, obj.Name(), obj.Type().(*types.Signature))
				case *types.TypeName:
					named := obj.Type().(*types.Named)
					var methods []*types.Func
					if n := named.NumMethods(); n > 0 {
						methods = make([]*types.Func, n)
						for i := 0; i < n; i++ {
							methods[i] = named.Method(i)
						}
					}
					obj = types.NewTypeName(obj.Pos(), patch, obj.Name(), nil)
					types.NewNamed(obj.(*types.TypeName), named.Underlying(), methods)
				default:
					continue
				}
				pkg.Scope().Insert(obj)
			}
		}
	}
	return pkg, nil
}

func (c *Context) ParseDir(dir string) (*Package, error) {
	pkgs, err := parser.ParseDirEx(c.FileSet, dir, parser.Config{
		ClassKind: ClassKind,
	})
	if err != nil {
		return nil, err
	}
	return c.loadPackage(dir, pkgs)
}

func (c *Context) ParseFSDir(fs parser.FileSystem, dir string) (*Package, error) {
	pkgs, err := parser.ParseFSDir(c.FileSet, fs, dir, parser.Config{
		ClassKind: ClassKind,
	})
	if err != nil {
		return nil, err
	}
	return c.loadPackage(dir, pkgs)
}

func (c *Context) ParseFile(fname string, src interface{}) (*Package, error) {
	srcDir, _ := filepath.Split(fname)
	fnameRmGox := fname
	ext := filepath.Ext(fname)
	var err error
	var isProj, isClass, isNormalGox, rmGox bool
	switch ext {
	case ".go", ".gop":
	case ".gox":
		isClass = true
		t := fname[:len(fname)-4]
		if c := filepath.Ext(t); c != "" {
			ext, fnameRmGox, rmGox = c, t, true
		} else {
			isNormalGox = true
		}
		fallthrough
	default:
		if !isNormalGox {
			if isProj, isClass = ClassKind(fnameRmGox); !isClass {
				if rmGox {
					return nil, fmt.Errorf("not found Go+ class by ext %q for %q", ext, fname)
				}
				return nil, nil
			}
		}
	}
	if err != nil {
		return nil, err
	}
	mode := parser.ParseComments
	if isClass {
		mode |= parser.ParseGoPlusClass
	}
	f, err := parser.ParseFile(c.FileSet, fname, src, mode)
	if err != nil {
		return nil, err
	}
	f.IsProj, f.IsClass, f.IsNormalGox = isProj, isClass, isNormalGox
	name := f.Name.Name
	pkgs := map[string]*ast.Package{
		name: &ast.Package{
			Name: name,
			Files: map[string]*ast.File{
				fname: f,
			},
		},
	}
	return c.loadPackage(srcDir, pkgs)
}

func (c *Context) loadPackage(srcDir string, apkgs map[string]*ast.Package) (*Package, error) {
	var pkgs []*gogen.Package
	for _, apkg := range apkgs {
		if c.Context.Mode&igop.DisableCustomBuiltin == 0 {
			if f, err := igop.ParseBuiltin(c.FileSet, apkg.Name); err == nil {
				apkg.GoFiles = map[string]*goast.File{"_igop_builtin.go": f}
			}
		}
		conf := &cl.Config{Fset: c.FileSet}
		conf.Importer = c
		conf.LookupClass = func(ext string) (c *cl.Project, ok bool) {
			c, ok = projects[ext]
			return
		}
		if c.Context.IsEvalMode() {
			conf.NoSkipConstant = true
		}
		out, err := cl.NewPackage("", apkg, conf)
		if err != nil {
			return nil, err
		}
		pkgs = append(pkgs, out)
	}
	return &Package{c.FileSet, pkgs}, nil
}
