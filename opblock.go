package gossa

import (
	"fmt"
	"go/token"
	"go/types"
	"reflect"
	"sync/atomic"

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
	Instrs []func(fr *frame, k *int)
}

func findExternFunc(interp *Interp, fn *ssa.Function) (ext reflect.Value, ok bool) {
	ext, ok = externValues[fn.String()]
	if !ok && fn.Pkg != nil {
		if pkg, found := interp.installed(fn.Pkg.Pkg.Path()); found {
			if recv := fn.Signature.Recv(); recv == nil {
				ext, ok = pkg.Funcs[fn.Name()]
			} else if typ, found := interp.loader.LookupReflect(recv.Type()); found {
				if m, found := reflectx.MethodByName(typ, fn.Name()); found {
					ext, ok = m.Func, true
				}
			}
		}
	}
	return
}

func makeInstr(interp *Interp, instr ssa.Instruction) func(fr *frame, k *int) {
	switch instr := instr.(type) {
	case *ssa.Alloc:
		if instr.Heap {
			typ := interp.preToType(instr.Type()).Elem()
			return func(fr *frame, k *int) {
				fr.env[instr] = reflect.New(typ).Interface()
			}
		} else {
			return func(fr *frame, k *int) {
				v := reflect.ValueOf(fr.env[instr])
				SetValue(v.Elem(), fr.locals[instr])
			}
		}
	case *ssa.Phi:
		return func(fr *frame, k *int) {
			for i, pred := range instr.Block().Preds {
				if fr.prevBlock == pred {
					fr.env[instr] = fr.get(instr.Edges[i])
					break
				}
			}
		}
	case *ssa.Call:
		pos := instr.Pos()
		if instr.Call.Method == nil {
			switch fn := instr.Call.Value.(type) {
			case *ssa.Function:
				if fn.Blocks == nil {
					ext, ok := findExternFunc(interp, fn)
					if !ok {
						// skip pkg.init
						if fn.Pkg != nil && fn.Name() == "init" && fn.Type().String() == "func()" {
							return nil
						}
						panic(fmt.Errorf("no code for function: %v", fn))
					}
					nargs := len(instr.Call.Args)
					return func(fr *frame, k *int) {
						args := make([]value, nargs, nargs)
						for i := 0; i < nargs; i++ {
							v := fr.get(instr.Call.Args[i])
							if fn, ok := v.(*ssa.Function); ok {
								v = interp.toFunc(fr, interp.toType(fn.Type()), fn).Interface()
							}
							args[i] = v
						}
						fr.env[instr] = interp.callReflect(fr, pos, ext, args, nil)
					}
				}
			case *ssa.Builtin:
				nargs := len(instr.Call.Args)
				return func(fr *frame, k *int) {
					args := make([]value, nargs, nargs)
					for i := 0; i < nargs; i++ {
						v := fr.get(instr.Call.Args[i])
						if fn, ok := v.(*ssa.Function); ok {
							v = interp.toFunc(fr, interp.toType(fn.Type()), fn).Interface()
						}
						args[i] = v
					}
					fr.env[instr] = interp.callBuiltin(fr, pos, fn, args, instr.Call.Args)
				}
			case *ssa.MakeClosure:
			default:
				typ := interp.preToType(instr.Call.Value.Type())
				if typ.Kind() != reflect.Func {
					panic("unsupport")
				}
				// nargs := len(instr.Call.Args)
			}
		}
		return func(fr *frame, k *int) {
			fn, args := interp.prepareCall(fr, &instr.Call)
			fr.env[instr] = interp.call(fr, instr.Pos(), fn, args, instr.Call.Args)
		}
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
		return func(fr *frame, k *int) {
			x := fr.get(instr.X)
			var v reflect.Value
			switch f := x.(type) {
			case *ssa.Function:
				v = interp.toFunc(fr, interp.toType(f.Type()), f)
			default:
				v = reflect.ValueOf(x)
			}
			if !v.IsValid() {
				fr.env[instr] = reflect.New(typ).Elem()
			} else {
				fr.env[instr] = v.Convert(typ).Interface()
			}
		}
	case *ssa.Convert:
		typ := interp.preToType(instr.Type())
		return func(fr *frame, k *int) {
			x := fr.get(instr.X)
			fr.env[instr] = convert(x, typ)
		}
	case *ssa.MakeInterface:
		typ := interp.preToType(instr.Type())
		xtyp := interp.preToType(instr.X.Type())
		return func(fr *frame, k *int) {
			v := reflect.New(typ).Elem()
			x := fr.get(instr.X)
			var vx reflect.Value
			switch x.(type) {
			case *ssa.Function:
				vx = interp.toFunc(fr, xtyp, x)
			case nil:
				vx = reflect.New(xtyp).Elem()
			default:
				vx = reflect.ValueOf(x)
				if xtyp != vx.Type() {
					vx = reflect.ValueOf(convert(x, xtyp))
				}
			}
			SetValue(v, vx)
			fr.env[instr] = v.Interface()
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
			fr.env[instr] = interp.toFunc(fr, typ, c).Interface()
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
		var reserve int
		return func(fr *frame, k *int) {
			if instr.Reserve != nil {
				reserve = asInt(fr.get(instr.Reserve))
			}
			if st, ok := key.Underlying().(*types.Struct); ok && hasUnderscore(st) {
				fr.mapUnderscoreKey[typ] = true
			}
			fr.env[instr] = reflect.MakeMapWithSize(rtyp, reserve).Interface()
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
		return func(fr *frame, k *int) {
			fr.env[instr] = slice(fr, instr)
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
			if v.Kind() == reflect.Ptr {
				v = v.Elem()
			}
			fr.env[instr] = v.Index(asInt(idx)).Interface()
		}
	case *ssa.Lookup:
		return func(fr *frame, k *int) {
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
		return func(fr *frame, k *int) {
			fr.env[instr] = rangeIter(fr.get(instr.X), instr.X.Type())
		}
	case *ssa.Next:
		return func(fr *frame, k *int) {
			fr.env[instr] = fr.get(instr.Iter).(iter).next()
		}
	case *ssa.TypeAssert:
		return func(fr *frame, k *int) {
			v := fr.get(instr.X)
			if fn, ok := v.(*ssa.Function); ok {
				typ := interp.toType(fn.Type())
				v = interp.toFunc(fr, typ, fn).Interface()
			}
			fr.env[instr] = typeAssert(interp, instr, v)
		}
	case *ssa.Extract:
		return func(fr *frame, k *int) {
			fr.env[instr] = fr.get(instr.Tuple).(tuple)[instr.Index]
		}
	// Instructions executed for effect
	case *ssa.Jump:
		return func(fr *frame, k *int) {
			fr.prevBlock, fr.block = fr.block, fr.block.Succs[0]
			fr.fnBlock = interp.blocks[fr.block]
			fr.prevFnBlock = interp.blocks[fr.prevBlock]
			*k = kJump
		}
	case *ssa.If:
		return func(fr *frame, k *int) {
			succ := 1
			if v := fr.get(instr.Cond); reflect.ValueOf(v).Bool() {
				succ = 0
			}
			fr.prevBlock, fr.block = fr.block, fr.block.Succs[succ]
			fr.fnBlock = interp.blocks[fr.block]
			fr.prevFnBlock = interp.blocks[fr.prevBlock]
			*k = kJump
		}
	case *ssa.Return:
		switch n := len(instr.Results); n {
		case 0:
			return func(fr *frame, k *int) {
				fr.block = nil
				fr.fnBlock = nil
				*k = kReturn
			}
		case 1:
			return func(fr *frame, k *int) {
				fr.result = fr.get(instr.Results[0])
				fr.block = nil
				fr.fnBlock = nil
				*k = kReturn
			}
		default:
			res := make([]value, n, n)
			return func(fr *frame, k *int) {
				for i := 0; i < n; i++ {
					res[i] = fr.get(instr.Results[i])
				}
				fr.result = tuple(res)
				fr.block = nil
				fr.fnBlock = nil
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
		return func(fr *frame, k *int) {
			x := reflect.ValueOf(fr.get(instr.Addr))
			val := fr.get(instr.Val)
			switch fn := val.(type) {
			case *ssa.Function:
				f := interp.toFunc(fr, interp.toType(fn.Type()), fn)
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
		return func(fr *frame, k *int) {
			vm := reflect.ValueOf(fr.get(instr.Map))
			vk := reflect.ValueOf(fr.get(instr.Key))
			v := fr.get(instr.Value)
			if fn, ok := v.(*ssa.Function); ok {
				typ := interp.toType(fn.Type())
				v = interp.toFunc(fr, typ, fn).Interface()
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
