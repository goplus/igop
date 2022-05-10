package gossa

import (
	"reflect"

	"github.com/goplus/gossa/internal/basic"
	"golang.org/x/tools/go/ssa"
)

func makeBinOpSHL(pfn *function, instr *ssa.BinOp) func(fr *frame) {
	ir := pfn.regIndex(instr)
	ix, kx, vx := pfn.regIndex3(instr.X)
	iy, ky, vy := pfn.regIndex3(instr.Y)
	xtyp := pfn.Interp.preToType(instr.X.Type())
	ytyp := pfn.Interp.preToType(instr.Y.Type())
	xkind := xtyp.Kind()
	ykind := ytyp.Kind()
	if kx == kindConst && ky == kindConst {
		t := basic.TypeOfType(xtyp)
		switch xkind {
		case reflect.Int:
			x := basic.Int(vx)
			switch ykind {
			case reflect.Int:
				v := basic.Make(t, x<<basic.Int(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			case reflect.Int8:
				v := basic.Make(t, x<<basic.Int8(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			case reflect.Int16:
				v := basic.Make(t, x<<basic.Int16(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			case reflect.Int32:
				v := basic.Make(t, x<<basic.Int32(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			case reflect.Int64:
				v := basic.Make(t, x<<basic.Int64(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			case reflect.Uint:
				v := basic.Make(t, x<<basic.Uint(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			case reflect.Uint8:
				v := basic.Make(t, x<<basic.Uint8(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			case reflect.Uint16:
				v := basic.Make(t, x<<basic.Uint16(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			case reflect.Uint32:
				v := basic.Make(t, x<<basic.Uint32(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			case reflect.Uint64:
				v := basic.Make(t, x<<basic.Uint64(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			case reflect.Uintptr:
				v := basic.Make(t, x<<basic.Uintptr(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			}
		case reflect.Int8:
			x := basic.Int8(vx)
			switch ykind {
			case reflect.Int:
				v := basic.Make(t, x<<basic.Int(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			case reflect.Int8:
				v := basic.Make(t, x<<basic.Int8(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			case reflect.Int16:
				v := basic.Make(t, x<<basic.Int16(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			case reflect.Int32:
				v := basic.Make(t, x<<basic.Int32(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			case reflect.Int64:
				v := basic.Make(t, x<<basic.Int64(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			case reflect.Uint:
				v := basic.Make(t, x<<basic.Uint(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			case reflect.Uint8:
				v := basic.Make(t, x<<basic.Uint8(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			case reflect.Uint16:
				v := basic.Make(t, x<<basic.Uint16(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			case reflect.Uint32:
				v := basic.Make(t, x<<basic.Uint32(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			case reflect.Uint64:
				v := basic.Make(t, x<<basic.Uint64(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			case reflect.Uintptr:
				v := basic.Make(t, x<<basic.Uintptr(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			}
		case reflect.Int16:
			x := basic.Int16(vx)
			switch ykind {
			case reflect.Int:
				v := basic.Make(t, x<<basic.Int(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			case reflect.Int8:
				v := basic.Make(t, x<<basic.Int8(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			case reflect.Int16:
				v := basic.Make(t, x<<basic.Int16(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			case reflect.Int32:
				v := basic.Make(t, x<<basic.Int32(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			case reflect.Int64:
				v := basic.Make(t, x<<basic.Int64(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			case reflect.Uint:
				v := basic.Make(t, x<<basic.Uint(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			case reflect.Uint8:
				v := basic.Make(t, x<<basic.Uint8(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			case reflect.Uint16:
				v := basic.Make(t, x<<basic.Uint16(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			case reflect.Uint32:
				v := basic.Make(t, x<<basic.Uint32(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			case reflect.Uint64:
				v := basic.Make(t, x<<basic.Uint64(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			case reflect.Uintptr:
				v := basic.Make(t, x<<basic.Uintptr(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			}
		case reflect.Int32:
			x := basic.Int32(vx)
			switch ykind {
			case reflect.Int:
				v := basic.Make(t, x<<basic.Int(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			case reflect.Int8:
				v := basic.Make(t, x<<basic.Int8(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			case reflect.Int16:
				v := basic.Make(t, x<<basic.Int16(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			case reflect.Int32:
				v := basic.Make(t, x<<basic.Int32(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			case reflect.Int64:
				v := basic.Make(t, x<<basic.Int64(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			case reflect.Uint:
				v := basic.Make(t, x<<basic.Uint(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			case reflect.Uint8:
				v := basic.Make(t, x<<basic.Uint8(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			case reflect.Uint16:
				v := basic.Make(t, x<<basic.Uint16(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			case reflect.Uint32:
				v := basic.Make(t, x<<basic.Uint32(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			case reflect.Uint64:
				v := basic.Make(t, x<<basic.Uint64(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			case reflect.Uintptr:
				v := basic.Make(t, x<<basic.Uintptr(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			}
		case reflect.Int64:
			x := basic.Int64(vx)
			switch ykind {
			case reflect.Int:
				v := basic.Make(t, x<<basic.Int(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			case reflect.Int8:
				v := basic.Make(t, x<<basic.Int8(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			case reflect.Int16:
				v := basic.Make(t, x<<basic.Int16(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			case reflect.Int32:
				v := basic.Make(t, x<<basic.Int32(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			case reflect.Int64:
				v := basic.Make(t, x<<basic.Int64(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			case reflect.Uint:
				v := basic.Make(t, x<<basic.Uint(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			case reflect.Uint8:
				v := basic.Make(t, x<<basic.Uint8(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			case reflect.Uint16:
				v := basic.Make(t, x<<basic.Uint16(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			case reflect.Uint32:
				v := basic.Make(t, x<<basic.Uint32(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			case reflect.Uint64:
				v := basic.Make(t, x<<basic.Uint64(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			case reflect.Uintptr:
				v := basic.Make(t, x<<basic.Uintptr(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			}
		case reflect.Uint:
			x := basic.Uint(vx)
			switch ykind {
			case reflect.Int:
				v := basic.Make(t, x<<basic.Int(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			case reflect.Int8:
				v := basic.Make(t, x<<basic.Int8(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			case reflect.Int16:
				v := basic.Make(t, x<<basic.Int16(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			case reflect.Int32:
				v := basic.Make(t, x<<basic.Int32(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			case reflect.Int64:
				v := basic.Make(t, x<<basic.Int64(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			case reflect.Uint:
				v := basic.Make(t, x<<basic.Uint(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			case reflect.Uint8:
				v := basic.Make(t, x<<basic.Uint8(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			case reflect.Uint16:
				v := basic.Make(t, x<<basic.Uint16(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			case reflect.Uint32:
				v := basic.Make(t, x<<basic.Uint32(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			case reflect.Uint64:
				v := basic.Make(t, x<<basic.Uint64(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			case reflect.Uintptr:
				v := basic.Make(t, x<<basic.Uintptr(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			}
		case reflect.Uint8:
			x := basic.Uint8(vx)
			switch ykind {
			case reflect.Int:
				v := basic.Make(t, x<<basic.Int(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			case reflect.Int8:
				v := basic.Make(t, x<<basic.Int8(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			case reflect.Int16:
				v := basic.Make(t, x<<basic.Int16(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			case reflect.Int32:
				v := basic.Make(t, x<<basic.Int32(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			case reflect.Int64:
				v := basic.Make(t, x<<basic.Int64(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			case reflect.Uint:
				v := basic.Make(t, x<<basic.Uint(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			case reflect.Uint8:
				v := basic.Make(t, x<<basic.Uint8(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			case reflect.Uint16:
				v := basic.Make(t, x<<basic.Uint16(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			case reflect.Uint32:
				v := basic.Make(t, x<<basic.Uint32(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			case reflect.Uint64:
				v := basic.Make(t, x<<basic.Uint64(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			case reflect.Uintptr:
				v := basic.Make(t, x<<basic.Uintptr(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			}
		case reflect.Uint16:
			x := basic.Uint16(vx)
			switch ykind {
			case reflect.Int:
				v := basic.Make(t, x<<basic.Int(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			case reflect.Int8:
				v := basic.Make(t, x<<basic.Int8(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			case reflect.Int16:
				v := basic.Make(t, x<<basic.Int16(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			case reflect.Int32:
				v := basic.Make(t, x<<basic.Int32(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			case reflect.Int64:
				v := basic.Make(t, x<<basic.Int64(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			case reflect.Uint:
				v := basic.Make(t, x<<basic.Uint(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			case reflect.Uint8:
				v := basic.Make(t, x<<basic.Uint8(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			case reflect.Uint16:
				v := basic.Make(t, x<<basic.Uint16(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			case reflect.Uint32:
				v := basic.Make(t, x<<basic.Uint32(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			case reflect.Uint64:
				v := basic.Make(t, x<<basic.Uint64(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			case reflect.Uintptr:
				v := basic.Make(t, x<<basic.Uintptr(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			}
		case reflect.Uint32:
			x := basic.Uint32(vx)
			switch ykind {
			case reflect.Int:
				v := basic.Make(t, x<<basic.Int(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			case reflect.Int8:
				v := basic.Make(t, x<<basic.Int8(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			case reflect.Int16:
				v := basic.Make(t, x<<basic.Int16(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			case reflect.Int32:
				v := basic.Make(t, x<<basic.Int32(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			case reflect.Int64:
				v := basic.Make(t, x<<basic.Int64(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			case reflect.Uint:
				v := basic.Make(t, x<<basic.Uint(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			case reflect.Uint8:
				v := basic.Make(t, x<<basic.Uint8(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			case reflect.Uint16:
				v := basic.Make(t, x<<basic.Uint16(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			case reflect.Uint32:
				v := basic.Make(t, x<<basic.Uint32(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			case reflect.Uint64:
				v := basic.Make(t, x<<basic.Uint64(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			case reflect.Uintptr:
				v := basic.Make(t, x<<basic.Uintptr(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			}
		case reflect.Uint64:
			x := basic.Uint64(vx)
			switch ykind {
			case reflect.Int:
				v := basic.Make(t, x<<basic.Int(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			case reflect.Int8:
				v := basic.Make(t, x<<basic.Int8(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			case reflect.Int16:
				v := basic.Make(t, x<<basic.Int16(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			case reflect.Int32:
				v := basic.Make(t, x<<basic.Int32(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			case reflect.Int64:
				v := basic.Make(t, x<<basic.Int64(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			case reflect.Uint:
				v := basic.Make(t, x<<basic.Uint(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			case reflect.Uint8:
				v := basic.Make(t, x<<basic.Uint8(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			case reflect.Uint16:
				v := basic.Make(t, x<<basic.Uint16(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			case reflect.Uint32:
				v := basic.Make(t, x<<basic.Uint32(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			case reflect.Uint64:
				v := basic.Make(t, x<<basic.Uint64(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			case reflect.Uintptr:
				v := basic.Make(t, x<<basic.Uintptr(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			}
		case reflect.Uintptr:
			x := basic.Uintptr(vx)
			switch ykind {
			case reflect.Int:
				v := basic.Make(t, x<<basic.Int(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			case reflect.Int8:
				v := basic.Make(t, x<<basic.Int8(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			case reflect.Int16:
				v := basic.Make(t, x<<basic.Int16(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			case reflect.Int32:
				v := basic.Make(t, x<<basic.Int32(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			case reflect.Int64:
				v := basic.Make(t, x<<basic.Int64(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			case reflect.Uint:
				v := basic.Make(t, x<<basic.Uint(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			case reflect.Uint8:
				v := basic.Make(t, x<<basic.Uint8(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			case reflect.Uint16:
				v := basic.Make(t, x<<basic.Uint16(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			case reflect.Uint32:
				v := basic.Make(t, x<<basic.Uint32(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			case reflect.Uint64:
				v := basic.Make(t, x<<basic.Uint64(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			case reflect.Uintptr:
				v := basic.Make(t, x<<basic.Uintptr(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			}
		}
	}
	if xtyp.PkgPath() == "" {

		switch xkind {
		case reflect.Int:
			if kx == kindConst {
				x := vx.(int)
				switch ykind {
				case reflect.Int:
					return func(fr *frame) {
						y := fr.int(iy)
						fr.setReg(ir, x<<y)
					}
				case reflect.Int8:
					return func(fr *frame) {
						y := fr.int8(iy)
						fr.setReg(ir, x<<y)
					}
				case reflect.Int16:
					return func(fr *frame) {
						y := fr.int16(iy)
						fr.setReg(ir, x<<y)
					}
				case reflect.Int32:
					return func(fr *frame) {
						y := fr.int32(iy)
						fr.setReg(ir, x<<y)
					}
				case reflect.Int64:
					return func(fr *frame) {
						y := fr.int64(iy)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint:
					return func(fr *frame) {
						y := fr.uint(iy)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint8:
					return func(fr *frame) {
						y := fr.uint8(iy)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint16:
					return func(fr *frame) {
						y := fr.uint16(iy)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint32:
					return func(fr *frame) {
						y := fr.uint32(iy)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint64:
					return func(fr *frame) {
						y := fr.uint64(iy)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uintptr:
					return func(fr *frame) {
						y := fr.uintptr(iy)
						fr.setReg(ir, x<<y)
					}
				}
			} else if ky == kindConst {
				switch ykind {
				case reflect.Int:
					y := basic.Int(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(int)
						fr.setReg(ir, x<<y)
					}
				case reflect.Int8:
					y := basic.Int8(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(int)
						fr.setReg(ir, x<<y)
					}
				case reflect.Int16:
					y := basic.Int16(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(int)
						fr.setReg(ir, x<<y)
					}
				case reflect.Int32:
					y := basic.Int32(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(int)
						fr.setReg(ir, x<<y)
					}
				case reflect.Int64:
					y := basic.Int64(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(int)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint:
					y := basic.Uint(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(int)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint8:
					y := basic.Uint8(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(int)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint16:
					y := basic.Uint16(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(int)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint32:
					y := basic.Uint32(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(int)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint64:
					y := basic.Uint64(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(int)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uintptr:
					y := basic.Uintptr(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(int)
						fr.setReg(ir, x<<y)
					}
				}
			} else {
				switch ykind {
				case reflect.Int:
					return func(fr *frame) {
						x := fr.int(ix)
						y := fr.int(iy)
						fr.setReg(ir, x<<y)
					}
				case reflect.Int8:
					return func(fr *frame) {
						x := fr.int(ix)
						y := fr.int8(iy)
						fr.setReg(ir, x<<y)
					}
				case reflect.Int16:
					return func(fr *frame) {
						x := fr.int(ix)
						y := fr.int16(iy)
						fr.setReg(ir, x<<y)
					}
				case reflect.Int32:
					return func(fr *frame) {
						x := fr.int(ix)
						y := fr.int32(iy)
						fr.setReg(ir, x<<y)
					}
				case reflect.Int64:
					return func(fr *frame) {
						x := fr.int(ix)
						y := fr.int64(iy)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint:
					return func(fr *frame) {
						x := fr.int(ix)
						y := fr.uint(iy)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint8:
					return func(fr *frame) {
						x := fr.int(ix)
						y := fr.uint8(iy)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint16:
					return func(fr *frame) {
						x := fr.int(ix)
						y := fr.uint16(iy)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint32:
					return func(fr *frame) {
						x := fr.int(ix)
						y := fr.uint32(iy)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint64:
					return func(fr *frame) {
						x := fr.int(ix)
						y := fr.uint64(iy)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uintptr:
					return func(fr *frame) {
						x := fr.int(ix)
						y := fr.uintptr(iy)
						fr.setReg(ir, x<<y)
					}
				}
			}
		case reflect.Int8:
			if kx == kindConst {
				x := vx.(int8)
				switch ykind {
				case reflect.Int:
					return func(fr *frame) {
						y := fr.int(iy)
						fr.setReg(ir, x<<y)
					}
				case reflect.Int8:
					return func(fr *frame) {
						y := fr.int8(iy)
						fr.setReg(ir, x<<y)
					}
				case reflect.Int16:
					return func(fr *frame) {
						y := fr.int16(iy)
						fr.setReg(ir, x<<y)
					}
				case reflect.Int32:
					return func(fr *frame) {
						y := fr.int32(iy)
						fr.setReg(ir, x<<y)
					}
				case reflect.Int64:
					return func(fr *frame) {
						y := fr.int64(iy)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint:
					return func(fr *frame) {
						y := fr.uint(iy)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint8:
					return func(fr *frame) {
						y := fr.uint8(iy)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint16:
					return func(fr *frame) {
						y := fr.uint16(iy)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint32:
					return func(fr *frame) {
						y := fr.uint32(iy)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint64:
					return func(fr *frame) {
						y := fr.uint64(iy)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uintptr:
					return func(fr *frame) {
						y := fr.uintptr(iy)
						fr.setReg(ir, x<<y)
					}
				}
			} else if ky == kindConst {
				switch ykind {
				case reflect.Int:
					y := basic.Int(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(int8)
						fr.setReg(ir, x<<y)
					}
				case reflect.Int8:
					y := basic.Int8(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(int8)
						fr.setReg(ir, x<<y)
					}
				case reflect.Int16:
					y := basic.Int16(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(int8)
						fr.setReg(ir, x<<y)
					}
				case reflect.Int32:
					y := basic.Int32(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(int8)
						fr.setReg(ir, x<<y)
					}
				case reflect.Int64:
					y := basic.Int64(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(int8)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint:
					y := basic.Uint(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(int8)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint8:
					y := basic.Uint8(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(int8)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint16:
					y := basic.Uint16(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(int8)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint32:
					y := basic.Uint32(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(int8)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint64:
					y := basic.Uint64(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(int8)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uintptr:
					y := basic.Uintptr(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(int8)
						fr.setReg(ir, x<<y)
					}
				}
			} else {
				switch ykind {
				case reflect.Int:
					return func(fr *frame) {
						x := fr.int8(ix)
						y := fr.int(iy)
						fr.setReg(ir, x<<y)
					}
				case reflect.Int8:
					return func(fr *frame) {
						x := fr.int8(ix)
						y := fr.int8(iy)
						fr.setReg(ir, x<<y)
					}
				case reflect.Int16:
					return func(fr *frame) {
						x := fr.int8(ix)
						y := fr.int16(iy)
						fr.setReg(ir, x<<y)
					}
				case reflect.Int32:
					return func(fr *frame) {
						x := fr.int8(ix)
						y := fr.int32(iy)
						fr.setReg(ir, x<<y)
					}
				case reflect.Int64:
					return func(fr *frame) {
						x := fr.int8(ix)
						y := fr.int64(iy)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint:
					return func(fr *frame) {
						x := fr.int8(ix)
						y := fr.uint(iy)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint8:
					return func(fr *frame) {
						x := fr.int8(ix)
						y := fr.uint8(iy)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint16:
					return func(fr *frame) {
						x := fr.int8(ix)
						y := fr.uint16(iy)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint32:
					return func(fr *frame) {
						x := fr.int8(ix)
						y := fr.uint32(iy)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint64:
					return func(fr *frame) {
						x := fr.int8(ix)
						y := fr.uint64(iy)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uintptr:
					return func(fr *frame) {
						x := fr.int8(ix)
						y := fr.uintptr(iy)
						fr.setReg(ir, x<<y)
					}
				}
			}
		case reflect.Int16:
			if kx == kindConst {
				x := vx.(int16)
				switch ykind {
				case reflect.Int:
					return func(fr *frame) {
						y := fr.int(iy)
						fr.setReg(ir, x<<y)
					}
				case reflect.Int8:
					return func(fr *frame) {
						y := fr.int8(iy)
						fr.setReg(ir, x<<y)
					}
				case reflect.Int16:
					return func(fr *frame) {
						y := fr.int16(iy)
						fr.setReg(ir, x<<y)
					}
				case reflect.Int32:
					return func(fr *frame) {
						y := fr.int32(iy)
						fr.setReg(ir, x<<y)
					}
				case reflect.Int64:
					return func(fr *frame) {
						y := fr.int64(iy)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint:
					return func(fr *frame) {
						y := fr.uint(iy)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint8:
					return func(fr *frame) {
						y := fr.uint8(iy)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint16:
					return func(fr *frame) {
						y := fr.uint16(iy)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint32:
					return func(fr *frame) {
						y := fr.uint32(iy)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint64:
					return func(fr *frame) {
						y := fr.uint64(iy)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uintptr:
					return func(fr *frame) {
						y := fr.uintptr(iy)
						fr.setReg(ir, x<<y)
					}
				}
			} else if ky == kindConst {
				switch ykind {
				case reflect.Int:
					y := basic.Int(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(int16)
						fr.setReg(ir, x<<y)
					}
				case reflect.Int8:
					y := basic.Int8(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(int16)
						fr.setReg(ir, x<<y)
					}
				case reflect.Int16:
					y := basic.Int16(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(int16)
						fr.setReg(ir, x<<y)
					}
				case reflect.Int32:
					y := basic.Int32(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(int16)
						fr.setReg(ir, x<<y)
					}
				case reflect.Int64:
					y := basic.Int64(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(int16)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint:
					y := basic.Uint(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(int16)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint8:
					y := basic.Uint8(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(int16)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint16:
					y := basic.Uint16(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(int16)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint32:
					y := basic.Uint32(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(int16)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint64:
					y := basic.Uint64(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(int16)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uintptr:
					y := basic.Uintptr(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(int16)
						fr.setReg(ir, x<<y)
					}
				}
			} else {
				switch ykind {
				case reflect.Int:
					return func(fr *frame) {
						x := fr.int16(ix)
						y := fr.int(iy)
						fr.setReg(ir, x<<y)
					}
				case reflect.Int8:
					return func(fr *frame) {
						x := fr.int16(ix)
						y := fr.int8(iy)
						fr.setReg(ir, x<<y)
					}
				case reflect.Int16:
					return func(fr *frame) {
						x := fr.int16(ix)
						y := fr.int16(iy)
						fr.setReg(ir, x<<y)
					}
				case reflect.Int32:
					return func(fr *frame) {
						x := fr.int16(ix)
						y := fr.int32(iy)
						fr.setReg(ir, x<<y)
					}
				case reflect.Int64:
					return func(fr *frame) {
						x := fr.int16(ix)
						y := fr.int64(iy)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint:
					return func(fr *frame) {
						x := fr.int16(ix)
						y := fr.uint(iy)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint8:
					return func(fr *frame) {
						x := fr.int16(ix)
						y := fr.uint8(iy)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint16:
					return func(fr *frame) {
						x := fr.int16(ix)
						y := fr.uint16(iy)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint32:
					return func(fr *frame) {
						x := fr.int16(ix)
						y := fr.uint32(iy)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint64:
					return func(fr *frame) {
						x := fr.int16(ix)
						y := fr.uint64(iy)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uintptr:
					return func(fr *frame) {
						x := fr.int16(ix)
						y := fr.uintptr(iy)
						fr.setReg(ir, x<<y)
					}
				}
			}
		case reflect.Int32:
			if kx == kindConst {
				x := vx.(int32)
				switch ykind {
				case reflect.Int:
					return func(fr *frame) {
						y := fr.int(iy)
						fr.setReg(ir, x<<y)
					}
				case reflect.Int8:
					return func(fr *frame) {
						y := fr.int8(iy)
						fr.setReg(ir, x<<y)
					}
				case reflect.Int16:
					return func(fr *frame) {
						y := fr.int16(iy)
						fr.setReg(ir, x<<y)
					}
				case reflect.Int32:
					return func(fr *frame) {
						y := fr.int32(iy)
						fr.setReg(ir, x<<y)
					}
				case reflect.Int64:
					return func(fr *frame) {
						y := fr.int64(iy)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint:
					return func(fr *frame) {
						y := fr.uint(iy)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint8:
					return func(fr *frame) {
						y := fr.uint8(iy)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint16:
					return func(fr *frame) {
						y := fr.uint16(iy)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint32:
					return func(fr *frame) {
						y := fr.uint32(iy)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint64:
					return func(fr *frame) {
						y := fr.uint64(iy)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uintptr:
					return func(fr *frame) {
						y := fr.uintptr(iy)
						fr.setReg(ir, x<<y)
					}
				}
			} else if ky == kindConst {
				switch ykind {
				case reflect.Int:
					y := basic.Int(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(int32)
						fr.setReg(ir, x<<y)
					}
				case reflect.Int8:
					y := basic.Int8(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(int32)
						fr.setReg(ir, x<<y)
					}
				case reflect.Int16:
					y := basic.Int16(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(int32)
						fr.setReg(ir, x<<y)
					}
				case reflect.Int32:
					y := basic.Int32(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(int32)
						fr.setReg(ir, x<<y)
					}
				case reflect.Int64:
					y := basic.Int64(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(int32)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint:
					y := basic.Uint(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(int32)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint8:
					y := basic.Uint8(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(int32)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint16:
					y := basic.Uint16(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(int32)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint32:
					y := basic.Uint32(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(int32)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint64:
					y := basic.Uint64(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(int32)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uintptr:
					y := basic.Uintptr(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(int32)
						fr.setReg(ir, x<<y)
					}
				}
			} else {
				switch ykind {
				case reflect.Int:
					return func(fr *frame) {
						x := fr.int32(ix)
						y := fr.int(iy)
						fr.setReg(ir, x<<y)
					}
				case reflect.Int8:
					return func(fr *frame) {
						x := fr.int32(ix)
						y := fr.int8(iy)
						fr.setReg(ir, x<<y)
					}
				case reflect.Int16:
					return func(fr *frame) {
						x := fr.int32(ix)
						y := fr.int16(iy)
						fr.setReg(ir, x<<y)
					}
				case reflect.Int32:
					return func(fr *frame) {
						x := fr.int32(ix)
						y := fr.int32(iy)
						fr.setReg(ir, x<<y)
					}
				case reflect.Int64:
					return func(fr *frame) {
						x := fr.int32(ix)
						y := fr.int64(iy)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint:
					return func(fr *frame) {
						x := fr.int32(ix)
						y := fr.uint(iy)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint8:
					return func(fr *frame) {
						x := fr.int32(ix)
						y := fr.uint8(iy)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint16:
					return func(fr *frame) {
						x := fr.int32(ix)
						y := fr.uint16(iy)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint32:
					return func(fr *frame) {
						x := fr.int32(ix)
						y := fr.uint32(iy)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint64:
					return func(fr *frame) {
						x := fr.int32(ix)
						y := fr.uint64(iy)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uintptr:
					return func(fr *frame) {
						x := fr.int32(ix)
						y := fr.uintptr(iy)
						fr.setReg(ir, x<<y)
					}
				}
			}
		case reflect.Int64:
			if kx == kindConst {
				x := vx.(int64)
				switch ykind {
				case reflect.Int:
					return func(fr *frame) {
						y := fr.int(iy)
						fr.setReg(ir, x<<y)
					}
				case reflect.Int8:
					return func(fr *frame) {
						y := fr.int8(iy)
						fr.setReg(ir, x<<y)
					}
				case reflect.Int16:
					return func(fr *frame) {
						y := fr.int16(iy)
						fr.setReg(ir, x<<y)
					}
				case reflect.Int32:
					return func(fr *frame) {
						y := fr.int32(iy)
						fr.setReg(ir, x<<y)
					}
				case reflect.Int64:
					return func(fr *frame) {
						y := fr.int64(iy)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint:
					return func(fr *frame) {
						y := fr.uint(iy)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint8:
					return func(fr *frame) {
						y := fr.uint8(iy)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint16:
					return func(fr *frame) {
						y := fr.uint16(iy)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint32:
					return func(fr *frame) {
						y := fr.uint32(iy)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint64:
					return func(fr *frame) {
						y := fr.uint64(iy)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uintptr:
					return func(fr *frame) {
						y := fr.uintptr(iy)
						fr.setReg(ir, x<<y)
					}
				}
			} else if ky == kindConst {
				switch ykind {
				case reflect.Int:
					y := basic.Int(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(int64)
						fr.setReg(ir, x<<y)
					}
				case reflect.Int8:
					y := basic.Int8(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(int64)
						fr.setReg(ir, x<<y)
					}
				case reflect.Int16:
					y := basic.Int16(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(int64)
						fr.setReg(ir, x<<y)
					}
				case reflect.Int32:
					y := basic.Int32(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(int64)
						fr.setReg(ir, x<<y)
					}
				case reflect.Int64:
					y := basic.Int64(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(int64)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint:
					y := basic.Uint(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(int64)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint8:
					y := basic.Uint8(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(int64)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint16:
					y := basic.Uint16(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(int64)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint32:
					y := basic.Uint32(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(int64)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint64:
					y := basic.Uint64(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(int64)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uintptr:
					y := basic.Uintptr(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(int64)
						fr.setReg(ir, x<<y)
					}
				}
			} else {
				switch ykind {
				case reflect.Int:
					return func(fr *frame) {
						x := fr.int64(ix)
						y := fr.int(iy)
						fr.setReg(ir, x<<y)
					}
				case reflect.Int8:
					return func(fr *frame) {
						x := fr.int64(ix)
						y := fr.int8(iy)
						fr.setReg(ir, x<<y)
					}
				case reflect.Int16:
					return func(fr *frame) {
						x := fr.int64(ix)
						y := fr.int16(iy)
						fr.setReg(ir, x<<y)
					}
				case reflect.Int32:
					return func(fr *frame) {
						x := fr.int64(ix)
						y := fr.int32(iy)
						fr.setReg(ir, x<<y)
					}
				case reflect.Int64:
					return func(fr *frame) {
						x := fr.int64(ix)
						y := fr.int64(iy)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint:
					return func(fr *frame) {
						x := fr.int64(ix)
						y := fr.uint(iy)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint8:
					return func(fr *frame) {
						x := fr.int64(ix)
						y := fr.uint8(iy)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint16:
					return func(fr *frame) {
						x := fr.int64(ix)
						y := fr.uint16(iy)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint32:
					return func(fr *frame) {
						x := fr.int64(ix)
						y := fr.uint32(iy)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint64:
					return func(fr *frame) {
						x := fr.int64(ix)
						y := fr.uint64(iy)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uintptr:
					return func(fr *frame) {
						x := fr.int64(ix)
						y := fr.uintptr(iy)
						fr.setReg(ir, x<<y)
					}
				}
			}
		case reflect.Uint:
			if kx == kindConst {
				x := vx.(uint)
				switch ykind {
				case reflect.Int:
					return func(fr *frame) {
						y := fr.int(iy)
						fr.setReg(ir, x<<y)
					}
				case reflect.Int8:
					return func(fr *frame) {
						y := fr.int8(iy)
						fr.setReg(ir, x<<y)
					}
				case reflect.Int16:
					return func(fr *frame) {
						y := fr.int16(iy)
						fr.setReg(ir, x<<y)
					}
				case reflect.Int32:
					return func(fr *frame) {
						y := fr.int32(iy)
						fr.setReg(ir, x<<y)
					}
				case reflect.Int64:
					return func(fr *frame) {
						y := fr.int64(iy)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint:
					return func(fr *frame) {
						y := fr.uint(iy)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint8:
					return func(fr *frame) {
						y := fr.uint8(iy)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint16:
					return func(fr *frame) {
						y := fr.uint16(iy)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint32:
					return func(fr *frame) {
						y := fr.uint32(iy)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint64:
					return func(fr *frame) {
						y := fr.uint64(iy)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uintptr:
					return func(fr *frame) {
						y := fr.uintptr(iy)
						fr.setReg(ir, x<<y)
					}
				}
			} else if ky == kindConst {
				switch ykind {
				case reflect.Int:
					y := basic.Int(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(uint)
						fr.setReg(ir, x<<y)
					}
				case reflect.Int8:
					y := basic.Int8(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(uint)
						fr.setReg(ir, x<<y)
					}
				case reflect.Int16:
					y := basic.Int16(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(uint)
						fr.setReg(ir, x<<y)
					}
				case reflect.Int32:
					y := basic.Int32(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(uint)
						fr.setReg(ir, x<<y)
					}
				case reflect.Int64:
					y := basic.Int64(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(uint)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint:
					y := basic.Uint(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(uint)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint8:
					y := basic.Uint8(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(uint)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint16:
					y := basic.Uint16(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(uint)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint32:
					y := basic.Uint32(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(uint)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint64:
					y := basic.Uint64(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(uint)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uintptr:
					y := basic.Uintptr(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(uint)
						fr.setReg(ir, x<<y)
					}
				}
			} else {
				switch ykind {
				case reflect.Int:
					return func(fr *frame) {
						x := fr.uint(ix)
						y := fr.int(iy)
						fr.setReg(ir, x<<y)
					}
				case reflect.Int8:
					return func(fr *frame) {
						x := fr.uint(ix)
						y := fr.int8(iy)
						fr.setReg(ir, x<<y)
					}
				case reflect.Int16:
					return func(fr *frame) {
						x := fr.uint(ix)
						y := fr.int16(iy)
						fr.setReg(ir, x<<y)
					}
				case reflect.Int32:
					return func(fr *frame) {
						x := fr.uint(ix)
						y := fr.int32(iy)
						fr.setReg(ir, x<<y)
					}
				case reflect.Int64:
					return func(fr *frame) {
						x := fr.uint(ix)
						y := fr.int64(iy)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint:
					return func(fr *frame) {
						x := fr.uint(ix)
						y := fr.uint(iy)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint8:
					return func(fr *frame) {
						x := fr.uint(ix)
						y := fr.uint8(iy)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint16:
					return func(fr *frame) {
						x := fr.uint(ix)
						y := fr.uint16(iy)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint32:
					return func(fr *frame) {
						x := fr.uint(ix)
						y := fr.uint32(iy)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint64:
					return func(fr *frame) {
						x := fr.uint(ix)
						y := fr.uint64(iy)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uintptr:
					return func(fr *frame) {
						x := fr.uint(ix)
						y := fr.uintptr(iy)
						fr.setReg(ir, x<<y)
					}
				}
			}
		case reflect.Uint8:
			if kx == kindConst {
				x := vx.(uint8)
				switch ykind {
				case reflect.Int:
					return func(fr *frame) {
						y := fr.int(iy)
						fr.setReg(ir, x<<y)
					}
				case reflect.Int8:
					return func(fr *frame) {
						y := fr.int8(iy)
						fr.setReg(ir, x<<y)
					}
				case reflect.Int16:
					return func(fr *frame) {
						y := fr.int16(iy)
						fr.setReg(ir, x<<y)
					}
				case reflect.Int32:
					return func(fr *frame) {
						y := fr.int32(iy)
						fr.setReg(ir, x<<y)
					}
				case reflect.Int64:
					return func(fr *frame) {
						y := fr.int64(iy)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint:
					return func(fr *frame) {
						y := fr.uint(iy)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint8:
					return func(fr *frame) {
						y := fr.uint8(iy)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint16:
					return func(fr *frame) {
						y := fr.uint16(iy)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint32:
					return func(fr *frame) {
						y := fr.uint32(iy)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint64:
					return func(fr *frame) {
						y := fr.uint64(iy)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uintptr:
					return func(fr *frame) {
						y := fr.uintptr(iy)
						fr.setReg(ir, x<<y)
					}
				}
			} else if ky == kindConst {
				switch ykind {
				case reflect.Int:
					y := basic.Int(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(uint8)
						fr.setReg(ir, x<<y)
					}
				case reflect.Int8:
					y := basic.Int8(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(uint8)
						fr.setReg(ir, x<<y)
					}
				case reflect.Int16:
					y := basic.Int16(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(uint8)
						fr.setReg(ir, x<<y)
					}
				case reflect.Int32:
					y := basic.Int32(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(uint8)
						fr.setReg(ir, x<<y)
					}
				case reflect.Int64:
					y := basic.Int64(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(uint8)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint:
					y := basic.Uint(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(uint8)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint8:
					y := basic.Uint8(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(uint8)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint16:
					y := basic.Uint16(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(uint8)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint32:
					y := basic.Uint32(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(uint8)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint64:
					y := basic.Uint64(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(uint8)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uintptr:
					y := basic.Uintptr(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(uint8)
						fr.setReg(ir, x<<y)
					}
				}
			} else {
				switch ykind {
				case reflect.Int:
					return func(fr *frame) {
						x := fr.uint8(ix)
						y := fr.int(iy)
						fr.setReg(ir, x<<y)
					}
				case reflect.Int8:
					return func(fr *frame) {
						x := fr.uint8(ix)
						y := fr.int8(iy)
						fr.setReg(ir, x<<y)
					}
				case reflect.Int16:
					return func(fr *frame) {
						x := fr.uint8(ix)
						y := fr.int16(iy)
						fr.setReg(ir, x<<y)
					}
				case reflect.Int32:
					return func(fr *frame) {
						x := fr.uint8(ix)
						y := fr.int32(iy)
						fr.setReg(ir, x<<y)
					}
				case reflect.Int64:
					return func(fr *frame) {
						x := fr.uint8(ix)
						y := fr.int64(iy)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint:
					return func(fr *frame) {
						x := fr.uint8(ix)
						y := fr.uint(iy)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint8:
					return func(fr *frame) {
						x := fr.uint8(ix)
						y := fr.uint8(iy)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint16:
					return func(fr *frame) {
						x := fr.uint8(ix)
						y := fr.uint16(iy)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint32:
					return func(fr *frame) {
						x := fr.uint8(ix)
						y := fr.uint32(iy)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint64:
					return func(fr *frame) {
						x := fr.uint8(ix)
						y := fr.uint64(iy)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uintptr:
					return func(fr *frame) {
						x := fr.uint8(ix)
						y := fr.uintptr(iy)
						fr.setReg(ir, x<<y)
					}
				}
			}
		case reflect.Uint16:
			if kx == kindConst {
				x := vx.(uint16)
				switch ykind {
				case reflect.Int:
					return func(fr *frame) {
						y := fr.int(iy)
						fr.setReg(ir, x<<y)
					}
				case reflect.Int8:
					return func(fr *frame) {
						y := fr.int8(iy)
						fr.setReg(ir, x<<y)
					}
				case reflect.Int16:
					return func(fr *frame) {
						y := fr.int16(iy)
						fr.setReg(ir, x<<y)
					}
				case reflect.Int32:
					return func(fr *frame) {
						y := fr.int32(iy)
						fr.setReg(ir, x<<y)
					}
				case reflect.Int64:
					return func(fr *frame) {
						y := fr.int64(iy)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint:
					return func(fr *frame) {
						y := fr.uint(iy)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint8:
					return func(fr *frame) {
						y := fr.uint8(iy)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint16:
					return func(fr *frame) {
						y := fr.uint16(iy)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint32:
					return func(fr *frame) {
						y := fr.uint32(iy)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint64:
					return func(fr *frame) {
						y := fr.uint64(iy)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uintptr:
					return func(fr *frame) {
						y := fr.uintptr(iy)
						fr.setReg(ir, x<<y)
					}
				}
			} else if ky == kindConst {
				switch ykind {
				case reflect.Int:
					y := basic.Int(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(uint16)
						fr.setReg(ir, x<<y)
					}
				case reflect.Int8:
					y := basic.Int8(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(uint16)
						fr.setReg(ir, x<<y)
					}
				case reflect.Int16:
					y := basic.Int16(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(uint16)
						fr.setReg(ir, x<<y)
					}
				case reflect.Int32:
					y := basic.Int32(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(uint16)
						fr.setReg(ir, x<<y)
					}
				case reflect.Int64:
					y := basic.Int64(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(uint16)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint:
					y := basic.Uint(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(uint16)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint8:
					y := basic.Uint8(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(uint16)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint16:
					y := basic.Uint16(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(uint16)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint32:
					y := basic.Uint32(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(uint16)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint64:
					y := basic.Uint64(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(uint16)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uintptr:
					y := basic.Uintptr(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(uint16)
						fr.setReg(ir, x<<y)
					}
				}
			} else {
				switch ykind {
				case reflect.Int:
					return func(fr *frame) {
						x := fr.uint16(ix)
						y := fr.int(iy)
						fr.setReg(ir, x<<y)
					}
				case reflect.Int8:
					return func(fr *frame) {
						x := fr.uint16(ix)
						y := fr.int8(iy)
						fr.setReg(ir, x<<y)
					}
				case reflect.Int16:
					return func(fr *frame) {
						x := fr.uint16(ix)
						y := fr.int16(iy)
						fr.setReg(ir, x<<y)
					}
				case reflect.Int32:
					return func(fr *frame) {
						x := fr.uint16(ix)
						y := fr.int32(iy)
						fr.setReg(ir, x<<y)
					}
				case reflect.Int64:
					return func(fr *frame) {
						x := fr.uint16(ix)
						y := fr.int64(iy)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint:
					return func(fr *frame) {
						x := fr.uint16(ix)
						y := fr.uint(iy)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint8:
					return func(fr *frame) {
						x := fr.uint16(ix)
						y := fr.uint8(iy)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint16:
					return func(fr *frame) {
						x := fr.uint16(ix)
						y := fr.uint16(iy)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint32:
					return func(fr *frame) {
						x := fr.uint16(ix)
						y := fr.uint32(iy)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint64:
					return func(fr *frame) {
						x := fr.uint16(ix)
						y := fr.uint64(iy)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uintptr:
					return func(fr *frame) {
						x := fr.uint16(ix)
						y := fr.uintptr(iy)
						fr.setReg(ir, x<<y)
					}
				}
			}
		case reflect.Uint32:
			if kx == kindConst {
				x := vx.(uint32)
				switch ykind {
				case reflect.Int:
					return func(fr *frame) {
						y := fr.int(iy)
						fr.setReg(ir, x<<y)
					}
				case reflect.Int8:
					return func(fr *frame) {
						y := fr.int8(iy)
						fr.setReg(ir, x<<y)
					}
				case reflect.Int16:
					return func(fr *frame) {
						y := fr.int16(iy)
						fr.setReg(ir, x<<y)
					}
				case reflect.Int32:
					return func(fr *frame) {
						y := fr.int32(iy)
						fr.setReg(ir, x<<y)
					}
				case reflect.Int64:
					return func(fr *frame) {
						y := fr.int64(iy)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint:
					return func(fr *frame) {
						y := fr.uint(iy)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint8:
					return func(fr *frame) {
						y := fr.uint8(iy)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint16:
					return func(fr *frame) {
						y := fr.uint16(iy)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint32:
					return func(fr *frame) {
						y := fr.uint32(iy)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint64:
					return func(fr *frame) {
						y := fr.uint64(iy)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uintptr:
					return func(fr *frame) {
						y := fr.uintptr(iy)
						fr.setReg(ir, x<<y)
					}
				}
			} else if ky == kindConst {
				switch ykind {
				case reflect.Int:
					y := basic.Int(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(uint32)
						fr.setReg(ir, x<<y)
					}
				case reflect.Int8:
					y := basic.Int8(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(uint32)
						fr.setReg(ir, x<<y)
					}
				case reflect.Int16:
					y := basic.Int16(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(uint32)
						fr.setReg(ir, x<<y)
					}
				case reflect.Int32:
					y := basic.Int32(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(uint32)
						fr.setReg(ir, x<<y)
					}
				case reflect.Int64:
					y := basic.Int64(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(uint32)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint:
					y := basic.Uint(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(uint32)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint8:
					y := basic.Uint8(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(uint32)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint16:
					y := basic.Uint16(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(uint32)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint32:
					y := basic.Uint32(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(uint32)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint64:
					y := basic.Uint64(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(uint32)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uintptr:
					y := basic.Uintptr(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(uint32)
						fr.setReg(ir, x<<y)
					}
				}
			} else {
				switch ykind {
				case reflect.Int:
					return func(fr *frame) {
						x := fr.uint32(ix)
						y := fr.int(iy)
						fr.setReg(ir, x<<y)
					}
				case reflect.Int8:
					return func(fr *frame) {
						x := fr.uint32(ix)
						y := fr.int8(iy)
						fr.setReg(ir, x<<y)
					}
				case reflect.Int16:
					return func(fr *frame) {
						x := fr.uint32(ix)
						y := fr.int16(iy)
						fr.setReg(ir, x<<y)
					}
				case reflect.Int32:
					return func(fr *frame) {
						x := fr.uint32(ix)
						y := fr.int32(iy)
						fr.setReg(ir, x<<y)
					}
				case reflect.Int64:
					return func(fr *frame) {
						x := fr.uint32(ix)
						y := fr.int64(iy)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint:
					return func(fr *frame) {
						x := fr.uint32(ix)
						y := fr.uint(iy)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint8:
					return func(fr *frame) {
						x := fr.uint32(ix)
						y := fr.uint8(iy)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint16:
					return func(fr *frame) {
						x := fr.uint32(ix)
						y := fr.uint16(iy)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint32:
					return func(fr *frame) {
						x := fr.uint32(ix)
						y := fr.uint32(iy)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint64:
					return func(fr *frame) {
						x := fr.uint32(ix)
						y := fr.uint64(iy)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uintptr:
					return func(fr *frame) {
						x := fr.uint32(ix)
						y := fr.uintptr(iy)
						fr.setReg(ir, x<<y)
					}
				}
			}
		case reflect.Uint64:
			if kx == kindConst {
				x := vx.(uint64)
				switch ykind {
				case reflect.Int:
					return func(fr *frame) {
						y := fr.int(iy)
						fr.setReg(ir, x<<y)
					}
				case reflect.Int8:
					return func(fr *frame) {
						y := fr.int8(iy)
						fr.setReg(ir, x<<y)
					}
				case reflect.Int16:
					return func(fr *frame) {
						y := fr.int16(iy)
						fr.setReg(ir, x<<y)
					}
				case reflect.Int32:
					return func(fr *frame) {
						y := fr.int32(iy)
						fr.setReg(ir, x<<y)
					}
				case reflect.Int64:
					return func(fr *frame) {
						y := fr.int64(iy)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint:
					return func(fr *frame) {
						y := fr.uint(iy)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint8:
					return func(fr *frame) {
						y := fr.uint8(iy)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint16:
					return func(fr *frame) {
						y := fr.uint16(iy)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint32:
					return func(fr *frame) {
						y := fr.uint32(iy)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint64:
					return func(fr *frame) {
						y := fr.uint64(iy)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uintptr:
					return func(fr *frame) {
						y := fr.uintptr(iy)
						fr.setReg(ir, x<<y)
					}
				}
			} else if ky == kindConst {
				switch ykind {
				case reflect.Int:
					y := basic.Int(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(uint64)
						fr.setReg(ir, x<<y)
					}
				case reflect.Int8:
					y := basic.Int8(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(uint64)
						fr.setReg(ir, x<<y)
					}
				case reflect.Int16:
					y := basic.Int16(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(uint64)
						fr.setReg(ir, x<<y)
					}
				case reflect.Int32:
					y := basic.Int32(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(uint64)
						fr.setReg(ir, x<<y)
					}
				case reflect.Int64:
					y := basic.Int64(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(uint64)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint:
					y := basic.Uint(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(uint64)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint8:
					y := basic.Uint8(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(uint64)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint16:
					y := basic.Uint16(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(uint64)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint32:
					y := basic.Uint32(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(uint64)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint64:
					y := basic.Uint64(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(uint64)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uintptr:
					y := basic.Uintptr(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(uint64)
						fr.setReg(ir, x<<y)
					}
				}
			} else {
				switch ykind {
				case reflect.Int:
					return func(fr *frame) {
						x := fr.uint64(ix)
						y := fr.int(iy)
						fr.setReg(ir, x<<y)
					}
				case reflect.Int8:
					return func(fr *frame) {
						x := fr.uint64(ix)
						y := fr.int8(iy)
						fr.setReg(ir, x<<y)
					}
				case reflect.Int16:
					return func(fr *frame) {
						x := fr.uint64(ix)
						y := fr.int16(iy)
						fr.setReg(ir, x<<y)
					}
				case reflect.Int32:
					return func(fr *frame) {
						x := fr.uint64(ix)
						y := fr.int32(iy)
						fr.setReg(ir, x<<y)
					}
				case reflect.Int64:
					return func(fr *frame) {
						x := fr.uint64(ix)
						y := fr.int64(iy)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint:
					return func(fr *frame) {
						x := fr.uint64(ix)
						y := fr.uint(iy)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint8:
					return func(fr *frame) {
						x := fr.uint64(ix)
						y := fr.uint8(iy)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint16:
					return func(fr *frame) {
						x := fr.uint64(ix)
						y := fr.uint16(iy)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint32:
					return func(fr *frame) {
						x := fr.uint64(ix)
						y := fr.uint32(iy)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint64:
					return func(fr *frame) {
						x := fr.uint64(ix)
						y := fr.uint64(iy)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uintptr:
					return func(fr *frame) {
						x := fr.uint64(ix)
						y := fr.uintptr(iy)
						fr.setReg(ir, x<<y)
					}
				}
			}
		case reflect.Uintptr:
			if kx == kindConst {
				x := vx.(uintptr)
				switch ykind {
				case reflect.Int:
					return func(fr *frame) {
						y := fr.int(iy)
						fr.setReg(ir, x<<y)
					}
				case reflect.Int8:
					return func(fr *frame) {
						y := fr.int8(iy)
						fr.setReg(ir, x<<y)
					}
				case reflect.Int16:
					return func(fr *frame) {
						y := fr.int16(iy)
						fr.setReg(ir, x<<y)
					}
				case reflect.Int32:
					return func(fr *frame) {
						y := fr.int32(iy)
						fr.setReg(ir, x<<y)
					}
				case reflect.Int64:
					return func(fr *frame) {
						y := fr.int64(iy)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint:
					return func(fr *frame) {
						y := fr.uint(iy)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint8:
					return func(fr *frame) {
						y := fr.uint8(iy)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint16:
					return func(fr *frame) {
						y := fr.uint16(iy)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint32:
					return func(fr *frame) {
						y := fr.uint32(iy)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint64:
					return func(fr *frame) {
						y := fr.uint64(iy)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uintptr:
					return func(fr *frame) {
						y := fr.uintptr(iy)
						fr.setReg(ir, x<<y)
					}
				}
			} else if ky == kindConst {
				switch ykind {
				case reflect.Int:
					y := basic.Int(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(uintptr)
						fr.setReg(ir, x<<y)
					}
				case reflect.Int8:
					y := basic.Int8(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(uintptr)
						fr.setReg(ir, x<<y)
					}
				case reflect.Int16:
					y := basic.Int16(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(uintptr)
						fr.setReg(ir, x<<y)
					}
				case reflect.Int32:
					y := basic.Int32(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(uintptr)
						fr.setReg(ir, x<<y)
					}
				case reflect.Int64:
					y := basic.Int64(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(uintptr)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint:
					y := basic.Uint(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(uintptr)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint8:
					y := basic.Uint8(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(uintptr)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint16:
					y := basic.Uint16(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(uintptr)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint32:
					y := basic.Uint32(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(uintptr)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint64:
					y := basic.Uint64(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(uintptr)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uintptr:
					y := basic.Uintptr(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(uintptr)
						fr.setReg(ir, x<<y)
					}
				}
			} else {
				switch ykind {
				case reflect.Int:
					return func(fr *frame) {
						x := fr.uintptr(ix)
						y := fr.int(iy)
						fr.setReg(ir, x<<y)
					}
				case reflect.Int8:
					return func(fr *frame) {
						x := fr.uintptr(ix)
						y := fr.int8(iy)
						fr.setReg(ir, x<<y)
					}
				case reflect.Int16:
					return func(fr *frame) {
						x := fr.uintptr(ix)
						y := fr.int16(iy)
						fr.setReg(ir, x<<y)
					}
				case reflect.Int32:
					return func(fr *frame) {
						x := fr.uintptr(ix)
						y := fr.int32(iy)
						fr.setReg(ir, x<<y)
					}
				case reflect.Int64:
					return func(fr *frame) {
						x := fr.uintptr(ix)
						y := fr.int64(iy)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint:
					return func(fr *frame) {
						x := fr.uintptr(ix)
						y := fr.uint(iy)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint8:
					return func(fr *frame) {
						x := fr.uintptr(ix)
						y := fr.uint8(iy)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint16:
					return func(fr *frame) {
						x := fr.uintptr(ix)
						y := fr.uint16(iy)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint32:
					return func(fr *frame) {
						x := fr.uintptr(ix)
						y := fr.uint32(iy)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint64:
					return func(fr *frame) {
						x := fr.uintptr(ix)
						y := fr.uint64(iy)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uintptr:
					return func(fr *frame) {
						x := fr.uintptr(ix)
						y := fr.uintptr(iy)
						fr.setReg(ir, x<<y)
					}
				}
			}
		}
	} else {
		t := basic.TypeOfType(xtyp)
		switch xkind {
		case reflect.Int:
			if kx == kindConst {
				x := basic.Int(vx)
				switch ykind {
				case reflect.Int:
					return func(fr *frame) {
						y := fr.int(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int8:
					return func(fr *frame) {
						y := fr.int8(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int16:
					return func(fr *frame) {
						y := fr.int16(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int32:
					return func(fr *frame) {
						y := fr.int32(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int64:
					return func(fr *frame) {
						y := fr.int64(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint:
					return func(fr *frame) {
						y := fr.uint(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint8:
					return func(fr *frame) {
						y := fr.uint8(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint16:
					return func(fr *frame) {
						y := fr.uint16(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint32:
					return func(fr *frame) {
						y := fr.uint32(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint64:
					return func(fr *frame) {
						y := fr.uint64(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uintptr:
					return func(fr *frame) {
						y := fr.uintptr(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				}
			} else if ky == kindConst {
				switch ykind {
				case reflect.Int:
					y := basic.Int(vy)
					return func(fr *frame) {
						x := fr.int(ix)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int8:
					y := basic.Int8(vy)
					return func(fr *frame) {
						x := fr.int(ix)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int16:
					y := basic.Int16(vy)
					return func(fr *frame) {
						x := fr.int(ix)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int32:
					y := basic.Int32(vy)
					return func(fr *frame) {
						x := fr.int(ix)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int64:
					y := basic.Int64(vy)
					return func(fr *frame) {
						x := fr.int(ix)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint:
					y := basic.Uint(vy)
					return func(fr *frame) {
						x := fr.int(ix)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint8:
					y := basic.Uint8(vy)
					return func(fr *frame) {
						x := fr.int(ix)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint16:
					y := basic.Uint16(vy)
					return func(fr *frame) {
						x := fr.int(ix)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint32:
					y := basic.Uint32(vy)
					return func(fr *frame) {
						x := fr.int(ix)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint64:
					y := basic.Uint64(vy)
					return func(fr *frame) {
						x := fr.int(ix)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uintptr:
					y := basic.Uintptr(vy)
					return func(fr *frame) {
						x := fr.int(ix)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				}
			} else {
				switch ykind {
				case reflect.Int:
					return func(fr *frame) {
						x := fr.int(ix)
						y := fr.int(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int8:
					return func(fr *frame) {
						x := fr.int(ix)
						y := fr.int8(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int16:
					return func(fr *frame) {
						x := fr.int(ix)
						y := fr.int16(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int32:
					return func(fr *frame) {
						x := fr.int(ix)
						y := fr.int32(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int64:
					return func(fr *frame) {
						x := fr.int(ix)
						y := fr.int64(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint:
					return func(fr *frame) {
						x := fr.int(ix)
						y := fr.uint(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint8:
					return func(fr *frame) {
						x := fr.int(ix)
						y := fr.uint8(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint16:
					return func(fr *frame) {
						x := fr.int(ix)
						y := fr.uint16(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint32:
					return func(fr *frame) {
						x := fr.int(ix)
						y := fr.uint32(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint64:
					return func(fr *frame) {
						x := fr.int(ix)
						y := fr.uint64(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uintptr:
					return func(fr *frame) {
						x := fr.int(ix)
						y := fr.uintptr(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				}
			}
		case reflect.Int8:
			if kx == kindConst {
				x := basic.Int8(vx)
				switch ykind {
				case reflect.Int:
					return func(fr *frame) {
						y := fr.int(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int8:
					return func(fr *frame) {
						y := fr.int8(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int16:
					return func(fr *frame) {
						y := fr.int16(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int32:
					return func(fr *frame) {
						y := fr.int32(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int64:
					return func(fr *frame) {
						y := fr.int64(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint:
					return func(fr *frame) {
						y := fr.uint(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint8:
					return func(fr *frame) {
						y := fr.uint8(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint16:
					return func(fr *frame) {
						y := fr.uint16(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint32:
					return func(fr *frame) {
						y := fr.uint32(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint64:
					return func(fr *frame) {
						y := fr.uint64(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uintptr:
					return func(fr *frame) {
						y := fr.uintptr(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				}
			} else if ky == kindConst {
				switch ykind {
				case reflect.Int:
					y := basic.Int(vy)
					return func(fr *frame) {
						x := fr.int8(ix)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int8:
					y := basic.Int8(vy)
					return func(fr *frame) {
						x := fr.int8(ix)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int16:
					y := basic.Int16(vy)
					return func(fr *frame) {
						x := fr.int8(ix)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int32:
					y := basic.Int32(vy)
					return func(fr *frame) {
						x := fr.int8(ix)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int64:
					y := basic.Int64(vy)
					return func(fr *frame) {
						x := fr.int8(ix)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint:
					y := basic.Uint(vy)
					return func(fr *frame) {
						x := fr.int8(ix)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint8:
					y := basic.Uint8(vy)
					return func(fr *frame) {
						x := fr.int8(ix)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint16:
					y := basic.Uint16(vy)
					return func(fr *frame) {
						x := fr.int8(ix)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint32:
					y := basic.Uint32(vy)
					return func(fr *frame) {
						x := fr.int8(ix)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint64:
					y := basic.Uint64(vy)
					return func(fr *frame) {
						x := fr.int8(ix)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uintptr:
					y := basic.Uintptr(vy)
					return func(fr *frame) {
						x := fr.int8(ix)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				}
			} else {
				switch ykind {
				case reflect.Int:
					return func(fr *frame) {
						x := fr.int8(ix)
						y := fr.int(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int8:
					return func(fr *frame) {
						x := fr.int8(ix)
						y := fr.int8(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int16:
					return func(fr *frame) {
						x := fr.int8(ix)
						y := fr.int16(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int32:
					return func(fr *frame) {
						x := fr.int8(ix)
						y := fr.int32(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int64:
					return func(fr *frame) {
						x := fr.int8(ix)
						y := fr.int64(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint:
					return func(fr *frame) {
						x := fr.int8(ix)
						y := fr.uint(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint8:
					return func(fr *frame) {
						x := fr.int8(ix)
						y := fr.uint8(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint16:
					return func(fr *frame) {
						x := fr.int8(ix)
						y := fr.uint16(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint32:
					return func(fr *frame) {
						x := fr.int8(ix)
						y := fr.uint32(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint64:
					return func(fr *frame) {
						x := fr.int8(ix)
						y := fr.uint64(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uintptr:
					return func(fr *frame) {
						x := fr.int8(ix)
						y := fr.uintptr(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				}
			}
		case reflect.Int16:
			if kx == kindConst {
				x := basic.Int16(vx)
				switch ykind {
				case reflect.Int:
					return func(fr *frame) {
						y := fr.int(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int8:
					return func(fr *frame) {
						y := fr.int8(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int16:
					return func(fr *frame) {
						y := fr.int16(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int32:
					return func(fr *frame) {
						y := fr.int32(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int64:
					return func(fr *frame) {
						y := fr.int64(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint:
					return func(fr *frame) {
						y := fr.uint(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint8:
					return func(fr *frame) {
						y := fr.uint8(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint16:
					return func(fr *frame) {
						y := fr.uint16(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint32:
					return func(fr *frame) {
						y := fr.uint32(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint64:
					return func(fr *frame) {
						y := fr.uint64(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uintptr:
					return func(fr *frame) {
						y := fr.uintptr(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				}
			} else if ky == kindConst {
				switch ykind {
				case reflect.Int:
					y := basic.Int(vy)
					return func(fr *frame) {
						x := fr.int16(ix)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int8:
					y := basic.Int8(vy)
					return func(fr *frame) {
						x := fr.int16(ix)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int16:
					y := basic.Int16(vy)
					return func(fr *frame) {
						x := fr.int16(ix)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int32:
					y := basic.Int32(vy)
					return func(fr *frame) {
						x := fr.int16(ix)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int64:
					y := basic.Int64(vy)
					return func(fr *frame) {
						x := fr.int16(ix)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint:
					y := basic.Uint(vy)
					return func(fr *frame) {
						x := fr.int16(ix)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint8:
					y := basic.Uint8(vy)
					return func(fr *frame) {
						x := fr.int16(ix)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint16:
					y := basic.Uint16(vy)
					return func(fr *frame) {
						x := fr.int16(ix)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint32:
					y := basic.Uint32(vy)
					return func(fr *frame) {
						x := fr.int16(ix)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint64:
					y := basic.Uint64(vy)
					return func(fr *frame) {
						x := fr.int16(ix)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uintptr:
					y := basic.Uintptr(vy)
					return func(fr *frame) {
						x := fr.int16(ix)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				}
			} else {
				switch ykind {
				case reflect.Int:
					return func(fr *frame) {
						x := fr.int16(ix)
						y := fr.int(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int8:
					return func(fr *frame) {
						x := fr.int16(ix)
						y := fr.int8(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int16:
					return func(fr *frame) {
						x := fr.int16(ix)
						y := fr.int16(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int32:
					return func(fr *frame) {
						x := fr.int16(ix)
						y := fr.int32(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int64:
					return func(fr *frame) {
						x := fr.int16(ix)
						y := fr.int64(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint:
					return func(fr *frame) {
						x := fr.int16(ix)
						y := fr.uint(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint8:
					return func(fr *frame) {
						x := fr.int16(ix)
						y := fr.uint8(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint16:
					return func(fr *frame) {
						x := fr.int16(ix)
						y := fr.uint16(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint32:
					return func(fr *frame) {
						x := fr.int16(ix)
						y := fr.uint32(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint64:
					return func(fr *frame) {
						x := fr.int16(ix)
						y := fr.uint64(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uintptr:
					return func(fr *frame) {
						x := fr.int16(ix)
						y := fr.uintptr(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				}
			}
		case reflect.Int32:
			if kx == kindConst {
				x := basic.Int32(vx)
				switch ykind {
				case reflect.Int:
					return func(fr *frame) {
						y := fr.int(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int8:
					return func(fr *frame) {
						y := fr.int8(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int16:
					return func(fr *frame) {
						y := fr.int16(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int32:
					return func(fr *frame) {
						y := fr.int32(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int64:
					return func(fr *frame) {
						y := fr.int64(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint:
					return func(fr *frame) {
						y := fr.uint(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint8:
					return func(fr *frame) {
						y := fr.uint8(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint16:
					return func(fr *frame) {
						y := fr.uint16(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint32:
					return func(fr *frame) {
						y := fr.uint32(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint64:
					return func(fr *frame) {
						y := fr.uint64(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uintptr:
					return func(fr *frame) {
						y := fr.uintptr(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				}
			} else if ky == kindConst {
				switch ykind {
				case reflect.Int:
					y := basic.Int(vy)
					return func(fr *frame) {
						x := fr.int32(ix)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int8:
					y := basic.Int8(vy)
					return func(fr *frame) {
						x := fr.int32(ix)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int16:
					y := basic.Int16(vy)
					return func(fr *frame) {
						x := fr.int32(ix)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int32:
					y := basic.Int32(vy)
					return func(fr *frame) {
						x := fr.int32(ix)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int64:
					y := basic.Int64(vy)
					return func(fr *frame) {
						x := fr.int32(ix)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint:
					y := basic.Uint(vy)
					return func(fr *frame) {
						x := fr.int32(ix)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint8:
					y := basic.Uint8(vy)
					return func(fr *frame) {
						x := fr.int32(ix)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint16:
					y := basic.Uint16(vy)
					return func(fr *frame) {
						x := fr.int32(ix)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint32:
					y := basic.Uint32(vy)
					return func(fr *frame) {
						x := fr.int32(ix)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint64:
					y := basic.Uint64(vy)
					return func(fr *frame) {
						x := fr.int32(ix)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uintptr:
					y := basic.Uintptr(vy)
					return func(fr *frame) {
						x := fr.int32(ix)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				}
			} else {
				switch ykind {
				case reflect.Int:
					return func(fr *frame) {
						x := fr.int32(ix)
						y := fr.int(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int8:
					return func(fr *frame) {
						x := fr.int32(ix)
						y := fr.int8(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int16:
					return func(fr *frame) {
						x := fr.int32(ix)
						y := fr.int16(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int32:
					return func(fr *frame) {
						x := fr.int32(ix)
						y := fr.int32(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int64:
					return func(fr *frame) {
						x := fr.int32(ix)
						y := fr.int64(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint:
					return func(fr *frame) {
						x := fr.int32(ix)
						y := fr.uint(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint8:
					return func(fr *frame) {
						x := fr.int32(ix)
						y := fr.uint8(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint16:
					return func(fr *frame) {
						x := fr.int32(ix)
						y := fr.uint16(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint32:
					return func(fr *frame) {
						x := fr.int32(ix)
						y := fr.uint32(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint64:
					return func(fr *frame) {
						x := fr.int32(ix)
						y := fr.uint64(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uintptr:
					return func(fr *frame) {
						x := fr.int32(ix)
						y := fr.uintptr(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				}
			}
		case reflect.Int64:
			if kx == kindConst {
				x := basic.Int64(vx)
				switch ykind {
				case reflect.Int:
					return func(fr *frame) {
						y := fr.int(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int8:
					return func(fr *frame) {
						y := fr.int8(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int16:
					return func(fr *frame) {
						y := fr.int16(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int32:
					return func(fr *frame) {
						y := fr.int32(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int64:
					return func(fr *frame) {
						y := fr.int64(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint:
					return func(fr *frame) {
						y := fr.uint(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint8:
					return func(fr *frame) {
						y := fr.uint8(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint16:
					return func(fr *frame) {
						y := fr.uint16(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint32:
					return func(fr *frame) {
						y := fr.uint32(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint64:
					return func(fr *frame) {
						y := fr.uint64(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uintptr:
					return func(fr *frame) {
						y := fr.uintptr(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				}
			} else if ky == kindConst {
				switch ykind {
				case reflect.Int:
					y := basic.Int(vy)
					return func(fr *frame) {
						x := fr.int64(ix)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int8:
					y := basic.Int8(vy)
					return func(fr *frame) {
						x := fr.int64(ix)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int16:
					y := basic.Int16(vy)
					return func(fr *frame) {
						x := fr.int64(ix)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int32:
					y := basic.Int32(vy)
					return func(fr *frame) {
						x := fr.int64(ix)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int64:
					y := basic.Int64(vy)
					return func(fr *frame) {
						x := fr.int64(ix)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint:
					y := basic.Uint(vy)
					return func(fr *frame) {
						x := fr.int64(ix)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint8:
					y := basic.Uint8(vy)
					return func(fr *frame) {
						x := fr.int64(ix)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint16:
					y := basic.Uint16(vy)
					return func(fr *frame) {
						x := fr.int64(ix)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint32:
					y := basic.Uint32(vy)
					return func(fr *frame) {
						x := fr.int64(ix)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint64:
					y := basic.Uint64(vy)
					return func(fr *frame) {
						x := fr.int64(ix)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uintptr:
					y := basic.Uintptr(vy)
					return func(fr *frame) {
						x := fr.int64(ix)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				}
			} else {
				switch ykind {
				case reflect.Int:
					return func(fr *frame) {
						x := fr.int64(ix)
						y := fr.int(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int8:
					return func(fr *frame) {
						x := fr.int64(ix)
						y := fr.int8(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int16:
					return func(fr *frame) {
						x := fr.int64(ix)
						y := fr.int16(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int32:
					return func(fr *frame) {
						x := fr.int64(ix)
						y := fr.int32(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int64:
					return func(fr *frame) {
						x := fr.int64(ix)
						y := fr.int64(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint:
					return func(fr *frame) {
						x := fr.int64(ix)
						y := fr.uint(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint8:
					return func(fr *frame) {
						x := fr.int64(ix)
						y := fr.uint8(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint16:
					return func(fr *frame) {
						x := fr.int64(ix)
						y := fr.uint16(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint32:
					return func(fr *frame) {
						x := fr.int64(ix)
						y := fr.uint32(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint64:
					return func(fr *frame) {
						x := fr.int64(ix)
						y := fr.uint64(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uintptr:
					return func(fr *frame) {
						x := fr.int64(ix)
						y := fr.uintptr(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				}
			}
		case reflect.Uint:
			if kx == kindConst {
				x := basic.Uint(vx)
				switch ykind {
				case reflect.Int:
					return func(fr *frame) {
						y := fr.int(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int8:
					return func(fr *frame) {
						y := fr.int8(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int16:
					return func(fr *frame) {
						y := fr.int16(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int32:
					return func(fr *frame) {
						y := fr.int32(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int64:
					return func(fr *frame) {
						y := fr.int64(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint:
					return func(fr *frame) {
						y := fr.uint(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint8:
					return func(fr *frame) {
						y := fr.uint8(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint16:
					return func(fr *frame) {
						y := fr.uint16(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint32:
					return func(fr *frame) {
						y := fr.uint32(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint64:
					return func(fr *frame) {
						y := fr.uint64(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uintptr:
					return func(fr *frame) {
						y := fr.uintptr(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				}
			} else if ky == kindConst {
				switch ykind {
				case reflect.Int:
					y := basic.Int(vy)
					return func(fr *frame) {
						x := fr.uint(ix)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int8:
					y := basic.Int8(vy)
					return func(fr *frame) {
						x := fr.uint(ix)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int16:
					y := basic.Int16(vy)
					return func(fr *frame) {
						x := fr.uint(ix)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int32:
					y := basic.Int32(vy)
					return func(fr *frame) {
						x := fr.uint(ix)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int64:
					y := basic.Int64(vy)
					return func(fr *frame) {
						x := fr.uint(ix)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint:
					y := basic.Uint(vy)
					return func(fr *frame) {
						x := fr.uint(ix)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint8:
					y := basic.Uint8(vy)
					return func(fr *frame) {
						x := fr.uint(ix)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint16:
					y := basic.Uint16(vy)
					return func(fr *frame) {
						x := fr.uint(ix)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint32:
					y := basic.Uint32(vy)
					return func(fr *frame) {
						x := fr.uint(ix)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint64:
					y := basic.Uint64(vy)
					return func(fr *frame) {
						x := fr.uint(ix)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uintptr:
					y := basic.Uintptr(vy)
					return func(fr *frame) {
						x := fr.uint(ix)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				}
			} else {
				switch ykind {
				case reflect.Int:
					return func(fr *frame) {
						x := fr.uint(ix)
						y := fr.int(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int8:
					return func(fr *frame) {
						x := fr.uint(ix)
						y := fr.int8(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int16:
					return func(fr *frame) {
						x := fr.uint(ix)
						y := fr.int16(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int32:
					return func(fr *frame) {
						x := fr.uint(ix)
						y := fr.int32(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int64:
					return func(fr *frame) {
						x := fr.uint(ix)
						y := fr.int64(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint:
					return func(fr *frame) {
						x := fr.uint(ix)
						y := fr.uint(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint8:
					return func(fr *frame) {
						x := fr.uint(ix)
						y := fr.uint8(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint16:
					return func(fr *frame) {
						x := fr.uint(ix)
						y := fr.uint16(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint32:
					return func(fr *frame) {
						x := fr.uint(ix)
						y := fr.uint32(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint64:
					return func(fr *frame) {
						x := fr.uint(ix)
						y := fr.uint64(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uintptr:
					return func(fr *frame) {
						x := fr.uint(ix)
						y := fr.uintptr(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				}
			}
		case reflect.Uint8:
			if kx == kindConst {
				x := basic.Uint8(vx)
				switch ykind {
				case reflect.Int:
					return func(fr *frame) {
						y := fr.int(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int8:
					return func(fr *frame) {
						y := fr.int8(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int16:
					return func(fr *frame) {
						y := fr.int16(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int32:
					return func(fr *frame) {
						y := fr.int32(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int64:
					return func(fr *frame) {
						y := fr.int64(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint:
					return func(fr *frame) {
						y := fr.uint(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint8:
					return func(fr *frame) {
						y := fr.uint8(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint16:
					return func(fr *frame) {
						y := fr.uint16(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint32:
					return func(fr *frame) {
						y := fr.uint32(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint64:
					return func(fr *frame) {
						y := fr.uint64(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uintptr:
					return func(fr *frame) {
						y := fr.uintptr(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				}
			} else if ky == kindConst {
				switch ykind {
				case reflect.Int:
					y := basic.Int(vy)
					return func(fr *frame) {
						x := fr.uint8(ix)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int8:
					y := basic.Int8(vy)
					return func(fr *frame) {
						x := fr.uint8(ix)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int16:
					y := basic.Int16(vy)
					return func(fr *frame) {
						x := fr.uint8(ix)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int32:
					y := basic.Int32(vy)
					return func(fr *frame) {
						x := fr.uint8(ix)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int64:
					y := basic.Int64(vy)
					return func(fr *frame) {
						x := fr.uint8(ix)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint:
					y := basic.Uint(vy)
					return func(fr *frame) {
						x := fr.uint8(ix)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint8:
					y := basic.Uint8(vy)
					return func(fr *frame) {
						x := fr.uint8(ix)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint16:
					y := basic.Uint16(vy)
					return func(fr *frame) {
						x := fr.uint8(ix)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint32:
					y := basic.Uint32(vy)
					return func(fr *frame) {
						x := fr.uint8(ix)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint64:
					y := basic.Uint64(vy)
					return func(fr *frame) {
						x := fr.uint8(ix)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uintptr:
					y := basic.Uintptr(vy)
					return func(fr *frame) {
						x := fr.uint8(ix)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				}
			} else {
				switch ykind {
				case reflect.Int:
					return func(fr *frame) {
						x := fr.uint8(ix)
						y := fr.int(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int8:
					return func(fr *frame) {
						x := fr.uint8(ix)
						y := fr.int8(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int16:
					return func(fr *frame) {
						x := fr.uint8(ix)
						y := fr.int16(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int32:
					return func(fr *frame) {
						x := fr.uint8(ix)
						y := fr.int32(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int64:
					return func(fr *frame) {
						x := fr.uint8(ix)
						y := fr.int64(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint:
					return func(fr *frame) {
						x := fr.uint8(ix)
						y := fr.uint(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint8:
					return func(fr *frame) {
						x := fr.uint8(ix)
						y := fr.uint8(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint16:
					return func(fr *frame) {
						x := fr.uint8(ix)
						y := fr.uint16(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint32:
					return func(fr *frame) {
						x := fr.uint8(ix)
						y := fr.uint32(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint64:
					return func(fr *frame) {
						x := fr.uint8(ix)
						y := fr.uint64(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uintptr:
					return func(fr *frame) {
						x := fr.uint8(ix)
						y := fr.uintptr(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				}
			}
		case reflect.Uint16:
			if kx == kindConst {
				x := basic.Uint16(vx)
				switch ykind {
				case reflect.Int:
					return func(fr *frame) {
						y := fr.int(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int8:
					return func(fr *frame) {
						y := fr.int8(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int16:
					return func(fr *frame) {
						y := fr.int16(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int32:
					return func(fr *frame) {
						y := fr.int32(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int64:
					return func(fr *frame) {
						y := fr.int64(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint:
					return func(fr *frame) {
						y := fr.uint(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint8:
					return func(fr *frame) {
						y := fr.uint8(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint16:
					return func(fr *frame) {
						y := fr.uint16(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint32:
					return func(fr *frame) {
						y := fr.uint32(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint64:
					return func(fr *frame) {
						y := fr.uint64(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uintptr:
					return func(fr *frame) {
						y := fr.uintptr(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				}
			} else if ky == kindConst {
				switch ykind {
				case reflect.Int:
					y := basic.Int(vy)
					return func(fr *frame) {
						x := fr.uint16(ix)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int8:
					y := basic.Int8(vy)
					return func(fr *frame) {
						x := fr.uint16(ix)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int16:
					y := basic.Int16(vy)
					return func(fr *frame) {
						x := fr.uint16(ix)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int32:
					y := basic.Int32(vy)
					return func(fr *frame) {
						x := fr.uint16(ix)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int64:
					y := basic.Int64(vy)
					return func(fr *frame) {
						x := fr.uint16(ix)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint:
					y := basic.Uint(vy)
					return func(fr *frame) {
						x := fr.uint16(ix)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint8:
					y := basic.Uint8(vy)
					return func(fr *frame) {
						x := fr.uint16(ix)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint16:
					y := basic.Uint16(vy)
					return func(fr *frame) {
						x := fr.uint16(ix)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint32:
					y := basic.Uint32(vy)
					return func(fr *frame) {
						x := fr.uint16(ix)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint64:
					y := basic.Uint64(vy)
					return func(fr *frame) {
						x := fr.uint16(ix)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uintptr:
					y := basic.Uintptr(vy)
					return func(fr *frame) {
						x := fr.uint16(ix)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				}
			} else {
				switch ykind {
				case reflect.Int:
					return func(fr *frame) {
						x := fr.uint16(ix)
						y := fr.int(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int8:
					return func(fr *frame) {
						x := fr.uint16(ix)
						y := fr.int8(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int16:
					return func(fr *frame) {
						x := fr.uint16(ix)
						y := fr.int16(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int32:
					return func(fr *frame) {
						x := fr.uint16(ix)
						y := fr.int32(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int64:
					return func(fr *frame) {
						x := fr.uint16(ix)
						y := fr.int64(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint:
					return func(fr *frame) {
						x := fr.uint16(ix)
						y := fr.uint(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint8:
					return func(fr *frame) {
						x := fr.uint16(ix)
						y := fr.uint8(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint16:
					return func(fr *frame) {
						x := fr.uint16(ix)
						y := fr.uint16(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint32:
					return func(fr *frame) {
						x := fr.uint16(ix)
						y := fr.uint32(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint64:
					return func(fr *frame) {
						x := fr.uint16(ix)
						y := fr.uint64(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uintptr:
					return func(fr *frame) {
						x := fr.uint16(ix)
						y := fr.uintptr(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				}
			}
		case reflect.Uint32:
			if kx == kindConst {
				x := basic.Uint32(vx)
				switch ykind {
				case reflect.Int:
					return func(fr *frame) {
						y := fr.int(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int8:
					return func(fr *frame) {
						y := fr.int8(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int16:
					return func(fr *frame) {
						y := fr.int16(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int32:
					return func(fr *frame) {
						y := fr.int32(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int64:
					return func(fr *frame) {
						y := fr.int64(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint:
					return func(fr *frame) {
						y := fr.uint(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint8:
					return func(fr *frame) {
						y := fr.uint8(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint16:
					return func(fr *frame) {
						y := fr.uint16(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint32:
					return func(fr *frame) {
						y := fr.uint32(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint64:
					return func(fr *frame) {
						y := fr.uint64(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uintptr:
					return func(fr *frame) {
						y := fr.uintptr(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				}
			} else if ky == kindConst {
				switch ykind {
				case reflect.Int:
					y := basic.Int(vy)
					return func(fr *frame) {
						x := fr.uint32(ix)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int8:
					y := basic.Int8(vy)
					return func(fr *frame) {
						x := fr.uint32(ix)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int16:
					y := basic.Int16(vy)
					return func(fr *frame) {
						x := fr.uint32(ix)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int32:
					y := basic.Int32(vy)
					return func(fr *frame) {
						x := fr.uint32(ix)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int64:
					y := basic.Int64(vy)
					return func(fr *frame) {
						x := fr.uint32(ix)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint:
					y := basic.Uint(vy)
					return func(fr *frame) {
						x := fr.uint32(ix)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint8:
					y := basic.Uint8(vy)
					return func(fr *frame) {
						x := fr.uint32(ix)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint16:
					y := basic.Uint16(vy)
					return func(fr *frame) {
						x := fr.uint32(ix)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint32:
					y := basic.Uint32(vy)
					return func(fr *frame) {
						x := fr.uint32(ix)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint64:
					y := basic.Uint64(vy)
					return func(fr *frame) {
						x := fr.uint32(ix)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uintptr:
					y := basic.Uintptr(vy)
					return func(fr *frame) {
						x := fr.uint32(ix)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				}
			} else {
				switch ykind {
				case reflect.Int:
					return func(fr *frame) {
						x := fr.uint32(ix)
						y := fr.int(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int8:
					return func(fr *frame) {
						x := fr.uint32(ix)
						y := fr.int8(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int16:
					return func(fr *frame) {
						x := fr.uint32(ix)
						y := fr.int16(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int32:
					return func(fr *frame) {
						x := fr.uint32(ix)
						y := fr.int32(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int64:
					return func(fr *frame) {
						x := fr.uint32(ix)
						y := fr.int64(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint:
					return func(fr *frame) {
						x := fr.uint32(ix)
						y := fr.uint(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint8:
					return func(fr *frame) {
						x := fr.uint32(ix)
						y := fr.uint8(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint16:
					return func(fr *frame) {
						x := fr.uint32(ix)
						y := fr.uint16(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint32:
					return func(fr *frame) {
						x := fr.uint32(ix)
						y := fr.uint32(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint64:
					return func(fr *frame) {
						x := fr.uint32(ix)
						y := fr.uint64(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uintptr:
					return func(fr *frame) {
						x := fr.uint32(ix)
						y := fr.uintptr(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				}
			}
		case reflect.Uint64:
			if kx == kindConst {
				x := basic.Uint64(vx)
				switch ykind {
				case reflect.Int:
					return func(fr *frame) {
						y := fr.int(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int8:
					return func(fr *frame) {
						y := fr.int8(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int16:
					return func(fr *frame) {
						y := fr.int16(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int32:
					return func(fr *frame) {
						y := fr.int32(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int64:
					return func(fr *frame) {
						y := fr.int64(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint:
					return func(fr *frame) {
						y := fr.uint(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint8:
					return func(fr *frame) {
						y := fr.uint8(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint16:
					return func(fr *frame) {
						y := fr.uint16(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint32:
					return func(fr *frame) {
						y := fr.uint32(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint64:
					return func(fr *frame) {
						y := fr.uint64(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uintptr:
					return func(fr *frame) {
						y := fr.uintptr(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				}
			} else if ky == kindConst {
				switch ykind {
				case reflect.Int:
					y := basic.Int(vy)
					return func(fr *frame) {
						x := fr.uint64(ix)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int8:
					y := basic.Int8(vy)
					return func(fr *frame) {
						x := fr.uint64(ix)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int16:
					y := basic.Int16(vy)
					return func(fr *frame) {
						x := fr.uint64(ix)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int32:
					y := basic.Int32(vy)
					return func(fr *frame) {
						x := fr.uint64(ix)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int64:
					y := basic.Int64(vy)
					return func(fr *frame) {
						x := fr.uint64(ix)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint:
					y := basic.Uint(vy)
					return func(fr *frame) {
						x := fr.uint64(ix)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint8:
					y := basic.Uint8(vy)
					return func(fr *frame) {
						x := fr.uint64(ix)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint16:
					y := basic.Uint16(vy)
					return func(fr *frame) {
						x := fr.uint64(ix)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint32:
					y := basic.Uint32(vy)
					return func(fr *frame) {
						x := fr.uint64(ix)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint64:
					y := basic.Uint64(vy)
					return func(fr *frame) {
						x := fr.uint64(ix)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uintptr:
					y := basic.Uintptr(vy)
					return func(fr *frame) {
						x := fr.uint64(ix)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				}
			} else {
				switch ykind {
				case reflect.Int:
					return func(fr *frame) {
						x := fr.uint64(ix)
						y := fr.int(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int8:
					return func(fr *frame) {
						x := fr.uint64(ix)
						y := fr.int8(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int16:
					return func(fr *frame) {
						x := fr.uint64(ix)
						y := fr.int16(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int32:
					return func(fr *frame) {
						x := fr.uint64(ix)
						y := fr.int32(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int64:
					return func(fr *frame) {
						x := fr.uint64(ix)
						y := fr.int64(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint:
					return func(fr *frame) {
						x := fr.uint64(ix)
						y := fr.uint(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint8:
					return func(fr *frame) {
						x := fr.uint64(ix)
						y := fr.uint8(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint16:
					return func(fr *frame) {
						x := fr.uint64(ix)
						y := fr.uint16(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint32:
					return func(fr *frame) {
						x := fr.uint64(ix)
						y := fr.uint32(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint64:
					return func(fr *frame) {
						x := fr.uint64(ix)
						y := fr.uint64(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uintptr:
					return func(fr *frame) {
						x := fr.uint64(ix)
						y := fr.uintptr(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				}
			}
		case reflect.Uintptr:
			if kx == kindConst {
				x := basic.Uintptr(vx)
				switch ykind {
				case reflect.Int:
					return func(fr *frame) {
						y := fr.int(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int8:
					return func(fr *frame) {
						y := fr.int8(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int16:
					return func(fr *frame) {
						y := fr.int16(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int32:
					return func(fr *frame) {
						y := fr.int32(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int64:
					return func(fr *frame) {
						y := fr.int64(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint:
					return func(fr *frame) {
						y := fr.uint(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint8:
					return func(fr *frame) {
						y := fr.uint8(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint16:
					return func(fr *frame) {
						y := fr.uint16(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint32:
					return func(fr *frame) {
						y := fr.uint32(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint64:
					return func(fr *frame) {
						y := fr.uint64(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uintptr:
					return func(fr *frame) {
						y := fr.uintptr(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				}
			} else if ky == kindConst {
				switch ykind {
				case reflect.Int:
					y := basic.Int(vy)
					return func(fr *frame) {
						x := fr.uintptr(ix)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int8:
					y := basic.Int8(vy)
					return func(fr *frame) {
						x := fr.uintptr(ix)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int16:
					y := basic.Int16(vy)
					return func(fr *frame) {
						x := fr.uintptr(ix)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int32:
					y := basic.Int32(vy)
					return func(fr *frame) {
						x := fr.uintptr(ix)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int64:
					y := basic.Int64(vy)
					return func(fr *frame) {
						x := fr.uintptr(ix)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint:
					y := basic.Uint(vy)
					return func(fr *frame) {
						x := fr.uintptr(ix)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint8:
					y := basic.Uint8(vy)
					return func(fr *frame) {
						x := fr.uintptr(ix)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint16:
					y := basic.Uint16(vy)
					return func(fr *frame) {
						x := fr.uintptr(ix)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint32:
					y := basic.Uint32(vy)
					return func(fr *frame) {
						x := fr.uintptr(ix)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint64:
					y := basic.Uint64(vy)
					return func(fr *frame) {
						x := fr.uintptr(ix)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uintptr:
					y := basic.Uintptr(vy)
					return func(fr *frame) {
						x := fr.uintptr(ix)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				}
			} else {
				switch ykind {
				case reflect.Int:
					return func(fr *frame) {
						x := fr.uintptr(ix)
						y := fr.int(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int8:
					return func(fr *frame) {
						x := fr.uintptr(ix)
						y := fr.int8(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int16:
					return func(fr *frame) {
						x := fr.uintptr(ix)
						y := fr.int16(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int32:
					return func(fr *frame) {
						x := fr.uintptr(ix)
						y := fr.int32(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int64:
					return func(fr *frame) {
						x := fr.uintptr(ix)
						y := fr.int64(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint:
					return func(fr *frame) {
						x := fr.uintptr(ix)
						y := fr.uint(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint8:
					return func(fr *frame) {
						x := fr.uintptr(ix)
						y := fr.uint8(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint16:
					return func(fr *frame) {
						x := fr.uintptr(ix)
						y := fr.uint16(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint32:
					return func(fr *frame) {
						x := fr.uintptr(ix)
						y := fr.uint32(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint64:
					return func(fr *frame) {
						x := fr.uintptr(ix)
						y := fr.uint64(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uintptr:
					return func(fr *frame) {
						x := fr.uintptr(ix)
						y := fr.uintptr(iy)
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				}
			}
		}
	}
	panic("unreachable")
}

func makeBinOpSHR(pfn *function, instr *ssa.BinOp) func(fr *frame) {
	ir := pfn.regIndex(instr)
	ix, kx, vx := pfn.regIndex3(instr.X)
	iy, ky, vy := pfn.regIndex3(instr.Y)
	xtyp := pfn.Interp.preToType(instr.X.Type())
	ytyp := pfn.Interp.preToType(instr.Y.Type())
	xkind := xtyp.Kind()
	ykind := ytyp.Kind()
	if kx == kindConst && ky == kindConst {
		t := basic.TypeOfType(xtyp)
		switch xkind {
		case reflect.Int:
			x := basic.Int(vx)
			switch ykind {
			case reflect.Int:
				v := basic.Make(t, x>>basic.Int(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			case reflect.Int8:
				v := basic.Make(t, x>>basic.Int8(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			case reflect.Int16:
				v := basic.Make(t, x>>basic.Int16(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			case reflect.Int32:
				v := basic.Make(t, x>>basic.Int32(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			case reflect.Int64:
				v := basic.Make(t, x>>basic.Int64(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			case reflect.Uint:
				v := basic.Make(t, x>>basic.Uint(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			case reflect.Uint8:
				v := basic.Make(t, x>>basic.Uint8(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			case reflect.Uint16:
				v := basic.Make(t, x>>basic.Uint16(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			case reflect.Uint32:
				v := basic.Make(t, x>>basic.Uint32(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			case reflect.Uint64:
				v := basic.Make(t, x>>basic.Uint64(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			case reflect.Uintptr:
				v := basic.Make(t, x>>basic.Uintptr(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			}
		case reflect.Int8:
			x := basic.Int8(vx)
			switch ykind {
			case reflect.Int:
				v := basic.Make(t, x>>basic.Int(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			case reflect.Int8:
				v := basic.Make(t, x>>basic.Int8(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			case reflect.Int16:
				v := basic.Make(t, x>>basic.Int16(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			case reflect.Int32:
				v := basic.Make(t, x>>basic.Int32(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			case reflect.Int64:
				v := basic.Make(t, x>>basic.Int64(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			case reflect.Uint:
				v := basic.Make(t, x>>basic.Uint(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			case reflect.Uint8:
				v := basic.Make(t, x>>basic.Uint8(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			case reflect.Uint16:
				v := basic.Make(t, x>>basic.Uint16(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			case reflect.Uint32:
				v := basic.Make(t, x>>basic.Uint32(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			case reflect.Uint64:
				v := basic.Make(t, x>>basic.Uint64(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			case reflect.Uintptr:
				v := basic.Make(t, x>>basic.Uintptr(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			}
		case reflect.Int16:
			x := basic.Int16(vx)
			switch ykind {
			case reflect.Int:
				v := basic.Make(t, x>>basic.Int(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			case reflect.Int8:
				v := basic.Make(t, x>>basic.Int8(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			case reflect.Int16:
				v := basic.Make(t, x>>basic.Int16(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			case reflect.Int32:
				v := basic.Make(t, x>>basic.Int32(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			case reflect.Int64:
				v := basic.Make(t, x>>basic.Int64(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			case reflect.Uint:
				v := basic.Make(t, x>>basic.Uint(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			case reflect.Uint8:
				v := basic.Make(t, x>>basic.Uint8(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			case reflect.Uint16:
				v := basic.Make(t, x>>basic.Uint16(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			case reflect.Uint32:
				v := basic.Make(t, x>>basic.Uint32(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			case reflect.Uint64:
				v := basic.Make(t, x>>basic.Uint64(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			case reflect.Uintptr:
				v := basic.Make(t, x>>basic.Uintptr(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			}
		case reflect.Int32:
			x := basic.Int32(vx)
			switch ykind {
			case reflect.Int:
				v := basic.Make(t, x>>basic.Int(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			case reflect.Int8:
				v := basic.Make(t, x>>basic.Int8(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			case reflect.Int16:
				v := basic.Make(t, x>>basic.Int16(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			case reflect.Int32:
				v := basic.Make(t, x>>basic.Int32(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			case reflect.Int64:
				v := basic.Make(t, x>>basic.Int64(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			case reflect.Uint:
				v := basic.Make(t, x>>basic.Uint(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			case reflect.Uint8:
				v := basic.Make(t, x>>basic.Uint8(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			case reflect.Uint16:
				v := basic.Make(t, x>>basic.Uint16(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			case reflect.Uint32:
				v := basic.Make(t, x>>basic.Uint32(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			case reflect.Uint64:
				v := basic.Make(t, x>>basic.Uint64(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			case reflect.Uintptr:
				v := basic.Make(t, x>>basic.Uintptr(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			}
		case reflect.Int64:
			x := basic.Int64(vx)
			switch ykind {
			case reflect.Int:
				v := basic.Make(t, x>>basic.Int(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			case reflect.Int8:
				v := basic.Make(t, x>>basic.Int8(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			case reflect.Int16:
				v := basic.Make(t, x>>basic.Int16(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			case reflect.Int32:
				v := basic.Make(t, x>>basic.Int32(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			case reflect.Int64:
				v := basic.Make(t, x>>basic.Int64(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			case reflect.Uint:
				v := basic.Make(t, x>>basic.Uint(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			case reflect.Uint8:
				v := basic.Make(t, x>>basic.Uint8(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			case reflect.Uint16:
				v := basic.Make(t, x>>basic.Uint16(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			case reflect.Uint32:
				v := basic.Make(t, x>>basic.Uint32(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			case reflect.Uint64:
				v := basic.Make(t, x>>basic.Uint64(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			case reflect.Uintptr:
				v := basic.Make(t, x>>basic.Uintptr(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			}
		case reflect.Uint:
			x := basic.Uint(vx)
			switch ykind {
			case reflect.Int:
				v := basic.Make(t, x>>basic.Int(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			case reflect.Int8:
				v := basic.Make(t, x>>basic.Int8(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			case reflect.Int16:
				v := basic.Make(t, x>>basic.Int16(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			case reflect.Int32:
				v := basic.Make(t, x>>basic.Int32(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			case reflect.Int64:
				v := basic.Make(t, x>>basic.Int64(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			case reflect.Uint:
				v := basic.Make(t, x>>basic.Uint(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			case reflect.Uint8:
				v := basic.Make(t, x>>basic.Uint8(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			case reflect.Uint16:
				v := basic.Make(t, x>>basic.Uint16(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			case reflect.Uint32:
				v := basic.Make(t, x>>basic.Uint32(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			case reflect.Uint64:
				v := basic.Make(t, x>>basic.Uint64(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			case reflect.Uintptr:
				v := basic.Make(t, x>>basic.Uintptr(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			}
		case reflect.Uint8:
			x := basic.Uint8(vx)
			switch ykind {
			case reflect.Int:
				v := basic.Make(t, x>>basic.Int(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			case reflect.Int8:
				v := basic.Make(t, x>>basic.Int8(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			case reflect.Int16:
				v := basic.Make(t, x>>basic.Int16(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			case reflect.Int32:
				v := basic.Make(t, x>>basic.Int32(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			case reflect.Int64:
				v := basic.Make(t, x>>basic.Int64(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			case reflect.Uint:
				v := basic.Make(t, x>>basic.Uint(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			case reflect.Uint8:
				v := basic.Make(t, x>>basic.Uint8(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			case reflect.Uint16:
				v := basic.Make(t, x>>basic.Uint16(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			case reflect.Uint32:
				v := basic.Make(t, x>>basic.Uint32(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			case reflect.Uint64:
				v := basic.Make(t, x>>basic.Uint64(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			case reflect.Uintptr:
				v := basic.Make(t, x>>basic.Uintptr(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			}
		case reflect.Uint16:
			x := basic.Uint16(vx)
			switch ykind {
			case reflect.Int:
				v := basic.Make(t, x>>basic.Int(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			case reflect.Int8:
				v := basic.Make(t, x>>basic.Int8(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			case reflect.Int16:
				v := basic.Make(t, x>>basic.Int16(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			case reflect.Int32:
				v := basic.Make(t, x>>basic.Int32(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			case reflect.Int64:
				v := basic.Make(t, x>>basic.Int64(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			case reflect.Uint:
				v := basic.Make(t, x>>basic.Uint(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			case reflect.Uint8:
				v := basic.Make(t, x>>basic.Uint8(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			case reflect.Uint16:
				v := basic.Make(t, x>>basic.Uint16(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			case reflect.Uint32:
				v := basic.Make(t, x>>basic.Uint32(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			case reflect.Uint64:
				v := basic.Make(t, x>>basic.Uint64(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			case reflect.Uintptr:
				v := basic.Make(t, x>>basic.Uintptr(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			}
		case reflect.Uint32:
			x := basic.Uint32(vx)
			switch ykind {
			case reflect.Int:
				v := basic.Make(t, x>>basic.Int(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			case reflect.Int8:
				v := basic.Make(t, x>>basic.Int8(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			case reflect.Int16:
				v := basic.Make(t, x>>basic.Int16(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			case reflect.Int32:
				v := basic.Make(t, x>>basic.Int32(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			case reflect.Int64:
				v := basic.Make(t, x>>basic.Int64(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			case reflect.Uint:
				v := basic.Make(t, x>>basic.Uint(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			case reflect.Uint8:
				v := basic.Make(t, x>>basic.Uint8(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			case reflect.Uint16:
				v := basic.Make(t, x>>basic.Uint16(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			case reflect.Uint32:
				v := basic.Make(t, x>>basic.Uint32(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			case reflect.Uint64:
				v := basic.Make(t, x>>basic.Uint64(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			case reflect.Uintptr:
				v := basic.Make(t, x>>basic.Uintptr(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			}
		case reflect.Uint64:
			x := basic.Uint64(vx)
			switch ykind {
			case reflect.Int:
				v := basic.Make(t, x>>basic.Int(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			case reflect.Int8:
				v := basic.Make(t, x>>basic.Int8(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			case reflect.Int16:
				v := basic.Make(t, x>>basic.Int16(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			case reflect.Int32:
				v := basic.Make(t, x>>basic.Int32(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			case reflect.Int64:
				v := basic.Make(t, x>>basic.Int64(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			case reflect.Uint:
				v := basic.Make(t, x>>basic.Uint(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			case reflect.Uint8:
				v := basic.Make(t, x>>basic.Uint8(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			case reflect.Uint16:
				v := basic.Make(t, x>>basic.Uint16(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			case reflect.Uint32:
				v := basic.Make(t, x>>basic.Uint32(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			case reflect.Uint64:
				v := basic.Make(t, x>>basic.Uint64(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			case reflect.Uintptr:
				v := basic.Make(t, x>>basic.Uintptr(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			}
		case reflect.Uintptr:
			x := basic.Uintptr(vx)
			switch ykind {
			case reflect.Int:
				v := basic.Make(t, x>>basic.Int(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			case reflect.Int8:
				v := basic.Make(t, x>>basic.Int8(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			case reflect.Int16:
				v := basic.Make(t, x>>basic.Int16(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			case reflect.Int32:
				v := basic.Make(t, x>>basic.Int32(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			case reflect.Int64:
				v := basic.Make(t, x>>basic.Int64(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			case reflect.Uint:
				v := basic.Make(t, x>>basic.Uint(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			case reflect.Uint8:
				v := basic.Make(t, x>>basic.Uint8(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			case reflect.Uint16:
				v := basic.Make(t, x>>basic.Uint16(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			case reflect.Uint32:
				v := basic.Make(t, x>>basic.Uint32(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			case reflect.Uint64:
				v := basic.Make(t, x>>basic.Uint64(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			case reflect.Uintptr:
				v := basic.Make(t, x>>basic.Uintptr(vy))
				return func(fr *frame) {
					fr.setReg(ir, v)
				}
			}
		}
	}
	if xtyp.PkgPath() == "" {

		switch xkind {
		case reflect.Int:
			if kx == kindConst {
				x := vx.(int)
				switch ykind {
				case reflect.Int:
					return func(fr *frame) {
						y := fr.int(iy)
						fr.setReg(ir, x>>y)
					}
				case reflect.Int8:
					return func(fr *frame) {
						y := fr.int8(iy)
						fr.setReg(ir, x>>y)
					}
				case reflect.Int16:
					return func(fr *frame) {
						y := fr.int16(iy)
						fr.setReg(ir, x>>y)
					}
				case reflect.Int32:
					return func(fr *frame) {
						y := fr.int32(iy)
						fr.setReg(ir, x>>y)
					}
				case reflect.Int64:
					return func(fr *frame) {
						y := fr.int64(iy)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint:
					return func(fr *frame) {
						y := fr.uint(iy)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint8:
					return func(fr *frame) {
						y := fr.uint8(iy)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint16:
					return func(fr *frame) {
						y := fr.uint16(iy)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint32:
					return func(fr *frame) {
						y := fr.uint32(iy)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint64:
					return func(fr *frame) {
						y := fr.uint64(iy)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uintptr:
					return func(fr *frame) {
						y := fr.uintptr(iy)
						fr.setReg(ir, x>>y)
					}
				}
			} else if ky == kindConst {
				switch ykind {
				case reflect.Int:
					y := basic.Int(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(int)
						fr.setReg(ir, x>>y)
					}
				case reflect.Int8:
					y := basic.Int8(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(int)
						fr.setReg(ir, x>>y)
					}
				case reflect.Int16:
					y := basic.Int16(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(int)
						fr.setReg(ir, x>>y)
					}
				case reflect.Int32:
					y := basic.Int32(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(int)
						fr.setReg(ir, x>>y)
					}
				case reflect.Int64:
					y := basic.Int64(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(int)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint:
					y := basic.Uint(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(int)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint8:
					y := basic.Uint8(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(int)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint16:
					y := basic.Uint16(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(int)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint32:
					y := basic.Uint32(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(int)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint64:
					y := basic.Uint64(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(int)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uintptr:
					y := basic.Uintptr(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(int)
						fr.setReg(ir, x>>y)
					}
				}
			} else {
				switch ykind {
				case reflect.Int:
					return func(fr *frame) {
						x := fr.int(ix)
						y := fr.int(iy)
						fr.setReg(ir, x>>y)
					}
				case reflect.Int8:
					return func(fr *frame) {
						x := fr.int(ix)
						y := fr.int8(iy)
						fr.setReg(ir, x>>y)
					}
				case reflect.Int16:
					return func(fr *frame) {
						x := fr.int(ix)
						y := fr.int16(iy)
						fr.setReg(ir, x>>y)
					}
				case reflect.Int32:
					return func(fr *frame) {
						x := fr.int(ix)
						y := fr.int32(iy)
						fr.setReg(ir, x>>y)
					}
				case reflect.Int64:
					return func(fr *frame) {
						x := fr.int(ix)
						y := fr.int64(iy)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint:
					return func(fr *frame) {
						x := fr.int(ix)
						y := fr.uint(iy)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint8:
					return func(fr *frame) {
						x := fr.int(ix)
						y := fr.uint8(iy)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint16:
					return func(fr *frame) {
						x := fr.int(ix)
						y := fr.uint16(iy)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint32:
					return func(fr *frame) {
						x := fr.int(ix)
						y := fr.uint32(iy)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint64:
					return func(fr *frame) {
						x := fr.int(ix)
						y := fr.uint64(iy)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uintptr:
					return func(fr *frame) {
						x := fr.int(ix)
						y := fr.uintptr(iy)
						fr.setReg(ir, x>>y)
					}
				}
			}
		case reflect.Int8:
			if kx == kindConst {
				x := vx.(int8)
				switch ykind {
				case reflect.Int:
					return func(fr *frame) {
						y := fr.int(iy)
						fr.setReg(ir, x>>y)
					}
				case reflect.Int8:
					return func(fr *frame) {
						y := fr.int8(iy)
						fr.setReg(ir, x>>y)
					}
				case reflect.Int16:
					return func(fr *frame) {
						y := fr.int16(iy)
						fr.setReg(ir, x>>y)
					}
				case reflect.Int32:
					return func(fr *frame) {
						y := fr.int32(iy)
						fr.setReg(ir, x>>y)
					}
				case reflect.Int64:
					return func(fr *frame) {
						y := fr.int64(iy)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint:
					return func(fr *frame) {
						y := fr.uint(iy)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint8:
					return func(fr *frame) {
						y := fr.uint8(iy)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint16:
					return func(fr *frame) {
						y := fr.uint16(iy)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint32:
					return func(fr *frame) {
						y := fr.uint32(iy)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint64:
					return func(fr *frame) {
						y := fr.uint64(iy)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uintptr:
					return func(fr *frame) {
						y := fr.uintptr(iy)
						fr.setReg(ir, x>>y)
					}
				}
			} else if ky == kindConst {
				switch ykind {
				case reflect.Int:
					y := basic.Int(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(int8)
						fr.setReg(ir, x>>y)
					}
				case reflect.Int8:
					y := basic.Int8(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(int8)
						fr.setReg(ir, x>>y)
					}
				case reflect.Int16:
					y := basic.Int16(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(int8)
						fr.setReg(ir, x>>y)
					}
				case reflect.Int32:
					y := basic.Int32(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(int8)
						fr.setReg(ir, x>>y)
					}
				case reflect.Int64:
					y := basic.Int64(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(int8)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint:
					y := basic.Uint(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(int8)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint8:
					y := basic.Uint8(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(int8)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint16:
					y := basic.Uint16(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(int8)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint32:
					y := basic.Uint32(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(int8)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint64:
					y := basic.Uint64(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(int8)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uintptr:
					y := basic.Uintptr(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(int8)
						fr.setReg(ir, x>>y)
					}
				}
			} else {
				switch ykind {
				case reflect.Int:
					return func(fr *frame) {
						x := fr.int8(ix)
						y := fr.int(iy)
						fr.setReg(ir, x>>y)
					}
				case reflect.Int8:
					return func(fr *frame) {
						x := fr.int8(ix)
						y := fr.int8(iy)
						fr.setReg(ir, x>>y)
					}
				case reflect.Int16:
					return func(fr *frame) {
						x := fr.int8(ix)
						y := fr.int16(iy)
						fr.setReg(ir, x>>y)
					}
				case reflect.Int32:
					return func(fr *frame) {
						x := fr.int8(ix)
						y := fr.int32(iy)
						fr.setReg(ir, x>>y)
					}
				case reflect.Int64:
					return func(fr *frame) {
						x := fr.int8(ix)
						y := fr.int64(iy)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint:
					return func(fr *frame) {
						x := fr.int8(ix)
						y := fr.uint(iy)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint8:
					return func(fr *frame) {
						x := fr.int8(ix)
						y := fr.uint8(iy)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint16:
					return func(fr *frame) {
						x := fr.int8(ix)
						y := fr.uint16(iy)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint32:
					return func(fr *frame) {
						x := fr.int8(ix)
						y := fr.uint32(iy)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint64:
					return func(fr *frame) {
						x := fr.int8(ix)
						y := fr.uint64(iy)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uintptr:
					return func(fr *frame) {
						x := fr.int8(ix)
						y := fr.uintptr(iy)
						fr.setReg(ir, x>>y)
					}
				}
			}
		case reflect.Int16:
			if kx == kindConst {
				x := vx.(int16)
				switch ykind {
				case reflect.Int:
					return func(fr *frame) {
						y := fr.int(iy)
						fr.setReg(ir, x>>y)
					}
				case reflect.Int8:
					return func(fr *frame) {
						y := fr.int8(iy)
						fr.setReg(ir, x>>y)
					}
				case reflect.Int16:
					return func(fr *frame) {
						y := fr.int16(iy)
						fr.setReg(ir, x>>y)
					}
				case reflect.Int32:
					return func(fr *frame) {
						y := fr.int32(iy)
						fr.setReg(ir, x>>y)
					}
				case reflect.Int64:
					return func(fr *frame) {
						y := fr.int64(iy)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint:
					return func(fr *frame) {
						y := fr.uint(iy)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint8:
					return func(fr *frame) {
						y := fr.uint8(iy)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint16:
					return func(fr *frame) {
						y := fr.uint16(iy)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint32:
					return func(fr *frame) {
						y := fr.uint32(iy)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint64:
					return func(fr *frame) {
						y := fr.uint64(iy)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uintptr:
					return func(fr *frame) {
						y := fr.uintptr(iy)
						fr.setReg(ir, x>>y)
					}
				}
			} else if ky == kindConst {
				switch ykind {
				case reflect.Int:
					y := basic.Int(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(int16)
						fr.setReg(ir, x>>y)
					}
				case reflect.Int8:
					y := basic.Int8(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(int16)
						fr.setReg(ir, x>>y)
					}
				case reflect.Int16:
					y := basic.Int16(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(int16)
						fr.setReg(ir, x>>y)
					}
				case reflect.Int32:
					y := basic.Int32(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(int16)
						fr.setReg(ir, x>>y)
					}
				case reflect.Int64:
					y := basic.Int64(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(int16)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint:
					y := basic.Uint(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(int16)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint8:
					y := basic.Uint8(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(int16)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint16:
					y := basic.Uint16(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(int16)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint32:
					y := basic.Uint32(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(int16)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint64:
					y := basic.Uint64(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(int16)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uintptr:
					y := basic.Uintptr(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(int16)
						fr.setReg(ir, x>>y)
					}
				}
			} else {
				switch ykind {
				case reflect.Int:
					return func(fr *frame) {
						x := fr.int16(ix)
						y := fr.int(iy)
						fr.setReg(ir, x>>y)
					}
				case reflect.Int8:
					return func(fr *frame) {
						x := fr.int16(ix)
						y := fr.int8(iy)
						fr.setReg(ir, x>>y)
					}
				case reflect.Int16:
					return func(fr *frame) {
						x := fr.int16(ix)
						y := fr.int16(iy)
						fr.setReg(ir, x>>y)
					}
				case reflect.Int32:
					return func(fr *frame) {
						x := fr.int16(ix)
						y := fr.int32(iy)
						fr.setReg(ir, x>>y)
					}
				case reflect.Int64:
					return func(fr *frame) {
						x := fr.int16(ix)
						y := fr.int64(iy)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint:
					return func(fr *frame) {
						x := fr.int16(ix)
						y := fr.uint(iy)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint8:
					return func(fr *frame) {
						x := fr.int16(ix)
						y := fr.uint8(iy)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint16:
					return func(fr *frame) {
						x := fr.int16(ix)
						y := fr.uint16(iy)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint32:
					return func(fr *frame) {
						x := fr.int16(ix)
						y := fr.uint32(iy)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint64:
					return func(fr *frame) {
						x := fr.int16(ix)
						y := fr.uint64(iy)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uintptr:
					return func(fr *frame) {
						x := fr.int16(ix)
						y := fr.uintptr(iy)
						fr.setReg(ir, x>>y)
					}
				}
			}
		case reflect.Int32:
			if kx == kindConst {
				x := vx.(int32)
				switch ykind {
				case reflect.Int:
					return func(fr *frame) {
						y := fr.int(iy)
						fr.setReg(ir, x>>y)
					}
				case reflect.Int8:
					return func(fr *frame) {
						y := fr.int8(iy)
						fr.setReg(ir, x>>y)
					}
				case reflect.Int16:
					return func(fr *frame) {
						y := fr.int16(iy)
						fr.setReg(ir, x>>y)
					}
				case reflect.Int32:
					return func(fr *frame) {
						y := fr.int32(iy)
						fr.setReg(ir, x>>y)
					}
				case reflect.Int64:
					return func(fr *frame) {
						y := fr.int64(iy)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint:
					return func(fr *frame) {
						y := fr.uint(iy)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint8:
					return func(fr *frame) {
						y := fr.uint8(iy)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint16:
					return func(fr *frame) {
						y := fr.uint16(iy)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint32:
					return func(fr *frame) {
						y := fr.uint32(iy)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint64:
					return func(fr *frame) {
						y := fr.uint64(iy)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uintptr:
					return func(fr *frame) {
						y := fr.uintptr(iy)
						fr.setReg(ir, x>>y)
					}
				}
			} else if ky == kindConst {
				switch ykind {
				case reflect.Int:
					y := basic.Int(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(int32)
						fr.setReg(ir, x>>y)
					}
				case reflect.Int8:
					y := basic.Int8(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(int32)
						fr.setReg(ir, x>>y)
					}
				case reflect.Int16:
					y := basic.Int16(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(int32)
						fr.setReg(ir, x>>y)
					}
				case reflect.Int32:
					y := basic.Int32(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(int32)
						fr.setReg(ir, x>>y)
					}
				case reflect.Int64:
					y := basic.Int64(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(int32)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint:
					y := basic.Uint(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(int32)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint8:
					y := basic.Uint8(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(int32)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint16:
					y := basic.Uint16(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(int32)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint32:
					y := basic.Uint32(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(int32)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint64:
					y := basic.Uint64(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(int32)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uintptr:
					y := basic.Uintptr(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(int32)
						fr.setReg(ir, x>>y)
					}
				}
			} else {
				switch ykind {
				case reflect.Int:
					return func(fr *frame) {
						x := fr.int32(ix)
						y := fr.int(iy)
						fr.setReg(ir, x>>y)
					}
				case reflect.Int8:
					return func(fr *frame) {
						x := fr.int32(ix)
						y := fr.int8(iy)
						fr.setReg(ir, x>>y)
					}
				case reflect.Int16:
					return func(fr *frame) {
						x := fr.int32(ix)
						y := fr.int16(iy)
						fr.setReg(ir, x>>y)
					}
				case reflect.Int32:
					return func(fr *frame) {
						x := fr.int32(ix)
						y := fr.int32(iy)
						fr.setReg(ir, x>>y)
					}
				case reflect.Int64:
					return func(fr *frame) {
						x := fr.int32(ix)
						y := fr.int64(iy)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint:
					return func(fr *frame) {
						x := fr.int32(ix)
						y := fr.uint(iy)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint8:
					return func(fr *frame) {
						x := fr.int32(ix)
						y := fr.uint8(iy)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint16:
					return func(fr *frame) {
						x := fr.int32(ix)
						y := fr.uint16(iy)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint32:
					return func(fr *frame) {
						x := fr.int32(ix)
						y := fr.uint32(iy)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint64:
					return func(fr *frame) {
						x := fr.int32(ix)
						y := fr.uint64(iy)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uintptr:
					return func(fr *frame) {
						x := fr.int32(ix)
						y := fr.uintptr(iy)
						fr.setReg(ir, x>>y)
					}
				}
			}
		case reflect.Int64:
			if kx == kindConst {
				x := vx.(int64)
				switch ykind {
				case reflect.Int:
					return func(fr *frame) {
						y := fr.int(iy)
						fr.setReg(ir, x>>y)
					}
				case reflect.Int8:
					return func(fr *frame) {
						y := fr.int8(iy)
						fr.setReg(ir, x>>y)
					}
				case reflect.Int16:
					return func(fr *frame) {
						y := fr.int16(iy)
						fr.setReg(ir, x>>y)
					}
				case reflect.Int32:
					return func(fr *frame) {
						y := fr.int32(iy)
						fr.setReg(ir, x>>y)
					}
				case reflect.Int64:
					return func(fr *frame) {
						y := fr.int64(iy)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint:
					return func(fr *frame) {
						y := fr.uint(iy)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint8:
					return func(fr *frame) {
						y := fr.uint8(iy)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint16:
					return func(fr *frame) {
						y := fr.uint16(iy)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint32:
					return func(fr *frame) {
						y := fr.uint32(iy)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint64:
					return func(fr *frame) {
						y := fr.uint64(iy)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uintptr:
					return func(fr *frame) {
						y := fr.uintptr(iy)
						fr.setReg(ir, x>>y)
					}
				}
			} else if ky == kindConst {
				switch ykind {
				case reflect.Int:
					y := basic.Int(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(int64)
						fr.setReg(ir, x>>y)
					}
				case reflect.Int8:
					y := basic.Int8(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(int64)
						fr.setReg(ir, x>>y)
					}
				case reflect.Int16:
					y := basic.Int16(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(int64)
						fr.setReg(ir, x>>y)
					}
				case reflect.Int32:
					y := basic.Int32(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(int64)
						fr.setReg(ir, x>>y)
					}
				case reflect.Int64:
					y := basic.Int64(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(int64)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint:
					y := basic.Uint(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(int64)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint8:
					y := basic.Uint8(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(int64)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint16:
					y := basic.Uint16(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(int64)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint32:
					y := basic.Uint32(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(int64)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint64:
					y := basic.Uint64(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(int64)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uintptr:
					y := basic.Uintptr(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(int64)
						fr.setReg(ir, x>>y)
					}
				}
			} else {
				switch ykind {
				case reflect.Int:
					return func(fr *frame) {
						x := fr.int64(ix)
						y := fr.int(iy)
						fr.setReg(ir, x>>y)
					}
				case reflect.Int8:
					return func(fr *frame) {
						x := fr.int64(ix)
						y := fr.int8(iy)
						fr.setReg(ir, x>>y)
					}
				case reflect.Int16:
					return func(fr *frame) {
						x := fr.int64(ix)
						y := fr.int16(iy)
						fr.setReg(ir, x>>y)
					}
				case reflect.Int32:
					return func(fr *frame) {
						x := fr.int64(ix)
						y := fr.int32(iy)
						fr.setReg(ir, x>>y)
					}
				case reflect.Int64:
					return func(fr *frame) {
						x := fr.int64(ix)
						y := fr.int64(iy)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint:
					return func(fr *frame) {
						x := fr.int64(ix)
						y := fr.uint(iy)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint8:
					return func(fr *frame) {
						x := fr.int64(ix)
						y := fr.uint8(iy)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint16:
					return func(fr *frame) {
						x := fr.int64(ix)
						y := fr.uint16(iy)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint32:
					return func(fr *frame) {
						x := fr.int64(ix)
						y := fr.uint32(iy)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint64:
					return func(fr *frame) {
						x := fr.int64(ix)
						y := fr.uint64(iy)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uintptr:
					return func(fr *frame) {
						x := fr.int64(ix)
						y := fr.uintptr(iy)
						fr.setReg(ir, x>>y)
					}
				}
			}
		case reflect.Uint:
			if kx == kindConst {
				x := vx.(uint)
				switch ykind {
				case reflect.Int:
					return func(fr *frame) {
						y := fr.int(iy)
						fr.setReg(ir, x>>y)
					}
				case reflect.Int8:
					return func(fr *frame) {
						y := fr.int8(iy)
						fr.setReg(ir, x>>y)
					}
				case reflect.Int16:
					return func(fr *frame) {
						y := fr.int16(iy)
						fr.setReg(ir, x>>y)
					}
				case reflect.Int32:
					return func(fr *frame) {
						y := fr.int32(iy)
						fr.setReg(ir, x>>y)
					}
				case reflect.Int64:
					return func(fr *frame) {
						y := fr.int64(iy)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint:
					return func(fr *frame) {
						y := fr.uint(iy)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint8:
					return func(fr *frame) {
						y := fr.uint8(iy)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint16:
					return func(fr *frame) {
						y := fr.uint16(iy)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint32:
					return func(fr *frame) {
						y := fr.uint32(iy)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint64:
					return func(fr *frame) {
						y := fr.uint64(iy)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uintptr:
					return func(fr *frame) {
						y := fr.uintptr(iy)
						fr.setReg(ir, x>>y)
					}
				}
			} else if ky == kindConst {
				switch ykind {
				case reflect.Int:
					y := basic.Int(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(uint)
						fr.setReg(ir, x>>y)
					}
				case reflect.Int8:
					y := basic.Int8(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(uint)
						fr.setReg(ir, x>>y)
					}
				case reflect.Int16:
					y := basic.Int16(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(uint)
						fr.setReg(ir, x>>y)
					}
				case reflect.Int32:
					y := basic.Int32(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(uint)
						fr.setReg(ir, x>>y)
					}
				case reflect.Int64:
					y := basic.Int64(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(uint)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint:
					y := basic.Uint(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(uint)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint8:
					y := basic.Uint8(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(uint)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint16:
					y := basic.Uint16(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(uint)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint32:
					y := basic.Uint32(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(uint)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint64:
					y := basic.Uint64(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(uint)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uintptr:
					y := basic.Uintptr(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(uint)
						fr.setReg(ir, x>>y)
					}
				}
			} else {
				switch ykind {
				case reflect.Int:
					return func(fr *frame) {
						x := fr.uint(ix)
						y := fr.int(iy)
						fr.setReg(ir, x>>y)
					}
				case reflect.Int8:
					return func(fr *frame) {
						x := fr.uint(ix)
						y := fr.int8(iy)
						fr.setReg(ir, x>>y)
					}
				case reflect.Int16:
					return func(fr *frame) {
						x := fr.uint(ix)
						y := fr.int16(iy)
						fr.setReg(ir, x>>y)
					}
				case reflect.Int32:
					return func(fr *frame) {
						x := fr.uint(ix)
						y := fr.int32(iy)
						fr.setReg(ir, x>>y)
					}
				case reflect.Int64:
					return func(fr *frame) {
						x := fr.uint(ix)
						y := fr.int64(iy)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint:
					return func(fr *frame) {
						x := fr.uint(ix)
						y := fr.uint(iy)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint8:
					return func(fr *frame) {
						x := fr.uint(ix)
						y := fr.uint8(iy)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint16:
					return func(fr *frame) {
						x := fr.uint(ix)
						y := fr.uint16(iy)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint32:
					return func(fr *frame) {
						x := fr.uint(ix)
						y := fr.uint32(iy)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint64:
					return func(fr *frame) {
						x := fr.uint(ix)
						y := fr.uint64(iy)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uintptr:
					return func(fr *frame) {
						x := fr.uint(ix)
						y := fr.uintptr(iy)
						fr.setReg(ir, x>>y)
					}
				}
			}
		case reflect.Uint8:
			if kx == kindConst {
				x := vx.(uint8)
				switch ykind {
				case reflect.Int:
					return func(fr *frame) {
						y := fr.int(iy)
						fr.setReg(ir, x>>y)
					}
				case reflect.Int8:
					return func(fr *frame) {
						y := fr.int8(iy)
						fr.setReg(ir, x>>y)
					}
				case reflect.Int16:
					return func(fr *frame) {
						y := fr.int16(iy)
						fr.setReg(ir, x>>y)
					}
				case reflect.Int32:
					return func(fr *frame) {
						y := fr.int32(iy)
						fr.setReg(ir, x>>y)
					}
				case reflect.Int64:
					return func(fr *frame) {
						y := fr.int64(iy)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint:
					return func(fr *frame) {
						y := fr.uint(iy)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint8:
					return func(fr *frame) {
						y := fr.uint8(iy)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint16:
					return func(fr *frame) {
						y := fr.uint16(iy)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint32:
					return func(fr *frame) {
						y := fr.uint32(iy)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint64:
					return func(fr *frame) {
						y := fr.uint64(iy)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uintptr:
					return func(fr *frame) {
						y := fr.uintptr(iy)
						fr.setReg(ir, x>>y)
					}
				}
			} else if ky == kindConst {
				switch ykind {
				case reflect.Int:
					y := basic.Int(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(uint8)
						fr.setReg(ir, x>>y)
					}
				case reflect.Int8:
					y := basic.Int8(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(uint8)
						fr.setReg(ir, x>>y)
					}
				case reflect.Int16:
					y := basic.Int16(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(uint8)
						fr.setReg(ir, x>>y)
					}
				case reflect.Int32:
					y := basic.Int32(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(uint8)
						fr.setReg(ir, x>>y)
					}
				case reflect.Int64:
					y := basic.Int64(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(uint8)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint:
					y := basic.Uint(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(uint8)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint8:
					y := basic.Uint8(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(uint8)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint16:
					y := basic.Uint16(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(uint8)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint32:
					y := basic.Uint32(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(uint8)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint64:
					y := basic.Uint64(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(uint8)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uintptr:
					y := basic.Uintptr(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(uint8)
						fr.setReg(ir, x>>y)
					}
				}
			} else {
				switch ykind {
				case reflect.Int:
					return func(fr *frame) {
						x := fr.uint8(ix)
						y := fr.int(iy)
						fr.setReg(ir, x>>y)
					}
				case reflect.Int8:
					return func(fr *frame) {
						x := fr.uint8(ix)
						y := fr.int8(iy)
						fr.setReg(ir, x>>y)
					}
				case reflect.Int16:
					return func(fr *frame) {
						x := fr.uint8(ix)
						y := fr.int16(iy)
						fr.setReg(ir, x>>y)
					}
				case reflect.Int32:
					return func(fr *frame) {
						x := fr.uint8(ix)
						y := fr.int32(iy)
						fr.setReg(ir, x>>y)
					}
				case reflect.Int64:
					return func(fr *frame) {
						x := fr.uint8(ix)
						y := fr.int64(iy)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint:
					return func(fr *frame) {
						x := fr.uint8(ix)
						y := fr.uint(iy)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint8:
					return func(fr *frame) {
						x := fr.uint8(ix)
						y := fr.uint8(iy)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint16:
					return func(fr *frame) {
						x := fr.uint8(ix)
						y := fr.uint16(iy)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint32:
					return func(fr *frame) {
						x := fr.uint8(ix)
						y := fr.uint32(iy)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint64:
					return func(fr *frame) {
						x := fr.uint8(ix)
						y := fr.uint64(iy)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uintptr:
					return func(fr *frame) {
						x := fr.uint8(ix)
						y := fr.uintptr(iy)
						fr.setReg(ir, x>>y)
					}
				}
			}
		case reflect.Uint16:
			if kx == kindConst {
				x := vx.(uint16)
				switch ykind {
				case reflect.Int:
					return func(fr *frame) {
						y := fr.int(iy)
						fr.setReg(ir, x>>y)
					}
				case reflect.Int8:
					return func(fr *frame) {
						y := fr.int8(iy)
						fr.setReg(ir, x>>y)
					}
				case reflect.Int16:
					return func(fr *frame) {
						y := fr.int16(iy)
						fr.setReg(ir, x>>y)
					}
				case reflect.Int32:
					return func(fr *frame) {
						y := fr.int32(iy)
						fr.setReg(ir, x>>y)
					}
				case reflect.Int64:
					return func(fr *frame) {
						y := fr.int64(iy)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint:
					return func(fr *frame) {
						y := fr.uint(iy)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint8:
					return func(fr *frame) {
						y := fr.uint8(iy)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint16:
					return func(fr *frame) {
						y := fr.uint16(iy)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint32:
					return func(fr *frame) {
						y := fr.uint32(iy)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint64:
					return func(fr *frame) {
						y := fr.uint64(iy)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uintptr:
					return func(fr *frame) {
						y := fr.uintptr(iy)
						fr.setReg(ir, x>>y)
					}
				}
			} else if ky == kindConst {
				switch ykind {
				case reflect.Int:
					y := basic.Int(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(uint16)
						fr.setReg(ir, x>>y)
					}
				case reflect.Int8:
					y := basic.Int8(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(uint16)
						fr.setReg(ir, x>>y)
					}
				case reflect.Int16:
					y := basic.Int16(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(uint16)
						fr.setReg(ir, x>>y)
					}
				case reflect.Int32:
					y := basic.Int32(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(uint16)
						fr.setReg(ir, x>>y)
					}
				case reflect.Int64:
					y := basic.Int64(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(uint16)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint:
					y := basic.Uint(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(uint16)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint8:
					y := basic.Uint8(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(uint16)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint16:
					y := basic.Uint16(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(uint16)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint32:
					y := basic.Uint32(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(uint16)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint64:
					y := basic.Uint64(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(uint16)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uintptr:
					y := basic.Uintptr(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(uint16)
						fr.setReg(ir, x>>y)
					}
				}
			} else {
				switch ykind {
				case reflect.Int:
					return func(fr *frame) {
						x := fr.uint16(ix)
						y := fr.int(iy)
						fr.setReg(ir, x>>y)
					}
				case reflect.Int8:
					return func(fr *frame) {
						x := fr.uint16(ix)
						y := fr.int8(iy)
						fr.setReg(ir, x>>y)
					}
				case reflect.Int16:
					return func(fr *frame) {
						x := fr.uint16(ix)
						y := fr.int16(iy)
						fr.setReg(ir, x>>y)
					}
				case reflect.Int32:
					return func(fr *frame) {
						x := fr.uint16(ix)
						y := fr.int32(iy)
						fr.setReg(ir, x>>y)
					}
				case reflect.Int64:
					return func(fr *frame) {
						x := fr.uint16(ix)
						y := fr.int64(iy)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint:
					return func(fr *frame) {
						x := fr.uint16(ix)
						y := fr.uint(iy)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint8:
					return func(fr *frame) {
						x := fr.uint16(ix)
						y := fr.uint8(iy)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint16:
					return func(fr *frame) {
						x := fr.uint16(ix)
						y := fr.uint16(iy)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint32:
					return func(fr *frame) {
						x := fr.uint16(ix)
						y := fr.uint32(iy)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint64:
					return func(fr *frame) {
						x := fr.uint16(ix)
						y := fr.uint64(iy)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uintptr:
					return func(fr *frame) {
						x := fr.uint16(ix)
						y := fr.uintptr(iy)
						fr.setReg(ir, x>>y)
					}
				}
			}
		case reflect.Uint32:
			if kx == kindConst {
				x := vx.(uint32)
				switch ykind {
				case reflect.Int:
					return func(fr *frame) {
						y := fr.int(iy)
						fr.setReg(ir, x>>y)
					}
				case reflect.Int8:
					return func(fr *frame) {
						y := fr.int8(iy)
						fr.setReg(ir, x>>y)
					}
				case reflect.Int16:
					return func(fr *frame) {
						y := fr.int16(iy)
						fr.setReg(ir, x>>y)
					}
				case reflect.Int32:
					return func(fr *frame) {
						y := fr.int32(iy)
						fr.setReg(ir, x>>y)
					}
				case reflect.Int64:
					return func(fr *frame) {
						y := fr.int64(iy)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint:
					return func(fr *frame) {
						y := fr.uint(iy)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint8:
					return func(fr *frame) {
						y := fr.uint8(iy)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint16:
					return func(fr *frame) {
						y := fr.uint16(iy)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint32:
					return func(fr *frame) {
						y := fr.uint32(iy)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint64:
					return func(fr *frame) {
						y := fr.uint64(iy)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uintptr:
					return func(fr *frame) {
						y := fr.uintptr(iy)
						fr.setReg(ir, x>>y)
					}
				}
			} else if ky == kindConst {
				switch ykind {
				case reflect.Int:
					y := basic.Int(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(uint32)
						fr.setReg(ir, x>>y)
					}
				case reflect.Int8:
					y := basic.Int8(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(uint32)
						fr.setReg(ir, x>>y)
					}
				case reflect.Int16:
					y := basic.Int16(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(uint32)
						fr.setReg(ir, x>>y)
					}
				case reflect.Int32:
					y := basic.Int32(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(uint32)
						fr.setReg(ir, x>>y)
					}
				case reflect.Int64:
					y := basic.Int64(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(uint32)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint:
					y := basic.Uint(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(uint32)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint8:
					y := basic.Uint8(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(uint32)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint16:
					y := basic.Uint16(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(uint32)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint32:
					y := basic.Uint32(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(uint32)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint64:
					y := basic.Uint64(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(uint32)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uintptr:
					y := basic.Uintptr(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(uint32)
						fr.setReg(ir, x>>y)
					}
				}
			} else {
				switch ykind {
				case reflect.Int:
					return func(fr *frame) {
						x := fr.uint32(ix)
						y := fr.int(iy)
						fr.setReg(ir, x>>y)
					}
				case reflect.Int8:
					return func(fr *frame) {
						x := fr.uint32(ix)
						y := fr.int8(iy)
						fr.setReg(ir, x>>y)
					}
				case reflect.Int16:
					return func(fr *frame) {
						x := fr.uint32(ix)
						y := fr.int16(iy)
						fr.setReg(ir, x>>y)
					}
				case reflect.Int32:
					return func(fr *frame) {
						x := fr.uint32(ix)
						y := fr.int32(iy)
						fr.setReg(ir, x>>y)
					}
				case reflect.Int64:
					return func(fr *frame) {
						x := fr.uint32(ix)
						y := fr.int64(iy)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint:
					return func(fr *frame) {
						x := fr.uint32(ix)
						y := fr.uint(iy)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint8:
					return func(fr *frame) {
						x := fr.uint32(ix)
						y := fr.uint8(iy)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint16:
					return func(fr *frame) {
						x := fr.uint32(ix)
						y := fr.uint16(iy)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint32:
					return func(fr *frame) {
						x := fr.uint32(ix)
						y := fr.uint32(iy)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint64:
					return func(fr *frame) {
						x := fr.uint32(ix)
						y := fr.uint64(iy)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uintptr:
					return func(fr *frame) {
						x := fr.uint32(ix)
						y := fr.uintptr(iy)
						fr.setReg(ir, x>>y)
					}
				}
			}
		case reflect.Uint64:
			if kx == kindConst {
				x := vx.(uint64)
				switch ykind {
				case reflect.Int:
					return func(fr *frame) {
						y := fr.int(iy)
						fr.setReg(ir, x>>y)
					}
				case reflect.Int8:
					return func(fr *frame) {
						y := fr.int8(iy)
						fr.setReg(ir, x>>y)
					}
				case reflect.Int16:
					return func(fr *frame) {
						y := fr.int16(iy)
						fr.setReg(ir, x>>y)
					}
				case reflect.Int32:
					return func(fr *frame) {
						y := fr.int32(iy)
						fr.setReg(ir, x>>y)
					}
				case reflect.Int64:
					return func(fr *frame) {
						y := fr.int64(iy)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint:
					return func(fr *frame) {
						y := fr.uint(iy)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint8:
					return func(fr *frame) {
						y := fr.uint8(iy)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint16:
					return func(fr *frame) {
						y := fr.uint16(iy)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint32:
					return func(fr *frame) {
						y := fr.uint32(iy)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint64:
					return func(fr *frame) {
						y := fr.uint64(iy)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uintptr:
					return func(fr *frame) {
						y := fr.uintptr(iy)
						fr.setReg(ir, x>>y)
					}
				}
			} else if ky == kindConst {
				switch ykind {
				case reflect.Int:
					y := basic.Int(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(uint64)
						fr.setReg(ir, x>>y)
					}
				case reflect.Int8:
					y := basic.Int8(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(uint64)
						fr.setReg(ir, x>>y)
					}
				case reflect.Int16:
					y := basic.Int16(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(uint64)
						fr.setReg(ir, x>>y)
					}
				case reflect.Int32:
					y := basic.Int32(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(uint64)
						fr.setReg(ir, x>>y)
					}
				case reflect.Int64:
					y := basic.Int64(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(uint64)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint:
					y := basic.Uint(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(uint64)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint8:
					y := basic.Uint8(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(uint64)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint16:
					y := basic.Uint16(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(uint64)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint32:
					y := basic.Uint32(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(uint64)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint64:
					y := basic.Uint64(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(uint64)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uintptr:
					y := basic.Uintptr(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(uint64)
						fr.setReg(ir, x>>y)
					}
				}
			} else {
				switch ykind {
				case reflect.Int:
					return func(fr *frame) {
						x := fr.uint64(ix)
						y := fr.int(iy)
						fr.setReg(ir, x>>y)
					}
				case reflect.Int8:
					return func(fr *frame) {
						x := fr.uint64(ix)
						y := fr.int8(iy)
						fr.setReg(ir, x>>y)
					}
				case reflect.Int16:
					return func(fr *frame) {
						x := fr.uint64(ix)
						y := fr.int16(iy)
						fr.setReg(ir, x>>y)
					}
				case reflect.Int32:
					return func(fr *frame) {
						x := fr.uint64(ix)
						y := fr.int32(iy)
						fr.setReg(ir, x>>y)
					}
				case reflect.Int64:
					return func(fr *frame) {
						x := fr.uint64(ix)
						y := fr.int64(iy)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint:
					return func(fr *frame) {
						x := fr.uint64(ix)
						y := fr.uint(iy)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint8:
					return func(fr *frame) {
						x := fr.uint64(ix)
						y := fr.uint8(iy)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint16:
					return func(fr *frame) {
						x := fr.uint64(ix)
						y := fr.uint16(iy)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint32:
					return func(fr *frame) {
						x := fr.uint64(ix)
						y := fr.uint32(iy)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint64:
					return func(fr *frame) {
						x := fr.uint64(ix)
						y := fr.uint64(iy)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uintptr:
					return func(fr *frame) {
						x := fr.uint64(ix)
						y := fr.uintptr(iy)
						fr.setReg(ir, x>>y)
					}
				}
			}
		case reflect.Uintptr:
			if kx == kindConst {
				x := vx.(uintptr)
				switch ykind {
				case reflect.Int:
					return func(fr *frame) {
						y := fr.int(iy)
						fr.setReg(ir, x>>y)
					}
				case reflect.Int8:
					return func(fr *frame) {
						y := fr.int8(iy)
						fr.setReg(ir, x>>y)
					}
				case reflect.Int16:
					return func(fr *frame) {
						y := fr.int16(iy)
						fr.setReg(ir, x>>y)
					}
				case reflect.Int32:
					return func(fr *frame) {
						y := fr.int32(iy)
						fr.setReg(ir, x>>y)
					}
				case reflect.Int64:
					return func(fr *frame) {
						y := fr.int64(iy)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint:
					return func(fr *frame) {
						y := fr.uint(iy)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint8:
					return func(fr *frame) {
						y := fr.uint8(iy)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint16:
					return func(fr *frame) {
						y := fr.uint16(iy)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint32:
					return func(fr *frame) {
						y := fr.uint32(iy)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint64:
					return func(fr *frame) {
						y := fr.uint64(iy)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uintptr:
					return func(fr *frame) {
						y := fr.uintptr(iy)
						fr.setReg(ir, x>>y)
					}
				}
			} else if ky == kindConst {
				switch ykind {
				case reflect.Int:
					y := basic.Int(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(uintptr)
						fr.setReg(ir, x>>y)
					}
				case reflect.Int8:
					y := basic.Int8(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(uintptr)
						fr.setReg(ir, x>>y)
					}
				case reflect.Int16:
					y := basic.Int16(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(uintptr)
						fr.setReg(ir, x>>y)
					}
				case reflect.Int32:
					y := basic.Int32(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(uintptr)
						fr.setReg(ir, x>>y)
					}
				case reflect.Int64:
					y := basic.Int64(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(uintptr)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint:
					y := basic.Uint(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(uintptr)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint8:
					y := basic.Uint8(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(uintptr)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint16:
					y := basic.Uint16(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(uintptr)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint32:
					y := basic.Uint32(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(uintptr)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint64:
					y := basic.Uint64(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(uintptr)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uintptr:
					y := basic.Uintptr(vy)
					return func(fr *frame) {
						x := fr.reg(ix).(uintptr)
						fr.setReg(ir, x>>y)
					}
				}
			} else {
				switch ykind {
				case reflect.Int:
					return func(fr *frame) {
						x := fr.uintptr(ix)
						y := fr.int(iy)
						fr.setReg(ir, x>>y)
					}
				case reflect.Int8:
					return func(fr *frame) {
						x := fr.uintptr(ix)
						y := fr.int8(iy)
						fr.setReg(ir, x>>y)
					}
				case reflect.Int16:
					return func(fr *frame) {
						x := fr.uintptr(ix)
						y := fr.int16(iy)
						fr.setReg(ir, x>>y)
					}
				case reflect.Int32:
					return func(fr *frame) {
						x := fr.uintptr(ix)
						y := fr.int32(iy)
						fr.setReg(ir, x>>y)
					}
				case reflect.Int64:
					return func(fr *frame) {
						x := fr.uintptr(ix)
						y := fr.int64(iy)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint:
					return func(fr *frame) {
						x := fr.uintptr(ix)
						y := fr.uint(iy)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint8:
					return func(fr *frame) {
						x := fr.uintptr(ix)
						y := fr.uint8(iy)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint16:
					return func(fr *frame) {
						x := fr.uintptr(ix)
						y := fr.uint16(iy)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint32:
					return func(fr *frame) {
						x := fr.uintptr(ix)
						y := fr.uint32(iy)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint64:
					return func(fr *frame) {
						x := fr.uintptr(ix)
						y := fr.uint64(iy)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uintptr:
					return func(fr *frame) {
						x := fr.uintptr(ix)
						y := fr.uintptr(iy)
						fr.setReg(ir, x>>y)
					}
				}
			}
		}
	} else {
		t := basic.TypeOfType(xtyp)
		switch xkind {
		case reflect.Int:
			if kx == kindConst {
				x := basic.Int(vx)
				switch ykind {
				case reflect.Int:
					return func(fr *frame) {
						y := fr.int(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int8:
					return func(fr *frame) {
						y := fr.int8(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int16:
					return func(fr *frame) {
						y := fr.int16(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int32:
					return func(fr *frame) {
						y := fr.int32(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int64:
					return func(fr *frame) {
						y := fr.int64(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint:
					return func(fr *frame) {
						y := fr.uint(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint8:
					return func(fr *frame) {
						y := fr.uint8(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint16:
					return func(fr *frame) {
						y := fr.uint16(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint32:
					return func(fr *frame) {
						y := fr.uint32(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint64:
					return func(fr *frame) {
						y := fr.uint64(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uintptr:
					return func(fr *frame) {
						y := fr.uintptr(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				}
			} else if ky == kindConst {
				switch ykind {
				case reflect.Int:
					y := basic.Int(vy)
					return func(fr *frame) {
						x := fr.int(ix)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int8:
					y := basic.Int8(vy)
					return func(fr *frame) {
						x := fr.int(ix)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int16:
					y := basic.Int16(vy)
					return func(fr *frame) {
						x := fr.int(ix)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int32:
					y := basic.Int32(vy)
					return func(fr *frame) {
						x := fr.int(ix)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int64:
					y := basic.Int64(vy)
					return func(fr *frame) {
						x := fr.int(ix)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint:
					y := basic.Uint(vy)
					return func(fr *frame) {
						x := fr.int(ix)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint8:
					y := basic.Uint8(vy)
					return func(fr *frame) {
						x := fr.int(ix)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint16:
					y := basic.Uint16(vy)
					return func(fr *frame) {
						x := fr.int(ix)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint32:
					y := basic.Uint32(vy)
					return func(fr *frame) {
						x := fr.int(ix)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint64:
					y := basic.Uint64(vy)
					return func(fr *frame) {
						x := fr.int(ix)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uintptr:
					y := basic.Uintptr(vy)
					return func(fr *frame) {
						x := fr.int(ix)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				}
			} else {
				switch ykind {
				case reflect.Int:
					return func(fr *frame) {
						x := fr.int(ix)
						y := fr.int(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int8:
					return func(fr *frame) {
						x := fr.int(ix)
						y := fr.int8(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int16:
					return func(fr *frame) {
						x := fr.int(ix)
						y := fr.int16(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int32:
					return func(fr *frame) {
						x := fr.int(ix)
						y := fr.int32(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int64:
					return func(fr *frame) {
						x := fr.int(ix)
						y := fr.int64(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint:
					return func(fr *frame) {
						x := fr.int(ix)
						y := fr.uint(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint8:
					return func(fr *frame) {
						x := fr.int(ix)
						y := fr.uint8(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint16:
					return func(fr *frame) {
						x := fr.int(ix)
						y := fr.uint16(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint32:
					return func(fr *frame) {
						x := fr.int(ix)
						y := fr.uint32(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint64:
					return func(fr *frame) {
						x := fr.int(ix)
						y := fr.uint64(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uintptr:
					return func(fr *frame) {
						x := fr.int(ix)
						y := fr.uintptr(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				}
			}
		case reflect.Int8:
			if kx == kindConst {
				x := basic.Int8(vx)
				switch ykind {
				case reflect.Int:
					return func(fr *frame) {
						y := fr.int(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int8:
					return func(fr *frame) {
						y := fr.int8(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int16:
					return func(fr *frame) {
						y := fr.int16(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int32:
					return func(fr *frame) {
						y := fr.int32(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int64:
					return func(fr *frame) {
						y := fr.int64(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint:
					return func(fr *frame) {
						y := fr.uint(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint8:
					return func(fr *frame) {
						y := fr.uint8(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint16:
					return func(fr *frame) {
						y := fr.uint16(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint32:
					return func(fr *frame) {
						y := fr.uint32(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint64:
					return func(fr *frame) {
						y := fr.uint64(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uintptr:
					return func(fr *frame) {
						y := fr.uintptr(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				}
			} else if ky == kindConst {
				switch ykind {
				case reflect.Int:
					y := basic.Int(vy)
					return func(fr *frame) {
						x := fr.int8(ix)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int8:
					y := basic.Int8(vy)
					return func(fr *frame) {
						x := fr.int8(ix)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int16:
					y := basic.Int16(vy)
					return func(fr *frame) {
						x := fr.int8(ix)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int32:
					y := basic.Int32(vy)
					return func(fr *frame) {
						x := fr.int8(ix)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int64:
					y := basic.Int64(vy)
					return func(fr *frame) {
						x := fr.int8(ix)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint:
					y := basic.Uint(vy)
					return func(fr *frame) {
						x := fr.int8(ix)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint8:
					y := basic.Uint8(vy)
					return func(fr *frame) {
						x := fr.int8(ix)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint16:
					y := basic.Uint16(vy)
					return func(fr *frame) {
						x := fr.int8(ix)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint32:
					y := basic.Uint32(vy)
					return func(fr *frame) {
						x := fr.int8(ix)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint64:
					y := basic.Uint64(vy)
					return func(fr *frame) {
						x := fr.int8(ix)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uintptr:
					y := basic.Uintptr(vy)
					return func(fr *frame) {
						x := fr.int8(ix)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				}
			} else {
				switch ykind {
				case reflect.Int:
					return func(fr *frame) {
						x := fr.int8(ix)
						y := fr.int(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int8:
					return func(fr *frame) {
						x := fr.int8(ix)
						y := fr.int8(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int16:
					return func(fr *frame) {
						x := fr.int8(ix)
						y := fr.int16(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int32:
					return func(fr *frame) {
						x := fr.int8(ix)
						y := fr.int32(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int64:
					return func(fr *frame) {
						x := fr.int8(ix)
						y := fr.int64(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint:
					return func(fr *frame) {
						x := fr.int8(ix)
						y := fr.uint(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint8:
					return func(fr *frame) {
						x := fr.int8(ix)
						y := fr.uint8(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint16:
					return func(fr *frame) {
						x := fr.int8(ix)
						y := fr.uint16(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint32:
					return func(fr *frame) {
						x := fr.int8(ix)
						y := fr.uint32(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint64:
					return func(fr *frame) {
						x := fr.int8(ix)
						y := fr.uint64(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uintptr:
					return func(fr *frame) {
						x := fr.int8(ix)
						y := fr.uintptr(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				}
			}
		case reflect.Int16:
			if kx == kindConst {
				x := basic.Int16(vx)
				switch ykind {
				case reflect.Int:
					return func(fr *frame) {
						y := fr.int(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int8:
					return func(fr *frame) {
						y := fr.int8(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int16:
					return func(fr *frame) {
						y := fr.int16(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int32:
					return func(fr *frame) {
						y := fr.int32(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int64:
					return func(fr *frame) {
						y := fr.int64(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint:
					return func(fr *frame) {
						y := fr.uint(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint8:
					return func(fr *frame) {
						y := fr.uint8(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint16:
					return func(fr *frame) {
						y := fr.uint16(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint32:
					return func(fr *frame) {
						y := fr.uint32(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint64:
					return func(fr *frame) {
						y := fr.uint64(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uintptr:
					return func(fr *frame) {
						y := fr.uintptr(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				}
			} else if ky == kindConst {
				switch ykind {
				case reflect.Int:
					y := basic.Int(vy)
					return func(fr *frame) {
						x := fr.int16(ix)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int8:
					y := basic.Int8(vy)
					return func(fr *frame) {
						x := fr.int16(ix)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int16:
					y := basic.Int16(vy)
					return func(fr *frame) {
						x := fr.int16(ix)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int32:
					y := basic.Int32(vy)
					return func(fr *frame) {
						x := fr.int16(ix)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int64:
					y := basic.Int64(vy)
					return func(fr *frame) {
						x := fr.int16(ix)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint:
					y := basic.Uint(vy)
					return func(fr *frame) {
						x := fr.int16(ix)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint8:
					y := basic.Uint8(vy)
					return func(fr *frame) {
						x := fr.int16(ix)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint16:
					y := basic.Uint16(vy)
					return func(fr *frame) {
						x := fr.int16(ix)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint32:
					y := basic.Uint32(vy)
					return func(fr *frame) {
						x := fr.int16(ix)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint64:
					y := basic.Uint64(vy)
					return func(fr *frame) {
						x := fr.int16(ix)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uintptr:
					y := basic.Uintptr(vy)
					return func(fr *frame) {
						x := fr.int16(ix)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				}
			} else {
				switch ykind {
				case reflect.Int:
					return func(fr *frame) {
						x := fr.int16(ix)
						y := fr.int(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int8:
					return func(fr *frame) {
						x := fr.int16(ix)
						y := fr.int8(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int16:
					return func(fr *frame) {
						x := fr.int16(ix)
						y := fr.int16(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int32:
					return func(fr *frame) {
						x := fr.int16(ix)
						y := fr.int32(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int64:
					return func(fr *frame) {
						x := fr.int16(ix)
						y := fr.int64(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint:
					return func(fr *frame) {
						x := fr.int16(ix)
						y := fr.uint(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint8:
					return func(fr *frame) {
						x := fr.int16(ix)
						y := fr.uint8(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint16:
					return func(fr *frame) {
						x := fr.int16(ix)
						y := fr.uint16(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint32:
					return func(fr *frame) {
						x := fr.int16(ix)
						y := fr.uint32(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint64:
					return func(fr *frame) {
						x := fr.int16(ix)
						y := fr.uint64(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uintptr:
					return func(fr *frame) {
						x := fr.int16(ix)
						y := fr.uintptr(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				}
			}
		case reflect.Int32:
			if kx == kindConst {
				x := basic.Int32(vx)
				switch ykind {
				case reflect.Int:
					return func(fr *frame) {
						y := fr.int(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int8:
					return func(fr *frame) {
						y := fr.int8(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int16:
					return func(fr *frame) {
						y := fr.int16(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int32:
					return func(fr *frame) {
						y := fr.int32(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int64:
					return func(fr *frame) {
						y := fr.int64(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint:
					return func(fr *frame) {
						y := fr.uint(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint8:
					return func(fr *frame) {
						y := fr.uint8(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint16:
					return func(fr *frame) {
						y := fr.uint16(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint32:
					return func(fr *frame) {
						y := fr.uint32(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint64:
					return func(fr *frame) {
						y := fr.uint64(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uintptr:
					return func(fr *frame) {
						y := fr.uintptr(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				}
			} else if ky == kindConst {
				switch ykind {
				case reflect.Int:
					y := basic.Int(vy)
					return func(fr *frame) {
						x := fr.int32(ix)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int8:
					y := basic.Int8(vy)
					return func(fr *frame) {
						x := fr.int32(ix)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int16:
					y := basic.Int16(vy)
					return func(fr *frame) {
						x := fr.int32(ix)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int32:
					y := basic.Int32(vy)
					return func(fr *frame) {
						x := fr.int32(ix)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int64:
					y := basic.Int64(vy)
					return func(fr *frame) {
						x := fr.int32(ix)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint:
					y := basic.Uint(vy)
					return func(fr *frame) {
						x := fr.int32(ix)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint8:
					y := basic.Uint8(vy)
					return func(fr *frame) {
						x := fr.int32(ix)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint16:
					y := basic.Uint16(vy)
					return func(fr *frame) {
						x := fr.int32(ix)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint32:
					y := basic.Uint32(vy)
					return func(fr *frame) {
						x := fr.int32(ix)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint64:
					y := basic.Uint64(vy)
					return func(fr *frame) {
						x := fr.int32(ix)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uintptr:
					y := basic.Uintptr(vy)
					return func(fr *frame) {
						x := fr.int32(ix)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				}
			} else {
				switch ykind {
				case reflect.Int:
					return func(fr *frame) {
						x := fr.int32(ix)
						y := fr.int(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int8:
					return func(fr *frame) {
						x := fr.int32(ix)
						y := fr.int8(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int16:
					return func(fr *frame) {
						x := fr.int32(ix)
						y := fr.int16(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int32:
					return func(fr *frame) {
						x := fr.int32(ix)
						y := fr.int32(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int64:
					return func(fr *frame) {
						x := fr.int32(ix)
						y := fr.int64(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint:
					return func(fr *frame) {
						x := fr.int32(ix)
						y := fr.uint(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint8:
					return func(fr *frame) {
						x := fr.int32(ix)
						y := fr.uint8(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint16:
					return func(fr *frame) {
						x := fr.int32(ix)
						y := fr.uint16(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint32:
					return func(fr *frame) {
						x := fr.int32(ix)
						y := fr.uint32(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint64:
					return func(fr *frame) {
						x := fr.int32(ix)
						y := fr.uint64(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uintptr:
					return func(fr *frame) {
						x := fr.int32(ix)
						y := fr.uintptr(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				}
			}
		case reflect.Int64:
			if kx == kindConst {
				x := basic.Int64(vx)
				switch ykind {
				case reflect.Int:
					return func(fr *frame) {
						y := fr.int(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int8:
					return func(fr *frame) {
						y := fr.int8(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int16:
					return func(fr *frame) {
						y := fr.int16(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int32:
					return func(fr *frame) {
						y := fr.int32(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int64:
					return func(fr *frame) {
						y := fr.int64(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint:
					return func(fr *frame) {
						y := fr.uint(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint8:
					return func(fr *frame) {
						y := fr.uint8(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint16:
					return func(fr *frame) {
						y := fr.uint16(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint32:
					return func(fr *frame) {
						y := fr.uint32(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint64:
					return func(fr *frame) {
						y := fr.uint64(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uintptr:
					return func(fr *frame) {
						y := fr.uintptr(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				}
			} else if ky == kindConst {
				switch ykind {
				case reflect.Int:
					y := basic.Int(vy)
					return func(fr *frame) {
						x := fr.int64(ix)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int8:
					y := basic.Int8(vy)
					return func(fr *frame) {
						x := fr.int64(ix)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int16:
					y := basic.Int16(vy)
					return func(fr *frame) {
						x := fr.int64(ix)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int32:
					y := basic.Int32(vy)
					return func(fr *frame) {
						x := fr.int64(ix)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int64:
					y := basic.Int64(vy)
					return func(fr *frame) {
						x := fr.int64(ix)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint:
					y := basic.Uint(vy)
					return func(fr *frame) {
						x := fr.int64(ix)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint8:
					y := basic.Uint8(vy)
					return func(fr *frame) {
						x := fr.int64(ix)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint16:
					y := basic.Uint16(vy)
					return func(fr *frame) {
						x := fr.int64(ix)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint32:
					y := basic.Uint32(vy)
					return func(fr *frame) {
						x := fr.int64(ix)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint64:
					y := basic.Uint64(vy)
					return func(fr *frame) {
						x := fr.int64(ix)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uintptr:
					y := basic.Uintptr(vy)
					return func(fr *frame) {
						x := fr.int64(ix)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				}
			} else {
				switch ykind {
				case reflect.Int:
					return func(fr *frame) {
						x := fr.int64(ix)
						y := fr.int(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int8:
					return func(fr *frame) {
						x := fr.int64(ix)
						y := fr.int8(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int16:
					return func(fr *frame) {
						x := fr.int64(ix)
						y := fr.int16(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int32:
					return func(fr *frame) {
						x := fr.int64(ix)
						y := fr.int32(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int64:
					return func(fr *frame) {
						x := fr.int64(ix)
						y := fr.int64(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint:
					return func(fr *frame) {
						x := fr.int64(ix)
						y := fr.uint(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint8:
					return func(fr *frame) {
						x := fr.int64(ix)
						y := fr.uint8(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint16:
					return func(fr *frame) {
						x := fr.int64(ix)
						y := fr.uint16(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint32:
					return func(fr *frame) {
						x := fr.int64(ix)
						y := fr.uint32(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint64:
					return func(fr *frame) {
						x := fr.int64(ix)
						y := fr.uint64(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uintptr:
					return func(fr *frame) {
						x := fr.int64(ix)
						y := fr.uintptr(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				}
			}
		case reflect.Uint:
			if kx == kindConst {
				x := basic.Uint(vx)
				switch ykind {
				case reflect.Int:
					return func(fr *frame) {
						y := fr.int(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int8:
					return func(fr *frame) {
						y := fr.int8(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int16:
					return func(fr *frame) {
						y := fr.int16(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int32:
					return func(fr *frame) {
						y := fr.int32(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int64:
					return func(fr *frame) {
						y := fr.int64(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint:
					return func(fr *frame) {
						y := fr.uint(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint8:
					return func(fr *frame) {
						y := fr.uint8(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint16:
					return func(fr *frame) {
						y := fr.uint16(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint32:
					return func(fr *frame) {
						y := fr.uint32(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint64:
					return func(fr *frame) {
						y := fr.uint64(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uintptr:
					return func(fr *frame) {
						y := fr.uintptr(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				}
			} else if ky == kindConst {
				switch ykind {
				case reflect.Int:
					y := basic.Int(vy)
					return func(fr *frame) {
						x := fr.uint(ix)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int8:
					y := basic.Int8(vy)
					return func(fr *frame) {
						x := fr.uint(ix)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int16:
					y := basic.Int16(vy)
					return func(fr *frame) {
						x := fr.uint(ix)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int32:
					y := basic.Int32(vy)
					return func(fr *frame) {
						x := fr.uint(ix)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int64:
					y := basic.Int64(vy)
					return func(fr *frame) {
						x := fr.uint(ix)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint:
					y := basic.Uint(vy)
					return func(fr *frame) {
						x := fr.uint(ix)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint8:
					y := basic.Uint8(vy)
					return func(fr *frame) {
						x := fr.uint(ix)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint16:
					y := basic.Uint16(vy)
					return func(fr *frame) {
						x := fr.uint(ix)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint32:
					y := basic.Uint32(vy)
					return func(fr *frame) {
						x := fr.uint(ix)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint64:
					y := basic.Uint64(vy)
					return func(fr *frame) {
						x := fr.uint(ix)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uintptr:
					y := basic.Uintptr(vy)
					return func(fr *frame) {
						x := fr.uint(ix)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				}
			} else {
				switch ykind {
				case reflect.Int:
					return func(fr *frame) {
						x := fr.uint(ix)
						y := fr.int(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int8:
					return func(fr *frame) {
						x := fr.uint(ix)
						y := fr.int8(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int16:
					return func(fr *frame) {
						x := fr.uint(ix)
						y := fr.int16(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int32:
					return func(fr *frame) {
						x := fr.uint(ix)
						y := fr.int32(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int64:
					return func(fr *frame) {
						x := fr.uint(ix)
						y := fr.int64(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint:
					return func(fr *frame) {
						x := fr.uint(ix)
						y := fr.uint(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint8:
					return func(fr *frame) {
						x := fr.uint(ix)
						y := fr.uint8(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint16:
					return func(fr *frame) {
						x := fr.uint(ix)
						y := fr.uint16(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint32:
					return func(fr *frame) {
						x := fr.uint(ix)
						y := fr.uint32(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint64:
					return func(fr *frame) {
						x := fr.uint(ix)
						y := fr.uint64(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uintptr:
					return func(fr *frame) {
						x := fr.uint(ix)
						y := fr.uintptr(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				}
			}
		case reflect.Uint8:
			if kx == kindConst {
				x := basic.Uint8(vx)
				switch ykind {
				case reflect.Int:
					return func(fr *frame) {
						y := fr.int(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int8:
					return func(fr *frame) {
						y := fr.int8(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int16:
					return func(fr *frame) {
						y := fr.int16(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int32:
					return func(fr *frame) {
						y := fr.int32(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int64:
					return func(fr *frame) {
						y := fr.int64(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint:
					return func(fr *frame) {
						y := fr.uint(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint8:
					return func(fr *frame) {
						y := fr.uint8(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint16:
					return func(fr *frame) {
						y := fr.uint16(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint32:
					return func(fr *frame) {
						y := fr.uint32(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint64:
					return func(fr *frame) {
						y := fr.uint64(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uintptr:
					return func(fr *frame) {
						y := fr.uintptr(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				}
			} else if ky == kindConst {
				switch ykind {
				case reflect.Int:
					y := basic.Int(vy)
					return func(fr *frame) {
						x := fr.uint8(ix)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int8:
					y := basic.Int8(vy)
					return func(fr *frame) {
						x := fr.uint8(ix)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int16:
					y := basic.Int16(vy)
					return func(fr *frame) {
						x := fr.uint8(ix)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int32:
					y := basic.Int32(vy)
					return func(fr *frame) {
						x := fr.uint8(ix)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int64:
					y := basic.Int64(vy)
					return func(fr *frame) {
						x := fr.uint8(ix)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint:
					y := basic.Uint(vy)
					return func(fr *frame) {
						x := fr.uint8(ix)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint8:
					y := basic.Uint8(vy)
					return func(fr *frame) {
						x := fr.uint8(ix)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint16:
					y := basic.Uint16(vy)
					return func(fr *frame) {
						x := fr.uint8(ix)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint32:
					y := basic.Uint32(vy)
					return func(fr *frame) {
						x := fr.uint8(ix)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint64:
					y := basic.Uint64(vy)
					return func(fr *frame) {
						x := fr.uint8(ix)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uintptr:
					y := basic.Uintptr(vy)
					return func(fr *frame) {
						x := fr.uint8(ix)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				}
			} else {
				switch ykind {
				case reflect.Int:
					return func(fr *frame) {
						x := fr.uint8(ix)
						y := fr.int(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int8:
					return func(fr *frame) {
						x := fr.uint8(ix)
						y := fr.int8(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int16:
					return func(fr *frame) {
						x := fr.uint8(ix)
						y := fr.int16(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int32:
					return func(fr *frame) {
						x := fr.uint8(ix)
						y := fr.int32(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int64:
					return func(fr *frame) {
						x := fr.uint8(ix)
						y := fr.int64(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint:
					return func(fr *frame) {
						x := fr.uint8(ix)
						y := fr.uint(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint8:
					return func(fr *frame) {
						x := fr.uint8(ix)
						y := fr.uint8(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint16:
					return func(fr *frame) {
						x := fr.uint8(ix)
						y := fr.uint16(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint32:
					return func(fr *frame) {
						x := fr.uint8(ix)
						y := fr.uint32(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint64:
					return func(fr *frame) {
						x := fr.uint8(ix)
						y := fr.uint64(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uintptr:
					return func(fr *frame) {
						x := fr.uint8(ix)
						y := fr.uintptr(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				}
			}
		case reflect.Uint16:
			if kx == kindConst {
				x := basic.Uint16(vx)
				switch ykind {
				case reflect.Int:
					return func(fr *frame) {
						y := fr.int(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int8:
					return func(fr *frame) {
						y := fr.int8(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int16:
					return func(fr *frame) {
						y := fr.int16(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int32:
					return func(fr *frame) {
						y := fr.int32(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int64:
					return func(fr *frame) {
						y := fr.int64(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint:
					return func(fr *frame) {
						y := fr.uint(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint8:
					return func(fr *frame) {
						y := fr.uint8(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint16:
					return func(fr *frame) {
						y := fr.uint16(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint32:
					return func(fr *frame) {
						y := fr.uint32(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint64:
					return func(fr *frame) {
						y := fr.uint64(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uintptr:
					return func(fr *frame) {
						y := fr.uintptr(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				}
			} else if ky == kindConst {
				switch ykind {
				case reflect.Int:
					y := basic.Int(vy)
					return func(fr *frame) {
						x := fr.uint16(ix)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int8:
					y := basic.Int8(vy)
					return func(fr *frame) {
						x := fr.uint16(ix)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int16:
					y := basic.Int16(vy)
					return func(fr *frame) {
						x := fr.uint16(ix)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int32:
					y := basic.Int32(vy)
					return func(fr *frame) {
						x := fr.uint16(ix)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int64:
					y := basic.Int64(vy)
					return func(fr *frame) {
						x := fr.uint16(ix)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint:
					y := basic.Uint(vy)
					return func(fr *frame) {
						x := fr.uint16(ix)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint8:
					y := basic.Uint8(vy)
					return func(fr *frame) {
						x := fr.uint16(ix)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint16:
					y := basic.Uint16(vy)
					return func(fr *frame) {
						x := fr.uint16(ix)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint32:
					y := basic.Uint32(vy)
					return func(fr *frame) {
						x := fr.uint16(ix)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint64:
					y := basic.Uint64(vy)
					return func(fr *frame) {
						x := fr.uint16(ix)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uintptr:
					y := basic.Uintptr(vy)
					return func(fr *frame) {
						x := fr.uint16(ix)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				}
			} else {
				switch ykind {
				case reflect.Int:
					return func(fr *frame) {
						x := fr.uint16(ix)
						y := fr.int(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int8:
					return func(fr *frame) {
						x := fr.uint16(ix)
						y := fr.int8(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int16:
					return func(fr *frame) {
						x := fr.uint16(ix)
						y := fr.int16(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int32:
					return func(fr *frame) {
						x := fr.uint16(ix)
						y := fr.int32(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int64:
					return func(fr *frame) {
						x := fr.uint16(ix)
						y := fr.int64(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint:
					return func(fr *frame) {
						x := fr.uint16(ix)
						y := fr.uint(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint8:
					return func(fr *frame) {
						x := fr.uint16(ix)
						y := fr.uint8(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint16:
					return func(fr *frame) {
						x := fr.uint16(ix)
						y := fr.uint16(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint32:
					return func(fr *frame) {
						x := fr.uint16(ix)
						y := fr.uint32(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint64:
					return func(fr *frame) {
						x := fr.uint16(ix)
						y := fr.uint64(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uintptr:
					return func(fr *frame) {
						x := fr.uint16(ix)
						y := fr.uintptr(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				}
			}
		case reflect.Uint32:
			if kx == kindConst {
				x := basic.Uint32(vx)
				switch ykind {
				case reflect.Int:
					return func(fr *frame) {
						y := fr.int(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int8:
					return func(fr *frame) {
						y := fr.int8(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int16:
					return func(fr *frame) {
						y := fr.int16(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int32:
					return func(fr *frame) {
						y := fr.int32(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int64:
					return func(fr *frame) {
						y := fr.int64(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint:
					return func(fr *frame) {
						y := fr.uint(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint8:
					return func(fr *frame) {
						y := fr.uint8(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint16:
					return func(fr *frame) {
						y := fr.uint16(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint32:
					return func(fr *frame) {
						y := fr.uint32(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint64:
					return func(fr *frame) {
						y := fr.uint64(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uintptr:
					return func(fr *frame) {
						y := fr.uintptr(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				}
			} else if ky == kindConst {
				switch ykind {
				case reflect.Int:
					y := basic.Int(vy)
					return func(fr *frame) {
						x := fr.uint32(ix)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int8:
					y := basic.Int8(vy)
					return func(fr *frame) {
						x := fr.uint32(ix)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int16:
					y := basic.Int16(vy)
					return func(fr *frame) {
						x := fr.uint32(ix)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int32:
					y := basic.Int32(vy)
					return func(fr *frame) {
						x := fr.uint32(ix)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int64:
					y := basic.Int64(vy)
					return func(fr *frame) {
						x := fr.uint32(ix)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint:
					y := basic.Uint(vy)
					return func(fr *frame) {
						x := fr.uint32(ix)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint8:
					y := basic.Uint8(vy)
					return func(fr *frame) {
						x := fr.uint32(ix)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint16:
					y := basic.Uint16(vy)
					return func(fr *frame) {
						x := fr.uint32(ix)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint32:
					y := basic.Uint32(vy)
					return func(fr *frame) {
						x := fr.uint32(ix)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint64:
					y := basic.Uint64(vy)
					return func(fr *frame) {
						x := fr.uint32(ix)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uintptr:
					y := basic.Uintptr(vy)
					return func(fr *frame) {
						x := fr.uint32(ix)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				}
			} else {
				switch ykind {
				case reflect.Int:
					return func(fr *frame) {
						x := fr.uint32(ix)
						y := fr.int(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int8:
					return func(fr *frame) {
						x := fr.uint32(ix)
						y := fr.int8(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int16:
					return func(fr *frame) {
						x := fr.uint32(ix)
						y := fr.int16(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int32:
					return func(fr *frame) {
						x := fr.uint32(ix)
						y := fr.int32(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int64:
					return func(fr *frame) {
						x := fr.uint32(ix)
						y := fr.int64(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint:
					return func(fr *frame) {
						x := fr.uint32(ix)
						y := fr.uint(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint8:
					return func(fr *frame) {
						x := fr.uint32(ix)
						y := fr.uint8(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint16:
					return func(fr *frame) {
						x := fr.uint32(ix)
						y := fr.uint16(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint32:
					return func(fr *frame) {
						x := fr.uint32(ix)
						y := fr.uint32(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint64:
					return func(fr *frame) {
						x := fr.uint32(ix)
						y := fr.uint64(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uintptr:
					return func(fr *frame) {
						x := fr.uint32(ix)
						y := fr.uintptr(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				}
			}
		case reflect.Uint64:
			if kx == kindConst {
				x := basic.Uint64(vx)
				switch ykind {
				case reflect.Int:
					return func(fr *frame) {
						y := fr.int(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int8:
					return func(fr *frame) {
						y := fr.int8(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int16:
					return func(fr *frame) {
						y := fr.int16(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int32:
					return func(fr *frame) {
						y := fr.int32(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int64:
					return func(fr *frame) {
						y := fr.int64(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint:
					return func(fr *frame) {
						y := fr.uint(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint8:
					return func(fr *frame) {
						y := fr.uint8(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint16:
					return func(fr *frame) {
						y := fr.uint16(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint32:
					return func(fr *frame) {
						y := fr.uint32(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint64:
					return func(fr *frame) {
						y := fr.uint64(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uintptr:
					return func(fr *frame) {
						y := fr.uintptr(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				}
			} else if ky == kindConst {
				switch ykind {
				case reflect.Int:
					y := basic.Int(vy)
					return func(fr *frame) {
						x := fr.uint64(ix)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int8:
					y := basic.Int8(vy)
					return func(fr *frame) {
						x := fr.uint64(ix)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int16:
					y := basic.Int16(vy)
					return func(fr *frame) {
						x := fr.uint64(ix)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int32:
					y := basic.Int32(vy)
					return func(fr *frame) {
						x := fr.uint64(ix)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int64:
					y := basic.Int64(vy)
					return func(fr *frame) {
						x := fr.uint64(ix)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint:
					y := basic.Uint(vy)
					return func(fr *frame) {
						x := fr.uint64(ix)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint8:
					y := basic.Uint8(vy)
					return func(fr *frame) {
						x := fr.uint64(ix)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint16:
					y := basic.Uint16(vy)
					return func(fr *frame) {
						x := fr.uint64(ix)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint32:
					y := basic.Uint32(vy)
					return func(fr *frame) {
						x := fr.uint64(ix)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint64:
					y := basic.Uint64(vy)
					return func(fr *frame) {
						x := fr.uint64(ix)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uintptr:
					y := basic.Uintptr(vy)
					return func(fr *frame) {
						x := fr.uint64(ix)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				}
			} else {
				switch ykind {
				case reflect.Int:
					return func(fr *frame) {
						x := fr.uint64(ix)
						y := fr.int(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int8:
					return func(fr *frame) {
						x := fr.uint64(ix)
						y := fr.int8(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int16:
					return func(fr *frame) {
						x := fr.uint64(ix)
						y := fr.int16(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int32:
					return func(fr *frame) {
						x := fr.uint64(ix)
						y := fr.int32(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int64:
					return func(fr *frame) {
						x := fr.uint64(ix)
						y := fr.int64(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint:
					return func(fr *frame) {
						x := fr.uint64(ix)
						y := fr.uint(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint8:
					return func(fr *frame) {
						x := fr.uint64(ix)
						y := fr.uint8(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint16:
					return func(fr *frame) {
						x := fr.uint64(ix)
						y := fr.uint16(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint32:
					return func(fr *frame) {
						x := fr.uint64(ix)
						y := fr.uint32(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint64:
					return func(fr *frame) {
						x := fr.uint64(ix)
						y := fr.uint64(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uintptr:
					return func(fr *frame) {
						x := fr.uint64(ix)
						y := fr.uintptr(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				}
			}
		case reflect.Uintptr:
			if kx == kindConst {
				x := basic.Uintptr(vx)
				switch ykind {
				case reflect.Int:
					return func(fr *frame) {
						y := fr.int(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int8:
					return func(fr *frame) {
						y := fr.int8(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int16:
					return func(fr *frame) {
						y := fr.int16(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int32:
					return func(fr *frame) {
						y := fr.int32(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int64:
					return func(fr *frame) {
						y := fr.int64(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint:
					return func(fr *frame) {
						y := fr.uint(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint8:
					return func(fr *frame) {
						y := fr.uint8(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint16:
					return func(fr *frame) {
						y := fr.uint16(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint32:
					return func(fr *frame) {
						y := fr.uint32(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint64:
					return func(fr *frame) {
						y := fr.uint64(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uintptr:
					return func(fr *frame) {
						y := fr.uintptr(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				}
			} else if ky == kindConst {
				switch ykind {
				case reflect.Int:
					y := basic.Int(vy)
					return func(fr *frame) {
						x := fr.uintptr(ix)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int8:
					y := basic.Int8(vy)
					return func(fr *frame) {
						x := fr.uintptr(ix)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int16:
					y := basic.Int16(vy)
					return func(fr *frame) {
						x := fr.uintptr(ix)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int32:
					y := basic.Int32(vy)
					return func(fr *frame) {
						x := fr.uintptr(ix)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int64:
					y := basic.Int64(vy)
					return func(fr *frame) {
						x := fr.uintptr(ix)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint:
					y := basic.Uint(vy)
					return func(fr *frame) {
						x := fr.uintptr(ix)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint8:
					y := basic.Uint8(vy)
					return func(fr *frame) {
						x := fr.uintptr(ix)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint16:
					y := basic.Uint16(vy)
					return func(fr *frame) {
						x := fr.uintptr(ix)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint32:
					y := basic.Uint32(vy)
					return func(fr *frame) {
						x := fr.uintptr(ix)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint64:
					y := basic.Uint64(vy)
					return func(fr *frame) {
						x := fr.uintptr(ix)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uintptr:
					y := basic.Uintptr(vy)
					return func(fr *frame) {
						x := fr.uintptr(ix)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				}
			} else {
				switch ykind {
				case reflect.Int:
					return func(fr *frame) {
						x := fr.uintptr(ix)
						y := fr.int(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int8:
					return func(fr *frame) {
						x := fr.uintptr(ix)
						y := fr.int8(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int16:
					return func(fr *frame) {
						x := fr.uintptr(ix)
						y := fr.int16(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int32:
					return func(fr *frame) {
						x := fr.uintptr(ix)
						y := fr.int32(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int64:
					return func(fr *frame) {
						x := fr.uintptr(ix)
						y := fr.int64(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint:
					return func(fr *frame) {
						x := fr.uintptr(ix)
						y := fr.uint(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint8:
					return func(fr *frame) {
						x := fr.uintptr(ix)
						y := fr.uint8(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint16:
					return func(fr *frame) {
						x := fr.uintptr(ix)
						y := fr.uint16(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint32:
					return func(fr *frame) {
						x := fr.uintptr(ix)
						y := fr.uint32(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint64:
					return func(fr *frame) {
						x := fr.uintptr(ix)
						y := fr.uint64(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uintptr:
					return func(fr *frame) {
						x := fr.uintptr(ix)
						y := fr.uintptr(iy)
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				}
			}
		}
	}
	panic("unreachable")
}
