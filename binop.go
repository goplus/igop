package gossa

import (
	"reflect"

	"golang.org/x/tools/go/ssa"
)

func makeBinOpADD(pfn *function, instr *ssa.BinOp) func(fr *frame) {
	ir := pfn.regIndex(instr)
	ix, kx, vx := pfn.regIndex3(instr.X)
	iy, ky, vy := pfn.regIndex3(instr.Y)
	typ := pfn.Interp.preToType(instr.Type())
	if typ.PkgPath() == "" {
		switch typ.Kind() {
		case reflect.Int:
			if kx == kindConst {
				x := vx.(int)
				return func(fr *frame) {
					fr.setReg(ir, x+fr.reg(iy).(int))
				}
			} else if ky == kindConst {
				y := vy.(int)
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(int)+y)
				}
			} else {
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(int)+fr.reg(iy).(int))
				}
			}
		case reflect.Int8:
			if kx == kindConst {
				x := vx.(int8)
				return func(fr *frame) {
					fr.setReg(ir, x+fr.reg(iy).(int8))
				}
			} else if ky == kindConst {
				y := vy.(int8)
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(int8)+y)
				}
			} else {
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(int8)+fr.reg(iy).(int8))
				}
			}
		case reflect.Int16:
			if kx == kindConst {
				x := vx.(int16)
				return func(fr *frame) {
					fr.setReg(ir, x+fr.reg(iy).(int16))
				}
			} else if ky == kindConst {
				y := vy.(int16)
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(int16)+y)
				}
			} else {
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(int16)+fr.reg(iy).(int16))
				}
			}
		case reflect.Int32:
			if kx == kindConst {
				x := vx.(int32)
				return func(fr *frame) {
					fr.setReg(ir, x+fr.reg(iy).(int32))
				}
			} else if ky == kindConst {
				y := vy.(int32)
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(int32)+y)
				}
			} else {
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(int32)+fr.reg(iy).(int32))
				}
			}
		case reflect.Int64:
			if kx == kindConst {
				x := vx.(int64)
				return func(fr *frame) {
					fr.setReg(ir, x+fr.reg(iy).(int64))
				}
			} else if ky == kindConst {
				y := vy.(int64)
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(int64)+y)
				}
			} else {
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(int64)+fr.reg(iy).(int64))
				}
			}
		case reflect.Uint:
			if kx == kindConst {
				x := vx.(uint)
				return func(fr *frame) {
					fr.setReg(ir, x+fr.reg(iy).(uint))
				}
			} else if ky == kindConst {
				y := vy.(uint)
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(uint)+y)
				}
			} else {
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(uint)+fr.reg(iy).(uint))
				}
			}
		case reflect.Uint8:
			if kx == kindConst {
				x := vx.(uint8)
				return func(fr *frame) {
					fr.setReg(ir, x+fr.reg(iy).(uint8))
				}
			} else if ky == kindConst {
				y := vy.(uint8)
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(uint8)+y)
				}
			} else {
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(uint8)+fr.reg(iy).(uint8))
				}
			}
		case reflect.Uint16:
			if kx == kindConst {
				x := vx.(uint16)
				return func(fr *frame) {
					fr.setReg(ir, x+fr.reg(iy).(uint16))
				}
			} else if ky == kindConst {
				y := vy.(uint16)
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(uint16)+y)
				}
			} else {
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(uint16)+fr.reg(iy).(uint16))
				}
			}
		case reflect.Uint32:
			if kx == kindConst {
				x := vx.(uint32)
				return func(fr *frame) {
					fr.setReg(ir, x+fr.reg(iy).(uint32))
				}
			} else if ky == kindConst {
				y := vy.(uint32)
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(uint32)+y)
				}
			} else {
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(uint32)+fr.reg(iy).(uint32))
				}
			}
		case reflect.Uint64:
			if kx == kindConst {
				x := vx.(uint64)
				return func(fr *frame) {
					fr.setReg(ir, x+fr.reg(iy).(uint64))
				}
			} else if ky == kindConst {
				y := vy.(uint64)
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(uint64)+y)
				}
			} else {
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(uint64)+fr.reg(iy).(uint64))
				}
			}
		case reflect.Uintptr:
			if kx == kindConst {
				x := vx.(uintptr)
				return func(fr *frame) {
					fr.setReg(ir, x+fr.reg(iy).(uintptr))
				}
			} else if ky == kindConst {
				y := vy.(uintptr)
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(uintptr)+y)
				}
			} else {
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(uintptr)+fr.reg(iy).(uintptr))
				}
			}
		case reflect.Float32:
			if kx == kindConst {
				x := vx.(float32)
				return func(fr *frame) {
					fr.setReg(ir, x+fr.reg(iy).(float32))
				}
			} else if ky == kindConst {
				y := vy.(float32)
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(float32)+y)
				}
			} else {
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(float32)+fr.reg(iy).(float32))
				}
			}
		case reflect.Float64:
			if kx == kindConst {
				x := vx.(float64)
				return func(fr *frame) {
					fr.setReg(ir, x+fr.reg(iy).(float64))
				}
			} else if ky == kindConst {
				y := vy.(float64)
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(float64)+y)
				}
			} else {
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(float64)+fr.reg(iy).(float64))
				}
			}
		case reflect.Complex64:
			if kx == kindConst {
				x := vx.(complex64)
				return func(fr *frame) {
					fr.setReg(ir, x+fr.reg(iy).(complex64))
				}
			} else if ky == kindConst {
				y := vy.(complex64)
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(complex64)+y)
				}
			} else {
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(complex64)+fr.reg(iy).(complex64))
				}
			}
		case reflect.Complex128:
			if kx == kindConst {
				x := vx.(complex128)
				return func(fr *frame) {
					fr.setReg(ir, x+fr.reg(iy).(complex128))
				}
			} else if ky == kindConst {
				y := vy.(complex128)
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(complex128)+y)
				}
			} else {
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(complex128)+fr.reg(iy).(complex128))
				}
			}
		case reflect.String:
			if kx == kindConst {
				x := vx.(string)
				return func(fr *frame) {
					fr.setReg(ir, x+fr.reg(iy).(string))
				}
			} else if ky == kindConst {
				y := vy.(string)
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(string)+y)
				}
			} else {
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(string)+fr.reg(iy).(string))
				}
			}
		}
	} else {
		r := reflect.New(typ).Elem()
		switch typ.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			if kx == kindConst {
				x := reflect.ValueOf(vx).Int()
				return func(fr *frame) {
					y := reflect.ValueOf(fr.reg(iy)).Int()
					r.SetInt(x + y)
					fr.setReg(ir, r.Interface())
				}
			} else if ky == kindConst {
				y := reflect.ValueOf(vy).Int()
				return func(fr *frame) {
					x := reflect.ValueOf(fr.reg(ix)).Int()
					r.SetInt(x + y)
					fr.setReg(ir, r.Interface())
				}
			} else {
				return func(fr *frame) {
					x := reflect.ValueOf(fr.reg(ix)).Int()
					y := reflect.ValueOf(fr.reg(iy)).Int()
					r.SetInt(x + y)
					fr.setReg(ir, r.Interface())
				}
			}
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
			if kx == kindConst {
				x := reflect.ValueOf(vx).Uint()
				return func(fr *frame) {
					y := reflect.ValueOf(fr.reg(iy)).Uint()
					r.SetUint(x + y)
					fr.setReg(ir, r.Interface())
				}
			} else if ky == kindConst {
				y := reflect.ValueOf(vy).Uint()
				return func(fr *frame) {
					x := reflect.ValueOf(fr.reg(ix)).Uint()
					r.SetUint(x + y)
					fr.setReg(ir, r.Interface())
				}
			} else {
				return func(fr *frame) {
					x := reflect.ValueOf(fr.reg(ix)).Uint()
					y := reflect.ValueOf(fr.reg(iy)).Uint()
					r.SetUint(x + y)
					fr.setReg(ir, r.Interface())
				}
			}
		case reflect.Float32, reflect.Float64:
			if kx == kindConst {
				x := reflect.ValueOf(vx).Float()
				return func(fr *frame) {
					y := reflect.ValueOf(fr.reg(iy)).Float()
					r.SetFloat(x + y)
					fr.setReg(ir, r.Interface())
				}
			} else if ky == kindConst {
				y := reflect.ValueOf(vy).Float()
				return func(fr *frame) {
					x := reflect.ValueOf(fr.reg(ix)).Float()
					r.SetFloat(x + y)
					fr.setReg(ir, r.Interface())
				}
			} else {
				return func(fr *frame) {
					x := reflect.ValueOf(fr.reg(ix)).Float()
					y := reflect.ValueOf(fr.reg(iy)).Float()
					r.SetFloat(x + y)
					fr.setReg(ir, r.Interface())
				}
			}
		case reflect.Complex64, reflect.Complex128:
			if kx == kindConst {
				x := reflect.ValueOf(vx).Complex()
				return func(fr *frame) {
					y := reflect.ValueOf(fr.reg(iy)).Complex()
					r.SetComplex(x + y)
					fr.setReg(ir, r.Interface())
				}
			} else if ky == kindConst {
				y := reflect.ValueOf(vy).Complex()
				return func(fr *frame) {
					x := reflect.ValueOf(fr.reg(ix)).Complex()
					r.SetComplex(x + y)
					fr.setReg(ir, r.Interface())
				}
			} else {
				return func(fr *frame) {
					x := reflect.ValueOf(fr.reg(ix)).Complex()
					y := reflect.ValueOf(fr.reg(iy)).Complex()
					r.SetComplex(x + y)
					fr.setReg(ir, r.Interface())
				}
			}
		case reflect.String:
			if kx == kindConst {
				x := reflect.ValueOf(vx).String()
				return func(fr *frame) {
					y := reflect.ValueOf(fr.reg(iy)).String()
					r.SetString(x + y)
					fr.setReg(ir, r.Interface())
				}
			} else if ky == kindConst {
				y := reflect.ValueOf(vy).String()
				return func(fr *frame) {
					x := reflect.ValueOf(fr.reg(ix)).String()
					r.SetString(x + y)
					fr.setReg(ir, r.Interface())
				}
			} else {
				return func(fr *frame) {
					x := reflect.ValueOf(fr.reg(ix)).String()
					y := reflect.ValueOf(fr.reg(iy)).String()
					r.SetString(x + y)
					fr.setReg(ir, r.Interface())
				}
			}
		}
	}
	panic("unreachable")
}

func makeBinOpSUB(pfn *function, instr *ssa.BinOp) func(fr *frame) {
	ir := pfn.regIndex(instr)
	ix, kx, vx := pfn.regIndex3(instr.X)
	iy, ky, vy := pfn.regIndex3(instr.Y)
	typ := pfn.Interp.preToType(instr.Type())
	if typ.PkgPath() == "" {
		switch typ.Kind() {
		case reflect.Int:
			if kx == kindConst {
				x := vx.(int)
				return func(fr *frame) {
					fr.setReg(ir, x-fr.reg(iy).(int))
				}
			} else if ky == kindConst {
				y := vy.(int)
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(int)-y)
				}
			} else {
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(int)-fr.reg(iy).(int))
				}
			}
		case reflect.Int8:
			if kx == kindConst {
				x := vx.(int8)
				return func(fr *frame) {
					fr.setReg(ir, x-fr.reg(iy).(int8))
				}
			} else if ky == kindConst {
				y := vy.(int8)
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(int8)-y)
				}
			} else {
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(int8)-fr.reg(iy).(int8))
				}
			}
		case reflect.Int16:
			if kx == kindConst {
				x := vx.(int16)
				return func(fr *frame) {
					fr.setReg(ir, x-fr.reg(iy).(int16))
				}
			} else if ky == kindConst {
				y := vy.(int16)
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(int16)-y)
				}
			} else {
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(int16)-fr.reg(iy).(int16))
				}
			}
		case reflect.Int32:
			if kx == kindConst {
				x := vx.(int32)
				return func(fr *frame) {
					fr.setReg(ir, x-fr.reg(iy).(int32))
				}
			} else if ky == kindConst {
				y := vy.(int32)
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(int32)-y)
				}
			} else {
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(int32)-fr.reg(iy).(int32))
				}
			}
		case reflect.Int64:
			if kx == kindConst {
				x := vx.(int64)
				return func(fr *frame) {
					fr.setReg(ir, x-fr.reg(iy).(int64))
				}
			} else if ky == kindConst {
				y := vy.(int64)
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(int64)-y)
				}
			} else {
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(int64)-fr.reg(iy).(int64))
				}
			}
		case reflect.Uint:
			if kx == kindConst {
				x := vx.(uint)
				return func(fr *frame) {
					fr.setReg(ir, x-fr.reg(iy).(uint))
				}
			} else if ky == kindConst {
				y := vy.(uint)
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(uint)-y)
				}
			} else {
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(uint)-fr.reg(iy).(uint))
				}
			}
		case reflect.Uint8:
			if kx == kindConst {
				x := vx.(uint8)
				return func(fr *frame) {
					fr.setReg(ir, x-fr.reg(iy).(uint8))
				}
			} else if ky == kindConst {
				y := vy.(uint8)
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(uint8)-y)
				}
			} else {
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(uint8)-fr.reg(iy).(uint8))
				}
			}
		case reflect.Uint16:
			if kx == kindConst {
				x := vx.(uint16)
				return func(fr *frame) {
					fr.setReg(ir, x-fr.reg(iy).(uint16))
				}
			} else if ky == kindConst {
				y := vy.(uint16)
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(uint16)-y)
				}
			} else {
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(uint16)-fr.reg(iy).(uint16))
				}
			}
		case reflect.Uint32:
			if kx == kindConst {
				x := vx.(uint32)
				return func(fr *frame) {
					fr.setReg(ir, x-fr.reg(iy).(uint32))
				}
			} else if ky == kindConst {
				y := vy.(uint32)
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(uint32)-y)
				}
			} else {
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(uint32)-fr.reg(iy).(uint32))
				}
			}
		case reflect.Uint64:
			if kx == kindConst {
				x := vx.(uint64)
				return func(fr *frame) {
					fr.setReg(ir, x-fr.reg(iy).(uint64))
				}
			} else if ky == kindConst {
				y := vy.(uint64)
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(uint64)-y)
				}
			} else {
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(uint64)-fr.reg(iy).(uint64))
				}
			}
		case reflect.Uintptr:
			if kx == kindConst {
				x := vx.(uintptr)
				return func(fr *frame) {
					fr.setReg(ir, x-fr.reg(iy).(uintptr))
				}
			} else if ky == kindConst {
				y := vy.(uintptr)
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(uintptr)-y)
				}
			} else {
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(uintptr)-fr.reg(iy).(uintptr))
				}
			}
		case reflect.Float32:
			if kx == kindConst {
				x := vx.(float32)
				return func(fr *frame) {
					fr.setReg(ir, x-fr.reg(iy).(float32))
				}
			} else if ky == kindConst {
				y := vy.(float32)
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(float32)-y)
				}
			} else {
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(float32)-fr.reg(iy).(float32))
				}
			}
		case reflect.Float64:
			if kx == kindConst {
				x := vx.(float64)
				return func(fr *frame) {
					fr.setReg(ir, x-fr.reg(iy).(float64))
				}
			} else if ky == kindConst {
				y := vy.(float64)
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(float64)-y)
				}
			} else {
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(float64)-fr.reg(iy).(float64))
				}
			}
		case reflect.Complex64:
			if kx == kindConst {
				x := vx.(complex64)
				return func(fr *frame) {
					fr.setReg(ir, x-fr.reg(iy).(complex64))
				}
			} else if ky == kindConst {
				y := vy.(complex64)
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(complex64)-y)
				}
			} else {
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(complex64)-fr.reg(iy).(complex64))
				}
			}
		case reflect.Complex128:
			if kx == kindConst {
				x := vx.(complex128)
				return func(fr *frame) {
					fr.setReg(ir, x-fr.reg(iy).(complex128))
				}
			} else if ky == kindConst {
				y := vy.(complex128)
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(complex128)-y)
				}
			} else {
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(complex128)-fr.reg(iy).(complex128))
				}
			}

		}
	} else {
		r := reflect.New(typ).Elem()
		switch typ.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			if kx == kindConst {
				x := reflect.ValueOf(vx).Int()
				return func(fr *frame) {
					y := reflect.ValueOf(fr.reg(iy)).Int()
					r.SetInt(x - y)
					fr.setReg(ir, r.Interface())
				}
			} else if ky == kindConst {
				y := reflect.ValueOf(vy).Int()
				return func(fr *frame) {
					x := reflect.ValueOf(fr.reg(ix)).Int()
					r.SetInt(x - y)
					fr.setReg(ir, r.Interface())
				}
			} else {
				return func(fr *frame) {
					x := reflect.ValueOf(fr.reg(ix)).Int()
					y := reflect.ValueOf(fr.reg(iy)).Int()
					r.SetInt(x - y)
					fr.setReg(ir, r.Interface())
				}
			}
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
			if kx == kindConst {
				x := reflect.ValueOf(vx).Uint()
				return func(fr *frame) {
					y := reflect.ValueOf(fr.reg(iy)).Uint()
					r.SetUint(x - y)
					fr.setReg(ir, r.Interface())
				}
			} else if ky == kindConst {
				y := reflect.ValueOf(vy).Uint()
				return func(fr *frame) {
					x := reflect.ValueOf(fr.reg(ix)).Uint()
					r.SetUint(x - y)
					fr.setReg(ir, r.Interface())
				}
			} else {
				return func(fr *frame) {
					x := reflect.ValueOf(fr.reg(ix)).Uint()
					y := reflect.ValueOf(fr.reg(iy)).Uint()
					r.SetUint(x - y)
					fr.setReg(ir, r.Interface())
				}
			}
		case reflect.Float32, reflect.Float64:
			if kx == kindConst {
				x := reflect.ValueOf(vx).Float()
				return func(fr *frame) {
					y := reflect.ValueOf(fr.reg(iy)).Float()
					r.SetFloat(x - y)
					fr.setReg(ir, r.Interface())
				}
			} else if ky == kindConst {
				y := reflect.ValueOf(vy).Float()
				return func(fr *frame) {
					x := reflect.ValueOf(fr.reg(ix)).Float()
					r.SetFloat(x - y)
					fr.setReg(ir, r.Interface())
				}
			} else {
				return func(fr *frame) {
					x := reflect.ValueOf(fr.reg(ix)).Float()
					y := reflect.ValueOf(fr.reg(iy)).Float()
					r.SetFloat(x - y)
					fr.setReg(ir, r.Interface())
				}
			}
		case reflect.Complex64, reflect.Complex128:
			if kx == kindConst {
				x := reflect.ValueOf(vx).Complex()
				return func(fr *frame) {
					y := reflect.ValueOf(fr.reg(iy)).Complex()
					r.SetComplex(x - y)
					fr.setReg(ir, r.Interface())
				}
			} else if ky == kindConst {
				y := reflect.ValueOf(vy).Complex()
				return func(fr *frame) {
					x := reflect.ValueOf(fr.reg(ix)).Complex()
					r.SetComplex(x - y)
					fr.setReg(ir, r.Interface())
				}
			} else {
				return func(fr *frame) {
					x := reflect.ValueOf(fr.reg(ix)).Complex()
					y := reflect.ValueOf(fr.reg(iy)).Complex()
					r.SetComplex(x - y)
					fr.setReg(ir, r.Interface())
				}
			}

		}
	}
	panic("unreachable")
}

func makeBinOpMUL(pfn *function, instr *ssa.BinOp) func(fr *frame) {
	ir := pfn.regIndex(instr)
	ix, kx, vx := pfn.regIndex3(instr.X)
	iy, ky, vy := pfn.regIndex3(instr.Y)
	typ := pfn.Interp.preToType(instr.Type())
	if typ.PkgPath() == "" {
		switch typ.Kind() {
		case reflect.Int:
			if kx == kindConst {
				x := vx.(int)
				return func(fr *frame) {
					fr.setReg(ir, x*fr.reg(iy).(int))
				}
			} else if ky == kindConst {
				y := vy.(int)
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(int)*y)
				}
			} else {
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(int)*fr.reg(iy).(int))
				}
			}
		case reflect.Int8:
			if kx == kindConst {
				x := vx.(int8)
				return func(fr *frame) {
					fr.setReg(ir, x*fr.reg(iy).(int8))
				}
			} else if ky == kindConst {
				y := vy.(int8)
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(int8)*y)
				}
			} else {
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(int8)*fr.reg(iy).(int8))
				}
			}
		case reflect.Int16:
			if kx == kindConst {
				x := vx.(int16)
				return func(fr *frame) {
					fr.setReg(ir, x*fr.reg(iy).(int16))
				}
			} else if ky == kindConst {
				y := vy.(int16)
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(int16)*y)
				}
			} else {
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(int16)*fr.reg(iy).(int16))
				}
			}
		case reflect.Int32:
			if kx == kindConst {
				x := vx.(int32)
				return func(fr *frame) {
					fr.setReg(ir, x*fr.reg(iy).(int32))
				}
			} else if ky == kindConst {
				y := vy.(int32)
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(int32)*y)
				}
			} else {
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(int32)*fr.reg(iy).(int32))
				}
			}
		case reflect.Int64:
			if kx == kindConst {
				x := vx.(int64)
				return func(fr *frame) {
					fr.setReg(ir, x*fr.reg(iy).(int64))
				}
			} else if ky == kindConst {
				y := vy.(int64)
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(int64)*y)
				}
			} else {
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(int64)*fr.reg(iy).(int64))
				}
			}
		case reflect.Uint:
			if kx == kindConst {
				x := vx.(uint)
				return func(fr *frame) {
					fr.setReg(ir, x*fr.reg(iy).(uint))
				}
			} else if ky == kindConst {
				y := vy.(uint)
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(uint)*y)
				}
			} else {
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(uint)*fr.reg(iy).(uint))
				}
			}
		case reflect.Uint8:
			if kx == kindConst {
				x := vx.(uint8)
				return func(fr *frame) {
					fr.setReg(ir, x*fr.reg(iy).(uint8))
				}
			} else if ky == kindConst {
				y := vy.(uint8)
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(uint8)*y)
				}
			} else {
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(uint8)*fr.reg(iy).(uint8))
				}
			}
		case reflect.Uint16:
			if kx == kindConst {
				x := vx.(uint16)
				return func(fr *frame) {
					fr.setReg(ir, x*fr.reg(iy).(uint16))
				}
			} else if ky == kindConst {
				y := vy.(uint16)
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(uint16)*y)
				}
			} else {
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(uint16)*fr.reg(iy).(uint16))
				}
			}
		case reflect.Uint32:
			if kx == kindConst {
				x := vx.(uint32)
				return func(fr *frame) {
					fr.setReg(ir, x*fr.reg(iy).(uint32))
				}
			} else if ky == kindConst {
				y := vy.(uint32)
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(uint32)*y)
				}
			} else {
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(uint32)*fr.reg(iy).(uint32))
				}
			}
		case reflect.Uint64:
			if kx == kindConst {
				x := vx.(uint64)
				return func(fr *frame) {
					fr.setReg(ir, x*fr.reg(iy).(uint64))
				}
			} else if ky == kindConst {
				y := vy.(uint64)
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(uint64)*y)
				}
			} else {
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(uint64)*fr.reg(iy).(uint64))
				}
			}
		case reflect.Uintptr:
			if kx == kindConst {
				x := vx.(uintptr)
				return func(fr *frame) {
					fr.setReg(ir, x*fr.reg(iy).(uintptr))
				}
			} else if ky == kindConst {
				y := vy.(uintptr)
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(uintptr)*y)
				}
			} else {
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(uintptr)*fr.reg(iy).(uintptr))
				}
			}
		case reflect.Float32:
			if kx == kindConst {
				x := vx.(float32)
				return func(fr *frame) {
					fr.setReg(ir, x*fr.reg(iy).(float32))
				}
			} else if ky == kindConst {
				y := vy.(float32)
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(float32)*y)
				}
			} else {
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(float32)*fr.reg(iy).(float32))
				}
			}
		case reflect.Float64:
			if kx == kindConst {
				x := vx.(float64)
				return func(fr *frame) {
					fr.setReg(ir, x*fr.reg(iy).(float64))
				}
			} else if ky == kindConst {
				y := vy.(float64)
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(float64)*y)
				}
			} else {
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(float64)*fr.reg(iy).(float64))
				}
			}
		case reflect.Complex64:
			if kx == kindConst {
				x := vx.(complex64)
				return func(fr *frame) {
					fr.setReg(ir, x*fr.reg(iy).(complex64))
				}
			} else if ky == kindConst {
				y := vy.(complex64)
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(complex64)*y)
				}
			} else {
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(complex64)*fr.reg(iy).(complex64))
				}
			}
		case reflect.Complex128:
			if kx == kindConst {
				x := vx.(complex128)
				return func(fr *frame) {
					fr.setReg(ir, x*fr.reg(iy).(complex128))
				}
			} else if ky == kindConst {
				y := vy.(complex128)
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(complex128)*y)
				}
			} else {
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(complex128)*fr.reg(iy).(complex128))
				}
			}

		}
	} else {
		r := reflect.New(typ).Elem()
		switch typ.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			if kx == kindConst {
				x := reflect.ValueOf(vx).Int()
				return func(fr *frame) {
					y := reflect.ValueOf(fr.reg(iy)).Int()
					r.SetInt(x * y)
					fr.setReg(ir, r.Interface())
				}
			} else if ky == kindConst {
				y := reflect.ValueOf(vy).Int()
				return func(fr *frame) {
					x := reflect.ValueOf(fr.reg(ix)).Int()
					r.SetInt(x * y)
					fr.setReg(ir, r.Interface())
				}
			} else {
				return func(fr *frame) {
					x := reflect.ValueOf(fr.reg(ix)).Int()
					y := reflect.ValueOf(fr.reg(iy)).Int()
					r.SetInt(x * y)
					fr.setReg(ir, r.Interface())
				}
			}
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
			if kx == kindConst {
				x := reflect.ValueOf(vx).Uint()
				return func(fr *frame) {
					y := reflect.ValueOf(fr.reg(iy)).Uint()
					r.SetUint(x * y)
					fr.setReg(ir, r.Interface())
				}
			} else if ky == kindConst {
				y := reflect.ValueOf(vy).Uint()
				return func(fr *frame) {
					x := reflect.ValueOf(fr.reg(ix)).Uint()
					r.SetUint(x * y)
					fr.setReg(ir, r.Interface())
				}
			} else {
				return func(fr *frame) {
					x := reflect.ValueOf(fr.reg(ix)).Uint()
					y := reflect.ValueOf(fr.reg(iy)).Uint()
					r.SetUint(x * y)
					fr.setReg(ir, r.Interface())
				}
			}
		case reflect.Float32, reflect.Float64:
			if kx == kindConst {
				x := reflect.ValueOf(vx).Float()
				return func(fr *frame) {
					y := reflect.ValueOf(fr.reg(iy)).Float()
					r.SetFloat(x * y)
					fr.setReg(ir, r.Interface())
				}
			} else if ky == kindConst {
				y := reflect.ValueOf(vy).Float()
				return func(fr *frame) {
					x := reflect.ValueOf(fr.reg(ix)).Float()
					r.SetFloat(x * y)
					fr.setReg(ir, r.Interface())
				}
			} else {
				return func(fr *frame) {
					x := reflect.ValueOf(fr.reg(ix)).Float()
					y := reflect.ValueOf(fr.reg(iy)).Float()
					r.SetFloat(x * y)
					fr.setReg(ir, r.Interface())
				}
			}
		case reflect.Complex64, reflect.Complex128:
			if kx == kindConst {
				x := reflect.ValueOf(vx).Complex()
				return func(fr *frame) {
					y := reflect.ValueOf(fr.reg(iy)).Complex()
					r.SetComplex(x * y)
					fr.setReg(ir, r.Interface())
				}
			} else if ky == kindConst {
				y := reflect.ValueOf(vy).Complex()
				return func(fr *frame) {
					x := reflect.ValueOf(fr.reg(ix)).Complex()
					r.SetComplex(x * y)
					fr.setReg(ir, r.Interface())
				}
			} else {
				return func(fr *frame) {
					x := reflect.ValueOf(fr.reg(ix)).Complex()
					y := reflect.ValueOf(fr.reg(iy)).Complex()
					r.SetComplex(x * y)
					fr.setReg(ir, r.Interface())
				}
			}

		}
	}
	panic("unreachable")
}

func makeBinOpQUO(pfn *function, instr *ssa.BinOp) func(fr *frame) {
	ir := pfn.regIndex(instr)
	ix, kx, vx := pfn.regIndex3(instr.X)
	iy, ky, vy := pfn.regIndex3(instr.Y)
	typ := pfn.Interp.preToType(instr.Type())
	if typ.PkgPath() == "" {
		switch typ.Kind() {
		case reflect.Int:
			if kx == kindConst {
				x := vx.(int)
				return func(fr *frame) {
					fr.setReg(ir, x/fr.reg(iy).(int))
				}
			} else if ky == kindConst {
				y := vy.(int)
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(int)/y)
				}
			} else {
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(int)/fr.reg(iy).(int))
				}
			}
		case reflect.Int8:
			if kx == kindConst {
				x := vx.(int8)
				return func(fr *frame) {
					fr.setReg(ir, x/fr.reg(iy).(int8))
				}
			} else if ky == kindConst {
				y := vy.(int8)
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(int8)/y)
				}
			} else {
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(int8)/fr.reg(iy).(int8))
				}
			}
		case reflect.Int16:
			if kx == kindConst {
				x := vx.(int16)
				return func(fr *frame) {
					fr.setReg(ir, x/fr.reg(iy).(int16))
				}
			} else if ky == kindConst {
				y := vy.(int16)
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(int16)/y)
				}
			} else {
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(int16)/fr.reg(iy).(int16))
				}
			}
		case reflect.Int32:
			if kx == kindConst {
				x := vx.(int32)
				return func(fr *frame) {
					fr.setReg(ir, x/fr.reg(iy).(int32))
				}
			} else if ky == kindConst {
				y := vy.(int32)
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(int32)/y)
				}
			} else {
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(int32)/fr.reg(iy).(int32))
				}
			}
		case reflect.Int64:
			if kx == kindConst {
				x := vx.(int64)
				return func(fr *frame) {
					fr.setReg(ir, x/fr.reg(iy).(int64))
				}
			} else if ky == kindConst {
				y := vy.(int64)
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(int64)/y)
				}
			} else {
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(int64)/fr.reg(iy).(int64))
				}
			}
		case reflect.Uint:
			if kx == kindConst {
				x := vx.(uint)
				return func(fr *frame) {
					fr.setReg(ir, x/fr.reg(iy).(uint))
				}
			} else if ky == kindConst {
				y := vy.(uint)
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(uint)/y)
				}
			} else {
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(uint)/fr.reg(iy).(uint))
				}
			}
		case reflect.Uint8:
			if kx == kindConst {
				x := vx.(uint8)
				return func(fr *frame) {
					fr.setReg(ir, x/fr.reg(iy).(uint8))
				}
			} else if ky == kindConst {
				y := vy.(uint8)
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(uint8)/y)
				}
			} else {
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(uint8)/fr.reg(iy).(uint8))
				}
			}
		case reflect.Uint16:
			if kx == kindConst {
				x := vx.(uint16)
				return func(fr *frame) {
					fr.setReg(ir, x/fr.reg(iy).(uint16))
				}
			} else if ky == kindConst {
				y := vy.(uint16)
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(uint16)/y)
				}
			} else {
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(uint16)/fr.reg(iy).(uint16))
				}
			}
		case reflect.Uint32:
			if kx == kindConst {
				x := vx.(uint32)
				return func(fr *frame) {
					fr.setReg(ir, x/fr.reg(iy).(uint32))
				}
			} else if ky == kindConst {
				y := vy.(uint32)
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(uint32)/y)
				}
			} else {
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(uint32)/fr.reg(iy).(uint32))
				}
			}
		case reflect.Uint64:
			if kx == kindConst {
				x := vx.(uint64)
				return func(fr *frame) {
					fr.setReg(ir, x/fr.reg(iy).(uint64))
				}
			} else if ky == kindConst {
				y := vy.(uint64)
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(uint64)/y)
				}
			} else {
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(uint64)/fr.reg(iy).(uint64))
				}
			}
		case reflect.Uintptr:
			if kx == kindConst {
				x := vx.(uintptr)
				return func(fr *frame) {
					fr.setReg(ir, x/fr.reg(iy).(uintptr))
				}
			} else if ky == kindConst {
				y := vy.(uintptr)
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(uintptr)/y)
				}
			} else {
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(uintptr)/fr.reg(iy).(uintptr))
				}
			}
		case reflect.Float32:
			if kx == kindConst {
				x := vx.(float32)
				return func(fr *frame) {
					fr.setReg(ir, x/fr.reg(iy).(float32))
				}
			} else if ky == kindConst {
				y := vy.(float32)
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(float32)/y)
				}
			} else {
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(float32)/fr.reg(iy).(float32))
				}
			}
		case reflect.Float64:
			if kx == kindConst {
				x := vx.(float64)
				return func(fr *frame) {
					fr.setReg(ir, x/fr.reg(iy).(float64))
				}
			} else if ky == kindConst {
				y := vy.(float64)
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(float64)/y)
				}
			} else {
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(float64)/fr.reg(iy).(float64))
				}
			}
		case reflect.Complex64:
			if kx == kindConst {
				x := vx.(complex64)
				return func(fr *frame) {
					fr.setReg(ir, x/fr.reg(iy).(complex64))
				}
			} else if ky == kindConst {
				y := vy.(complex64)
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(complex64)/y)
				}
			} else {
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(complex64)/fr.reg(iy).(complex64))
				}
			}
		case reflect.Complex128:
			if kx == kindConst {
				x := vx.(complex128)
				return func(fr *frame) {
					fr.setReg(ir, x/fr.reg(iy).(complex128))
				}
			} else if ky == kindConst {
				y := vy.(complex128)
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(complex128)/y)
				}
			} else {
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(complex128)/fr.reg(iy).(complex128))
				}
			}

		}
	} else {
		r := reflect.New(typ).Elem()
		switch typ.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			if kx == kindConst {
				x := reflect.ValueOf(vx).Int()
				return func(fr *frame) {
					y := reflect.ValueOf(fr.reg(iy)).Int()
					r.SetInt(x / y)
					fr.setReg(ir, r.Interface())
				}
			} else if ky == kindConst {
				y := reflect.ValueOf(vy).Int()
				return func(fr *frame) {
					x := reflect.ValueOf(fr.reg(ix)).Int()
					r.SetInt(x / y)
					fr.setReg(ir, r.Interface())
				}
			} else {
				return func(fr *frame) {
					x := reflect.ValueOf(fr.reg(ix)).Int()
					y := reflect.ValueOf(fr.reg(iy)).Int()
					r.SetInt(x / y)
					fr.setReg(ir, r.Interface())
				}
			}
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
			if kx == kindConst {
				x := reflect.ValueOf(vx).Uint()
				return func(fr *frame) {
					y := reflect.ValueOf(fr.reg(iy)).Uint()
					r.SetUint(x / y)
					fr.setReg(ir, r.Interface())
				}
			} else if ky == kindConst {
				y := reflect.ValueOf(vy).Uint()
				return func(fr *frame) {
					x := reflect.ValueOf(fr.reg(ix)).Uint()
					r.SetUint(x / y)
					fr.setReg(ir, r.Interface())
				}
			} else {
				return func(fr *frame) {
					x := reflect.ValueOf(fr.reg(ix)).Uint()
					y := reflect.ValueOf(fr.reg(iy)).Uint()
					r.SetUint(x / y)
					fr.setReg(ir, r.Interface())
				}
			}
		case reflect.Float32, reflect.Float64:
			if kx == kindConst {
				x := reflect.ValueOf(vx).Float()
				return func(fr *frame) {
					y := reflect.ValueOf(fr.reg(iy)).Float()
					r.SetFloat(x / y)
					fr.setReg(ir, r.Interface())
				}
			} else if ky == kindConst {
				y := reflect.ValueOf(vy).Float()
				return func(fr *frame) {
					x := reflect.ValueOf(fr.reg(ix)).Float()
					r.SetFloat(x / y)
					fr.setReg(ir, r.Interface())
				}
			} else {
				return func(fr *frame) {
					x := reflect.ValueOf(fr.reg(ix)).Float()
					y := reflect.ValueOf(fr.reg(iy)).Float()
					r.SetFloat(x / y)
					fr.setReg(ir, r.Interface())
				}
			}
		case reflect.Complex64, reflect.Complex128:
			if kx == kindConst {
				x := reflect.ValueOf(vx).Complex()
				return func(fr *frame) {
					y := reflect.ValueOf(fr.reg(iy)).Complex()
					r.SetComplex(x / y)
					fr.setReg(ir, r.Interface())
				}
			} else if ky == kindConst {
				y := reflect.ValueOf(vy).Complex()
				return func(fr *frame) {
					x := reflect.ValueOf(fr.reg(ix)).Complex()
					r.SetComplex(x / y)
					fr.setReg(ir, r.Interface())
				}
			} else {
				return func(fr *frame) {
					x := reflect.ValueOf(fr.reg(ix)).Complex()
					y := reflect.ValueOf(fr.reg(iy)).Complex()
					r.SetComplex(x / y)
					fr.setReg(ir, r.Interface())
				}
			}

		}
	}
	panic("unreachable")
}

func makeBinOpREM(pfn *function, instr *ssa.BinOp) func(fr *frame) {
	ir := pfn.regIndex(instr)
	ix, kx, vx := pfn.regIndex3(instr.X)
	iy, ky, vy := pfn.regIndex3(instr.Y)
	typ := pfn.Interp.preToType(instr.Type())
	if typ.PkgPath() == "" {
		switch typ.Kind() {
		case reflect.Int:
			if kx == kindConst {
				x := vx.(int)
				return func(fr *frame) {
					fr.setReg(ir, x%fr.reg(iy).(int))
				}
			} else if ky == kindConst {
				y := vy.(int)
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(int)%y)
				}
			} else {
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(int)%fr.reg(iy).(int))
				}
			}
		case reflect.Int8:
			if kx == kindConst {
				x := vx.(int8)
				return func(fr *frame) {
					fr.setReg(ir, x%fr.reg(iy).(int8))
				}
			} else if ky == kindConst {
				y := vy.(int8)
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(int8)%y)
				}
			} else {
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(int8)%fr.reg(iy).(int8))
				}
			}
		case reflect.Int16:
			if kx == kindConst {
				x := vx.(int16)
				return func(fr *frame) {
					fr.setReg(ir, x%fr.reg(iy).(int16))
				}
			} else if ky == kindConst {
				y := vy.(int16)
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(int16)%y)
				}
			} else {
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(int16)%fr.reg(iy).(int16))
				}
			}
		case reflect.Int32:
			if kx == kindConst {
				x := vx.(int32)
				return func(fr *frame) {
					fr.setReg(ir, x%fr.reg(iy).(int32))
				}
			} else if ky == kindConst {
				y := vy.(int32)
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(int32)%y)
				}
			} else {
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(int32)%fr.reg(iy).(int32))
				}
			}
		case reflect.Int64:
			if kx == kindConst {
				x := vx.(int64)
				return func(fr *frame) {
					fr.setReg(ir, x%fr.reg(iy).(int64))
				}
			} else if ky == kindConst {
				y := vy.(int64)
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(int64)%y)
				}
			} else {
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(int64)%fr.reg(iy).(int64))
				}
			}
		case reflect.Uint:
			if kx == kindConst {
				x := vx.(uint)
				return func(fr *frame) {
					fr.setReg(ir, x%fr.reg(iy).(uint))
				}
			} else if ky == kindConst {
				y := vy.(uint)
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(uint)%y)
				}
			} else {
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(uint)%fr.reg(iy).(uint))
				}
			}
		case reflect.Uint8:
			if kx == kindConst {
				x := vx.(uint8)
				return func(fr *frame) {
					fr.setReg(ir, x%fr.reg(iy).(uint8))
				}
			} else if ky == kindConst {
				y := vy.(uint8)
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(uint8)%y)
				}
			} else {
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(uint8)%fr.reg(iy).(uint8))
				}
			}
		case reflect.Uint16:
			if kx == kindConst {
				x := vx.(uint16)
				return func(fr *frame) {
					fr.setReg(ir, x%fr.reg(iy).(uint16))
				}
			} else if ky == kindConst {
				y := vy.(uint16)
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(uint16)%y)
				}
			} else {
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(uint16)%fr.reg(iy).(uint16))
				}
			}
		case reflect.Uint32:
			if kx == kindConst {
				x := vx.(uint32)
				return func(fr *frame) {
					fr.setReg(ir, x%fr.reg(iy).(uint32))
				}
			} else if ky == kindConst {
				y := vy.(uint32)
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(uint32)%y)
				}
			} else {
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(uint32)%fr.reg(iy).(uint32))
				}
			}
		case reflect.Uint64:
			if kx == kindConst {
				x := vx.(uint64)
				return func(fr *frame) {
					fr.setReg(ir, x%fr.reg(iy).(uint64))
				}
			} else if ky == kindConst {
				y := vy.(uint64)
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(uint64)%y)
				}
			} else {
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(uint64)%fr.reg(iy).(uint64))
				}
			}
		case reflect.Uintptr:
			if kx == kindConst {
				x := vx.(uintptr)
				return func(fr *frame) {
					fr.setReg(ir, x%fr.reg(iy).(uintptr))
				}
			} else if ky == kindConst {
				y := vy.(uintptr)
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(uintptr)%y)
				}
			} else {
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(uintptr)%fr.reg(iy).(uintptr))
				}
			}

		}
	} else {
		r := reflect.New(typ).Elem()
		switch typ.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			if kx == kindConst {
				x := reflect.ValueOf(vx).Int()
				return func(fr *frame) {
					y := reflect.ValueOf(fr.reg(iy)).Int()
					r.SetInt(x % y)
					fr.setReg(ir, r.Interface())
				}
			} else if ky == kindConst {
				y := reflect.ValueOf(vy).Int()
				return func(fr *frame) {
					x := reflect.ValueOf(fr.reg(ix)).Int()
					r.SetInt(x % y)
					fr.setReg(ir, r.Interface())
				}
			} else {
				return func(fr *frame) {
					x := reflect.ValueOf(fr.reg(ix)).Int()
					y := reflect.ValueOf(fr.reg(iy)).Int()
					r.SetInt(x % y)
					fr.setReg(ir, r.Interface())
				}
			}
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
			if kx == kindConst {
				x := reflect.ValueOf(vx).Uint()
				return func(fr *frame) {
					y := reflect.ValueOf(fr.reg(iy)).Uint()
					r.SetUint(x % y)
					fr.setReg(ir, r.Interface())
				}
			} else if ky == kindConst {
				y := reflect.ValueOf(vy).Uint()
				return func(fr *frame) {
					x := reflect.ValueOf(fr.reg(ix)).Uint()
					r.SetUint(x % y)
					fr.setReg(ir, r.Interface())
				}
			} else {
				return func(fr *frame) {
					x := reflect.ValueOf(fr.reg(ix)).Uint()
					y := reflect.ValueOf(fr.reg(iy)).Uint()
					r.SetUint(x % y)
					fr.setReg(ir, r.Interface())
				}
			}

		}
	}
	panic("unreachable")
}

func makeBinOpAND(pfn *function, instr *ssa.BinOp) func(fr *frame) {
	ir := pfn.regIndex(instr)
	ix, kx, vx := pfn.regIndex3(instr.X)
	iy, ky, vy := pfn.regIndex3(instr.Y)
	typ := pfn.Interp.preToType(instr.Type())
	if typ.PkgPath() == "" {
		switch typ.Kind() {
		case reflect.Int:
			if kx == kindConst {
				x := vx.(int)
				return func(fr *frame) {
					fr.setReg(ir, x&fr.reg(iy).(int))
				}
			} else if ky == kindConst {
				y := vy.(int)
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(int)&y)
				}
			} else {
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(int)&fr.reg(iy).(int))
				}
			}
		case reflect.Int8:
			if kx == kindConst {
				x := vx.(int8)
				return func(fr *frame) {
					fr.setReg(ir, x&fr.reg(iy).(int8))
				}
			} else if ky == kindConst {
				y := vy.(int8)
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(int8)&y)
				}
			} else {
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(int8)&fr.reg(iy).(int8))
				}
			}
		case reflect.Int16:
			if kx == kindConst {
				x := vx.(int16)
				return func(fr *frame) {
					fr.setReg(ir, x&fr.reg(iy).(int16))
				}
			} else if ky == kindConst {
				y := vy.(int16)
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(int16)&y)
				}
			} else {
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(int16)&fr.reg(iy).(int16))
				}
			}
		case reflect.Int32:
			if kx == kindConst {
				x := vx.(int32)
				return func(fr *frame) {
					fr.setReg(ir, x&fr.reg(iy).(int32))
				}
			} else if ky == kindConst {
				y := vy.(int32)
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(int32)&y)
				}
			} else {
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(int32)&fr.reg(iy).(int32))
				}
			}
		case reflect.Int64:
			if kx == kindConst {
				x := vx.(int64)
				return func(fr *frame) {
					fr.setReg(ir, x&fr.reg(iy).(int64))
				}
			} else if ky == kindConst {
				y := vy.(int64)
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(int64)&y)
				}
			} else {
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(int64)&fr.reg(iy).(int64))
				}
			}
		case reflect.Uint:
			if kx == kindConst {
				x := vx.(uint)
				return func(fr *frame) {
					fr.setReg(ir, x&fr.reg(iy).(uint))
				}
			} else if ky == kindConst {
				y := vy.(uint)
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(uint)&y)
				}
			} else {
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(uint)&fr.reg(iy).(uint))
				}
			}
		case reflect.Uint8:
			if kx == kindConst {
				x := vx.(uint8)
				return func(fr *frame) {
					fr.setReg(ir, x&fr.reg(iy).(uint8))
				}
			} else if ky == kindConst {
				y := vy.(uint8)
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(uint8)&y)
				}
			} else {
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(uint8)&fr.reg(iy).(uint8))
				}
			}
		case reflect.Uint16:
			if kx == kindConst {
				x := vx.(uint16)
				return func(fr *frame) {
					fr.setReg(ir, x&fr.reg(iy).(uint16))
				}
			} else if ky == kindConst {
				y := vy.(uint16)
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(uint16)&y)
				}
			} else {
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(uint16)&fr.reg(iy).(uint16))
				}
			}
		case reflect.Uint32:
			if kx == kindConst {
				x := vx.(uint32)
				return func(fr *frame) {
					fr.setReg(ir, x&fr.reg(iy).(uint32))
				}
			} else if ky == kindConst {
				y := vy.(uint32)
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(uint32)&y)
				}
			} else {
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(uint32)&fr.reg(iy).(uint32))
				}
			}
		case reflect.Uint64:
			if kx == kindConst {
				x := vx.(uint64)
				return func(fr *frame) {
					fr.setReg(ir, x&fr.reg(iy).(uint64))
				}
			} else if ky == kindConst {
				y := vy.(uint64)
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(uint64)&y)
				}
			} else {
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(uint64)&fr.reg(iy).(uint64))
				}
			}
		case reflect.Uintptr:
			if kx == kindConst {
				x := vx.(uintptr)
				return func(fr *frame) {
					fr.setReg(ir, x&fr.reg(iy).(uintptr))
				}
			} else if ky == kindConst {
				y := vy.(uintptr)
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(uintptr)&y)
				}
			} else {
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(uintptr)&fr.reg(iy).(uintptr))
				}
			}

		}
	} else {
		r := reflect.New(typ).Elem()
		switch typ.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			if kx == kindConst {
				x := reflect.ValueOf(vx).Int()
				return func(fr *frame) {
					y := reflect.ValueOf(fr.reg(iy)).Int()
					r.SetInt(x & y)
					fr.setReg(ir, r.Interface())
				}
			} else if ky == kindConst {
				y := reflect.ValueOf(vy).Int()
				return func(fr *frame) {
					x := reflect.ValueOf(fr.reg(ix)).Int()
					r.SetInt(x & y)
					fr.setReg(ir, r.Interface())
				}
			} else {
				return func(fr *frame) {
					x := reflect.ValueOf(fr.reg(ix)).Int()
					y := reflect.ValueOf(fr.reg(iy)).Int()
					r.SetInt(x & y)
					fr.setReg(ir, r.Interface())
				}
			}
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
			if kx == kindConst {
				x := reflect.ValueOf(vx).Uint()
				return func(fr *frame) {
					y := reflect.ValueOf(fr.reg(iy)).Uint()
					r.SetUint(x & y)
					fr.setReg(ir, r.Interface())
				}
			} else if ky == kindConst {
				y := reflect.ValueOf(vy).Uint()
				return func(fr *frame) {
					x := reflect.ValueOf(fr.reg(ix)).Uint()
					r.SetUint(x & y)
					fr.setReg(ir, r.Interface())
				}
			} else {
				return func(fr *frame) {
					x := reflect.ValueOf(fr.reg(ix)).Uint()
					y := reflect.ValueOf(fr.reg(iy)).Uint()
					r.SetUint(x & y)
					fr.setReg(ir, r.Interface())
				}
			}

		}
	}
	panic("unreachable")
}

func makeBinOpOR(pfn *function, instr *ssa.BinOp) func(fr *frame) {
	ir := pfn.regIndex(instr)
	ix, kx, vx := pfn.regIndex3(instr.X)
	iy, ky, vy := pfn.regIndex3(instr.Y)
	typ := pfn.Interp.preToType(instr.Type())
	if typ.PkgPath() == "" {
		switch typ.Kind() {
		case reflect.Int:
			if kx == kindConst {
				x := vx.(int)
				return func(fr *frame) {
					fr.setReg(ir, x|fr.reg(iy).(int))
				}
			} else if ky == kindConst {
				y := vy.(int)
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(int)|y)
				}
			} else {
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(int)|fr.reg(iy).(int))
				}
			}
		case reflect.Int8:
			if kx == kindConst {
				x := vx.(int8)
				return func(fr *frame) {
					fr.setReg(ir, x|fr.reg(iy).(int8))
				}
			} else if ky == kindConst {
				y := vy.(int8)
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(int8)|y)
				}
			} else {
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(int8)|fr.reg(iy).(int8))
				}
			}
		case reflect.Int16:
			if kx == kindConst {
				x := vx.(int16)
				return func(fr *frame) {
					fr.setReg(ir, x|fr.reg(iy).(int16))
				}
			} else if ky == kindConst {
				y := vy.(int16)
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(int16)|y)
				}
			} else {
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(int16)|fr.reg(iy).(int16))
				}
			}
		case reflect.Int32:
			if kx == kindConst {
				x := vx.(int32)
				return func(fr *frame) {
					fr.setReg(ir, x|fr.reg(iy).(int32))
				}
			} else if ky == kindConst {
				y := vy.(int32)
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(int32)|y)
				}
			} else {
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(int32)|fr.reg(iy).(int32))
				}
			}
		case reflect.Int64:
			if kx == kindConst {
				x := vx.(int64)
				return func(fr *frame) {
					fr.setReg(ir, x|fr.reg(iy).(int64))
				}
			} else if ky == kindConst {
				y := vy.(int64)
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(int64)|y)
				}
			} else {
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(int64)|fr.reg(iy).(int64))
				}
			}
		case reflect.Uint:
			if kx == kindConst {
				x := vx.(uint)
				return func(fr *frame) {
					fr.setReg(ir, x|fr.reg(iy).(uint))
				}
			} else if ky == kindConst {
				y := vy.(uint)
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(uint)|y)
				}
			} else {
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(uint)|fr.reg(iy).(uint))
				}
			}
		case reflect.Uint8:
			if kx == kindConst {
				x := vx.(uint8)
				return func(fr *frame) {
					fr.setReg(ir, x|fr.reg(iy).(uint8))
				}
			} else if ky == kindConst {
				y := vy.(uint8)
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(uint8)|y)
				}
			} else {
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(uint8)|fr.reg(iy).(uint8))
				}
			}
		case reflect.Uint16:
			if kx == kindConst {
				x := vx.(uint16)
				return func(fr *frame) {
					fr.setReg(ir, x|fr.reg(iy).(uint16))
				}
			} else if ky == kindConst {
				y := vy.(uint16)
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(uint16)|y)
				}
			} else {
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(uint16)|fr.reg(iy).(uint16))
				}
			}
		case reflect.Uint32:
			if kx == kindConst {
				x := vx.(uint32)
				return func(fr *frame) {
					fr.setReg(ir, x|fr.reg(iy).(uint32))
				}
			} else if ky == kindConst {
				y := vy.(uint32)
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(uint32)|y)
				}
			} else {
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(uint32)|fr.reg(iy).(uint32))
				}
			}
		case reflect.Uint64:
			if kx == kindConst {
				x := vx.(uint64)
				return func(fr *frame) {
					fr.setReg(ir, x|fr.reg(iy).(uint64))
				}
			} else if ky == kindConst {
				y := vy.(uint64)
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(uint64)|y)
				}
			} else {
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(uint64)|fr.reg(iy).(uint64))
				}
			}
		case reflect.Uintptr:
			if kx == kindConst {
				x := vx.(uintptr)
				return func(fr *frame) {
					fr.setReg(ir, x|fr.reg(iy).(uintptr))
				}
			} else if ky == kindConst {
				y := vy.(uintptr)
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(uintptr)|y)
				}
			} else {
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(uintptr)|fr.reg(iy).(uintptr))
				}
			}

		}
	} else {
		r := reflect.New(typ).Elem()
		switch typ.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			if kx == kindConst {
				x := reflect.ValueOf(vx).Int()
				return func(fr *frame) {
					y := reflect.ValueOf(fr.reg(iy)).Int()
					r.SetInt(x | y)
					fr.setReg(ir, r.Interface())
				}
			} else if ky == kindConst {
				y := reflect.ValueOf(vy).Int()
				return func(fr *frame) {
					x := reflect.ValueOf(fr.reg(ix)).Int()
					r.SetInt(x | y)
					fr.setReg(ir, r.Interface())
				}
			} else {
				return func(fr *frame) {
					x := reflect.ValueOf(fr.reg(ix)).Int()
					y := reflect.ValueOf(fr.reg(iy)).Int()
					r.SetInt(x | y)
					fr.setReg(ir, r.Interface())
				}
			}
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
			if kx == kindConst {
				x := reflect.ValueOf(vx).Uint()
				return func(fr *frame) {
					y := reflect.ValueOf(fr.reg(iy)).Uint()
					r.SetUint(x | y)
					fr.setReg(ir, r.Interface())
				}
			} else if ky == kindConst {
				y := reflect.ValueOf(vy).Uint()
				return func(fr *frame) {
					x := reflect.ValueOf(fr.reg(ix)).Uint()
					r.SetUint(x | y)
					fr.setReg(ir, r.Interface())
				}
			} else {
				return func(fr *frame) {
					x := reflect.ValueOf(fr.reg(ix)).Uint()
					y := reflect.ValueOf(fr.reg(iy)).Uint()
					r.SetUint(x | y)
					fr.setReg(ir, r.Interface())
				}
			}

		}
	}
	panic("unreachable")
}

func makeBinOpXOR(pfn *function, instr *ssa.BinOp) func(fr *frame) {
	ir := pfn.regIndex(instr)
	ix, kx, vx := pfn.regIndex3(instr.X)
	iy, ky, vy := pfn.regIndex3(instr.Y)
	typ := pfn.Interp.preToType(instr.Type())
	if typ.PkgPath() == "" {
		switch typ.Kind() {
		case reflect.Int:
			if kx == kindConst {
				x := vx.(int)
				return func(fr *frame) {
					fr.setReg(ir, x^fr.reg(iy).(int))
				}
			} else if ky == kindConst {
				y := vy.(int)
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(int)^y)
				}
			} else {
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(int)^fr.reg(iy).(int))
				}
			}
		case reflect.Int8:
			if kx == kindConst {
				x := vx.(int8)
				return func(fr *frame) {
					fr.setReg(ir, x^fr.reg(iy).(int8))
				}
			} else if ky == kindConst {
				y := vy.(int8)
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(int8)^y)
				}
			} else {
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(int8)^fr.reg(iy).(int8))
				}
			}
		case reflect.Int16:
			if kx == kindConst {
				x := vx.(int16)
				return func(fr *frame) {
					fr.setReg(ir, x^fr.reg(iy).(int16))
				}
			} else if ky == kindConst {
				y := vy.(int16)
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(int16)^y)
				}
			} else {
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(int16)^fr.reg(iy).(int16))
				}
			}
		case reflect.Int32:
			if kx == kindConst {
				x := vx.(int32)
				return func(fr *frame) {
					fr.setReg(ir, x^fr.reg(iy).(int32))
				}
			} else if ky == kindConst {
				y := vy.(int32)
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(int32)^y)
				}
			} else {
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(int32)^fr.reg(iy).(int32))
				}
			}
		case reflect.Int64:
			if kx == kindConst {
				x := vx.(int64)
				return func(fr *frame) {
					fr.setReg(ir, x^fr.reg(iy).(int64))
				}
			} else if ky == kindConst {
				y := vy.(int64)
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(int64)^y)
				}
			} else {
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(int64)^fr.reg(iy).(int64))
				}
			}
		case reflect.Uint:
			if kx == kindConst {
				x := vx.(uint)
				return func(fr *frame) {
					fr.setReg(ir, x^fr.reg(iy).(uint))
				}
			} else if ky == kindConst {
				y := vy.(uint)
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(uint)^y)
				}
			} else {
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(uint)^fr.reg(iy).(uint))
				}
			}
		case reflect.Uint8:
			if kx == kindConst {
				x := vx.(uint8)
				return func(fr *frame) {
					fr.setReg(ir, x^fr.reg(iy).(uint8))
				}
			} else if ky == kindConst {
				y := vy.(uint8)
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(uint8)^y)
				}
			} else {
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(uint8)^fr.reg(iy).(uint8))
				}
			}
		case reflect.Uint16:
			if kx == kindConst {
				x := vx.(uint16)
				return func(fr *frame) {
					fr.setReg(ir, x^fr.reg(iy).(uint16))
				}
			} else if ky == kindConst {
				y := vy.(uint16)
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(uint16)^y)
				}
			} else {
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(uint16)^fr.reg(iy).(uint16))
				}
			}
		case reflect.Uint32:
			if kx == kindConst {
				x := vx.(uint32)
				return func(fr *frame) {
					fr.setReg(ir, x^fr.reg(iy).(uint32))
				}
			} else if ky == kindConst {
				y := vy.(uint32)
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(uint32)^y)
				}
			} else {
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(uint32)^fr.reg(iy).(uint32))
				}
			}
		case reflect.Uint64:
			if kx == kindConst {
				x := vx.(uint64)
				return func(fr *frame) {
					fr.setReg(ir, x^fr.reg(iy).(uint64))
				}
			} else if ky == kindConst {
				y := vy.(uint64)
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(uint64)^y)
				}
			} else {
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(uint64)^fr.reg(iy).(uint64))
				}
			}
		case reflect.Uintptr:
			if kx == kindConst {
				x := vx.(uintptr)
				return func(fr *frame) {
					fr.setReg(ir, x^fr.reg(iy).(uintptr))
				}
			} else if ky == kindConst {
				y := vy.(uintptr)
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(uintptr)^y)
				}
			} else {
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(uintptr)^fr.reg(iy).(uintptr))
				}
			}

		}
	} else {
		r := reflect.New(typ).Elem()
		switch typ.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			if kx == kindConst {
				x := reflect.ValueOf(vx).Int()
				return func(fr *frame) {
					y := reflect.ValueOf(fr.reg(iy)).Int()
					r.SetInt(x ^ y)
					fr.setReg(ir, r.Interface())
				}
			} else if ky == kindConst {
				y := reflect.ValueOf(vy).Int()
				return func(fr *frame) {
					x := reflect.ValueOf(fr.reg(ix)).Int()
					r.SetInt(x ^ y)
					fr.setReg(ir, r.Interface())
				}
			} else {
				return func(fr *frame) {
					x := reflect.ValueOf(fr.reg(ix)).Int()
					y := reflect.ValueOf(fr.reg(iy)).Int()
					r.SetInt(x ^ y)
					fr.setReg(ir, r.Interface())
				}
			}
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
			if kx == kindConst {
				x := reflect.ValueOf(vx).Uint()
				return func(fr *frame) {
					y := reflect.ValueOf(fr.reg(iy)).Uint()
					r.SetUint(x ^ y)
					fr.setReg(ir, r.Interface())
				}
			} else if ky == kindConst {
				y := reflect.ValueOf(vy).Uint()
				return func(fr *frame) {
					x := reflect.ValueOf(fr.reg(ix)).Uint()
					r.SetUint(x ^ y)
					fr.setReg(ir, r.Interface())
				}
			} else {
				return func(fr *frame) {
					x := reflect.ValueOf(fr.reg(ix)).Uint()
					y := reflect.ValueOf(fr.reg(iy)).Uint()
					r.SetUint(x ^ y)
					fr.setReg(ir, r.Interface())
				}
			}

		}
	}
	panic("unreachable")
}

func makeBinOpANDNOT(pfn *function, instr *ssa.BinOp) func(fr *frame) {
	ir := pfn.regIndex(instr)
	ix, kx, vx := pfn.regIndex3(instr.X)
	iy, ky, vy := pfn.regIndex3(instr.Y)
	typ := pfn.Interp.preToType(instr.Type())
	if typ.PkgPath() == "" {
		switch typ.Kind() {
		case reflect.Int:
			if kx == kindConst {
				x := vx.(int)
				return func(fr *frame) {
					fr.setReg(ir, x&^fr.reg(iy).(int))
				}
			} else if ky == kindConst {
				y := vy.(int)
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(int)&^y)
				}
			} else {
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(int)&^fr.reg(iy).(int))
				}
			}
		case reflect.Int8:
			if kx == kindConst {
				x := vx.(int8)
				return func(fr *frame) {
					fr.setReg(ir, x&^fr.reg(iy).(int8))
				}
			} else if ky == kindConst {
				y := vy.(int8)
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(int8)&^y)
				}
			} else {
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(int8)&^fr.reg(iy).(int8))
				}
			}
		case reflect.Int16:
			if kx == kindConst {
				x := vx.(int16)
				return func(fr *frame) {
					fr.setReg(ir, x&^fr.reg(iy).(int16))
				}
			} else if ky == kindConst {
				y := vy.(int16)
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(int16)&^y)
				}
			} else {
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(int16)&^fr.reg(iy).(int16))
				}
			}
		case reflect.Int32:
			if kx == kindConst {
				x := vx.(int32)
				return func(fr *frame) {
					fr.setReg(ir, x&^fr.reg(iy).(int32))
				}
			} else if ky == kindConst {
				y := vy.(int32)
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(int32)&^y)
				}
			} else {
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(int32)&^fr.reg(iy).(int32))
				}
			}
		case reflect.Int64:
			if kx == kindConst {
				x := vx.(int64)
				return func(fr *frame) {
					fr.setReg(ir, x&^fr.reg(iy).(int64))
				}
			} else if ky == kindConst {
				y := vy.(int64)
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(int64)&^y)
				}
			} else {
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(int64)&^fr.reg(iy).(int64))
				}
			}
		case reflect.Uint:
			if kx == kindConst {
				x := vx.(uint)
				return func(fr *frame) {
					fr.setReg(ir, x&^fr.reg(iy).(uint))
				}
			} else if ky == kindConst {
				y := vy.(uint)
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(uint)&^y)
				}
			} else {
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(uint)&^fr.reg(iy).(uint))
				}
			}
		case reflect.Uint8:
			if kx == kindConst {
				x := vx.(uint8)
				return func(fr *frame) {
					fr.setReg(ir, x&^fr.reg(iy).(uint8))
				}
			} else if ky == kindConst {
				y := vy.(uint8)
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(uint8)&^y)
				}
			} else {
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(uint8)&^fr.reg(iy).(uint8))
				}
			}
		case reflect.Uint16:
			if kx == kindConst {
				x := vx.(uint16)
				return func(fr *frame) {
					fr.setReg(ir, x&^fr.reg(iy).(uint16))
				}
			} else if ky == kindConst {
				y := vy.(uint16)
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(uint16)&^y)
				}
			} else {
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(uint16)&^fr.reg(iy).(uint16))
				}
			}
		case reflect.Uint32:
			if kx == kindConst {
				x := vx.(uint32)
				return func(fr *frame) {
					fr.setReg(ir, x&^fr.reg(iy).(uint32))
				}
			} else if ky == kindConst {
				y := vy.(uint32)
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(uint32)&^y)
				}
			} else {
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(uint32)&^fr.reg(iy).(uint32))
				}
			}
		case reflect.Uint64:
			if kx == kindConst {
				x := vx.(uint64)
				return func(fr *frame) {
					fr.setReg(ir, x&^fr.reg(iy).(uint64))
				}
			} else if ky == kindConst {
				y := vy.(uint64)
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(uint64)&^y)
				}
			} else {
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(uint64)&^fr.reg(iy).(uint64))
				}
			}
		case reflect.Uintptr:
			if kx == kindConst {
				x := vx.(uintptr)
				return func(fr *frame) {
					fr.setReg(ir, x&^fr.reg(iy).(uintptr))
				}
			} else if ky == kindConst {
				y := vy.(uintptr)
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(uintptr)&^y)
				}
			} else {
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(uintptr)&^fr.reg(iy).(uintptr))
				}
			}

		}
	} else {
		r := reflect.New(typ).Elem()
		switch typ.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			if kx == kindConst {
				x := reflect.ValueOf(vx).Int()
				return func(fr *frame) {
					y := reflect.ValueOf(fr.reg(iy)).Int()
					r.SetInt(x &^ y)
					fr.setReg(ir, r.Interface())
				}
			} else if ky == kindConst {
				y := reflect.ValueOf(vy).Int()
				return func(fr *frame) {
					x := reflect.ValueOf(fr.reg(ix)).Int()
					r.SetInt(x &^ y)
					fr.setReg(ir, r.Interface())
				}
			} else {
				return func(fr *frame) {
					x := reflect.ValueOf(fr.reg(ix)).Int()
					y := reflect.ValueOf(fr.reg(iy)).Int()
					r.SetInt(x &^ y)
					fr.setReg(ir, r.Interface())
				}
			}
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
			if kx == kindConst {
				x := reflect.ValueOf(vx).Uint()
				return func(fr *frame) {
					y := reflect.ValueOf(fr.reg(iy)).Uint()
					r.SetUint(x &^ y)
					fr.setReg(ir, r.Interface())
				}
			} else if ky == kindConst {
				y := reflect.ValueOf(vy).Uint()
				return func(fr *frame) {
					x := reflect.ValueOf(fr.reg(ix)).Uint()
					r.SetUint(x &^ y)
					fr.setReg(ir, r.Interface())
				}
			} else {
				return func(fr *frame) {
					x := reflect.ValueOf(fr.reg(ix)).Uint()
					y := reflect.ValueOf(fr.reg(iy)).Uint()
					r.SetUint(x &^ y)
					fr.setReg(ir, r.Interface())
				}
			}

		}
	}
	panic("unreachable")
}

func makeBinOpLSS(pfn *function, instr *ssa.BinOp) func(fr *frame) {
	ir := pfn.regIndex(instr)
	ix, kx, vx := pfn.regIndex3(instr.X)
	iy, ky, vy := pfn.regIndex3(instr.Y)
	typ := pfn.Interp.preToType(instr.X.Type())
	if typ.PkgPath() == "" {
		switch typ.Kind() {
		case reflect.Int:
			if kx == kindConst {
				x := vx.(int)
				return func(fr *frame) {
					fr.setReg(ir, x < fr.reg(iy).(int))
				}
			} else if ky == kindConst {
				y := vy.(int)
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(int) < y)
				}
			} else {
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(int) < fr.reg(iy).(int))
				}
			}
		case reflect.Int8:
			if kx == kindConst {
				x := vx.(int8)
				return func(fr *frame) {
					fr.setReg(ir, x < fr.reg(iy).(int8))
				}
			} else if ky == kindConst {
				y := vy.(int8)
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(int8) < y)
				}
			} else {
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(int8) < fr.reg(iy).(int8))
				}
			}
		case reflect.Int16:
			if kx == kindConst {
				x := vx.(int16)
				return func(fr *frame) {
					fr.setReg(ir, x < fr.reg(iy).(int16))
				}
			} else if ky == kindConst {
				y := vy.(int16)
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(int16) < y)
				}
			} else {
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(int16) < fr.reg(iy).(int16))
				}
			}
		case reflect.Int32:
			if kx == kindConst {
				x := vx.(int32)
				return func(fr *frame) {
					fr.setReg(ir, x < fr.reg(iy).(int32))
				}
			} else if ky == kindConst {
				y := vy.(int32)
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(int32) < y)
				}
			} else {
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(int32) < fr.reg(iy).(int32))
				}
			}
		case reflect.Int64:
			if kx == kindConst {
				x := vx.(int64)
				return func(fr *frame) {
					fr.setReg(ir, x < fr.reg(iy).(int64))
				}
			} else if ky == kindConst {
				y := vy.(int64)
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(int64) < y)
				}
			} else {
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(int64) < fr.reg(iy).(int64))
				}
			}
		case reflect.Uint:
			if kx == kindConst {
				x := vx.(uint)
				return func(fr *frame) {
					fr.setReg(ir, x < fr.reg(iy).(uint))
				}
			} else if ky == kindConst {
				y := vy.(uint)
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(uint) < y)
				}
			} else {
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(uint) < fr.reg(iy).(uint))
				}
			}
		case reflect.Uint8:
			if kx == kindConst {
				x := vx.(uint8)
				return func(fr *frame) {
					fr.setReg(ir, x < fr.reg(iy).(uint8))
				}
			} else if ky == kindConst {
				y := vy.(uint8)
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(uint8) < y)
				}
			} else {
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(uint8) < fr.reg(iy).(uint8))
				}
			}
		case reflect.Uint16:
			if kx == kindConst {
				x := vx.(uint16)
				return func(fr *frame) {
					fr.setReg(ir, x < fr.reg(iy).(uint16))
				}
			} else if ky == kindConst {
				y := vy.(uint16)
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(uint16) < y)
				}
			} else {
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(uint16) < fr.reg(iy).(uint16))
				}
			}
		case reflect.Uint32:
			if kx == kindConst {
				x := vx.(uint32)
				return func(fr *frame) {
					fr.setReg(ir, x < fr.reg(iy).(uint32))
				}
			} else if ky == kindConst {
				y := vy.(uint32)
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(uint32) < y)
				}
			} else {
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(uint32) < fr.reg(iy).(uint32))
				}
			}
		case reflect.Uint64:
			if kx == kindConst {
				x := vx.(uint64)
				return func(fr *frame) {
					fr.setReg(ir, x < fr.reg(iy).(uint64))
				}
			} else if ky == kindConst {
				y := vy.(uint64)
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(uint64) < y)
				}
			} else {
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(uint64) < fr.reg(iy).(uint64))
				}
			}
		case reflect.Uintptr:
			if kx == kindConst {
				x := vx.(uintptr)
				return func(fr *frame) {
					fr.setReg(ir, x < fr.reg(iy).(uintptr))
				}
			} else if ky == kindConst {
				y := vy.(uintptr)
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(uintptr) < y)
				}
			} else {
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(uintptr) < fr.reg(iy).(uintptr))
				}
			}
		case reflect.Float32:
			if kx == kindConst {
				x := vx.(float32)
				return func(fr *frame) {
					fr.setReg(ir, x < fr.reg(iy).(float32))
				}
			} else if ky == kindConst {
				y := vy.(float32)
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(float32) < y)
				}
			} else {
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(float32) < fr.reg(iy).(float32))
				}
			}
		case reflect.Float64:
			if kx == kindConst {
				x := vx.(float64)
				return func(fr *frame) {
					fr.setReg(ir, x < fr.reg(iy).(float64))
				}
			} else if ky == kindConst {
				y := vy.(float64)
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(float64) < y)
				}
			} else {
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(float64) < fr.reg(iy).(float64))
				}
			}

		case reflect.String:
			if kx == kindConst {
				x := vx.(string)
				return func(fr *frame) {
					fr.setReg(ir, x < fr.reg(iy).(string))
				}
			} else if ky == kindConst {
				y := vy.(string)
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(string) < y)
				}
			} else {
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(string) < fr.reg(iy).(string))
				}
			}
		}
	} else {
		r := reflect.New(typ).Elem()
		switch typ.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			if kx == kindConst {
				x := reflect.ValueOf(vx).Int()
				return func(fr *frame) {
					y := reflect.ValueOf(fr.reg(iy)).Int()
					r.SetBool(x < y)
					fr.setReg(ir, r.Interface())
				}
			} else if ky == kindConst {
				y := reflect.ValueOf(vy).Int()
				return func(fr *frame) {
					x := reflect.ValueOf(fr.reg(ix)).Int()
					r.SetBool(x < y)
					fr.setReg(ir, r.Interface())
				}
			} else {
				return func(fr *frame) {
					x := reflect.ValueOf(fr.reg(ix)).Int()
					y := reflect.ValueOf(fr.reg(iy)).Int()
					r.SetBool(x < y)
					fr.setReg(ir, r.Interface())
				}
			}
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
			if kx == kindConst {
				x := reflect.ValueOf(vx).Uint()
				return func(fr *frame) {
					y := reflect.ValueOf(fr.reg(iy)).Uint()
					r.SetBool(x < y)
					fr.setReg(ir, r.Interface())
				}
			} else if ky == kindConst {
				y := reflect.ValueOf(vy).Uint()
				return func(fr *frame) {
					x := reflect.ValueOf(fr.reg(ix)).Uint()
					r.SetBool(x < y)
					fr.setReg(ir, r.Interface())
				}
			} else {
				return func(fr *frame) {
					x := reflect.ValueOf(fr.reg(ix)).Uint()
					y := reflect.ValueOf(fr.reg(iy)).Uint()
					r.SetBool(x < y)
					fr.setReg(ir, r.Interface())
				}
			}
		case reflect.Float32, reflect.Float64:
			if kx == kindConst {
				x := reflect.ValueOf(vx).Float()
				return func(fr *frame) {
					y := reflect.ValueOf(fr.reg(iy)).Float()
					r.SetBool(x < y)
					fr.setReg(ir, r.Interface())
				}
			} else if ky == kindConst {
				y := reflect.ValueOf(vy).Float()
				return func(fr *frame) {
					x := reflect.ValueOf(fr.reg(ix)).Float()
					r.SetBool(x < y)
					fr.setReg(ir, r.Interface())
				}
			} else {
				return func(fr *frame) {
					x := reflect.ValueOf(fr.reg(ix)).Float()
					y := reflect.ValueOf(fr.reg(iy)).Float()
					r.SetBool(x < y)
					fr.setReg(ir, r.Interface())
				}
			}

		case reflect.String:
			if kx == kindConst {
				x := reflect.ValueOf(vx).String()
				return func(fr *frame) {
					y := reflect.ValueOf(fr.reg(iy)).String()
					r.SetBool(x < y)
					fr.setReg(ir, r.Interface())
				}
			} else if ky == kindConst {
				y := reflect.ValueOf(vy).String()
				return func(fr *frame) {
					x := reflect.ValueOf(fr.reg(ix)).String()
					r.SetBool(x < y)
					fr.setReg(ir, r.Interface())
				}
			} else {
				return func(fr *frame) {
					x := reflect.ValueOf(fr.reg(ix)).String()
					y := reflect.ValueOf(fr.reg(iy)).String()
					r.SetBool(x < y)
					fr.setReg(ir, r.Interface())
				}
			}
		}
	}
	panic("unreachable")
}

func makeBinOpLEQ(pfn *function, instr *ssa.BinOp) func(fr *frame) {
	ir := pfn.regIndex(instr)
	ix, kx, vx := pfn.regIndex3(instr.X)
	iy, ky, vy := pfn.regIndex3(instr.Y)
	typ := pfn.Interp.preToType(instr.X.Type())
	if typ.PkgPath() == "" {
		switch typ.Kind() {
		case reflect.Int:
			if kx == kindConst {
				x := vx.(int)
				return func(fr *frame) {
					fr.setReg(ir, x <= fr.reg(iy).(int))
				}
			} else if ky == kindConst {
				y := vy.(int)
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(int) <= y)
				}
			} else {
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(int) <= fr.reg(iy).(int))
				}
			}
		case reflect.Int8:
			if kx == kindConst {
				x := vx.(int8)
				return func(fr *frame) {
					fr.setReg(ir, x <= fr.reg(iy).(int8))
				}
			} else if ky == kindConst {
				y := vy.(int8)
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(int8) <= y)
				}
			} else {
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(int8) <= fr.reg(iy).(int8))
				}
			}
		case reflect.Int16:
			if kx == kindConst {
				x := vx.(int16)
				return func(fr *frame) {
					fr.setReg(ir, x <= fr.reg(iy).(int16))
				}
			} else if ky == kindConst {
				y := vy.(int16)
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(int16) <= y)
				}
			} else {
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(int16) <= fr.reg(iy).(int16))
				}
			}
		case reflect.Int32:
			if kx == kindConst {
				x := vx.(int32)
				return func(fr *frame) {
					fr.setReg(ir, x <= fr.reg(iy).(int32))
				}
			} else if ky == kindConst {
				y := vy.(int32)
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(int32) <= y)
				}
			} else {
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(int32) <= fr.reg(iy).(int32))
				}
			}
		case reflect.Int64:
			if kx == kindConst {
				x := vx.(int64)
				return func(fr *frame) {
					fr.setReg(ir, x <= fr.reg(iy).(int64))
				}
			} else if ky == kindConst {
				y := vy.(int64)
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(int64) <= y)
				}
			} else {
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(int64) <= fr.reg(iy).(int64))
				}
			}
		case reflect.Uint:
			if kx == kindConst {
				x := vx.(uint)
				return func(fr *frame) {
					fr.setReg(ir, x <= fr.reg(iy).(uint))
				}
			} else if ky == kindConst {
				y := vy.(uint)
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(uint) <= y)
				}
			} else {
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(uint) <= fr.reg(iy).(uint))
				}
			}
		case reflect.Uint8:
			if kx == kindConst {
				x := vx.(uint8)
				return func(fr *frame) {
					fr.setReg(ir, x <= fr.reg(iy).(uint8))
				}
			} else if ky == kindConst {
				y := vy.(uint8)
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(uint8) <= y)
				}
			} else {
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(uint8) <= fr.reg(iy).(uint8))
				}
			}
		case reflect.Uint16:
			if kx == kindConst {
				x := vx.(uint16)
				return func(fr *frame) {
					fr.setReg(ir, x <= fr.reg(iy).(uint16))
				}
			} else if ky == kindConst {
				y := vy.(uint16)
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(uint16) <= y)
				}
			} else {
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(uint16) <= fr.reg(iy).(uint16))
				}
			}
		case reflect.Uint32:
			if kx == kindConst {
				x := vx.(uint32)
				return func(fr *frame) {
					fr.setReg(ir, x <= fr.reg(iy).(uint32))
				}
			} else if ky == kindConst {
				y := vy.(uint32)
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(uint32) <= y)
				}
			} else {
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(uint32) <= fr.reg(iy).(uint32))
				}
			}
		case reflect.Uint64:
			if kx == kindConst {
				x := vx.(uint64)
				return func(fr *frame) {
					fr.setReg(ir, x <= fr.reg(iy).(uint64))
				}
			} else if ky == kindConst {
				y := vy.(uint64)
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(uint64) <= y)
				}
			} else {
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(uint64) <= fr.reg(iy).(uint64))
				}
			}
		case reflect.Uintptr:
			if kx == kindConst {
				x := vx.(uintptr)
				return func(fr *frame) {
					fr.setReg(ir, x <= fr.reg(iy).(uintptr))
				}
			} else if ky == kindConst {
				y := vy.(uintptr)
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(uintptr) <= y)
				}
			} else {
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(uintptr) <= fr.reg(iy).(uintptr))
				}
			}
		case reflect.Float32:
			if kx == kindConst {
				x := vx.(float32)
				return func(fr *frame) {
					fr.setReg(ir, x <= fr.reg(iy).(float32))
				}
			} else if ky == kindConst {
				y := vy.(float32)
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(float32) <= y)
				}
			} else {
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(float32) <= fr.reg(iy).(float32))
				}
			}
		case reflect.Float64:
			if kx == kindConst {
				x := vx.(float64)
				return func(fr *frame) {
					fr.setReg(ir, x <= fr.reg(iy).(float64))
				}
			} else if ky == kindConst {
				y := vy.(float64)
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(float64) <= y)
				}
			} else {
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(float64) <= fr.reg(iy).(float64))
				}
			}

		case reflect.String:
			if kx == kindConst {
				x := vx.(string)
				return func(fr *frame) {
					fr.setReg(ir, x <= fr.reg(iy).(string))
				}
			} else if ky == kindConst {
				y := vy.(string)
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(string) <= y)
				}
			} else {
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(string) <= fr.reg(iy).(string))
				}
			}
		}
	} else {
		r := reflect.New(typ).Elem()
		switch typ.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			if kx == kindConst {
				x := reflect.ValueOf(vx).Int()
				return func(fr *frame) {
					y := reflect.ValueOf(fr.reg(iy)).Int()
					r.SetBool(x <= y)
					fr.setReg(ir, r.Interface())
				}
			} else if ky == kindConst {
				y := reflect.ValueOf(vy).Int()
				return func(fr *frame) {
					x := reflect.ValueOf(fr.reg(ix)).Int()
					r.SetBool(x <= y)
					fr.setReg(ir, r.Interface())
				}
			} else {
				return func(fr *frame) {
					x := reflect.ValueOf(fr.reg(ix)).Int()
					y := reflect.ValueOf(fr.reg(iy)).Int()
					r.SetBool(x <= y)
					fr.setReg(ir, r.Interface())
				}
			}
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
			if kx == kindConst {
				x := reflect.ValueOf(vx).Uint()
				return func(fr *frame) {
					y := reflect.ValueOf(fr.reg(iy)).Uint()
					r.SetBool(x <= y)
					fr.setReg(ir, r.Interface())
				}
			} else if ky == kindConst {
				y := reflect.ValueOf(vy).Uint()
				return func(fr *frame) {
					x := reflect.ValueOf(fr.reg(ix)).Uint()
					r.SetBool(x <= y)
					fr.setReg(ir, r.Interface())
				}
			} else {
				return func(fr *frame) {
					x := reflect.ValueOf(fr.reg(ix)).Uint()
					y := reflect.ValueOf(fr.reg(iy)).Uint()
					r.SetBool(x <= y)
					fr.setReg(ir, r.Interface())
				}
			}
		case reflect.Float32, reflect.Float64:
			if kx == kindConst {
				x := reflect.ValueOf(vx).Float()
				return func(fr *frame) {
					y := reflect.ValueOf(fr.reg(iy)).Float()
					r.SetBool(x <= y)
					fr.setReg(ir, r.Interface())
				}
			} else if ky == kindConst {
				y := reflect.ValueOf(vy).Float()
				return func(fr *frame) {
					x := reflect.ValueOf(fr.reg(ix)).Float()
					r.SetBool(x <= y)
					fr.setReg(ir, r.Interface())
				}
			} else {
				return func(fr *frame) {
					x := reflect.ValueOf(fr.reg(ix)).Float()
					y := reflect.ValueOf(fr.reg(iy)).Float()
					r.SetBool(x <= y)
					fr.setReg(ir, r.Interface())
				}
			}

		case reflect.String:
			if kx == kindConst {
				x := reflect.ValueOf(vx).String()
				return func(fr *frame) {
					y := reflect.ValueOf(fr.reg(iy)).String()
					r.SetBool(x <= y)
					fr.setReg(ir, r.Interface())
				}
			} else if ky == kindConst {
				y := reflect.ValueOf(vy).String()
				return func(fr *frame) {
					x := reflect.ValueOf(fr.reg(ix)).String()
					r.SetBool(x <= y)
					fr.setReg(ir, r.Interface())
				}
			} else {
				return func(fr *frame) {
					x := reflect.ValueOf(fr.reg(ix)).String()
					y := reflect.ValueOf(fr.reg(iy)).String()
					r.SetBool(x <= y)
					fr.setReg(ir, r.Interface())
				}
			}
		}
	}
	panic("unreachable")
}

func makeBinOpGTR(pfn *function, instr *ssa.BinOp) func(fr *frame) {
	ir := pfn.regIndex(instr)
	ix, kx, vx := pfn.regIndex3(instr.X)
	iy, ky, vy := pfn.regIndex3(instr.Y)
	typ := pfn.Interp.preToType(instr.X.Type())
	if typ.PkgPath() == "" {
		switch typ.Kind() {
		case reflect.Int:
			if kx == kindConst {
				x := vx.(int)
				return func(fr *frame) {
					fr.setReg(ir, x > fr.reg(iy).(int))
				}
			} else if ky == kindConst {
				y := vy.(int)
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(int) > y)
				}
			} else {
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(int) > fr.reg(iy).(int))
				}
			}
		case reflect.Int8:
			if kx == kindConst {
				x := vx.(int8)
				return func(fr *frame) {
					fr.setReg(ir, x > fr.reg(iy).(int8))
				}
			} else if ky == kindConst {
				y := vy.(int8)
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(int8) > y)
				}
			} else {
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(int8) > fr.reg(iy).(int8))
				}
			}
		case reflect.Int16:
			if kx == kindConst {
				x := vx.(int16)
				return func(fr *frame) {
					fr.setReg(ir, x > fr.reg(iy).(int16))
				}
			} else if ky == kindConst {
				y := vy.(int16)
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(int16) > y)
				}
			} else {
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(int16) > fr.reg(iy).(int16))
				}
			}
		case reflect.Int32:
			if kx == kindConst {
				x := vx.(int32)
				return func(fr *frame) {
					fr.setReg(ir, x > fr.reg(iy).(int32))
				}
			} else if ky == kindConst {
				y := vy.(int32)
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(int32) > y)
				}
			} else {
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(int32) > fr.reg(iy).(int32))
				}
			}
		case reflect.Int64:
			if kx == kindConst {
				x := vx.(int64)
				return func(fr *frame) {
					fr.setReg(ir, x > fr.reg(iy).(int64))
				}
			} else if ky == kindConst {
				y := vy.(int64)
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(int64) > y)
				}
			} else {
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(int64) > fr.reg(iy).(int64))
				}
			}
		case reflect.Uint:
			if kx == kindConst {
				x := vx.(uint)
				return func(fr *frame) {
					fr.setReg(ir, x > fr.reg(iy).(uint))
				}
			} else if ky == kindConst {
				y := vy.(uint)
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(uint) > y)
				}
			} else {
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(uint) > fr.reg(iy).(uint))
				}
			}
		case reflect.Uint8:
			if kx == kindConst {
				x := vx.(uint8)
				return func(fr *frame) {
					fr.setReg(ir, x > fr.reg(iy).(uint8))
				}
			} else if ky == kindConst {
				y := vy.(uint8)
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(uint8) > y)
				}
			} else {
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(uint8) > fr.reg(iy).(uint8))
				}
			}
		case reflect.Uint16:
			if kx == kindConst {
				x := vx.(uint16)
				return func(fr *frame) {
					fr.setReg(ir, x > fr.reg(iy).(uint16))
				}
			} else if ky == kindConst {
				y := vy.(uint16)
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(uint16) > y)
				}
			} else {
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(uint16) > fr.reg(iy).(uint16))
				}
			}
		case reflect.Uint32:
			if kx == kindConst {
				x := vx.(uint32)
				return func(fr *frame) {
					fr.setReg(ir, x > fr.reg(iy).(uint32))
				}
			} else if ky == kindConst {
				y := vy.(uint32)
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(uint32) > y)
				}
			} else {
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(uint32) > fr.reg(iy).(uint32))
				}
			}
		case reflect.Uint64:
			if kx == kindConst {
				x := vx.(uint64)
				return func(fr *frame) {
					fr.setReg(ir, x > fr.reg(iy).(uint64))
				}
			} else if ky == kindConst {
				y := vy.(uint64)
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(uint64) > y)
				}
			} else {
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(uint64) > fr.reg(iy).(uint64))
				}
			}
		case reflect.Uintptr:
			if kx == kindConst {
				x := vx.(uintptr)
				return func(fr *frame) {
					fr.setReg(ir, x > fr.reg(iy).(uintptr))
				}
			} else if ky == kindConst {
				y := vy.(uintptr)
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(uintptr) > y)
				}
			} else {
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(uintptr) > fr.reg(iy).(uintptr))
				}
			}
		case reflect.Float32:
			if kx == kindConst {
				x := vx.(float32)
				return func(fr *frame) {
					fr.setReg(ir, x > fr.reg(iy).(float32))
				}
			} else if ky == kindConst {
				y := vy.(float32)
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(float32) > y)
				}
			} else {
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(float32) > fr.reg(iy).(float32))
				}
			}
		case reflect.Float64:
			if kx == kindConst {
				x := vx.(float64)
				return func(fr *frame) {
					fr.setReg(ir, x > fr.reg(iy).(float64))
				}
			} else if ky == kindConst {
				y := vy.(float64)
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(float64) > y)
				}
			} else {
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(float64) > fr.reg(iy).(float64))
				}
			}

		case reflect.String:
			if kx == kindConst {
				x := vx.(string)
				return func(fr *frame) {
					fr.setReg(ir, x > fr.reg(iy).(string))
				}
			} else if ky == kindConst {
				y := vy.(string)
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(string) > y)
				}
			} else {
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(string) > fr.reg(iy).(string))
				}
			}
		}
	} else {
		r := reflect.New(typ).Elem()
		switch typ.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			if kx == kindConst {
				x := reflect.ValueOf(vx).Int()
				return func(fr *frame) {
					y := reflect.ValueOf(fr.reg(iy)).Int()
					r.SetBool(x > y)
					fr.setReg(ir, r.Interface())
				}
			} else if ky == kindConst {
				y := reflect.ValueOf(vy).Int()
				return func(fr *frame) {
					x := reflect.ValueOf(fr.reg(ix)).Int()
					r.SetBool(x > y)
					fr.setReg(ir, r.Interface())
				}
			} else {
				return func(fr *frame) {
					x := reflect.ValueOf(fr.reg(ix)).Int()
					y := reflect.ValueOf(fr.reg(iy)).Int()
					r.SetBool(x > y)
					fr.setReg(ir, r.Interface())
				}
			}
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
			if kx == kindConst {
				x := reflect.ValueOf(vx).Uint()
				return func(fr *frame) {
					y := reflect.ValueOf(fr.reg(iy)).Uint()
					r.SetBool(x > y)
					fr.setReg(ir, r.Interface())
				}
			} else if ky == kindConst {
				y := reflect.ValueOf(vy).Uint()
				return func(fr *frame) {
					x := reflect.ValueOf(fr.reg(ix)).Uint()
					r.SetBool(x > y)
					fr.setReg(ir, r.Interface())
				}
			} else {
				return func(fr *frame) {
					x := reflect.ValueOf(fr.reg(ix)).Uint()
					y := reflect.ValueOf(fr.reg(iy)).Uint()
					r.SetBool(x > y)
					fr.setReg(ir, r.Interface())
				}
			}
		case reflect.Float32, reflect.Float64:
			if kx == kindConst {
				x := reflect.ValueOf(vx).Float()
				return func(fr *frame) {
					y := reflect.ValueOf(fr.reg(iy)).Float()
					r.SetBool(x > y)
					fr.setReg(ir, r.Interface())
				}
			} else if ky == kindConst {
				y := reflect.ValueOf(vy).Float()
				return func(fr *frame) {
					x := reflect.ValueOf(fr.reg(ix)).Float()
					r.SetBool(x > y)
					fr.setReg(ir, r.Interface())
				}
			} else {
				return func(fr *frame) {
					x := reflect.ValueOf(fr.reg(ix)).Float()
					y := reflect.ValueOf(fr.reg(iy)).Float()
					r.SetBool(x > y)
					fr.setReg(ir, r.Interface())
				}
			}

		case reflect.String:
			if kx == kindConst {
				x := reflect.ValueOf(vx).String()
				return func(fr *frame) {
					y := reflect.ValueOf(fr.reg(iy)).String()
					r.SetBool(x > y)
					fr.setReg(ir, r.Interface())
				}
			} else if ky == kindConst {
				y := reflect.ValueOf(vy).String()
				return func(fr *frame) {
					x := reflect.ValueOf(fr.reg(ix)).String()
					r.SetBool(x > y)
					fr.setReg(ir, r.Interface())
				}
			} else {
				return func(fr *frame) {
					x := reflect.ValueOf(fr.reg(ix)).String()
					y := reflect.ValueOf(fr.reg(iy)).String()
					r.SetBool(x > y)
					fr.setReg(ir, r.Interface())
				}
			}
		}
	}
	panic("unreachable")
}

func makeBinOpGEQ(pfn *function, instr *ssa.BinOp) func(fr *frame) {
	ir := pfn.regIndex(instr)
	ix, kx, vx := pfn.regIndex3(instr.X)
	iy, ky, vy := pfn.regIndex3(instr.Y)
	typ := pfn.Interp.preToType(instr.X.Type())
	if typ.PkgPath() == "" {
		switch typ.Kind() {
		case reflect.Int:
			if kx == kindConst {
				x := vx.(int)
				return func(fr *frame) {
					fr.setReg(ir, x >= fr.reg(iy).(int))
				}
			} else if ky == kindConst {
				y := vy.(int)
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(int) >= y)
				}
			} else {
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(int) >= fr.reg(iy).(int))
				}
			}
		case reflect.Int8:
			if kx == kindConst {
				x := vx.(int8)
				return func(fr *frame) {
					fr.setReg(ir, x >= fr.reg(iy).(int8))
				}
			} else if ky == kindConst {
				y := vy.(int8)
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(int8) >= y)
				}
			} else {
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(int8) >= fr.reg(iy).(int8))
				}
			}
		case reflect.Int16:
			if kx == kindConst {
				x := vx.(int16)
				return func(fr *frame) {
					fr.setReg(ir, x >= fr.reg(iy).(int16))
				}
			} else if ky == kindConst {
				y := vy.(int16)
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(int16) >= y)
				}
			} else {
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(int16) >= fr.reg(iy).(int16))
				}
			}
		case reflect.Int32:
			if kx == kindConst {
				x := vx.(int32)
				return func(fr *frame) {
					fr.setReg(ir, x >= fr.reg(iy).(int32))
				}
			} else if ky == kindConst {
				y := vy.(int32)
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(int32) >= y)
				}
			} else {
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(int32) >= fr.reg(iy).(int32))
				}
			}
		case reflect.Int64:
			if kx == kindConst {
				x := vx.(int64)
				return func(fr *frame) {
					fr.setReg(ir, x >= fr.reg(iy).(int64))
				}
			} else if ky == kindConst {
				y := vy.(int64)
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(int64) >= y)
				}
			} else {
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(int64) >= fr.reg(iy).(int64))
				}
			}
		case reflect.Uint:
			if kx == kindConst {
				x := vx.(uint)
				return func(fr *frame) {
					fr.setReg(ir, x >= fr.reg(iy).(uint))
				}
			} else if ky == kindConst {
				y := vy.(uint)
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(uint) >= y)
				}
			} else {
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(uint) >= fr.reg(iy).(uint))
				}
			}
		case reflect.Uint8:
			if kx == kindConst {
				x := vx.(uint8)
				return func(fr *frame) {
					fr.setReg(ir, x >= fr.reg(iy).(uint8))
				}
			} else if ky == kindConst {
				y := vy.(uint8)
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(uint8) >= y)
				}
			} else {
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(uint8) >= fr.reg(iy).(uint8))
				}
			}
		case reflect.Uint16:
			if kx == kindConst {
				x := vx.(uint16)
				return func(fr *frame) {
					fr.setReg(ir, x >= fr.reg(iy).(uint16))
				}
			} else if ky == kindConst {
				y := vy.(uint16)
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(uint16) >= y)
				}
			} else {
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(uint16) >= fr.reg(iy).(uint16))
				}
			}
		case reflect.Uint32:
			if kx == kindConst {
				x := vx.(uint32)
				return func(fr *frame) {
					fr.setReg(ir, x >= fr.reg(iy).(uint32))
				}
			} else if ky == kindConst {
				y := vy.(uint32)
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(uint32) >= y)
				}
			} else {
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(uint32) >= fr.reg(iy).(uint32))
				}
			}
		case reflect.Uint64:
			if kx == kindConst {
				x := vx.(uint64)
				return func(fr *frame) {
					fr.setReg(ir, x >= fr.reg(iy).(uint64))
				}
			} else if ky == kindConst {
				y := vy.(uint64)
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(uint64) >= y)
				}
			} else {
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(uint64) >= fr.reg(iy).(uint64))
				}
			}
		case reflect.Uintptr:
			if kx == kindConst {
				x := vx.(uintptr)
				return func(fr *frame) {
					fr.setReg(ir, x >= fr.reg(iy).(uintptr))
				}
			} else if ky == kindConst {
				y := vy.(uintptr)
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(uintptr) >= y)
				}
			} else {
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(uintptr) >= fr.reg(iy).(uintptr))
				}
			}
		case reflect.Float32:
			if kx == kindConst {
				x := vx.(float32)
				return func(fr *frame) {
					fr.setReg(ir, x >= fr.reg(iy).(float32))
				}
			} else if ky == kindConst {
				y := vy.(float32)
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(float32) >= y)
				}
			} else {
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(float32) >= fr.reg(iy).(float32))
				}
			}
		case reflect.Float64:
			if kx == kindConst {
				x := vx.(float64)
				return func(fr *frame) {
					fr.setReg(ir, x >= fr.reg(iy).(float64))
				}
			} else if ky == kindConst {
				y := vy.(float64)
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(float64) >= y)
				}
			} else {
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(float64) >= fr.reg(iy).(float64))
				}
			}

		case reflect.String:
			if kx == kindConst {
				x := vx.(string)
				return func(fr *frame) {
					fr.setReg(ir, x >= fr.reg(iy).(string))
				}
			} else if ky == kindConst {
				y := vy.(string)
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(string) >= y)
				}
			} else {
				return func(fr *frame) {
					fr.setReg(ir, fr.reg(ix).(string) >= fr.reg(iy).(string))
				}
			}
		}
	} else {
		r := reflect.New(typ).Elem()
		switch typ.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			if kx == kindConst {
				x := reflect.ValueOf(vx).Int()
				return func(fr *frame) {
					y := reflect.ValueOf(fr.reg(iy)).Int()
					r.SetBool(x >= y)
					fr.setReg(ir, r.Interface())
				}
			} else if ky == kindConst {
				y := reflect.ValueOf(vy).Int()
				return func(fr *frame) {
					x := reflect.ValueOf(fr.reg(ix)).Int()
					r.SetBool(x >= y)
					fr.setReg(ir, r.Interface())
				}
			} else {
				return func(fr *frame) {
					x := reflect.ValueOf(fr.reg(ix)).Int()
					y := reflect.ValueOf(fr.reg(iy)).Int()
					r.SetBool(x >= y)
					fr.setReg(ir, r.Interface())
				}
			}
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
			if kx == kindConst {
				x := reflect.ValueOf(vx).Uint()
				return func(fr *frame) {
					y := reflect.ValueOf(fr.reg(iy)).Uint()
					r.SetBool(x >= y)
					fr.setReg(ir, r.Interface())
				}
			} else if ky == kindConst {
				y := reflect.ValueOf(vy).Uint()
				return func(fr *frame) {
					x := reflect.ValueOf(fr.reg(ix)).Uint()
					r.SetBool(x >= y)
					fr.setReg(ir, r.Interface())
				}
			} else {
				return func(fr *frame) {
					x := reflect.ValueOf(fr.reg(ix)).Uint()
					y := reflect.ValueOf(fr.reg(iy)).Uint()
					r.SetBool(x >= y)
					fr.setReg(ir, r.Interface())
				}
			}
		case reflect.Float32, reflect.Float64:
			if kx == kindConst {
				x := reflect.ValueOf(vx).Float()
				return func(fr *frame) {
					y := reflect.ValueOf(fr.reg(iy)).Float()
					r.SetBool(x >= y)
					fr.setReg(ir, r.Interface())
				}
			} else if ky == kindConst {
				y := reflect.ValueOf(vy).Float()
				return func(fr *frame) {
					x := reflect.ValueOf(fr.reg(ix)).Float()
					r.SetBool(x >= y)
					fr.setReg(ir, r.Interface())
				}
			} else {
				return func(fr *frame) {
					x := reflect.ValueOf(fr.reg(ix)).Float()
					y := reflect.ValueOf(fr.reg(iy)).Float()
					r.SetBool(x >= y)
					fr.setReg(ir, r.Interface())
				}
			}

		case reflect.String:
			if kx == kindConst {
				x := reflect.ValueOf(vx).String()
				return func(fr *frame) {
					y := reflect.ValueOf(fr.reg(iy)).String()
					r.SetBool(x >= y)
					fr.setReg(ir, r.Interface())
				}
			} else if ky == kindConst {
				y := reflect.ValueOf(vy).String()
				return func(fr *frame) {
					x := reflect.ValueOf(fr.reg(ix)).String()
					r.SetBool(x >= y)
					fr.setReg(ir, r.Interface())
				}
			} else {
				return func(fr *frame) {
					x := reflect.ValueOf(fr.reg(ix)).String()
					y := reflect.ValueOf(fr.reg(iy)).String()
					r.SetBool(x >= y)
					fr.setReg(ir, r.Interface())
				}
			}
		}
	}
	panic("unreachable")
}
