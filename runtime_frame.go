//go:build !go1.23
// +build !go1.23

/*
 * Copyright (c) 2025 The GoPlus Authors (goplus.org). All rights reserved.
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

package ixgo

import (
	"runtime"

	"golang.org/x/tools/go/ssa"
)

/*
	type Frames struct {
		// callers is a slice of PCs that have not yet been expanded to frames.
		callers []uintptr

		// frames is a slice of Frames that have yet to be returned.
		frames     []Frame
		frameStore [2]Frame
	}
*/

type runtimeFrames struct {
	callers    []uintptr
	frames     []runtime.Frame
	frameStore [2]runtime.Frame
}

func makeDefer(interp *Interp, pfn *function, instr *ssa.Defer) func(fr *frame) {
	iv, ia, ib := getCallIndex(pfn, &instr.Call)
	return func(fr *frame) {
		fn, args := interp.prepareCall(fr, &instr.Call, iv, ia, ib)
		fr._defer = &_defer{
			fn:      fn,
			args:    args,
			ssaArgs: instr.Call.Args,
			tail:    fr._defer,
		}
	}
}
