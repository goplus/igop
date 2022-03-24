package gossa

import (
	"fmt"
	"go/token"
	"go/types"
	"os"
	"reflect"
	"strings"
	"sync/atomic"
	"unsafe"

	"github.com/goplus/reflectx"
	"golang.org/x/tools/go/ssa"
)

/*
type Op int

const (
	OpInvalid Op = iota
	// Value-defining instructions
	OpAlloc
	OpPhi
	OpCall
	OpBinOp
	OpUnOp
	OpChangeType
	OpConvert
	OpChangeInterface
	OpSliceToArrayPointer
	OpMakeInterface
	OpMakeClosure
	OpMakeMap
	OpMakeChan
	OpMakeSlice
	OpSlice
	OpFieldAddr
	OpField
	OpIndexAddr
	OpIndex
	OpLookup
	OpSelect
	OpRange
	OpNext
	OpTypeAssert
	OpExtract
	// Instructions executed for effect
	OpJump
	OpIf
	OpReturn
	OpRunDefers
	OpPanic
	OpGo
	OpDefer
	OpSend
	OpStore
	OpMapUpdate
	OpDebugRef
)
*/

type FuncBlock struct {
	Index  int
	Instrs []func(fr *frame, k *int)
	Succs  []*FuncBlock
}

type Function struct {
	Fn               *ssa.Function
	MainBlock        *FuncBlock
	Recover          *FuncBlock
	mapUnderscoreKey map[types.Type]bool
}

func findExternFunc(interp *Interp, fn *ssa.Function) (ext reflect.Value, ok bool) {
	fnName := fn.String()
	if fnName == "os.Exit" {
		return reflect.ValueOf(func(code int) {
			if interp.exited {
				os.Exit(code)
			} else {
				panic(exitPanic(code))
			}
		}), true
	}
	ext, ok = externValues[fnName]
	if !ok && fn.Pkg != nil {
		if recv := fn.Signature.Recv(); recv == nil {
			if pkg, found := interp.installed(fn.Pkg.Pkg.Path()); found {
				ext, ok = pkg.Funcs[fn.Name()]
			}
		} else if typ, found := interp.loader.LookupReflect(recv.Type()); found {
			if m, found := reflectx.MethodByName(typ, fn.Name()); found {
				ext, ok = m.Func, true
			}
		}
	}
	return
}

func makeInstr(interp *Interp, pfn *Function, instr ssa.Instruction) func(fr *frame, k *int) {
	switch instr := instr.(type) {
	case *ssa.Alloc:
		if instr.Heap {
			typ := interp.preToType(instr.Type()).Elem()
			return func(fr *frame, k *int) {
				fr.env[instr] = reflect.New(typ).Interface()
			}
		} else {
			typ := interp.preToType(instr.Type()).Elem()
			elem := reflect.New(typ).Elem()
			return func(fr *frame, k *int) {
				if v, ok := fr.env[instr]; ok {
					SetValue(reflect.ValueOf(v).Elem(), elem)
				} else {
					fr.env[instr] = reflect.New(typ).Interface()
				}
			}
		}
	case *ssa.Phi:
		return func(fr *frame, k *int) {
			for i, pred := range instr.Block().Preds {
				if fr.pred == pred.Index {
					fr.env[instr] = fr.get(instr.Edges[i])
					break
				}
			}
		}
	case *ssa.Call:
		return makeCallInstr(interp, instr, &instr.Call)
	case *ssa.BinOp:
		switch instr.Op {
		case token.ADD:
			return func(fr *frame, k *int) {
				fr.env[instr] = opADD(fr.get(instr.X), fr.get(instr.Y))
			}
		case token.SUB:
			return func(fr *frame, k *int) {
				fr.env[instr] = opSUB(fr.get(instr.X), fr.get(instr.Y))
			}
		case token.MUL:
			return func(fr *frame, k *int) {
				fr.env[instr] = opMUL(fr.get(instr.X), fr.get(instr.Y))
			}
		case token.QUO:
			return func(fr *frame, k *int) {
				fr.env[instr] = opQuo(fr.get(instr.X), fr.get(instr.Y))
			}
		case token.REM:
			return func(fr *frame, k *int) {
				fr.env[instr] = opREM(fr.get(instr.X), fr.get(instr.Y))
			}
		case token.AND:
			return func(fr *frame, k *int) {
				fr.env[instr] = opAND(fr.get(instr.X), fr.get(instr.Y))
			}
		case token.OR:
			return func(fr *frame, k *int) {
				fr.env[instr] = opOR(fr.get(instr.X), fr.get(instr.Y))
			}
		case token.XOR:
			return func(fr *frame, k *int) {
				fr.env[instr] = opXOR(fr.get(instr.X), fr.get(instr.Y))
			}
		case token.AND_NOT:
			return func(fr *frame, k *int) {
				fr.env[instr] = opANDNOT(fr.get(instr.X), fr.get(instr.Y))
			}
		case token.SHL:
			if c, ok := instr.Y.(*ssa.Convert); ok {
				xtyp := interp.preToType(c.X.Type())
				xk := xtyp.Kind()
				if xk >= reflect.Int && xk <= reflect.Int64 {
					return func(fr *frame, k *int) {
						v := reflect.ValueOf(fr.get(c.X))
						if v.Int() < 0 {
							panic(runtimeError("negative shift amount"))
						}
						fr.env[instr] = opSHL(fr.get(instr.X), fr.get(instr.Y))
					}
				}
			}
			return func(fr *frame, k *int) {
				fr.env[instr] = opSHL(fr.get(instr.X), fr.get(instr.Y))
			}
		case token.SHR:
			if c, ok := instr.Y.(*ssa.Convert); ok {
				xtyp := interp.preToType(c.X.Type())
				xk := xtyp.Kind()
				if xk >= reflect.Int && xk <= reflect.Int64 {
					return func(fr *frame, k *int) {
						v := reflect.ValueOf(fr.get(c.X))
						if v.Int() < 0 {
							panic(runtimeError("negative shift amount"))
						}
						fr.env[instr] = opSHR(fr.get(instr.X), fr.get(instr.Y))
					}
				}
			}
			return func(fr *frame, k *int) {
				fr.env[instr] = opSHR(fr.get(instr.X), fr.get(instr.Y))
			}
		case token.LSS:
			return func(fr *frame, k *int) {
				fr.env[instr] = opLSS(fr.get(instr.X), fr.get(instr.Y))
			}
		case token.LEQ:
			return func(fr *frame, k *int) {
				fr.env[instr] = opLEQ(fr.get(instr.X), fr.get(instr.Y))
			}
		case token.EQL:
			return func(fr *frame, k *int) {
				fr.env[instr] = opEQL(instr, fr.get(instr.X), fr.get(instr.Y))
			}
		case token.NEQ:
			return func(fr *frame, k *int) {
				fr.env[instr] = !opEQL(instr, fr.get(instr.X), fr.get(instr.Y))
			}
		case token.GTR:
			return func(fr *frame, k *int) {
				fr.env[instr] = opGTR(fr.get(instr.X), fr.get(instr.Y))
			}
		case token.GEQ:
			return func(fr *frame, k *int) {
				fr.env[instr] = opGEQ(fr.get(instr.X), fr.get(instr.Y))
			}
		default:
			panic(fmt.Errorf("unreachable %v", instr.Op))
		}
	case *ssa.UnOp:
		return func(fr *frame, k *int) {
			fr.env[instr] = unop(instr, fr.get(instr.X))
		}
	case *ssa.ChangeInterface:
		return func(fr *frame, k *int) {
			fr.env[instr] = fr.get(instr.X)
		}
	case *ssa.ChangeType:
		typ := interp.preToType(instr.Type())
		switch f := instr.X.(type) {
		case *ssa.Function:
			fn := interp.makeFunc(nil, interp.preToType(f.Type()), f, nil)
			v := fn.Convert(typ).Interface()
			return func(fr *frame, k *int) {
				fr.env[instr] = v
			}
		default:
			return func(fr *frame, k *int) {
				x := fr.get(instr.X)
				if x == nil {
					fr.env[instr] = reflect.New(typ).Elem().Interface()
				} else {
					fr.env[instr] = reflect.ValueOf(x).Convert(typ).Interface()
				}
			}
		}
	case *ssa.Convert:
		return makeConvertInstr(interp, instr)
		// typ := interp.preToType(instr.Type())
		// return func(fr *frame, k *int) {
		// 	x := fr.get(instr.X)
		// 	fr.env[instr] = convert(x, typ)
		// }
	case *ssa.MakeInterface:
		typ := interp.preToType(instr.Type())
		xtyp := interp.preToType(instr.X.Type())
		switch f := instr.X.(type) {
		case *ssa.Function:
			fn := interp.makeFunc(nil, xtyp, f, nil)
			if typ == tyEmptyInterface {
				return func(fr *frame, k *int) {
					fr.env[instr] = fn.Interface()
				}
			}
			return func(fr *frame, k *int) {
				v := reflect.New(typ).Elem()
				SetValue(v, fn)
				fr.env[instr] = v.Interface()
			}
		default:
			if typ == tyEmptyInterface {
				return func(fr *frame, k *int) {
					x := fr.get(instr.X)
					fr.env[instr] = x
				}
			}
			return func(fr *frame, k *int) {
				v := reflect.New(typ).Elem()
				if x := fr.get(instr.X); x != nil {
					SetValue(v, reflect.ValueOf(x))
				}
				fr.env[instr] = v.Interface()
			}
		}
	case *ssa.MakeClosure:
		fn := instr.Fn.(*ssa.Function)
		typ := interp.preToType(fn.Type())
		return func(fr *frame, k *int) {
			var bindings []value
			for _, binding := range instr.Bindings {
				bindings = append(bindings, fr.get(binding))
			}
			c := &closure{fn, bindings}
			fr.env[instr] = interp.makeFunc(fr, typ, c.Fn, c.Env).Interface()
		}
	case *ssa.MakeChan:
		typ := interp.preToType(instr.Type())
		return func(fr *frame, k *int) {
			size := fr.get(instr.Size)
			buffer := asInt(size)
			if buffer < 0 {
				panic(runtimeError("makechan: size out of range"))
			}
			fr.env[instr] = reflect.MakeChan(typ, buffer).Interface()
		}
	case *ssa.MakeMap:
		typ := instr.Type()
		rtyp := interp.preToType(typ)
		key := typ.Underlying().(*types.Map).Key()
		if st, ok := key.Underlying().(*types.Struct); ok && hasUnderscore(st) {
			pfn.mapUnderscoreKey[typ] = true
		}
		if instr.Reserve == nil {
			return func(fr *frame, k *int) {
				fr.env[instr] = reflect.MakeMap(rtyp).Interface()
			}
		} else {
			return func(fr *frame, k *int) {
				reserve := asInt(fr.get(instr.Reserve))
				fr.env[instr] = reflect.MakeMapWithSize(rtyp, reserve).Interface()
			}
		}
	case *ssa.MakeSlice:
		typ := interp.preToType(instr.Type())
		return func(fr *frame, k *int) {
			Len := asInt(fr.get(instr.Len))
			if Len < 0 || Len >= maxMemLen {
				panic(runtimeError("makeslice: len out of range"))
			}
			Cap := asInt(fr.get(instr.Cap))
			if Cap < 0 || Cap >= maxMemLen {
				panic(runtimeError("makeslice: cap out of range"))
			}
			fr.env[instr] = reflect.MakeSlice(typ, Len, Cap).Interface()
		}
	case *ssa.Slice:
		typ := interp.preToType(instr.Type())
		isNamed := typ.Kind() == reflect.Slice && typ != reflect.SliceOf(typ.Elem())
		_, makesliceCheck := instr.X.(*ssa.Alloc)
		if isNamed {
			return func(fr *frame, k *int) {
				fr.env[instr] = slice(fr, instr, makesliceCheck).Convert(typ).Interface()
			}
		} else {
			return func(fr *frame, k *int) {
				fr.env[instr] = slice(fr, instr, makesliceCheck).Interface()
			}
		}
	case *ssa.FieldAddr:
		return func(fr *frame, k *int) {
			v, err := FieldAddr(fr.get(instr.X), instr.Field)
			if err != nil {
				panic(runtimeError(err.Error()))
			}
			fr.env[instr] = v
		}
	case *ssa.Field:
		return func(fr *frame, k *int) {
			v, err := Field(fr.get(instr.X), instr.Field)
			if err != nil {
				panic(runtimeError(err.Error()))
			}
			fr.env[instr] = v
		}
	case *ssa.IndexAddr:
		return func(fr *frame, k *int) {
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
		}
	case *ssa.Index:
		return func(fr *frame, k *int) {
			x := fr.get(instr.X)
			idx := fr.get(instr.Index)
			v := reflect.ValueOf(x)
			fr.env[instr] = v.Index(asInt(idx)).Interface()
		}
	case *ssa.Lookup:
		typ := interp.preToType(instr.X.Type())
		switch typ.Kind() {
		case reflect.String:
			return func(fr *frame, k *int) {
				v := fr.get(instr.X)
				idx := fr.get(instr.Index)
				fr.env[instr] = reflect.ValueOf(v).String()[asInt(idx)]
			}
		case reflect.Map:
			return func(fr *frame, k *int) {
				m := fr.get(instr.X)
				idx := fr.get(instr.Index)
				vm := reflect.ValueOf(m)
				v := vm.MapIndex(reflect.ValueOf(idx))
				ok := v.IsValid()
				var rv value
				if ok {
					rv = v.Interface()
				} else {
					rv = reflect.New(typ.Elem()).Elem().Interface()
				}
				if instr.CommaOk {
					fr.env[instr] = tuple{rv, ok}
				} else {
					fr.env[instr] = rv
				}
			}
		default:
			panic("unreachable")
		}
	case *ssa.Select:
		return func(fr *frame, k *int) {
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
			for n, st := range instr.States {
				if st.Dir == types.RecvOnly {
					var v value
					if n == chosen && recvOk {
						// No need to copy since send makes an unaliased copy.
						v = recv.Interface()
					} else {
						typ := interp.toType(st.Chan.Type())
						v = reflect.New(typ.Elem()).Elem().Interface()
					}
					r = append(r, v)
				}
			}
			fr.env[instr] = r
		}
	case *ssa.SliceToArrayPointer:
		typ := interp.preToType(instr.Type())
		return func(fr *frame, k *int) {
			x := fr.get(instr.X)
			v := reflect.ValueOf(x)
			vLen := v.Len()
			tLen := typ.Elem().Len()
			if tLen > vLen {
				panic(runtimeError(fmt.Sprintf("cannot convert slice with length %v to pointer to array with length %v", vLen, tLen)))
			}
			fr.env[instr] = v.Convert(typ).Interface()
		}
	case *ssa.Range:
		typ := interp.preToType(instr.X.Type())
		switch typ.Kind() {
		case reflect.String:
			return func(fr *frame, k *int) {
				v := fr.get(instr.X)
				fr.env[instr] = &stringIter{Reader: strings.NewReader(reflect.ValueOf(v).String())}
			}
		case reflect.Map:
			return func(fr *frame, k *int) {
				v := fr.get(instr.X)
				fr.env[instr] = &mapIter{iter: reflect.ValueOf(v).MapRange()}
			}
		default:
			panic("unreachable")
		}
	case *ssa.Next:
		if instr.IsString {
			return func(fr *frame, k *int) {
				fr.env[instr] = fr.get(instr.Iter).(*stringIter).next()
			}
		} else {
			return func(fr *frame, k *int) {
				fr.env[instr] = fr.get(instr.Iter).(*mapIter).next()
			}
		}
	case *ssa.TypeAssert:
		typ := interp.preToType(instr.AssertedType)
		if fn, ok := instr.X.(*ssa.Function); ok {
			f := interp.makeFunc(nil, interp.preToType(fn.Type()), fn, nil).Interface()
			return func(fr *frame, k *int) {
				fr.env[instr] = typeAssert(interp, instr, typ, f)
			}
		} else {
			return func(fr *frame, k *int) {
				v := fr.get(instr.X)
				fr.env[instr] = typeAssert(interp, instr, typ, v)
			}
		}
	case *ssa.Extract:
		if *instr.Referrers() == nil {
			return nil
		}
		return func(fr *frame, k *int) {
			fr.env[instr] = fr.get(instr.Tuple).(tuple)[instr.Index]
		}
	// Instructions executed for effect
	case *ssa.Jump:
		return func(fr *frame, k *int) {
			fr.pred, fr.block = fr.block.Index, fr.block.Succs[0]
			*k = kJump
		}
	case *ssa.If:
		switch instr.Cond.Type().(type) {
		case *types.Basic:
			return func(fr *frame, k *int) {
				succ := 1
				if v := fr.get(instr.Cond); v.(bool) {
					succ = 0
				}
				fr.pred, fr.block = fr.block.Index, fr.block.Succs[succ]
				*k = kJump
			}
		default:
			return func(fr *frame, k *int) {
				succ := 1
				if v := fr.get(instr.Cond); reflect.ValueOf(v).Bool() {
					succ = 0
				}
				fr.pred, fr.block = fr.block.Index, fr.block.Succs[succ]
				*k = kJump
			}
		}
	case *ssa.Return:
		switch n := len(instr.Results); n {
		case 0:
			return func(fr *frame, k *int) {
				fr.block = nil
				*k = kReturn
			}
		case 1:
			return func(fr *frame, k *int) {
				fr.result = fr.get(instr.Results[0])
				fr.block = nil
				*k = kReturn
			}
		default:
			return func(fr *frame, k *int) {
				res := make([]value, n, n)
				for i := 0; i < n; i++ {
					res[i] = fr.get(instr.Results[i])
				}
				fr.result = tuple(res)
				fr.block = nil
				*k = kReturn
			}
		}
	case *ssa.RunDefers:
		return func(fr *frame, k *int) {
			fr.runDefers()
		}
	case *ssa.Panic:
		return func(fr *frame, k *int) {
			panic(targetPanic{fr.get(instr.X)})
		}
	case *ssa.Go:
		return func(fr *frame, k *int) {
			fn, args := interp.prepareCall(fr, &instr.Call)
			atomic.AddInt32(&interp.goroutines, 1)
			go func() {
				interp.call(nil, instr.Pos(), fn, args, instr.Call.Args)
				atomic.AddInt32(&interp.goroutines, -1)
			}()
		}
	case *ssa.Defer:
		return func(fr *frame, k *int) {
			fn, args := interp.prepareCall(fr, &instr.Call)
			fr.defers = &deferred{
				fn:      fn,
				args:    args,
				ssaArgs: instr.Call.Args,
				instr:   instr,
				tail:    fr.defers,
			}
		}
	case *ssa.Send:
		return func(fr *frame, k *int) {
			c := fr.get(instr.Chan)
			x := fr.get(instr.X)
			ch := reflect.ValueOf(c)
			if x == nil {
				ch.Send(reflect.New(ch.Type().Elem()).Elem())
			} else {
				ch.Send(reflect.ValueOf(x))
			}
		}
	case *ssa.Store:
		// skip struct field _
		if addr, ok := instr.Addr.(*ssa.FieldAddr); ok {
			if s, ok := addr.X.Type().(*types.Pointer).Elem().(*types.Struct); ok {
				if s.Field(addr.Field).Name() == "_" {
					return nil
				}
			}
		}
		if fn, ok := instr.Val.(*ssa.Function); ok {
			v := interp.makeFunc(nil, interp.preToType(fn.Type()), fn, nil)
			return func(fr *frame, k *int) {
				x := reflect.ValueOf(fr.get(instr.Addr))
				SetValue(x.Elem(), v)
			}
		}
		return func(fr *frame, k *int) {
			x := reflect.ValueOf(fr.get(instr.Addr))
			val := fr.get(instr.Val)
			switch fn := val.(type) {
			case *ssa.Function:
				f := interp.makeFunc(fr, interp.toType(fn.Type()), fn, nil)
				SetValue(x.Elem(), f)
			default:
				v := reflect.ValueOf(val)
				if v.IsValid() {
					SetValue(x.Elem(), v)
				} else {
					SetValue(x.Elem(), reflect.New(x.Elem().Type()).Elem())
				}
			}
		}
	case *ssa.MapUpdate:
		if pfn.mapUnderscoreKey[instr.Map.Type()] {
			if fn, ok := instr.Value.(*ssa.Function); ok {
				v := interp.makeFunc(nil, interp.preToType(fn.Type()), fn, nil)
				return func(fr *frame, k *int) {
					vm := reflect.ValueOf(fr.get(instr.Map))
					vk := reflect.ValueOf(fr.get(instr.Key))
					for _, vv := range vm.MapKeys() {
						if equalStruct(vk, vv) {
							vk = vv
							break
						}
					}
					vm.SetMapIndex(vk, reflect.ValueOf(v))
				}
			}
			return func(fr *frame, k *int) {
				vm := reflect.ValueOf(fr.get(instr.Map))
				vk := reflect.ValueOf(fr.get(instr.Key))
				v := fr.get(instr.Value)
				for _, vv := range vm.MapKeys() {
					if equalStruct(vk, vv) {
						vk = vv
						break
					}
				}
				vm.SetMapIndex(vk, reflect.ValueOf(v))
			}
		} else {
			if fn, ok := instr.Value.(*ssa.Function); ok {
				v := interp.makeFunc(nil, interp.preToType(fn.Type()), fn, nil)
				return func(fr *frame, k *int) {
					vm := reflect.ValueOf(fr.get(instr.Map))
					vk := reflect.ValueOf(fr.get(instr.Key))
					vm.SetMapIndex(vk, reflect.ValueOf(v))
				}
			} else {
				return func(fr *frame, k *int) {
					vm := reflect.ValueOf(fr.get(instr.Map))
					vk := reflect.ValueOf(fr.get(instr.Key))
					v := fr.get(instr.Value)
					vm.SetMapIndex(vk, reflect.ValueOf(v))
				}
			}
		}
	case *ssa.DebugRef:
		if interp.fnDebug == nil {
			return nil
		} else {
			return func(fr *frame, k *int) {
				ref := &DebugInfo{DebugRef: instr, fset: interp.fset}
				ref.toValue = func() (*types.Var, interface{}, bool) {
					if v, ok := instr.Object().(*types.Var); ok {
						return v, fr.get(instr.X), true
					}
					return nil, nil, false
				}
				interp.fnDebug(ref)
			}
		}
	default:
		panic(fmt.Errorf("unreachable %T", instr))
	}
}

func makeConvertInstr(interp *Interp, instr *ssa.Convert) func(fr *frame, k *int) {
	typ := interp.preToType(instr.Type())
	xtyp := interp.preToType(instr.X.Type())
	vk := xtyp.Kind()
	switch typ.Kind() {
	case reflect.UnsafePointer:
		if vk == reflect.Uintptr {
			return func(fr *frame, k *int) {
				v := reflect.ValueOf(fr.get(instr.X))
				fr.env[instr] = toUnsafePointer(v)
			}
		} else if vk == reflect.Ptr {
			return func(fr *frame, k *int) {
				v := reflect.ValueOf(fr.get(instr.X))
				fr.env[instr] = unsafe.Pointer(v.Pointer())
			}
		}
	case reflect.Uintptr:
		if vk == reflect.UnsafePointer {
			return func(fr *frame, k *int) {
				v := reflect.ValueOf(fr.get(instr.X))
				fr.env[instr] = v.Pointer()
			}
		}
	case reflect.Ptr:
		if vk == reflect.UnsafePointer {
			return func(fr *frame, k *int) {
				v := reflect.ValueOf(fr.get(instr.X))
				fr.env[instr] = reflect.NewAt(typ.Elem(), unsafe.Pointer(v.Pointer())).Interface()
			}
		}
	case reflect.Slice:
		if vk == reflect.String {
			elem := typ.Elem()
			switch elem.Kind() {
			case reflect.Uint8:
				if elem.PkgPath() != "" {
					return func(fr *frame, k *int) {
						v := reflect.ValueOf(fr.get(instr.X))
						dst := reflect.New(typ).Elem()
						dst.SetBytes([]byte(v.String()))
						fr.env[instr] = dst.Interface()
					}
				}
			case reflect.Int32:
				if elem.PkgPath() != "" {
					return func(fr *frame, k *int) {
						v := reflect.ValueOf(fr.get(instr.X))
						dst := reflect.New(typ).Elem()
						*(*[]rune)((*reflectValue)(unsafe.Pointer(&dst)).ptr) = []rune(v.String())
						fr.env[instr] = dst.Interface()
					}
				}
			}
		}
	case reflect.String:
		if vk == reflect.Slice {
			elem := xtyp.Elem()
			switch elem.Kind() {
			case reflect.Uint8:
				if elem.PkgPath() != "" {
					return func(fr *frame, k *int) {
						v := reflect.ValueOf(fr.get(instr.X))
						v = reflect.ValueOf(string(v.Bytes()))
						fr.env[instr] = v.Convert(typ).Interface()
					}
				}
			case reflect.Int32:
				if elem.PkgPath() != "" {
					return func(fr *frame, k *int) {
						v := reflect.ValueOf(fr.get(instr.X))
						v = reflect.ValueOf(*(*[]rune)(((*reflectValue)(unsafe.Pointer(&v))).ptr))
						fr.env[instr] = v.Convert(typ).Interface()
					}
				}
			}
		}
	}
	return func(fr *frame, k *int) {
		v := reflect.ValueOf(fr.get(instr.X))
		fr.env[instr] = v.Convert(typ).Interface()
	}
}

func makeCallInstr1(interp *Interp, instr ssa.Value, call *ssa.CallCommon) func(fr *frame, k *int) {
	return func(fr *frame, k *int) {
		fn, args := interp.prepareCall(fr, call)
		fr.env[instr] = interp.call(fr, call.Pos(), fn, args, call.Args)
	}
}

func makeCallInstr(interp *Interp, instr ssa.Value, call *ssa.CallCommon) func(fr *frame, k *int) {
	pos := instr.Pos()
	switch fn := call.Value.(type) {
	case *ssa.Builtin:
		nargs := len(call.Args)
		return func(fr *frame, k *int) {
			args := make([]value, nargs, nargs)
			for i := 0; i < nargs; i++ {
				args[i] = fr.getParam(call.Args[i])
			}
			fr.env[instr] = interp.callBuiltin(fr, pos, fn, args, call.Args)
		}
	case *ssa.MakeClosure:
		pfn := fn.Fn.(*ssa.Function)
		nargs := len(call.Args)
		return func(fr *frame, k *int) {
			var env []value
			for _, binding := range fn.Bindings {
				env = append(env, fr.get(binding))
			}
			args := make([]value, nargs, nargs)
			for i := 0; i < nargs; i++ {
				args[i] = fr.getParam(call.Args[i])
			}
			fr.env[instr] = interp.callFunction(fr, pos, pfn, args, env)
		}
	case *ssa.Function:
		// "static func/method call"
		nargs := len(call.Args)
		if fn.Blocks == nil {
			ext, ok := findExternFunc(interp, fn)
			if !ok {
				// skip pkg.init
				if fn.Pkg != nil && fn.Name() == "init" {
					return nil
				}
				panic(fmt.Errorf("no code for function: %v", fn))
			}
			return func(fr *frame, k *int) {
				args := make([]value, nargs, nargs)
				for i := 0; i < nargs; i++ {
					args[i] = fr.getParam(call.Args[i])
				}
				fr.env[instr] = interp.callReflect(fr, pos, ext, args, nil)
			}
		}
		return func(fr *frame, k *int) {
			args := make([]value, nargs, nargs)
			for i := 0; i < nargs; i++ {
				args[i] = fr.getParam(call.Args[i])
			}
			fr.env[instr] = interp.callFunction(fr, pos, fn, args, nil)
		}
	}
	// "dynamic method call" // ("invoke" mode)
	if call.IsInvoke() {
		return makeCallMethodInstr(interp, instr, call)
	}
	// dynamic func call
	typ := interp.preToType(call.Value.Type())
	if typ.Kind() != reflect.Func {
		panic("unsupport")
	}
	nargs := len(call.Args)
	return func(fr *frame, k *int) {
		fn := fr.get(call.Value)
		args := make([]value, nargs, nargs)
		for i := 0; i < nargs; i++ {
			args[i] = fr.getParam(call.Args[i])
		}
		switch fn := fn.(type) {
		case *ssa.Function:
			fr.env[instr] = interp.callFunction(fr, pos, fn, args, nil)
		case *closure:
			fr.env[instr] = interp.callFunction(fr, pos, fn.Fn, args, fn.Env)
		case *ssa.Builtin:
			fr.env[instr] = interp.callBuiltin(fr, pos, fn, args, call.Args)
		default:
			fr.env[instr] = interp.callReflect(fr, pos, reflect.ValueOf(fn), args, nil)
		}
	}
}

var (
	typeOfType      = reflect.TypeOf(reflect.TypeOf(0))
	vfnMethod       = reflect.ValueOf(reflectx.MethodByIndex)
	vfnMethodByName = reflect.ValueOf(reflectx.MethodByName)
)

func findUserMethod(typ reflect.Type, name string) (ext reflect.Value, ok bool) {
	if m, ok := reflectx.MethodByName(typ, name); ok {
		return m.Func, true
	}
	return
}

func findExternMethod(typ reflect.Type, name string) (ext reflect.Value, ok bool) {
	if typ == typeOfType {
		switch name {
		case "Method":
			return vfnMethod, true
		case "MethodByName":
			return vfnMethodByName, true
		}
	}
	if m, ok := typ.MethodByName(name); ok {
		return m.Func, true
	}
	return
}

// skip unexport method
// func isSkipMethod(interp *Interp, pkg *types.Package, mname string) bool {
// 	if pkg != nil {
// 		if _, found := interp.installed(pkg.Path()); found {
// 			if !token.IsExported(mname) {
// 				return true
// 			}
// 		}
// 	}
// 	return false
// }

func (i *Interp) findMethod(typ reflect.Type, mname string) (fn *ssa.Function, ok bool) {
	if mset, mok := i.msets[typ]; mok {
		fn, ok = mset[mname]
	}
	return
}

func makeCallMethodInstr(interp *Interp, instr ssa.Value, call *ssa.CallCommon) func(fr *frame, k *int) {
	mname := call.Method.Name()
	nargs := len(call.Args)
	margs := nargs + 1
	pos := instr.Pos()
	var found bool
	var ext reflect.Value
	return func(fr *frame, k *int) {
		v := fr.get(call.Value)
		rtype := reflect.TypeOf(v)
		// find user type method *ssa.Function
		if mset, ok := interp.msets[rtype]; ok {
			if fn, ok := mset[mname]; ok {
				args := make([]value, margs, margs)
				args[0] = v
				for i := 0; i < nargs; i++ {
					args[i+1] = fr.getParam(call.Args[i])
				}
				fr.env[instr] = interp.callFunction(fr, pos, fn, args, nil)
				return
			}
			ext, found = findUserMethod(rtype, mname)
		} else {
			ext, found = findExternMethod(rtype, mname)
		}
		if !found {
			panic(fmt.Errorf("no code for method: %v.%v", rtype, mname))
		}
		args := make([]value, margs, margs)
		args[0] = v
		for i := 0; i < nargs; i++ {
			args[i+1] = fr.getParam(call.Args[i])
		}
		fr.env[instr] = interp.callReflect(fr, pos, ext, args, nil)
	}
}
