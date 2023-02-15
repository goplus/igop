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

package igop

import (
	"go/types"
	"regexp"
	"runtime"
	"strings"
	"unsafe"

	"golang.org/x/tools/go/ssa"
)

type funcinl struct {
	ones  uint32  // set to ^0 to distinguish from _func
	entry uintptr // entry of the real (the "outermost") frame
	name  string
	file  string
	line  int
}

func inlineFunc(entry uintptr) *funcinl {
	return &funcinl{ones: ^uint32(0), entry: entry}
}

func isInlineFunc(f *runtime.Func) bool {
	return (*funcinl)(unsafe.Pointer(f)).ones == ^uint32(0)
}

var (
	reFuncName = regexp.MustCompile("\\$(\\d+)")
)

func fixedFuncName(fn *ssa.Function) (string, bool) {
	name := fn.String()
	name = reFuncName.ReplaceAllString(name, ".func$1")
	if strings.HasPrefix(name, "(") {
		if pos := strings.LastIndex(name, ")"); pos != -1 {
			line := name[1:pos]
			if strings.HasPrefix(line, "*") {
				if dot := strings.LastIndex(line, "."); dot != -1 {
					line = line[1:dot+1] + "(*" + line[dot+1:] + ")"
				}
			}
			name = line + name[pos+1:]
		}
	}
	if strings.HasSuffix(name, "$bound") {
		return name[:len(name)-6] + "-fm", true
	} else if strings.HasSuffix(name, "$thunk") {
		name = name[:len(name)-6]
		if strings.HasPrefix(name, "struct{") {
			return name, true
		}
		if sig, ok := fn.Type().(*types.Signature); ok {
			if types.IsInterface(sig.Params().At(0).Type()) {
				return name, true
			}
		}
	}
	return name, false
}
