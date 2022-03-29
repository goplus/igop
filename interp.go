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
	"go/constant"
	"go/token"
	"go/types"
	"reflect"
	"runtime"
	"sync"
	"sync/atomic"
	"unsafe"

	"github.com/petermattis/goid"
	"golang.org/x/tools/go/ssa"
)

var (
	maxMemLen int
)

const intSize = 32 << (^uint(0) >> 63)

func init() {
	if intSize == 32 {
		maxMemLen = 1<<31 - 1
	} else {
		v := int64(1) << 59
		maxMemLen = int(v)
	}
}

type continuation = int

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
	fset         *token.FileSet
	prog         *ssa.Program        // the SSA program
	mainpkg      *ssa.Package        // the SSA main package
	globals      map[ssa.Value]value // addresses of global variables (immutable)
	mode         Mode                // interpreter options
	goroutines   int32               // atomically updated
	deferCount   int32
	exited       bool
	preloadTypes map[types.Type]reflect.Type
	deferMap     sync.Map
	loader       Loader
	record       *TypesRecord
	typesMutex   sync.RWMutex
	fnDebug      func(*DebugInfo)
	funcs        map[*ssa.Function]*Function
	msets        map[reflect.Type](map[string]*ssa.Function) // user defined type method sets
	//sizes        types.Sizes         // the effective type-sizing function
}

func (i *Interp) setDebug(fn func(*DebugInfo)) {
	i.fnDebug = fn
}

func (i *Interp) installed(path string) (pkg *Package, ok bool) {
	pkg, ok = i.loader.Installed(path)
	return
}

func (i *Interp) findType(rt reflect.Type, local bool) (types.Type, bool) {
	i.typesMutex.Lock()
	defer i.typesMutex.Unlock()
	if local {
		return i.record.LookupLocalTypes(rt)
	} else {
		return i.record.LookupTypes(rt)
	}
}

func (i *Interp) tryDeferFrame() *frame {
	if atomic.LoadInt32(&i.deferCount) != 0 {
		if f, ok := i.deferMap.Load(goid.Get()); ok {
			return f.(*frame)
		}
	}
	return nil
}

func (i *Interp) FindMethod(mtyp reflect.Type, fn *types.Func) func([]reflect.Value) []reflect.Value {
	typ := fn.Type().(*types.Signature).Recv().Type()
	if f := i.prog.LookupMethod(typ, fn.Pkg(), fn.Name()); f != nil {
		return func(args []reflect.Value) []reflect.Value {
			iargs := make([]value, len(args))
			for i := 0; i < len(args); i++ {
				iargs[i] = args[i].Interface()
			}
			r := i.callFunction(i.tryDeferFrame(), f, iargs, nil)
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
	name := fn.FullName()
	//	pkgPath := fn.Pkg().Path()
	if v, ok := externValues[name]; ok && v.Kind() == reflect.Func {
		return func(args []reflect.Value) []reflect.Value {
			return v.Call(args)
		}
	}
	// if pkg, ok := i.installed(pkgPath); ok {
	// 	if ext, ok := pkg.Methods[name]; ok {
	// 		return func(args []reflect.Value) []reflect.Value {
	// 			return ext.Call(args)
	// 		}
	// 	}
	// }
	panic(fmt.Sprintf("Not found func %v", fn))
	return nil
}

func (i *Interp) makeFunc(typ reflect.Type, fn *ssa.Function, env []value) reflect.Value {
	return reflect.MakeFunc(typ, func(args []reflect.Value) []reflect.Value {
		iargs := make([]value, len(args))
		for i := 0; i < len(args); i++ {
			iargs[i] = args[i].Interface()
		}
		r := i.callFunction(i.tryDeferFrame(), fn, iargs, env)
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
	interp    *Interp
	caller    *frame
	pfn       *Function
	defers    *deferred
	panicking *panicking
	block     *ssa.BasicBlock
	pc        int
	pred      int
	deferid   int64
	stack     []value
	results   []int
}

func (fr *frame) setReg(index int, v value) {
	fr.stack[index] = v
}

func (fr *frame) reg(index int) value {
	return fr.stack[index]
}

type panicking struct {
	value interface{}
}

// runDefer runs a deferred call d.
// It always returns normally, but may set or clear fr.panic.
//
func (fr *frame) runDefer(d *deferred) {
	var ok bool
	defer func() {
		if !ok {
			// Deferred call created a new state of panic.
			fr.panicking = &panicking{recover()}
		}
	}()
	fr.interp.call(fr, d.fn, d.args, d.ssaArgs)
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
	atomic.AddInt32(&fr.interp.deferCount, 1)
	fr.deferid = goid.Get()
	fr.interp.deferMap.Store(fr.deferid, fr)
	for d := fr.defers; d != nil; d = d.tail {
		fr.runDefer(d)
	}
	fr.interp.deferMap.Delete(fr.deferid)
	atomic.AddInt32(&fr.interp.deferCount, -1)
	fr.deferid = 0
	// runtime.Goexit() fr.panic == nil
	if fr.panicking != nil {
		panic(fr.panicking.value) // new panic, or still panicking
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

func hasUnderscore(st *types.Struct) bool {
	n := st.NumFields()
	for i := 0; i < n; i++ {
		if st.Field(i).Name() == "_" {
			return true
		}
	}
	return false
}

type DebugInfo struct {
	*ssa.DebugRef
	fset    *token.FileSet
	toValue func() (*types.Var, interface{}, bool) // var object value
}

func (i *DebugInfo) Position() token.Position {
	return i.fset.Position(i.Pos())
}

func (i *DebugInfo) AsVar() (*types.Var, interface{}, bool) {
	return i.toValue()
}

func (i *DebugInfo) AsFunc() (*types.Func, bool) {
	v, ok := i.Object().(*types.Func)
	return v, ok
}

// prepareCall determines the function value and argument values for a
// function call in a Call, Go or Defer instruction, performing
// interface method lookup if needed.
//
func (i *Interp) prepareCall(fr *frame, call *ssa.CallCommon, iv int, ia []int, ib []int) (fv value, args []value) {
	if call.Method == nil {
		switch f := call.Value.(type) {
		case *ssa.Builtin:
			fv = f
		case *ssa.Function:
			if f.Blocks == nil {
				ext, ok := findExternFunc(i, f)
				if !ok {
					// skip pkg.init
					if f.Pkg != nil && f.Name() == "init" {
						fv = func() {}
					} else {
						panic(fmt.Errorf("no code for function: %v", f))
					}
				} else {
					fv = ext
				}
			} else {
				fv = f
			}
		case *ssa.MakeClosure:
			var bindings []value
			for i, _ := range f.Bindings {
				bindings = append(bindings, fr.reg(ib[i]))
			}
			fv = &closure{f.Fn.(*ssa.Function), bindings}
		default:
			fv = fr.reg(iv)
		}
	} else {
		v := fr.reg(iv)
		rtype := reflect.TypeOf(v)
		mname := call.Method.Name()
		if mset, ok := i.msets[rtype]; ok {
			if f, ok := mset[mname]; ok {
				fv = f
			} else {
				ext, ok := findUserMethod(rtype, mname)
				if !ok {
					panic(fmt.Errorf("no code for method: %v.%v", rtype, mname))
				}
				fv = ext
			}
		} else {
			ext, ok := findExternMethod(rtype, mname)
			if !ok {
				panic(fmt.Errorf("no code for method: %v.%v", rtype, mname))
			}
			fv = ext
		}
		args = append(args, v)
	}
	for i, _ := range call.Args {
		v := fr.reg(ia[i])
		args = append(args, v)
	}
	return
}

// call interprets a call to a function (function, builtin or closure)
// fn with arguments args, returning its result.
// callpos is the position of the callsite.
//
func (i *Interp) call(caller *frame, fn value, args []value, ssaArgs []ssa.Value) value {
	switch fn := fn.(type) {
	case *ssa.Function:
		return i.callFunction(caller, fn, args, nil)
	case *closure:
		return i.callFunction(caller, fn.Fn, args, fn.Env)
	case *ssa.Builtin:
		return i.callBuiltin(caller, fn, args, ssaArgs)
	case reflect.Value:
		return i.callReflect(caller, fn, args, nil)
	default:
		return i.callReflect(caller, reflect.ValueOf(fn), args, nil)
	}
	panic(fmt.Sprintf("cannot call %T %v", fn, reflect.ValueOf(fn).Kind()))
}

func loc(fset *token.FileSet, pos token.Pos) string {
	if pos == token.NoPos {
		return ""
	}
	return " at " + fset.Position(pos).String()
}

func (i *Interp) callFunction(caller *frame, fn *ssa.Function, args []value, env []value) (result value) {
	fr := &frame{
		interp: i,
		caller: caller, // for panic/recover
		pfn:    i.funcs[fn],
	}
	if caller != nil {
		fr.deferid = caller.deferid
	}
	fr.stack = append([]value{}, fr.pfn.stack...)
	fr.block = fr.pfn.Fn.Blocks[0]
	var ip = 0
	for i := range fn.Params {
		fr.stack[ip] = args[i]
		ip++
	}
	for i := range fn.FreeVars {
		fr.stack[ip] = env[i]
		ip++
	}
	fr.run()
	// Destroy the locals to avoid accidental use after return.
	n := len(fr.results)
	if n == 1 {
		result = fr.stack[fr.results[0]]
	} else if n > 1 {
		res := make([]value, n, n)
		for i := 0; i < n; i++ {
			res[i] = fr.reg(fr.results[i])
		}
		result = tuple(res)
	}
	fr.stack = nil
	return
}

func (i *Interp) callReflect(caller *frame, fn reflect.Value, args []value, env []value) value {
	if caller != nil && caller.deferid != 0 {
		i.deferMap.Store(caller.deferid, caller)
	}
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
func (fr *frame) run() {
	if fr.pfn.Recover != nil {
		defer func() {
			if fr.pc == -1 {
				return // normal return
			}
			fr.panicking = &panicking{recover()}
			fr.runDefers()
			for _, fn := range fr.pfn.Recover {
				fn(fr)
			}
		}()
	}

	for fr.pc >= 0 {
		fn := fr.pfn.Instrs[fr.pc]
		fr.pc++
		fn(fr)
	}
}

// doRecover implements the recover() built-in.
func doRecover(caller *frame) value {
	// recover() must be exactly one level beneath the deferred
	// function (two levels beneath the panicking function) to
	// have any effect.  Thus we ignore both "defer recover()" and
	// "defer f() -> g() -> recover()".
	if caller.interp.mode&DisableRecover == 0 &&
		caller != nil && caller.panicking == nil &&
		caller.caller != nil && caller.caller.panicking != nil {
		p := caller.caller.panicking.value
		caller.caller.panicking = nil
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

func NewInterp(loader Loader, mainpkg *ssa.Package, mode Mode) (*Interp, error) {
	i := &Interp{
		fset:         mainpkg.Prog.Fset,
		prog:         mainpkg.Prog,
		mainpkg:      mainpkg,
		globals:      make(map[ssa.Value]value),
		mode:         mode,
		goroutines:   1,
		preloadTypes: make(map[types.Type]reflect.Type),
		funcs:        make(map[*ssa.Function]*Function),
		msets:        make(map[reflect.Type](map[string]*ssa.Function)),
	}
	i.loader = loader
	i.record = NewTypesRecord(i.loader, i)
	i.record.Load(mainpkg)

	var pkgs []*ssa.Package
	for _, pkg := range mainpkg.Prog.AllPackages() {
		// skip external pkg
		if pkg.Func("init").Blocks == nil {
			continue
		}
		pkgs = append(pkgs, pkg)
		// Initialize global storage.
		for _, m := range pkg.Members {
			switch v := m.(type) {
			case *ssa.Global:
				typ := i.preToType(deref(v.Type()))
				i.globals[v] = reflect.New(typ).Interface()
			}
		}
	}

	// static types check
	err := checkPackages(i, pkgs)
	if err != nil {
		return i, err
	}

	_, err = i.Run("init")
	if err != nil {
		err = fmt.Errorf("init error: %w", err)
	}
	return i, err
}

func (i *Interp) loadType(typ types.Type) {
	if _, ok := i.preloadTypes[typ]; !ok {
		i.preloadTypes[typ] = i.record.ToType(typ)
	}
}

func (i *Interp) preToType(typ types.Type) reflect.Type {
	if t, ok := i.preloadTypes[typ]; ok {
		return t
	}
	t := i.record.ToType(typ)
	i.preloadTypes[typ] = t
	return t
}

func (i *Interp) toType(typ types.Type) reflect.Type {
	if t, ok := i.preloadTypes[typ]; ok {
		return t
	}
	// log.Panicf("toType %v %p\n", typ, typ)
	i.typesMutex.Lock()
	defer i.typesMutex.Unlock()
	return i.record.ToType(typ)
}

func (i *Interp) RunFunc(name string, args ...Value) (r Value, err error) {
	defer func() {
		if i.mode&DisableRecover != 0 {
			return
		}
		switch p := recover().(type) {
		case nil:
			// nothing
		case exitPanic:
			// nothing
		case targetPanic:
			err = p
		case runtime.Error:
			err = p
		case string:
			err = plainError(p)
		case plainError:
			err = p
		default:
			err = fmt.Errorf("unexpected type: %T: %v", p, p)
		}
	}()
	if fn := i.mainpkg.Func(name); fn != nil {
		r = i.call(nil, fn, args, nil)
	} else {
		err = fmt.Errorf("no function %v", name)
	}
	return
}

func (i *Interp) Run(entry string) (exitCode int, err error) {
	// Top-level error handler.
	i.exited = false
	exitCode = 2
	defer func() {
		if i.exited {
			return
		}
		i.exited = true
		if i.mode&DisableRecover != 0 {
			return
		}
		switch p := recover().(type) {
		case nil:
			// nothing
		case exitPanic:
			exitCode = int(p)
		case targetPanic:
			err = p
		case runtime.Error:
			err = p
		case string:
			err = plainError(p)
		case plainError:
			err = p
		default:
			err = fmt.Errorf("unexpected type: %T: %v", p, p)
		}
	}()
	if mainFn := i.mainpkg.Func(entry); mainFn != nil {
		i.call(nil, mainFn, nil, nil)
		exitCode = 0
	} else {
		err = fmt.Errorf("no function %v", entry)
		exitCode = 1
	}
	return
}

func (i *Interp) GetFunc(key string) (interface{}, bool) {
	m, ok := i.mainpkg.Members[key]
	if !ok {
		return nil, false
	}
	fn, ok := m.(*ssa.Function)
	if !ok {
		return nil, false
	}
	return i.makeFunc(i.toType(fn.Type()), fn, nil).Interface(), true
}

func (i *Interp) GetVarAddr(key string) (interface{}, bool) {
	m, ok := i.mainpkg.Members[key]
	if !ok {
		return nil, false
	}
	v, ok := m.(*ssa.Global)
	if !ok {
		return nil, false
	}
	p, ok := i.globals[v]
	return p, ok
}

func (i *Interp) GetConst(key string) (constant.Value, bool) {
	m, ok := i.mainpkg.Members[key]
	if !ok {
		return nil, false
	}
	v, ok := m.(*ssa.NamedConst)
	if !ok {
		return nil, false
	}
	return v.Value.Value, true
}

func (i *Interp) GetType(key string) (reflect.Type, bool) {
	m, ok := i.mainpkg.Members[key]
	if !ok {
		return nil, false
	}
	t, ok := m.(*ssa.Type)
	if !ok {
		return nil, false
	}
	return i.toType(t.Type()), true
}

// deref returns a pointer's element type; otherwise it returns typ.
// TODO(adonovan): Import from ssa?
func deref(typ types.Type) types.Type {
	if p, ok := typ.Underlying().(*types.Pointer); ok {
		return p.Elem()
	}
	return typ
}
