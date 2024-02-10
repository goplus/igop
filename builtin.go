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
	"bytes"
	"fmt"
	"go/ast"
	"go/token"
	"go/types"
	"reflect"
	"strings"
	"unsafe"

	"golang.org/x/tools/go/ast/astutil"
	"golang.org/x/tools/go/ssa"
)

// callBuiltin interprets a call to builtin fn with arguments args,
// returning its result.
func (inter *Interp) callBuiltin(caller *frame, fn *ssa.Builtin, args []value, ssaArgs []ssa.Value) value {
	switch fnName := fn.Name(); fnName {
	case "append":
		if len(args) == 1 {
			return args[0]
		}
		v0 := reflect.ValueOf(args[0])
		v1 := reflect.ValueOf(args[1])
		// append([]byte, ...string) []byte
		if v1.Kind() == reflect.String {
			v1 = reflect.ValueOf([]byte(v1.String()))
		}
		i0 := v0.Len()
		i1 := v1.Len()
		if i0+i1 < i0 {
			panic(runtimeError(errAppendOutOfRange))
		}
		return reflect.AppendSlice(v0, v1).Interface()

	case "copy": // copy([]T, []T) int or copy([]byte, string) int
		return reflect.Copy(reflect.ValueOf(args[0]), reflect.ValueOf(args[1]))

	case "close": // close(chan T)
		reflect.ValueOf(args[0]).Close()
		return nil

	case "delete": // delete(map[K]value, K)
		reflect.ValueOf(args[0]).SetMapIndex(reflect.ValueOf(args[1]), reflect.Value{})
		return nil

	case "print", "println": // print(any, ...)
		ln := fn.Name() == "println"
		var buf bytes.Buffer
		for i, arg := range args {
			if i > 0 && ln {
				buf.WriteRune(' ')
			}
			if len(ssaArgs) > i {
				typ := inter.toType(ssaArgs[i].Type())
				if typ.Kind() == reflect.Interface {
					writeinterface(&buf, arg)
					continue
				}
			}
			writevalue(&buf, arg, inter.ctx.Mode&EnablePrintAny != 0)
		}
		if ln {
			buf.WriteRune('\n')
		}
		inter.ctx.writeOutput(buf.Bytes())
		return nil

	case "len":
		return reflect.ValueOf(args[0]).Len()

	case "cap":
		return reflect.ValueOf(args[0]).Cap()

	case "real":
		c := reflect.ValueOf(args[0])
		switch c.Kind() {
		case reflect.Complex64:
			return real(complex64(c.Complex()))
		case reflect.Complex128:
			return real(c.Complex())
		default:
			panic(fmt.Sprintf("real: illegal operand: %T", c))
		}

	case "imag":
		c := reflect.ValueOf(args[0])
		switch c.Kind() {
		case reflect.Complex64:
			return imag(complex64(c.Complex()))
		case reflect.Complex128:
			return imag(c.Complex())
		default:
			panic(fmt.Sprintf("imag: illegal operand: %T", c))
		}

	case "complex":
		r := reflect.ValueOf(args[0])
		i := reflect.ValueOf(args[1])
		switch r.Kind() {
		case reflect.Float32:
			return complex(float32(r.Float()), float32(i.Float()))
		case reflect.Float64:
			return complex(r.Float(), i.Float())
		default:
			panic(fmt.Sprintf("complex: illegal operand: %v", r.Kind()))
		}

	case "panic":
		// ssa.Panic handles most cases; this is only for "go
		// panic" or "defer panic".
		panic(PanicError{stack: debugStack(caller), Value: args[0]})

	case "recover":
		return doRecover(caller)

	case "ssa:wrapnilchk":
		recv := args[0]
		if reflect.ValueOf(recv).IsNil() {
			recvType := args[1]
			methodName := args[2]
			var info value
			if s, ok := recvType.(string); ok && strings.HasPrefix(s, "main.") {
				info = s[5:]
			} else {
				info = recvType
			}
			panic(plainError(fmt.Sprintf("value method %s.%s called using nil *%s pointer",
				recvType, methodName, info)))
		}
		return recv

	case "Add":
		ptr := args[0].(unsafe.Pointer)
		length := asInt(args[1])
		return unsafe.Pointer(uintptr(ptr) + uintptr(length))
	case "Slice":
		//func Slice(ptr *ArbitraryType, len IntegerType) []ArbitraryType
		//(*[len]ArbitraryType)(unsafe.Pointer(ptr))[:]
		ptr := reflect.ValueOf(args[0])
		length := asInt(args[1])
		etyp := ptr.Type().Elem()
		if ptr.IsNil() {
			if length == 0 {
				return reflect.New(reflect.SliceOf(etyp)).Elem().Interface()
			}
			panic(runtimeError("unsafe.Slice: ptr is nil and len is not zero"))
		}
		mem, overflow := mulUintptr(etyp.Size(), uintptr(length))
		if overflow || mem > -uintptr(ptr.Pointer()) {
			panic(runtimeError("unsafe.Slice: len out of range"))
		}
		typ := reflect.ArrayOf(length, etyp)
		v := reflect.NewAt(typ, unsafe.Pointer(ptr.Pointer()))
		return v.Elem().Slice(0, length).Interface()
	case "SliceData":
		// SliceData returns a pointer to the underlying array of the argument
		// slice.
		//   - If cap(slice) > 0, SliceData returns &slice[:1][0].
		//   - If slice == nil, SliceData returns nil.
		//   - Otherwise, SliceData returns a non-nil pointer to an
		//     unspecified memory address.
		// func SliceData(slice []ArbitraryType) *ArbitraryType
		v := reflect.ValueOf(args[0])
		if v.Cap() > 0 {
			return v.Slice(0, 1).Index(0).Addr().Interface()
		} else if v.IsNil() {
			return reflect.Zero(reflect.PtrTo(v.Type().Elem())).Interface()
		} else {
			return reflect.New(v.Type().Elem()).Interface()
		}
	case "String":
		// String returns a string value whose underlying bytes
		// start at ptr and whose length is len.
		//
		// The len argument must be of integer type or an untyped constant.
		// A constant len argument must be non-negative and representable by a value of type int;
		// if it is an untyped constant it is given type int.
		// At run time, if len is negative, or if ptr is nil and len is not zero,
		// a run-time panic occurs.
		//
		// Since Go strings are immutable, the bytes passed to String
		// must not be modified afterwards.
		// func String(ptr *byte, len IntegerType) string
		ptr := args[0].(*byte)
		length := asInt(args[1])
		if ptr == nil {
			if length == 0 {
				return ""
			}
			panic(runtimeError("unsafe.String: ptr is nil and len is not zero"))
		}
		mem, overflow := mulUintptr(1, uintptr(length))
		if overflow || mem > -uintptr(unsafe.Pointer(ptr)) {
			panic(runtimeError("unsafe.String: len out of range"))
		}
		sh := reflect.StringHeader{Data: uintptr(unsafe.Pointer(ptr)), Len: length}
		return *(*string)(unsafe.Pointer(&sh))
	case "StringData":
		// StringData returns a pointer to the underlying bytes of str.
		// For an empty string the return value is unspecified, and may be nil.
		//
		// Since Go strings are immutable, the bytes returned by StringData
		// must not be modified.
		// func StringData(str string) *byte
		s := args[0].(string)
		data := (*reflect.StringHeader)(unsafe.Pointer(&s)).Data
		return (*byte)(unsafe.Pointer(data))
	default:
		panic("unknown built-in: " + fnName)
	}
}

// callBuiltinDiscardsResult interprets a call to builtin fn with arguments args,
// discards its result.
func (inter *Interp) callBuiltinDiscardsResult(caller *frame, fn *ssa.Builtin, args []value, ssaArgs []ssa.Value) {
	switch fnName := fn.Name(); fnName {
	case "append":
		panic("discards result of " + fnName)

	case "copy": // copy([]T, []T) int or copy([]byte, string) int
		reflect.Copy(reflect.ValueOf(args[0]), reflect.ValueOf(args[1]))

	case "close": // close(chan T)
		reflect.ValueOf(args[0]).Close()

	case "delete": // delete(map[K]value, K)
		reflect.ValueOf(args[0]).SetMapIndex(reflect.ValueOf(args[1]), reflect.Value{})

	case "print", "println": // print(any, ...)
		ln := fn.Name() == "println"
		var buf bytes.Buffer
		for i, arg := range args {
			if i > 0 && ln {
				buf.WriteRune(' ')
			}
			if len(ssaArgs) > i {
				typ := inter.toType(ssaArgs[i].Type())
				if typ.Kind() == reflect.Interface {
					writeinterface(&buf, arg)
					continue
				}
			}
			writevalue(&buf, arg, inter.ctx.Mode&EnablePrintAny != 0)
		}
		if ln {
			buf.WriteRune('\n')
		}
		inter.ctx.writeOutput(buf.Bytes())

	case "len":
		panic("discards result of " + fnName)

	case "cap":
		panic("discards result of " + fnName)

	case "real":
		panic("discards result of " + fnName)

	case "imag":
		panic("discards result of " + fnName)

	case "complex":
		panic("discards result of " + fnName)

	case "panic":
		// ssa.Panic handles most cases; this is only for "go
		// panic" or "defer panic".
		panic(PanicError{stack: debugStack(caller), Value: args[0]})

	case "recover":
		doRecover(caller)

	case "ssa:wrapnilchk":
		recv := args[0]
		if reflect.ValueOf(recv).IsNil() {
			recvType := args[1]
			methodName := args[2]
			var info value
			if s, ok := recvType.(string); ok && strings.HasPrefix(s, "main.") {
				info = s[5:]
			} else {
				info = recvType
			}
			panic(plainError(fmt.Sprintf("value method %s.%s called using nil *%s pointer",
				recvType, methodName, info)))
		}

	case "Add":
		panic("discards result of " + fnName)

	case "Slice":
		//func Slice(ptr *ArbitraryType, len IntegerType) []ArbitraryType
		//(*[len]ArbitraryType)(unsafe.Pointer(ptr))[:]
		panic("discards result of " + fnName)

	case "SliceData":
		panic("discards result of " + fnName)

	case "String":
		panic("discards result of " + fnName)

	case "StringData":
		panic("discards result of " + fnName)

	default:
		panic("unknown built-in: " + fnName)
	}
}

// callBuiltin interprets a call to builtin fn with arguments args,
// returning its result.
func (interp *Interp) callBuiltinByStack(caller *frame, fn string, ssaArgs []ssa.Value, ir register, ia []register) {
	switch fn {
	case "append":
		if len(ia) == 1 {
			caller.copyReg(ir, ia[0])
			return
		}
		arg0 := caller.reg(ia[0])
		arg1 := caller.reg(ia[1])
		v0 := reflect.ValueOf(arg0)
		v1 := reflect.ValueOf(arg1)
		// append([]byte, ...string) []byte
		if v1.Kind() == reflect.String {
			v1 = reflect.ValueOf([]byte(v1.String()))
		}
		i0 := v0.Len()
		i1 := v1.Len()
		if i0+i1 < i0 {
			panic(runtimeError(errAppendOutOfRange))
		}
		caller.setReg(ir, reflect.AppendSlice(v0, v1).Interface())

	case "copy": // copy([]T, []T) int or copy([]byte, string) int
		arg0 := caller.reg(ia[0])
		arg1 := caller.reg(ia[1])
		caller.setReg(ir, reflect.Copy(reflect.ValueOf(arg0), reflect.ValueOf(arg1)))

	case "close": // close(chan T)
		arg0 := caller.reg(ia[0])
		reflect.ValueOf(arg0).Close()

	case "delete": // delete(map[K]value, K)
		arg0 := caller.reg(ia[0])
		arg1 := caller.reg(ia[1])
		reflect.ValueOf(arg0).SetMapIndex(reflect.ValueOf(arg1), reflect.Value{})

	case "print", "println": // print(any, ...)
		ln := fn == "println"
		var buf bytes.Buffer
		for i := 0; i < len(ia); i++ {
			arg := caller.reg(ia[i])
			if i > 0 && ln {
				buf.WriteRune(' ')
			}
			if len(ssaArgs) > i {
				typ := interp.toType(ssaArgs[i].Type())
				if typ.Kind() == reflect.Interface {
					writeinterface(&buf, arg)
					continue
				}
			}
			writevalue(&buf, arg, interp.ctx.Mode&EnablePrintAny != 0)
		}
		if ln {
			buf.WriteRune('\n')
		}
		interp.ctx.writeOutput(buf.Bytes())

	case "len":
		arg0 := caller.reg(ia[0])
		caller.setReg(ir, reflect.ValueOf(arg0).Len())

	case "cap":
		arg0 := caller.reg(ia[0])
		caller.setReg(ir, reflect.ValueOf(arg0).Cap())

	case "real":
		arg0 := caller.reg(ia[0])
		c := reflect.ValueOf(arg0)
		switch c.Kind() {
		case reflect.Complex64:
			caller.setReg(ir, real(complex64(c.Complex())))
		case reflect.Complex128:
			caller.setReg(ir, real(c.Complex()))
		default:
			panic(fmt.Sprintf("real: illegal operand: %T", c))
		}

	case "imag":
		arg0 := caller.reg(ia[0])
		c := reflect.ValueOf(arg0)
		switch c.Kind() {
		case reflect.Complex64:
			caller.setReg(ir, imag(complex64(c.Complex())))
		case reflect.Complex128:
			caller.setReg(ir, imag(c.Complex()))
		default:
			panic(fmt.Sprintf("imag: illegal operand: %T", c))
		}

	case "complex":
		arg0 := caller.reg(ia[0])
		arg1 := caller.reg(ia[1])
		r := reflect.ValueOf(arg0)
		i := reflect.ValueOf(arg1)
		switch r.Kind() {
		case reflect.Float32:
			caller.setReg(ir, complex(float32(r.Float()), float32(i.Float())))
		case reflect.Float64:
			caller.setReg(ir, complex(r.Float(), i.Float()))
		default:
			panic(fmt.Sprintf("complex: illegal operand: %v", r.Kind()))
		}

	case "panic":
		// ssa.Panic handles most cases; this is only for "go
		// panic" or "defer panic".
		arg0 := caller.reg(ia[0])
		panic(PanicError{stack: debugStack(caller), Value: arg0})

	case "recover":
		caller.setReg(ir, doRecover(caller))

	case "ssa:wrapnilchk":
		recv := caller.reg(ia[0])
		if reflect.ValueOf(recv).IsNil() {
			recvType := caller.reg(ia[1])
			methodName := caller.reg(ia[2])
			var info value
			if s, ok := recvType.(string); ok && strings.HasPrefix(s, "main.") {
				info = s[5:]
			} else {
				info = recvType
			}
			panic(plainError(fmt.Sprintf("value method %s.%s called using nil *%s pointer",
				recvType, methodName, info)))
		}
		caller.setReg(ir, recv)

	case "Add":
		arg0 := caller.reg(ia[0])
		arg1 := caller.reg(ia[1])
		ptr := arg0.(unsafe.Pointer)
		length := asInt(arg1)
		caller.setReg(ir, unsafe.Pointer(uintptr(ptr)+uintptr(length)))
	case "Slice":
		//func Slice(ptr *ArbitraryType, len IntegerType) []ArbitraryType
		//(*[len]ArbitraryType)(unsafe.Pointer(ptr))[:]
		arg0 := caller.reg(ia[0])
		arg1 := caller.reg(ia[1])
		ptr := reflect.ValueOf(arg0)
		etyp := ptr.Type().Elem()
		length := asInt(arg1)
		if ptr.IsNil() {
			if length == 0 {
				caller.setReg(ir, reflect.New(reflect.SliceOf(etyp)).Elem().Interface())
				return
			}
			panic(runtimeError("unsafe.Slice: ptr is nil and len is not zero"))
		}
		mem, overflow := mulUintptr(etyp.Size(), uintptr(length))
		if overflow || mem > -uintptr(ptr.Pointer()) {
			panic(runtimeError("unsafe.Slice: len out of range"))
		}
		typ := reflect.ArrayOf(length, etyp)
		v := reflect.NewAt(typ, unsafe.Pointer(ptr.Pointer()))
		caller.setReg(ir, v.Elem().Slice(0, length).Interface())
	case "SliceData":
		// SliceData returns a pointer to the underlying array of the argument
		// slice.
		//   - If cap(slice) > 0, SliceData returns &slice[:1][0].
		//   - If slice == nil, SliceData returns nil.
		//   - Otherwise, SliceData returns a non-nil pointer to an
		//     unspecified memory address.
		// func SliceData(slice []ArbitraryType) *ArbitraryType
		arg0 := caller.reg(ia[0])
		v := reflect.ValueOf(arg0)
		if v.Cap() > 0 {
			caller.setReg(ir, v.Slice(0, 1).Index(0).Addr().Interface())
		} else if v.IsNil() {
			caller.setReg(ir, reflect.Zero(reflect.PtrTo(v.Type().Elem())).Interface())
		} else {
			caller.setReg(ir, reflect.New(v.Type().Elem()).Interface())
		}
	case "String":
		// String returns a string value whose underlying bytes
		// start at ptr and whose length is len.
		//
		// The len argument must be of integer type or an untyped constant.
		// A constant len argument must be non-negative and representable by a value of type int;
		// if it is an untyped constant it is given type int.
		// At run time, if len is negative, or if ptr is nil and len is not zero,
		// a run-time panic occurs.
		//
		// Since Go strings are immutable, the bytes passed to String
		// must not be modified afterwards.
		// func String(ptr *byte, len IntegerType) string
		ptr := caller.reg(ia[0]).(*byte)
		length := asInt(caller.reg(ia[1]))
		if ptr == nil {
			if length == 0 {
				caller.setReg(ir, "")
				return
			}
			panic(runtimeError("unsafe.String: ptr is nil and len is not zero"))
		}
		mem, overflow := mulUintptr(1, uintptr(length))
		if overflow || mem > -uintptr(unsafe.Pointer(ptr)) {
			panic(runtimeError("unsafe.String: len out of range"))
		}
		sh := reflect.StringHeader{Data: uintptr(unsafe.Pointer(ptr)), Len: length}
		caller.setReg(ir, *(*string)(unsafe.Pointer(&sh)))
	case "StringData":
		// StringData returns a pointer to the underlying bytes of str.
		// For an empty string the return value is unspecified, and may be nil.
		//
		// Since Go strings are immutable, the bytes returned by StringData
		// must not be modified.
		// func StringData(str string) *byte
		s := caller.string(ia[0])
		data := (*reflect.StringHeader)(unsafe.Pointer(&s)).Data
		caller.setReg(ir, (*byte)(unsafe.Pointer(data)))
	case "Sizeof": // instance of generic function
		typ := reflect.TypeOf(caller.reg(ia[0]))
		caller.setReg(ir, uintptr(typ.Size()))
	case "Alignof": // instance of generic function
		typ := reflect.TypeOf(caller.reg(ia[0]))
		caller.setReg(ir, uintptr(typ.Align()))
	case "Offsetof": // instance of generic function
		offset, err := builtinOffsetof(caller.pfn, caller.ipc-1)
		if err != nil {
			panic(err)
		}
		caller.setReg(ir, uintptr(offset))
	case "clear":
		arg0 := caller.reg(ia[0])
		valueClear(reflect.ValueOf(arg0))
	case "max":
		v := reflect.ValueOf(caller.reg(ia[0]))
		for _, i := range ia {
			arg := reflect.ValueOf(caller.reg(i))
			if i > 0 {
				switch v.Kind() {
				case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
					if v.Int() < arg.Int() {
						v = arg
					}
				case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
					if v.Uint() < arg.Uint() {
						v = arg
					}
				case reflect.Float32, reflect.Float64:
					if v.Float() < arg.Float() {
						v = arg
					}
				case reflect.String:
					if v.String() < arg.String() {
						v = arg
					}
				}
			}
		}
		caller.setReg(ir, v.Interface())
	case "min":
		v := reflect.ValueOf(caller.reg(ia[0]))
		for _, i := range ia {
			arg := reflect.ValueOf(caller.reg(i))
			if i > 0 {
				switch v.Kind() {
				case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
					if v.Int() > arg.Int() {
						v = arg
					}
				case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
					if v.Uint() > arg.Uint() {
						v = arg
					}
				case reflect.Float32, reflect.Float64:
					if v.Float() > arg.Float() {
						v = arg
					}
				case reflect.String:
					if v.String() > arg.String() {
						v = arg
					}
				}
			}
		}
		caller.setReg(ir, v.Interface())
	default:
		panic("unknown built-in: " + fn)
	}
}

func valueClear(v reflect.Value) {
	reflect.ValueOf(v).MethodByName("Clear").Call(nil)
}

const ptrSize = 4 << (^uintptr(0) >> 63)

const maxUintptr = ^uintptr(0)

// mulUintptr returns a * b and whether the multiplication overflowed.
// On supported platforms this is an intrinsic lowered by the compiler.
func mulUintptr(a, b uintptr) (uintptr, bool) {
	if a|b < 1<<(4*ptrSize) || a == 0 {
		return a * b, false
	}
	overflow := b > maxUintptr/a
	return a * b, overflow
}

func builtinOffsetof(pfn *function, pc int) (int64, error) {
	pos := pfn.ssaInstrs[pc].Pos()
	paths, info, ok := pathEnclosingInterval(pfn.Interp.ctx, pos)
	if !ok {
		return -1, plainError("unsafe.Offsetof not found code")
	}
	call, ok := paths[0].(*ast.CallExpr)
	if !ok {
		return -1, plainError("unsafe.Offsetof not found call")
	}
	selx, ok := call.Args[0].(*ast.SelectorExpr)
	if !ok {
		return -1, plainError("unsafe.Offsetof not found selector expr")
	}
	sel, _ := info.Selections[selx]
	if sel == nil {
		return -1, plainError("unsafe.Offsetof not found selector type")
	}
	instrs, found := foundFieldAddr(pfn, pc)
	if !found || len(sel.Index()) > len(instrs) {
		return -1, plainError("unsafe.Offsetof not found FieldAddr instr")
	}
	instr := instrs[len(sel.Index())-1]
	return selOffsetof(pfn.Interp.ctx.sizes, instr.X.Type().Underlying().(*types.Pointer).Elem().Underlying().(*types.Struct), sel.Index(), selx.Sel.Name)
}

func foundFieldAddr(pfn *function, pc int) (instrs []*ssa.FieldAddr, found bool) {
	for pc > 0 {
		pc--
		if fd, ok := pfn.ssaInstrs[pc].(*ssa.FieldAddr); ok {
			found = true
			instrs = append(instrs, fd)
		} else if found {
			return
		}
	}
	return
}

func pathEnclosingInterval(ctx *Context, pos token.Pos) (path []ast.Node, info *types.Info, exact bool) {
	for _, sp := range ctx.pkgs {
		for _, file := range sp.Files {
			if pos >= file.Pos() && pos < file.End() {
				path, exact = astutil.PathEnclosingInterval(file, pos, pos)
				if exact {
					info = sp.Info
					return
				}
			}
		}
	}
	return
}

func selOffsetof(sizes types.Sizes, typ types.Type, index []int, sel string) (int64, error) {
	var o int64
	var indirectType int
	for n, i := range index {
		if n > 0 {
			if t, ok := typ.(*types.Pointer); ok {
				typ = t.Elem()
				indirectType = n
			}
			if t, ok := typ.(*types.Named); ok {
				typ = t.Underlying()
			}
		}
		s := typ.(*types.Struct)
		o += structOffsetsof(sizes, s)[i]
		typ = s.Field(i).Type()
	}
	if indirectType > 0 {
		return -1, fmt.Errorf("invalid argument: field %v is embedded via a pointer", sel)
	}
	return o, nil
}

func structOffsetsof(sizes types.Sizes, t *types.Struct) []int64 {
	var fields []*types.Var
	for i := 0; i < t.NumFields(); i++ {
		fields = append(fields, t.Field(i))
	}
	return sizes.Offsetsof(fields)
}
