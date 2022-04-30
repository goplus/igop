package gossa

import (
	"reflect"

	"github.com/goplus/gossa/internal/basic"
	"golang.org/x/tools/go/ssa"
)

func makeTypeChangeInstr(pfn *function, instr *ssa.ChangeType) func(fr *frame) {
	typ := pfn.Interp.preToType(instr.Type())
	ir := pfn.regIndex(instr)
	ix, kx, vx := pfn.regIndex3(instr.X)
	if kx.isStatic() {
		var v interface{}
		if vx == nil {
			v = reflect.New(typ).Elem().Interface()
		} else {
			v = reflect.ValueOf(vx).Convert(typ).Interface()
		}
		return func(fr *frame) {
			fr.setReg(ir, v)
		}
	}
	xtyp := pfn.Interp.preToType(instr.X.Type())
	isBasic := typ.PkgPath() == ""
	if kind := typ.Kind(); kind == xtyp.Kind() {
		if kind == reflect.Ptr {
			t := basic.TypeOfType(typ)
			return func(fr *frame) {
				x := fr.reg(ix)
				fr.setReg(ir, basic.ConvertPtr(t, x))
			}
		}
		if isBasic {
			switch kind {
			case reflect.Bool:
				return func(fr *frame) {
					x := fr.reg(ix)
					fr.setReg(ir, basic.Bool(x))
				}
			case reflect.Int:
				return func(fr *frame) {
					x := fr.reg(ix)
					fr.setReg(ir, basic.Int(x))
				}
			case reflect.Int8:
				return func(fr *frame) {
					x := fr.reg(ix)
					fr.setReg(ir, basic.Int8(x))
				}
			case reflect.Int16:
				return func(fr *frame) {
					x := fr.reg(ix)
					fr.setReg(ir, basic.Int16(x))
				}
			case reflect.Int32:
				return func(fr *frame) {
					x := fr.reg(ix)
					fr.setReg(ir, basic.Int32(x))
				}
			case reflect.Int64:
				return func(fr *frame) {
					x := fr.reg(ix)
					fr.setReg(ir, basic.Int64(x))
				}
			case reflect.Uint:
				return func(fr *frame) {
					x := fr.reg(ix)
					fr.setReg(ir, basic.Uint(x))
				}
			case reflect.Uint8:
				return func(fr *frame) {
					x := fr.reg(ix)
					fr.setReg(ir, basic.Uint8(x))
				}
			case reflect.Uint16:
				return func(fr *frame) {
					x := fr.reg(ix)
					fr.setReg(ir, basic.Uint16(x))
				}
			case reflect.Uint32:
				return func(fr *frame) {
					x := fr.reg(ix)
					fr.setReg(ir, basic.Uint32(x))
				}
			case reflect.Uint64:
				return func(fr *frame) {
					x := fr.reg(ix)
					fr.setReg(ir, basic.Uint64(x))
				}
			case reflect.Uintptr:
				return func(fr *frame) {
					x := fr.reg(ix)
					fr.setReg(ir, basic.Uintptr(x))
				}
			case reflect.Float32:
				return func(fr *frame) {
					x := fr.reg(ix)
					fr.setReg(ir, basic.Float32(x))
				}
			case reflect.Float64:
				return func(fr *frame) {
					x := fr.reg(ix)
					fr.setReg(ir, basic.Float64(x))
				}
			case reflect.Complex64:
				return func(fr *frame) {
					x := fr.reg(ix)
					fr.setReg(ir, basic.Complex64(x))
				}
			case reflect.Complex128:
				return func(fr *frame) {
					x := fr.reg(ix)
					fr.setReg(ir, basic.Complex128(x))
				}
			case reflect.String:
				return func(fr *frame) {
					x := fr.reg(ix)
					fr.setReg(ir, basic.String(x))
				}
			}
		} else {
			t := basic.TypeOfType(typ)
			switch kind {
			case reflect.Bool:
				return func(fr *frame) {
					x := fr.reg(ix)
					fr.setReg(ir, basic.ConvertBool(t, x))
				}
			case reflect.Int:
				return func(fr *frame) {
					x := fr.reg(ix)
					fr.setReg(ir, basic.ConvertInt(t, x))
				}
			case reflect.Int8:
				return func(fr *frame) {
					x := fr.reg(ix)
					fr.setReg(ir, basic.ConvertInt8(t, x))
				}
			case reflect.Int16:
				return func(fr *frame) {
					x := fr.reg(ix)
					fr.setReg(ir, basic.ConvertInt16(t, x))
				}
			case reflect.Int32:
				return func(fr *frame) {
					x := fr.reg(ix)
					fr.setReg(ir, basic.ConvertInt32(t, x))
				}
			case reflect.Int64:
				return func(fr *frame) {
					x := fr.reg(ix)
					fr.setReg(ir, basic.ConvertInt64(t, x))
				}
			case reflect.Uint:
				return func(fr *frame) {
					x := fr.reg(ix)
					fr.setReg(ir, basic.ConvertUint(t, x))
				}
			case reflect.Uint8:
				return func(fr *frame) {
					x := fr.reg(ix)
					fr.setReg(ir, basic.ConvertUint8(t, x))
				}
			case reflect.Uint16:
				return func(fr *frame) {
					x := fr.reg(ix)
					fr.setReg(ir, basic.ConvertUint16(t, x))
				}
			case reflect.Uint32:
				return func(fr *frame) {
					x := fr.reg(ix)
					fr.setReg(ir, basic.ConvertUint32(t, x))
				}
			case reflect.Uint64:
				return func(fr *frame) {
					x := fr.reg(ix)
					fr.setReg(ir, basic.ConvertUint64(t, x))
				}
			case reflect.Uintptr:
				return func(fr *frame) {
					x := fr.reg(ix)
					fr.setReg(ir, basic.ConvertUintptr(t, x))
				}
			case reflect.Float32:
				return func(fr *frame) {
					x := fr.reg(ix)
					fr.setReg(ir, basic.ConvertFloat32(t, x))
				}
			case reflect.Float64:
				return func(fr *frame) {
					x := fr.reg(ix)
					fr.setReg(ir, basic.ConvertFloat64(t, x))
				}
			case reflect.Complex64:
				return func(fr *frame) {
					x := fr.reg(ix)
					fr.setReg(ir, basic.ConvertComplex64(t, x))
				}
			case reflect.Complex128:
				return func(fr *frame) {
					x := fr.reg(ix)
					fr.setReg(ir, basic.ConvertComplex128(t, x))
				}
			case reflect.String:
				return func(fr *frame) {
					x := fr.reg(ix)
					fr.setReg(ir, basic.ConvertString(t, x))
				}
			}
		}
	}
	return func(fr *frame) {
		x := fr.reg(ix)
		if x == nil {
			fr.setReg(ir, reflect.New(typ).Elem().Interface())
		} else {
			fr.setReg(ir, reflect.ValueOf(x).Convert(typ).Interface())
		}
	}
}
