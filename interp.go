// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package ssa/interp defines an interpreter for the SSA
// representation of Go programs.
//
// This interpreter is provided as an adjunct for testing the SSA
// construction algorithm.  Its purpose is to provide a minimal
// metacircular implementation of the dynamic semantics of each SSA
// instruction.  It is not, and will never be, a production-quality Go
// interpreter.
//
// The following is a partial list of Go features that are currently
// unsupported or incomplete in the interpreter.
//
// * Unsafe operations, including all uses of unsafe.Pointer, are
// impossible to support given the "boxed" value representation we
// have chosen.
//
// * The reflect package is only partially implemented.
//
// * The "testing" package is no longer supported because it
// depends on low-level details that change too often.
//
// * "sync/atomic" operations are not atomic due to the "boxed" value
// representation: it is not possible to read, modify and write an
// interface value atomically. As a consequence, Mutexes are currently
// broken.
//
// * recover is only partially implemented.  Also, the interpreter
// makes no attempt to distinguish target panics from interpreter
// crashes.
//
// * the sizes of the int, uint and uintptr types in the target
// program are assumed to be the same as those of the interpreter
// itself.
//
// * all values occupy space, even those of types defined by the spec
// to have zero size, e.g. struct{}.  This can cause asymptotic
// performance degradation.
//
// * os.Exit is implemented using panic, causing deferred functions to
// run.
package gossa

import (
	"fmt"
	"go/token"
	"go/types"
	"os"
	"reflect"
	"runtime"
	"strings"
	"sync"
	"sync/atomic"
	"unsafe"

	"github.com/goplus/reflectx"
	"github.com/goplus/xtypes"
	"golang.org/x/tools/go/ssa"
)

type continuation int

const (
	kNext continuation = iota
	kReturn
	kJump
)

type plainError string

func (e plainError) Error() string {
	return string(e)
}

type runtimeError string

func (e runtimeError) RuntimeError() {}

func (e runtimeError) Error() string {
	return "runtime error: " + string(e)
}

// State shared between all interpreted goroutines.
type Interp struct {
	prog        *ssa.Program        // the SSA program
	mainpkg     *ssa.Package        // the SSA main package
	globals     map[ssa.Value]value // addresses of global variables (immutable)
	mode        Mode                // interpreter options
	sizes       types.Sizes         // the effective type-sizing function
	goroutines  int32               // atomically updated
	ctx         xtypes.Context
	types       map[types.Type]reflect.Type
	caller      *frame
	inst        Loader
	record      *TypesRecord
	typesMutex  sync.RWMutex
	callerMutex sync.RWMutex
}

func (i *Interp) findType(t reflect.Type) (types.Type, bool) {
	i.typesMutex.Lock()
	defer i.typesMutex.Unlock()
	return i.record.LookupTypes(t)
	// typ, ok := i.inst.rcache[t]
	// if !ok {
	// 	log.Println("~~~~~", typ)
	// }
	// return typ, ok

	// return i.inst.ToType(t), true
}

// func (i *Interp) findTypeHelper(t reflect.Type) (types.Type, bool) {
// 	if rt, ok := i.inst.rcache[t]; ok {
// 		return rt, true
// 	}
// 	for k, v := range i.types {
// 		if v == t {
// 			return k, true
// 		}
// 	}
// 	if t.Kind() == reflect.Ptr {
// 		if typ, ok := i.findTypeHelper(t.Elem()); ok {
// 			return types.NewPointer(typ), true
// 		}
// 	}
// 	return nil, false
// }

func (i *Interp) FindMethod(mtyp reflect.Type, fn *types.Func) func([]reflect.Value) []reflect.Value {
	typ := fn.Type().(*types.Signature).Recv().Type()
	if f := i.prog.LookupMethod(typ, fn.Pkg(), fn.Name()); f != nil {
		return func(args []reflect.Value) []reflect.Value {
			iargs := make([]value, len(args))
			for i := 0; i < len(args); i++ {
				iargs[i] = args[i].Interface()
			}
			i.callerMutex.RLock()
			caller := i.caller
			i.callerMutex.RUnlock()
			r := call(i, caller, token.NoPos, f, iargs, nil)
			switch mtyp.NumOut() {
			case 0:
				return nil
			case 1:
				if r == nil {
					return []reflect.Value{reflect.New(mtyp.Out(0)).Elem()}
				} else {
					return []reflect.Value{reflect.ValueOf(r)}
				}
			default:
				v, ok := r.(tuple)
				if !ok {
					panic(fmt.Errorf("result type must tuple: %T", v))
				}
				res := make([]reflect.Value, len(v))
				for j := 0; j < len(v); j++ {
					if v[j] == nil {
						res[j] = reflect.New(mtyp.Out(j)).Elem()
					} else {
						res[j] = reflect.ValueOf(v[j])
					}
				}
				return res
			}
		}
	}
	if v, ok := externValues[fn.FullName()]; ok && v.Kind() == reflect.Func {
		return func(args []reflect.Value) []reflect.Value {
			return v.Call(args)
		}
	}
	panic(fmt.Sprintf("Not found func %v", fn))
	return nil
}

func ptrCount(s string) (string, int) {
	for i, v := range s {
		if v != '*' {
			return s[i:], i
		}
	}
	panic("failed ptrCount " + s)
}

func isUntyped(typ types.Type) bool {
	t, ok := typ.Underlying().(*types.Basic)
	return ok && t.Info()&types.IsUntyped != 0
}

func (i *Interp) toType(typ types.Type) reflect.Type {
	i.typesMutex.Lock()
	defer i.typesMutex.Unlock()
	return i.record.ToType(typ)
}

func (fr *frame) toFunc(typ reflect.Type, fn value) reflect.Value {
	return reflect.MakeFunc(typ, func(args []reflect.Value) []reflect.Value {
		iargs := make([]value, len(args))
		for i := 0; i < len(args); i++ {
			iargs[i] = args[i].Interface()
		}
		r := call(fr.i, fr, token.NoPos, fn, iargs, nil)
		if v, ok := r.(tuple); ok {
			res := make([]reflect.Value, len(v))
			for i := 0; i < len(v); i++ {
				if v[i] == nil {
					res[i] = reflect.New(typ.Out(i)).Elem()
				} else {
					res[i] = reflect.ValueOf(v[i])
				}
			}
			return res
		} else if typ.NumOut() == 1 {
			if r != nil {
				return []reflect.Value{reflect.ValueOf(r)}
			} else {
				return []reflect.Value{reflect.New(typ.Out(0)).Elem()}
			}
		}
		return nil
	})
}

type deferred struct {
	fn      value
	args    []value
	ssaArgs []ssa.Value
	instr   *ssa.Defer
	tail    *deferred
}

type frame struct {
	i                *Interp
	caller           *frame
	fn               *ssa.Function
	block, prevBlock *ssa.BasicBlock
	env              map[ssa.Value]value // dynamic values of SSA variables
	locals           map[ssa.Value]reflect.Value
	mapUnderscoreKey map[types.Type]bool
	defers           *deferred
	result           value
	panicking        bool
	panic            interface{}
}

func (fr *frame) get(key ssa.Value) value {
	if key == nil {
		return nil
	}
	// if ck, ok := key.(*ssa.Const); ok {
	// 	c := constValue(fr.i, ck)
	// 	if c == nil {
	// 		return c
	// 	}
	// 	typ := fr.i.toType(key.Type())
	// 	if typ.PkgPath() == "" {
	// 		return c
	// 	}
	// 	v := reflect.New(typ).Elem()
	// 	SetValue(v, reflect.ValueOf(c))
	// 	return v.Interface()
	// }
	if key.Parent() == nil {
		path := key.String()
		if paths := strings.Split(path, "."); len(paths) == 2 {
			if pkg, ok := registerPkgs[paths[0]]; ok {
				if ext, ok := pkg.Vars[path]; ok {
					return ext.Interface()
				}
			}
		}
		if ext, ok := externValues[key.String()]; ok {
			if fr.i.mode&EnableTracing != 0 {
				fmt.Fprintln(os.Stderr, "\t(external)")
			}
			return ext.Interface()
		}
	}
	switch key := key.(type) {
	case *ssa.Function:
		return key
	case *ssa.Builtin:
		return key
	case *ssa.Const:
		c := constValue(fr.i, key)
		if c == nil {
			return c
		}
		typ := fr.i.toType(key.Type())
		if typ.PkgPath() == "" {
			return c
		}
		v := reflect.New(typ).Elem()
		SetValue(v, reflect.ValueOf(c))
		return v.Interface()
	case *ssa.Global:
		if r, ok := fr.i.globals[key]; ok {
			return r
		}
	}
	if r, ok := fr.env[key]; ok {
		return r
	}
	panic(fmt.Sprintf("get: no value for %T: %v", key, key.String()))
}

// runDefer runs a deferred call d.
// It always returns normally, but may set or clear fr.panic.
//
func (fr *frame) runDefer(d *deferred) {
	if fr.i.mode&EnableTracing != 0 {
		fmt.Fprintf(os.Stderr, "%s: invoking deferred function call\n",
			fr.i.prog.Fset.Position(d.instr.Pos()))
	}
	var ok bool
	defer func() {
		if !ok {
			// Deferred call created a new state of panic.
			fr.panicking = true
			fr.panic = recover()
		}
	}()
	call(fr.i, fr, d.instr.Pos(), d.fn, d.args, d.ssaArgs)
	ok = true
}

// runDefers executes fr's deferred function calls in LIFO order.
//
// On entry, fr.panicking indicates a state of panic; if
// true, fr.panic contains the panic value.
//
// On completion, if a deferred call started a panic, or if no
// deferred call recovered from a previous state of panic, then
// runDefers itself panics after the last deferred call has run.
//
// If there was no initial state of panic, or it was recovered from,
// runDefers returns normally.
//
func (fr *frame) runDefers() {
	for d := fr.defers; d != nil; d = d.tail {
		fr.runDefer(d)
	}
	fr.defers = nil
	// runtime.Goexit() fr.panic == nil
	if fr.panicking && fr.panic != nil {
		panic(fr.panic) // new panic, or still panicking
	}
}

// lookupMethod returns the method set for type typ, which may be one
// of the interpreter's fake types.
func lookupMethod(i *Interp, typ types.Type, meth *types.Func) *ssa.Function {
	return i.prog.LookupMethod(typ, meth.Pkg(), meth.Name())
}

func SetValue(v reflect.Value, x reflect.Value) {
	switch v.Kind() {
	case reflect.Bool:
		v.SetBool(x.Bool())
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		v.SetInt(x.Int())
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		v.SetUint(x.Uint())
	case reflect.Uintptr:
		v.SetUint(x.Uint())
	case reflect.Float32, reflect.Float64:
		v.SetFloat(x.Float())
	case reflect.Complex64, reflect.Complex128:
		v.SetComplex(x.Complex())
	case reflect.String:
		v.SetString(x.String())
	case reflect.UnsafePointer:
		v.SetPointer(unsafe.Pointer(x.Pointer()))
	default:
		v.Set(x)
	}
}

// 0 = 32bit, 1 = 64bit
const ptrSizeIndex = (^uintptr(0) >> 63)

var (
	maxMem int
)

func init() {
	if ptrSizeIndex == 0 {
		maxMem = 1<<31 - 1
	} else {
		maxMem = 1 << 59
	}
}

func hasUnderscore(st *types.Struct) bool {
	n := st.NumFields()
	for i := 0; i < n; i++ {
		if st.Field(i).Name() == "_" {
			return true
		}
	}
	return false
}

// visitInstr interprets a single ssa.Instruction within the activation
// record frame.  It returns a continuation value indicating where to
// read the next instruction from.
func visitInstr(fr *frame, instr ssa.Instruction) (func(), continuation) {
	if fr.i.mode&EnableDumpInstr != 0 {
		fmt.Fprintf(os.Stderr, "Instr %T %+v\n", instr, instr)
	}
	switch instr := instr.(type) {
	case *ssa.DebugRef:
		// no-op
	case *ssa.UnOp:
		fr.env[instr] = unop(instr, fr.get(instr.X))

	case *ssa.BinOp:
		if instr.Op == token.SHR || instr.Op == token.SHL {
			if c, ok := instr.Y.(*ssa.Convert); ok {
				v := reflect.ValueOf(fr.get(c.X))
				vk := v.Kind()
				if vk >= reflect.Int && vk <= reflect.Int64 {
					if v.Int() < 0 {
						panic(runtimeError("negative shift amount"))
					}
				}
			}
		}
		fr.env[instr] = binop(instr, instr.X.Type(), fr.get(instr.X), fr.get(instr.Y))

	case *ssa.Call:
		return func() {
			fn, args := prepareCall(fr, &instr.Call)
			fr.env[instr] = call(fr.i, fr, instr.Pos(), fn, args, instr.Call.Args)
		}, kNext

	case *ssa.ChangeInterface:
		fr.env[instr] = fr.get(instr.X)

	case *ssa.ChangeType:
		typ := fr.i.toType(instr.Type())
		x := fr.get(instr.X)
		var v reflect.Value
		switch f := x.(type) {
		case *ssa.Function:
			v = fr.toFunc(fr.i.toType(f.Type()), f)
		default:
			v = reflect.ValueOf(x)
		}
		if !v.IsValid() {
			fr.env[instr] = reflect.New(typ).Elem()
		} else {
			fr.env[instr] = v.Convert(typ).Interface()
		}
		//fr.env[instr] = fr.get(instr.X)

	case *ssa.Convert:
		typ := fr.i.toType(instr.Type())
		x := fr.get(instr.X)
		fr.env[instr] = convert(x, typ)
		//fr.env[instr] = conv(fr.i, instr.Type(), instr.X.Type(), fr.get(instr.X))

	case *ssa.MakeInterface:
		typ := fr.i.toType(instr.Type())
		i := reflect.New(typ).Elem()
		xtyp := fr.i.toType(instr.X.Type())
		x := fr.get(instr.X)
		var vx reflect.Value
		switch x.(type) {
		case *ssa.Function:
			vx = fr.toFunc(xtyp, x)
		default:
			vx = reflect.ValueOf(x)
			if xtyp != vx.Type() {
				vx = reflect.ValueOf(convert(x, xtyp))
			}
		}
		SetValue(i, vx)
		fr.env[instr] = i.Interface()
		//fr.env[instr] = iface{t: instr.X.Type(), v: fr.get(instr.X)}

	case *ssa.Extract:
		fr.env[instr] = fr.get(instr.Tuple).(tuple)[instr.Index]

	case *ssa.Slice:
		fr.env[instr] = slice(fr, instr)

	case *ssa.Return:
		switch len(instr.Results) {
		case 0:
		case 1:
			fr.result = fr.get(instr.Results[0])
		default:
			var res []value
			for _, r := range instr.Results {
				res = append(res, fr.get(r))
			}
			fr.result = tuple(res)
		}
		fr.block = nil
		return nil, kReturn

	case *ssa.RunDefers:
		fr.runDefers()

	case *ssa.Panic:
		panic(targetPanic{fr.get(instr.X)})

	case *ssa.Send:
		c := fr.get(instr.Chan)
		x := fr.get(instr.X)
		ch := reflect.ValueOf(c)
		if x == nil {
			ch.Send(reflect.New(ch.Type().Elem()).Elem())
		} else {
			ch.Send(reflect.ValueOf(x))
		}
		//fr.get(instr.Chan).(chan value) <- fr.get(instr.X)

	case *ssa.Store:
		// skip struct field _
		if addr, ok := instr.Addr.(*ssa.FieldAddr); ok {
			if s, ok := addr.X.Type().(*types.Pointer).Elem().(*types.Struct); ok {
				if s.Field(addr.Field).Name() == "_" {
					break
				}
			}
		}
		x := reflect.ValueOf(fr.get(instr.Addr))
		val := fr.get(instr.Val)
		switch fn := val.(type) {
		case *ssa.Function:
			f := fr.toFunc(fr.i.toType(fn.Type()), fn)
			SetValue(x.Elem(), f)
		default:
			v := reflect.ValueOf(val)
			if v.IsValid() {
				SetValue(x.Elem(), v)
			} else {
				SetValue(x.Elem(), reflect.New(x.Elem().Type()).Elem())
			}
		}
		//store(deref(instr.Addr.Type()), fr.get(instr.Addr).(*value), fr.get(instr.Val))

	case *ssa.If:
		succ := 1
		if v := fr.get(instr.Cond); reflect.ValueOf(v).Bool() {
			succ = 0
		}
		fr.prevBlock, fr.block = fr.block, fr.block.Succs[succ]
		return nil, kJump

	case *ssa.Jump:
		fr.prevBlock, fr.block = fr.block, fr.block.Succs[0]
		return nil, kJump

	case *ssa.Defer:
		fn, args := prepareCall(fr, &instr.Call)
		fr.defers = &deferred{
			fn:      fn,
			args:    args,
			ssaArgs: instr.Call.Args,
			instr:   instr,
			tail:    fr.defers,
		}

	case *ssa.Go:
		fn, args := prepareCall(fr, &instr.Call)
		atomic.AddInt32(&fr.i.goroutines, 1)
		go func() {
			call(fr.i, nil, instr.Pos(), fn, args, instr.Call.Args)
			atomic.AddInt32(&fr.i.goroutines, -1)
		}()

	case *ssa.MakeChan:
		typ := fr.i.toType(instr.Type())
		size := fr.get(instr.Size)
		buffer := asInt(size)
		if buffer < 0 {
			panic(runtimeError("makechan: size out of range"))
		}
		fr.env[instr] = reflect.MakeChan(typ, buffer).Interface()
		//fr.env[instr] = make(chan value, asInt(fr.get(instr.Size)))

	case *ssa.Alloc:
		typ := fr.i.toType(instr.Type()).Elem() //deref(instr.Type()))
		//var addr *value
		if instr.Heap {
			// new
			//addr = new(value)
			//fr.env[instr] = addr
			fr.env[instr] = reflect.New(typ).Interface()
		} else {
			//fr.env[instr] = fr.locals[instr]
			// local
			//addr = fr.env[instr].(*value)
			v := reflect.ValueOf(fr.env[instr])
			SetValue(v.Elem(), fr.locals[instr])
			//SetValue(v.Elem(), reflect.Zero(typ))
		}
		//*addr = zero(deref(instr.Type()))

	case *ssa.MakeSlice:
		typ := fr.i.toType(instr.Type())
		Len := asInt(fr.get(instr.Len))
		if Len < 0 || Len >= maxMem {
			panic(runtimeError("makeslice: len out of range"))
		}
		Cap := asInt(fr.get(instr.Cap))
		if Cap < 0 || Cap >= maxMem {
			panic(runtimeError("makeslice: cap out of range"))
		}
		fr.env[instr] = reflect.MakeSlice(typ, Len, Cap).Interface()
		// slice := make([]value, asInt(fr.get(instr.Cap)))
		// tElt := instr.Type().Underlying().(*types.Slice).Elem()
		// for i := range slice {
		// 	slice[i] = zero(tElt)
		// }
		// fr.env[instr] = slice[:asInt(fr.get(instr.Len))]

	case *ssa.MakeMap:
		typ := instr.Type()
		reserve := 0
		if instr.Reserve != nil {
			reserve = asInt(fr.get(instr.Reserve))
		}
		key := typ.Underlying().(*types.Map).Key()
		if st, ok := key.Underlying().(*types.Struct); ok && hasUnderscore(st) {
			fr.mapUnderscoreKey[typ] = true
		}
		fr.env[instr] = reflect.MakeMapWithSize(fr.i.toType(typ), reserve).Interface()
		//fr.env[instr] = makeMap(instr.Type().Underlying().(*types.Map).Key(), reserve)

	case *ssa.Range:
		fr.env[instr] = rangeIter(fr.get(instr.X), instr.X.Type())

	case *ssa.Next:
		fr.env[instr] = fr.get(instr.Iter).(iter).next()

	case *ssa.FieldAddr:
		// v := reflect.ValueOf(fr.get(instr.X)).Elem()
		// fr.env[instr] = reflectx.Field(v, instr.Field).Addr().Interface()
		//fr.env[instr] = &(*fr.get(instr.X).(*value)).(structure)[instr.Field]
		v, err := FieldAddr(fr.get(instr.X), instr.Field)
		if err != nil {
			panic(runtimeError(err.Error()))
		}
		fr.env[instr] = v
	case *ssa.Field:
		// v := reflect.ValueOf(fr.get(instr.X))
		// for v.Kind() == reflect.Ptr {
		// 	v = v.Elem()
		// }
		//fr.env[instr] = reflectx.Field(v, instr.Field).Interface()
		//fr.env[instr] = fr.get(instr.X).(structure)[instr.Field]
		v, err := Field(fr.get(instr.X), instr.Field)
		if err != nil {
			panic(runtimeError(err.Error()))
		}
		fr.env[instr] = v

	case *ssa.IndexAddr:
		x := fr.get(instr.X)
		idx := fr.get(instr.Index)
		v := reflect.ValueOf(x)
		if v.Kind() == reflect.Ptr {
			v = v.Elem()
		}
		switch v.Kind() {
		case reflect.Slice:
		case reflect.Array:
		case reflect.Invalid:
			panic(runtimeError("invalid memory address or nil pointer dereference"))
		default:
			panic(fmt.Sprintf("unexpected x type in IndexAddr: %T", x))
		}
		index := asInt(idx)
		if index < 0 {
			panic(runtimeError(fmt.Sprintf("index out of range [%v]", index)))
		} else if length := v.Len(); index >= length {
			panic(runtimeError(fmt.Sprintf("index out of range [%v] with length %v", index, length)))
		}
		fr.env[instr] = v.Index(index).Addr().Interface()
		// switch x := x.(type) {
		// case []value:
		// 	fr.env[instr] = &x[asInt(idx)]
		// case *value: // *array
		// 	fr.env[instr] = &(*x).(array)[asInt(idx)]
		// default:
		// 	panic(fmt.Sprintf("unexpected x type in IndexAddr: %T", x))
		// }

	case *ssa.Index:
		x := fr.get(instr.X)
		idx := fr.get(instr.Index)
		v := reflect.ValueOf(x)
		if v.Kind() == reflect.Ptr {
			v = v.Elem()
		}
		fr.env[instr] = v.Index(asInt(idx)).Interface()

	case *ssa.Lookup:
		m := fr.get(instr.X)
		idx := fr.get(instr.Index)
		if s, ok := m.(string); ok {
			fr.env[instr] = s[asInt(idx)]
		} else {
			vm := reflect.ValueOf(m)
			v := vm.MapIndex(reflect.ValueOf(idx))
			ok := v.IsValid()
			var rv value
			if ok {
				rv = v.Interface()
			} else {
				rv = reflect.New(vm.Type().Elem()).Elem().Interface()
			}
			if instr.CommaOk {
				fr.env[instr] = tuple{rv, ok}
			} else {
				fr.env[instr] = rv
			}
		}
		//fr.env[instr] = lookup(instr, fr.get(instr.X), fr.get(instr.Index))

	case *ssa.MapUpdate:
		vm := reflect.ValueOf(fr.get(instr.Map))
		vk := reflect.ValueOf(fr.get(instr.Key))
		v := fr.get(instr.Value)
		if fn, ok := v.(*ssa.Function); ok {
			typ := fr.i.toType(fn.Type())
			v = fr.toFunc(typ, fn).Interface()
		}
		if fr.mapUnderscoreKey[instr.Map.Type()] {
			for _, vv := range vm.MapKeys() {
				if equalStruct(vk, vv) {
					vk = vv
					break
				}
			}
		}
		vm.SetMapIndex(vk, reflect.ValueOf(v))

	case *ssa.TypeAssert:
		v := fr.get(instr.X)
		if fn, ok := v.(*ssa.Function); ok {
			typ := fr.i.toType(fn.Type())
			v = fr.toFunc(typ, fn).Interface()
		}
		fr.env[instr] = typeAssert(fr.i, instr, v)

	case *ssa.MakeClosure:
		var bindings []value
		for _, binding := range instr.Bindings {
			bindings = append(bindings, fr.get(binding))
		}
		//fr.env[instr] = &closure{instr.Fn.(*ssa.Function), bindings}
		c := &closure{instr.Fn.(*ssa.Function), bindings}
		typ := fr.i.toType(c.Fn.Type())
		fr.env[instr] = fr.toFunc(typ, c).Interface()

	case *ssa.Phi:
		for i, pred := range instr.Block().Preds {
			if fr.prevBlock == pred {
				fr.env[instr] = fr.get(instr.Edges[i])
				break
			}
		}

	case *ssa.Select:
		var cases []reflect.SelectCase
		if !instr.Blocking {
			cases = append(cases, reflect.SelectCase{
				Dir: reflect.SelectDefault,
			})
		}
		for _, state := range instr.States {
			var dir reflect.SelectDir
			if state.Dir == types.RecvOnly {
				dir = reflect.SelectRecv
			} else {
				dir = reflect.SelectSend
			}
			ch := reflect.ValueOf(fr.get(state.Chan))
			var send reflect.Value
			if state.Send != nil {
				v := fr.get(state.Send)
				if v == nil {
					send = reflect.New(ch.Type().Elem()).Elem()
				} else {
					send = reflect.ValueOf(v)
				}
			}
			cases = append(cases, reflect.SelectCase{
				Dir:  dir,
				Chan: ch,
				Send: send,
			})
		}
		chosen, recv, recvOk := reflect.Select(cases)
		if !instr.Blocking {
			chosen-- // default case should have index -1.
		}
		r := tuple{chosen, recvOk}
		for i, st := range instr.States {
			if st.Dir == types.RecvOnly {
				var v value
				if i == chosen && recvOk {
					// No need to copy since send makes an unaliased copy.
					v = recv.Interface()
				} else {
					typ := fr.i.toType(st.Chan.Type())
					v = reflect.New(typ.Elem()).Elem().Interface()
					//v = zero(st.Chan.Type().Underlying().(*types.Chan).Elem())
				}
				r = append(r, v)
			}
		}
		fr.env[instr] = r
	case *ssa.SliceToArrayPointer:
		typ := fr.i.toType(instr.Type())
		x := fr.get(instr.X)
		v := reflect.ValueOf(x)
		vLen := v.Len()
		tLen := typ.Elem().Len()
		if tLen > vLen {
			panic(runtimeError(fmt.Sprintf("cannot convert slice with length %v to pointer to array with length %v", vLen, tLen)))
		}
		fr.env[instr] = v.Convert(typ).Interface()
	default:
		panic(fmt.Sprintf("unexpected instruction: %T", instr))
	}

	// if val, ok := instr.(ssa.Value); ok {
	// 	fmt.Println(toString(fr.env[val])) // debugging
	// }

	return nil, kNext
}

// prepareCall determines the function value and argument values for a
// function call in a Call, Go or Defer instruction, performing
// interface method lookup if needed.
//
func prepareCall(fr *frame, call *ssa.CallCommon) (fn value, args []value) {
	v := fr.get(call.Value)
	if call.Method == nil {
		// Function call.
		fn = v
	} else {
		// Interface method invocation.
		//vt, ok := call.Value.(*ssa.MakeInterface)
		// recv := v.(iface)
		rv := reflect.ValueOf(v)
		t, ok := fr.i.findType(rv.Type())
		if ok {
			if f := lookupMethod(fr.i, t, call.Method); f == nil {
				// Unreachable in well-typed programs.
				panic(fmt.Sprintf("method set for dynamic type %v does not contain %s", t, call.Method))
			} else {
				fn = f
			}
		} else {
			rtype := rv.Type()
			mname := call.Method.Name()
			if s := rtype.String(); (s == "*reflect.rtype" || s == "reflect.Type") &&
				mname == "Method" || mname == "MethodByName" {
				if mname == "Method" {
					fn = reflectx.MethodByIndex
				} else {
					fn = reflectx.MethodByName
				}
			} else {
				if f, ok := rtype.MethodByName(mname); ok {
					fn = f.Func.Interface()
				} else {
					panic("method invoked on nil interface")
				}
			}
		}
		args = append(args, v)
	}
	for _, arg := range call.Args {
		v := fr.get(arg)
		if fn, ok := v.(*ssa.Function); ok {
			v = fr.toFunc(fr.i.toType(fn.Type()), fn).Interface()
		}
		args = append(args, v)
	}
	return
}

// call interprets a call to a function (function, builtin or closure)
// fn with arguments args, returning its result.
// callpos is the position of the callsite.
//
func call(i *Interp, caller *frame, callpos token.Pos, fn value, args []value, ssaArgs []ssa.Value) value {
	i.callerMutex.Lock()
	i.caller = caller
	i.callerMutex.Unlock()
	switch fn := fn.(type) {
	case *ssa.Function:
		if fn == nil {
			panic("call of nil function") // nil of func type
		}
		return callSSA(i, caller, callpos, fn, args, nil)
	case *closure:
		if fn.Fn == nil {
			panic("call of nil closure function") // nil of func type
		}
		return callSSA(i, caller, callpos, fn.Fn, args, fn.Env)
	case *ssa.Builtin:
		return callBuiltin(i, caller, callpos, fn, args, ssaArgs)
	default:
		if f := reflect.ValueOf(fn); f.Kind() == reflect.Func {
			return callReflect(i, caller, callpos, f, args, nil)
		}
	}
	panic(fmt.Sprintf("cannot call %T %v", fn, reflect.ValueOf(fn).Kind()))
}

func loc(fset *token.FileSet, pos token.Pos) string {
	if pos == token.NoPos {
		return ""
	}
	return " at " + fset.Position(pos).String()
}

// callSSA interprets a call to function fn with arguments args,
// and lexical environment env, returning its result.
// callpos is the position of the callsite.
//
func callSSA(i *Interp, caller *frame, callpos token.Pos, fn *ssa.Function, args []value, env []value) value {
	if i.mode&EnableTracing != 0 {
		fset := fn.Prog.Fset
		// TODO(adonovan): fix: loc() lies for external functions.
		fmt.Fprintf(os.Stderr, "Entering %s%s.\n", fn, loc(fset, fn.Pos()))
		suffix := ""
		if caller != nil {
			suffix = ", resuming " + caller.fn.String() + loc(fset, callpos)
		}
		defer fmt.Fprintf(os.Stderr, "Leaving %s%s.\n", fn, suffix)
	}
	fr := &frame{
		i:      i,
		caller: caller, // for panic/recover
		fn:     fn,
	}
	if fn.Parent() == nil {
		name := fn.String()
		pkgPath := fn.Pkg.Pkg.Path()
		if pkg, ok := registerPkgs[pkgPath]; ok {
			if ext, ok := pkg.Funcs[name]; ok {
				if i.mode&EnableTracing != 0 {
					fmt.Fprintln(os.Stderr, "\t(external func)")
				}
				return callReflect(i, caller, callpos, ext, args, nil)
			}
			if ext, ok := pkg.Methods[name]; ok {
				if i.mode&EnableTracing != 0 {
					fmt.Fprintln(os.Stderr, "\t(external method)")
				}
				return callReflect(i, caller, callpos, ext, args, nil)
			}
		}
		if ext := externValues[name]; ext.Kind() == reflect.Func {
			if i.mode&EnableTracing != 0 {
				fmt.Fprintln(os.Stderr, "\t(external)")
			}
			return callReflect(i, caller, callpos, ext, args, nil)
			//			return ext(fr, args)
		}
		if fn.Blocks == nil {
			// check unexport method
			if fn.Signature.Recv() != nil {
				v := reflect.ValueOf(args[0])
				if f, ok := v.Type().MethodByName(fn.Name()); ok {
					return callReflect(i, caller, callpos, f.Func, args, nil)
				}
			}
			if fn.Name() == "init" && fn.Type().String() == "func()" {
				return true
			}
			panic("no code for function: " + name)
		}
	}
	fr.env = make(map[ssa.Value]value)
	fr.block = fn.Blocks[0]
	fr.locals = make(map[ssa.Value]reflect.Value)
	fr.mapUnderscoreKey = make(map[types.Type]bool)
	for _, l := range fn.Locals {
		typ := fr.i.toType(deref(l.Type()))
		fr.locals[l] = reflect.New(typ).Elem()   //zero(deref(l.Type()))
		fr.env[l] = reflect.New(typ).Interface() //&fr.locals[i]
	}
	for i, p := range fn.Params {
		fr.env[p] = args[i]
	}
	for i, fv := range fn.FreeVars {
		fr.env[fv] = env[i]
	}
	for fr.block != nil {
		runFrame(fr)
	}
	// Destroy the locals to avoid accidental use after return.
	fr.env = nil
	fr.block = nil
	fr.locals = nil
	return fr.result
}

func callReflect(i *Interp, caller *frame, callpos token.Pos, fn reflect.Value, args []value, env []value) value {
	var ins []reflect.Value
	typ := fn.Type()
	isVariadic := fn.Type().IsVariadic()
	if isVariadic {
		for i := 0; i < len(args)-1; i++ {
			if args[i] == nil {
				ins = append(ins, reflect.New(typ.In(i)).Elem())
			} else {
				ins = append(ins, reflect.ValueOf(args[i]))
			}
		}
		ins = append(ins, reflect.ValueOf(args[len(args)-1]))
	} else {
		ins = make([]reflect.Value, len(args), len(args))
		for i := 0; i < len(args); i++ {
			if args[i] == nil {
				ins[i] = reflect.New(typ.In(i)).Elem()
			} else {
				ins[i] = reflect.ValueOf(args[i])
			}
		}
	}
	var results []reflect.Value
	if isVariadic {
		results = fn.CallSlice(ins)
	} else {
		results = fn.Call(ins)
	}
	switch len(results) {
	case 0:
		return nil
	case 1:
		return results[0].Interface()
	default:
		var res []value
		for _, r := range results {
			res = append(res, r.Interface())
		}
		return tuple(res)
	}
}

// runFrame executes SSA instructions starting at fr.block and
// continuing until a return, a panic, or a recovered panic.
//
// After a panic, runFrame panics.
//
// After a normal return, fr.result contains the result of the call
// and fr.block is nil.
//
// A recovered panic in a function without named return parameters
// (NRPs) becomes a normal return of the zero value of the function's
// result type.
//
// After a recovered panic in a function with NRPs, fr.result is
// undefined and fr.block contains the block at which to resume
// control.
//
func runFrame(fr *frame) {
	defer func() {
		if fr.block == nil {
			return // normal return
		}
		if fr.i.mode&DisableRecover != 0 {
			return // let interpreter crash
		}
		fr.panicking = true
		fr.panic = recover()
		if fr.i.mode&EnableTracing != 0 {
			fmt.Fprintf(os.Stderr, "Panicking: %T %v.\n", fr.panic, fr.panic)
		}
		fr.runDefers()
		fr.block = fr.fn.Recover
	}()

	for {
		if fr.i.mode&EnableTracing != 0 {
			fmt.Fprintf(os.Stderr, ".%s:\n", fr.block)
		}
	block:
		for _, instr := range fr.block.Instrs {
			if fr.i.mode&EnableTracing != 0 {
				if v, ok := instr.(ssa.Value); ok {
					fmt.Fprintln(os.Stderr, "\t", v.Name(), "=", instr)
				} else {
					fmt.Fprintln(os.Stderr, "\t", instr)
				}
			}
			fn, cond := visitInstr(fr, instr)
			if fn != nil {
				fn()
			}
			switch cond {
			case kReturn:
				return
			case kNext:
				// no-op
			case kJump:
				break block
			}
		}
	}
}

// doRecover implements the recover() built-in.
func doRecover(caller *frame) value {
	// recover() must be exactly one level beneath the deferred
	// function (two levels beneath the panicking function) to
	// have any effect.  Thus we ignore both "defer recover()" and
	// "defer f() -> g() -> recover()".
	if caller.i.mode&DisableRecover == 0 &&
		caller != nil && !caller.panicking &&
		caller.caller != nil && caller.caller.panicking {
		caller.caller.panicking = false
		p := caller.caller.panic
		caller.caller.panic = nil
		// TODO(adonovan): support runtime.Goexit.
		switch p := p.(type) {
		case targetPanic:
			// The target program explicitly called panic().
			return p.v
		case runtime.Error:
			// The interpreter encountered a runtime error.
			return p
			//return iface{caller.i.runtimeErrorString, p.Error()}
		case string:
			return p
		case plainError:
			return p
		case runtimeError:
			return p
		case *reflect.ValueError:
			return p
		default:
			panic(fmt.Sprintf("unexpected panic type %T in target call to recover()", p))
		}
	}
	return nil //iface{}
}

// setGlobal sets the value of a system-initialized global variable.
func setGlobal(i *Interp, pkg *ssa.Package, name string, v value) {
	// if g, ok := i.globals[pkg.Var(name)]; ok {
	// 	*g = v
	// 	return
	// }
	panic("no global variable: " + pkg.Pkg.Path() + "." + name)
}

// Interpret interprets the Go program whose main package is mainpkg.
// mode specifies various interpreter options.  filename and args are
// the initial values of os.Args for the target program.  sizes is the
// effective type-sizing function for this program.
//
// Interpret returns the exit code of the program: 2 for panic (like
// gc does), or the argument to os.Exit for normal termination.
//
// The SSA program must include the "runtime" package.
//

func NewInterp(inst Loader, mainpkg *ssa.Package, mode Mode) *Interp {
	i := &Interp{
		prog:       mainpkg.Prog,
		mainpkg:    mainpkg,
		globals:    make(map[ssa.Value]value),
		mode:       mode,
		goroutines: 1,
		types:      make(map[types.Type]reflect.Type),
	}
	i.inst = inst
	i.record = NewTypesRecord(i.inst, i)
	i.record.Load(mainpkg)
	for _, pkg := range i.prog.AllPackages() {
		if _, ok := externPackages[pkg.Pkg.Path()]; ok {
			if i.mode&EnableDumpPackage != 0 {
				fmt.Fprintln(os.Stderr, "initialize", pkg, "(extern)")
			}
			continue
		}
		if i.mode&EnableDumpPackage != 0 {
			fmt.Fprintln(os.Stderr, "initialize", pkg)
		}
		// Initialize global storage.
		for _, m := range pkg.Members {
			switch v := m.(type) {
			case *ssa.Global:
				typ := i.toType(deref(v.Type()))
				i.globals[v] = reflect.New(typ).Interface()
			}
		}
	}
	return i
}

func (i *Interp) Run(entry string) (r value, exitCode int) {
	// Top-level error handler.
	exitCode = 2
	defer func() {
		if exitCode != 2 || i.mode&DisableRecover != 0 {
			return
		}
		switch p := recover().(type) {
		case exitPanic:
			exitCode = int(p)
			return
		case targetPanic:
			fmt.Fprintln(os.Stderr, "panic:", toString(p.v))
		case runtime.Error:
			fmt.Fprintln(os.Stderr, "panic:", p.Error())
		case string:
			fmt.Fprintln(os.Stderr, "panic:", p)
		case plainError:
			fmt.Fprintln(os.Stderr, "panic:", p.Error())
		default:
			fmt.Fprintf(os.Stderr, "panic: unexpected type: %T: %v\n", p, p)
		}
	}()
	if mainFn := i.mainpkg.Func(entry); mainFn != nil {
		r = call(i, nil, token.NoPos, mainFn, nil, nil)
		exitCode = 0
	} else {
		fmt.Fprintln(os.Stderr, "No function.", entry)
		exitCode = 1
	}
	return
}

// remove
func _Interpret(mainpkg *ssa.Package, mode Mode, entry string) (exitCode int) {
	return
	// i := &Interp{
	// 	prog:       mainpkg.Prog,
	// 	globals:    make(map[ssa.Value]value),
	// 	mode:       mode,
	// 	goroutines: 1,
	// 	types:      make(map[types.Type]reflect.Type),
	// }
	// i.ctx = xtypes.NewContext(func(mtyp reflect.Type, fn *types.Func) func([]reflect.Value) []reflect.Value {
	// 	typ := fn.Type().(*types.Signature).Recv().Type()
	// 	if f := i.prog.LookupMethod(typ, fn.Pkg(), fn.Name()); f != nil {
	// 		return func(args []reflect.Value) []reflect.Value {
	// 			iargs := make([]value, len(args))
	// 			for i := 0; i < len(args); i++ {
	// 				iargs[i] = args[i].Interface()
	// 			}
	// 			i.callerMutex.RLock()
	// 			caller := i.caller
	// 			i.callerMutex.RUnlock()
	// 			r := call(i, caller, token.NoPos, f, iargs, nil)
	// 			switch mtyp.NumOut() {
	// 			case 0:
	// 				return nil
	// 			case 1:
	// 				if r == nil {
	// 					return []reflect.Value{reflect.New(mtyp.Out(0)).Elem()}
	// 				} else {
	// 					return []reflect.Value{reflect.ValueOf(r)}
	// 				}
	// 			default:
	// 				v, ok := r.(tuple)
	// 				if !ok {
	// 					panic(fmt.Errorf("result type must tuple: %T", v))
	// 				}
	// 				res := make([]reflect.Value, len(v))
	// 				for j := 0; j < len(v); j++ {
	// 					if v[j] == nil {
	// 						res[j] = reflect.New(mtyp.Out(j)).Elem()
	// 					} else {
	// 						res[j] = reflect.ValueOf(v[j])
	// 					}
	// 				}
	// 				return res
	// 			}
	// 		}
	// 	}
	// 	if v, ok := externValues[fn.FullName()]; ok && v.Kind() == reflect.Func {
	// 		return func(args []reflect.Value) []reflect.Value {
	// 			return v.Call(args)
	// 		}
	// 	}
	// 	panic(fmt.Sprintf("Not found func %v", fn))
	// 	return nil
	// }, func(name *types.TypeName) (reflect.Type, bool) {
	// 	if t := i.inst.tcache.At(name.Type()); t != nil {
	// 		return t.(reflect.Type), true
	// 	}
	// 	if typ, ok := externTypes[name.Type().String()]; ok {
	// 		return typ, true
	// 	}
	// 	return nil, false
	// }, func(typ types.Type) (reflect.Type, bool) {
	// 	if t := i.inst.tcache.At(typ); t != nil {
	// 		return t.(reflect.Type), true
	// 	}
	// 	if t, ok := i.types[typ]; ok {
	// 		return t, true
	// 	}
	// 	for k, v := range i.types {
	// 		if types.Identical(k, typ) {
	// 			return v, true
	// 		}
	// 	}
	// 	return nil, false
	// })

	// for _, pkg := range i.prog.AllPackages() {
	// 	if _, ok := externPackages[pkg.Pkg.Path()]; ok {
	// 		if i.mode&EnableDumpPackage != 0 {
	// 			fmt.Fprintln(os.Stderr, "initialize", pkg, "(extern)")
	// 		}
	// 		continue
	// 	}
	// 	if i.mode&EnableDumpPackage != 0 {
	// 		fmt.Fprintln(os.Stderr, "initialize", pkg)
	// 	}
	// 	// Initialize global storage.
	// 	for _, m := range pkg.Members {
	// 		switch v := m.(type) {
	// 		case *ssa.Global:
	// 			typ := i.toType(deref(v.Type()))
	// 			i.globals[v] = reflect.New(typ).Interface()
	// 		}
	// 	}
	// }

	// // Top-level error handler.
	// exitCode = 2
	// defer func() {
	// 	if exitCode != 2 || i.mode&DisableRecover != 0 {
	// 		return
	// 	}
	// 	switch p := recover().(type) {
	// 	case exitPanic:
	// 		exitCode = int(p)
	// 		return
	// 	case targetPanic:
	// 		fmt.Fprintln(os.Stderr, "panic:", toString(p.v))
	// 	case runtime.Error:
	// 		fmt.Fprintln(os.Stderr, "panic:", p.Error())
	// 	case string:
	// 		fmt.Fprintln(os.Stderr, "panic:", p)
	// 	case plainError:
	// 		fmt.Fprintln(os.Stderr, "panic:", p.Error())
	// 	default:
	// 		fmt.Fprintf(os.Stderr, "panic: unexpected type: %T: %v\n", p, p)
	// 	}

	// 	// TODO(adonovan): dump panicking interpreter goroutine?
	// 	// buf := make([]byte, 0x10000)
	// 	// runtime.Stack(buf, false)
	// 	// fmt.Fprintln(os.Stderr, string(buf))
	// 	// (Or dump panicking target goroutine?)
	// }()

	// // Run!
	// call(i, nil, token.NoPos, mainpkg.Func("init"), nil, nil)
	// if mainFn := mainpkg.Func(entry); mainFn != nil {
	// 	call(i, nil, token.NoPos, mainFn, nil, nil)
	// 	exitCode = 0
	// } else {
	// 	fmt.Fprintln(os.Stderr, "No main function.")
	// 	exitCode = 1
	// }
	// return
}

// deref returns a pointer's element type; otherwise it returns typ.
// TODO(adonovan): Import from ssa?
func deref(typ types.Type) types.Type {
	if p, ok := typ.Underlying().(*types.Pointer); ok {
		return p.Elem()
	}
	return typ
}
