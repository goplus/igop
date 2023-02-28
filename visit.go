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
	"fmt"
	"go/token"
	"log"
	"reflect"
	"strings"

	"golang.org/x/tools/go/ssa"
)

const (
	fnBase = 100
)

func checkPackages(intp *Interp, pkgs []*ssa.Package) (err error) {
	if intp.ctx.Mode&DisableRecover == 0 {
		defer func() {
			if v := recover(); v != nil {
				err = v.(error)
			}
		}()
	}
	visit := visitor{
		intp: intp,
		prog: intp.mainpkg.Prog,
		pkgs: make(map[*ssa.Package]bool),
		seen: make(map[*ssa.Function]bool),
		base: fnBase,
	}
	for _, pkg := range pkgs {
		visit.pkgs[pkg] = true
	}
	visit.program()
	return
}

type visitor struct {
	intp *Interp
	prog *ssa.Program
	pkgs map[*ssa.Package]bool
	seen map[*ssa.Function]bool
	base int
}

func (visit *visitor) program() {
	chks := make(map[string]bool)
	chks[""] = true // anonymous struct embed named type
	for pkg := range visit.pkgs {
		chks[pkg.Pkg.Path()] = true
		for _, mem := range pkg.Members {
			if fn, ok := mem.(*ssa.Function); ok {
				visit.function(fn)
			}
		}
	}
	isExtern := func(typ reflect.Type) bool {
		if typ.Kind() == reflect.Ptr {
			typ = typ.Elem()
		}
		return !chks[typ.PkgPath()]
	}
	for _, T := range visit.prog.RuntimeTypes() {
		typ := visit.intp.preToType(T)
		// skip extern type
		if isExtern(typ) {
			continue
		}
		mmap := make(map[string]*ssa.Function)
		mset := visit.prog.MethodSets.MethodSet(T)
		for i, n := 0, mset.Len(); i < n; i++ {
			sel := mset.At(i)
			obj := sel.Obj()
			// skip embed extern type method
			if pkg := obj.Pkg(); pkg != nil && !chks[pkg.Path()] {
				continue
			}
			fn := visit.prog.MethodValue(sel)
			mmap[obj.Name()] = fn
			visit.function(fn)
		}
		visit.intp.msets[typ] = mmap
	}
}

func (visit *visitor) function(fn *ssa.Function) {
	if visit.seen[fn] {
		return
	}
	if hasTypeParam(fn.Type()) {
		return
	}
	visit.seen[fn] = true
	fnPath := fn.String()
	if f, ok := visit.intp.ctx.override[fnPath]; ok &&
		visit.intp.preToType(fn.Type()) == f.Type() {
		fn.Blocks = nil
		return
	}
	if fn.Blocks == nil {
		if _, ok := visit.pkgs[fn.Pkg]; ok {
			if _, ok = findExternFunc(visit.intp, fn); !ok {
				if visit.intp.ctx.Mode&EnableNoStrict != 0 {
					typ := visit.intp.preToType(fn.Type())
					numOut := typ.NumOut()
					if numOut == 0 {
						externValues[fnPath] = reflect.MakeFunc(typ, func(args []reflect.Value) (results []reflect.Value) {
							return
						})
					} else {
						externValues[fnPath] = reflect.MakeFunc(typ, func(args []reflect.Value) (results []reflect.Value) {
							results = make([]reflect.Value, numOut)
							for i := 0; i < numOut; i++ {
								results[i] = reflect.New(typ.Out(i)).Elem()
							}
							return
						})
					}
					println(fmt.Sprintf("igop warning: %v: %v missing function body", visit.intp.ctx.FileSet.Position(fn.Pos()), fnPath))
					return
				}
				panic(fmt.Errorf("%v: missing function body", visit.intp.ctx.FileSet.Position(fn.Pos())))
			}
		}
		return
	}
	if len(fn.TypeArgs()) != 0 {
		visit.intp.record.EnterInstance(fn)
		defer visit.intp.record.LeaveInstance(fn)
	}
	visit.intp.loadType(fn.Type())
	for _, alloc := range fn.Locals {
		visit.intp.loadType(alloc.Type())
		visit.intp.loadType(deref(alloc.Type()))
	}
	pfn := visit.intp.loadFunction(fn)
	for _, p := range fn.Params {
		pfn.regIndex(p)
	}
	for _, p := range fn.FreeVars {
		pfn.regIndex(p)
	}
	var buf [32]*ssa.Value // avoid alloc in common case
	for _, b := range fn.Blocks {
		Instrs := make([]func(*frame), len(b.Instrs))
		ssaInstrs := make([]ssa.Instruction, len(b.Instrs))
		var index int
		n := len(b.Instrs)
		for i := 0; i < n; i++ {
			instr := b.Instrs[i]
			ops := instr.Operands(buf[:0])
			switch instr := instr.(type) {
			case *ssa.Alloc:
				visit.intp.loadType(instr.Type())
				visit.intp.loadType(deref(instr.Type()))
			case *ssa.Next:
				// skip *ssa.opaqueType: iter
				ops = nil
			case *ssa.Extract:
				// skip
				ops = nil
			case *ssa.TypeAssert:
				visit.intp.loadType(instr.AssertedType)
			case *ssa.MakeChan:
				visit.intp.loadType(instr.Type())
			case *ssa.MakeMap:
				visit.intp.loadType(instr.Type())
			case *ssa.MakeSlice:
				visit.intp.loadType(instr.Type())
			case *ssa.SliceToArrayPointer:
				visit.intp.loadType(instr.Type())
			case *ssa.Convert:
				visit.intp.loadType(instr.Type())
			case *ssa.ChangeType:
				visit.intp.loadType(instr.Type())
			case *ssa.MakeInterface:
				visit.intp.loadType(instr.Type())
			}
			for _, op := range ops {
				switch v := (*op).(type) {
				case *ssa.Function:
					visit.function(v)
				case nil:
					// skip
				default:
					visit.intp.loadType(v.Type())
				}
			}
			ifn := makeInstr(visit.intp, pfn, instr)
			if ifn == nil {
				continue
			}
			if visit.intp.ctx.evalMode && fn.String() == "main.init" {
				if visit.intp.ctx.evalInit == nil {
					visit.intp.ctx.evalInit = make(map[string]bool)
				}
				if call, ok := instr.(*ssa.Call); ok {
					key := call.String()
					if strings.HasPrefix(key, "init#") {
						if visit.intp.ctx.evalInit[key] {
							ifn = func(fr *frame) {}
						} else {
							visit.intp.ctx.evalInit[key] = true
						}
					}
				}
			}
			if visit.intp.ctx.evalCallFn != nil {
				if call, ok := instr.(*ssa.Call); ok {
					ir := pfn.regIndex(call)
					results := call.Call.Signature().Results()
					ofn := ifn
					switch results.Len() {
					case 0:
						ifn = func(fr *frame) {
							ofn(fr)
							visit.intp.ctx.evalCallFn(visit.intp, call)
						}
					case 1:
						ifn = func(fr *frame) {
							ofn(fr)
							visit.intp.ctx.evalCallFn(visit.intp, call, fr.reg(ir))
						}
					default:
						ifn = func(fr *frame) {
							ofn(fr)
							r := fr.reg(ir).(tuple)
							visit.intp.ctx.evalCallFn(visit.intp, call, r...)
						}
					}
				}
			}
			if visit.intp.ctx.Mode&EnableTracing != 0 {
				ofn := ifn
				ifn = func(fr *frame) {
					if v, ok := instr.(ssa.Value); ok {
						log.Printf("\t%-20T %v = %-40v\t%v\n", instr, v.Name(), instr, v.Type())
					} else {
						log.Printf("\t%-20T %v\n", instr, instr)
					}
					ofn(fr)
				}
				if index == 0 {
					ofn := ifn
					bi := b.Index
					common := b.Comment
					ifn = func(fr *frame) {
						log.Printf(".%v %v\n", bi, common)
						ofn(fr)
					}
				}
				if index == 0 && b.Index == 0 {
					ofn := ifn
					ifn = func(fr *frame) {
						log.Printf("Entering %v%v.", fr.pfn.Fn, loc(fr.pfn.Interp.ctx.FileSet, fr.pfn.Fn.Pos()))
						ofn(fr)
					}
				}
				if _, ok := instr.(*ssa.Return); ok {
					ofn := ifn
					ifn = func(fr *frame) {
						ofn(fr)
						var caller ssa.Instruction
						if fr.caller != nil {
							caller = fr.caller.pfn.InstrForPC(fr.caller.ipc - 1)
						}
						if caller == nil {
							log.Printf("Leaving %v.\n", fr.pfn.Fn)
						} else {
							log.Printf("Leaving %v, resuming %v call %v%v.\n",
								fr.pfn.Fn, fr.caller.pfn.Fn, caller, loc(fr.pfn.Interp.ctx.FileSet, caller.Pos()))
						}
					}
				}
			}
			Instrs[index] = ifn
			ssaInstrs[index] = instr
			index++
		}
		offset := len(pfn.Instrs)
		pfn.Blocks = append(pfn.Blocks, offset)
		pfn.Instrs = append(pfn.Instrs, Instrs[:index]...)
		pfn.ssaInstrs = append(pfn.ssaInstrs, ssaInstrs[:index]...)
		if b == fn.Recover && visit.intp.ctx.Mode&DisableRecover == 0 {
			pfn.Recover = pfn.Instrs[offset:]
		}
	}
	pfn.base = visit.base
	visit.base += len(pfn.ssaInstrs) + 2
	pfn.initPool()
}

func loc(fset *token.FileSet, pos token.Pos) string {
	if pos == token.NoPos {
		return ""
	}
	return " at " + fset.Position(pos).String()
}
