package gossa

import (
	"reflect"

	"github.com/goplus/gossa/internal/basic"
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
			if kx == kindConst && ky == kindConst {
				v := vx.(int) + vy.(int)
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := vx.(int)
				return func(fr *frame) { fr.setReg(ir, x+fr.reg(iy).(int)) }
			} else if ky == kindConst {
				y := vy.(int)
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int)+y) }
			} else {
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int)+fr.reg(iy).(int)) }
			}
		case reflect.Int8:
			if kx == kindConst && ky == kindConst {
				v := vx.(int8) + vy.(int8)
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := vx.(int8)
				return func(fr *frame) { fr.setReg(ir, x+fr.reg(iy).(int8)) }
			} else if ky == kindConst {
				y := vy.(int8)
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int8)+y) }
			} else {
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int8)+fr.reg(iy).(int8)) }
			}
		case reflect.Int16:
			if kx == kindConst && ky == kindConst {
				v := vx.(int16) + vy.(int16)
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := vx.(int16)
				return func(fr *frame) { fr.setReg(ir, x+fr.reg(iy).(int16)) }
			} else if ky == kindConst {
				y := vy.(int16)
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int16)+y) }
			} else {
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int16)+fr.reg(iy).(int16)) }
			}
		case reflect.Int32:
			if kx == kindConst && ky == kindConst {
				v := vx.(int32) + vy.(int32)
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := vx.(int32)
				return func(fr *frame) { fr.setReg(ir, x+fr.reg(iy).(int32)) }
			} else if ky == kindConst {
				y := vy.(int32)
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int32)+y) }
			} else {
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int32)+fr.reg(iy).(int32)) }
			}
		case reflect.Int64:
			if kx == kindConst && ky == kindConst {
				v := vx.(int64) + vy.(int64)
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := vx.(int64)
				return func(fr *frame) { fr.setReg(ir, x+fr.reg(iy).(int64)) }
			} else if ky == kindConst {
				y := vy.(int64)
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int64)+y) }
			} else {
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int64)+fr.reg(iy).(int64)) }
			}
		case reflect.Uint:
			if kx == kindConst && ky == kindConst {
				v := vx.(uint) + vy.(uint)
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := vx.(uint)
				return func(fr *frame) { fr.setReg(ir, x+fr.reg(iy).(uint)) }
			} else if ky == kindConst {
				y := vy.(uint)
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint)+y) }
			} else {
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint)+fr.reg(iy).(uint)) }
			}
		case reflect.Uint8:
			if kx == kindConst && ky == kindConst {
				v := vx.(uint8) + vy.(uint8)
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := vx.(uint8)
				return func(fr *frame) { fr.setReg(ir, x+fr.reg(iy).(uint8)) }
			} else if ky == kindConst {
				y := vy.(uint8)
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint8)+y) }
			} else {
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint8)+fr.reg(iy).(uint8)) }
			}
		case reflect.Uint16:
			if kx == kindConst && ky == kindConst {
				v := vx.(uint16) + vy.(uint16)
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := vx.(uint16)
				return func(fr *frame) { fr.setReg(ir, x+fr.reg(iy).(uint16)) }
			} else if ky == kindConst {
				y := vy.(uint16)
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint16)+y) }
			} else {
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint16)+fr.reg(iy).(uint16)) }
			}
		case reflect.Uint32:
			if kx == kindConst && ky == kindConst {
				v := vx.(uint32) + vy.(uint32)
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := vx.(uint32)
				return func(fr *frame) { fr.setReg(ir, x+fr.reg(iy).(uint32)) }
			} else if ky == kindConst {
				y := vy.(uint32)
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint32)+y) }
			} else {
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint32)+fr.reg(iy).(uint32)) }
			}
		case reflect.Uint64:
			if kx == kindConst && ky == kindConst {
				v := vx.(uint64) + vy.(uint64)
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := vx.(uint64)
				return func(fr *frame) { fr.setReg(ir, x+fr.reg(iy).(uint64)) }
			} else if ky == kindConst {
				y := vy.(uint64)
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint64)+y) }
			} else {
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint64)+fr.reg(iy).(uint64)) }
			}
		case reflect.Uintptr:
			if kx == kindConst && ky == kindConst {
				v := vx.(uintptr) + vy.(uintptr)
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := vx.(uintptr)
				return func(fr *frame) { fr.setReg(ir, x+fr.reg(iy).(uintptr)) }
			} else if ky == kindConst {
				y := vy.(uintptr)
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uintptr)+y) }
			} else {
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uintptr)+fr.reg(iy).(uintptr)) }
			}
		case reflect.Float32:
			if kx == kindConst && ky == kindConst {
				v := vx.(float32) + vy.(float32)
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := vx.(float32)
				return func(fr *frame) { fr.setReg(ir, x+fr.reg(iy).(float32)) }
			} else if ky == kindConst {
				y := vy.(float32)
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(float32)+y) }
			} else {
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(float32)+fr.reg(iy).(float32)) }
			}
		case reflect.Float64:
			if kx == kindConst && ky == kindConst {
				v := vx.(float64) + vy.(float64)
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := vx.(float64)
				return func(fr *frame) { fr.setReg(ir, x+fr.reg(iy).(float64)) }
			} else if ky == kindConst {
				y := vy.(float64)
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(float64)+y) }
			} else {
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(float64)+fr.reg(iy).(float64)) }
			}
		case reflect.Complex64:
			if kx == kindConst && ky == kindConst {
				v := vx.(complex64) + vy.(complex64)
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := vx.(complex64)
				return func(fr *frame) { fr.setReg(ir, x+fr.reg(iy).(complex64)) }
			} else if ky == kindConst {
				y := vy.(complex64)
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(complex64)+y) }
			} else {
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(complex64)+fr.reg(iy).(complex64)) }
			}
		case reflect.Complex128:
			if kx == kindConst && ky == kindConst {
				v := vx.(complex128) + vy.(complex128)
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := vx.(complex128)
				return func(fr *frame) { fr.setReg(ir, x+fr.reg(iy).(complex128)) }
			} else if ky == kindConst {
				y := vy.(complex128)
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(complex128)+y) }
			} else {
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(complex128)+fr.reg(iy).(complex128)) }
			}
		case reflect.String:
			if kx == kindConst && ky == kindConst {
				v := vx.(string) + vy.(string)
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := vx.(string)
				return func(fr *frame) { fr.setReg(ir, x+fr.reg(iy).(string)) }
			} else if ky == kindConst {
				y := vy.(string)
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(string)+y) }
			} else {
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(string)+fr.reg(iy).(string)) }
			}
		}
	} else {
		t := basic.TypeOfType(typ)
		switch typ.Kind() {
		case reflect.Int:
			if kx == kindConst && ky == kindConst {
				v := basic.Make(t, basic.Int(vx)+basic.Int(vy))
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := basic.Int(vx)
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, x+fr.int(iy))) }
			} else if ky == kindConst {
				y := basic.Int(vy)
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, fr.int(ix)+y)) }
			} else {
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, fr.int(ix)+fr.int(iy))) }
			}
		case reflect.Int8:
			if kx == kindConst && ky == kindConst {
				v := basic.Make(t, basic.Int8(vx)+basic.Int8(vy))
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := basic.Int8(vx)
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, x+fr.int8(iy))) }
			} else if ky == kindConst {
				y := basic.Int8(vy)
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, fr.int8(ix)+y)) }
			} else {
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, fr.int8(ix)+fr.int8(iy))) }
			}
		case reflect.Int16:
			if kx == kindConst && ky == kindConst {
				v := basic.Make(t, basic.Int16(vx)+basic.Int16(vy))
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := basic.Int16(vx)
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, x+fr.int16(iy))) }
			} else if ky == kindConst {
				y := basic.Int16(vy)
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, fr.int16(ix)+y)) }
			} else {
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, fr.int16(ix)+fr.int16(iy))) }
			}
		case reflect.Int32:
			if kx == kindConst && ky == kindConst {
				v := basic.Make(t, basic.Int32(vx)+basic.Int32(vy))
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := basic.Int32(vx)
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, x+fr.int32(iy))) }
			} else if ky == kindConst {
				y := basic.Int32(vy)
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, fr.int32(ix)+y)) }
			} else {
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, fr.int32(ix)+fr.int32(iy))) }
			}
		case reflect.Int64:
			if kx == kindConst && ky == kindConst {
				v := basic.Make(t, basic.Int64(vx)+basic.Int64(vy))
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := basic.Int64(vx)
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, x+fr.int64(iy))) }
			} else if ky == kindConst {
				y := basic.Int64(vy)
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, fr.int64(ix)+y)) }
			} else {
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, fr.int64(ix)+fr.int64(iy))) }
			}
		case reflect.Uint:
			if kx == kindConst && ky == kindConst {
				v := basic.Make(t, basic.Uint(vx)+basic.Uint(vy))
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := basic.Uint(vx)
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, x+fr.uint(iy))) }
			} else if ky == kindConst {
				y := basic.Uint(vy)
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, fr.uint(ix)+y)) }
			} else {
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, fr.uint(ix)+fr.uint(iy))) }
			}
		case reflect.Uint8:
			if kx == kindConst && ky == kindConst {
				v := basic.Make(t, basic.Uint8(vx)+basic.Uint8(vy))
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := basic.Uint8(vx)
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, x+fr.uint8(iy))) }
			} else if ky == kindConst {
				y := basic.Uint8(vy)
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, fr.uint8(ix)+y)) }
			} else {
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, fr.uint8(ix)+fr.uint8(iy))) }
			}
		case reflect.Uint16:
			if kx == kindConst && ky == kindConst {
				v := basic.Make(t, basic.Uint16(vx)+basic.Uint16(vy))
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := basic.Uint16(vx)
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, x+fr.uint16(iy))) }
			} else if ky == kindConst {
				y := basic.Uint16(vy)
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, fr.uint16(ix)+y)) }
			} else {
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, fr.uint16(ix)+fr.uint16(iy))) }
			}
		case reflect.Uint32:
			if kx == kindConst && ky == kindConst {
				v := basic.Make(t, basic.Uint32(vx)+basic.Uint32(vy))
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := basic.Uint32(vx)
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, x+fr.uint32(iy))) }
			} else if ky == kindConst {
				y := basic.Uint32(vy)
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, fr.uint32(ix)+y)) }
			} else {
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, fr.uint32(ix)+fr.uint32(iy))) }
			}
		case reflect.Uint64:
			if kx == kindConst && ky == kindConst {
				v := basic.Make(t, basic.Uint64(vx)+basic.Uint64(vy))
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := basic.Uint64(vx)
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, x+fr.uint64(iy))) }
			} else if ky == kindConst {
				y := basic.Uint64(vy)
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, fr.uint64(ix)+y)) }
			} else {
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, fr.uint64(ix)+fr.uint64(iy))) }
			}
		case reflect.Uintptr:
			if kx == kindConst && ky == kindConst {
				v := basic.Make(t, basic.Uintptr(vx)+basic.Uintptr(vy))
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := basic.Uintptr(vx)
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, x+fr.uintptr(iy))) }
			} else if ky == kindConst {
				y := basic.Uintptr(vy)
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, fr.uintptr(ix)+y)) }
			} else {
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, fr.uintptr(ix)+fr.uintptr(iy))) }
			}
		case reflect.Float32:
			if kx == kindConst && ky == kindConst {
				v := basic.Make(t, basic.Float32(vx)+basic.Float32(vy))
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := basic.Float32(vx)
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, x+fr.float32(iy))) }
			} else if ky == kindConst {
				y := basic.Float32(vy)
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, fr.float32(ix)+y)) }
			} else {
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, fr.float32(ix)+fr.float32(iy))) }
			}
		case reflect.Float64:
			if kx == kindConst && ky == kindConst {
				v := basic.Make(t, basic.Float64(vx)+basic.Float64(vy))
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := basic.Float64(vx)
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, x+fr.float64(iy))) }
			} else if ky == kindConst {
				y := basic.Float64(vy)
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, fr.float64(ix)+y)) }
			} else {
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, fr.float64(ix)+fr.float64(iy))) }
			}
		case reflect.Complex64:
			if kx == kindConst && ky == kindConst {
				v := basic.Make(t, basic.Complex64(vx)+basic.Complex64(vy))
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := basic.Complex64(vx)
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, x+fr.complex64(iy))) }
			} else if ky == kindConst {
				y := basic.Complex64(vy)
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, fr.complex64(ix)+y)) }
			} else {
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, fr.complex64(ix)+fr.complex64(iy))) }
			}
		case reflect.Complex128:
			if kx == kindConst && ky == kindConst {
				v := basic.Make(t, basic.Complex128(vx)+basic.Complex128(vy))
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := basic.Complex128(vx)
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, x+fr.complex128(iy))) }
			} else if ky == kindConst {
				y := basic.Complex128(vy)
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, fr.complex128(ix)+y)) }
			} else {
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, fr.complex128(ix)+fr.complex128(iy))) }
			}
		case reflect.String:
			if kx == kindConst && ky == kindConst {
				v := basic.Make(t, basic.String(vx)+basic.String(vy))
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := basic.String(vx)
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, x+fr.string(iy))) }
			} else if ky == kindConst {
				y := basic.String(vy)
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, fr.string(ix)+y)) }
			} else {
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, fr.string(ix)+fr.string(iy))) }
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
			if kx == kindConst && ky == kindConst {
				v := vx.(int) - vy.(int)
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := vx.(int)
				return func(fr *frame) { fr.setReg(ir, x-fr.reg(iy).(int)) }
			} else if ky == kindConst {
				y := vy.(int)
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int)-y) }
			} else {
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int)-fr.reg(iy).(int)) }
			}
		case reflect.Int8:
			if kx == kindConst && ky == kindConst {
				v := vx.(int8) - vy.(int8)
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := vx.(int8)
				return func(fr *frame) { fr.setReg(ir, x-fr.reg(iy).(int8)) }
			} else if ky == kindConst {
				y := vy.(int8)
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int8)-y) }
			} else {
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int8)-fr.reg(iy).(int8)) }
			}
		case reflect.Int16:
			if kx == kindConst && ky == kindConst {
				v := vx.(int16) - vy.(int16)
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := vx.(int16)
				return func(fr *frame) { fr.setReg(ir, x-fr.reg(iy).(int16)) }
			} else if ky == kindConst {
				y := vy.(int16)
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int16)-y) }
			} else {
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int16)-fr.reg(iy).(int16)) }
			}
		case reflect.Int32:
			if kx == kindConst && ky == kindConst {
				v := vx.(int32) - vy.(int32)
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := vx.(int32)
				return func(fr *frame) { fr.setReg(ir, x-fr.reg(iy).(int32)) }
			} else if ky == kindConst {
				y := vy.(int32)
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int32)-y) }
			} else {
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int32)-fr.reg(iy).(int32)) }
			}
		case reflect.Int64:
			if kx == kindConst && ky == kindConst {
				v := vx.(int64) - vy.(int64)
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := vx.(int64)
				return func(fr *frame) { fr.setReg(ir, x-fr.reg(iy).(int64)) }
			} else if ky == kindConst {
				y := vy.(int64)
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int64)-y) }
			} else {
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int64)-fr.reg(iy).(int64)) }
			}
		case reflect.Uint:
			if kx == kindConst && ky == kindConst {
				v := vx.(uint) - vy.(uint)
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := vx.(uint)
				return func(fr *frame) { fr.setReg(ir, x-fr.reg(iy).(uint)) }
			} else if ky == kindConst {
				y := vy.(uint)
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint)-y) }
			} else {
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint)-fr.reg(iy).(uint)) }
			}
		case reflect.Uint8:
			if kx == kindConst && ky == kindConst {
				v := vx.(uint8) - vy.(uint8)
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := vx.(uint8)
				return func(fr *frame) { fr.setReg(ir, x-fr.reg(iy).(uint8)) }
			} else if ky == kindConst {
				y := vy.(uint8)
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint8)-y) }
			} else {
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint8)-fr.reg(iy).(uint8)) }
			}
		case reflect.Uint16:
			if kx == kindConst && ky == kindConst {
				v := vx.(uint16) - vy.(uint16)
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := vx.(uint16)
				return func(fr *frame) { fr.setReg(ir, x-fr.reg(iy).(uint16)) }
			} else if ky == kindConst {
				y := vy.(uint16)
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint16)-y) }
			} else {
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint16)-fr.reg(iy).(uint16)) }
			}
		case reflect.Uint32:
			if kx == kindConst && ky == kindConst {
				v := vx.(uint32) - vy.(uint32)
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := vx.(uint32)
				return func(fr *frame) { fr.setReg(ir, x-fr.reg(iy).(uint32)) }
			} else if ky == kindConst {
				y := vy.(uint32)
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint32)-y) }
			} else {
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint32)-fr.reg(iy).(uint32)) }
			}
		case reflect.Uint64:
			if kx == kindConst && ky == kindConst {
				v := vx.(uint64) - vy.(uint64)
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := vx.(uint64)
				return func(fr *frame) { fr.setReg(ir, x-fr.reg(iy).(uint64)) }
			} else if ky == kindConst {
				y := vy.(uint64)
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint64)-y) }
			} else {
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint64)-fr.reg(iy).(uint64)) }
			}
		case reflect.Uintptr:
			if kx == kindConst && ky == kindConst {
				v := vx.(uintptr) - vy.(uintptr)
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := vx.(uintptr)
				return func(fr *frame) { fr.setReg(ir, x-fr.reg(iy).(uintptr)) }
			} else if ky == kindConst {
				y := vy.(uintptr)
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uintptr)-y) }
			} else {
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uintptr)-fr.reg(iy).(uintptr)) }
			}
		case reflect.Float32:
			if kx == kindConst && ky == kindConst {
				v := vx.(float32) - vy.(float32)
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := vx.(float32)
				return func(fr *frame) { fr.setReg(ir, x-fr.reg(iy).(float32)) }
			} else if ky == kindConst {
				y := vy.(float32)
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(float32)-y) }
			} else {
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(float32)-fr.reg(iy).(float32)) }
			}
		case reflect.Float64:
			if kx == kindConst && ky == kindConst {
				v := vx.(float64) - vy.(float64)
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := vx.(float64)
				return func(fr *frame) { fr.setReg(ir, x-fr.reg(iy).(float64)) }
			} else if ky == kindConst {
				y := vy.(float64)
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(float64)-y) }
			} else {
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(float64)-fr.reg(iy).(float64)) }
			}
		case reflect.Complex64:
			if kx == kindConst && ky == kindConst {
				v := vx.(complex64) - vy.(complex64)
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := vx.(complex64)
				return func(fr *frame) { fr.setReg(ir, x-fr.reg(iy).(complex64)) }
			} else if ky == kindConst {
				y := vy.(complex64)
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(complex64)-y) }
			} else {
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(complex64)-fr.reg(iy).(complex64)) }
			}
		case reflect.Complex128:
			if kx == kindConst && ky == kindConst {
				v := vx.(complex128) - vy.(complex128)
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := vx.(complex128)
				return func(fr *frame) { fr.setReg(ir, x-fr.reg(iy).(complex128)) }
			} else if ky == kindConst {
				y := vy.(complex128)
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(complex128)-y) }
			} else {
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(complex128)-fr.reg(iy).(complex128)) }
			}
		}
	} else {
		t := basic.TypeOfType(typ)
		switch typ.Kind() {
		case reflect.Int:
			if kx == kindConst && ky == kindConst {
				v := basic.Make(t, basic.Int(vx)-basic.Int(vy))
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := basic.Int(vx)
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, x-fr.int(iy))) }
			} else if ky == kindConst {
				y := basic.Int(vy)
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, fr.int(ix)-y)) }
			} else {
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, fr.int(ix)-fr.int(iy))) }
			}
		case reflect.Int8:
			if kx == kindConst && ky == kindConst {
				v := basic.Make(t, basic.Int8(vx)-basic.Int8(vy))
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := basic.Int8(vx)
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, x-fr.int8(iy))) }
			} else if ky == kindConst {
				y := basic.Int8(vy)
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, fr.int8(ix)-y)) }
			} else {
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, fr.int8(ix)-fr.int8(iy))) }
			}
		case reflect.Int16:
			if kx == kindConst && ky == kindConst {
				v := basic.Make(t, basic.Int16(vx)-basic.Int16(vy))
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := basic.Int16(vx)
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, x-fr.int16(iy))) }
			} else if ky == kindConst {
				y := basic.Int16(vy)
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, fr.int16(ix)-y)) }
			} else {
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, fr.int16(ix)-fr.int16(iy))) }
			}
		case reflect.Int32:
			if kx == kindConst && ky == kindConst {
				v := basic.Make(t, basic.Int32(vx)-basic.Int32(vy))
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := basic.Int32(vx)
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, x-fr.int32(iy))) }
			} else if ky == kindConst {
				y := basic.Int32(vy)
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, fr.int32(ix)-y)) }
			} else {
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, fr.int32(ix)-fr.int32(iy))) }
			}
		case reflect.Int64:
			if kx == kindConst && ky == kindConst {
				v := basic.Make(t, basic.Int64(vx)-basic.Int64(vy))
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := basic.Int64(vx)
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, x-fr.int64(iy))) }
			} else if ky == kindConst {
				y := basic.Int64(vy)
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, fr.int64(ix)-y)) }
			} else {
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, fr.int64(ix)-fr.int64(iy))) }
			}
		case reflect.Uint:
			if kx == kindConst && ky == kindConst {
				v := basic.Make(t, basic.Uint(vx)-basic.Uint(vy))
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := basic.Uint(vx)
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, x-fr.uint(iy))) }
			} else if ky == kindConst {
				y := basic.Uint(vy)
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, fr.uint(ix)-y)) }
			} else {
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, fr.uint(ix)-fr.uint(iy))) }
			}
		case reflect.Uint8:
			if kx == kindConst && ky == kindConst {
				v := basic.Make(t, basic.Uint8(vx)-basic.Uint8(vy))
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := basic.Uint8(vx)
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, x-fr.uint8(iy))) }
			} else if ky == kindConst {
				y := basic.Uint8(vy)
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, fr.uint8(ix)-y)) }
			} else {
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, fr.uint8(ix)-fr.uint8(iy))) }
			}
		case reflect.Uint16:
			if kx == kindConst && ky == kindConst {
				v := basic.Make(t, basic.Uint16(vx)-basic.Uint16(vy))
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := basic.Uint16(vx)
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, x-fr.uint16(iy))) }
			} else if ky == kindConst {
				y := basic.Uint16(vy)
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, fr.uint16(ix)-y)) }
			} else {
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, fr.uint16(ix)-fr.uint16(iy))) }
			}
		case reflect.Uint32:
			if kx == kindConst && ky == kindConst {
				v := basic.Make(t, basic.Uint32(vx)-basic.Uint32(vy))
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := basic.Uint32(vx)
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, x-fr.uint32(iy))) }
			} else if ky == kindConst {
				y := basic.Uint32(vy)
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, fr.uint32(ix)-y)) }
			} else {
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, fr.uint32(ix)-fr.uint32(iy))) }
			}
		case reflect.Uint64:
			if kx == kindConst && ky == kindConst {
				v := basic.Make(t, basic.Uint64(vx)-basic.Uint64(vy))
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := basic.Uint64(vx)
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, x-fr.uint64(iy))) }
			} else if ky == kindConst {
				y := basic.Uint64(vy)
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, fr.uint64(ix)-y)) }
			} else {
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, fr.uint64(ix)-fr.uint64(iy))) }
			}
		case reflect.Uintptr:
			if kx == kindConst && ky == kindConst {
				v := basic.Make(t, basic.Uintptr(vx)-basic.Uintptr(vy))
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := basic.Uintptr(vx)
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, x-fr.uintptr(iy))) }
			} else if ky == kindConst {
				y := basic.Uintptr(vy)
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, fr.uintptr(ix)-y)) }
			} else {
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, fr.uintptr(ix)-fr.uintptr(iy))) }
			}
		case reflect.Float32:
			if kx == kindConst && ky == kindConst {
				v := basic.Make(t, basic.Float32(vx)-basic.Float32(vy))
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := basic.Float32(vx)
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, x-fr.float32(iy))) }
			} else if ky == kindConst {
				y := basic.Float32(vy)
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, fr.float32(ix)-y)) }
			} else {
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, fr.float32(ix)-fr.float32(iy))) }
			}
		case reflect.Float64:
			if kx == kindConst && ky == kindConst {
				v := basic.Make(t, basic.Float64(vx)-basic.Float64(vy))
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := basic.Float64(vx)
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, x-fr.float64(iy))) }
			} else if ky == kindConst {
				y := basic.Float64(vy)
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, fr.float64(ix)-y)) }
			} else {
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, fr.float64(ix)-fr.float64(iy))) }
			}
		case reflect.Complex64:
			if kx == kindConst && ky == kindConst {
				v := basic.Make(t, basic.Complex64(vx)-basic.Complex64(vy))
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := basic.Complex64(vx)
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, x-fr.complex64(iy))) }
			} else if ky == kindConst {
				y := basic.Complex64(vy)
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, fr.complex64(ix)-y)) }
			} else {
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, fr.complex64(ix)-fr.complex64(iy))) }
			}
		case reflect.Complex128:
			if kx == kindConst && ky == kindConst {
				v := basic.Make(t, basic.Complex128(vx)-basic.Complex128(vy))
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := basic.Complex128(vx)
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, x-fr.complex128(iy))) }
			} else if ky == kindConst {
				y := basic.Complex128(vy)
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, fr.complex128(ix)-y)) }
			} else {
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, fr.complex128(ix)-fr.complex128(iy))) }
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
			if kx == kindConst && ky == kindConst {
				v := vx.(int) * vy.(int)
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := vx.(int)
				return func(fr *frame) { fr.setReg(ir, x*fr.reg(iy).(int)) }
			} else if ky == kindConst {
				y := vy.(int)
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int)*y) }
			} else {
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int)*fr.reg(iy).(int)) }
			}
		case reflect.Int8:
			if kx == kindConst && ky == kindConst {
				v := vx.(int8) * vy.(int8)
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := vx.(int8)
				return func(fr *frame) { fr.setReg(ir, x*fr.reg(iy).(int8)) }
			} else if ky == kindConst {
				y := vy.(int8)
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int8)*y) }
			} else {
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int8)*fr.reg(iy).(int8)) }
			}
		case reflect.Int16:
			if kx == kindConst && ky == kindConst {
				v := vx.(int16) * vy.(int16)
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := vx.(int16)
				return func(fr *frame) { fr.setReg(ir, x*fr.reg(iy).(int16)) }
			} else if ky == kindConst {
				y := vy.(int16)
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int16)*y) }
			} else {
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int16)*fr.reg(iy).(int16)) }
			}
		case reflect.Int32:
			if kx == kindConst && ky == kindConst {
				v := vx.(int32) * vy.(int32)
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := vx.(int32)
				return func(fr *frame) { fr.setReg(ir, x*fr.reg(iy).(int32)) }
			} else if ky == kindConst {
				y := vy.(int32)
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int32)*y) }
			} else {
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int32)*fr.reg(iy).(int32)) }
			}
		case reflect.Int64:
			if kx == kindConst && ky == kindConst {
				v := vx.(int64) * vy.(int64)
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := vx.(int64)
				return func(fr *frame) { fr.setReg(ir, x*fr.reg(iy).(int64)) }
			} else if ky == kindConst {
				y := vy.(int64)
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int64)*y) }
			} else {
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int64)*fr.reg(iy).(int64)) }
			}
		case reflect.Uint:
			if kx == kindConst && ky == kindConst {
				v := vx.(uint) * vy.(uint)
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := vx.(uint)
				return func(fr *frame) { fr.setReg(ir, x*fr.reg(iy).(uint)) }
			} else if ky == kindConst {
				y := vy.(uint)
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint)*y) }
			} else {
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint)*fr.reg(iy).(uint)) }
			}
		case reflect.Uint8:
			if kx == kindConst && ky == kindConst {
				v := vx.(uint8) * vy.(uint8)
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := vx.(uint8)
				return func(fr *frame) { fr.setReg(ir, x*fr.reg(iy).(uint8)) }
			} else if ky == kindConst {
				y := vy.(uint8)
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint8)*y) }
			} else {
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint8)*fr.reg(iy).(uint8)) }
			}
		case reflect.Uint16:
			if kx == kindConst && ky == kindConst {
				v := vx.(uint16) * vy.(uint16)
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := vx.(uint16)
				return func(fr *frame) { fr.setReg(ir, x*fr.reg(iy).(uint16)) }
			} else if ky == kindConst {
				y := vy.(uint16)
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint16)*y) }
			} else {
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint16)*fr.reg(iy).(uint16)) }
			}
		case reflect.Uint32:
			if kx == kindConst && ky == kindConst {
				v := vx.(uint32) * vy.(uint32)
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := vx.(uint32)
				return func(fr *frame) { fr.setReg(ir, x*fr.reg(iy).(uint32)) }
			} else if ky == kindConst {
				y := vy.(uint32)
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint32)*y) }
			} else {
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint32)*fr.reg(iy).(uint32)) }
			}
		case reflect.Uint64:
			if kx == kindConst && ky == kindConst {
				v := vx.(uint64) * vy.(uint64)
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := vx.(uint64)
				return func(fr *frame) { fr.setReg(ir, x*fr.reg(iy).(uint64)) }
			} else if ky == kindConst {
				y := vy.(uint64)
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint64)*y) }
			} else {
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint64)*fr.reg(iy).(uint64)) }
			}
		case reflect.Uintptr:
			if kx == kindConst && ky == kindConst {
				v := vx.(uintptr) * vy.(uintptr)
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := vx.(uintptr)
				return func(fr *frame) { fr.setReg(ir, x*fr.reg(iy).(uintptr)) }
			} else if ky == kindConst {
				y := vy.(uintptr)
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uintptr)*y) }
			} else {
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uintptr)*fr.reg(iy).(uintptr)) }
			}
		case reflect.Float32:
			if kx == kindConst && ky == kindConst {
				v := vx.(float32) * vy.(float32)
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := vx.(float32)
				return func(fr *frame) { fr.setReg(ir, x*fr.reg(iy).(float32)) }
			} else if ky == kindConst {
				y := vy.(float32)
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(float32)*y) }
			} else {
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(float32)*fr.reg(iy).(float32)) }
			}
		case reflect.Float64:
			if kx == kindConst && ky == kindConst {
				v := vx.(float64) * vy.(float64)
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := vx.(float64)
				return func(fr *frame) { fr.setReg(ir, x*fr.reg(iy).(float64)) }
			} else if ky == kindConst {
				y := vy.(float64)
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(float64)*y) }
			} else {
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(float64)*fr.reg(iy).(float64)) }
			}
		case reflect.Complex64:
			if kx == kindConst && ky == kindConst {
				v := vx.(complex64) * vy.(complex64)
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := vx.(complex64)
				return func(fr *frame) { fr.setReg(ir, x*fr.reg(iy).(complex64)) }
			} else if ky == kindConst {
				y := vy.(complex64)
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(complex64)*y) }
			} else {
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(complex64)*fr.reg(iy).(complex64)) }
			}
		case reflect.Complex128:
			if kx == kindConst && ky == kindConst {
				v := vx.(complex128) * vy.(complex128)
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := vx.(complex128)
				return func(fr *frame) { fr.setReg(ir, x*fr.reg(iy).(complex128)) }
			} else if ky == kindConst {
				y := vy.(complex128)
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(complex128)*y) }
			} else {
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(complex128)*fr.reg(iy).(complex128)) }
			}
		}
	} else {
		t := basic.TypeOfType(typ)
		switch typ.Kind() {
		case reflect.Int:
			if kx == kindConst && ky == kindConst {
				v := basic.Make(t, basic.Int(vx)*basic.Int(vy))
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := basic.Int(vx)
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, x*fr.int(iy))) }
			} else if ky == kindConst {
				y := basic.Int(vy)
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, fr.int(ix)*y)) }
			} else {
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, fr.int(ix)*fr.int(iy))) }
			}
		case reflect.Int8:
			if kx == kindConst && ky == kindConst {
				v := basic.Make(t, basic.Int8(vx)*basic.Int8(vy))
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := basic.Int8(vx)
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, x*fr.int8(iy))) }
			} else if ky == kindConst {
				y := basic.Int8(vy)
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, fr.int8(ix)*y)) }
			} else {
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, fr.int8(ix)*fr.int8(iy))) }
			}
		case reflect.Int16:
			if kx == kindConst && ky == kindConst {
				v := basic.Make(t, basic.Int16(vx)*basic.Int16(vy))
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := basic.Int16(vx)
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, x*fr.int16(iy))) }
			} else if ky == kindConst {
				y := basic.Int16(vy)
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, fr.int16(ix)*y)) }
			} else {
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, fr.int16(ix)*fr.int16(iy))) }
			}
		case reflect.Int32:
			if kx == kindConst && ky == kindConst {
				v := basic.Make(t, basic.Int32(vx)*basic.Int32(vy))
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := basic.Int32(vx)
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, x*fr.int32(iy))) }
			} else if ky == kindConst {
				y := basic.Int32(vy)
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, fr.int32(ix)*y)) }
			} else {
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, fr.int32(ix)*fr.int32(iy))) }
			}
		case reflect.Int64:
			if kx == kindConst && ky == kindConst {
				v := basic.Make(t, basic.Int64(vx)*basic.Int64(vy))
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := basic.Int64(vx)
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, x*fr.int64(iy))) }
			} else if ky == kindConst {
				y := basic.Int64(vy)
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, fr.int64(ix)*y)) }
			} else {
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, fr.int64(ix)*fr.int64(iy))) }
			}
		case reflect.Uint:
			if kx == kindConst && ky == kindConst {
				v := basic.Make(t, basic.Uint(vx)*basic.Uint(vy))
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := basic.Uint(vx)
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, x*fr.uint(iy))) }
			} else if ky == kindConst {
				y := basic.Uint(vy)
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, fr.uint(ix)*y)) }
			} else {
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, fr.uint(ix)*fr.uint(iy))) }
			}
		case reflect.Uint8:
			if kx == kindConst && ky == kindConst {
				v := basic.Make(t, basic.Uint8(vx)*basic.Uint8(vy))
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := basic.Uint8(vx)
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, x*fr.uint8(iy))) }
			} else if ky == kindConst {
				y := basic.Uint8(vy)
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, fr.uint8(ix)*y)) }
			} else {
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, fr.uint8(ix)*fr.uint8(iy))) }
			}
		case reflect.Uint16:
			if kx == kindConst && ky == kindConst {
				v := basic.Make(t, basic.Uint16(vx)*basic.Uint16(vy))
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := basic.Uint16(vx)
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, x*fr.uint16(iy))) }
			} else if ky == kindConst {
				y := basic.Uint16(vy)
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, fr.uint16(ix)*y)) }
			} else {
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, fr.uint16(ix)*fr.uint16(iy))) }
			}
		case reflect.Uint32:
			if kx == kindConst && ky == kindConst {
				v := basic.Make(t, basic.Uint32(vx)*basic.Uint32(vy))
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := basic.Uint32(vx)
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, x*fr.uint32(iy))) }
			} else if ky == kindConst {
				y := basic.Uint32(vy)
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, fr.uint32(ix)*y)) }
			} else {
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, fr.uint32(ix)*fr.uint32(iy))) }
			}
		case reflect.Uint64:
			if kx == kindConst && ky == kindConst {
				v := basic.Make(t, basic.Uint64(vx)*basic.Uint64(vy))
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := basic.Uint64(vx)
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, x*fr.uint64(iy))) }
			} else if ky == kindConst {
				y := basic.Uint64(vy)
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, fr.uint64(ix)*y)) }
			} else {
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, fr.uint64(ix)*fr.uint64(iy))) }
			}
		case reflect.Uintptr:
			if kx == kindConst && ky == kindConst {
				v := basic.Make(t, basic.Uintptr(vx)*basic.Uintptr(vy))
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := basic.Uintptr(vx)
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, x*fr.uintptr(iy))) }
			} else if ky == kindConst {
				y := basic.Uintptr(vy)
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, fr.uintptr(ix)*y)) }
			} else {
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, fr.uintptr(ix)*fr.uintptr(iy))) }
			}
		case reflect.Float32:
			if kx == kindConst && ky == kindConst {
				v := basic.Make(t, basic.Float32(vx)*basic.Float32(vy))
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := basic.Float32(vx)
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, x*fr.float32(iy))) }
			} else if ky == kindConst {
				y := basic.Float32(vy)
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, fr.float32(ix)*y)) }
			} else {
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, fr.float32(ix)*fr.float32(iy))) }
			}
		case reflect.Float64:
			if kx == kindConst && ky == kindConst {
				v := basic.Make(t, basic.Float64(vx)*basic.Float64(vy))
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := basic.Float64(vx)
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, x*fr.float64(iy))) }
			} else if ky == kindConst {
				y := basic.Float64(vy)
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, fr.float64(ix)*y)) }
			} else {
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, fr.float64(ix)*fr.float64(iy))) }
			}
		case reflect.Complex64:
			if kx == kindConst && ky == kindConst {
				v := basic.Make(t, basic.Complex64(vx)*basic.Complex64(vy))
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := basic.Complex64(vx)
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, x*fr.complex64(iy))) }
			} else if ky == kindConst {
				y := basic.Complex64(vy)
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, fr.complex64(ix)*y)) }
			} else {
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, fr.complex64(ix)*fr.complex64(iy))) }
			}
		case reflect.Complex128:
			if kx == kindConst && ky == kindConst {
				v := basic.Make(t, basic.Complex128(vx)*basic.Complex128(vy))
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := basic.Complex128(vx)
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, x*fr.complex128(iy))) }
			} else if ky == kindConst {
				y := basic.Complex128(vy)
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, fr.complex128(ix)*y)) }
			} else {
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, fr.complex128(ix)*fr.complex128(iy))) }
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
			if kx == kindConst && ky == kindConst {
				x := vx.(int)
				y := vy.(int)
				if y == 0 {
					return func(fr *frame) { fr.setReg(ir, x/y) }
				}
				v := x / y
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := vx.(int)
				return func(fr *frame) { fr.setReg(ir, x/fr.reg(iy).(int)) }
			} else if ky == kindConst {
				y := vy.(int)
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int)/y) }
			} else {
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int)/fr.reg(iy).(int)) }
			}
		case reflect.Int8:
			if kx == kindConst && ky == kindConst {
				x := vx.(int8)
				y := vy.(int8)
				if y == 0 {
					return func(fr *frame) { fr.setReg(ir, x/y) }
				}
				v := x / y
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := vx.(int8)
				return func(fr *frame) { fr.setReg(ir, x/fr.reg(iy).(int8)) }
			} else if ky == kindConst {
				y := vy.(int8)
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int8)/y) }
			} else {
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int8)/fr.reg(iy).(int8)) }
			}
		case reflect.Int16:
			if kx == kindConst && ky == kindConst {
				x := vx.(int16)
				y := vy.(int16)
				if y == 0 {
					return func(fr *frame) { fr.setReg(ir, x/y) }
				}
				v := x / y
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := vx.(int16)
				return func(fr *frame) { fr.setReg(ir, x/fr.reg(iy).(int16)) }
			} else if ky == kindConst {
				y := vy.(int16)
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int16)/y) }
			} else {
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int16)/fr.reg(iy).(int16)) }
			}
		case reflect.Int32:
			if kx == kindConst && ky == kindConst {
				x := vx.(int32)
				y := vy.(int32)
				if y == 0 {
					return func(fr *frame) { fr.setReg(ir, x/y) }
				}
				v := x / y
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := vx.(int32)
				return func(fr *frame) { fr.setReg(ir, x/fr.reg(iy).(int32)) }
			} else if ky == kindConst {
				y := vy.(int32)
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int32)/y) }
			} else {
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int32)/fr.reg(iy).(int32)) }
			}
		case reflect.Int64:
			if kx == kindConst && ky == kindConst {
				x := vx.(int64)
				y := vy.(int64)
				if y == 0 {
					return func(fr *frame) { fr.setReg(ir, x/y) }
				}
				v := x / y
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := vx.(int64)
				return func(fr *frame) { fr.setReg(ir, x/fr.reg(iy).(int64)) }
			} else if ky == kindConst {
				y := vy.(int64)
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int64)/y) }
			} else {
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int64)/fr.reg(iy).(int64)) }
			}
		case reflect.Uint:
			if kx == kindConst && ky == kindConst {
				x := vx.(uint)
				y := vy.(uint)
				if y == 0 {
					return func(fr *frame) { fr.setReg(ir, x/y) }
				}
				v := x / y
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := vx.(uint)
				return func(fr *frame) { fr.setReg(ir, x/fr.reg(iy).(uint)) }
			} else if ky == kindConst {
				y := vy.(uint)
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint)/y) }
			} else {
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint)/fr.reg(iy).(uint)) }
			}
		case reflect.Uint8:
			if kx == kindConst && ky == kindConst {
				x := vx.(uint8)
				y := vy.(uint8)
				if y == 0 {
					return func(fr *frame) { fr.setReg(ir, x/y) }
				}
				v := x / y
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := vx.(uint8)
				return func(fr *frame) { fr.setReg(ir, x/fr.reg(iy).(uint8)) }
			} else if ky == kindConst {
				y := vy.(uint8)
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint8)/y) }
			} else {
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint8)/fr.reg(iy).(uint8)) }
			}
		case reflect.Uint16:
			if kx == kindConst && ky == kindConst {
				x := vx.(uint16)
				y := vy.(uint16)
				if y == 0 {
					return func(fr *frame) { fr.setReg(ir, x/y) }
				}
				v := x / y
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := vx.(uint16)
				return func(fr *frame) { fr.setReg(ir, x/fr.reg(iy).(uint16)) }
			} else if ky == kindConst {
				y := vy.(uint16)
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint16)/y) }
			} else {
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint16)/fr.reg(iy).(uint16)) }
			}
		case reflect.Uint32:
			if kx == kindConst && ky == kindConst {
				x := vx.(uint32)
				y := vy.(uint32)
				if y == 0 {
					return func(fr *frame) { fr.setReg(ir, x/y) }
				}
				v := x / y
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := vx.(uint32)
				return func(fr *frame) { fr.setReg(ir, x/fr.reg(iy).(uint32)) }
			} else if ky == kindConst {
				y := vy.(uint32)
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint32)/y) }
			} else {
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint32)/fr.reg(iy).(uint32)) }
			}
		case reflect.Uint64:
			if kx == kindConst && ky == kindConst {
				x := vx.(uint64)
				y := vy.(uint64)
				if y == 0 {
					return func(fr *frame) { fr.setReg(ir, x/y) }
				}
				v := x / y
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := vx.(uint64)
				return func(fr *frame) { fr.setReg(ir, x/fr.reg(iy).(uint64)) }
			} else if ky == kindConst {
				y := vy.(uint64)
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint64)/y) }
			} else {
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint64)/fr.reg(iy).(uint64)) }
			}
		case reflect.Uintptr:
			if kx == kindConst && ky == kindConst {
				x := vx.(uintptr)
				y := vy.(uintptr)
				if y == 0 {
					return func(fr *frame) { fr.setReg(ir, x/y) }
				}
				v := x / y
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := vx.(uintptr)
				return func(fr *frame) { fr.setReg(ir, x/fr.reg(iy).(uintptr)) }
			} else if ky == kindConst {
				y := vy.(uintptr)
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uintptr)/y) }
			} else {
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uintptr)/fr.reg(iy).(uintptr)) }
			}
		case reflect.Float32:
			if kx == kindConst && ky == kindConst {
				v := vx.(float32) / vy.(float32)
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := vx.(float32)
				return func(fr *frame) { fr.setReg(ir, x/fr.reg(iy).(float32)) }
			} else if ky == kindConst {
				y := vy.(float32)
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(float32)/y) }
			} else {
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(float32)/fr.reg(iy).(float32)) }
			}
		case reflect.Float64:
			if kx == kindConst && ky == kindConst {
				v := vx.(float64) / vy.(float64)
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := vx.(float64)
				return func(fr *frame) { fr.setReg(ir, x/fr.reg(iy).(float64)) }
			} else if ky == kindConst {
				y := vy.(float64)
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(float64)/y) }
			} else {
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(float64)/fr.reg(iy).(float64)) }
			}
		case reflect.Complex64:
			if kx == kindConst && ky == kindConst {
				v := vx.(complex64) / vy.(complex64)
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := vx.(complex64)
				return func(fr *frame) { fr.setReg(ir, x/fr.reg(iy).(complex64)) }
			} else if ky == kindConst {
				y := vy.(complex64)
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(complex64)/y) }
			} else {
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(complex64)/fr.reg(iy).(complex64)) }
			}
		case reflect.Complex128:
			if kx == kindConst && ky == kindConst {
				v := vx.(complex128) / vy.(complex128)
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := vx.(complex128)
				return func(fr *frame) { fr.setReg(ir, x/fr.reg(iy).(complex128)) }
			} else if ky == kindConst {
				y := vy.(complex128)
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(complex128)/y) }
			} else {
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(complex128)/fr.reg(iy).(complex128)) }
			}
		}
	} else {
		t := basic.TypeOfType(typ)
		switch typ.Kind() {
		case reflect.Int:
			if kx == kindConst && ky == kindConst {
				x := basic.Int(vx)
				y := basic.Int(vy)
				if y == 0 {
					return func(fr *frame) { fr.setReg(ir, basic.Make(t, x/y)) }
				}
				v := basic.Make(t, x/y)
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := basic.Int(vx)
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, x/fr.int(iy))) }
			} else if ky == kindConst {
				y := basic.Int(vy)
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, fr.int(ix)/y)) }
			} else {
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, fr.int(ix)/fr.int(iy))) }
			}
		case reflect.Int8:
			if kx == kindConst && ky == kindConst {
				x := basic.Int8(vx)
				y := basic.Int8(vy)
				if y == 0 {
					return func(fr *frame) { fr.setReg(ir, basic.Make(t, x/y)) }
				}
				v := basic.Make(t, x/y)
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := basic.Int8(vx)
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, x/fr.int8(iy))) }
			} else if ky == kindConst {
				y := basic.Int8(vy)
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, fr.int8(ix)/y)) }
			} else {
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, fr.int8(ix)/fr.int8(iy))) }
			}
		case reflect.Int16:
			if kx == kindConst && ky == kindConst {
				x := basic.Int16(vx)
				y := basic.Int16(vy)
				if y == 0 {
					return func(fr *frame) { fr.setReg(ir, basic.Make(t, x/y)) }
				}
				v := basic.Make(t, x/y)
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := basic.Int16(vx)
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, x/fr.int16(iy))) }
			} else if ky == kindConst {
				y := basic.Int16(vy)
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, fr.int16(ix)/y)) }
			} else {
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, fr.int16(ix)/fr.int16(iy))) }
			}
		case reflect.Int32:
			if kx == kindConst && ky == kindConst {
				x := basic.Int32(vx)
				y := basic.Int32(vy)
				if y == 0 {
					return func(fr *frame) { fr.setReg(ir, basic.Make(t, x/y)) }
				}
				v := basic.Make(t, x/y)
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := basic.Int32(vx)
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, x/fr.int32(iy))) }
			} else if ky == kindConst {
				y := basic.Int32(vy)
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, fr.int32(ix)/y)) }
			} else {
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, fr.int32(ix)/fr.int32(iy))) }
			}
		case reflect.Int64:
			if kx == kindConst && ky == kindConst {
				x := basic.Int64(vx)
				y := basic.Int64(vy)
				if y == 0 {
					return func(fr *frame) { fr.setReg(ir, basic.Make(t, x/y)) }
				}
				v := basic.Make(t, x/y)
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := basic.Int64(vx)
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, x/fr.int64(iy))) }
			} else if ky == kindConst {
				y := basic.Int64(vy)
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, fr.int64(ix)/y)) }
			} else {
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, fr.int64(ix)/fr.int64(iy))) }
			}
		case reflect.Uint:
			if kx == kindConst && ky == kindConst {
				x := basic.Uint(vx)
				y := basic.Uint(vy)
				if y == 0 {
					return func(fr *frame) { fr.setReg(ir, basic.Make(t, x/y)) }
				}
				v := basic.Make(t, x/y)
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := basic.Uint(vx)
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, x/fr.uint(iy))) }
			} else if ky == kindConst {
				y := basic.Uint(vy)
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, fr.uint(ix)/y)) }
			} else {
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, fr.uint(ix)/fr.uint(iy))) }
			}
		case reflect.Uint8:
			if kx == kindConst && ky == kindConst {
				x := basic.Uint8(vx)
				y := basic.Uint8(vy)
				if y == 0 {
					return func(fr *frame) { fr.setReg(ir, basic.Make(t, x/y)) }
				}
				v := basic.Make(t, x/y)
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := basic.Uint8(vx)
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, x/fr.uint8(iy))) }
			} else if ky == kindConst {
				y := basic.Uint8(vy)
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, fr.uint8(ix)/y)) }
			} else {
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, fr.uint8(ix)/fr.uint8(iy))) }
			}
		case reflect.Uint16:
			if kx == kindConst && ky == kindConst {
				x := basic.Uint16(vx)
				y := basic.Uint16(vy)
				if y == 0 {
					return func(fr *frame) { fr.setReg(ir, basic.Make(t, x/y)) }
				}
				v := basic.Make(t, x/y)
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := basic.Uint16(vx)
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, x/fr.uint16(iy))) }
			} else if ky == kindConst {
				y := basic.Uint16(vy)
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, fr.uint16(ix)/y)) }
			} else {
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, fr.uint16(ix)/fr.uint16(iy))) }
			}
		case reflect.Uint32:
			if kx == kindConst && ky == kindConst {
				x := basic.Uint32(vx)
				y := basic.Uint32(vy)
				if y == 0 {
					return func(fr *frame) { fr.setReg(ir, basic.Make(t, x/y)) }
				}
				v := basic.Make(t, x/y)
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := basic.Uint32(vx)
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, x/fr.uint32(iy))) }
			} else if ky == kindConst {
				y := basic.Uint32(vy)
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, fr.uint32(ix)/y)) }
			} else {
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, fr.uint32(ix)/fr.uint32(iy))) }
			}
		case reflect.Uint64:
			if kx == kindConst && ky == kindConst {
				x := basic.Uint64(vx)
				y := basic.Uint64(vy)
				if y == 0 {
					return func(fr *frame) { fr.setReg(ir, basic.Make(t, x/y)) }
				}
				v := basic.Make(t, x/y)
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := basic.Uint64(vx)
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, x/fr.uint64(iy))) }
			} else if ky == kindConst {
				y := basic.Uint64(vy)
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, fr.uint64(ix)/y)) }
			} else {
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, fr.uint64(ix)/fr.uint64(iy))) }
			}
		case reflect.Uintptr:
			if kx == kindConst && ky == kindConst {
				x := basic.Uintptr(vx)
				y := basic.Uintptr(vy)
				if y == 0 {
					return func(fr *frame) { fr.setReg(ir, basic.Make(t, x/y)) }
				}
				v := basic.Make(t, x/y)
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := basic.Uintptr(vx)
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, x/fr.uintptr(iy))) }
			} else if ky == kindConst {
				y := basic.Uintptr(vy)
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, fr.uintptr(ix)/y)) }
			} else {
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, fr.uintptr(ix)/fr.uintptr(iy))) }
			}
		case reflect.Float32:
			if kx == kindConst && ky == kindConst {
				v := basic.Make(t, basic.Float32(vx)/basic.Float32(vy))
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := basic.Float32(vx)
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, x/fr.float32(iy))) }
			} else if ky == kindConst {
				y := basic.Float32(vy)
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, fr.float32(ix)/y)) }
			} else {
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, fr.float32(ix)/fr.float32(iy))) }
			}
		case reflect.Float64:
			if kx == kindConst && ky == kindConst {
				v := basic.Make(t, basic.Float64(vx)/basic.Float64(vy))
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := basic.Float64(vx)
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, x/fr.float64(iy))) }
			} else if ky == kindConst {
				y := basic.Float64(vy)
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, fr.float64(ix)/y)) }
			} else {
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, fr.float64(ix)/fr.float64(iy))) }
			}
		case reflect.Complex64:
			if kx == kindConst && ky == kindConst {
				v := basic.Make(t, basic.Complex64(vx)/basic.Complex64(vy))
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := basic.Complex64(vx)
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, x/fr.complex64(iy))) }
			} else if ky == kindConst {
				y := basic.Complex64(vy)
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, fr.complex64(ix)/y)) }
			} else {
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, fr.complex64(ix)/fr.complex64(iy))) }
			}
		case reflect.Complex128:
			if kx == kindConst && ky == kindConst {
				v := basic.Make(t, basic.Complex128(vx)/basic.Complex128(vy))
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := basic.Complex128(vx)
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, x/fr.complex128(iy))) }
			} else if ky == kindConst {
				y := basic.Complex128(vy)
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, fr.complex128(ix)/y)) }
			} else {
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, fr.complex128(ix)/fr.complex128(iy))) }
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
			if kx == kindConst && ky == kindConst {
				v := vx.(int) % vy.(int)
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := vx.(int)
				return func(fr *frame) { fr.setReg(ir, x%fr.reg(iy).(int)) }
			} else if ky == kindConst {
				y := vy.(int)
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int)%y) }
			} else {
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int)%fr.reg(iy).(int)) }
			}
		case reflect.Int8:
			if kx == kindConst && ky == kindConst {
				v := vx.(int8) % vy.(int8)
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := vx.(int8)
				return func(fr *frame) { fr.setReg(ir, x%fr.reg(iy).(int8)) }
			} else if ky == kindConst {
				y := vy.(int8)
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int8)%y) }
			} else {
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int8)%fr.reg(iy).(int8)) }
			}
		case reflect.Int16:
			if kx == kindConst && ky == kindConst {
				v := vx.(int16) % vy.(int16)
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := vx.(int16)
				return func(fr *frame) { fr.setReg(ir, x%fr.reg(iy).(int16)) }
			} else if ky == kindConst {
				y := vy.(int16)
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int16)%y) }
			} else {
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int16)%fr.reg(iy).(int16)) }
			}
		case reflect.Int32:
			if kx == kindConst && ky == kindConst {
				v := vx.(int32) % vy.(int32)
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := vx.(int32)
				return func(fr *frame) { fr.setReg(ir, x%fr.reg(iy).(int32)) }
			} else if ky == kindConst {
				y := vy.(int32)
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int32)%y) }
			} else {
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int32)%fr.reg(iy).(int32)) }
			}
		case reflect.Int64:
			if kx == kindConst && ky == kindConst {
				v := vx.(int64) % vy.(int64)
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := vx.(int64)
				return func(fr *frame) { fr.setReg(ir, x%fr.reg(iy).(int64)) }
			} else if ky == kindConst {
				y := vy.(int64)
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int64)%y) }
			} else {
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int64)%fr.reg(iy).(int64)) }
			}
		case reflect.Uint:
			if kx == kindConst && ky == kindConst {
				v := vx.(uint) % vy.(uint)
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := vx.(uint)
				return func(fr *frame) { fr.setReg(ir, x%fr.reg(iy).(uint)) }
			} else if ky == kindConst {
				y := vy.(uint)
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint)%y) }
			} else {
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint)%fr.reg(iy).(uint)) }
			}
		case reflect.Uint8:
			if kx == kindConst && ky == kindConst {
				v := vx.(uint8) % vy.(uint8)
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := vx.(uint8)
				return func(fr *frame) { fr.setReg(ir, x%fr.reg(iy).(uint8)) }
			} else if ky == kindConst {
				y := vy.(uint8)
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint8)%y) }
			} else {
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint8)%fr.reg(iy).(uint8)) }
			}
		case reflect.Uint16:
			if kx == kindConst && ky == kindConst {
				v := vx.(uint16) % vy.(uint16)
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := vx.(uint16)
				return func(fr *frame) { fr.setReg(ir, x%fr.reg(iy).(uint16)) }
			} else if ky == kindConst {
				y := vy.(uint16)
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint16)%y) }
			} else {
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint16)%fr.reg(iy).(uint16)) }
			}
		case reflect.Uint32:
			if kx == kindConst && ky == kindConst {
				v := vx.(uint32) % vy.(uint32)
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := vx.(uint32)
				return func(fr *frame) { fr.setReg(ir, x%fr.reg(iy).(uint32)) }
			} else if ky == kindConst {
				y := vy.(uint32)
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint32)%y) }
			} else {
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint32)%fr.reg(iy).(uint32)) }
			}
		case reflect.Uint64:
			if kx == kindConst && ky == kindConst {
				v := vx.(uint64) % vy.(uint64)
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := vx.(uint64)
				return func(fr *frame) { fr.setReg(ir, x%fr.reg(iy).(uint64)) }
			} else if ky == kindConst {
				y := vy.(uint64)
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint64)%y) }
			} else {
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint64)%fr.reg(iy).(uint64)) }
			}
		case reflect.Uintptr:
			if kx == kindConst && ky == kindConst {
				v := vx.(uintptr) % vy.(uintptr)
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := vx.(uintptr)
				return func(fr *frame) { fr.setReg(ir, x%fr.reg(iy).(uintptr)) }
			} else if ky == kindConst {
				y := vy.(uintptr)
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uintptr)%y) }
			} else {
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uintptr)%fr.reg(iy).(uintptr)) }
			}
		}
	} else {
		t := basic.TypeOfType(typ)
		switch typ.Kind() {
		case reflect.Int:
			if kx == kindConst && ky == kindConst {
				v := basic.Make(t, basic.Int(vx)%basic.Int(vy))
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := basic.Int(vx)
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, x%fr.int(iy))) }
			} else if ky == kindConst {
				y := basic.Int(vy)
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, fr.int(ix)%y)) }
			} else {
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, fr.int(ix)%fr.int(iy))) }
			}
		case reflect.Int8:
			if kx == kindConst && ky == kindConst {
				v := basic.Make(t, basic.Int8(vx)%basic.Int8(vy))
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := basic.Int8(vx)
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, x%fr.int8(iy))) }
			} else if ky == kindConst {
				y := basic.Int8(vy)
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, fr.int8(ix)%y)) }
			} else {
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, fr.int8(ix)%fr.int8(iy))) }
			}
		case reflect.Int16:
			if kx == kindConst && ky == kindConst {
				v := basic.Make(t, basic.Int16(vx)%basic.Int16(vy))
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := basic.Int16(vx)
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, x%fr.int16(iy))) }
			} else if ky == kindConst {
				y := basic.Int16(vy)
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, fr.int16(ix)%y)) }
			} else {
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, fr.int16(ix)%fr.int16(iy))) }
			}
		case reflect.Int32:
			if kx == kindConst && ky == kindConst {
				v := basic.Make(t, basic.Int32(vx)%basic.Int32(vy))
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := basic.Int32(vx)
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, x%fr.int32(iy))) }
			} else if ky == kindConst {
				y := basic.Int32(vy)
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, fr.int32(ix)%y)) }
			} else {
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, fr.int32(ix)%fr.int32(iy))) }
			}
		case reflect.Int64:
			if kx == kindConst && ky == kindConst {
				v := basic.Make(t, basic.Int64(vx)%basic.Int64(vy))
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := basic.Int64(vx)
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, x%fr.int64(iy))) }
			} else if ky == kindConst {
				y := basic.Int64(vy)
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, fr.int64(ix)%y)) }
			} else {
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, fr.int64(ix)%fr.int64(iy))) }
			}
		case reflect.Uint:
			if kx == kindConst && ky == kindConst {
				v := basic.Make(t, basic.Uint(vx)%basic.Uint(vy))
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := basic.Uint(vx)
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, x%fr.uint(iy))) }
			} else if ky == kindConst {
				y := basic.Uint(vy)
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, fr.uint(ix)%y)) }
			} else {
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, fr.uint(ix)%fr.uint(iy))) }
			}
		case reflect.Uint8:
			if kx == kindConst && ky == kindConst {
				v := basic.Make(t, basic.Uint8(vx)%basic.Uint8(vy))
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := basic.Uint8(vx)
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, x%fr.uint8(iy))) }
			} else if ky == kindConst {
				y := basic.Uint8(vy)
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, fr.uint8(ix)%y)) }
			} else {
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, fr.uint8(ix)%fr.uint8(iy))) }
			}
		case reflect.Uint16:
			if kx == kindConst && ky == kindConst {
				v := basic.Make(t, basic.Uint16(vx)%basic.Uint16(vy))
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := basic.Uint16(vx)
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, x%fr.uint16(iy))) }
			} else if ky == kindConst {
				y := basic.Uint16(vy)
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, fr.uint16(ix)%y)) }
			} else {
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, fr.uint16(ix)%fr.uint16(iy))) }
			}
		case reflect.Uint32:
			if kx == kindConst && ky == kindConst {
				v := basic.Make(t, basic.Uint32(vx)%basic.Uint32(vy))
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := basic.Uint32(vx)
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, x%fr.uint32(iy))) }
			} else if ky == kindConst {
				y := basic.Uint32(vy)
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, fr.uint32(ix)%y)) }
			} else {
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, fr.uint32(ix)%fr.uint32(iy))) }
			}
		case reflect.Uint64:
			if kx == kindConst && ky == kindConst {
				v := basic.Make(t, basic.Uint64(vx)%basic.Uint64(vy))
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := basic.Uint64(vx)
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, x%fr.uint64(iy))) }
			} else if ky == kindConst {
				y := basic.Uint64(vy)
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, fr.uint64(ix)%y)) }
			} else {
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, fr.uint64(ix)%fr.uint64(iy))) }
			}
		case reflect.Uintptr:
			if kx == kindConst && ky == kindConst {
				v := basic.Make(t, basic.Uintptr(vx)%basic.Uintptr(vy))
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := basic.Uintptr(vx)
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, x%fr.uintptr(iy))) }
			} else if ky == kindConst {
				y := basic.Uintptr(vy)
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, fr.uintptr(ix)%y)) }
			} else {
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, fr.uintptr(ix)%fr.uintptr(iy))) }
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
			if kx == kindConst && ky == kindConst {
				v := vx.(int) & vy.(int)
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := vx.(int)
				return func(fr *frame) { fr.setReg(ir, x&fr.reg(iy).(int)) }
			} else if ky == kindConst {
				y := vy.(int)
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int)&y) }
			} else {
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int)&fr.reg(iy).(int)) }
			}
		case reflect.Int8:
			if kx == kindConst && ky == kindConst {
				v := vx.(int8) & vy.(int8)
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := vx.(int8)
				return func(fr *frame) { fr.setReg(ir, x&fr.reg(iy).(int8)) }
			} else if ky == kindConst {
				y := vy.(int8)
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int8)&y) }
			} else {
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int8)&fr.reg(iy).(int8)) }
			}
		case reflect.Int16:
			if kx == kindConst && ky == kindConst {
				v := vx.(int16) & vy.(int16)
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := vx.(int16)
				return func(fr *frame) { fr.setReg(ir, x&fr.reg(iy).(int16)) }
			} else if ky == kindConst {
				y := vy.(int16)
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int16)&y) }
			} else {
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int16)&fr.reg(iy).(int16)) }
			}
		case reflect.Int32:
			if kx == kindConst && ky == kindConst {
				v := vx.(int32) & vy.(int32)
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := vx.(int32)
				return func(fr *frame) { fr.setReg(ir, x&fr.reg(iy).(int32)) }
			} else if ky == kindConst {
				y := vy.(int32)
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int32)&y) }
			} else {
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int32)&fr.reg(iy).(int32)) }
			}
		case reflect.Int64:
			if kx == kindConst && ky == kindConst {
				v := vx.(int64) & vy.(int64)
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := vx.(int64)
				return func(fr *frame) { fr.setReg(ir, x&fr.reg(iy).(int64)) }
			} else if ky == kindConst {
				y := vy.(int64)
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int64)&y) }
			} else {
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int64)&fr.reg(iy).(int64)) }
			}
		case reflect.Uint:
			if kx == kindConst && ky == kindConst {
				v := vx.(uint) & vy.(uint)
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := vx.(uint)
				return func(fr *frame) { fr.setReg(ir, x&fr.reg(iy).(uint)) }
			} else if ky == kindConst {
				y := vy.(uint)
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint)&y) }
			} else {
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint)&fr.reg(iy).(uint)) }
			}
		case reflect.Uint8:
			if kx == kindConst && ky == kindConst {
				v := vx.(uint8) & vy.(uint8)
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := vx.(uint8)
				return func(fr *frame) { fr.setReg(ir, x&fr.reg(iy).(uint8)) }
			} else if ky == kindConst {
				y := vy.(uint8)
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint8)&y) }
			} else {
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint8)&fr.reg(iy).(uint8)) }
			}
		case reflect.Uint16:
			if kx == kindConst && ky == kindConst {
				v := vx.(uint16) & vy.(uint16)
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := vx.(uint16)
				return func(fr *frame) { fr.setReg(ir, x&fr.reg(iy).(uint16)) }
			} else if ky == kindConst {
				y := vy.(uint16)
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint16)&y) }
			} else {
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint16)&fr.reg(iy).(uint16)) }
			}
		case reflect.Uint32:
			if kx == kindConst && ky == kindConst {
				v := vx.(uint32) & vy.(uint32)
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := vx.(uint32)
				return func(fr *frame) { fr.setReg(ir, x&fr.reg(iy).(uint32)) }
			} else if ky == kindConst {
				y := vy.(uint32)
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint32)&y) }
			} else {
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint32)&fr.reg(iy).(uint32)) }
			}
		case reflect.Uint64:
			if kx == kindConst && ky == kindConst {
				v := vx.(uint64) & vy.(uint64)
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := vx.(uint64)
				return func(fr *frame) { fr.setReg(ir, x&fr.reg(iy).(uint64)) }
			} else if ky == kindConst {
				y := vy.(uint64)
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint64)&y) }
			} else {
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint64)&fr.reg(iy).(uint64)) }
			}
		case reflect.Uintptr:
			if kx == kindConst && ky == kindConst {
				v := vx.(uintptr) & vy.(uintptr)
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := vx.(uintptr)
				return func(fr *frame) { fr.setReg(ir, x&fr.reg(iy).(uintptr)) }
			} else if ky == kindConst {
				y := vy.(uintptr)
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uintptr)&y) }
			} else {
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uintptr)&fr.reg(iy).(uintptr)) }
			}
		}
	} else {
		t := basic.TypeOfType(typ)
		switch typ.Kind() {
		case reflect.Int:
			if kx == kindConst && ky == kindConst {
				v := basic.Make(t, basic.Int(vx)&basic.Int(vy))
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := basic.Int(vx)
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, x&fr.int(iy))) }
			} else if ky == kindConst {
				y := basic.Int(vy)
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, fr.int(ix)&y)) }
			} else {
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, fr.int(ix)&fr.int(iy))) }
			}
		case reflect.Int8:
			if kx == kindConst && ky == kindConst {
				v := basic.Make(t, basic.Int8(vx)&basic.Int8(vy))
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := basic.Int8(vx)
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, x&fr.int8(iy))) }
			} else if ky == kindConst {
				y := basic.Int8(vy)
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, fr.int8(ix)&y)) }
			} else {
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, fr.int8(ix)&fr.int8(iy))) }
			}
		case reflect.Int16:
			if kx == kindConst && ky == kindConst {
				v := basic.Make(t, basic.Int16(vx)&basic.Int16(vy))
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := basic.Int16(vx)
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, x&fr.int16(iy))) }
			} else if ky == kindConst {
				y := basic.Int16(vy)
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, fr.int16(ix)&y)) }
			} else {
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, fr.int16(ix)&fr.int16(iy))) }
			}
		case reflect.Int32:
			if kx == kindConst && ky == kindConst {
				v := basic.Make(t, basic.Int32(vx)&basic.Int32(vy))
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := basic.Int32(vx)
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, x&fr.int32(iy))) }
			} else if ky == kindConst {
				y := basic.Int32(vy)
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, fr.int32(ix)&y)) }
			} else {
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, fr.int32(ix)&fr.int32(iy))) }
			}
		case reflect.Int64:
			if kx == kindConst && ky == kindConst {
				v := basic.Make(t, basic.Int64(vx)&basic.Int64(vy))
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := basic.Int64(vx)
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, x&fr.int64(iy))) }
			} else if ky == kindConst {
				y := basic.Int64(vy)
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, fr.int64(ix)&y)) }
			} else {
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, fr.int64(ix)&fr.int64(iy))) }
			}
		case reflect.Uint:
			if kx == kindConst && ky == kindConst {
				v := basic.Make(t, basic.Uint(vx)&basic.Uint(vy))
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := basic.Uint(vx)
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, x&fr.uint(iy))) }
			} else if ky == kindConst {
				y := basic.Uint(vy)
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, fr.uint(ix)&y)) }
			} else {
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, fr.uint(ix)&fr.uint(iy))) }
			}
		case reflect.Uint8:
			if kx == kindConst && ky == kindConst {
				v := basic.Make(t, basic.Uint8(vx)&basic.Uint8(vy))
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := basic.Uint8(vx)
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, x&fr.uint8(iy))) }
			} else if ky == kindConst {
				y := basic.Uint8(vy)
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, fr.uint8(ix)&y)) }
			} else {
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, fr.uint8(ix)&fr.uint8(iy))) }
			}
		case reflect.Uint16:
			if kx == kindConst && ky == kindConst {
				v := basic.Make(t, basic.Uint16(vx)&basic.Uint16(vy))
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := basic.Uint16(vx)
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, x&fr.uint16(iy))) }
			} else if ky == kindConst {
				y := basic.Uint16(vy)
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, fr.uint16(ix)&y)) }
			} else {
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, fr.uint16(ix)&fr.uint16(iy))) }
			}
		case reflect.Uint32:
			if kx == kindConst && ky == kindConst {
				v := basic.Make(t, basic.Uint32(vx)&basic.Uint32(vy))
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := basic.Uint32(vx)
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, x&fr.uint32(iy))) }
			} else if ky == kindConst {
				y := basic.Uint32(vy)
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, fr.uint32(ix)&y)) }
			} else {
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, fr.uint32(ix)&fr.uint32(iy))) }
			}
		case reflect.Uint64:
			if kx == kindConst && ky == kindConst {
				v := basic.Make(t, basic.Uint64(vx)&basic.Uint64(vy))
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := basic.Uint64(vx)
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, x&fr.uint64(iy))) }
			} else if ky == kindConst {
				y := basic.Uint64(vy)
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, fr.uint64(ix)&y)) }
			} else {
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, fr.uint64(ix)&fr.uint64(iy))) }
			}
		case reflect.Uintptr:
			if kx == kindConst && ky == kindConst {
				v := basic.Make(t, basic.Uintptr(vx)&basic.Uintptr(vy))
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := basic.Uintptr(vx)
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, x&fr.uintptr(iy))) }
			} else if ky == kindConst {
				y := basic.Uintptr(vy)
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, fr.uintptr(ix)&y)) }
			} else {
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, fr.uintptr(ix)&fr.uintptr(iy))) }
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
			if kx == kindConst && ky == kindConst {
				v := vx.(int) | vy.(int)
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := vx.(int)
				return func(fr *frame) { fr.setReg(ir, x|fr.reg(iy).(int)) }
			} else if ky == kindConst {
				y := vy.(int)
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int)|y) }
			} else {
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int)|fr.reg(iy).(int)) }
			}
		case reflect.Int8:
			if kx == kindConst && ky == kindConst {
				v := vx.(int8) | vy.(int8)
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := vx.(int8)
				return func(fr *frame) { fr.setReg(ir, x|fr.reg(iy).(int8)) }
			} else if ky == kindConst {
				y := vy.(int8)
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int8)|y) }
			} else {
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int8)|fr.reg(iy).(int8)) }
			}
		case reflect.Int16:
			if kx == kindConst && ky == kindConst {
				v := vx.(int16) | vy.(int16)
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := vx.(int16)
				return func(fr *frame) { fr.setReg(ir, x|fr.reg(iy).(int16)) }
			} else if ky == kindConst {
				y := vy.(int16)
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int16)|y) }
			} else {
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int16)|fr.reg(iy).(int16)) }
			}
		case reflect.Int32:
			if kx == kindConst && ky == kindConst {
				v := vx.(int32) | vy.(int32)
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := vx.(int32)
				return func(fr *frame) { fr.setReg(ir, x|fr.reg(iy).(int32)) }
			} else if ky == kindConst {
				y := vy.(int32)
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int32)|y) }
			} else {
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int32)|fr.reg(iy).(int32)) }
			}
		case reflect.Int64:
			if kx == kindConst && ky == kindConst {
				v := vx.(int64) | vy.(int64)
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := vx.(int64)
				return func(fr *frame) { fr.setReg(ir, x|fr.reg(iy).(int64)) }
			} else if ky == kindConst {
				y := vy.(int64)
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int64)|y) }
			} else {
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int64)|fr.reg(iy).(int64)) }
			}
		case reflect.Uint:
			if kx == kindConst && ky == kindConst {
				v := vx.(uint) | vy.(uint)
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := vx.(uint)
				return func(fr *frame) { fr.setReg(ir, x|fr.reg(iy).(uint)) }
			} else if ky == kindConst {
				y := vy.(uint)
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint)|y) }
			} else {
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint)|fr.reg(iy).(uint)) }
			}
		case reflect.Uint8:
			if kx == kindConst && ky == kindConst {
				v := vx.(uint8) | vy.(uint8)
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := vx.(uint8)
				return func(fr *frame) { fr.setReg(ir, x|fr.reg(iy).(uint8)) }
			} else if ky == kindConst {
				y := vy.(uint8)
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint8)|y) }
			} else {
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint8)|fr.reg(iy).(uint8)) }
			}
		case reflect.Uint16:
			if kx == kindConst && ky == kindConst {
				v := vx.(uint16) | vy.(uint16)
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := vx.(uint16)
				return func(fr *frame) { fr.setReg(ir, x|fr.reg(iy).(uint16)) }
			} else if ky == kindConst {
				y := vy.(uint16)
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint16)|y) }
			} else {
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint16)|fr.reg(iy).(uint16)) }
			}
		case reflect.Uint32:
			if kx == kindConst && ky == kindConst {
				v := vx.(uint32) | vy.(uint32)
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := vx.(uint32)
				return func(fr *frame) { fr.setReg(ir, x|fr.reg(iy).(uint32)) }
			} else if ky == kindConst {
				y := vy.(uint32)
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint32)|y) }
			} else {
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint32)|fr.reg(iy).(uint32)) }
			}
		case reflect.Uint64:
			if kx == kindConst && ky == kindConst {
				v := vx.(uint64) | vy.(uint64)
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := vx.(uint64)
				return func(fr *frame) { fr.setReg(ir, x|fr.reg(iy).(uint64)) }
			} else if ky == kindConst {
				y := vy.(uint64)
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint64)|y) }
			} else {
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint64)|fr.reg(iy).(uint64)) }
			}
		case reflect.Uintptr:
			if kx == kindConst && ky == kindConst {
				v := vx.(uintptr) | vy.(uintptr)
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := vx.(uintptr)
				return func(fr *frame) { fr.setReg(ir, x|fr.reg(iy).(uintptr)) }
			} else if ky == kindConst {
				y := vy.(uintptr)
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uintptr)|y) }
			} else {
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uintptr)|fr.reg(iy).(uintptr)) }
			}
		}
	} else {
		t := basic.TypeOfType(typ)
		switch typ.Kind() {
		case reflect.Int:
			if kx == kindConst && ky == kindConst {
				v := basic.Make(t, basic.Int(vx)|basic.Int(vy))
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := basic.Int(vx)
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, x|fr.int(iy))) }
			} else if ky == kindConst {
				y := basic.Int(vy)
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, fr.int(ix)|y)) }
			} else {
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, fr.int(ix)|fr.int(iy))) }
			}
		case reflect.Int8:
			if kx == kindConst && ky == kindConst {
				v := basic.Make(t, basic.Int8(vx)|basic.Int8(vy))
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := basic.Int8(vx)
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, x|fr.int8(iy))) }
			} else if ky == kindConst {
				y := basic.Int8(vy)
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, fr.int8(ix)|y)) }
			} else {
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, fr.int8(ix)|fr.int8(iy))) }
			}
		case reflect.Int16:
			if kx == kindConst && ky == kindConst {
				v := basic.Make(t, basic.Int16(vx)|basic.Int16(vy))
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := basic.Int16(vx)
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, x|fr.int16(iy))) }
			} else if ky == kindConst {
				y := basic.Int16(vy)
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, fr.int16(ix)|y)) }
			} else {
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, fr.int16(ix)|fr.int16(iy))) }
			}
		case reflect.Int32:
			if kx == kindConst && ky == kindConst {
				v := basic.Make(t, basic.Int32(vx)|basic.Int32(vy))
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := basic.Int32(vx)
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, x|fr.int32(iy))) }
			} else if ky == kindConst {
				y := basic.Int32(vy)
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, fr.int32(ix)|y)) }
			} else {
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, fr.int32(ix)|fr.int32(iy))) }
			}
		case reflect.Int64:
			if kx == kindConst && ky == kindConst {
				v := basic.Make(t, basic.Int64(vx)|basic.Int64(vy))
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := basic.Int64(vx)
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, x|fr.int64(iy))) }
			} else if ky == kindConst {
				y := basic.Int64(vy)
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, fr.int64(ix)|y)) }
			} else {
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, fr.int64(ix)|fr.int64(iy))) }
			}
		case reflect.Uint:
			if kx == kindConst && ky == kindConst {
				v := basic.Make(t, basic.Uint(vx)|basic.Uint(vy))
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := basic.Uint(vx)
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, x|fr.uint(iy))) }
			} else if ky == kindConst {
				y := basic.Uint(vy)
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, fr.uint(ix)|y)) }
			} else {
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, fr.uint(ix)|fr.uint(iy))) }
			}
		case reflect.Uint8:
			if kx == kindConst && ky == kindConst {
				v := basic.Make(t, basic.Uint8(vx)|basic.Uint8(vy))
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := basic.Uint8(vx)
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, x|fr.uint8(iy))) }
			} else if ky == kindConst {
				y := basic.Uint8(vy)
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, fr.uint8(ix)|y)) }
			} else {
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, fr.uint8(ix)|fr.uint8(iy))) }
			}
		case reflect.Uint16:
			if kx == kindConst && ky == kindConst {
				v := basic.Make(t, basic.Uint16(vx)|basic.Uint16(vy))
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := basic.Uint16(vx)
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, x|fr.uint16(iy))) }
			} else if ky == kindConst {
				y := basic.Uint16(vy)
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, fr.uint16(ix)|y)) }
			} else {
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, fr.uint16(ix)|fr.uint16(iy))) }
			}
		case reflect.Uint32:
			if kx == kindConst && ky == kindConst {
				v := basic.Make(t, basic.Uint32(vx)|basic.Uint32(vy))
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := basic.Uint32(vx)
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, x|fr.uint32(iy))) }
			} else if ky == kindConst {
				y := basic.Uint32(vy)
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, fr.uint32(ix)|y)) }
			} else {
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, fr.uint32(ix)|fr.uint32(iy))) }
			}
		case reflect.Uint64:
			if kx == kindConst && ky == kindConst {
				v := basic.Make(t, basic.Uint64(vx)|basic.Uint64(vy))
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := basic.Uint64(vx)
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, x|fr.uint64(iy))) }
			} else if ky == kindConst {
				y := basic.Uint64(vy)
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, fr.uint64(ix)|y)) }
			} else {
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, fr.uint64(ix)|fr.uint64(iy))) }
			}
		case reflect.Uintptr:
			if kx == kindConst && ky == kindConst {
				v := basic.Make(t, basic.Uintptr(vx)|basic.Uintptr(vy))
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := basic.Uintptr(vx)
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, x|fr.uintptr(iy))) }
			} else if ky == kindConst {
				y := basic.Uintptr(vy)
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, fr.uintptr(ix)|y)) }
			} else {
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, fr.uintptr(ix)|fr.uintptr(iy))) }
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
			if kx == kindConst && ky == kindConst {
				v := vx.(int) ^ vy.(int)
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := vx.(int)
				return func(fr *frame) { fr.setReg(ir, x^fr.reg(iy).(int)) }
			} else if ky == kindConst {
				y := vy.(int)
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int)^y) }
			} else {
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int)^fr.reg(iy).(int)) }
			}
		case reflect.Int8:
			if kx == kindConst && ky == kindConst {
				v := vx.(int8) ^ vy.(int8)
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := vx.(int8)
				return func(fr *frame) { fr.setReg(ir, x^fr.reg(iy).(int8)) }
			} else if ky == kindConst {
				y := vy.(int8)
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int8)^y) }
			} else {
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int8)^fr.reg(iy).(int8)) }
			}
		case reflect.Int16:
			if kx == kindConst && ky == kindConst {
				v := vx.(int16) ^ vy.(int16)
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := vx.(int16)
				return func(fr *frame) { fr.setReg(ir, x^fr.reg(iy).(int16)) }
			} else if ky == kindConst {
				y := vy.(int16)
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int16)^y) }
			} else {
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int16)^fr.reg(iy).(int16)) }
			}
		case reflect.Int32:
			if kx == kindConst && ky == kindConst {
				v := vx.(int32) ^ vy.(int32)
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := vx.(int32)
				return func(fr *frame) { fr.setReg(ir, x^fr.reg(iy).(int32)) }
			} else if ky == kindConst {
				y := vy.(int32)
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int32)^y) }
			} else {
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int32)^fr.reg(iy).(int32)) }
			}
		case reflect.Int64:
			if kx == kindConst && ky == kindConst {
				v := vx.(int64) ^ vy.(int64)
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := vx.(int64)
				return func(fr *frame) { fr.setReg(ir, x^fr.reg(iy).(int64)) }
			} else if ky == kindConst {
				y := vy.(int64)
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int64)^y) }
			} else {
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int64)^fr.reg(iy).(int64)) }
			}
		case reflect.Uint:
			if kx == kindConst && ky == kindConst {
				v := vx.(uint) ^ vy.(uint)
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := vx.(uint)
				return func(fr *frame) { fr.setReg(ir, x^fr.reg(iy).(uint)) }
			} else if ky == kindConst {
				y := vy.(uint)
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint)^y) }
			} else {
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint)^fr.reg(iy).(uint)) }
			}
		case reflect.Uint8:
			if kx == kindConst && ky == kindConst {
				v := vx.(uint8) ^ vy.(uint8)
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := vx.(uint8)
				return func(fr *frame) { fr.setReg(ir, x^fr.reg(iy).(uint8)) }
			} else if ky == kindConst {
				y := vy.(uint8)
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint8)^y) }
			} else {
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint8)^fr.reg(iy).(uint8)) }
			}
		case reflect.Uint16:
			if kx == kindConst && ky == kindConst {
				v := vx.(uint16) ^ vy.(uint16)
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := vx.(uint16)
				return func(fr *frame) { fr.setReg(ir, x^fr.reg(iy).(uint16)) }
			} else if ky == kindConst {
				y := vy.(uint16)
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint16)^y) }
			} else {
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint16)^fr.reg(iy).(uint16)) }
			}
		case reflect.Uint32:
			if kx == kindConst && ky == kindConst {
				v := vx.(uint32) ^ vy.(uint32)
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := vx.(uint32)
				return func(fr *frame) { fr.setReg(ir, x^fr.reg(iy).(uint32)) }
			} else if ky == kindConst {
				y := vy.(uint32)
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint32)^y) }
			} else {
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint32)^fr.reg(iy).(uint32)) }
			}
		case reflect.Uint64:
			if kx == kindConst && ky == kindConst {
				v := vx.(uint64) ^ vy.(uint64)
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := vx.(uint64)
				return func(fr *frame) { fr.setReg(ir, x^fr.reg(iy).(uint64)) }
			} else if ky == kindConst {
				y := vy.(uint64)
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint64)^y) }
			} else {
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint64)^fr.reg(iy).(uint64)) }
			}
		case reflect.Uintptr:
			if kx == kindConst && ky == kindConst {
				v := vx.(uintptr) ^ vy.(uintptr)
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := vx.(uintptr)
				return func(fr *frame) { fr.setReg(ir, x^fr.reg(iy).(uintptr)) }
			} else if ky == kindConst {
				y := vy.(uintptr)
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uintptr)^y) }
			} else {
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uintptr)^fr.reg(iy).(uintptr)) }
			}
		}
	} else {
		t := basic.TypeOfType(typ)
		switch typ.Kind() {
		case reflect.Int:
			if kx == kindConst && ky == kindConst {
				v := basic.Make(t, basic.Int(vx)^basic.Int(vy))
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := basic.Int(vx)
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, x^fr.int(iy))) }
			} else if ky == kindConst {
				y := basic.Int(vy)
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, fr.int(ix)^y)) }
			} else {
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, fr.int(ix)^fr.int(iy))) }
			}
		case reflect.Int8:
			if kx == kindConst && ky == kindConst {
				v := basic.Make(t, basic.Int8(vx)^basic.Int8(vy))
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := basic.Int8(vx)
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, x^fr.int8(iy))) }
			} else if ky == kindConst {
				y := basic.Int8(vy)
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, fr.int8(ix)^y)) }
			} else {
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, fr.int8(ix)^fr.int8(iy))) }
			}
		case reflect.Int16:
			if kx == kindConst && ky == kindConst {
				v := basic.Make(t, basic.Int16(vx)^basic.Int16(vy))
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := basic.Int16(vx)
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, x^fr.int16(iy))) }
			} else if ky == kindConst {
				y := basic.Int16(vy)
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, fr.int16(ix)^y)) }
			} else {
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, fr.int16(ix)^fr.int16(iy))) }
			}
		case reflect.Int32:
			if kx == kindConst && ky == kindConst {
				v := basic.Make(t, basic.Int32(vx)^basic.Int32(vy))
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := basic.Int32(vx)
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, x^fr.int32(iy))) }
			} else if ky == kindConst {
				y := basic.Int32(vy)
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, fr.int32(ix)^y)) }
			} else {
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, fr.int32(ix)^fr.int32(iy))) }
			}
		case reflect.Int64:
			if kx == kindConst && ky == kindConst {
				v := basic.Make(t, basic.Int64(vx)^basic.Int64(vy))
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := basic.Int64(vx)
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, x^fr.int64(iy))) }
			} else if ky == kindConst {
				y := basic.Int64(vy)
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, fr.int64(ix)^y)) }
			} else {
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, fr.int64(ix)^fr.int64(iy))) }
			}
		case reflect.Uint:
			if kx == kindConst && ky == kindConst {
				v := basic.Make(t, basic.Uint(vx)^basic.Uint(vy))
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := basic.Uint(vx)
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, x^fr.uint(iy))) }
			} else if ky == kindConst {
				y := basic.Uint(vy)
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, fr.uint(ix)^y)) }
			} else {
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, fr.uint(ix)^fr.uint(iy))) }
			}
		case reflect.Uint8:
			if kx == kindConst && ky == kindConst {
				v := basic.Make(t, basic.Uint8(vx)^basic.Uint8(vy))
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := basic.Uint8(vx)
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, x^fr.uint8(iy))) }
			} else if ky == kindConst {
				y := basic.Uint8(vy)
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, fr.uint8(ix)^y)) }
			} else {
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, fr.uint8(ix)^fr.uint8(iy))) }
			}
		case reflect.Uint16:
			if kx == kindConst && ky == kindConst {
				v := basic.Make(t, basic.Uint16(vx)^basic.Uint16(vy))
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := basic.Uint16(vx)
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, x^fr.uint16(iy))) }
			} else if ky == kindConst {
				y := basic.Uint16(vy)
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, fr.uint16(ix)^y)) }
			} else {
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, fr.uint16(ix)^fr.uint16(iy))) }
			}
		case reflect.Uint32:
			if kx == kindConst && ky == kindConst {
				v := basic.Make(t, basic.Uint32(vx)^basic.Uint32(vy))
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := basic.Uint32(vx)
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, x^fr.uint32(iy))) }
			} else if ky == kindConst {
				y := basic.Uint32(vy)
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, fr.uint32(ix)^y)) }
			} else {
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, fr.uint32(ix)^fr.uint32(iy))) }
			}
		case reflect.Uint64:
			if kx == kindConst && ky == kindConst {
				v := basic.Make(t, basic.Uint64(vx)^basic.Uint64(vy))
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := basic.Uint64(vx)
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, x^fr.uint64(iy))) }
			} else if ky == kindConst {
				y := basic.Uint64(vy)
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, fr.uint64(ix)^y)) }
			} else {
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, fr.uint64(ix)^fr.uint64(iy))) }
			}
		case reflect.Uintptr:
			if kx == kindConst && ky == kindConst {
				v := basic.Make(t, basic.Uintptr(vx)^basic.Uintptr(vy))
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := basic.Uintptr(vx)
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, x^fr.uintptr(iy))) }
			} else if ky == kindConst {
				y := basic.Uintptr(vy)
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, fr.uintptr(ix)^y)) }
			} else {
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, fr.uintptr(ix)^fr.uintptr(iy))) }
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
			if kx == kindConst && ky == kindConst {
				v := vx.(int) &^ vy.(int)
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := vx.(int)
				return func(fr *frame) { fr.setReg(ir, x&^fr.reg(iy).(int)) }
			} else if ky == kindConst {
				y := vy.(int)
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int)&^y) }
			} else {
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int)&^fr.reg(iy).(int)) }
			}
		case reflect.Int8:
			if kx == kindConst && ky == kindConst {
				v := vx.(int8) &^ vy.(int8)
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := vx.(int8)
				return func(fr *frame) { fr.setReg(ir, x&^fr.reg(iy).(int8)) }
			} else if ky == kindConst {
				y := vy.(int8)
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int8)&^y) }
			} else {
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int8)&^fr.reg(iy).(int8)) }
			}
		case reflect.Int16:
			if kx == kindConst && ky == kindConst {
				v := vx.(int16) &^ vy.(int16)
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := vx.(int16)
				return func(fr *frame) { fr.setReg(ir, x&^fr.reg(iy).(int16)) }
			} else if ky == kindConst {
				y := vy.(int16)
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int16)&^y) }
			} else {
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int16)&^fr.reg(iy).(int16)) }
			}
		case reflect.Int32:
			if kx == kindConst && ky == kindConst {
				v := vx.(int32) &^ vy.(int32)
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := vx.(int32)
				return func(fr *frame) { fr.setReg(ir, x&^fr.reg(iy).(int32)) }
			} else if ky == kindConst {
				y := vy.(int32)
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int32)&^y) }
			} else {
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int32)&^fr.reg(iy).(int32)) }
			}
		case reflect.Int64:
			if kx == kindConst && ky == kindConst {
				v := vx.(int64) &^ vy.(int64)
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := vx.(int64)
				return func(fr *frame) { fr.setReg(ir, x&^fr.reg(iy).(int64)) }
			} else if ky == kindConst {
				y := vy.(int64)
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int64)&^y) }
			} else {
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int64)&^fr.reg(iy).(int64)) }
			}
		case reflect.Uint:
			if kx == kindConst && ky == kindConst {
				v := vx.(uint) &^ vy.(uint)
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := vx.(uint)
				return func(fr *frame) { fr.setReg(ir, x&^fr.reg(iy).(uint)) }
			} else if ky == kindConst {
				y := vy.(uint)
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint)&^y) }
			} else {
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint)&^fr.reg(iy).(uint)) }
			}
		case reflect.Uint8:
			if kx == kindConst && ky == kindConst {
				v := vx.(uint8) &^ vy.(uint8)
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := vx.(uint8)
				return func(fr *frame) { fr.setReg(ir, x&^fr.reg(iy).(uint8)) }
			} else if ky == kindConst {
				y := vy.(uint8)
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint8)&^y) }
			} else {
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint8)&^fr.reg(iy).(uint8)) }
			}
		case reflect.Uint16:
			if kx == kindConst && ky == kindConst {
				v := vx.(uint16) &^ vy.(uint16)
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := vx.(uint16)
				return func(fr *frame) { fr.setReg(ir, x&^fr.reg(iy).(uint16)) }
			} else if ky == kindConst {
				y := vy.(uint16)
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint16)&^y) }
			} else {
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint16)&^fr.reg(iy).(uint16)) }
			}
		case reflect.Uint32:
			if kx == kindConst && ky == kindConst {
				v := vx.(uint32) &^ vy.(uint32)
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := vx.(uint32)
				return func(fr *frame) { fr.setReg(ir, x&^fr.reg(iy).(uint32)) }
			} else if ky == kindConst {
				y := vy.(uint32)
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint32)&^y) }
			} else {
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint32)&^fr.reg(iy).(uint32)) }
			}
		case reflect.Uint64:
			if kx == kindConst && ky == kindConst {
				v := vx.(uint64) &^ vy.(uint64)
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := vx.(uint64)
				return func(fr *frame) { fr.setReg(ir, x&^fr.reg(iy).(uint64)) }
			} else if ky == kindConst {
				y := vy.(uint64)
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint64)&^y) }
			} else {
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint64)&^fr.reg(iy).(uint64)) }
			}
		case reflect.Uintptr:
			if kx == kindConst && ky == kindConst {
				v := vx.(uintptr) &^ vy.(uintptr)
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := vx.(uintptr)
				return func(fr *frame) { fr.setReg(ir, x&^fr.reg(iy).(uintptr)) }
			} else if ky == kindConst {
				y := vy.(uintptr)
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uintptr)&^y) }
			} else {
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uintptr)&^fr.reg(iy).(uintptr)) }
			}
		}
	} else {
		t := basic.TypeOfType(typ)
		switch typ.Kind() {
		case reflect.Int:
			if kx == kindConst && ky == kindConst {
				v := basic.Make(t, basic.Int(vx)&^basic.Int(vy))
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := basic.Int(vx)
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, x&^fr.int(iy))) }
			} else if ky == kindConst {
				y := basic.Int(vy)
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, fr.int(ix)&^y)) }
			} else {
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, fr.int(ix)&^fr.int(iy))) }
			}
		case reflect.Int8:
			if kx == kindConst && ky == kindConst {
				v := basic.Make(t, basic.Int8(vx)&^basic.Int8(vy))
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := basic.Int8(vx)
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, x&^fr.int8(iy))) }
			} else if ky == kindConst {
				y := basic.Int8(vy)
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, fr.int8(ix)&^y)) }
			} else {
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, fr.int8(ix)&^fr.int8(iy))) }
			}
		case reflect.Int16:
			if kx == kindConst && ky == kindConst {
				v := basic.Make(t, basic.Int16(vx)&^basic.Int16(vy))
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := basic.Int16(vx)
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, x&^fr.int16(iy))) }
			} else if ky == kindConst {
				y := basic.Int16(vy)
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, fr.int16(ix)&^y)) }
			} else {
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, fr.int16(ix)&^fr.int16(iy))) }
			}
		case reflect.Int32:
			if kx == kindConst && ky == kindConst {
				v := basic.Make(t, basic.Int32(vx)&^basic.Int32(vy))
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := basic.Int32(vx)
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, x&^fr.int32(iy))) }
			} else if ky == kindConst {
				y := basic.Int32(vy)
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, fr.int32(ix)&^y)) }
			} else {
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, fr.int32(ix)&^fr.int32(iy))) }
			}
		case reflect.Int64:
			if kx == kindConst && ky == kindConst {
				v := basic.Make(t, basic.Int64(vx)&^basic.Int64(vy))
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := basic.Int64(vx)
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, x&^fr.int64(iy))) }
			} else if ky == kindConst {
				y := basic.Int64(vy)
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, fr.int64(ix)&^y)) }
			} else {
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, fr.int64(ix)&^fr.int64(iy))) }
			}
		case reflect.Uint:
			if kx == kindConst && ky == kindConst {
				v := basic.Make(t, basic.Uint(vx)&^basic.Uint(vy))
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := basic.Uint(vx)
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, x&^fr.uint(iy))) }
			} else if ky == kindConst {
				y := basic.Uint(vy)
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, fr.uint(ix)&^y)) }
			} else {
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, fr.uint(ix)&^fr.uint(iy))) }
			}
		case reflect.Uint8:
			if kx == kindConst && ky == kindConst {
				v := basic.Make(t, basic.Uint8(vx)&^basic.Uint8(vy))
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := basic.Uint8(vx)
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, x&^fr.uint8(iy))) }
			} else if ky == kindConst {
				y := basic.Uint8(vy)
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, fr.uint8(ix)&^y)) }
			} else {
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, fr.uint8(ix)&^fr.uint8(iy))) }
			}
		case reflect.Uint16:
			if kx == kindConst && ky == kindConst {
				v := basic.Make(t, basic.Uint16(vx)&^basic.Uint16(vy))
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := basic.Uint16(vx)
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, x&^fr.uint16(iy))) }
			} else if ky == kindConst {
				y := basic.Uint16(vy)
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, fr.uint16(ix)&^y)) }
			} else {
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, fr.uint16(ix)&^fr.uint16(iy))) }
			}
		case reflect.Uint32:
			if kx == kindConst && ky == kindConst {
				v := basic.Make(t, basic.Uint32(vx)&^basic.Uint32(vy))
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := basic.Uint32(vx)
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, x&^fr.uint32(iy))) }
			} else if ky == kindConst {
				y := basic.Uint32(vy)
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, fr.uint32(ix)&^y)) }
			} else {
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, fr.uint32(ix)&^fr.uint32(iy))) }
			}
		case reflect.Uint64:
			if kx == kindConst && ky == kindConst {
				v := basic.Make(t, basic.Uint64(vx)&^basic.Uint64(vy))
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := basic.Uint64(vx)
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, x&^fr.uint64(iy))) }
			} else if ky == kindConst {
				y := basic.Uint64(vy)
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, fr.uint64(ix)&^y)) }
			} else {
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, fr.uint64(ix)&^fr.uint64(iy))) }
			}
		case reflect.Uintptr:
			if kx == kindConst && ky == kindConst {
				v := basic.Make(t, basic.Uintptr(vx)&^basic.Uintptr(vy))
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := basic.Uintptr(vx)
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, x&^fr.uintptr(iy))) }
			} else if ky == kindConst {
				y := basic.Uintptr(vy)
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, fr.uintptr(ix)&^y)) }
			} else {
				return func(fr *frame) { fr.setReg(ir, basic.Make(t, fr.uintptr(ix)&^fr.uintptr(iy))) }
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
			if kx == kindConst && ky == kindConst {
				v := vx.(int) < vy.(int)
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := vx.(int)
				return func(fr *frame) { fr.setReg(ir, x < fr.reg(iy).(int)) }
			} else if ky == kindConst {
				y := vy.(int)
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int) < y) }
			} else {
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int) < fr.reg(iy).(int)) }
			}
		case reflect.Int8:
			if kx == kindConst && ky == kindConst {
				v := vx.(int8) < vy.(int8)
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := vx.(int8)
				return func(fr *frame) { fr.setReg(ir, x < fr.reg(iy).(int8)) }
			} else if ky == kindConst {
				y := vy.(int8)
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int8) < y) }
			} else {
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int8) < fr.reg(iy).(int8)) }
			}
		case reflect.Int16:
			if kx == kindConst && ky == kindConst {
				v := vx.(int16) < vy.(int16)
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := vx.(int16)
				return func(fr *frame) { fr.setReg(ir, x < fr.reg(iy).(int16)) }
			} else if ky == kindConst {
				y := vy.(int16)
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int16) < y) }
			} else {
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int16) < fr.reg(iy).(int16)) }
			}
		case reflect.Int32:
			if kx == kindConst && ky == kindConst {
				v := vx.(int32) < vy.(int32)
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := vx.(int32)
				return func(fr *frame) { fr.setReg(ir, x < fr.reg(iy).(int32)) }
			} else if ky == kindConst {
				y := vy.(int32)
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int32) < y) }
			} else {
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int32) < fr.reg(iy).(int32)) }
			}
		case reflect.Int64:
			if kx == kindConst && ky == kindConst {
				v := vx.(int64) < vy.(int64)
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := vx.(int64)
				return func(fr *frame) { fr.setReg(ir, x < fr.reg(iy).(int64)) }
			} else if ky == kindConst {
				y := vy.(int64)
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int64) < y) }
			} else {
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int64) < fr.reg(iy).(int64)) }
			}
		case reflect.Uint:
			if kx == kindConst && ky == kindConst {
				v := vx.(uint) < vy.(uint)
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := vx.(uint)
				return func(fr *frame) { fr.setReg(ir, x < fr.reg(iy).(uint)) }
			} else if ky == kindConst {
				y := vy.(uint)
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint) < y) }
			} else {
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint) < fr.reg(iy).(uint)) }
			}
		case reflect.Uint8:
			if kx == kindConst && ky == kindConst {
				v := vx.(uint8) < vy.(uint8)
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := vx.(uint8)
				return func(fr *frame) { fr.setReg(ir, x < fr.reg(iy).(uint8)) }
			} else if ky == kindConst {
				y := vy.(uint8)
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint8) < y) }
			} else {
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint8) < fr.reg(iy).(uint8)) }
			}
		case reflect.Uint16:
			if kx == kindConst && ky == kindConst {
				v := vx.(uint16) < vy.(uint16)
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := vx.(uint16)
				return func(fr *frame) { fr.setReg(ir, x < fr.reg(iy).(uint16)) }
			} else if ky == kindConst {
				y := vy.(uint16)
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint16) < y) }
			} else {
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint16) < fr.reg(iy).(uint16)) }
			}
		case reflect.Uint32:
			if kx == kindConst && ky == kindConst {
				v := vx.(uint32) < vy.(uint32)
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := vx.(uint32)
				return func(fr *frame) { fr.setReg(ir, x < fr.reg(iy).(uint32)) }
			} else if ky == kindConst {
				y := vy.(uint32)
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint32) < y) }
			} else {
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint32) < fr.reg(iy).(uint32)) }
			}
		case reflect.Uint64:
			if kx == kindConst && ky == kindConst {
				v := vx.(uint64) < vy.(uint64)
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := vx.(uint64)
				return func(fr *frame) { fr.setReg(ir, x < fr.reg(iy).(uint64)) }
			} else if ky == kindConst {
				y := vy.(uint64)
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint64) < y) }
			} else {
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint64) < fr.reg(iy).(uint64)) }
			}
		case reflect.Uintptr:
			if kx == kindConst && ky == kindConst {
				v := vx.(uintptr) < vy.(uintptr)
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := vx.(uintptr)
				return func(fr *frame) { fr.setReg(ir, x < fr.reg(iy).(uintptr)) }
			} else if ky == kindConst {
				y := vy.(uintptr)
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uintptr) < y) }
			} else {
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uintptr) < fr.reg(iy).(uintptr)) }
			}
		case reflect.Float32:
			if kx == kindConst && ky == kindConst {
				v := vx.(float32) < vy.(float32)
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := vx.(float32)
				return func(fr *frame) { fr.setReg(ir, x < fr.reg(iy).(float32)) }
			} else if ky == kindConst {
				y := vy.(float32)
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(float32) < y) }
			} else {
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(float32) < fr.reg(iy).(float32)) }
			}
		case reflect.Float64:
			if kx == kindConst && ky == kindConst {
				v := vx.(float64) < vy.(float64)
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := vx.(float64)
				return func(fr *frame) { fr.setReg(ir, x < fr.reg(iy).(float64)) }
			} else if ky == kindConst {
				y := vy.(float64)
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(float64) < y) }
			} else {
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(float64) < fr.reg(iy).(float64)) }
			}
		case reflect.String:
			if kx == kindConst && ky == kindConst {
				v := vx.(string) < vy.(string)
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := vx.(string)
				return func(fr *frame) { fr.setReg(ir, x < fr.reg(iy).(string)) }
			} else if ky == kindConst {
				y := vy.(string)
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(string) < y) }
			} else {
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(string) < fr.reg(iy).(string)) }
			}
		}
	} else {
		switch typ.Kind() {
		case reflect.Int:
			if kx == kindConst && ky == kindConst {
				v := basic.Int(vx) < basic.Int(vy)
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := basic.Int(vx)
				return func(fr *frame) { fr.setReg(ir, x < fr.int(iy)) }
			} else if ky == kindConst {
				y := basic.Int(vy)
				return func(fr *frame) { fr.setReg(ir, fr.int(ix) < y) }
			} else {
				return func(fr *frame) { fr.setReg(ir, fr.int(ix) < fr.int(iy)) }
			}
		case reflect.Int8:
			if kx == kindConst && ky == kindConst {
				v := basic.Int8(vx) < basic.Int8(vy)
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := basic.Int8(vx)
				return func(fr *frame) { fr.setReg(ir, x < fr.int8(iy)) }
			} else if ky == kindConst {
				y := basic.Int8(vy)
				return func(fr *frame) { fr.setReg(ir, fr.int8(ix) < y) }
			} else {
				return func(fr *frame) { fr.setReg(ir, fr.int8(ix) < fr.int8(iy)) }
			}
		case reflect.Int16:
			if kx == kindConst && ky == kindConst {
				v := basic.Int16(vx) < basic.Int16(vy)
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := basic.Int16(vx)
				return func(fr *frame) { fr.setReg(ir, x < fr.int16(iy)) }
			} else if ky == kindConst {
				y := basic.Int16(vy)
				return func(fr *frame) { fr.setReg(ir, fr.int16(ix) < y) }
			} else {
				return func(fr *frame) { fr.setReg(ir, fr.int16(ix) < fr.int16(iy)) }
			}
		case reflect.Int32:
			if kx == kindConst && ky == kindConst {
				v := basic.Int32(vx) < basic.Int32(vy)
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := basic.Int32(vx)
				return func(fr *frame) { fr.setReg(ir, x < fr.int32(iy)) }
			} else if ky == kindConst {
				y := basic.Int32(vy)
				return func(fr *frame) { fr.setReg(ir, fr.int32(ix) < y) }
			} else {
				return func(fr *frame) { fr.setReg(ir, fr.int32(ix) < fr.int32(iy)) }
			}
		case reflect.Int64:
			if kx == kindConst && ky == kindConst {
				v := basic.Int64(vx) < basic.Int64(vy)
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := basic.Int64(vx)
				return func(fr *frame) { fr.setReg(ir, x < fr.int64(iy)) }
			} else if ky == kindConst {
				y := basic.Int64(vy)
				return func(fr *frame) { fr.setReg(ir, fr.int64(ix) < y) }
			} else {
				return func(fr *frame) { fr.setReg(ir, fr.int64(ix) < fr.int64(iy)) }
			}
		case reflect.Uint:
			if kx == kindConst && ky == kindConst {
				v := basic.Uint(vx) < basic.Uint(vy)
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := basic.Uint(vx)
				return func(fr *frame) { fr.setReg(ir, x < fr.uint(iy)) }
			} else if ky == kindConst {
				y := basic.Uint(vy)
				return func(fr *frame) { fr.setReg(ir, fr.uint(ix) < y) }
			} else {
				return func(fr *frame) { fr.setReg(ir, fr.uint(ix) < fr.uint(iy)) }
			}
		case reflect.Uint8:
			if kx == kindConst && ky == kindConst {
				v := basic.Uint8(vx) < basic.Uint8(vy)
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := basic.Uint8(vx)
				return func(fr *frame) { fr.setReg(ir, x < fr.uint8(iy)) }
			} else if ky == kindConst {
				y := basic.Uint8(vy)
				return func(fr *frame) { fr.setReg(ir, fr.uint8(ix) < y) }
			} else {
				return func(fr *frame) { fr.setReg(ir, fr.uint8(ix) < fr.uint8(iy)) }
			}
		case reflect.Uint16:
			if kx == kindConst && ky == kindConst {
				v := basic.Uint16(vx) < basic.Uint16(vy)
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := basic.Uint16(vx)
				return func(fr *frame) { fr.setReg(ir, x < fr.uint16(iy)) }
			} else if ky == kindConst {
				y := basic.Uint16(vy)
				return func(fr *frame) { fr.setReg(ir, fr.uint16(ix) < y) }
			} else {
				return func(fr *frame) { fr.setReg(ir, fr.uint16(ix) < fr.uint16(iy)) }
			}
		case reflect.Uint32:
			if kx == kindConst && ky == kindConst {
				v := basic.Uint32(vx) < basic.Uint32(vy)
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := basic.Uint32(vx)
				return func(fr *frame) { fr.setReg(ir, x < fr.uint32(iy)) }
			} else if ky == kindConst {
				y := basic.Uint32(vy)
				return func(fr *frame) { fr.setReg(ir, fr.uint32(ix) < y) }
			} else {
				return func(fr *frame) { fr.setReg(ir, fr.uint32(ix) < fr.uint32(iy)) }
			}
		case reflect.Uint64:
			if kx == kindConst && ky == kindConst {
				v := basic.Uint64(vx) < basic.Uint64(vy)
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := basic.Uint64(vx)
				return func(fr *frame) { fr.setReg(ir, x < fr.uint64(iy)) }
			} else if ky == kindConst {
				y := basic.Uint64(vy)
				return func(fr *frame) { fr.setReg(ir, fr.uint64(ix) < y) }
			} else {
				return func(fr *frame) { fr.setReg(ir, fr.uint64(ix) < fr.uint64(iy)) }
			}
		case reflect.Uintptr:
			if kx == kindConst && ky == kindConst {
				v := basic.Uintptr(vx) < basic.Uintptr(vy)
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := basic.Uintptr(vx)
				return func(fr *frame) { fr.setReg(ir, x < fr.uintptr(iy)) }
			} else if ky == kindConst {
				y := basic.Uintptr(vy)
				return func(fr *frame) { fr.setReg(ir, fr.uintptr(ix) < y) }
			} else {
				return func(fr *frame) { fr.setReg(ir, fr.uintptr(ix) < fr.uintptr(iy)) }
			}
		case reflect.Float32:
			if kx == kindConst && ky == kindConst {
				v := basic.Float32(vx) < basic.Float32(vy)
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := basic.Float32(vx)
				return func(fr *frame) { fr.setReg(ir, x < fr.float32(iy)) }
			} else if ky == kindConst {
				y := basic.Float32(vy)
				return func(fr *frame) { fr.setReg(ir, fr.float32(ix) < y) }
			} else {
				return func(fr *frame) { fr.setReg(ir, fr.float32(ix) < fr.float32(iy)) }
			}
		case reflect.Float64:
			if kx == kindConst && ky == kindConst {
				v := basic.Float64(vx) < basic.Float64(vy)
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := basic.Float64(vx)
				return func(fr *frame) { fr.setReg(ir, x < fr.float64(iy)) }
			} else if ky == kindConst {
				y := basic.Float64(vy)
				return func(fr *frame) { fr.setReg(ir, fr.float64(ix) < y) }
			} else {
				return func(fr *frame) { fr.setReg(ir, fr.float64(ix) < fr.float64(iy)) }
			}
		case reflect.String:
			if kx == kindConst && ky == kindConst {
				v := basic.String(vx) < basic.String(vy)
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := basic.String(vx)
				return func(fr *frame) { fr.setReg(ir, x < fr.string(iy)) }
			} else if ky == kindConst {
				y := basic.String(vy)
				return func(fr *frame) { fr.setReg(ir, fr.string(ix) < y) }
			} else {
				return func(fr *frame) { fr.setReg(ir, fr.string(ix) < fr.string(iy)) }
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
			if kx == kindConst && ky == kindConst {
				v := vx.(int) <= vy.(int)
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := vx.(int)
				return func(fr *frame) { fr.setReg(ir, x <= fr.reg(iy).(int)) }
			} else if ky == kindConst {
				y := vy.(int)
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int) <= y) }
			} else {
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int) <= fr.reg(iy).(int)) }
			}
		case reflect.Int8:
			if kx == kindConst && ky == kindConst {
				v := vx.(int8) <= vy.(int8)
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := vx.(int8)
				return func(fr *frame) { fr.setReg(ir, x <= fr.reg(iy).(int8)) }
			} else if ky == kindConst {
				y := vy.(int8)
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int8) <= y) }
			} else {
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int8) <= fr.reg(iy).(int8)) }
			}
		case reflect.Int16:
			if kx == kindConst && ky == kindConst {
				v := vx.(int16) <= vy.(int16)
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := vx.(int16)
				return func(fr *frame) { fr.setReg(ir, x <= fr.reg(iy).(int16)) }
			} else if ky == kindConst {
				y := vy.(int16)
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int16) <= y) }
			} else {
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int16) <= fr.reg(iy).(int16)) }
			}
		case reflect.Int32:
			if kx == kindConst && ky == kindConst {
				v := vx.(int32) <= vy.(int32)
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := vx.(int32)
				return func(fr *frame) { fr.setReg(ir, x <= fr.reg(iy).(int32)) }
			} else if ky == kindConst {
				y := vy.(int32)
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int32) <= y) }
			} else {
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int32) <= fr.reg(iy).(int32)) }
			}
		case reflect.Int64:
			if kx == kindConst && ky == kindConst {
				v := vx.(int64) <= vy.(int64)
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := vx.(int64)
				return func(fr *frame) { fr.setReg(ir, x <= fr.reg(iy).(int64)) }
			} else if ky == kindConst {
				y := vy.(int64)
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int64) <= y) }
			} else {
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int64) <= fr.reg(iy).(int64)) }
			}
		case reflect.Uint:
			if kx == kindConst && ky == kindConst {
				v := vx.(uint) <= vy.(uint)
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := vx.(uint)
				return func(fr *frame) { fr.setReg(ir, x <= fr.reg(iy).(uint)) }
			} else if ky == kindConst {
				y := vy.(uint)
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint) <= y) }
			} else {
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint) <= fr.reg(iy).(uint)) }
			}
		case reflect.Uint8:
			if kx == kindConst && ky == kindConst {
				v := vx.(uint8) <= vy.(uint8)
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := vx.(uint8)
				return func(fr *frame) { fr.setReg(ir, x <= fr.reg(iy).(uint8)) }
			} else if ky == kindConst {
				y := vy.(uint8)
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint8) <= y) }
			} else {
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint8) <= fr.reg(iy).(uint8)) }
			}
		case reflect.Uint16:
			if kx == kindConst && ky == kindConst {
				v := vx.(uint16) <= vy.(uint16)
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := vx.(uint16)
				return func(fr *frame) { fr.setReg(ir, x <= fr.reg(iy).(uint16)) }
			} else if ky == kindConst {
				y := vy.(uint16)
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint16) <= y) }
			} else {
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint16) <= fr.reg(iy).(uint16)) }
			}
		case reflect.Uint32:
			if kx == kindConst && ky == kindConst {
				v := vx.(uint32) <= vy.(uint32)
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := vx.(uint32)
				return func(fr *frame) { fr.setReg(ir, x <= fr.reg(iy).(uint32)) }
			} else if ky == kindConst {
				y := vy.(uint32)
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint32) <= y) }
			} else {
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint32) <= fr.reg(iy).(uint32)) }
			}
		case reflect.Uint64:
			if kx == kindConst && ky == kindConst {
				v := vx.(uint64) <= vy.(uint64)
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := vx.(uint64)
				return func(fr *frame) { fr.setReg(ir, x <= fr.reg(iy).(uint64)) }
			} else if ky == kindConst {
				y := vy.(uint64)
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint64) <= y) }
			} else {
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint64) <= fr.reg(iy).(uint64)) }
			}
		case reflect.Uintptr:
			if kx == kindConst && ky == kindConst {
				v := vx.(uintptr) <= vy.(uintptr)
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := vx.(uintptr)
				return func(fr *frame) { fr.setReg(ir, x <= fr.reg(iy).(uintptr)) }
			} else if ky == kindConst {
				y := vy.(uintptr)
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uintptr) <= y) }
			} else {
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uintptr) <= fr.reg(iy).(uintptr)) }
			}
		case reflect.Float32:
			if kx == kindConst && ky == kindConst {
				v := vx.(float32) <= vy.(float32)
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := vx.(float32)
				return func(fr *frame) { fr.setReg(ir, x <= fr.reg(iy).(float32)) }
			} else if ky == kindConst {
				y := vy.(float32)
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(float32) <= y) }
			} else {
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(float32) <= fr.reg(iy).(float32)) }
			}
		case reflect.Float64:
			if kx == kindConst && ky == kindConst {
				v := vx.(float64) <= vy.(float64)
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := vx.(float64)
				return func(fr *frame) { fr.setReg(ir, x <= fr.reg(iy).(float64)) }
			} else if ky == kindConst {
				y := vy.(float64)
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(float64) <= y) }
			} else {
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(float64) <= fr.reg(iy).(float64)) }
			}
		case reflect.String:
			if kx == kindConst && ky == kindConst {
				v := vx.(string) <= vy.(string)
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := vx.(string)
				return func(fr *frame) { fr.setReg(ir, x <= fr.reg(iy).(string)) }
			} else if ky == kindConst {
				y := vy.(string)
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(string) <= y) }
			} else {
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(string) <= fr.reg(iy).(string)) }
			}
		}
	} else {
		switch typ.Kind() {
		case reflect.Int:
			if kx == kindConst && ky == kindConst {
				v := basic.Int(vx) <= basic.Int(vy)
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := basic.Int(vx)
				return func(fr *frame) { fr.setReg(ir, x <= fr.int(iy)) }
			} else if ky == kindConst {
				y := basic.Int(vy)
				return func(fr *frame) { fr.setReg(ir, fr.int(ix) <= y) }
			} else {
				return func(fr *frame) { fr.setReg(ir, fr.int(ix) <= fr.int(iy)) }
			}
		case reflect.Int8:
			if kx == kindConst && ky == kindConst {
				v := basic.Int8(vx) <= basic.Int8(vy)
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := basic.Int8(vx)
				return func(fr *frame) { fr.setReg(ir, x <= fr.int8(iy)) }
			} else if ky == kindConst {
				y := basic.Int8(vy)
				return func(fr *frame) { fr.setReg(ir, fr.int8(ix) <= y) }
			} else {
				return func(fr *frame) { fr.setReg(ir, fr.int8(ix) <= fr.int8(iy)) }
			}
		case reflect.Int16:
			if kx == kindConst && ky == kindConst {
				v := basic.Int16(vx) <= basic.Int16(vy)
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := basic.Int16(vx)
				return func(fr *frame) { fr.setReg(ir, x <= fr.int16(iy)) }
			} else if ky == kindConst {
				y := basic.Int16(vy)
				return func(fr *frame) { fr.setReg(ir, fr.int16(ix) <= y) }
			} else {
				return func(fr *frame) { fr.setReg(ir, fr.int16(ix) <= fr.int16(iy)) }
			}
		case reflect.Int32:
			if kx == kindConst && ky == kindConst {
				v := basic.Int32(vx) <= basic.Int32(vy)
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := basic.Int32(vx)
				return func(fr *frame) { fr.setReg(ir, x <= fr.int32(iy)) }
			} else if ky == kindConst {
				y := basic.Int32(vy)
				return func(fr *frame) { fr.setReg(ir, fr.int32(ix) <= y) }
			} else {
				return func(fr *frame) { fr.setReg(ir, fr.int32(ix) <= fr.int32(iy)) }
			}
		case reflect.Int64:
			if kx == kindConst && ky == kindConst {
				v := basic.Int64(vx) <= basic.Int64(vy)
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := basic.Int64(vx)
				return func(fr *frame) { fr.setReg(ir, x <= fr.int64(iy)) }
			} else if ky == kindConst {
				y := basic.Int64(vy)
				return func(fr *frame) { fr.setReg(ir, fr.int64(ix) <= y) }
			} else {
				return func(fr *frame) { fr.setReg(ir, fr.int64(ix) <= fr.int64(iy)) }
			}
		case reflect.Uint:
			if kx == kindConst && ky == kindConst {
				v := basic.Uint(vx) <= basic.Uint(vy)
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := basic.Uint(vx)
				return func(fr *frame) { fr.setReg(ir, x <= fr.uint(iy)) }
			} else if ky == kindConst {
				y := basic.Uint(vy)
				return func(fr *frame) { fr.setReg(ir, fr.uint(ix) <= y) }
			} else {
				return func(fr *frame) { fr.setReg(ir, fr.uint(ix) <= fr.uint(iy)) }
			}
		case reflect.Uint8:
			if kx == kindConst && ky == kindConst {
				v := basic.Uint8(vx) <= basic.Uint8(vy)
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := basic.Uint8(vx)
				return func(fr *frame) { fr.setReg(ir, x <= fr.uint8(iy)) }
			} else if ky == kindConst {
				y := basic.Uint8(vy)
				return func(fr *frame) { fr.setReg(ir, fr.uint8(ix) <= y) }
			} else {
				return func(fr *frame) { fr.setReg(ir, fr.uint8(ix) <= fr.uint8(iy)) }
			}
		case reflect.Uint16:
			if kx == kindConst && ky == kindConst {
				v := basic.Uint16(vx) <= basic.Uint16(vy)
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := basic.Uint16(vx)
				return func(fr *frame) { fr.setReg(ir, x <= fr.uint16(iy)) }
			} else if ky == kindConst {
				y := basic.Uint16(vy)
				return func(fr *frame) { fr.setReg(ir, fr.uint16(ix) <= y) }
			} else {
				return func(fr *frame) { fr.setReg(ir, fr.uint16(ix) <= fr.uint16(iy)) }
			}
		case reflect.Uint32:
			if kx == kindConst && ky == kindConst {
				v := basic.Uint32(vx) <= basic.Uint32(vy)
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := basic.Uint32(vx)
				return func(fr *frame) { fr.setReg(ir, x <= fr.uint32(iy)) }
			} else if ky == kindConst {
				y := basic.Uint32(vy)
				return func(fr *frame) { fr.setReg(ir, fr.uint32(ix) <= y) }
			} else {
				return func(fr *frame) { fr.setReg(ir, fr.uint32(ix) <= fr.uint32(iy)) }
			}
		case reflect.Uint64:
			if kx == kindConst && ky == kindConst {
				v := basic.Uint64(vx) <= basic.Uint64(vy)
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := basic.Uint64(vx)
				return func(fr *frame) { fr.setReg(ir, x <= fr.uint64(iy)) }
			} else if ky == kindConst {
				y := basic.Uint64(vy)
				return func(fr *frame) { fr.setReg(ir, fr.uint64(ix) <= y) }
			} else {
				return func(fr *frame) { fr.setReg(ir, fr.uint64(ix) <= fr.uint64(iy)) }
			}
		case reflect.Uintptr:
			if kx == kindConst && ky == kindConst {
				v := basic.Uintptr(vx) <= basic.Uintptr(vy)
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := basic.Uintptr(vx)
				return func(fr *frame) { fr.setReg(ir, x <= fr.uintptr(iy)) }
			} else if ky == kindConst {
				y := basic.Uintptr(vy)
				return func(fr *frame) { fr.setReg(ir, fr.uintptr(ix) <= y) }
			} else {
				return func(fr *frame) { fr.setReg(ir, fr.uintptr(ix) <= fr.uintptr(iy)) }
			}
		case reflect.Float32:
			if kx == kindConst && ky == kindConst {
				v := basic.Float32(vx) <= basic.Float32(vy)
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := basic.Float32(vx)
				return func(fr *frame) { fr.setReg(ir, x <= fr.float32(iy)) }
			} else if ky == kindConst {
				y := basic.Float32(vy)
				return func(fr *frame) { fr.setReg(ir, fr.float32(ix) <= y) }
			} else {
				return func(fr *frame) { fr.setReg(ir, fr.float32(ix) <= fr.float32(iy)) }
			}
		case reflect.Float64:
			if kx == kindConst && ky == kindConst {
				v := basic.Float64(vx) <= basic.Float64(vy)
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := basic.Float64(vx)
				return func(fr *frame) { fr.setReg(ir, x <= fr.float64(iy)) }
			} else if ky == kindConst {
				y := basic.Float64(vy)
				return func(fr *frame) { fr.setReg(ir, fr.float64(ix) <= y) }
			} else {
				return func(fr *frame) { fr.setReg(ir, fr.float64(ix) <= fr.float64(iy)) }
			}
		case reflect.String:
			if kx == kindConst && ky == kindConst {
				v := basic.String(vx) <= basic.String(vy)
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := basic.String(vx)
				return func(fr *frame) { fr.setReg(ir, x <= fr.string(iy)) }
			} else if ky == kindConst {
				y := basic.String(vy)
				return func(fr *frame) { fr.setReg(ir, fr.string(ix) <= y) }
			} else {
				return func(fr *frame) { fr.setReg(ir, fr.string(ix) <= fr.string(iy)) }
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
			if kx == kindConst && ky == kindConst {
				v := vx.(int) > vy.(int)
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := vx.(int)
				return func(fr *frame) { fr.setReg(ir, x > fr.reg(iy).(int)) }
			} else if ky == kindConst {
				y := vy.(int)
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int) > y) }
			} else {
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int) > fr.reg(iy).(int)) }
			}
		case reflect.Int8:
			if kx == kindConst && ky == kindConst {
				v := vx.(int8) > vy.(int8)
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := vx.(int8)
				return func(fr *frame) { fr.setReg(ir, x > fr.reg(iy).(int8)) }
			} else if ky == kindConst {
				y := vy.(int8)
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int8) > y) }
			} else {
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int8) > fr.reg(iy).(int8)) }
			}
		case reflect.Int16:
			if kx == kindConst && ky == kindConst {
				v := vx.(int16) > vy.(int16)
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := vx.(int16)
				return func(fr *frame) { fr.setReg(ir, x > fr.reg(iy).(int16)) }
			} else if ky == kindConst {
				y := vy.(int16)
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int16) > y) }
			} else {
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int16) > fr.reg(iy).(int16)) }
			}
		case reflect.Int32:
			if kx == kindConst && ky == kindConst {
				v := vx.(int32) > vy.(int32)
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := vx.(int32)
				return func(fr *frame) { fr.setReg(ir, x > fr.reg(iy).(int32)) }
			} else if ky == kindConst {
				y := vy.(int32)
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int32) > y) }
			} else {
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int32) > fr.reg(iy).(int32)) }
			}
		case reflect.Int64:
			if kx == kindConst && ky == kindConst {
				v := vx.(int64) > vy.(int64)
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := vx.(int64)
				return func(fr *frame) { fr.setReg(ir, x > fr.reg(iy).(int64)) }
			} else if ky == kindConst {
				y := vy.(int64)
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int64) > y) }
			} else {
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int64) > fr.reg(iy).(int64)) }
			}
		case reflect.Uint:
			if kx == kindConst && ky == kindConst {
				v := vx.(uint) > vy.(uint)
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := vx.(uint)
				return func(fr *frame) { fr.setReg(ir, x > fr.reg(iy).(uint)) }
			} else if ky == kindConst {
				y := vy.(uint)
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint) > y) }
			} else {
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint) > fr.reg(iy).(uint)) }
			}
		case reflect.Uint8:
			if kx == kindConst && ky == kindConst {
				v := vx.(uint8) > vy.(uint8)
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := vx.(uint8)
				return func(fr *frame) { fr.setReg(ir, x > fr.reg(iy).(uint8)) }
			} else if ky == kindConst {
				y := vy.(uint8)
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint8) > y) }
			} else {
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint8) > fr.reg(iy).(uint8)) }
			}
		case reflect.Uint16:
			if kx == kindConst && ky == kindConst {
				v := vx.(uint16) > vy.(uint16)
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := vx.(uint16)
				return func(fr *frame) { fr.setReg(ir, x > fr.reg(iy).(uint16)) }
			} else if ky == kindConst {
				y := vy.(uint16)
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint16) > y) }
			} else {
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint16) > fr.reg(iy).(uint16)) }
			}
		case reflect.Uint32:
			if kx == kindConst && ky == kindConst {
				v := vx.(uint32) > vy.(uint32)
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := vx.(uint32)
				return func(fr *frame) { fr.setReg(ir, x > fr.reg(iy).(uint32)) }
			} else if ky == kindConst {
				y := vy.(uint32)
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint32) > y) }
			} else {
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint32) > fr.reg(iy).(uint32)) }
			}
		case reflect.Uint64:
			if kx == kindConst && ky == kindConst {
				v := vx.(uint64) > vy.(uint64)
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := vx.(uint64)
				return func(fr *frame) { fr.setReg(ir, x > fr.reg(iy).(uint64)) }
			} else if ky == kindConst {
				y := vy.(uint64)
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint64) > y) }
			} else {
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint64) > fr.reg(iy).(uint64)) }
			}
		case reflect.Uintptr:
			if kx == kindConst && ky == kindConst {
				v := vx.(uintptr) > vy.(uintptr)
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := vx.(uintptr)
				return func(fr *frame) { fr.setReg(ir, x > fr.reg(iy).(uintptr)) }
			} else if ky == kindConst {
				y := vy.(uintptr)
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uintptr) > y) }
			} else {
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uintptr) > fr.reg(iy).(uintptr)) }
			}
		case reflect.Float32:
			if kx == kindConst && ky == kindConst {
				v := vx.(float32) > vy.(float32)
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := vx.(float32)
				return func(fr *frame) { fr.setReg(ir, x > fr.reg(iy).(float32)) }
			} else if ky == kindConst {
				y := vy.(float32)
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(float32) > y) }
			} else {
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(float32) > fr.reg(iy).(float32)) }
			}
		case reflect.Float64:
			if kx == kindConst && ky == kindConst {
				v := vx.(float64) > vy.(float64)
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := vx.(float64)
				return func(fr *frame) { fr.setReg(ir, x > fr.reg(iy).(float64)) }
			} else if ky == kindConst {
				y := vy.(float64)
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(float64) > y) }
			} else {
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(float64) > fr.reg(iy).(float64)) }
			}
		case reflect.String:
			if kx == kindConst && ky == kindConst {
				v := vx.(string) > vy.(string)
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := vx.(string)
				return func(fr *frame) { fr.setReg(ir, x > fr.reg(iy).(string)) }
			} else if ky == kindConst {
				y := vy.(string)
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(string) > y) }
			} else {
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(string) > fr.reg(iy).(string)) }
			}
		}
	} else {
		switch typ.Kind() {
		case reflect.Int:
			if kx == kindConst && ky == kindConst {
				v := basic.Int(vx) > basic.Int(vy)
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := basic.Int(vx)
				return func(fr *frame) { fr.setReg(ir, x > fr.int(iy)) }
			} else if ky == kindConst {
				y := basic.Int(vy)
				return func(fr *frame) { fr.setReg(ir, fr.int(ix) > y) }
			} else {
				return func(fr *frame) { fr.setReg(ir, fr.int(ix) > fr.int(iy)) }
			}
		case reflect.Int8:
			if kx == kindConst && ky == kindConst {
				v := basic.Int8(vx) > basic.Int8(vy)
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := basic.Int8(vx)
				return func(fr *frame) { fr.setReg(ir, x > fr.int8(iy)) }
			} else if ky == kindConst {
				y := basic.Int8(vy)
				return func(fr *frame) { fr.setReg(ir, fr.int8(ix) > y) }
			} else {
				return func(fr *frame) { fr.setReg(ir, fr.int8(ix) > fr.int8(iy)) }
			}
		case reflect.Int16:
			if kx == kindConst && ky == kindConst {
				v := basic.Int16(vx) > basic.Int16(vy)
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := basic.Int16(vx)
				return func(fr *frame) { fr.setReg(ir, x > fr.int16(iy)) }
			} else if ky == kindConst {
				y := basic.Int16(vy)
				return func(fr *frame) { fr.setReg(ir, fr.int16(ix) > y) }
			} else {
				return func(fr *frame) { fr.setReg(ir, fr.int16(ix) > fr.int16(iy)) }
			}
		case reflect.Int32:
			if kx == kindConst && ky == kindConst {
				v := basic.Int32(vx) > basic.Int32(vy)
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := basic.Int32(vx)
				return func(fr *frame) { fr.setReg(ir, x > fr.int32(iy)) }
			} else if ky == kindConst {
				y := basic.Int32(vy)
				return func(fr *frame) { fr.setReg(ir, fr.int32(ix) > y) }
			} else {
				return func(fr *frame) { fr.setReg(ir, fr.int32(ix) > fr.int32(iy)) }
			}
		case reflect.Int64:
			if kx == kindConst && ky == kindConst {
				v := basic.Int64(vx) > basic.Int64(vy)
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := basic.Int64(vx)
				return func(fr *frame) { fr.setReg(ir, x > fr.int64(iy)) }
			} else if ky == kindConst {
				y := basic.Int64(vy)
				return func(fr *frame) { fr.setReg(ir, fr.int64(ix) > y) }
			} else {
				return func(fr *frame) { fr.setReg(ir, fr.int64(ix) > fr.int64(iy)) }
			}
		case reflect.Uint:
			if kx == kindConst && ky == kindConst {
				v := basic.Uint(vx) > basic.Uint(vy)
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := basic.Uint(vx)
				return func(fr *frame) { fr.setReg(ir, x > fr.uint(iy)) }
			} else if ky == kindConst {
				y := basic.Uint(vy)
				return func(fr *frame) { fr.setReg(ir, fr.uint(ix) > y) }
			} else {
				return func(fr *frame) { fr.setReg(ir, fr.uint(ix) > fr.uint(iy)) }
			}
		case reflect.Uint8:
			if kx == kindConst && ky == kindConst {
				v := basic.Uint8(vx) > basic.Uint8(vy)
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := basic.Uint8(vx)
				return func(fr *frame) { fr.setReg(ir, x > fr.uint8(iy)) }
			} else if ky == kindConst {
				y := basic.Uint8(vy)
				return func(fr *frame) { fr.setReg(ir, fr.uint8(ix) > y) }
			} else {
				return func(fr *frame) { fr.setReg(ir, fr.uint8(ix) > fr.uint8(iy)) }
			}
		case reflect.Uint16:
			if kx == kindConst && ky == kindConst {
				v := basic.Uint16(vx) > basic.Uint16(vy)
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := basic.Uint16(vx)
				return func(fr *frame) { fr.setReg(ir, x > fr.uint16(iy)) }
			} else if ky == kindConst {
				y := basic.Uint16(vy)
				return func(fr *frame) { fr.setReg(ir, fr.uint16(ix) > y) }
			} else {
				return func(fr *frame) { fr.setReg(ir, fr.uint16(ix) > fr.uint16(iy)) }
			}
		case reflect.Uint32:
			if kx == kindConst && ky == kindConst {
				v := basic.Uint32(vx) > basic.Uint32(vy)
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := basic.Uint32(vx)
				return func(fr *frame) { fr.setReg(ir, x > fr.uint32(iy)) }
			} else if ky == kindConst {
				y := basic.Uint32(vy)
				return func(fr *frame) { fr.setReg(ir, fr.uint32(ix) > y) }
			} else {
				return func(fr *frame) { fr.setReg(ir, fr.uint32(ix) > fr.uint32(iy)) }
			}
		case reflect.Uint64:
			if kx == kindConst && ky == kindConst {
				v := basic.Uint64(vx) > basic.Uint64(vy)
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := basic.Uint64(vx)
				return func(fr *frame) { fr.setReg(ir, x > fr.uint64(iy)) }
			} else if ky == kindConst {
				y := basic.Uint64(vy)
				return func(fr *frame) { fr.setReg(ir, fr.uint64(ix) > y) }
			} else {
				return func(fr *frame) { fr.setReg(ir, fr.uint64(ix) > fr.uint64(iy)) }
			}
		case reflect.Uintptr:
			if kx == kindConst && ky == kindConst {
				v := basic.Uintptr(vx) > basic.Uintptr(vy)
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := basic.Uintptr(vx)
				return func(fr *frame) { fr.setReg(ir, x > fr.uintptr(iy)) }
			} else if ky == kindConst {
				y := basic.Uintptr(vy)
				return func(fr *frame) { fr.setReg(ir, fr.uintptr(ix) > y) }
			} else {
				return func(fr *frame) { fr.setReg(ir, fr.uintptr(ix) > fr.uintptr(iy)) }
			}
		case reflect.Float32:
			if kx == kindConst && ky == kindConst {
				v := basic.Float32(vx) > basic.Float32(vy)
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := basic.Float32(vx)
				return func(fr *frame) { fr.setReg(ir, x > fr.float32(iy)) }
			} else if ky == kindConst {
				y := basic.Float32(vy)
				return func(fr *frame) { fr.setReg(ir, fr.float32(ix) > y) }
			} else {
				return func(fr *frame) { fr.setReg(ir, fr.float32(ix) > fr.float32(iy)) }
			}
		case reflect.Float64:
			if kx == kindConst && ky == kindConst {
				v := basic.Float64(vx) > basic.Float64(vy)
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := basic.Float64(vx)
				return func(fr *frame) { fr.setReg(ir, x > fr.float64(iy)) }
			} else if ky == kindConst {
				y := basic.Float64(vy)
				return func(fr *frame) { fr.setReg(ir, fr.float64(ix) > y) }
			} else {
				return func(fr *frame) { fr.setReg(ir, fr.float64(ix) > fr.float64(iy)) }
			}
		case reflect.String:
			if kx == kindConst && ky == kindConst {
				v := basic.String(vx) > basic.String(vy)
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := basic.String(vx)
				return func(fr *frame) { fr.setReg(ir, x > fr.string(iy)) }
			} else if ky == kindConst {
				y := basic.String(vy)
				return func(fr *frame) { fr.setReg(ir, fr.string(ix) > y) }
			} else {
				return func(fr *frame) { fr.setReg(ir, fr.string(ix) > fr.string(iy)) }
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
			if kx == kindConst && ky == kindConst {
				v := vx.(int) >= vy.(int)
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := vx.(int)
				return func(fr *frame) { fr.setReg(ir, x >= fr.reg(iy).(int)) }
			} else if ky == kindConst {
				y := vy.(int)
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int) >= y) }
			} else {
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int) >= fr.reg(iy).(int)) }
			}
		case reflect.Int8:
			if kx == kindConst && ky == kindConst {
				v := vx.(int8) >= vy.(int8)
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := vx.(int8)
				return func(fr *frame) { fr.setReg(ir, x >= fr.reg(iy).(int8)) }
			} else if ky == kindConst {
				y := vy.(int8)
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int8) >= y) }
			} else {
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int8) >= fr.reg(iy).(int8)) }
			}
		case reflect.Int16:
			if kx == kindConst && ky == kindConst {
				v := vx.(int16) >= vy.(int16)
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := vx.(int16)
				return func(fr *frame) { fr.setReg(ir, x >= fr.reg(iy).(int16)) }
			} else if ky == kindConst {
				y := vy.(int16)
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int16) >= y) }
			} else {
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int16) >= fr.reg(iy).(int16)) }
			}
		case reflect.Int32:
			if kx == kindConst && ky == kindConst {
				v := vx.(int32) >= vy.(int32)
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := vx.(int32)
				return func(fr *frame) { fr.setReg(ir, x >= fr.reg(iy).(int32)) }
			} else if ky == kindConst {
				y := vy.(int32)
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int32) >= y) }
			} else {
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int32) >= fr.reg(iy).(int32)) }
			}
		case reflect.Int64:
			if kx == kindConst && ky == kindConst {
				v := vx.(int64) >= vy.(int64)
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := vx.(int64)
				return func(fr *frame) { fr.setReg(ir, x >= fr.reg(iy).(int64)) }
			} else if ky == kindConst {
				y := vy.(int64)
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int64) >= y) }
			} else {
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(int64) >= fr.reg(iy).(int64)) }
			}
		case reflect.Uint:
			if kx == kindConst && ky == kindConst {
				v := vx.(uint) >= vy.(uint)
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := vx.(uint)
				return func(fr *frame) { fr.setReg(ir, x >= fr.reg(iy).(uint)) }
			} else if ky == kindConst {
				y := vy.(uint)
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint) >= y) }
			} else {
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint) >= fr.reg(iy).(uint)) }
			}
		case reflect.Uint8:
			if kx == kindConst && ky == kindConst {
				v := vx.(uint8) >= vy.(uint8)
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := vx.(uint8)
				return func(fr *frame) { fr.setReg(ir, x >= fr.reg(iy).(uint8)) }
			} else if ky == kindConst {
				y := vy.(uint8)
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint8) >= y) }
			} else {
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint8) >= fr.reg(iy).(uint8)) }
			}
		case reflect.Uint16:
			if kx == kindConst && ky == kindConst {
				v := vx.(uint16) >= vy.(uint16)
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := vx.(uint16)
				return func(fr *frame) { fr.setReg(ir, x >= fr.reg(iy).(uint16)) }
			} else if ky == kindConst {
				y := vy.(uint16)
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint16) >= y) }
			} else {
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint16) >= fr.reg(iy).(uint16)) }
			}
		case reflect.Uint32:
			if kx == kindConst && ky == kindConst {
				v := vx.(uint32) >= vy.(uint32)
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := vx.(uint32)
				return func(fr *frame) { fr.setReg(ir, x >= fr.reg(iy).(uint32)) }
			} else if ky == kindConst {
				y := vy.(uint32)
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint32) >= y) }
			} else {
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint32) >= fr.reg(iy).(uint32)) }
			}
		case reflect.Uint64:
			if kx == kindConst && ky == kindConst {
				v := vx.(uint64) >= vy.(uint64)
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := vx.(uint64)
				return func(fr *frame) { fr.setReg(ir, x >= fr.reg(iy).(uint64)) }
			} else if ky == kindConst {
				y := vy.(uint64)
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint64) >= y) }
			} else {
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uint64) >= fr.reg(iy).(uint64)) }
			}
		case reflect.Uintptr:
			if kx == kindConst && ky == kindConst {
				v := vx.(uintptr) >= vy.(uintptr)
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := vx.(uintptr)
				return func(fr *frame) { fr.setReg(ir, x >= fr.reg(iy).(uintptr)) }
			} else if ky == kindConst {
				y := vy.(uintptr)
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uintptr) >= y) }
			} else {
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(uintptr) >= fr.reg(iy).(uintptr)) }
			}
		case reflect.Float32:
			if kx == kindConst && ky == kindConst {
				v := vx.(float32) >= vy.(float32)
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := vx.(float32)
				return func(fr *frame) { fr.setReg(ir, x >= fr.reg(iy).(float32)) }
			} else if ky == kindConst {
				y := vy.(float32)
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(float32) >= y) }
			} else {
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(float32) >= fr.reg(iy).(float32)) }
			}
		case reflect.Float64:
			if kx == kindConst && ky == kindConst {
				v := vx.(float64) >= vy.(float64)
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := vx.(float64)
				return func(fr *frame) { fr.setReg(ir, x >= fr.reg(iy).(float64)) }
			} else if ky == kindConst {
				y := vy.(float64)
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(float64) >= y) }
			} else {
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(float64) >= fr.reg(iy).(float64)) }
			}
		case reflect.String:
			if kx == kindConst && ky == kindConst {
				v := vx.(string) >= vy.(string)
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := vx.(string)
				return func(fr *frame) { fr.setReg(ir, x >= fr.reg(iy).(string)) }
			} else if ky == kindConst {
				y := vy.(string)
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(string) >= y) }
			} else {
				return func(fr *frame) { fr.setReg(ir, fr.reg(ix).(string) >= fr.reg(iy).(string)) }
			}
		}
	} else {
		switch typ.Kind() {
		case reflect.Int:
			if kx == kindConst && ky == kindConst {
				v := basic.Int(vx) >= basic.Int(vy)
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := basic.Int(vx)
				return func(fr *frame) { fr.setReg(ir, x >= fr.int(iy)) }
			} else if ky == kindConst {
				y := basic.Int(vy)
				return func(fr *frame) { fr.setReg(ir, fr.int(ix) >= y) }
			} else {
				return func(fr *frame) { fr.setReg(ir, fr.int(ix) >= fr.int(iy)) }
			}
		case reflect.Int8:
			if kx == kindConst && ky == kindConst {
				v := basic.Int8(vx) >= basic.Int8(vy)
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := basic.Int8(vx)
				return func(fr *frame) { fr.setReg(ir, x >= fr.int8(iy)) }
			} else if ky == kindConst {
				y := basic.Int8(vy)
				return func(fr *frame) { fr.setReg(ir, fr.int8(ix) >= y) }
			} else {
				return func(fr *frame) { fr.setReg(ir, fr.int8(ix) >= fr.int8(iy)) }
			}
		case reflect.Int16:
			if kx == kindConst && ky == kindConst {
				v := basic.Int16(vx) >= basic.Int16(vy)
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := basic.Int16(vx)
				return func(fr *frame) { fr.setReg(ir, x >= fr.int16(iy)) }
			} else if ky == kindConst {
				y := basic.Int16(vy)
				return func(fr *frame) { fr.setReg(ir, fr.int16(ix) >= y) }
			} else {
				return func(fr *frame) { fr.setReg(ir, fr.int16(ix) >= fr.int16(iy)) }
			}
		case reflect.Int32:
			if kx == kindConst && ky == kindConst {
				v := basic.Int32(vx) >= basic.Int32(vy)
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := basic.Int32(vx)
				return func(fr *frame) { fr.setReg(ir, x >= fr.int32(iy)) }
			} else if ky == kindConst {
				y := basic.Int32(vy)
				return func(fr *frame) { fr.setReg(ir, fr.int32(ix) >= y) }
			} else {
				return func(fr *frame) { fr.setReg(ir, fr.int32(ix) >= fr.int32(iy)) }
			}
		case reflect.Int64:
			if kx == kindConst && ky == kindConst {
				v := basic.Int64(vx) >= basic.Int64(vy)
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := basic.Int64(vx)
				return func(fr *frame) { fr.setReg(ir, x >= fr.int64(iy)) }
			} else if ky == kindConst {
				y := basic.Int64(vy)
				return func(fr *frame) { fr.setReg(ir, fr.int64(ix) >= y) }
			} else {
				return func(fr *frame) { fr.setReg(ir, fr.int64(ix) >= fr.int64(iy)) }
			}
		case reflect.Uint:
			if kx == kindConst && ky == kindConst {
				v := basic.Uint(vx) >= basic.Uint(vy)
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := basic.Uint(vx)
				return func(fr *frame) { fr.setReg(ir, x >= fr.uint(iy)) }
			} else if ky == kindConst {
				y := basic.Uint(vy)
				return func(fr *frame) { fr.setReg(ir, fr.uint(ix) >= y) }
			} else {
				return func(fr *frame) { fr.setReg(ir, fr.uint(ix) >= fr.uint(iy)) }
			}
		case reflect.Uint8:
			if kx == kindConst && ky == kindConst {
				v := basic.Uint8(vx) >= basic.Uint8(vy)
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := basic.Uint8(vx)
				return func(fr *frame) { fr.setReg(ir, x >= fr.uint8(iy)) }
			} else if ky == kindConst {
				y := basic.Uint8(vy)
				return func(fr *frame) { fr.setReg(ir, fr.uint8(ix) >= y) }
			} else {
				return func(fr *frame) { fr.setReg(ir, fr.uint8(ix) >= fr.uint8(iy)) }
			}
		case reflect.Uint16:
			if kx == kindConst && ky == kindConst {
				v := basic.Uint16(vx) >= basic.Uint16(vy)
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := basic.Uint16(vx)
				return func(fr *frame) { fr.setReg(ir, x >= fr.uint16(iy)) }
			} else if ky == kindConst {
				y := basic.Uint16(vy)
				return func(fr *frame) { fr.setReg(ir, fr.uint16(ix) >= y) }
			} else {
				return func(fr *frame) { fr.setReg(ir, fr.uint16(ix) >= fr.uint16(iy)) }
			}
		case reflect.Uint32:
			if kx == kindConst && ky == kindConst {
				v := basic.Uint32(vx) >= basic.Uint32(vy)
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := basic.Uint32(vx)
				return func(fr *frame) { fr.setReg(ir, x >= fr.uint32(iy)) }
			} else if ky == kindConst {
				y := basic.Uint32(vy)
				return func(fr *frame) { fr.setReg(ir, fr.uint32(ix) >= y) }
			} else {
				return func(fr *frame) { fr.setReg(ir, fr.uint32(ix) >= fr.uint32(iy)) }
			}
		case reflect.Uint64:
			if kx == kindConst && ky == kindConst {
				v := basic.Uint64(vx) >= basic.Uint64(vy)
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := basic.Uint64(vx)
				return func(fr *frame) { fr.setReg(ir, x >= fr.uint64(iy)) }
			} else if ky == kindConst {
				y := basic.Uint64(vy)
				return func(fr *frame) { fr.setReg(ir, fr.uint64(ix) >= y) }
			} else {
				return func(fr *frame) { fr.setReg(ir, fr.uint64(ix) >= fr.uint64(iy)) }
			}
		case reflect.Uintptr:
			if kx == kindConst && ky == kindConst {
				v := basic.Uintptr(vx) >= basic.Uintptr(vy)
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := basic.Uintptr(vx)
				return func(fr *frame) { fr.setReg(ir, x >= fr.uintptr(iy)) }
			} else if ky == kindConst {
				y := basic.Uintptr(vy)
				return func(fr *frame) { fr.setReg(ir, fr.uintptr(ix) >= y) }
			} else {
				return func(fr *frame) { fr.setReg(ir, fr.uintptr(ix) >= fr.uintptr(iy)) }
			}
		case reflect.Float32:
			if kx == kindConst && ky == kindConst {
				v := basic.Float32(vx) >= basic.Float32(vy)
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := basic.Float32(vx)
				return func(fr *frame) { fr.setReg(ir, x >= fr.float32(iy)) }
			} else if ky == kindConst {
				y := basic.Float32(vy)
				return func(fr *frame) { fr.setReg(ir, fr.float32(ix) >= y) }
			} else {
				return func(fr *frame) { fr.setReg(ir, fr.float32(ix) >= fr.float32(iy)) }
			}
		case reflect.Float64:
			if kx == kindConst && ky == kindConst {
				v := basic.Float64(vx) >= basic.Float64(vy)
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := basic.Float64(vx)
				return func(fr *frame) { fr.setReg(ir, x >= fr.float64(iy)) }
			} else if ky == kindConst {
				y := basic.Float64(vy)
				return func(fr *frame) { fr.setReg(ir, fr.float64(ix) >= y) }
			} else {
				return func(fr *frame) { fr.setReg(ir, fr.float64(ix) >= fr.float64(iy)) }
			}
		case reflect.String:
			if kx == kindConst && ky == kindConst {
				v := basic.String(vx) >= basic.String(vy)
				return func(fr *frame) { fr.setReg(ir, v) }
			} else if kx == kindConst {
				x := basic.String(vx)
				return func(fr *frame) { fr.setReg(ir, x >= fr.string(iy)) }
			} else if ky == kindConst {
				y := basic.String(vy)
				return func(fr *frame) { fr.setReg(ir, fr.string(ix) >= y) }
			} else {
				return func(fr *frame) { fr.setReg(ir, fr.string(ix) >= fr.string(iy)) }
			}
		}
	}
	panic("unreachable")
}
