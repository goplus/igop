//go:build go1.18
// +build go1.18

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

package export

import (
	"go/ast"
	"go/types"
)

func hasTypeParam(t types.Type) bool {
	switch t1 := t.(type) {
	case *types.TypeParam:
		return true
	case *types.Named:
		if i, ok := t.Underlying().(*types.Interface); ok {
			for n := 0; n < i.NumEmbeddeds(); n++ {
				if hasTypeParam(i.EmbeddedType(n)) {
					return true
				}
			}
		}
		return t1.TypeParams() != nil
	case *types.Signature:
		return t1.TypeParams() != nil
	case *types.Union:
		return true
	}
	return false
}

func recvHasTypeParam(expr ast.Expr) bool {
retry:
	switch v := expr.(type) {
	case *ast.IndexExpr, *ast.IndexListExpr:
		return true
	case *ast.ParenExpr:
		expr = v.X
		goto retry
	case *ast.StarExpr:
		expr = v.X
		goto retry
	}
	return false
}

func funcHasTypeParams(fn *ast.FuncDecl) bool {
	if fn.Type.TypeParams != nil {
		return true
	}
	if fn.Recv != nil && len(fn.Recv.List) == 1 && recvHasTypeParam(fn.Recv.List[0].Type) {
		return true
	}
	return false
}
