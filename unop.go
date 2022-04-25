package gossa

import (
	"reflect"

	"github.com/goplus/gossa/basic"
	"golang.org/x/tools/go/ssa"
)

func makeUnOpSUB(pfn *function, instr *ssa.UnOp) func(fr *frame) {
	ir := pfn.regIndex(instr)
	ix, kx, vx := pfn.regIndex3(instr.X)
	typ := pfn.Interp.preToType(instr.Type())

	if typ.PkgPath() == "" {
		switch typ.Kind() {
		case reflect.Int:
			if kx == kindConst {
				v := -vx.(int)
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			} else {
				return func(fr *frame) {
					v := -fr.reg(ix).(int)
					fr.setReg(ir, v)
				}
			}
		case reflect.Int8:
			if kx == kindConst {
				v := -vx.(int8)
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			} else {
				return func(fr *frame) {
					v := -fr.reg(ix).(int8)
					fr.setReg(ir, v)
				}
			}
		case reflect.Int16:
			if kx == kindConst {
				v := -vx.(int16)
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			} else {
				return func(fr *frame) {
					v := -fr.reg(ix).(int16)
					fr.setReg(ir, v)
				}
			}
		case reflect.Int32:
			if kx == kindConst {
				v := -vx.(int32)
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			} else {
				return func(fr *frame) {
					v := -fr.reg(ix).(int32)
					fr.setReg(ir, v)
				}
			}
		case reflect.Int64:
			if kx == kindConst {
				v := -vx.(int64)
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			} else {
				return func(fr *frame) {
					v := -fr.reg(ix).(int64)
					fr.setReg(ir, v)
				}
			}
		case reflect.Uint:
			if kx == kindConst {
				v := -vx.(uint)
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			} else {
				return func(fr *frame) {
					v := -fr.reg(ix).(uint)
					fr.setReg(ir, v)
				}
			}
		case reflect.Uint8:
			if kx == kindConst {
				v := -vx.(uint8)
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			} else {
				return func(fr *frame) {
					v := -fr.reg(ix).(uint8)
					fr.setReg(ir, v)
				}
			}
		case reflect.Uint16:
			if kx == kindConst {
				v := -vx.(uint16)
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			} else {
				return func(fr *frame) {
					v := -fr.reg(ix).(uint16)
					fr.setReg(ir, v)
				}
			}
		case reflect.Uint32:
			if kx == kindConst {
				v := -vx.(uint32)
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			} else {
				return func(fr *frame) {
					v := -fr.reg(ix).(uint32)
					fr.setReg(ir, v)
				}
			}
		case reflect.Uint64:
			if kx == kindConst {
				v := -vx.(uint64)
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			} else {
				return func(fr *frame) {
					v := -fr.reg(ix).(uint64)
					fr.setReg(ir, v)
				}
			}
		case reflect.Uintptr:
			if kx == kindConst {
				v := -vx.(uintptr)
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			} else {
				return func(fr *frame) {
					v := -fr.reg(ix).(uintptr)
					fr.setReg(ir, v)
				}
			}
		case reflect.Float32:
			if kx == kindConst {
				v := -vx.(float32)
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			} else {
				return func(fr *frame) {
					v := -fr.reg(ix).(float32)
					fr.setReg(ir, v)
				}
			}
		case reflect.Float64:
			if kx == kindConst {
				v := -vx.(float64)
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			} else {
				return func(fr *frame) {
					v := -fr.reg(ix).(float64)
					fr.setReg(ir, v)
				}
			}
		case reflect.Complex64:
			if kx == kindConst {
				v := -vx.(complex64)
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			} else {
				return func(fr *frame) {
					v := -fr.reg(ix).(complex64)
					fr.setReg(ir, v)
				}
			}
		case reflect.Complex128:
			if kx == kindConst {
				v := -vx.(complex128)
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			} else {
				return func(fr *frame) {
					v := -fr.reg(ix).(complex128)
					fr.setReg(ir, v)
				}
			}
		}
	} else {
		switch typ.Kind() {
		case reflect.Int:
			if kx == kindConst {
				v := basic.NegInt(vx)
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			} else {
				return func(fr *frame) {
					v := basic.NegInt(fr.reg(ix))
					fr.setReg(ir, v)
				}
			}
		case reflect.Int8:
			if kx == kindConst {
				v := basic.NegInt8(vx)
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			} else {
				return func(fr *frame) {
					v := basic.NegInt8(fr.reg(ix))
					fr.setReg(ir, v)
				}
			}
		case reflect.Int16:
			if kx == kindConst {
				v := basic.NegInt16(vx)
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			} else {
				return func(fr *frame) {
					v := basic.NegInt16(fr.reg(ix))
					fr.setReg(ir, v)
				}
			}
		case reflect.Int32:
			if kx == kindConst {
				v := basic.NegInt32(vx)
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			} else {
				return func(fr *frame) {
					v := basic.NegInt32(fr.reg(ix))
					fr.setReg(ir, v)
				}
			}
		case reflect.Int64:
			if kx == kindConst {
				v := basic.NegInt64(vx)
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			} else {
				return func(fr *frame) {
					v := basic.NegInt64(fr.reg(ix))
					fr.setReg(ir, v)
				}
			}
		case reflect.Uint:
			if kx == kindConst {
				v := basic.NegUint(vx)
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			} else {
				return func(fr *frame) {
					v := basic.NegUint(fr.reg(ix))
					fr.setReg(ir, v)
				}
			}
		case reflect.Uint8:
			if kx == kindConst {
				v := basic.NegUint8(vx)
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			} else {
				return func(fr *frame) {
					v := basic.NegUint8(fr.reg(ix))
					fr.setReg(ir, v)
				}
			}
		case reflect.Uint16:
			if kx == kindConst {
				v := basic.NegUint16(vx)
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			} else {
				return func(fr *frame) {
					v := basic.NegUint16(fr.reg(ix))
					fr.setReg(ir, v)
				}
			}
		case reflect.Uint32:
			if kx == kindConst {
				v := basic.NegUint32(vx)
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			} else {
				return func(fr *frame) {
					v := basic.NegUint32(fr.reg(ix))
					fr.setReg(ir, v)
				}
			}
		case reflect.Uint64:
			if kx == kindConst {
				v := basic.NegUint64(vx)
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			} else {
				return func(fr *frame) {
					v := basic.NegUint64(fr.reg(ix))
					fr.setReg(ir, v)
				}
			}
		case reflect.Uintptr:
			if kx == kindConst {
				v := basic.NegUintptr(vx)
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			} else {
				return func(fr *frame) {
					v := basic.NegUintptr(fr.reg(ix))
					fr.setReg(ir, v)
				}
			}
		case reflect.Float32:
			if kx == kindConst {
				v := basic.NegFloat32(vx)
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			} else {
				return func(fr *frame) {
					v := basic.NegFloat32(fr.reg(ix))
					fr.setReg(ir, v)
				}
			}
		case reflect.Float64:
			if kx == kindConst {
				v := basic.NegFloat64(vx)
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			} else {
				return func(fr *frame) {
					v := basic.NegFloat64(fr.reg(ix))
					fr.setReg(ir, v)
				}
			}
		case reflect.Complex64:
			if kx == kindConst {
				v := basic.NegComplex64(vx)
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			} else {
				return func(fr *frame) {
					v := basic.NegComplex64(fr.reg(ix))
					fr.setReg(ir, v)
				}
			}
		case reflect.Complex128:
			if kx == kindConst {
				v := basic.NegComplex128(vx)
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			} else {
				return func(fr *frame) {
					v := basic.NegComplex128(fr.reg(ix))
					fr.setReg(ir, v)
				}
			}
		}
	}
	panic("unreachable")
}
func makeUnOpXOR(pfn *function, instr *ssa.UnOp) func(fr *frame) {
	ir := pfn.regIndex(instr)
	ix, kx, vx := pfn.regIndex3(instr.X)
	typ := pfn.Interp.preToType(instr.Type())

	if typ.PkgPath() == "" {
		switch typ.Kind() {
		case reflect.Int:
			if kx == kindConst {
				v := ^vx.(int)
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			} else {
				return func(fr *frame) {
					v := ^fr.reg(ix).(int)
					fr.setReg(ir, v)
				}
			}
		case reflect.Int8:
			if kx == kindConst {
				v := ^vx.(int8)
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			} else {
				return func(fr *frame) {
					v := ^fr.reg(ix).(int8)
					fr.setReg(ir, v)
				}
			}
		case reflect.Int16:
			if kx == kindConst {
				v := ^vx.(int16)
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			} else {
				return func(fr *frame) {
					v := ^fr.reg(ix).(int16)
					fr.setReg(ir, v)
				}
			}
		case reflect.Int32:
			if kx == kindConst {
				v := ^vx.(int32)
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			} else {
				return func(fr *frame) {
					v := ^fr.reg(ix).(int32)
					fr.setReg(ir, v)
				}
			}
		case reflect.Int64:
			if kx == kindConst {
				v := ^vx.(int64)
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			} else {
				return func(fr *frame) {
					v := ^fr.reg(ix).(int64)
					fr.setReg(ir, v)
				}
			}
		case reflect.Uint:
			if kx == kindConst {
				v := ^vx.(uint)
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			} else {
				return func(fr *frame) {
					v := ^fr.reg(ix).(uint)
					fr.setReg(ir, v)
				}
			}
		case reflect.Uint8:
			if kx == kindConst {
				v := ^vx.(uint8)
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			} else {
				return func(fr *frame) {
					v := ^fr.reg(ix).(uint8)
					fr.setReg(ir, v)
				}
			}
		case reflect.Uint16:
			if kx == kindConst {
				v := ^vx.(uint16)
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			} else {
				return func(fr *frame) {
					v := ^fr.reg(ix).(uint16)
					fr.setReg(ir, v)
				}
			}
		case reflect.Uint32:
			if kx == kindConst {
				v := ^vx.(uint32)
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			} else {
				return func(fr *frame) {
					v := ^fr.reg(ix).(uint32)
					fr.setReg(ir, v)
				}
			}
		case reflect.Uint64:
			if kx == kindConst {
				v := ^vx.(uint64)
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			} else {
				return func(fr *frame) {
					v := ^fr.reg(ix).(uint64)
					fr.setReg(ir, v)
				}
			}
		case reflect.Uintptr:
			if kx == kindConst {
				v := ^vx.(uintptr)
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			} else {
				return func(fr *frame) {
					v := ^fr.reg(ix).(uintptr)
					fr.setReg(ir, v)
				}
			}
		}
	} else {
		switch typ.Kind() {
		case reflect.Int:
			if kx == kindConst {
				v := basic.XorInt(vx)
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			} else {
				return func(fr *frame) {
					v := basic.XorInt(fr.reg(ix))
					fr.setReg(ir, v)
				}
			}
		case reflect.Int8:
			if kx == kindConst {
				v := basic.XorInt8(vx)
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			} else {
				return func(fr *frame) {
					v := basic.XorInt8(fr.reg(ix))
					fr.setReg(ir, v)
				}
			}
		case reflect.Int16:
			if kx == kindConst {
				v := basic.XorInt16(vx)
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			} else {
				return func(fr *frame) {
					v := basic.XorInt16(fr.reg(ix))
					fr.setReg(ir, v)
				}
			}
		case reflect.Int32:
			if kx == kindConst {
				v := basic.XorInt32(vx)
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			} else {
				return func(fr *frame) {
					v := basic.XorInt32(fr.reg(ix))
					fr.setReg(ir, v)
				}
			}
		case reflect.Int64:
			if kx == kindConst {
				v := basic.XorInt64(vx)
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			} else {
				return func(fr *frame) {
					v := basic.XorInt64(fr.reg(ix))
					fr.setReg(ir, v)
				}
			}
		case reflect.Uint:
			if kx == kindConst {
				v := basic.XorUint(vx)
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			} else {
				return func(fr *frame) {
					v := basic.XorUint(fr.reg(ix))
					fr.setReg(ir, v)
				}
			}
		case reflect.Uint8:
			if kx == kindConst {
				v := basic.XorUint8(vx)
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			} else {
				return func(fr *frame) {
					v := basic.XorUint8(fr.reg(ix))
					fr.setReg(ir, v)
				}
			}
		case reflect.Uint16:
			if kx == kindConst {
				v := basic.XorUint16(vx)
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			} else {
				return func(fr *frame) {
					v := basic.XorUint16(fr.reg(ix))
					fr.setReg(ir, v)
				}
			}
		case reflect.Uint32:
			if kx == kindConst {
				v := basic.XorUint32(vx)
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			} else {
				return func(fr *frame) {
					v := basic.XorUint32(fr.reg(ix))
					fr.setReg(ir, v)
				}
			}
		case reflect.Uint64:
			if kx == kindConst {
				v := basic.XorUint64(vx)
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			} else {
				return func(fr *frame) {
					v := basic.XorUint64(fr.reg(ix))
					fr.setReg(ir, v)
				}
			}
		case reflect.Uintptr:
			if kx == kindConst {
				v := basic.XorUintptr(vx)
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			} else {
				return func(fr *frame) {
					v := basic.XorUintptr(fr.reg(ix))
					fr.setReg(ir, v)
				}
			}
		}
	}
	panic("unreachable")
}
