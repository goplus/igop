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
	"fmt"
	"go/ast"
	"go/token"
	"strings"
)

type LinkSym struct {
	Kind     ast.ObjKind
	PkgPath  string
	Name     string
	Linkname Linkname
}

func (l *LinkSym) String() string {
	return l.PkgPath + "." + l.Name + "->" + l.Linkname.PkgPath + "." + l.Linkname.Name
}

type Linkname struct {
	PkgPath string
	Name    string
	Recv    string
	Method  string
}

// ParseLinkname parse ast files go:linkname
// //go:linkname <localname> <importpath>.<name>
// //go:linkname <localname> <importpath>.<type>.<name>
// //go:linkname <localname> <importpath>.<(*type)>.<name>
// //go:linkname <localname> linkname indicate by runtime package
func ParseLinkname(fset *token.FileSet, pkgPath string, files []*ast.File) ([]*LinkSym, error) {
	var links []*LinkSym
	for _, file := range files {
		var hasUnsafe bool
		for _, imp := range file.Imports {
			if imp.Path.Value == `"unsafe"` {
				hasUnsafe = true
			}
		}
		for _, cg := range file.Comments {
			for _, c := range cg.List {
				link, err := parseLinknameComment(pkgPath, file, c, hasUnsafe)
				if err != nil {
					return nil, fmt.Errorf("%s: %w", fset.Position(c.Pos()), err)
				}
				if link != nil {
					links = append(links, link)
				}
			}
		}
	}
	return links, nil
}

func parseLinknameComment(pkgPath string, file *ast.File, comment *ast.Comment, hasUnsafe bool) (*LinkSym, error) {
	if !strings.HasPrefix(comment.Text, "//go:linkname ") {
		return nil, nil
	}
	if !hasUnsafe {
		return nil, fmt.Errorf(`//go:linkname only allowed in Go files that import "unsafe"`)
	}
	fields := strings.Fields(comment.Text)
	if n := len(fields); n != 3 {
		if n == 2 {
			// //go:linkname <localname>
			return nil, nil
		}
		return nil, fmt.Errorf(`usage: //go:linkname localname [linkname]`)
	}

	localName := fields[1]
	linkPkg, linkName := "", fields[2]
	if pos := strings.LastIndexByte(linkName, '/'); pos != -1 {
		if idx := strings.IndexByte(linkName[pos+1:], '.'); idx != -1 {
			linkPkg, linkName = linkName[0:pos+idx+1], linkName[pos+idx+2:]
		}
	} else if idx := strings.IndexByte(linkName, '.'); idx != -1 {
		linkPkg, linkName = linkName[0:idx], linkName[idx+1:]
	}

	obj := file.Scope.Lookup(localName)
	if obj == nil || (obj.Kind != ast.Fun && obj.Kind != ast.Var) {
		return nil, fmt.Errorf("//go:linkname must refer to declared function or variable")
	}

	if pkgPath == linkPkg && localName == linkName {
		return nil, nil
	}

	var recv, method string
	pos := strings.IndexByte(linkName, '.')
	if pos != -1 {
		recv, method = linkName[:pos], linkName[pos+1:]
		size := len(recv)
		if size > 2 && recv[0] == '(' && recv[size-1] == ')' {
			recv = recv[1 : size-1]
		}
	}
	return &LinkSym{
		Kind:     obj.Kind,
		PkgPath:  pkgPath,
		Name:     localName,
		Linkname: Linkname{PkgPath: linkPkg, Name: linkName, Recv: recv, Method: method},
	}, nil
}
