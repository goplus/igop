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

//go:generate go run ../cmd/qexp -outdir ../pkg github.com/goplus/gop/builtin
//go:generate go run ../cmd/qexp -outdir ../pkg github.com/goplus/gop/builtin/ng
//go:generate go run ../cmd/qexp -outdir ../pkg github.com/goplus/gop/builtin/iox
//go:generate go run ../cmd/qexp -outdir ../pkg github.com/qiniu/x/errors

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
	"github.com/goplus/gox"
	"github.com/goplus/igop"

	_ "github.com/goplus/igop/pkg/bufio"
	_ "github.com/goplus/igop/pkg/fmt"
	_ "github.com/goplus/igop/pkg/github.com/goplus/gop/builtin"
	_ "github.com/goplus/igop/pkg/github.com/goplus/gop/builtin/iox"
	_ "github.com/goplus/igop/pkg/github.com/goplus/gop/builtin/ng"
	_ "github.com/goplus/igop/pkg/github.com/qiniu/x/errors"
	_ "github.com/goplus/igop/pkg/io"
	_ "github.com/goplus/igop/pkg/log"
	_ "github.com/goplus/igop/pkg/math/big"
	_ "github.com/goplus/igop/pkg/math/bits"
	_ "github.com/goplus/igop/pkg/os"
	_ "github.com/goplus/igop/pkg/strconv"
	_ "github.com/goplus/igop/pkg/strings"
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
	igop.RegisterFileProcess(".gop", BuildFile)
	RegisterClassFileType(".gmx", ".spx", "github.com/goplus/spx", "math")
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
			err = fmt.Errorf("compile %v failed. %v", dir, err)
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
			err = fmt.Errorf("compile %v failed. %v", dir, err)
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
	Pkg  *gox.Package
}

func (p *Package) ToSource() ([]byte, error) {
	var buf bytes.Buffer
	if err := gox.WriteTo(&buf, p.Pkg); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (p *Package) ToAst() *goast.File {
	return gox.ASTFile(p.Pkg)
}

type Context struct {
	ctx  *igop.Context
	fset *token.FileSet
	gop  igop.Loader
}

func IsClass(ext string) (isProj bool, ok bool) {
	if cls, ok := classfile[ext]; ok {
		return ext == cls.ProjExt, true
	}
	return
}

func NewContext(ctx *igop.Context) *Context {
	if ctx.IsEvalMode() {
		ctx = igop.NewContext(0)
	}
	return &Context{ctx: ctx, fset: token.NewFileSet(), gop: igop.NewTypesLoader(ctx, 0)}
}

func isGopPackage(path string) bool {
	if pkg, ok := igop.LookupPackage(path); ok {
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

func (c *Context) ParseDir(dir string) (*Package, error) {
	pkgs, err := parser.ParseDirEx(c.fset, dir, parser.Config{
		IsClass: IsClass,
	})
	if err != nil {
		return nil, err
	}
	return c.loadPackage(dir, pkgs)
}

func (c *Context) ParseFSDir(fs parser.FileSystem, dir string) (*Package, error) {
	pkgs, err := parser.ParseFSDir(c.fset, fs, dir, parser.Config{
		IsClass: IsClass,
	})
	if err != nil {
		return nil, err
	}
	return c.loadPackage(dir, pkgs)
}

func (c *Context) ParseFile(filename string, src interface{}) (*Package, error) {
	srcDir, _ := filepath.Split(filename)
	f, err := parser.ParseFile(c.fset, filename, src, 0)
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
	return c.loadPackage(srcDir, pkgs)
}

func (c *Context) loadPackage(srcDir string, pkgs map[string]*ast.Package) (*Package, error) {
	mainPkg, ok := pkgs["main"]
	if !ok {
		return nil, fmt.Errorf("not a main package")
	}
	if f, err := igop.ParseBuiltin(c.fset, "main"); err == nil {
		mainPkg.GoFiles = map[string]*goast.File{"_igop_builtin.go": f}
	}
	conf := &cl.Config{
		WorkingDir: srcDir, TargetDir: srcDir, Fset: c.fset}
	conf.Importer = c
	conf.LookupClass = func(ext string) (c *cl.Class, ok bool) {
		c, ok = classfile[ext]
		return
	}
	if c.ctx.IsEvalMode() {
		conf.NoSkipConstant = true
	}
	out, err := cl.NewPackage("", mainPkg, conf)
	if err != nil {
		return nil, err
	}
	return &Package{c.fset, out}, nil
}
