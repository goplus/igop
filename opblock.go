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
	Interp           *Interp
	Fn               *ssa.Function
	MainBlock        *FuncBlock
	Recover          *FuncBlock
	mapUnderscoreKey map[types.Type]bool
	index            map[ssa.Value]int
	stack            []value
}

func (p *Function) regIndex(v ssa.Value) int {
	if i, ok := p.index[v]; ok {
		return i
	}
	i := len(p.stack)
	p.index[v] = i
	switch v := v.(type) {
	case *ssa.Const:
		p.stack = append(p.stack, constToValue(p.Interp, v))
	case *ssa.Global:
		g, _ := globalToValue(p.Interp, v)
		p.stack = append(p.stack, g)
	case *ssa.Function:
		typ := p.Interp.preToType(v.Type())
		fn := p.Interp.makeFunc(nil, typ, v, nil).Interface()
		p.stack = append(p.stack, fn)
	default:
		p.stack = append(p.stack, nil)
	}
	return i
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
			ir := pfn.regIndex(instr)
			return func(fr *frame, k *int) {
				//fr.env[instr] = reflect.New(typ).Interface()
				fr.setReg(ir, reflect.New(typ).Interface())
			}
		} else {
			typ := interp.preToType(instr.Type()).Elem()
			elem := reflect.New(typ).Elem()
			ir := pfn.regIndex(instr)
			return func(fr *frame, k *int) {
				// if v, ok := fr.env[instr]; ok {
				// 	SetValue(reflect.ValueOf(v).Elem(), elem)
				// } else {
				// 	fr.env[instr] = reflect.New(typ).Interface()
				// }
				if v := fr.reg(ir); v != nil {
					SetValue(reflect.ValueOf(v).Elem(), elem)
				} else {
					fr.setReg(ir, reflect.New(typ).Interface())
				}
			}
		}
	case *ssa.Phi:
		ir := pfn.regIndex(instr)
		ie := make([]int, len(instr.Edges))
		for i, v := range instr.Edges {
			ie[i] = pfn.regIndex(v)
		}
		return func(fr *frame, k *int) {
			for i, pred := range instr.Block().Preds {
				if fr.pred == pred.Index {
					//fr.env[instr] = fr.get(instr.Edges[i])
					fr.setReg(ir, fr.reg(ie[i]))
					break
				}
			}
		}
	case *ssa.Call:
		return makeCallInstr(pfn, interp, instr, &instr.Call)
	case *ssa.BinOp:
		ir := pfn.regIndex(instr)
		ix := pfn.regIndex(instr.X)
		iy := pfn.regIndex(instr.Y)
		switch instr.Op {
		case token.ADD:
			return func(fr *frame, k *int) {
				//fr.env[instr] = opADD(fr.get(instr.X), fr.get(instr.Y))
				fr.setReg(ir, opADD(fr.reg(ix), fr.reg(iy)))
			}
		case token.SUB:
			return func(fr *frame, k *int) {
				// fr.env[instr] = opSUB(fr.get(instr.X), fr.get(instr.Y))
				fr.setReg(ir, opSUB(fr.reg(ix), fr.reg(iy)))
			}
		case token.MUL:
			return func(fr *frame, k *int) {
				//fr.env[instr] = opMUL(fr.get(instr.X), fr.get(instr.Y))
				fr.setReg(ir, opMUL(fr.reg(ix), fr.reg(iy)))
			}
		case token.QUO:
			return func(fr *frame, k *int) {
				// fr.env[instr] = opQuo(fr.get(instr.X), fr.get(instr.Y))
				fr.setReg(ir, opQuo(fr.reg(ix), fr.reg(iy)))
			}
		case token.REM:
			return func(fr *frame, k *int) {
				// fr.env[instr] = opREM(fr.get(instr.X), fr.get(instr.Y))
				fr.setReg(ir, opREM(fr.reg(ix), fr.reg(iy)))
			}
		case token.AND:
			return func(fr *frame, k *int) {
				// fr.env[instr] = opAND(fr.get(instr.X), fr.get(instr.Y))
				fr.setReg(ir, opAND(fr.reg(ix), fr.reg(iy)))
			}
		case token.OR:
			return func(fr *frame, k *int) {
				// fr.env[instr] = opOR(fr.get(instr.X), fr.get(instr.Y))
				fr.setReg(ir, opOR(fr.reg(ix), fr.reg(iy)))
			}
		case token.XOR:
			return func(fr *frame, k *int) {
				// fr.env[instr] = opXOR(fr.get(instr.X), fr.get(instr.Y))
				fr.setReg(ir, opXOR(fr.reg(ix), fr.reg(iy)))
			}
		case token.AND_NOT:
			return func(fr *frame, k *int) {
				// fr.env[instr] = opANDNOT(fr.get(instr.X), fr.get(instr.Y))
				fr.setReg(ir, opANDNOT(fr.reg(ix), fr.reg(iy)))
			}
		case token.SHL:
			if c, ok := instr.Y.(*ssa.Convert); ok {
				xtyp := interp.preToType(c.X.Type())
				xk := xtyp.Kind()
				ic := pfn.regIndex(c.X)
				if xk >= reflect.Int && xk <= reflect.Int64 {
					return func(fr *frame, k *int) {
						// v := reflect.ValueOf(fr.get(c.X))
						v := reflect.ValueOf(fr.reg(ic))
						if v.Int() < 0 {
							panic(runtimeError("negative shift amount"))
						}
						// fr.env[instr] = opSHL(fr.get(instr.X), fr.get(instr.Y))
						fr.setReg(ir, opSHL(fr.reg(ix), fr.reg(iy)))
					}
				}
			}
			return func(fr *frame, k *int) {
				// fr.env[instr] = opSHL(fr.get(instr.X), fr.get(instr.Y))
				fr.setReg(ir, opSHL(fr.reg(ix), fr.reg(iy)))
			}
		case token.SHR:
			if c, ok := instr.Y.(*ssa.Convert); ok {
				xtyp := interp.preToType(c.X.Type())
				xk := xtyp.Kind()
				ic := pfn.regIndex(c.X)
				if xk >= reflect.Int && xk <= reflect.Int64 {
					return func(fr *frame, k *int) {
						// v := reflect.ValueOf(fr.get(c.X))
						v := reflect.ValueOf(fr.reg(ic))
						if v.Int() < 0 {
							panic(runtimeError("negative shift amount"))
						}
						// fr.env[instr] = opSHR(fr.get(instr.X), fr.get(instr.Y))
						fr.setReg(ir, opSHR(fr.reg(ix), fr.reg(iy)))
					}
				}
			}
			return func(fr *frame, k *int) {
				// fr.env[instr] = opSHR(fr.get(instr.X), fr.get(instr.Y))
				fr.setReg(ir, opSHR(fr.reg(ix), fr.reg(iy)))
			}
		case token.LSS:
			return func(fr *frame, k *int) {
				// fr.env[instr] = opLSS(fr.get(instr.X), fr.get(instr.Y))
				fr.setReg(ir, opLSS(fr.reg(ix), fr.reg(iy)))
			}
		case token.LEQ:
			return func(fr *frame, k *int) {
				// fr.env[instr] = opLEQ(fr.get(instr.X), fr.get(instr.Y))
				fr.setReg(ir, opLEQ(fr.reg(ix), fr.reg(iy)))
			}
		case token.EQL:
			return func(fr *frame, k *int) {
				// fr.env[instr] = opEQL(instr, fr.get(instr.X), fr.get(instr.Y))
				fr.setReg(ir, opEQL(instr, fr.reg(ix), fr.reg(iy)))
			}
		case token.NEQ:
			return func(fr *frame, k *int) {
				// fr.env[instr] = !opEQL(instr, fr.get(instr.X), fr.get(instr.Y))
				fr.setReg(ir, !opEQL(instr, fr.reg(ix), fr.reg(iy)))
			}
		case token.GTR:
			return func(fr *frame, k *int) {
				// fr.env[instr] = opGTR(fr.get(instr.X), fr.get(instr.Y))
				fr.setReg(ir, opGTR(fr.reg(ix), fr.reg(iy)))
			}
		case token.GEQ:
			return func(fr *frame, k *int) {
				// fr.env[instr] = opGEQ(fr.get(instr.X), fr.get(instr.Y))
				fr.setReg(ir, opGEQ(fr.reg(ix), fr.reg(iy)))
			}
		default:
			panic(fmt.Errorf("unreachable %v", instr.Op))
		}
	case *ssa.UnOp:
		ir := pfn.regIndex(instr)
		ix := pfn.regIndex(instr.X)
		return func(fr *frame, k *int) {
			// fr.env[instr] = unop(instr, fr.get(instr.X))
			fr.setReg(ir, unop(instr, fr.reg(ix)))
		}
	case *ssa.ChangeInterface:
		ir := pfn.regIndex(instr)
		ix := pfn.regIndex(instr.X)
		return func(fr *frame, k *int) {
			// fr.env[instr] = fr.get(instr.X)
			fr.setReg(ir, fr.reg(ix))
		}
	case *ssa.ChangeType:
		typ := interp.preToType(instr.Type())
		ir := pfn.regIndex(instr)
		switch f := instr.X.(type) {
		case *ssa.Function:
			fn := interp.makeFunc(nil, interp.preToType(f.Type()), f, nil)
			v := fn.Convert(typ).Interface()
			return func(fr *frame, k *int) {
				// fr.env[instr] = v
				fr.setReg(ir, v)
			}
		default:
			ix := pfn.regIndex(instr.X)
			return func(fr *frame, k *int) {
				// x := fr.get(instr.X)
				// if x == nil {
				// 	fr.env[instr] = reflect.New(typ).Elem().Interface()
				// } else {
				// 	fr.env[instr] = reflect.ValueOf(x).Convert(typ).Interface()
				// }
				x := fr.reg(ix)
				if x == nil {
					fr.setReg(ir, reflect.New(typ).Elem().Interface())
				} else {
					fr.setReg(ir, reflect.ValueOf(x).Convert(typ).Interface())
				}
			}
		}
	case *ssa.Convert:
		return makeConvertInstr(pfn, interp, instr)
		// typ := interp.preToType(instr.Type())
		// return func(fr *frame, k *int) {
		// 	x := fr.get(instr.X)
		// 	fr.env[instr] = convert(x, typ)
		// }
	case *ssa.MakeInterface:
		typ := interp.preToType(instr.Type())
		xtyp := interp.preToType(instr.X.Type())
		ir := pfn.regIndex(instr)
		ix := pfn.regIndex(instr.X)
		switch f := instr.X.(type) {
		case *ssa.Function:
			fn := interp.makeFunc(nil, xtyp, f, nil)
			if typ == tyEmptyInterface {
				return func(fr *frame, k *int) {
					// fr.env[instr] = fn.Interface()
					fr.setReg(ir, fn.Interface())
				}
			}
			return func(fr *frame, k *int) {
				v := reflect.New(typ).Elem()
				SetValue(v, fn)
				// fr.env[instr] = v.Interface()
				fr.setReg(ir, v.Interface())
			}
		default:
			if typ == tyEmptyInterface {
				return func(fr *frame, k *int) {
					// x := fr.get(instr.X)
					// fr.env[instr] = x
					fr.setReg(ir, fr.reg(ix))
				}
			}
			return func(fr *frame, k *int) {
				v := reflect.New(typ).Elem()
				// if x := fr.get(instr.X); x != nil {
				if x := fr.reg(ix); x != nil {
					SetValue(v, reflect.ValueOf(x))
				}
				// fr.env[instr] = v.Interface()
				fr.setReg(ir, v.Interface())
			}
		}
	case *ssa.MakeClosure:
		fn := instr.Fn.(*ssa.Function)
		typ := interp.preToType(fn.Type())
		ir := pfn.regIndex(instr)
		ib := make([]int, len(instr.Bindings))
		for i, v := range instr.Bindings {
			ib[i] = pfn.regIndex(v)
		}
		return func(fr *frame, k *int) {
			var bindings []value
			for i, _ := range instr.Bindings {
				// bindings = append(bindings, fr.get(binding))
				bindings = append(bindings, fr.reg(ib[i]))
			}
			c := &closure{fn, bindings}
			// fr.env[instr] = interp.makeFunc(fr, typ, c.Fn, c.Env).Interface()
			fr.setReg(ir, interp.makeFunc(fr, typ, c.Fn, c.Env).Interface())
		}
	case *ssa.MakeChan:
		typ := interp.preToType(instr.Type())
		ir := pfn.regIndex(instr)
		is := pfn.regIndex(instr.Size)
		return func(fr *frame, k *int) {
			// size := fr.get(instr.Size)
			size := fr.reg(is)
			buffer := asInt(size)
			if buffer < 0 {
				panic(runtimeError("makechan: size out of range"))
			}
			// fr.env[instr] = reflect.MakeChan(typ, buffer).Interface()
			fr.setReg(ir, reflect.MakeChan(typ, buffer).Interface())
		}
	case *ssa.MakeMap:
		typ := instr.Type()
		rtyp := interp.preToType(typ)
		key := typ.Underlying().(*types.Map).Key()
		if st, ok := key.Underlying().(*types.Struct); ok && hasUnderscore(st) {
			pfn.mapUnderscoreKey[typ] = true
		}
		ir := pfn.regIndex(instr)
		if instr.Reserve == nil {
			return func(fr *frame, k *int) {
				// fr.env[instr] = reflect.MakeMap(rtyp).Interface()
				fr.setReg(ir, reflect.MakeMap(rtyp).Interface())
			}
		} else {
			iv := pfn.regIndex(instr.Reserve)
			return func(fr *frame, k *int) {
				// reserve := asInt(fr.get(instr.Reserve))
				// fr.env[instr] = reflect.MakeMapWithSize(rtyp, reserve).Interface()
				reserve := asInt(fr.reg(iv))
				fr.setReg(ir, reflect.MakeMapWithSize(rtyp, reserve).Interface())
			}
		}
	case *ssa.MakeSlice:
		typ := interp.preToType(instr.Type())
		ir := pfn.regIndex(instr)
		il := pfn.regIndex(instr.Len)
		ic := pfn.regIndex(instr.Cap)
		return func(fr *frame, k *int) {
			// Len := asInt(fr.get(instr.Len))
			Len := asInt(fr.reg(il))
			if Len < 0 || Len >= maxMemLen {
				panic(runtimeError("makeslice: len out of range"))
			}
			// Cap := asInt(fr.get(instr.Cap))
			Cap := asInt(fr.reg(ic))
			if Cap < 0 || Cap >= maxMemLen {
				panic(runtimeError("makeslice: cap out of range"))
			}
			// fr.env[instr] = reflect.MakeSlice(typ, Len, Cap).Interface()
			fr.setReg(ir, reflect.MakeSlice(typ, Len, Cap).Interface())
		}
	case *ssa.Slice:
		typ := interp.preToType(instr.Type())
		isNamed := typ.Kind() == reflect.Slice && typ != reflect.SliceOf(typ.Elem())
		_, makesliceCheck := instr.X.(*ssa.Alloc)
		ir := pfn.regIndex(instr)
		ix := pfn.regIndex(instr.X)
		ih := pfn.regIndex(instr.High)
		il := pfn.regIndex(instr.Low)
		im := pfn.regIndex(instr.Max)
		if isNamed {
			return func(fr *frame, k *int) {
				// fr.env[instr] = slice(fr, instr, makesliceCheck).Convert(typ).Interface()
				fr.setReg(ir, slice(fr, instr, makesliceCheck, ix, ih, il, im).Convert(typ).Interface())
			}
		} else {
			return func(fr *frame, k *int) {
				// fr.env[instr] = slice(fr, instr, makesliceCheck).Interface()
				fr.setReg(ir, slice(fr, instr, makesliceCheck, ix, ih, il, im).Interface())
			}
		}
	case *ssa.FieldAddr:
		ir := pfn.regIndex(instr)
		ix := pfn.regIndex(instr.X)
		return func(fr *frame, k *int) {
			// v, err := FieldAddr(fr.get(instr.X), instr.Field)
			v, err := FieldAddr(fr.reg(ix), instr.Field)
			if err != nil {
				panic(runtimeError(err.Error()))
			}
			// fr.env[instr] = v
			fr.setReg(ir, v)
		}
	case *ssa.Field:
		ir := pfn.regIndex(instr)
		ix := pfn.regIndex(instr.X)
		return func(fr *frame, k *int) {
			// v, err := Field(fr.get(instr.X), instr.Field)
			v, err := Field(fr.reg(ix), instr.Field)
			if err != nil {
				panic(runtimeError(err.Error()))
			}
			fr.setReg(ir, v)
		}
	case *ssa.IndexAddr:
		ir := pfn.regIndex(instr)
		ix := pfn.regIndex(instr.X)
		ii := pfn.regIndex(instr.Index)
		return func(fr *frame, k *int) {
			// x := fr.get(instr.X)
			// idx := fr.get(instr.Index)
			x := fr.reg(ix)
			idx := fr.reg(ii)
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
			// fr.env[instr] = v.Index(index).Addr().Interface()
			fr.setReg(ir, v.Index(index).Addr().Interface())
		}
	case *ssa.Index:
		ir := pfn.regIndex(instr)
		ix := pfn.regIndex(instr.X)
		ii := pfn.regIndex(instr.Index)
		return func(fr *frame, k *int) {
			// x := fr.get(instr.X)
			// idx := fr.get(instr.Index)
			// v := reflect.ValueOf(x)
			// fr.env[instr] = v.Index(asInt(idx)).Interface()
			x := fr.reg(ix)
			idx := fr.reg(ii)
			v := reflect.ValueOf(x)
			fr.setReg(ir, v.Index(asInt(idx)).Interface())
		}
	case *ssa.Lookup:
		typ := interp.preToType(instr.X.Type())
		ir := pfn.regIndex(instr)
		ix := pfn.regIndex(instr.X)
		ii := pfn.regIndex(instr.Index)
		switch typ.Kind() {
		case reflect.String:
			return func(fr *frame, k *int) {
				// v := fr.get(instr.X)
				// idx := fr.get(instr.Index)
				// fr.env[instr] = reflect.ValueOf(v).String()[asInt(idx)]
				v := fr.reg(ix)
				idx := fr.reg(ii)
				fr.setReg(ir, reflect.ValueOf(v).String()[asInt(idx)])
			}
		case reflect.Map:
			return func(fr *frame, k *int) {
				// m := fr.get(instr.X)
				// idx := fr.get(instr.Index)
				m := fr.reg(ix)
				idx := fr.reg(ii)
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
					// fr.env[instr] = tuple{rv, ok}
					fr.setReg(ir, tuple{rv, ok})
				} else {
					// fr.env[instr] = rv
					fr.setReg(ir, rv)
				}
			}
		default:
			panic("unreachable")
		}
	case *ssa.Select:
		ir := pfn.regIndex(instr)
		ic := make([]int, len(instr.States))
		is := make([]int, len(instr.States))
		for i, state := range instr.States {
			ic[i] = pfn.regIndex(state.Chan)
			if state.Send != nil {
				is[i] = pfn.regIndex(state.Send)
			}
		}
		return func(fr *frame, k *int) {
			var cases []reflect.SelectCase
			if !instr.Blocking {
				cases = append(cases, reflect.SelectCase{
					Dir: reflect.SelectDefault,
				})
			}
			for i, state := range instr.States {
				var dir reflect.SelectDir
				if state.Dir == types.RecvOnly {
					dir = reflect.SelectRecv
				} else {
					dir = reflect.SelectSend
				}
				// ch := reflect.ValueOf(fr.get(state.Chan))
				ch := reflect.ValueOf(fr.reg(ic[i]))
				var send reflect.Value
				if state.Send != nil {
					// v := fr.get(state.Send)
					v := fr.reg(is[i])
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
			// fr.env[instr] = r
			fr.setReg(ir, r)
		}
	case *ssa.SliceToArrayPointer:
		typ := interp.preToType(instr.Type())
		ir := pfn.regIndex(instr)
		ix := pfn.regIndex(instr.X)
		return func(fr *frame, k *int) {
			// x := fr.get(instr.X)
			x := fr.reg(ix)
			v := reflect.ValueOf(x)
			vLen := v.Len()
			tLen := typ.Elem().Len()
			if tLen > vLen {
				panic(runtimeError(fmt.Sprintf("cannot convert slice with length %v to pointer to array with length %v", vLen, tLen)))
			}
			// fr.env[instr] = v.Convert(typ).Interface()
			fr.setReg(ir, v.Convert(typ).Interface())
		}
	case *ssa.Range:
		typ := interp.preToType(instr.X.Type())
		ir := pfn.regIndex(instr)
		ix := pfn.regIndex(instr.X)
		switch typ.Kind() {
		case reflect.String:
			return func(fr *frame, k *int) {
				// v := fr.get(instr.X)
				// fr.env[instr] = &stringIter{Reader: strings.NewReader(reflect.ValueOf(v).String())}
				v := fr.reg(ix)
				fr.setReg(ir, &stringIter{Reader: strings.NewReader(reflect.ValueOf(v).String())})
			}
		case reflect.Map:
			return func(fr *frame, k *int) {
				// v := fr.get(instr.X)
				// fr.env[instr] = &mapIter{iter: reflect.ValueOf(v).MapRange()}
				v := fr.reg(ix)
				fr.setReg(ir, &mapIter{iter: reflect.ValueOf(v).MapRange()})
			}
		default:
			panic("unreachable")
		}
	case *ssa.Next:
		ir := pfn.regIndex(instr)
		ii := pfn.regIndex(instr.Iter)
		if instr.IsString {
			return func(fr *frame, k *int) {
				// fr.env[instr] = fr.get(instr.Iter).(*stringIter).next()
				fr.setReg(ir, fr.reg(ii).(*stringIter).next())
			}
		} else {
			return func(fr *frame, k *int) {
				// fr.env[instr] = fr.get(instr.Iter).(*mapIter).next()
				fr.setReg(ir, fr.reg(ii).(*mapIter).next())
			}
		}
	case *ssa.TypeAssert:
		typ := interp.preToType(instr.AssertedType)
		ir := pfn.regIndex(instr)
		if fn, ok := instr.X.(*ssa.Function); ok {
			f := interp.makeFunc(nil, interp.preToType(fn.Type()), fn, nil).Interface()
			return func(fr *frame, k *int) {
				// fr.env[instr] = typeAssert(interp, instr, typ, f)
				fr.setReg(ir, typeAssert(interp, instr, typ, f))
			}
		} else {
			ix := pfn.regIndex(instr.X)
			return func(fr *frame, k *int) {
				// v := fr.get(instr.X)
				// fr.env[instr] = typeAssert(interp, instr, typ, v)
				v := fr.reg(ix)
				fr.setReg(ir, typeAssert(interp, instr, typ, v))
			}
		}
	case *ssa.Extract:
		if *instr.Referrers() == nil {
			return nil
		}
		ir := pfn.regIndex(instr)
		it := pfn.regIndex(instr.Tuple)
		return func(fr *frame, k *int) {
			// fr.env[instr] = fr.get(instr.Tuple).(tuple)[instr.Index]
			fr.setReg(ir, fr.reg(it).(tuple)[instr.Index])
		}
	// Instructions executed for effect
	case *ssa.Jump:
		return func(fr *frame, k *int) {
			fr.pred, fr.block = fr.block.Index, fr.block.Succs[0]
			*k = kJump
		}
	case *ssa.If:
		ic := pfn.regIndex(instr.Cond)
		switch instr.Cond.Type().(type) {
		case *types.Basic:
			return func(fr *frame, k *int) {
				succ := 1
				// if v := fr.get(instr.Cond); v.(bool) {
				if v := fr.reg(ic); v.(bool) {
					succ = 0
				}
				fr.pred, fr.block = fr.block.Index, fr.block.Succs[succ]
				*k = kJump
			}
		default:
			return func(fr *frame, k *int) {
				succ := 1
				// if v := fr.get(instr.Cond); reflect.ValueOf(v).Bool() {
				if v := fr.reg(ic); reflect.ValueOf(v).Bool() {
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
			ir := pfn.regIndex(instr.Results[0])
			return func(fr *frame, k *int) {
				// fr.result = fr.get(instr.Results[0])
				fr.result = fr.reg(ir)
				fr.block = nil
				*k = kReturn
			}
		default:
			ir := make([]int, n, n)
			for i, v := range instr.Results {
				ir[i] = pfn.regIndex(v)
			}
			return func(fr *frame, k *int) {
				res := make([]value, n, n)
				for i := 0; i < n; i++ {
					// res[i] = fr.get(instr.Results[i])
					res[i] = fr.reg(ir[i])
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
		ix := pfn.regIndex(instr.X)
		return func(fr *frame, k *int) {
			// panic(targetPanic{fr.get(instr.X)})
			panic(targetPanic{fr.reg(ix)})
		}
	case *ssa.Go:
		iv, ia, ib := getCallIndex(pfn, &instr.Call)
		return func(fr *frame, k *int) {
			fn, args := interp.prepareCall(fr, &instr.Call, iv, ia, ib)
			atomic.AddInt32(&interp.goroutines, 1)
			go func() {
				interp.call(nil, instr.Pos(), fn, args, instr.Call.Args)
				atomic.AddInt32(&interp.goroutines, -1)
			}()
		}
	case *ssa.Defer:
		iv, ia, ib := getCallIndex(pfn, &instr.Call)
		return func(fr *frame, k *int) {
			fn, args := interp.prepareCall(fr, &instr.Call, iv, ia, ib)
			fr.defers = &deferred{
				fn:      fn,
				args:    args,
				ssaArgs: instr.Call.Args,
				instr:   instr,
				tail:    fr.defers,
			}
		}
	case *ssa.Send:
		ic := pfn.regIndex(instr.Chan)
		ix := pfn.regIndex(instr.X)
		return func(fr *frame, k *int) {
			// c := fr.get(instr.Chan)
			// x := fr.get(instr.X)
			c := fr.reg(ic)
			x := fr.reg(ix)
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
		ia := pfn.regIndex(instr.Addr)
		if fn, ok := instr.Val.(*ssa.Function); ok {
			v := interp.makeFunc(nil, interp.preToType(fn.Type()), fn, nil)
			return func(fr *frame, k *int) {
				// x := reflect.ValueOf(fr.get(instr.Addr))
				x := reflect.ValueOf(fr.reg(ia))
				SetValue(x.Elem(), v)
			}
		}
		iv := pfn.regIndex(instr.Val)
		return func(fr *frame, k *int) {
			// x := reflect.ValueOf(fr.get(instr.Addr))
			// val := fr.get(instr.Val)
			x := reflect.ValueOf(fr.reg(ia))
			val := fr.reg(iv)
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
		iv := pfn.regIndex(instr.Value)
		im := pfn.regIndex(instr.Map)
		ik := pfn.regIndex(instr.Key)
		if pfn.mapUnderscoreKey[instr.Map.Type()] {
			if fn, ok := instr.Value.(*ssa.Function); ok {
				v := interp.makeFunc(nil, interp.preToType(fn.Type()), fn, nil)
				return func(fr *frame, k *int) {
					// vm := reflect.ValueOf(fr.get(instr.Map))
					// vk := reflect.ValueOf(fr.get(instr.Key))
					vm := reflect.ValueOf(fr.reg(im))
					vk := reflect.ValueOf(fr.reg(ik))
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
				// vm := reflect.ValueOf(fr.get(instr.Map))
				// vk := reflect.ValueOf(fr.get(instr.Key))
				vm := reflect.ValueOf(fr.reg(im))
				vk := reflect.ValueOf(fr.reg(ik))
				// v := fr.get(instr.Value)
				v := fr.reg(iv)
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
					// vm := reflect.ValueOf(fr.get(instr.Map))
					// vk := reflect.ValueOf(fr.get(instr.Key))
					vm := reflect.ValueOf(fr.reg(im))
					vk := reflect.ValueOf(fr.reg(ik))
					vm.SetMapIndex(vk, reflect.ValueOf(v))
				}
			} else {
				return func(fr *frame, k *int) {
					// vm := reflect.ValueOf(fr.get(instr.Map))
					// vk := reflect.ValueOf(fr.get(instr.Key))
					// v := fr.get(instr.Value)
					vm := reflect.ValueOf(fr.reg(im))
					vk := reflect.ValueOf(fr.reg(ik))
					v := fr.reg(iv)
					vm.SetMapIndex(vk, reflect.ValueOf(v))
				}
			}
		}
	case *ssa.DebugRef:
		if interp.fnDebug == nil {
			return nil
		} else {
			ix := pfn.regIndex(instr.X)
			return func(fr *frame, k *int) {
				ref := &DebugInfo{DebugRef: instr, fset: interp.fset}
				ref.toValue = func() (*types.Var, interface{}, bool) {
					if v, ok := instr.Object().(*types.Var); ok {
						// return v, fr.get(instr.X), true
						return v, fr.reg(ix), true
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

func getCallIndex(pfn *Function, call *ssa.CallCommon) (iv int, ia []int, ib []int) {
	iv = pfn.regIndex(call.Value)
	ia = make([]int, len(call.Args), len(call.Args))
	for i, v := range call.Args {
		ia[i] = pfn.regIndex(v)
	}
	if f, ok := call.Value.(*ssa.MakeClosure); ok {
		ib = make([]int, len(f.Bindings), len(f.Bindings))
		for i, binding := range f.Bindings {
			ib[i] = pfn.regIndex(binding)
		}
	}
	return
}

func makeConvertInstr(pfn *Function, interp *Interp, instr *ssa.Convert) func(fr *frame, k *int) {
	typ := interp.preToType(instr.Type())
	xtyp := interp.preToType(instr.X.Type())
	vk := xtyp.Kind()
	ir := pfn.regIndex(instr)
	ix := pfn.regIndex(instr.X)
	switch typ.Kind() {
	case reflect.UnsafePointer:
		if vk == reflect.Uintptr {
			return func(fr *frame, k *int) {
				// v := reflect.ValueOf(fr.get(instr.X))
				// fr.env[instr] = toUnsafePointer(v)
				v := reflect.ValueOf(fr.reg(ix))
				fr.setReg(ir, toUnsafePointer(v))
			}
		} else if vk == reflect.Ptr {
			return func(fr *frame, k *int) {
				// v := reflect.ValueOf(fr.get(instr.X))
				// fr.env[instr] = unsafe.Pointer(v.Pointer())
				v := reflect.ValueOf(fr.reg(ix))
				fr.setReg(ir, unsafe.Pointer(v.Pointer()))
			}
		}
	case reflect.Uintptr:
		if vk == reflect.UnsafePointer {
			return func(fr *frame, k *int) {
				// v := reflect.ValueOf(fr.get(instr.X))
				// fr.env[instr] = v.Pointer()
				v := reflect.ValueOf(fr.reg(ix))
				fr.setReg(ir, v.Pointer())
			}
		}
	case reflect.Ptr:
		if vk == reflect.UnsafePointer {
			return func(fr *frame, k *int) {
				// v := reflect.ValueOf(fr.get(instr.X))
				// fr.env[instr] = reflect.NewAt(typ.Elem(), unsafe.Pointer(v.Pointer())).Interface()
				v := reflect.ValueOf(fr.reg(ix))
				fr.setReg(ir, reflect.NewAt(typ.Elem(), unsafe.Pointer(v.Pointer())).Interface())
			}
		}
	case reflect.Slice:
		if vk == reflect.String {
			elem := typ.Elem()
			switch elem.Kind() {
			case reflect.Uint8:
				if elem.PkgPath() != "" {
					return func(fr *frame, k *int) {
						// v := reflect.ValueOf(fr.get(instr.X))
						// dst := reflect.New(typ).Elem()
						// dst.SetBytes([]byte(v.String()))
						// fr.env[instr] = dst.Interface()
						v := reflect.ValueOf(fr.reg(ix))
						dst := reflect.New(typ).Elem()
						dst.SetBytes([]byte(v.String()))
						fr.setReg(ir, dst.Interface())
					}
				}
			case reflect.Int32:
				if elem.PkgPath() != "" {
					return func(fr *frame, k *int) {
						// v := reflect.ValueOf(fr.get(instr.X))
						// dst := reflect.New(typ).Elem()
						// *(*[]rune)((*reflectValue)(unsafe.Pointer(&dst)).ptr) = []rune(v.String())
						// fr.env[instr] = dst.Interface()
						v := reflect.ValueOf(fr.reg(ix))
						dst := reflect.New(typ).Elem()
						*(*[]rune)((*reflectValue)(unsafe.Pointer(&dst)).ptr) = []rune(v.String())
						fr.setReg(ir, dst.Interface())
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
						// v := reflect.ValueOf(fr.get(instr.X))
						// v = reflect.ValueOf(string(v.Bytes()))
						// fr.env[instr] = v.Convert(typ).Interface()
						v := reflect.ValueOf(fr.reg(ix))
						v = reflect.ValueOf(string(v.Bytes()))
						fr.setReg(ir, v.Convert(typ).Interface())
					}
				}
			case reflect.Int32:
				if elem.PkgPath() != "" {
					return func(fr *frame, k *int) {
						// v := reflect.ValueOf(fr.get(instr.X))
						// v = reflect.ValueOf(*(*[]rune)(((*reflectValue)(unsafe.Pointer(&v))).ptr))
						// fr.env[instr] = v.Convert(typ).Interface()
						v := reflect.ValueOf(fr.reg(ix))
						v = reflect.ValueOf(*(*[]rune)(((*reflectValue)(unsafe.Pointer(&v))).ptr))
						fr.setReg(ir, v.Convert(typ).Interface())
					}
				}
			}
		}
	}
	return func(fr *frame, k *int) {
		// v := reflect.ValueOf(fr.get(instr.X))
		// fr.env[instr] = v.Convert(typ).Interface()
		v := reflect.ValueOf(fr.reg(ix))
		fr.setReg(ir, v.Convert(typ).Interface())
	}
}

func makeCallInstr1(pfn *Function, interp *Interp, instr ssa.Value, call *ssa.CallCommon) func(fr *frame, k *int) {
	ir := pfn.regIndex(instr)
	iv, ia, ib := getCallIndex(pfn, call)
	return func(fr *frame, k *int) {
		fn, args := interp.prepareCall(fr, call, iv, ia, ib)
		// fr.env[instr] = interp.call(fr, call.Pos(), fn, args, call.Args)
		fr.setReg(ir, interp.call(fr, call.Pos(), fn, args, call.Args))
	}
}

func makeCallInstr(pfn *Function, interp *Interp, instr ssa.Value, call *ssa.CallCommon) func(fr *frame, k *int) {
	pos := instr.Pos()
	ir := pfn.regIndex(instr)
	iv, ia, ib := getCallIndex(pfn, call)
	switch fn := call.Value.(type) {
	case *ssa.Builtin:
		nargs := len(call.Args)
		return func(fr *frame, k *int) {
			args := make([]value, nargs, nargs)
			for i := 0; i < nargs; i++ {
				args[i] = fr.reg(ia[i])
			}
			fr.setReg(ir, interp.callBuiltin(fr, pos, fn, args, call.Args))
		}
	case *ssa.MakeClosure:
		pfn := fn.Fn.(*ssa.Function)
		nargs := len(call.Args)
		nenv := len(fn.Bindings)
		return func(fr *frame, k *int) {
			env := make([]value, nenv, nenv)
			for i := 0; i < nenv; i++ {
				env[i] = fr.reg(ib[i])
			}
			args := make([]value, nargs, nargs)
			for i := 0; i < nargs; i++ {
				args[i] = fr.reg(ia[i])
			}
			fr.setReg(ir, interp.callFunction(fr, pos, pfn, args, env))
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
					args[i] = fr.reg(ia[i])
				}
				fr.setReg(ir, interp.callReflect(fr, pos, ext, args, nil))
			}
		}
		return func(fr *frame, k *int) {
			args := make([]value, nargs, nargs)
			for i := 0; i < nargs; i++ {
				args[i] = fr.reg(ia[i])
			}
			fr.setReg(ir, interp.callFunction(fr, pos, fn, args, nil))
		}
	}
	// "dynamic method call" // ("invoke" mode)
	if call.IsInvoke() {
		return makeCallMethodInstr(interp, instr, call, ir, iv, ia)
	}
	// dynamic func call
	typ := interp.preToType(call.Value.Type())
	if typ.Kind() != reflect.Func {
		panic("unsupport")
	}
	nargs := len(call.Args)
	return func(fr *frame, k *int) {
		fn := fr.reg(iv)
		args := make([]value, nargs, nargs)
		for i := 0; i < nargs; i++ {
			args[i] = fr.reg(ia[i])
		}
		switch fn := fn.(type) {
		case *ssa.Function:
			fr.setReg(ir, interp.callFunction(fr, pos, fn, args, nil))
		case *closure:
			fr.setReg(ir, interp.callFunction(fr, pos, fn.Fn, args, fn.Env))
		case *ssa.Builtin:
			fr.setReg(ir, interp.callBuiltin(fr, pos, fn, args, call.Args))
		default:
			fr.setReg(ir, interp.callReflect(fr, pos, reflect.ValueOf(fn), args, nil))
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

func makeCallMethodInstr(interp *Interp, instr ssa.Value, call *ssa.CallCommon, ir int, iv int, ia []int) func(fr *frame, k *int) {
	mname := call.Method.Name()
	nargs := len(call.Args)
	margs := nargs + 1
	pos := instr.Pos()
	var found bool
	var ext reflect.Value
	return func(fr *frame, k *int) {
		v := fr.reg(iv)
		rtype := reflect.TypeOf(v)
		// find user type method *ssa.Function
		if mset, ok := interp.msets[rtype]; ok {
			if fn, ok := mset[mname]; ok {
				args := make([]value, margs, margs)
				args[0] = v
				for i := 0; i < nargs; i++ {
					args[i+1] = fr.reg(ia[i])
				}
				fr.setReg(ir, interp.callFunction(fr, pos, fn, args, nil))
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
			args[i+1] = fr.reg(ia[i])
		}
		fr.setReg(ir, interp.callReflect(fr, pos, ext, args, nil))
	}
}
