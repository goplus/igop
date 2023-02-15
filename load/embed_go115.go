//go:build !go1.16
// +build !go1.16

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

package load

import (
	"go/ast"
	"go/build"
	"go/token"
)

func Embed(bp *build.Package, fset *token.FileSet, files []*ast.File, test bool, xtest bool) (*ast.File, error) {
	return nil, nil
}

func EmbedFiles(pkgName string, dir string, fset *token.FileSet, files []*ast.File) (*ast.File, error) {
	return nil, nil
}
