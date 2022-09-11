//go:build !go1.16
// +build !go1.16

package load

import (
	"go/ast"
	"go/build"
	"go/token"
)

func Embed(bp *build.Package, fset *token.FileSet, files []*ast.File, test bool, xtest bool) ([]byte, error) {
	return nil, nil
}
