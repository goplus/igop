package gossa

import (
	"go/ast"
	"go/parser"
	"go/token"
	"go/types"

	"golang.org/x/tools/go/ssa"
)

type Context struct {
	Loader Loader
}

func NewContext() *Context {
	return &Context{
		Loader: NewTypesLoader(),
	}
}

func (c *Context) LoadFile(filename string, src interface{}, mode ssa.BuilderMode) (*ssa.Package, error) {
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, filename, src, parser.AllErrors|parser.ParseComments)
	if err != nil {
		return nil, err
	}
	return c.LoadAstFile(fset, file, mode)
}

func (c *Context) LoadAstFile(fset *token.FileSet, file *ast.File, mode ssa.BuilderMode) (*ssa.Package, error) {
	pkg := types.NewPackage(file.Name.Name, "")
	ssapkg, _, err := BuildPackage(c.Loader, fset, pkg, []*ast.File{file}, mode)
	if err != nil {
		return nil, err
	}
	ssapkg.Build()
	return ssapkg, nil
}

func (c *Context) LoadAstPackage(fset *token.FileSet, apkg *ast.Package, mode ssa.BuilderMode) (*ssa.Package, error) {
	pkg := types.NewPackage(apkg.Name, "")
	var files []*ast.File
	for _, f := range apkg.Files {
		files = append(files, f)
	}
	ssapkg, _, err := BuildPackage(c.Loader, fset, pkg, files, mode)
	if err != nil {
		return nil, err
	}
	ssapkg.Build()
	return ssapkg, nil
}
