package gossa

import (
	"reflect"
	"unsafe"

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
	kind := typ.Kind()
	switch kind {
	case reflect.Ptr, reflect.Chan, reflect.Map, reflect.Func, reflect.Slice:
		t := basic.TypeOfType(typ)
		return func(fr *frame) {
			x := fr.reg(ix)
			fr.setReg(ir, basic.ConvertPtr(t, x))
		}
	case reflect.Struct, reflect.Array:
		t := basic.TypeOfType(typ)
		return func(fr *frame) {
			x := fr.reg(ix)
			fr.setReg(ir, basic.ConvertDirect(t, x))
		}
	case reflect.Interface:
		return func(fr *frame) {
			x := fr.reg(ix)
			if x == nil {
				fr.setReg(ir, reflect.New(typ).Elem().Interface())
			} else {
				fr.setReg(ir, reflect.ValueOf(x).Convert(typ).Interface())
			}
		}
	case reflect.Bool:
		if typ.PkgPath() == "" {
			return func(fr *frame) {
				x := fr.reg(ix)
				fr.setReg(ir, basic.Bool(x))
			}
		} else {
			t := basic.TypeOfType(typ)
			return func(fr *frame) {
				x := fr.reg(ix)
				fr.setReg(ir, basic.ConvertBool(t, x))
			}
		}
	case reflect.Int:
		if typ.PkgPath() == "" {
			return func(fr *frame) {
				x := fr.reg(ix)
				fr.setReg(ir, basic.Int(x))
			}
		} else {
			t := basic.TypeOfType(typ)
			return func(fr *frame) {
				x := fr.reg(ix)
				fr.setReg(ir, basic.ConvertInt(t, x))
			}
		}
	case reflect.Int8:
		if typ.PkgPath() == "" {
			return func(fr *frame) {
				x := fr.reg(ix)
				fr.setReg(ir, basic.Int8(x))
			}
		} else {
			t := basic.TypeOfType(typ)
			return func(fr *frame) {
				x := fr.reg(ix)
				fr.setReg(ir, basic.ConvertInt8(t, x))
			}
		}
	case reflect.Int16:
		if typ.PkgPath() == "" {
			return func(fr *frame) {
				x := fr.reg(ix)
				fr.setReg(ir, basic.Int16(x))
			}
		} else {
			t := basic.TypeOfType(typ)
			return func(fr *frame) {
				x := fr.reg(ix)
				fr.setReg(ir, basic.ConvertInt16(t, x))
			}
		}
	case reflect.Int32:
		if typ.PkgPath() == "" {
			return func(fr *frame) {
				x := fr.reg(ix)
				fr.setReg(ir, basic.Int32(x))
			}
		} else {
			t := basic.TypeOfType(typ)
			return func(fr *frame) {
				x := fr.reg(ix)
				fr.setReg(ir, basic.ConvertInt32(t, x))
			}
		}
	case reflect.Int64:
		if typ.PkgPath() == "" {
			return func(fr *frame) {
				x := fr.reg(ix)
				fr.setReg(ir, basic.Int64(x))
			}
		} else {
			t := basic.TypeOfType(typ)
			return func(fr *frame) {
				x := fr.reg(ix)
				fr.setReg(ir, basic.ConvertInt64(t, x))
			}
		}
	case reflect.Uint:
		if typ.PkgPath() == "" {
			return func(fr *frame) {
				x := fr.reg(ix)
				fr.setReg(ir, basic.Uint(x))
			}
		} else {
			t := basic.TypeOfType(typ)
			return func(fr *frame) {
				x := fr.reg(ix)
				fr.setReg(ir, basic.ConvertUint(t, x))
			}
		}
	case reflect.Uint8:
		if typ.PkgPath() == "" {
			return func(fr *frame) {
				x := fr.reg(ix)
				fr.setReg(ir, basic.Uint8(x))
			}
		} else {
			t := basic.TypeOfType(typ)
			return func(fr *frame) {
				x := fr.reg(ix)
				fr.setReg(ir, basic.ConvertUint8(t, x))
			}
		}
	case reflect.Uint16:
		if typ.PkgPath() == "" {
			return func(fr *frame) {
				x := fr.reg(ix)
				fr.setReg(ir, basic.Uint16(x))
			}
		} else {
			t := basic.TypeOfType(typ)
			return func(fr *frame) {
				x := fr.reg(ix)
				fr.setReg(ir, basic.ConvertUint16(t, x))
			}
		}
	case reflect.Uint32:
		if typ.PkgPath() == "" {
			return func(fr *frame) {
				x := fr.reg(ix)
				fr.setReg(ir, basic.Uint32(x))
			}
		} else {
			t := basic.TypeOfType(typ)
			return func(fr *frame) {
				x := fr.reg(ix)
				fr.setReg(ir, basic.ConvertUint32(t, x))
			}
		}
	case reflect.Uint64:
		if typ.PkgPath() == "" {
			return func(fr *frame) {
				x := fr.reg(ix)
				fr.setReg(ir, basic.Uint64(x))
			}
		} else {
			t := basic.TypeOfType(typ)
			return func(fr *frame) {
				x := fr.reg(ix)
				fr.setReg(ir, basic.ConvertUint64(t, x))
			}
		}
	case reflect.Uintptr:
		if typ.PkgPath() == "" {
			return func(fr *frame) {
				x := fr.reg(ix)
				fr.setReg(ir, basic.Uintptr(x))
			}
		} else {
			t := basic.TypeOfType(typ)
			return func(fr *frame) {
				x := fr.reg(ix)
				fr.setReg(ir, basic.ConvertUintptr(t, x))
			}
		}
	case reflect.Float32:
		if typ.PkgPath() == "" {
			return func(fr *frame) {
				x := fr.reg(ix)
				fr.setReg(ir, basic.Float32(x))
			}
		} else {
			t := basic.TypeOfType(typ)
			return func(fr *frame) {
				x := fr.reg(ix)
				fr.setReg(ir, basic.ConvertFloat32(t, x))
			}
		}
	case reflect.Float64:
		if typ.PkgPath() == "" {
			return func(fr *frame) {
				x := fr.reg(ix)
				fr.setReg(ir, basic.Float64(x))
			}
		} else {
			t := basic.TypeOfType(typ)
			return func(fr *frame) {
				x := fr.reg(ix)
				fr.setReg(ir, basic.ConvertFloat64(t, x))
			}
		}
	case reflect.Complex64:
		if typ.PkgPath() == "" {
			return func(fr *frame) {
				x := fr.reg(ix)
				fr.setReg(ir, basic.Complex64(x))
			}
		} else {
			t := basic.TypeOfType(typ)
			return func(fr *frame) {
				x := fr.reg(ix)
				fr.setReg(ir, basic.ConvertComplex64(t, x))
			}
		}
	case reflect.Complex128:
		if typ.PkgPath() == "" {
			return func(fr *frame) {
				x := fr.reg(ix)
				fr.setReg(ir, basic.Complex128(x))
			}
		} else {
			t := basic.TypeOfType(typ)
			return func(fr *frame) {
				x := fr.reg(ix)
				fr.setReg(ir, basic.ConvertComplex128(t, x))
			}
		}
	case reflect.String:
		if typ.PkgPath() == "" {
			return func(fr *frame) {
				x := fr.reg(ix)
				fr.setReg(ir, basic.String(x))
			}
		} else {
			t := basic.TypeOfType(typ)
			return func(fr *frame) {
				x := fr.reg(ix)
				fr.setReg(ir, basic.ConvertString(t, x))
			}
		}
	}
	panic("unreachable")
}

func makeConvertInstr(pfn *function, interp *Interp, instr *ssa.Convert) func(fr *frame) {
	typ := interp.preToType(instr.Type())
	xtyp := interp.preToType(instr.X.Type())
	vk := xtyp.Kind()
	ir := pfn.regIndex(instr)
	ix, kx, vx := pfn.regIndex3(instr.X)
	switch typ.Kind() {
	case reflect.UnsafePointer:
		if vk == reflect.Uintptr {
			return func(fr *frame) {
				v := reflect.ValueOf(fr.reg(ix))
				fr.setReg(ir, toUnsafePointer(v))
			}
		} else if vk == reflect.Ptr {
			return func(fr *frame) {
				v := reflect.ValueOf(fr.reg(ix))
				fr.setReg(ir, unsafe.Pointer(v.Pointer()))
			}
		}
	case reflect.Uintptr:
		if vk == reflect.UnsafePointer {
			return func(fr *frame) {
				v := reflect.ValueOf(fr.reg(ix))
				fr.setReg(ir, v.Pointer())
			}
		}
	case reflect.Ptr:
		if vk == reflect.UnsafePointer {
			return func(fr *frame) {
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
					return func(fr *frame) {
						v := reflect.ValueOf(fr.reg(ix))
						dst := reflect.New(typ).Elem()
						dst.SetBytes([]byte(v.String()))
						fr.setReg(ir, dst.Interface())
					}
				}
			case reflect.Int32:
				if elem.PkgPath() != "" {
					return func(fr *frame) {
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
					return func(fr *frame) {
						v := reflect.ValueOf(fr.reg(ix))
						v = reflect.ValueOf(string(v.Bytes()))
						fr.setReg(ir, v.Convert(typ).Interface())
					}
				}
			case reflect.Int32:
				if elem.PkgPath() != "" {
					return func(fr *frame) {
						v := reflect.ValueOf(fr.reg(ix))
						v = reflect.ValueOf(*(*[]rune)(((*reflectValue)(unsafe.Pointer(&v))).ptr))
						fr.setReg(ir, v.Convert(typ).Interface())
					}
				}
			}
		}
	}
	if kx.isStatic() {
		v := reflect.ValueOf(vx)
		return func(fr *frame) {
			fr.setReg(ir, v.Convert(typ).Interface())
		}
	}
	return func(fr *frame) {
		v := reflect.ValueOf(fr.reg(ix))
		fr.setReg(ir, v.Convert(typ).Interface())
	}
}
