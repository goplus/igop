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
package interp

import (
	"fmt"
	"go/token"
	"go/types"
	"os"
	"reflect"
	"runtime"
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

// Mode is a bitmask of options affecting the interpreter.
type Mode uint

const (
	DisableRecover  Mode = 1 << iota // Disable recover() in target programs; show interpreter crash instead.
	EnableTracing                    // Print a trace of all instructions as they are interpreted.
	EnableDumpInstr                  //Print instr type & value
)

type methodSet map[string]*ssa.Function

// State shared between all interpreted goroutines.
type interpreter struct {
	osArgs             []value             // the value of os.Args
	prog               *ssa.Program        // the SSA program
	globals            map[ssa.Value]value // addresses of global variables (immutable)
	mode               Mode                // interpreter options
	reflectPackage     *ssa.Package        // the fake reflect package
	errorMethods       methodSet           // the method set of reflect.error, which implements the error interface.
	rtypeMethods       methodSet           // the method set of rtype, which implements the reflect.Type interface.
	runtimeErrorString types.Type          // the runtime.errorString type
	sizes              types.Sizes         // the effective type-sizing function
	goroutines         int32               // atomically updated
	ctx                xtypes.Context
	types              map[types.Type]reflect.Type
}

func (i *interpreter) findType(t reflect.Type) (types.Type, bool) {
	for k, v := range i.types {
		if v == t {
			return k, true
		}
	}
	if t.Kind() == reflect.Ptr {
		if typ, ok := i.findType(t.Elem()); ok {
			return types.NewPointer(typ), true
		}
	}
	return nil, false
}

func (i *interpreter) toType(typ types.Type) reflect.Type {
	if t, ok := i.types[typ]; ok {
		return t
	}
	t, err := xtypes.ToType(typ, i.ctx)
	if err != nil {
		panic(fmt.Sprintf("toType %v error: %v", typ, err))
	}
	i.types[typ] = t
	return t
}

func (fr *frame) toFunc(typ reflect.Type, fn value) reflect.Value {
	return reflect.MakeFunc(typ, func(args []reflect.Value) []reflect.Value {
		iargs := make([]value, len(args))
		for i := 0; i < len(args); i++ {
			iargs[i] = args[i].Interface()
		}
		r := call(fr.i, fr, token.NoPos, fn, iargs)
		if v, ok := r.(tuple); ok {
			res := make([]reflect.Value, len(v))
			for i := 0; i < len(v); i++ {
				res[i] = reflect.ValueOf(v[i])
			}
			return res
		} else if r != nil {
			return []reflect.Value{reflect.ValueOf(r)}
		}
		return nil
	})
}

type deferred struct {
	fn    value
	args  []value
	instr *ssa.Defer
	tail  *deferred
}

type frame struct {
	i                *interpreter
	caller           *frame
	fn               *ssa.Function
	block, prevBlock *ssa.BasicBlock
	env              map[ssa.Value]value // dynamic values of SSA variables
	locals           []value
	defers           *deferred
	result           value
	panicking        bool
	panic            interface{}
}

func (fr *frame) get(key ssa.Value) value {
	switch key := key.(type) {
	case nil:
		// Hack; simplifies handling of optional attributes
		// such as ssa.Slice.{Low,High}.
		return nil
	case *ssa.Function, *ssa.Builtin:
		return key
	case *ssa.Const:
		c := constValue(fr.i, key)
		typ := fr.i.toType(key.Type())
		if typ.PkgPath() == "" {
			return c
		}
		if c == nil {
			return reflect.Zero(typ).Interface()
		}
		v := reflect.New(typ).Elem()
		reflectx.SetValue(v, reflect.ValueOf(c))
		return v.Interface()
	case *ssa.Global:
		if r, ok := fr.i.globals[key]; ok {
			return r
		}
	}
	if r, ok := fr.env[key]; ok {
		return r
	}
	panic(fmt.Sprintf("get: no value for %T: %v", key, key.Name()))
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
	call(fr.i, fr, d.instr.Pos(), d.fn, d.args)
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
	if fr.panicking {
		panic(fr.panic) // new panic, or still panicking
	}
}

// lookupMethod returns the method set for type typ, which may be one
// of the interpreter's fake types.
func lookupMethod(i *interpreter, typ types.Type, meth *types.Func) *ssa.Function {
	switch typ {
	case rtypeType:
		return i.rtypeMethods[meth.Id()]
	case errorType:
		return i.errorMethods[meth.Id()]
	}
	return i.prog.LookupMethod(typ, meth.Pkg(), meth.Name())
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
		fr.env[instr] = binop(instr.Op, instr.X.Type(), fr.get(instr.X), fr.get(instr.Y))

	case *ssa.Call:
		return func() {
			fn, args := prepareCall(fr, &instr.Call)
			fr.env[instr] = call(fr.i, fr, instr.Pos(), fn, args)
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
		if v.Kind() == reflect.Ptr {
			fr.env[instr] = reflect.NewAt(typ, unsafe.Pointer(v.Pointer())).Interface()
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
		reflectx.SetValue(i, reflect.ValueOf(fr.get(instr.X)))
		fr.env[instr] = i.Interface()
		//fr.env[instr] = iface{t: instr.X.Type(), v: fr.get(instr.X)}

	case *ssa.Extract:
		fr.env[instr] = fr.get(instr.Tuple).(tuple)[instr.Index]

	case *ssa.Slice:
		fr.env[instr] = slice(fr.get(instr.X), fr.get(instr.Low), fr.get(instr.High), fr.get(instr.Max))

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
			ch.Send(reflect.Zero(ch.Type().Elem()))
		} else {
			ch.Send(reflect.ValueOf(x))
		}
		//fr.get(instr.Chan).(chan value) <- fr.get(instr.X)

	case *ssa.Store:
		x := reflect.ValueOf(fr.get(instr.Addr))
		val := fr.get(instr.Val)
		switch fn := val.(type) {
		case *ssa.Function:
			f := fr.toFunc(fr.i.toType(fn.Type()), fn)
			reflectx.SetValue(x.Elem(), f)
		default:
			v := reflect.ValueOf(val)
			reflectx.SetValue(x.Elem(), v)
		}
		//store(deref(instr.Addr.Type()), fr.get(instr.Addr).(*value), fr.get(instr.Val))

	case *ssa.If:
		succ := 1
		if fr.get(instr.Cond).(bool) {
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
			fn:    fn,
			args:  args,
			instr: instr,
			tail:  fr.defers,
		}

	case *ssa.Go:
		fn, args := prepareCall(fr, &instr.Call)
		atomic.AddInt32(&fr.i.goroutines, 1)
		go func() {
			call(fr.i, nil, instr.Pos(), fn, args)
			atomic.AddInt32(&fr.i.goroutines, -1)
		}()

	case *ssa.MakeChan:
		typ := fr.i.toType(instr.Type())
		size := fr.get(instr.Size)
		fr.env[instr] = reflect.MakeChan(typ, asInt(size)).Interface()
		//fr.env[instr] = make(chan value, asInt(fr.get(instr.Size)))

	case *ssa.Alloc:
		typ := fr.i.toType(deref(instr.Type()))
		//var addr *value
		if instr.Heap {
			// new
			//addr = new(value)
			//fr.env[instr] = addr
			fr.env[instr] = reflect.New(typ).Interface()
		} else {
			// local
			//addr = fr.env[instr].(*value)
			v := reflect.ValueOf(fr.env[instr])
			reflectx.SetValue(v.Elem(), reflect.Zero(typ))
		}
		//*addr = zero(deref(instr.Type()))

	case *ssa.MakeSlice:
		typ := fr.i.toType(instr.Type())
		Len := asInt(fr.get(instr.Len))
		Cap := asInt(fr.get(instr.Cap))
		fr.env[instr] = reflect.MakeSlice(typ, Len, Cap).Interface()
		// slice := make([]value, asInt(fr.get(instr.Cap)))
		// tElt := instr.Type().Underlying().(*types.Slice).Elem()
		// for i := range slice {
		// 	slice[i] = zero(tElt)
		// }
		// fr.env[instr] = slice[:asInt(fr.get(instr.Len))]

	case *ssa.MakeMap:
		typ := fr.i.toType(instr.Type())
		reserve := 0
		if instr.Reserve != nil {
			reserve = asInt(fr.get(instr.Reserve))
		}
		fr.env[instr] = reflect.MakeMapWithSize(typ, reserve).Interface()
		//fr.env[instr] = makeMap(instr.Type().Underlying().(*types.Map).Key(), reserve)

	case *ssa.Range:
		fr.env[instr] = rangeIter(fr.get(instr.X), instr.X.Type())

	case *ssa.Next:
		fr.env[instr] = fr.get(instr.Iter).(iter).next()

	case *ssa.FieldAddr:
		v := reflect.ValueOf(fr.get(instr.X)).Elem()
		fr.env[instr] = reflectx.Field(v, instr.Field).Addr().Interface()
		//fr.env[instr] = &(*fr.get(instr.X).(*value)).(structure)[instr.Field]

	case *ssa.Field:
		v := reflect.ValueOf(fr.get(instr.X))
		for v.Kind() == reflect.Ptr {
			v = v.Elem()
		}
		fr.env[instr] = reflectx.Field(v, instr.Field).Interface()
		//fr.env[instr] = fr.get(instr.X).(structure)[instr.Field]

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
		default:
			panic(fmt.Sprintf("unexpected x type in IndexAddr: %T", x))
		}
		fr.env[instr] = v.Index(asInt(idx)).Addr().Interface()
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
				rv = reflect.Zero(vm.Type().Elem()).Interface()
			}
			if instr.CommaOk {
				fr.env[instr] = tuple{rv, ok}
			} else {
				fr.env[instr] = rv
			}
		}
		//fr.env[instr] = lookup(instr, fr.get(instr.X), fr.get(instr.Index))

	case *ssa.MapUpdate:
		m := fr.get(instr.Map)
		key := fr.get(instr.Key)
		v := fr.get(instr.Value)
		reflect.ValueOf(m).SetMapIndex(reflect.ValueOf(key), reflect.ValueOf(v))
		// switch m := m.(type) {
		// case map[value]value:
		// 	m[key] = v
		// case *hashmap:
		// 	m.insert(key.(hashable), v)
		// default:
		// 	panic(fmt.Sprintf("illegal map type: %T", m))
		// }

	case *ssa.TypeAssert:
		fr.env[instr] = typeAssert(fr.i, instr, fr.get(instr.X))

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
					send = reflect.Zero(ch.Type().Elem())
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
					v = recv.Interface().(value)
				} else {
					typ := fr.i.toType(st.Chan.Type()).Elem()
					v = reflect.Zero(typ).Interface()
					//v = zero(st.Chan.Type().Underlying().(*types.Chan).Elem())
				}
				r = append(r, v)
			}
		}
		fr.env[instr] = r

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
		t, ok := fr.i.findType(reflect.TypeOf(v))
		if !ok {
			panic("method invoked on nil interface")
		}
		if f := lookupMethod(fr.i, t, call.Method); f == nil {
			// Unreachable in well-typed programs.
			panic(fmt.Sprintf("method set for dynamic type %v does not contain %s", t, call.Method))
		} else {
			fn = f
		}
		args = append(args, v)
	}
	for _, arg := range call.Args {
		args = append(args, fr.get(arg))
	}
	return
}

// call interprets a call to a function (function, builtin or closure)
// fn with arguments args, returning its result.
// callpos is the position of the callsite.
//
func call(i *interpreter, caller *frame, callpos token.Pos, fn value, args []value) value {
	switch fn := fn.(type) {
	case *ssa.Function:
		if fn == nil {
			panic("call of nil function") // nil of func type
		}
		return callSSA(i, caller, callpos, fn, args, nil)
	case *closure:
		return callSSA(i, caller, callpos, fn.Fn, args, fn.Env)
	case *ssa.Builtin:
		return callBuiltin(caller, callpos, fn, args)
	default:
		if f := reflect.ValueOf(fn); f.Kind() == reflect.Func {
			vargs := make([]reflect.Value, len(args))
			for i := 0; i < len(args); i++ {
				vargs[i] = reflect.ValueOf(args[i])
			}
			var results []reflect.Value
			if f.Type().IsVariadic() {
				results = f.CallSlice(vargs)
			} else {
				results = f.Call(vargs)
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
func callSSA(i *interpreter, caller *frame, callpos token.Pos, fn *ssa.Function, args []value, env []value) value {
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
		if ext := externals[name]; ext != nil {
			if i.mode&EnableTracing != 0 {
				fmt.Fprintln(os.Stderr, "\t(external)")
			}
			return ext(fr, args)
		}
		if fn.Blocks == nil {
			panic("no code for function: " + name)
		}
	}
	fr.env = make(map[ssa.Value]value)
	fr.block = fn.Blocks[0]
	fr.locals = make([]value, len(fn.Locals))
	for i, l := range fn.Locals {
		typ := fr.i.toType(deref(l.Type()))
		v := reflect.New(typ).Interface()
		fr.locals[i] = v //zero(deref(l.Type()))
		fr.env[l] = v    //&fr.locals[i]
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
	for i := range fn.Locals {
		fr.locals[i] = bad{}
	}
	return fr.result
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

type errorString string

func (e errorString) RuntimeError() {}

func (e errorString) Error() string {
	return "runtime error: " + string(e)
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
			// The interpreter explicitly called panic().
			return p
			//return iface{caller.i.runtimeErrorString, p}
		case *reflect.ValueError:
			return p
		default:
			panic(fmt.Sprintf("unexpected panic type %T in target call to recover()", p))
		}
	}
	return nil //iface{}
}

// setGlobal sets the value of a system-initialized global variable.
func setGlobal(i *interpreter, pkg *ssa.Package, name string, v value) {
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
func Interpret(mainpkg *ssa.Package, mode Mode, sizes types.Sizes, filename string, args []string) (exitCode int) {
	i := &interpreter{
		prog:       mainpkg.Prog,
		globals:    make(map[ssa.Value]value),
		mode:       mode,
		sizes:      sizes,
		goroutines: 1,
		types:      make(map[types.Type]reflect.Type),
	}
	i.ctx = xtypes.NewContext(func(fn *types.Func) func([]reflect.Value) []reflect.Value {
		typ := fn.Type().(*types.Signature).Recv().Type()
		if f := i.prog.LookupMethod(typ, fn.Pkg(), fn.Name()); f != nil {
			return func(args []reflect.Value) []reflect.Value {
				iargs := make([]value, len(args))
				for i := 0; i < len(args); i++ {
					iargs[i] = args[i].Interface()
				}
				r := call(i, nil, token.NoPos, f, iargs)
				if v, ok := r.(tuple); ok {
					res := make([]reflect.Value, len(v))
					for i := 0; i < len(v); i++ {
						res[i] = reflect.ValueOf(v[i])
					}
					return res
				} else if r != nil {
					return []reflect.Value{reflect.ValueOf(r)}
				}
				return nil
			}
		}
		return nil
	})
	runtimePkg := i.prog.ImportedPackage("runtime")
	if runtimePkg == nil {
		panic("ssa.Program doesn't include runtime package")
	}
	i.runtimeErrorString = runtimePkg.Type("errorString").Object().Type()

	initReflect(i)

	i.osArgs = append(i.osArgs, filename)
	for _, arg := range args {
		i.osArgs = append(i.osArgs, arg)
	}

	for _, pkg := range i.prog.AllPackages() {
		// Initialize global storage.
		for _, m := range pkg.Members {
			switch v := m.(type) {
			case *ssa.Global:
				// cell := zero(deref(v.Type()))
				// i.globals[v] = &cell
				pkgPath := v.Pkg.Pkg.Path()
				if pkgPath == "main" || pkgPath == "unsafe" {
					typ := i.toType(deref(v.Type()))
					i.globals[v] = reflect.New(typ).Interface()
				}
			}
		}
	}

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
		default:
			fmt.Fprintf(os.Stderr, "panic: unexpected type: %T: %v\n", p, p)
		}

		// TODO(adonovan): dump panicking interpreter goroutine?
		// buf := make([]byte, 0x10000)
		// runtime.Stack(buf, false)
		// fmt.Fprintln(os.Stderr, string(buf))
		// (Or dump panicking target goroutine?)
	}()

	// Run!
	call(i, nil, token.NoPos, mainpkg.Func("init"), nil)
	if mainFn := mainpkg.Func("main"); mainFn != nil {
		call(i, nil, token.NoPos, mainFn, nil)
		exitCode = 0
	} else {
		fmt.Fprintln(os.Stderr, "No main function.")
		exitCode = 1
	}
	return
}

// deref returns a pointer's element type; otherwise it returns typ.
// TODO(adonovan): Import from ssa?
func deref(typ types.Type) types.Type {
	if p, ok := typ.Underlying().(*types.Pointer); ok {
		return p.Elem()
	}
	return typ
}
