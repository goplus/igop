package gossa

import (
	"reflect"

	"github.com/goplus/gossa/basic"
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
	xIsBasic := xtyp.PkgPath() == ""
	yIsBasic := ytyp.PkgPath() == ""
	if xIsBasic {
		switch xkind {
		case reflect.Int:
			if kx == kindConst && ky == kindConst {
				x := vx.(int)
				switch ykind {
				case reflect.Int:
					var y int
					if yIsBasic {
						y = vy.(int)
					} else {
						y = basic.Int(vy)
					}
					v := x << y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Int8:
					var y int8
					if yIsBasic {
						y = vy.(int8)
					} else {
						y = basic.Int8(vy)
					}
					v := x << y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Int16:
					var y int16
					if yIsBasic {
						y = vy.(int16)
					} else {
						y = basic.Int16(vy)
					}
					v := x << y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Int32:
					var y int32
					if yIsBasic {
						y = vy.(int32)
					} else {
						y = basic.Int32(vy)
					}
					v := x << y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Int64:
					var y int64
					if yIsBasic {
						y = vy.(int64)
					} else {
						y = basic.Int64(vy)
					}
					v := x << y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uint:
					var y uint
					if yIsBasic {
						y = vy.(uint)
					} else {
						y = basic.Uint(vy)
					}
					v := x << y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uint8:
					var y uint8
					if yIsBasic {
						y = vy.(uint8)
					} else {
						y = basic.Uint8(vy)
					}
					v := x << y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uint16:
					var y uint16
					if yIsBasic {
						y = vy.(uint16)
					} else {
						y = basic.Uint16(vy)
					}
					v := x << y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uint32:
					var y uint32
					if yIsBasic {
						y = vy.(uint32)
					} else {
						y = basic.Uint32(vy)
					}
					v := x << y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uint64:
					var y uint64
					if yIsBasic {
						y = vy.(uint64)
					} else {
						y = basic.Uint64(vy)
					}
					v := x << y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uintptr:
					var y uintptr
					if yIsBasic {
						y = vy.(uintptr)
					} else {
						y = basic.Uintptr(vy)
					}
					v := x << y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				}
			} else if kx == kindConst {
				x := vx.(int)
				switch ykind {
				case reflect.Int:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(int)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						y := basic.Int(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				case reflect.Int8:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(int8)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						y := basic.Int8(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				case reflect.Int16:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(int16)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						y := basic.Int16(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				case reflect.Int32:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(int32)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						y := basic.Int32(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				case reflect.Int64:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(int64)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						y := basic.Int64(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uint)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						y := basic.Uint(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint8:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uint8)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						y := basic.Uint8(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint16:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uint16)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						y := basic.Uint16(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint32:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uint32)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						y := basic.Uint32(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint64:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uint64)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						y := basic.Uint64(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				case reflect.Uintptr:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uintptr)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						y := basic.Uintptr(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				}
			} else if ky == kindConst {
				switch ykind {
				case reflect.Int:
					var y int
					if yIsBasic {
						y = vy.(int)
					} else {
						y = basic.Int(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(int)
						fr.setReg(ir, x<<y)
					}
				case reflect.Int8:
					var y int8
					if yIsBasic {
						y = vy.(int8)
					} else {
						y = basic.Int8(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(int)
						fr.setReg(ir, x<<y)
					}
				case reflect.Int16:
					var y int16
					if yIsBasic {
						y = vy.(int16)
					} else {
						y = basic.Int16(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(int)
						fr.setReg(ir, x<<y)
					}
				case reflect.Int32:
					var y int32
					if yIsBasic {
						y = vy.(int32)
					} else {
						y = basic.Int32(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(int)
						fr.setReg(ir, x<<y)
					}
				case reflect.Int64:
					var y int64
					if yIsBasic {
						y = vy.(int64)
					} else {
						y = basic.Int64(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(int)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint:
					var y uint
					if yIsBasic {
						y = vy.(uint)
					} else {
						y = basic.Uint(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(int)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint8:
					var y uint8
					if yIsBasic {
						y = vy.(uint8)
					} else {
						y = basic.Uint8(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(int)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint16:
					var y uint16
					if yIsBasic {
						y = vy.(uint16)
					} else {
						y = basic.Uint16(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(int)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint32:
					var y uint32
					if yIsBasic {
						y = vy.(uint32)
					} else {
						y = basic.Uint32(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(int)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint64:
					var y uint64
					if yIsBasic {
						y = vy.(uint64)
					} else {
						y = basic.Uint64(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(int)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uintptr:
					var y uintptr
					if yIsBasic {
						y = vy.(uintptr)
					} else {
						y = basic.Uintptr(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(int)
						fr.setReg(ir, x<<y)
					}
				}
			} else {
				switch ykind {
				case reflect.Int:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(int)
							y := fr.reg(iy).(int)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(int)
						y := basic.Int(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				case reflect.Int8:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(int)
							y := fr.reg(iy).(int8)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(int)
						y := basic.Int8(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				case reflect.Int16:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(int)
							y := fr.reg(iy).(int16)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(int)
						y := basic.Int16(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				case reflect.Int32:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(int)
							y := fr.reg(iy).(int32)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(int)
						y := basic.Int32(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				case reflect.Int64:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(int)
							y := fr.reg(iy).(int64)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(int)
						y := basic.Int64(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(int)
							y := fr.reg(iy).(uint)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(int)
						y := basic.Uint(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint8:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(int)
							y := fr.reg(iy).(uint8)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(int)
						y := basic.Uint8(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint16:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(int)
							y := fr.reg(iy).(uint16)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(int)
						y := basic.Uint16(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint32:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(int)
							y := fr.reg(iy).(uint32)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(int)
						y := basic.Uint32(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint64:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(int)
							y := fr.reg(iy).(uint64)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(int)
						y := basic.Uint64(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				case reflect.Uintptr:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(int)
							y := fr.reg(iy).(uintptr)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(int)
						y := basic.Uintptr(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				}
			}
		case reflect.Int8:
			if kx == kindConst && ky == kindConst {
				x := vx.(int8)
				switch ykind {
				case reflect.Int:
					var y int
					if yIsBasic {
						y = vy.(int)
					} else {
						y = basic.Int(vy)
					}
					v := x << y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Int8:
					var y int8
					if yIsBasic {
						y = vy.(int8)
					} else {
						y = basic.Int8(vy)
					}
					v := x << y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Int16:
					var y int16
					if yIsBasic {
						y = vy.(int16)
					} else {
						y = basic.Int16(vy)
					}
					v := x << y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Int32:
					var y int32
					if yIsBasic {
						y = vy.(int32)
					} else {
						y = basic.Int32(vy)
					}
					v := x << y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Int64:
					var y int64
					if yIsBasic {
						y = vy.(int64)
					} else {
						y = basic.Int64(vy)
					}
					v := x << y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uint:
					var y uint
					if yIsBasic {
						y = vy.(uint)
					} else {
						y = basic.Uint(vy)
					}
					v := x << y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uint8:
					var y uint8
					if yIsBasic {
						y = vy.(uint8)
					} else {
						y = basic.Uint8(vy)
					}
					v := x << y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uint16:
					var y uint16
					if yIsBasic {
						y = vy.(uint16)
					} else {
						y = basic.Uint16(vy)
					}
					v := x << y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uint32:
					var y uint32
					if yIsBasic {
						y = vy.(uint32)
					} else {
						y = basic.Uint32(vy)
					}
					v := x << y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uint64:
					var y uint64
					if yIsBasic {
						y = vy.(uint64)
					} else {
						y = basic.Uint64(vy)
					}
					v := x << y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uintptr:
					var y uintptr
					if yIsBasic {
						y = vy.(uintptr)
					} else {
						y = basic.Uintptr(vy)
					}
					v := x << y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				}
			} else if kx == kindConst {
				x := vx.(int8)
				switch ykind {
				case reflect.Int:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(int)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						y := basic.Int(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				case reflect.Int8:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(int8)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						y := basic.Int8(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				case reflect.Int16:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(int16)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						y := basic.Int16(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				case reflect.Int32:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(int32)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						y := basic.Int32(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				case reflect.Int64:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(int64)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						y := basic.Int64(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uint)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						y := basic.Uint(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint8:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uint8)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						y := basic.Uint8(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint16:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uint16)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						y := basic.Uint16(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint32:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uint32)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						y := basic.Uint32(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint64:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uint64)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						y := basic.Uint64(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				case reflect.Uintptr:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uintptr)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						y := basic.Uintptr(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				}
			} else if ky == kindConst {
				switch ykind {
				case reflect.Int:
					var y int
					if yIsBasic {
						y = vy.(int)
					} else {
						y = basic.Int(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(int8)
						fr.setReg(ir, x<<y)
					}
				case reflect.Int8:
					var y int8
					if yIsBasic {
						y = vy.(int8)
					} else {
						y = basic.Int8(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(int8)
						fr.setReg(ir, x<<y)
					}
				case reflect.Int16:
					var y int16
					if yIsBasic {
						y = vy.(int16)
					} else {
						y = basic.Int16(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(int8)
						fr.setReg(ir, x<<y)
					}
				case reflect.Int32:
					var y int32
					if yIsBasic {
						y = vy.(int32)
					} else {
						y = basic.Int32(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(int8)
						fr.setReg(ir, x<<y)
					}
				case reflect.Int64:
					var y int64
					if yIsBasic {
						y = vy.(int64)
					} else {
						y = basic.Int64(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(int8)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint:
					var y uint
					if yIsBasic {
						y = vy.(uint)
					} else {
						y = basic.Uint(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(int8)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint8:
					var y uint8
					if yIsBasic {
						y = vy.(uint8)
					} else {
						y = basic.Uint8(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(int8)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint16:
					var y uint16
					if yIsBasic {
						y = vy.(uint16)
					} else {
						y = basic.Uint16(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(int8)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint32:
					var y uint32
					if yIsBasic {
						y = vy.(uint32)
					} else {
						y = basic.Uint32(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(int8)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint64:
					var y uint64
					if yIsBasic {
						y = vy.(uint64)
					} else {
						y = basic.Uint64(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(int8)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uintptr:
					var y uintptr
					if yIsBasic {
						y = vy.(uintptr)
					} else {
						y = basic.Uintptr(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(int8)
						fr.setReg(ir, x<<y)
					}
				}
			} else {
				switch ykind {
				case reflect.Int:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(int8)
							y := fr.reg(iy).(int)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(int8)
						y := basic.Int(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				case reflect.Int8:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(int8)
							y := fr.reg(iy).(int8)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(int8)
						y := basic.Int8(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				case reflect.Int16:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(int8)
							y := fr.reg(iy).(int16)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(int8)
						y := basic.Int16(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				case reflect.Int32:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(int8)
							y := fr.reg(iy).(int32)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(int8)
						y := basic.Int32(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				case reflect.Int64:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(int8)
							y := fr.reg(iy).(int64)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(int8)
						y := basic.Int64(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(int8)
							y := fr.reg(iy).(uint)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(int8)
						y := basic.Uint(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint8:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(int8)
							y := fr.reg(iy).(uint8)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(int8)
						y := basic.Uint8(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint16:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(int8)
							y := fr.reg(iy).(uint16)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(int8)
						y := basic.Uint16(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint32:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(int8)
							y := fr.reg(iy).(uint32)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(int8)
						y := basic.Uint32(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint64:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(int8)
							y := fr.reg(iy).(uint64)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(int8)
						y := basic.Uint64(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				case reflect.Uintptr:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(int8)
							y := fr.reg(iy).(uintptr)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(int8)
						y := basic.Uintptr(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				}
			}
		case reflect.Int16:
			if kx == kindConst && ky == kindConst {
				x := vx.(int16)
				switch ykind {
				case reflect.Int:
					var y int
					if yIsBasic {
						y = vy.(int)
					} else {
						y = basic.Int(vy)
					}
					v := x << y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Int8:
					var y int8
					if yIsBasic {
						y = vy.(int8)
					} else {
						y = basic.Int8(vy)
					}
					v := x << y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Int16:
					var y int16
					if yIsBasic {
						y = vy.(int16)
					} else {
						y = basic.Int16(vy)
					}
					v := x << y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Int32:
					var y int32
					if yIsBasic {
						y = vy.(int32)
					} else {
						y = basic.Int32(vy)
					}
					v := x << y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Int64:
					var y int64
					if yIsBasic {
						y = vy.(int64)
					} else {
						y = basic.Int64(vy)
					}
					v := x << y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uint:
					var y uint
					if yIsBasic {
						y = vy.(uint)
					} else {
						y = basic.Uint(vy)
					}
					v := x << y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uint8:
					var y uint8
					if yIsBasic {
						y = vy.(uint8)
					} else {
						y = basic.Uint8(vy)
					}
					v := x << y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uint16:
					var y uint16
					if yIsBasic {
						y = vy.(uint16)
					} else {
						y = basic.Uint16(vy)
					}
					v := x << y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uint32:
					var y uint32
					if yIsBasic {
						y = vy.(uint32)
					} else {
						y = basic.Uint32(vy)
					}
					v := x << y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uint64:
					var y uint64
					if yIsBasic {
						y = vy.(uint64)
					} else {
						y = basic.Uint64(vy)
					}
					v := x << y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uintptr:
					var y uintptr
					if yIsBasic {
						y = vy.(uintptr)
					} else {
						y = basic.Uintptr(vy)
					}
					v := x << y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				}
			} else if kx == kindConst {
				x := vx.(int16)
				switch ykind {
				case reflect.Int:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(int)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						y := basic.Int(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				case reflect.Int8:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(int8)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						y := basic.Int8(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				case reflect.Int16:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(int16)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						y := basic.Int16(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				case reflect.Int32:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(int32)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						y := basic.Int32(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				case reflect.Int64:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(int64)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						y := basic.Int64(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uint)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						y := basic.Uint(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint8:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uint8)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						y := basic.Uint8(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint16:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uint16)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						y := basic.Uint16(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint32:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uint32)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						y := basic.Uint32(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint64:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uint64)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						y := basic.Uint64(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				case reflect.Uintptr:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uintptr)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						y := basic.Uintptr(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				}
			} else if ky == kindConst {
				switch ykind {
				case reflect.Int:
					var y int
					if yIsBasic {
						y = vy.(int)
					} else {
						y = basic.Int(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(int16)
						fr.setReg(ir, x<<y)
					}
				case reflect.Int8:
					var y int8
					if yIsBasic {
						y = vy.(int8)
					} else {
						y = basic.Int8(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(int16)
						fr.setReg(ir, x<<y)
					}
				case reflect.Int16:
					var y int16
					if yIsBasic {
						y = vy.(int16)
					} else {
						y = basic.Int16(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(int16)
						fr.setReg(ir, x<<y)
					}
				case reflect.Int32:
					var y int32
					if yIsBasic {
						y = vy.(int32)
					} else {
						y = basic.Int32(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(int16)
						fr.setReg(ir, x<<y)
					}
				case reflect.Int64:
					var y int64
					if yIsBasic {
						y = vy.(int64)
					} else {
						y = basic.Int64(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(int16)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint:
					var y uint
					if yIsBasic {
						y = vy.(uint)
					} else {
						y = basic.Uint(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(int16)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint8:
					var y uint8
					if yIsBasic {
						y = vy.(uint8)
					} else {
						y = basic.Uint8(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(int16)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint16:
					var y uint16
					if yIsBasic {
						y = vy.(uint16)
					} else {
						y = basic.Uint16(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(int16)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint32:
					var y uint32
					if yIsBasic {
						y = vy.(uint32)
					} else {
						y = basic.Uint32(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(int16)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint64:
					var y uint64
					if yIsBasic {
						y = vy.(uint64)
					} else {
						y = basic.Uint64(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(int16)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uintptr:
					var y uintptr
					if yIsBasic {
						y = vy.(uintptr)
					} else {
						y = basic.Uintptr(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(int16)
						fr.setReg(ir, x<<y)
					}
				}
			} else {
				switch ykind {
				case reflect.Int:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(int16)
							y := fr.reg(iy).(int)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(int16)
						y := basic.Int(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				case reflect.Int8:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(int16)
							y := fr.reg(iy).(int8)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(int16)
						y := basic.Int8(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				case reflect.Int16:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(int16)
							y := fr.reg(iy).(int16)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(int16)
						y := basic.Int16(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				case reflect.Int32:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(int16)
							y := fr.reg(iy).(int32)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(int16)
						y := basic.Int32(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				case reflect.Int64:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(int16)
							y := fr.reg(iy).(int64)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(int16)
						y := basic.Int64(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(int16)
							y := fr.reg(iy).(uint)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(int16)
						y := basic.Uint(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint8:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(int16)
							y := fr.reg(iy).(uint8)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(int16)
						y := basic.Uint8(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint16:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(int16)
							y := fr.reg(iy).(uint16)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(int16)
						y := basic.Uint16(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint32:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(int16)
							y := fr.reg(iy).(uint32)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(int16)
						y := basic.Uint32(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint64:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(int16)
							y := fr.reg(iy).(uint64)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(int16)
						y := basic.Uint64(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				case reflect.Uintptr:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(int16)
							y := fr.reg(iy).(uintptr)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(int16)
						y := basic.Uintptr(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				}
			}
		case reflect.Int32:
			if kx == kindConst && ky == kindConst {
				x := vx.(int32)
				switch ykind {
				case reflect.Int:
					var y int
					if yIsBasic {
						y = vy.(int)
					} else {
						y = basic.Int(vy)
					}
					v := x << y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Int8:
					var y int8
					if yIsBasic {
						y = vy.(int8)
					} else {
						y = basic.Int8(vy)
					}
					v := x << y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Int16:
					var y int16
					if yIsBasic {
						y = vy.(int16)
					} else {
						y = basic.Int16(vy)
					}
					v := x << y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Int32:
					var y int32
					if yIsBasic {
						y = vy.(int32)
					} else {
						y = basic.Int32(vy)
					}
					v := x << y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Int64:
					var y int64
					if yIsBasic {
						y = vy.(int64)
					} else {
						y = basic.Int64(vy)
					}
					v := x << y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uint:
					var y uint
					if yIsBasic {
						y = vy.(uint)
					} else {
						y = basic.Uint(vy)
					}
					v := x << y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uint8:
					var y uint8
					if yIsBasic {
						y = vy.(uint8)
					} else {
						y = basic.Uint8(vy)
					}
					v := x << y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uint16:
					var y uint16
					if yIsBasic {
						y = vy.(uint16)
					} else {
						y = basic.Uint16(vy)
					}
					v := x << y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uint32:
					var y uint32
					if yIsBasic {
						y = vy.(uint32)
					} else {
						y = basic.Uint32(vy)
					}
					v := x << y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uint64:
					var y uint64
					if yIsBasic {
						y = vy.(uint64)
					} else {
						y = basic.Uint64(vy)
					}
					v := x << y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uintptr:
					var y uintptr
					if yIsBasic {
						y = vy.(uintptr)
					} else {
						y = basic.Uintptr(vy)
					}
					v := x << y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				}
			} else if kx == kindConst {
				x := vx.(int32)
				switch ykind {
				case reflect.Int:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(int)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						y := basic.Int(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				case reflect.Int8:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(int8)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						y := basic.Int8(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				case reflect.Int16:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(int16)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						y := basic.Int16(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				case reflect.Int32:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(int32)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						y := basic.Int32(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				case reflect.Int64:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(int64)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						y := basic.Int64(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uint)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						y := basic.Uint(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint8:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uint8)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						y := basic.Uint8(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint16:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uint16)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						y := basic.Uint16(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint32:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uint32)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						y := basic.Uint32(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint64:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uint64)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						y := basic.Uint64(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				case reflect.Uintptr:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uintptr)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						y := basic.Uintptr(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				}
			} else if ky == kindConst {
				switch ykind {
				case reflect.Int:
					var y int
					if yIsBasic {
						y = vy.(int)
					} else {
						y = basic.Int(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(int32)
						fr.setReg(ir, x<<y)
					}
				case reflect.Int8:
					var y int8
					if yIsBasic {
						y = vy.(int8)
					} else {
						y = basic.Int8(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(int32)
						fr.setReg(ir, x<<y)
					}
				case reflect.Int16:
					var y int16
					if yIsBasic {
						y = vy.(int16)
					} else {
						y = basic.Int16(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(int32)
						fr.setReg(ir, x<<y)
					}
				case reflect.Int32:
					var y int32
					if yIsBasic {
						y = vy.(int32)
					} else {
						y = basic.Int32(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(int32)
						fr.setReg(ir, x<<y)
					}
				case reflect.Int64:
					var y int64
					if yIsBasic {
						y = vy.(int64)
					} else {
						y = basic.Int64(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(int32)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint:
					var y uint
					if yIsBasic {
						y = vy.(uint)
					} else {
						y = basic.Uint(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(int32)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint8:
					var y uint8
					if yIsBasic {
						y = vy.(uint8)
					} else {
						y = basic.Uint8(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(int32)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint16:
					var y uint16
					if yIsBasic {
						y = vy.(uint16)
					} else {
						y = basic.Uint16(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(int32)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint32:
					var y uint32
					if yIsBasic {
						y = vy.(uint32)
					} else {
						y = basic.Uint32(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(int32)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint64:
					var y uint64
					if yIsBasic {
						y = vy.(uint64)
					} else {
						y = basic.Uint64(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(int32)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uintptr:
					var y uintptr
					if yIsBasic {
						y = vy.(uintptr)
					} else {
						y = basic.Uintptr(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(int32)
						fr.setReg(ir, x<<y)
					}
				}
			} else {
				switch ykind {
				case reflect.Int:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(int32)
							y := fr.reg(iy).(int)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(int32)
						y := basic.Int(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				case reflect.Int8:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(int32)
							y := fr.reg(iy).(int8)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(int32)
						y := basic.Int8(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				case reflect.Int16:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(int32)
							y := fr.reg(iy).(int16)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(int32)
						y := basic.Int16(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				case reflect.Int32:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(int32)
							y := fr.reg(iy).(int32)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(int32)
						y := basic.Int32(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				case reflect.Int64:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(int32)
							y := fr.reg(iy).(int64)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(int32)
						y := basic.Int64(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(int32)
							y := fr.reg(iy).(uint)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(int32)
						y := basic.Uint(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint8:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(int32)
							y := fr.reg(iy).(uint8)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(int32)
						y := basic.Uint8(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint16:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(int32)
							y := fr.reg(iy).(uint16)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(int32)
						y := basic.Uint16(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint32:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(int32)
							y := fr.reg(iy).(uint32)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(int32)
						y := basic.Uint32(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint64:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(int32)
							y := fr.reg(iy).(uint64)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(int32)
						y := basic.Uint64(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				case reflect.Uintptr:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(int32)
							y := fr.reg(iy).(uintptr)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(int32)
						y := basic.Uintptr(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				}
			}
		case reflect.Int64:
			if kx == kindConst && ky == kindConst {
				x := vx.(int64)
				switch ykind {
				case reflect.Int:
					var y int
					if yIsBasic {
						y = vy.(int)
					} else {
						y = basic.Int(vy)
					}
					v := x << y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Int8:
					var y int8
					if yIsBasic {
						y = vy.(int8)
					} else {
						y = basic.Int8(vy)
					}
					v := x << y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Int16:
					var y int16
					if yIsBasic {
						y = vy.(int16)
					} else {
						y = basic.Int16(vy)
					}
					v := x << y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Int32:
					var y int32
					if yIsBasic {
						y = vy.(int32)
					} else {
						y = basic.Int32(vy)
					}
					v := x << y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Int64:
					var y int64
					if yIsBasic {
						y = vy.(int64)
					} else {
						y = basic.Int64(vy)
					}
					v := x << y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uint:
					var y uint
					if yIsBasic {
						y = vy.(uint)
					} else {
						y = basic.Uint(vy)
					}
					v := x << y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uint8:
					var y uint8
					if yIsBasic {
						y = vy.(uint8)
					} else {
						y = basic.Uint8(vy)
					}
					v := x << y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uint16:
					var y uint16
					if yIsBasic {
						y = vy.(uint16)
					} else {
						y = basic.Uint16(vy)
					}
					v := x << y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uint32:
					var y uint32
					if yIsBasic {
						y = vy.(uint32)
					} else {
						y = basic.Uint32(vy)
					}
					v := x << y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uint64:
					var y uint64
					if yIsBasic {
						y = vy.(uint64)
					} else {
						y = basic.Uint64(vy)
					}
					v := x << y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uintptr:
					var y uintptr
					if yIsBasic {
						y = vy.(uintptr)
					} else {
						y = basic.Uintptr(vy)
					}
					v := x << y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				}
			} else if kx == kindConst {
				x := vx.(int64)
				switch ykind {
				case reflect.Int:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(int)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						y := basic.Int(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				case reflect.Int8:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(int8)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						y := basic.Int8(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				case reflect.Int16:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(int16)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						y := basic.Int16(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				case reflect.Int32:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(int32)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						y := basic.Int32(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				case reflect.Int64:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(int64)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						y := basic.Int64(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uint)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						y := basic.Uint(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint8:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uint8)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						y := basic.Uint8(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint16:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uint16)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						y := basic.Uint16(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint32:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uint32)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						y := basic.Uint32(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint64:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uint64)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						y := basic.Uint64(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				case reflect.Uintptr:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uintptr)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						y := basic.Uintptr(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				}
			} else if ky == kindConst {
				switch ykind {
				case reflect.Int:
					var y int
					if yIsBasic {
						y = vy.(int)
					} else {
						y = basic.Int(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(int64)
						fr.setReg(ir, x<<y)
					}
				case reflect.Int8:
					var y int8
					if yIsBasic {
						y = vy.(int8)
					} else {
						y = basic.Int8(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(int64)
						fr.setReg(ir, x<<y)
					}
				case reflect.Int16:
					var y int16
					if yIsBasic {
						y = vy.(int16)
					} else {
						y = basic.Int16(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(int64)
						fr.setReg(ir, x<<y)
					}
				case reflect.Int32:
					var y int32
					if yIsBasic {
						y = vy.(int32)
					} else {
						y = basic.Int32(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(int64)
						fr.setReg(ir, x<<y)
					}
				case reflect.Int64:
					var y int64
					if yIsBasic {
						y = vy.(int64)
					} else {
						y = basic.Int64(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(int64)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint:
					var y uint
					if yIsBasic {
						y = vy.(uint)
					} else {
						y = basic.Uint(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(int64)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint8:
					var y uint8
					if yIsBasic {
						y = vy.(uint8)
					} else {
						y = basic.Uint8(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(int64)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint16:
					var y uint16
					if yIsBasic {
						y = vy.(uint16)
					} else {
						y = basic.Uint16(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(int64)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint32:
					var y uint32
					if yIsBasic {
						y = vy.(uint32)
					} else {
						y = basic.Uint32(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(int64)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint64:
					var y uint64
					if yIsBasic {
						y = vy.(uint64)
					} else {
						y = basic.Uint64(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(int64)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uintptr:
					var y uintptr
					if yIsBasic {
						y = vy.(uintptr)
					} else {
						y = basic.Uintptr(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(int64)
						fr.setReg(ir, x<<y)
					}
				}
			} else {
				switch ykind {
				case reflect.Int:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(int64)
							y := fr.reg(iy).(int)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(int64)
						y := basic.Int(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				case reflect.Int8:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(int64)
							y := fr.reg(iy).(int8)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(int64)
						y := basic.Int8(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				case reflect.Int16:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(int64)
							y := fr.reg(iy).(int16)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(int64)
						y := basic.Int16(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				case reflect.Int32:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(int64)
							y := fr.reg(iy).(int32)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(int64)
						y := basic.Int32(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				case reflect.Int64:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(int64)
							y := fr.reg(iy).(int64)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(int64)
						y := basic.Int64(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(int64)
							y := fr.reg(iy).(uint)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(int64)
						y := basic.Uint(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint8:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(int64)
							y := fr.reg(iy).(uint8)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(int64)
						y := basic.Uint8(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint16:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(int64)
							y := fr.reg(iy).(uint16)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(int64)
						y := basic.Uint16(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint32:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(int64)
							y := fr.reg(iy).(uint32)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(int64)
						y := basic.Uint32(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint64:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(int64)
							y := fr.reg(iy).(uint64)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(int64)
						y := basic.Uint64(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				case reflect.Uintptr:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(int64)
							y := fr.reg(iy).(uintptr)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(int64)
						y := basic.Uintptr(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				}
			}
		case reflect.Uint:
			if kx == kindConst && ky == kindConst {
				x := vx.(uint)
				switch ykind {
				case reflect.Int:
					var y int
					if yIsBasic {
						y = vy.(int)
					} else {
						y = basic.Int(vy)
					}
					v := x << y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Int8:
					var y int8
					if yIsBasic {
						y = vy.(int8)
					} else {
						y = basic.Int8(vy)
					}
					v := x << y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Int16:
					var y int16
					if yIsBasic {
						y = vy.(int16)
					} else {
						y = basic.Int16(vy)
					}
					v := x << y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Int32:
					var y int32
					if yIsBasic {
						y = vy.(int32)
					} else {
						y = basic.Int32(vy)
					}
					v := x << y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Int64:
					var y int64
					if yIsBasic {
						y = vy.(int64)
					} else {
						y = basic.Int64(vy)
					}
					v := x << y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uint:
					var y uint
					if yIsBasic {
						y = vy.(uint)
					} else {
						y = basic.Uint(vy)
					}
					v := x << y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uint8:
					var y uint8
					if yIsBasic {
						y = vy.(uint8)
					} else {
						y = basic.Uint8(vy)
					}
					v := x << y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uint16:
					var y uint16
					if yIsBasic {
						y = vy.(uint16)
					} else {
						y = basic.Uint16(vy)
					}
					v := x << y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uint32:
					var y uint32
					if yIsBasic {
						y = vy.(uint32)
					} else {
						y = basic.Uint32(vy)
					}
					v := x << y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uint64:
					var y uint64
					if yIsBasic {
						y = vy.(uint64)
					} else {
						y = basic.Uint64(vy)
					}
					v := x << y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uintptr:
					var y uintptr
					if yIsBasic {
						y = vy.(uintptr)
					} else {
						y = basic.Uintptr(vy)
					}
					v := x << y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				}
			} else if kx == kindConst {
				x := vx.(uint)
				switch ykind {
				case reflect.Int:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(int)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						y := basic.Int(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				case reflect.Int8:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(int8)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						y := basic.Int8(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				case reflect.Int16:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(int16)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						y := basic.Int16(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				case reflect.Int32:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(int32)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						y := basic.Int32(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				case reflect.Int64:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(int64)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						y := basic.Int64(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uint)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						y := basic.Uint(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint8:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uint8)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						y := basic.Uint8(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint16:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uint16)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						y := basic.Uint16(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint32:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uint32)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						y := basic.Uint32(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint64:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uint64)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						y := basic.Uint64(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				case reflect.Uintptr:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uintptr)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						y := basic.Uintptr(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				}
			} else if ky == kindConst {
				switch ykind {
				case reflect.Int:
					var y int
					if yIsBasic {
						y = vy.(int)
					} else {
						y = basic.Int(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uint)
						fr.setReg(ir, x<<y)
					}
				case reflect.Int8:
					var y int8
					if yIsBasic {
						y = vy.(int8)
					} else {
						y = basic.Int8(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uint)
						fr.setReg(ir, x<<y)
					}
				case reflect.Int16:
					var y int16
					if yIsBasic {
						y = vy.(int16)
					} else {
						y = basic.Int16(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uint)
						fr.setReg(ir, x<<y)
					}
				case reflect.Int32:
					var y int32
					if yIsBasic {
						y = vy.(int32)
					} else {
						y = basic.Int32(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uint)
						fr.setReg(ir, x<<y)
					}
				case reflect.Int64:
					var y int64
					if yIsBasic {
						y = vy.(int64)
					} else {
						y = basic.Int64(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uint)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint:
					var y uint
					if yIsBasic {
						y = vy.(uint)
					} else {
						y = basic.Uint(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uint)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint8:
					var y uint8
					if yIsBasic {
						y = vy.(uint8)
					} else {
						y = basic.Uint8(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uint)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint16:
					var y uint16
					if yIsBasic {
						y = vy.(uint16)
					} else {
						y = basic.Uint16(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uint)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint32:
					var y uint32
					if yIsBasic {
						y = vy.(uint32)
					} else {
						y = basic.Uint32(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uint)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint64:
					var y uint64
					if yIsBasic {
						y = vy.(uint64)
					} else {
						y = basic.Uint64(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uint)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uintptr:
					var y uintptr
					if yIsBasic {
						y = vy.(uintptr)
					} else {
						y = basic.Uintptr(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uint)
						fr.setReg(ir, x<<y)
					}
				}
			} else {
				switch ykind {
				case reflect.Int:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(uint)
							y := fr.reg(iy).(int)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uint)
						y := basic.Int(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				case reflect.Int8:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(uint)
							y := fr.reg(iy).(int8)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uint)
						y := basic.Int8(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				case reflect.Int16:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(uint)
							y := fr.reg(iy).(int16)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uint)
						y := basic.Int16(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				case reflect.Int32:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(uint)
							y := fr.reg(iy).(int32)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uint)
						y := basic.Int32(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				case reflect.Int64:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(uint)
							y := fr.reg(iy).(int64)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uint)
						y := basic.Int64(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(uint)
							y := fr.reg(iy).(uint)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uint)
						y := basic.Uint(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint8:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(uint)
							y := fr.reg(iy).(uint8)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uint)
						y := basic.Uint8(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint16:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(uint)
							y := fr.reg(iy).(uint16)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uint)
						y := basic.Uint16(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint32:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(uint)
							y := fr.reg(iy).(uint32)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uint)
						y := basic.Uint32(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint64:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(uint)
							y := fr.reg(iy).(uint64)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uint)
						y := basic.Uint64(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				case reflect.Uintptr:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(uint)
							y := fr.reg(iy).(uintptr)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uint)
						y := basic.Uintptr(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				}
			}
		case reflect.Uint8:
			if kx == kindConst && ky == kindConst {
				x := vx.(uint8)
				switch ykind {
				case reflect.Int:
					var y int
					if yIsBasic {
						y = vy.(int)
					} else {
						y = basic.Int(vy)
					}
					v := x << y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Int8:
					var y int8
					if yIsBasic {
						y = vy.(int8)
					} else {
						y = basic.Int8(vy)
					}
					v := x << y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Int16:
					var y int16
					if yIsBasic {
						y = vy.(int16)
					} else {
						y = basic.Int16(vy)
					}
					v := x << y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Int32:
					var y int32
					if yIsBasic {
						y = vy.(int32)
					} else {
						y = basic.Int32(vy)
					}
					v := x << y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Int64:
					var y int64
					if yIsBasic {
						y = vy.(int64)
					} else {
						y = basic.Int64(vy)
					}
					v := x << y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uint:
					var y uint
					if yIsBasic {
						y = vy.(uint)
					} else {
						y = basic.Uint(vy)
					}
					v := x << y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uint8:
					var y uint8
					if yIsBasic {
						y = vy.(uint8)
					} else {
						y = basic.Uint8(vy)
					}
					v := x << y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uint16:
					var y uint16
					if yIsBasic {
						y = vy.(uint16)
					} else {
						y = basic.Uint16(vy)
					}
					v := x << y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uint32:
					var y uint32
					if yIsBasic {
						y = vy.(uint32)
					} else {
						y = basic.Uint32(vy)
					}
					v := x << y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uint64:
					var y uint64
					if yIsBasic {
						y = vy.(uint64)
					} else {
						y = basic.Uint64(vy)
					}
					v := x << y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uintptr:
					var y uintptr
					if yIsBasic {
						y = vy.(uintptr)
					} else {
						y = basic.Uintptr(vy)
					}
					v := x << y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				}
			} else if kx == kindConst {
				x := vx.(uint8)
				switch ykind {
				case reflect.Int:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(int)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						y := basic.Int(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				case reflect.Int8:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(int8)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						y := basic.Int8(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				case reflect.Int16:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(int16)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						y := basic.Int16(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				case reflect.Int32:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(int32)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						y := basic.Int32(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				case reflect.Int64:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(int64)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						y := basic.Int64(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uint)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						y := basic.Uint(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint8:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uint8)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						y := basic.Uint8(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint16:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uint16)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						y := basic.Uint16(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint32:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uint32)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						y := basic.Uint32(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint64:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uint64)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						y := basic.Uint64(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				case reflect.Uintptr:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uintptr)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						y := basic.Uintptr(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				}
			} else if ky == kindConst {
				switch ykind {
				case reflect.Int:
					var y int
					if yIsBasic {
						y = vy.(int)
					} else {
						y = basic.Int(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uint8)
						fr.setReg(ir, x<<y)
					}
				case reflect.Int8:
					var y int8
					if yIsBasic {
						y = vy.(int8)
					} else {
						y = basic.Int8(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uint8)
						fr.setReg(ir, x<<y)
					}
				case reflect.Int16:
					var y int16
					if yIsBasic {
						y = vy.(int16)
					} else {
						y = basic.Int16(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uint8)
						fr.setReg(ir, x<<y)
					}
				case reflect.Int32:
					var y int32
					if yIsBasic {
						y = vy.(int32)
					} else {
						y = basic.Int32(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uint8)
						fr.setReg(ir, x<<y)
					}
				case reflect.Int64:
					var y int64
					if yIsBasic {
						y = vy.(int64)
					} else {
						y = basic.Int64(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uint8)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint:
					var y uint
					if yIsBasic {
						y = vy.(uint)
					} else {
						y = basic.Uint(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uint8)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint8:
					var y uint8
					if yIsBasic {
						y = vy.(uint8)
					} else {
						y = basic.Uint8(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uint8)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint16:
					var y uint16
					if yIsBasic {
						y = vy.(uint16)
					} else {
						y = basic.Uint16(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uint8)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint32:
					var y uint32
					if yIsBasic {
						y = vy.(uint32)
					} else {
						y = basic.Uint32(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uint8)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint64:
					var y uint64
					if yIsBasic {
						y = vy.(uint64)
					} else {
						y = basic.Uint64(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uint8)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uintptr:
					var y uintptr
					if yIsBasic {
						y = vy.(uintptr)
					} else {
						y = basic.Uintptr(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uint8)
						fr.setReg(ir, x<<y)
					}
				}
			} else {
				switch ykind {
				case reflect.Int:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(uint8)
							y := fr.reg(iy).(int)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uint8)
						y := basic.Int(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				case reflect.Int8:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(uint8)
							y := fr.reg(iy).(int8)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uint8)
						y := basic.Int8(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				case reflect.Int16:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(uint8)
							y := fr.reg(iy).(int16)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uint8)
						y := basic.Int16(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				case reflect.Int32:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(uint8)
							y := fr.reg(iy).(int32)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uint8)
						y := basic.Int32(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				case reflect.Int64:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(uint8)
							y := fr.reg(iy).(int64)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uint8)
						y := basic.Int64(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(uint8)
							y := fr.reg(iy).(uint)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uint8)
						y := basic.Uint(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint8:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(uint8)
							y := fr.reg(iy).(uint8)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uint8)
						y := basic.Uint8(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint16:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(uint8)
							y := fr.reg(iy).(uint16)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uint8)
						y := basic.Uint16(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint32:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(uint8)
							y := fr.reg(iy).(uint32)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uint8)
						y := basic.Uint32(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint64:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(uint8)
							y := fr.reg(iy).(uint64)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uint8)
						y := basic.Uint64(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				case reflect.Uintptr:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(uint8)
							y := fr.reg(iy).(uintptr)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uint8)
						y := basic.Uintptr(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				}
			}
		case reflect.Uint16:
			if kx == kindConst && ky == kindConst {
				x := vx.(uint16)
				switch ykind {
				case reflect.Int:
					var y int
					if yIsBasic {
						y = vy.(int)
					} else {
						y = basic.Int(vy)
					}
					v := x << y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Int8:
					var y int8
					if yIsBasic {
						y = vy.(int8)
					} else {
						y = basic.Int8(vy)
					}
					v := x << y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Int16:
					var y int16
					if yIsBasic {
						y = vy.(int16)
					} else {
						y = basic.Int16(vy)
					}
					v := x << y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Int32:
					var y int32
					if yIsBasic {
						y = vy.(int32)
					} else {
						y = basic.Int32(vy)
					}
					v := x << y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Int64:
					var y int64
					if yIsBasic {
						y = vy.(int64)
					} else {
						y = basic.Int64(vy)
					}
					v := x << y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uint:
					var y uint
					if yIsBasic {
						y = vy.(uint)
					} else {
						y = basic.Uint(vy)
					}
					v := x << y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uint8:
					var y uint8
					if yIsBasic {
						y = vy.(uint8)
					} else {
						y = basic.Uint8(vy)
					}
					v := x << y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uint16:
					var y uint16
					if yIsBasic {
						y = vy.(uint16)
					} else {
						y = basic.Uint16(vy)
					}
					v := x << y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uint32:
					var y uint32
					if yIsBasic {
						y = vy.(uint32)
					} else {
						y = basic.Uint32(vy)
					}
					v := x << y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uint64:
					var y uint64
					if yIsBasic {
						y = vy.(uint64)
					} else {
						y = basic.Uint64(vy)
					}
					v := x << y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uintptr:
					var y uintptr
					if yIsBasic {
						y = vy.(uintptr)
					} else {
						y = basic.Uintptr(vy)
					}
					v := x << y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				}
			} else if kx == kindConst {
				x := vx.(uint16)
				switch ykind {
				case reflect.Int:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(int)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						y := basic.Int(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				case reflect.Int8:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(int8)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						y := basic.Int8(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				case reflect.Int16:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(int16)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						y := basic.Int16(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				case reflect.Int32:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(int32)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						y := basic.Int32(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				case reflect.Int64:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(int64)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						y := basic.Int64(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uint)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						y := basic.Uint(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint8:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uint8)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						y := basic.Uint8(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint16:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uint16)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						y := basic.Uint16(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint32:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uint32)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						y := basic.Uint32(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint64:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uint64)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						y := basic.Uint64(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				case reflect.Uintptr:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uintptr)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						y := basic.Uintptr(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				}
			} else if ky == kindConst {
				switch ykind {
				case reflect.Int:
					var y int
					if yIsBasic {
						y = vy.(int)
					} else {
						y = basic.Int(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uint16)
						fr.setReg(ir, x<<y)
					}
				case reflect.Int8:
					var y int8
					if yIsBasic {
						y = vy.(int8)
					} else {
						y = basic.Int8(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uint16)
						fr.setReg(ir, x<<y)
					}
				case reflect.Int16:
					var y int16
					if yIsBasic {
						y = vy.(int16)
					} else {
						y = basic.Int16(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uint16)
						fr.setReg(ir, x<<y)
					}
				case reflect.Int32:
					var y int32
					if yIsBasic {
						y = vy.(int32)
					} else {
						y = basic.Int32(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uint16)
						fr.setReg(ir, x<<y)
					}
				case reflect.Int64:
					var y int64
					if yIsBasic {
						y = vy.(int64)
					} else {
						y = basic.Int64(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uint16)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint:
					var y uint
					if yIsBasic {
						y = vy.(uint)
					} else {
						y = basic.Uint(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uint16)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint8:
					var y uint8
					if yIsBasic {
						y = vy.(uint8)
					} else {
						y = basic.Uint8(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uint16)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint16:
					var y uint16
					if yIsBasic {
						y = vy.(uint16)
					} else {
						y = basic.Uint16(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uint16)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint32:
					var y uint32
					if yIsBasic {
						y = vy.(uint32)
					} else {
						y = basic.Uint32(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uint16)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint64:
					var y uint64
					if yIsBasic {
						y = vy.(uint64)
					} else {
						y = basic.Uint64(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uint16)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uintptr:
					var y uintptr
					if yIsBasic {
						y = vy.(uintptr)
					} else {
						y = basic.Uintptr(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uint16)
						fr.setReg(ir, x<<y)
					}
				}
			} else {
				switch ykind {
				case reflect.Int:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(uint16)
							y := fr.reg(iy).(int)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uint16)
						y := basic.Int(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				case reflect.Int8:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(uint16)
							y := fr.reg(iy).(int8)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uint16)
						y := basic.Int8(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				case reflect.Int16:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(uint16)
							y := fr.reg(iy).(int16)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uint16)
						y := basic.Int16(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				case reflect.Int32:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(uint16)
							y := fr.reg(iy).(int32)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uint16)
						y := basic.Int32(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				case reflect.Int64:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(uint16)
							y := fr.reg(iy).(int64)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uint16)
						y := basic.Int64(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(uint16)
							y := fr.reg(iy).(uint)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uint16)
						y := basic.Uint(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint8:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(uint16)
							y := fr.reg(iy).(uint8)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uint16)
						y := basic.Uint8(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint16:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(uint16)
							y := fr.reg(iy).(uint16)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uint16)
						y := basic.Uint16(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint32:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(uint16)
							y := fr.reg(iy).(uint32)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uint16)
						y := basic.Uint32(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint64:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(uint16)
							y := fr.reg(iy).(uint64)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uint16)
						y := basic.Uint64(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				case reflect.Uintptr:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(uint16)
							y := fr.reg(iy).(uintptr)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uint16)
						y := basic.Uintptr(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				}
			}
		case reflect.Uint32:
			if kx == kindConst && ky == kindConst {
				x := vx.(uint32)
				switch ykind {
				case reflect.Int:
					var y int
					if yIsBasic {
						y = vy.(int)
					} else {
						y = basic.Int(vy)
					}
					v := x << y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Int8:
					var y int8
					if yIsBasic {
						y = vy.(int8)
					} else {
						y = basic.Int8(vy)
					}
					v := x << y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Int16:
					var y int16
					if yIsBasic {
						y = vy.(int16)
					} else {
						y = basic.Int16(vy)
					}
					v := x << y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Int32:
					var y int32
					if yIsBasic {
						y = vy.(int32)
					} else {
						y = basic.Int32(vy)
					}
					v := x << y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Int64:
					var y int64
					if yIsBasic {
						y = vy.(int64)
					} else {
						y = basic.Int64(vy)
					}
					v := x << y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uint:
					var y uint
					if yIsBasic {
						y = vy.(uint)
					} else {
						y = basic.Uint(vy)
					}
					v := x << y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uint8:
					var y uint8
					if yIsBasic {
						y = vy.(uint8)
					} else {
						y = basic.Uint8(vy)
					}
					v := x << y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uint16:
					var y uint16
					if yIsBasic {
						y = vy.(uint16)
					} else {
						y = basic.Uint16(vy)
					}
					v := x << y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uint32:
					var y uint32
					if yIsBasic {
						y = vy.(uint32)
					} else {
						y = basic.Uint32(vy)
					}
					v := x << y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uint64:
					var y uint64
					if yIsBasic {
						y = vy.(uint64)
					} else {
						y = basic.Uint64(vy)
					}
					v := x << y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uintptr:
					var y uintptr
					if yIsBasic {
						y = vy.(uintptr)
					} else {
						y = basic.Uintptr(vy)
					}
					v := x << y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				}
			} else if kx == kindConst {
				x := vx.(uint32)
				switch ykind {
				case reflect.Int:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(int)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						y := basic.Int(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				case reflect.Int8:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(int8)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						y := basic.Int8(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				case reflect.Int16:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(int16)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						y := basic.Int16(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				case reflect.Int32:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(int32)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						y := basic.Int32(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				case reflect.Int64:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(int64)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						y := basic.Int64(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uint)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						y := basic.Uint(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint8:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uint8)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						y := basic.Uint8(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint16:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uint16)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						y := basic.Uint16(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint32:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uint32)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						y := basic.Uint32(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint64:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uint64)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						y := basic.Uint64(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				case reflect.Uintptr:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uintptr)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						y := basic.Uintptr(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				}
			} else if ky == kindConst {
				switch ykind {
				case reflect.Int:
					var y int
					if yIsBasic {
						y = vy.(int)
					} else {
						y = basic.Int(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uint32)
						fr.setReg(ir, x<<y)
					}
				case reflect.Int8:
					var y int8
					if yIsBasic {
						y = vy.(int8)
					} else {
						y = basic.Int8(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uint32)
						fr.setReg(ir, x<<y)
					}
				case reflect.Int16:
					var y int16
					if yIsBasic {
						y = vy.(int16)
					} else {
						y = basic.Int16(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uint32)
						fr.setReg(ir, x<<y)
					}
				case reflect.Int32:
					var y int32
					if yIsBasic {
						y = vy.(int32)
					} else {
						y = basic.Int32(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uint32)
						fr.setReg(ir, x<<y)
					}
				case reflect.Int64:
					var y int64
					if yIsBasic {
						y = vy.(int64)
					} else {
						y = basic.Int64(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uint32)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint:
					var y uint
					if yIsBasic {
						y = vy.(uint)
					} else {
						y = basic.Uint(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uint32)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint8:
					var y uint8
					if yIsBasic {
						y = vy.(uint8)
					} else {
						y = basic.Uint8(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uint32)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint16:
					var y uint16
					if yIsBasic {
						y = vy.(uint16)
					} else {
						y = basic.Uint16(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uint32)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint32:
					var y uint32
					if yIsBasic {
						y = vy.(uint32)
					} else {
						y = basic.Uint32(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uint32)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint64:
					var y uint64
					if yIsBasic {
						y = vy.(uint64)
					} else {
						y = basic.Uint64(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uint32)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uintptr:
					var y uintptr
					if yIsBasic {
						y = vy.(uintptr)
					} else {
						y = basic.Uintptr(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uint32)
						fr.setReg(ir, x<<y)
					}
				}
			} else {
				switch ykind {
				case reflect.Int:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(uint32)
							y := fr.reg(iy).(int)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uint32)
						y := basic.Int(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				case reflect.Int8:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(uint32)
							y := fr.reg(iy).(int8)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uint32)
						y := basic.Int8(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				case reflect.Int16:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(uint32)
							y := fr.reg(iy).(int16)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uint32)
						y := basic.Int16(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				case reflect.Int32:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(uint32)
							y := fr.reg(iy).(int32)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uint32)
						y := basic.Int32(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				case reflect.Int64:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(uint32)
							y := fr.reg(iy).(int64)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uint32)
						y := basic.Int64(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(uint32)
							y := fr.reg(iy).(uint)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uint32)
						y := basic.Uint(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint8:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(uint32)
							y := fr.reg(iy).(uint8)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uint32)
						y := basic.Uint8(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint16:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(uint32)
							y := fr.reg(iy).(uint16)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uint32)
						y := basic.Uint16(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint32:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(uint32)
							y := fr.reg(iy).(uint32)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uint32)
						y := basic.Uint32(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint64:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(uint32)
							y := fr.reg(iy).(uint64)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uint32)
						y := basic.Uint64(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				case reflect.Uintptr:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(uint32)
							y := fr.reg(iy).(uintptr)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uint32)
						y := basic.Uintptr(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				}
			}
		case reflect.Uint64:
			if kx == kindConst && ky == kindConst {
				x := vx.(uint64)
				switch ykind {
				case reflect.Int:
					var y int
					if yIsBasic {
						y = vy.(int)
					} else {
						y = basic.Int(vy)
					}
					v := x << y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Int8:
					var y int8
					if yIsBasic {
						y = vy.(int8)
					} else {
						y = basic.Int8(vy)
					}
					v := x << y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Int16:
					var y int16
					if yIsBasic {
						y = vy.(int16)
					} else {
						y = basic.Int16(vy)
					}
					v := x << y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Int32:
					var y int32
					if yIsBasic {
						y = vy.(int32)
					} else {
						y = basic.Int32(vy)
					}
					v := x << y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Int64:
					var y int64
					if yIsBasic {
						y = vy.(int64)
					} else {
						y = basic.Int64(vy)
					}
					v := x << y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uint:
					var y uint
					if yIsBasic {
						y = vy.(uint)
					} else {
						y = basic.Uint(vy)
					}
					v := x << y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uint8:
					var y uint8
					if yIsBasic {
						y = vy.(uint8)
					} else {
						y = basic.Uint8(vy)
					}
					v := x << y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uint16:
					var y uint16
					if yIsBasic {
						y = vy.(uint16)
					} else {
						y = basic.Uint16(vy)
					}
					v := x << y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uint32:
					var y uint32
					if yIsBasic {
						y = vy.(uint32)
					} else {
						y = basic.Uint32(vy)
					}
					v := x << y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uint64:
					var y uint64
					if yIsBasic {
						y = vy.(uint64)
					} else {
						y = basic.Uint64(vy)
					}
					v := x << y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uintptr:
					var y uintptr
					if yIsBasic {
						y = vy.(uintptr)
					} else {
						y = basic.Uintptr(vy)
					}
					v := x << y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				}
			} else if kx == kindConst {
				x := vx.(uint64)
				switch ykind {
				case reflect.Int:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(int)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						y := basic.Int(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				case reflect.Int8:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(int8)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						y := basic.Int8(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				case reflect.Int16:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(int16)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						y := basic.Int16(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				case reflect.Int32:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(int32)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						y := basic.Int32(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				case reflect.Int64:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(int64)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						y := basic.Int64(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uint)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						y := basic.Uint(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint8:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uint8)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						y := basic.Uint8(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint16:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uint16)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						y := basic.Uint16(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint32:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uint32)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						y := basic.Uint32(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint64:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uint64)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						y := basic.Uint64(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				case reflect.Uintptr:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uintptr)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						y := basic.Uintptr(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				}
			} else if ky == kindConst {
				switch ykind {
				case reflect.Int:
					var y int
					if yIsBasic {
						y = vy.(int)
					} else {
						y = basic.Int(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uint64)
						fr.setReg(ir, x<<y)
					}
				case reflect.Int8:
					var y int8
					if yIsBasic {
						y = vy.(int8)
					} else {
						y = basic.Int8(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uint64)
						fr.setReg(ir, x<<y)
					}
				case reflect.Int16:
					var y int16
					if yIsBasic {
						y = vy.(int16)
					} else {
						y = basic.Int16(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uint64)
						fr.setReg(ir, x<<y)
					}
				case reflect.Int32:
					var y int32
					if yIsBasic {
						y = vy.(int32)
					} else {
						y = basic.Int32(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uint64)
						fr.setReg(ir, x<<y)
					}
				case reflect.Int64:
					var y int64
					if yIsBasic {
						y = vy.(int64)
					} else {
						y = basic.Int64(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uint64)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint:
					var y uint
					if yIsBasic {
						y = vy.(uint)
					} else {
						y = basic.Uint(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uint64)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint8:
					var y uint8
					if yIsBasic {
						y = vy.(uint8)
					} else {
						y = basic.Uint8(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uint64)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint16:
					var y uint16
					if yIsBasic {
						y = vy.(uint16)
					} else {
						y = basic.Uint16(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uint64)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint32:
					var y uint32
					if yIsBasic {
						y = vy.(uint32)
					} else {
						y = basic.Uint32(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uint64)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint64:
					var y uint64
					if yIsBasic {
						y = vy.(uint64)
					} else {
						y = basic.Uint64(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uint64)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uintptr:
					var y uintptr
					if yIsBasic {
						y = vy.(uintptr)
					} else {
						y = basic.Uintptr(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uint64)
						fr.setReg(ir, x<<y)
					}
				}
			} else {
				switch ykind {
				case reflect.Int:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(uint64)
							y := fr.reg(iy).(int)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uint64)
						y := basic.Int(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				case reflect.Int8:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(uint64)
							y := fr.reg(iy).(int8)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uint64)
						y := basic.Int8(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				case reflect.Int16:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(uint64)
							y := fr.reg(iy).(int16)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uint64)
						y := basic.Int16(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				case reflect.Int32:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(uint64)
							y := fr.reg(iy).(int32)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uint64)
						y := basic.Int32(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				case reflect.Int64:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(uint64)
							y := fr.reg(iy).(int64)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uint64)
						y := basic.Int64(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(uint64)
							y := fr.reg(iy).(uint)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uint64)
						y := basic.Uint(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint8:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(uint64)
							y := fr.reg(iy).(uint8)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uint64)
						y := basic.Uint8(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint16:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(uint64)
							y := fr.reg(iy).(uint16)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uint64)
						y := basic.Uint16(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint32:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(uint64)
							y := fr.reg(iy).(uint32)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uint64)
						y := basic.Uint32(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint64:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(uint64)
							y := fr.reg(iy).(uint64)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uint64)
						y := basic.Uint64(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				case reflect.Uintptr:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(uint64)
							y := fr.reg(iy).(uintptr)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uint64)
						y := basic.Uintptr(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				}
			}
		case reflect.Uintptr:
			if kx == kindConst && ky == kindConst {
				x := vx.(uintptr)
				switch ykind {
				case reflect.Int:
					var y int
					if yIsBasic {
						y = vy.(int)
					} else {
						y = basic.Int(vy)
					}
					v := x << y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Int8:
					var y int8
					if yIsBasic {
						y = vy.(int8)
					} else {
						y = basic.Int8(vy)
					}
					v := x << y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Int16:
					var y int16
					if yIsBasic {
						y = vy.(int16)
					} else {
						y = basic.Int16(vy)
					}
					v := x << y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Int32:
					var y int32
					if yIsBasic {
						y = vy.(int32)
					} else {
						y = basic.Int32(vy)
					}
					v := x << y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Int64:
					var y int64
					if yIsBasic {
						y = vy.(int64)
					} else {
						y = basic.Int64(vy)
					}
					v := x << y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uint:
					var y uint
					if yIsBasic {
						y = vy.(uint)
					} else {
						y = basic.Uint(vy)
					}
					v := x << y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uint8:
					var y uint8
					if yIsBasic {
						y = vy.(uint8)
					} else {
						y = basic.Uint8(vy)
					}
					v := x << y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uint16:
					var y uint16
					if yIsBasic {
						y = vy.(uint16)
					} else {
						y = basic.Uint16(vy)
					}
					v := x << y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uint32:
					var y uint32
					if yIsBasic {
						y = vy.(uint32)
					} else {
						y = basic.Uint32(vy)
					}
					v := x << y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uint64:
					var y uint64
					if yIsBasic {
						y = vy.(uint64)
					} else {
						y = basic.Uint64(vy)
					}
					v := x << y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uintptr:
					var y uintptr
					if yIsBasic {
						y = vy.(uintptr)
					} else {
						y = basic.Uintptr(vy)
					}
					v := x << y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				}
			} else if kx == kindConst {
				x := vx.(uintptr)
				switch ykind {
				case reflect.Int:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(int)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						y := basic.Int(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				case reflect.Int8:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(int8)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						y := basic.Int8(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				case reflect.Int16:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(int16)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						y := basic.Int16(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				case reflect.Int32:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(int32)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						y := basic.Int32(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				case reflect.Int64:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(int64)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						y := basic.Int64(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uint)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						y := basic.Uint(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint8:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uint8)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						y := basic.Uint8(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint16:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uint16)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						y := basic.Uint16(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint32:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uint32)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						y := basic.Uint32(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint64:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uint64)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						y := basic.Uint64(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				case reflect.Uintptr:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uintptr)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						y := basic.Uintptr(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				}
			} else if ky == kindConst {
				switch ykind {
				case reflect.Int:
					var y int
					if yIsBasic {
						y = vy.(int)
					} else {
						y = basic.Int(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uintptr)
						fr.setReg(ir, x<<y)
					}
				case reflect.Int8:
					var y int8
					if yIsBasic {
						y = vy.(int8)
					} else {
						y = basic.Int8(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uintptr)
						fr.setReg(ir, x<<y)
					}
				case reflect.Int16:
					var y int16
					if yIsBasic {
						y = vy.(int16)
					} else {
						y = basic.Int16(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uintptr)
						fr.setReg(ir, x<<y)
					}
				case reflect.Int32:
					var y int32
					if yIsBasic {
						y = vy.(int32)
					} else {
						y = basic.Int32(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uintptr)
						fr.setReg(ir, x<<y)
					}
				case reflect.Int64:
					var y int64
					if yIsBasic {
						y = vy.(int64)
					} else {
						y = basic.Int64(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uintptr)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint:
					var y uint
					if yIsBasic {
						y = vy.(uint)
					} else {
						y = basic.Uint(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uintptr)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint8:
					var y uint8
					if yIsBasic {
						y = vy.(uint8)
					} else {
						y = basic.Uint8(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uintptr)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint16:
					var y uint16
					if yIsBasic {
						y = vy.(uint16)
					} else {
						y = basic.Uint16(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uintptr)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint32:
					var y uint32
					if yIsBasic {
						y = vy.(uint32)
					} else {
						y = basic.Uint32(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uintptr)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint64:
					var y uint64
					if yIsBasic {
						y = vy.(uint64)
					} else {
						y = basic.Uint64(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uintptr)
						fr.setReg(ir, x<<y)
					}
				case reflect.Uintptr:
					var y uintptr
					if yIsBasic {
						y = vy.(uintptr)
					} else {
						y = basic.Uintptr(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uintptr)
						fr.setReg(ir, x<<y)
					}
				}
			} else {
				switch ykind {
				case reflect.Int:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(uintptr)
							y := fr.reg(iy).(int)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uintptr)
						y := basic.Int(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				case reflect.Int8:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(uintptr)
							y := fr.reg(iy).(int8)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uintptr)
						y := basic.Int8(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				case reflect.Int16:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(uintptr)
							y := fr.reg(iy).(int16)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uintptr)
						y := basic.Int16(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				case reflect.Int32:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(uintptr)
							y := fr.reg(iy).(int32)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uintptr)
						y := basic.Int32(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				case reflect.Int64:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(uintptr)
							y := fr.reg(iy).(int64)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uintptr)
						y := basic.Int64(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(uintptr)
							y := fr.reg(iy).(uint)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uintptr)
						y := basic.Uint(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint8:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(uintptr)
							y := fr.reg(iy).(uint8)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uintptr)
						y := basic.Uint8(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint16:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(uintptr)
							y := fr.reg(iy).(uint16)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uintptr)
						y := basic.Uint16(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint32:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(uintptr)
							y := fr.reg(iy).(uint32)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uintptr)
						y := basic.Uint32(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				case reflect.Uint64:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(uintptr)
							y := fr.reg(iy).(uint64)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uintptr)
						y := basic.Uint64(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				case reflect.Uintptr:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(uintptr)
							y := fr.reg(iy).(uintptr)
							fr.setReg(ir, x<<y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uintptr)
						y := basic.Uintptr(fr.reg(iy))
						fr.setReg(ir, x<<y)
					}
				}
			}
		}
	} else {
		t := basic.TypeOfType(xtyp)
		switch xkind {
		case reflect.Int:
			if kx == kindConst && ky == kindConst {
				x := basic.Int(vx)
				switch ykind {
				case reflect.Int:
					var y int
					if yIsBasic {
						y = vy.(int)
					} else {
						y = basic.Int(vy)
					}
					v := basic.Make(t, x<<y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Int8:
					var y int8
					if yIsBasic {
						y = vy.(int8)
					} else {
						y = basic.Int8(vy)
					}
					v := basic.Make(t, x<<y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Int16:
					var y int16
					if yIsBasic {
						y = vy.(int16)
					} else {
						y = basic.Int16(vy)
					}
					v := basic.Make(t, x<<y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Int32:
					var y int32
					if yIsBasic {
						y = vy.(int32)
					} else {
						y = basic.Int32(vy)
					}
					v := basic.Make(t, x<<y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Int64:
					var y int64
					if yIsBasic {
						y = vy.(int64)
					} else {
						y = basic.Int64(vy)
					}
					v := basic.Make(t, x<<y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uint:
					var y uint
					if yIsBasic {
						y = vy.(uint)
					} else {
						y = basic.Uint(vy)
					}
					v := basic.Make(t, x<<y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uint8:
					var y uint8
					if yIsBasic {
						y = vy.(uint8)
					} else {
						y = basic.Uint8(vy)
					}
					v := basic.Make(t, x<<y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uint16:
					var y uint16
					if yIsBasic {
						y = vy.(uint16)
					} else {
						y = basic.Uint16(vy)
					}
					v := basic.Make(t, x<<y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uint32:
					var y uint32
					if yIsBasic {
						y = vy.(uint32)
					} else {
						y = basic.Uint32(vy)
					}
					v := basic.Make(t, x<<y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uint64:
					var y uint64
					if yIsBasic {
						y = vy.(uint64)
					} else {
						y = basic.Uint64(vy)
					}
					v := basic.Make(t, x<<y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uintptr:
					var y uintptr
					if yIsBasic {
						y = vy.(uintptr)
					} else {
						y = basic.Uintptr(vy)
					}
					v := basic.Make(t, x<<y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				}
			} else if kx == kindConst {
				x := basic.Int(vx)
				switch ykind {
				case reflect.Int:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(int)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						y := basic.Int(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int8:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(int8)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						y := basic.Int8(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int16:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(int16)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						y := basic.Int16(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int32:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(int32)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						y := basic.Int32(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int64:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(int64)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						y := basic.Int64(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uint)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						y := basic.Uint(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint8:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uint8)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						y := basic.Uint8(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint16:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uint16)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						y := basic.Uint16(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint32:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uint32)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						y := basic.Uint32(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint64:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uint64)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						y := basic.Uint64(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uintptr:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uintptr)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						y := basic.Uintptr(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				}
			} else if ky == kindConst {
				switch ykind {
				case reflect.Int:
					var y int
					if yIsBasic {
						y = vy.(int)
					} else {
						y = basic.Int(vy)
					}
					return func(fr *frame) {
						x := basic.Int(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int8:
					var y int8
					if yIsBasic {
						y = vy.(int8)
					} else {
						y = basic.Int8(vy)
					}
					return func(fr *frame) {
						x := basic.Int(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int16:
					var y int16
					if yIsBasic {
						y = vy.(int16)
					} else {
						y = basic.Int16(vy)
					}
					return func(fr *frame) {
						x := basic.Int(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int32:
					var y int32
					if yIsBasic {
						y = vy.(int32)
					} else {
						y = basic.Int32(vy)
					}
					return func(fr *frame) {
						x := basic.Int(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int64:
					var y int64
					if yIsBasic {
						y = vy.(int64)
					} else {
						y = basic.Int64(vy)
					}
					return func(fr *frame) {
						x := basic.Int(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint:
					var y uint
					if yIsBasic {
						y = vy.(uint)
					} else {
						y = basic.Uint(vy)
					}
					return func(fr *frame) {
						x := basic.Int(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint8:
					var y uint8
					if yIsBasic {
						y = vy.(uint8)
					} else {
						y = basic.Uint8(vy)
					}
					return func(fr *frame) {
						x := basic.Int(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint16:
					var y uint16
					if yIsBasic {
						y = vy.(uint16)
					} else {
						y = basic.Uint16(vy)
					}
					return func(fr *frame) {
						x := basic.Int(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint32:
					var y uint32
					if yIsBasic {
						y = vy.(uint32)
					} else {
						y = basic.Uint32(vy)
					}
					return func(fr *frame) {
						x := basic.Int(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint64:
					var y uint64
					if yIsBasic {
						y = vy.(uint64)
					} else {
						y = basic.Uint64(vy)
					}
					return func(fr *frame) {
						x := basic.Int(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uintptr:
					var y uintptr
					if yIsBasic {
						y = vy.(uintptr)
					} else {
						y = basic.Uintptr(vy)
					}
					return func(fr *frame) {
						x := basic.Int(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				}
			} else {
				switch ykind {
				case reflect.Int:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Int(fr.reg(ix))
							y := fr.reg(iy).(int)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						x := basic.Int(fr.reg(ix))
						y := basic.Int(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int8:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Int(fr.reg(ix))
							y := fr.reg(iy).(int8)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						x := basic.Int(fr.reg(ix))
						y := basic.Int8(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int16:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Int(fr.reg(ix))
							y := fr.reg(iy).(int16)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						x := basic.Int(fr.reg(ix))
						y := basic.Int16(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int32:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Int(fr.reg(ix))
							y := fr.reg(iy).(int32)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						x := basic.Int(fr.reg(ix))
						y := basic.Int32(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int64:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Int(fr.reg(ix))
							y := fr.reg(iy).(int64)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						x := basic.Int(fr.reg(ix))
						y := basic.Int64(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Int(fr.reg(ix))
							y := fr.reg(iy).(uint)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						x := basic.Int(fr.reg(ix))
						y := basic.Uint(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint8:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Int(fr.reg(ix))
							y := fr.reg(iy).(uint8)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						x := basic.Int(fr.reg(ix))
						y := basic.Uint8(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint16:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Int(fr.reg(ix))
							y := fr.reg(iy).(uint16)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						x := basic.Int(fr.reg(ix))
						y := basic.Uint16(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint32:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Int(fr.reg(ix))
							y := fr.reg(iy).(uint32)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						x := basic.Int(fr.reg(ix))
						y := basic.Uint32(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint64:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Int(fr.reg(ix))
							y := fr.reg(iy).(uint64)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						x := basic.Int(fr.reg(ix))
						y := basic.Uint64(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uintptr:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Int(fr.reg(ix))
							y := fr.reg(iy).(uintptr)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						x := basic.Int(fr.reg(ix))
						y := basic.Uintptr(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				}
			}
		case reflect.Int8:
			if kx == kindConst && ky == kindConst {
				x := basic.Int8(vx)
				switch ykind {
				case reflect.Int:
					var y int
					if yIsBasic {
						y = vy.(int)
					} else {
						y = basic.Int(vy)
					}
					v := basic.Make(t, x<<y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Int8:
					var y int8
					if yIsBasic {
						y = vy.(int8)
					} else {
						y = basic.Int8(vy)
					}
					v := basic.Make(t, x<<y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Int16:
					var y int16
					if yIsBasic {
						y = vy.(int16)
					} else {
						y = basic.Int16(vy)
					}
					v := basic.Make(t, x<<y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Int32:
					var y int32
					if yIsBasic {
						y = vy.(int32)
					} else {
						y = basic.Int32(vy)
					}
					v := basic.Make(t, x<<y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Int64:
					var y int64
					if yIsBasic {
						y = vy.(int64)
					} else {
						y = basic.Int64(vy)
					}
					v := basic.Make(t, x<<y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uint:
					var y uint
					if yIsBasic {
						y = vy.(uint)
					} else {
						y = basic.Uint(vy)
					}
					v := basic.Make(t, x<<y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uint8:
					var y uint8
					if yIsBasic {
						y = vy.(uint8)
					} else {
						y = basic.Uint8(vy)
					}
					v := basic.Make(t, x<<y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uint16:
					var y uint16
					if yIsBasic {
						y = vy.(uint16)
					} else {
						y = basic.Uint16(vy)
					}
					v := basic.Make(t, x<<y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uint32:
					var y uint32
					if yIsBasic {
						y = vy.(uint32)
					} else {
						y = basic.Uint32(vy)
					}
					v := basic.Make(t, x<<y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uint64:
					var y uint64
					if yIsBasic {
						y = vy.(uint64)
					} else {
						y = basic.Uint64(vy)
					}
					v := basic.Make(t, x<<y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uintptr:
					var y uintptr
					if yIsBasic {
						y = vy.(uintptr)
					} else {
						y = basic.Uintptr(vy)
					}
					v := basic.Make(t, x<<y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				}
			} else if kx == kindConst {
				x := basic.Int8(vx)
				switch ykind {
				case reflect.Int:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(int)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						y := basic.Int(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int8:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(int8)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						y := basic.Int8(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int16:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(int16)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						y := basic.Int16(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int32:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(int32)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						y := basic.Int32(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int64:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(int64)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						y := basic.Int64(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uint)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						y := basic.Uint(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint8:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uint8)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						y := basic.Uint8(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint16:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uint16)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						y := basic.Uint16(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint32:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uint32)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						y := basic.Uint32(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint64:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uint64)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						y := basic.Uint64(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uintptr:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uintptr)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						y := basic.Uintptr(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				}
			} else if ky == kindConst {
				switch ykind {
				case reflect.Int:
					var y int
					if yIsBasic {
						y = vy.(int)
					} else {
						y = basic.Int(vy)
					}
					return func(fr *frame) {
						x := basic.Int8(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int8:
					var y int8
					if yIsBasic {
						y = vy.(int8)
					} else {
						y = basic.Int8(vy)
					}
					return func(fr *frame) {
						x := basic.Int8(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int16:
					var y int16
					if yIsBasic {
						y = vy.(int16)
					} else {
						y = basic.Int16(vy)
					}
					return func(fr *frame) {
						x := basic.Int8(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int32:
					var y int32
					if yIsBasic {
						y = vy.(int32)
					} else {
						y = basic.Int32(vy)
					}
					return func(fr *frame) {
						x := basic.Int8(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int64:
					var y int64
					if yIsBasic {
						y = vy.(int64)
					} else {
						y = basic.Int64(vy)
					}
					return func(fr *frame) {
						x := basic.Int8(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint:
					var y uint
					if yIsBasic {
						y = vy.(uint)
					} else {
						y = basic.Uint(vy)
					}
					return func(fr *frame) {
						x := basic.Int8(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint8:
					var y uint8
					if yIsBasic {
						y = vy.(uint8)
					} else {
						y = basic.Uint8(vy)
					}
					return func(fr *frame) {
						x := basic.Int8(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint16:
					var y uint16
					if yIsBasic {
						y = vy.(uint16)
					} else {
						y = basic.Uint16(vy)
					}
					return func(fr *frame) {
						x := basic.Int8(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint32:
					var y uint32
					if yIsBasic {
						y = vy.(uint32)
					} else {
						y = basic.Uint32(vy)
					}
					return func(fr *frame) {
						x := basic.Int8(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint64:
					var y uint64
					if yIsBasic {
						y = vy.(uint64)
					} else {
						y = basic.Uint64(vy)
					}
					return func(fr *frame) {
						x := basic.Int8(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uintptr:
					var y uintptr
					if yIsBasic {
						y = vy.(uintptr)
					} else {
						y = basic.Uintptr(vy)
					}
					return func(fr *frame) {
						x := basic.Int8(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				}
			} else {
				switch ykind {
				case reflect.Int:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Int8(fr.reg(ix))
							y := fr.reg(iy).(int)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						x := basic.Int8(fr.reg(ix))
						y := basic.Int(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int8:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Int8(fr.reg(ix))
							y := fr.reg(iy).(int8)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						x := basic.Int8(fr.reg(ix))
						y := basic.Int8(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int16:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Int8(fr.reg(ix))
							y := fr.reg(iy).(int16)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						x := basic.Int8(fr.reg(ix))
						y := basic.Int16(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int32:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Int8(fr.reg(ix))
							y := fr.reg(iy).(int32)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						x := basic.Int8(fr.reg(ix))
						y := basic.Int32(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int64:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Int8(fr.reg(ix))
							y := fr.reg(iy).(int64)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						x := basic.Int8(fr.reg(ix))
						y := basic.Int64(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Int8(fr.reg(ix))
							y := fr.reg(iy).(uint)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						x := basic.Int8(fr.reg(ix))
						y := basic.Uint(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint8:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Int8(fr.reg(ix))
							y := fr.reg(iy).(uint8)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						x := basic.Int8(fr.reg(ix))
						y := basic.Uint8(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint16:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Int8(fr.reg(ix))
							y := fr.reg(iy).(uint16)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						x := basic.Int8(fr.reg(ix))
						y := basic.Uint16(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint32:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Int8(fr.reg(ix))
							y := fr.reg(iy).(uint32)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						x := basic.Int8(fr.reg(ix))
						y := basic.Uint32(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint64:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Int8(fr.reg(ix))
							y := fr.reg(iy).(uint64)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						x := basic.Int8(fr.reg(ix))
						y := basic.Uint64(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uintptr:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Int8(fr.reg(ix))
							y := fr.reg(iy).(uintptr)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						x := basic.Int8(fr.reg(ix))
						y := basic.Uintptr(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				}
			}
		case reflect.Int16:
			if kx == kindConst && ky == kindConst {
				x := basic.Int16(vx)
				switch ykind {
				case reflect.Int:
					var y int
					if yIsBasic {
						y = vy.(int)
					} else {
						y = basic.Int(vy)
					}
					v := basic.Make(t, x<<y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Int8:
					var y int8
					if yIsBasic {
						y = vy.(int8)
					} else {
						y = basic.Int8(vy)
					}
					v := basic.Make(t, x<<y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Int16:
					var y int16
					if yIsBasic {
						y = vy.(int16)
					} else {
						y = basic.Int16(vy)
					}
					v := basic.Make(t, x<<y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Int32:
					var y int32
					if yIsBasic {
						y = vy.(int32)
					} else {
						y = basic.Int32(vy)
					}
					v := basic.Make(t, x<<y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Int64:
					var y int64
					if yIsBasic {
						y = vy.(int64)
					} else {
						y = basic.Int64(vy)
					}
					v := basic.Make(t, x<<y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uint:
					var y uint
					if yIsBasic {
						y = vy.(uint)
					} else {
						y = basic.Uint(vy)
					}
					v := basic.Make(t, x<<y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uint8:
					var y uint8
					if yIsBasic {
						y = vy.(uint8)
					} else {
						y = basic.Uint8(vy)
					}
					v := basic.Make(t, x<<y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uint16:
					var y uint16
					if yIsBasic {
						y = vy.(uint16)
					} else {
						y = basic.Uint16(vy)
					}
					v := basic.Make(t, x<<y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uint32:
					var y uint32
					if yIsBasic {
						y = vy.(uint32)
					} else {
						y = basic.Uint32(vy)
					}
					v := basic.Make(t, x<<y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uint64:
					var y uint64
					if yIsBasic {
						y = vy.(uint64)
					} else {
						y = basic.Uint64(vy)
					}
					v := basic.Make(t, x<<y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uintptr:
					var y uintptr
					if yIsBasic {
						y = vy.(uintptr)
					} else {
						y = basic.Uintptr(vy)
					}
					v := basic.Make(t, x<<y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				}
			} else if kx == kindConst {
				x := basic.Int16(vx)
				switch ykind {
				case reflect.Int:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(int)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						y := basic.Int(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int8:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(int8)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						y := basic.Int8(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int16:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(int16)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						y := basic.Int16(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int32:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(int32)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						y := basic.Int32(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int64:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(int64)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						y := basic.Int64(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uint)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						y := basic.Uint(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint8:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uint8)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						y := basic.Uint8(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint16:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uint16)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						y := basic.Uint16(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint32:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uint32)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						y := basic.Uint32(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint64:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uint64)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						y := basic.Uint64(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uintptr:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uintptr)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						y := basic.Uintptr(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				}
			} else if ky == kindConst {
				switch ykind {
				case reflect.Int:
					var y int
					if yIsBasic {
						y = vy.(int)
					} else {
						y = basic.Int(vy)
					}
					return func(fr *frame) {
						x := basic.Int16(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int8:
					var y int8
					if yIsBasic {
						y = vy.(int8)
					} else {
						y = basic.Int8(vy)
					}
					return func(fr *frame) {
						x := basic.Int16(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int16:
					var y int16
					if yIsBasic {
						y = vy.(int16)
					} else {
						y = basic.Int16(vy)
					}
					return func(fr *frame) {
						x := basic.Int16(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int32:
					var y int32
					if yIsBasic {
						y = vy.(int32)
					} else {
						y = basic.Int32(vy)
					}
					return func(fr *frame) {
						x := basic.Int16(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int64:
					var y int64
					if yIsBasic {
						y = vy.(int64)
					} else {
						y = basic.Int64(vy)
					}
					return func(fr *frame) {
						x := basic.Int16(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint:
					var y uint
					if yIsBasic {
						y = vy.(uint)
					} else {
						y = basic.Uint(vy)
					}
					return func(fr *frame) {
						x := basic.Int16(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint8:
					var y uint8
					if yIsBasic {
						y = vy.(uint8)
					} else {
						y = basic.Uint8(vy)
					}
					return func(fr *frame) {
						x := basic.Int16(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint16:
					var y uint16
					if yIsBasic {
						y = vy.(uint16)
					} else {
						y = basic.Uint16(vy)
					}
					return func(fr *frame) {
						x := basic.Int16(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint32:
					var y uint32
					if yIsBasic {
						y = vy.(uint32)
					} else {
						y = basic.Uint32(vy)
					}
					return func(fr *frame) {
						x := basic.Int16(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint64:
					var y uint64
					if yIsBasic {
						y = vy.(uint64)
					} else {
						y = basic.Uint64(vy)
					}
					return func(fr *frame) {
						x := basic.Int16(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uintptr:
					var y uintptr
					if yIsBasic {
						y = vy.(uintptr)
					} else {
						y = basic.Uintptr(vy)
					}
					return func(fr *frame) {
						x := basic.Int16(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				}
			} else {
				switch ykind {
				case reflect.Int:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Int16(fr.reg(ix))
							y := fr.reg(iy).(int)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						x := basic.Int16(fr.reg(ix))
						y := basic.Int(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int8:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Int16(fr.reg(ix))
							y := fr.reg(iy).(int8)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						x := basic.Int16(fr.reg(ix))
						y := basic.Int8(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int16:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Int16(fr.reg(ix))
							y := fr.reg(iy).(int16)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						x := basic.Int16(fr.reg(ix))
						y := basic.Int16(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int32:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Int16(fr.reg(ix))
							y := fr.reg(iy).(int32)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						x := basic.Int16(fr.reg(ix))
						y := basic.Int32(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int64:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Int16(fr.reg(ix))
							y := fr.reg(iy).(int64)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						x := basic.Int16(fr.reg(ix))
						y := basic.Int64(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Int16(fr.reg(ix))
							y := fr.reg(iy).(uint)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						x := basic.Int16(fr.reg(ix))
						y := basic.Uint(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint8:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Int16(fr.reg(ix))
							y := fr.reg(iy).(uint8)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						x := basic.Int16(fr.reg(ix))
						y := basic.Uint8(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint16:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Int16(fr.reg(ix))
							y := fr.reg(iy).(uint16)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						x := basic.Int16(fr.reg(ix))
						y := basic.Uint16(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint32:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Int16(fr.reg(ix))
							y := fr.reg(iy).(uint32)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						x := basic.Int16(fr.reg(ix))
						y := basic.Uint32(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint64:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Int16(fr.reg(ix))
							y := fr.reg(iy).(uint64)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						x := basic.Int16(fr.reg(ix))
						y := basic.Uint64(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uintptr:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Int16(fr.reg(ix))
							y := fr.reg(iy).(uintptr)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						x := basic.Int16(fr.reg(ix))
						y := basic.Uintptr(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				}
			}
		case reflect.Int32:
			if kx == kindConst && ky == kindConst {
				x := basic.Int32(vx)
				switch ykind {
				case reflect.Int:
					var y int
					if yIsBasic {
						y = vy.(int)
					} else {
						y = basic.Int(vy)
					}
					v := basic.Make(t, x<<y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Int8:
					var y int8
					if yIsBasic {
						y = vy.(int8)
					} else {
						y = basic.Int8(vy)
					}
					v := basic.Make(t, x<<y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Int16:
					var y int16
					if yIsBasic {
						y = vy.(int16)
					} else {
						y = basic.Int16(vy)
					}
					v := basic.Make(t, x<<y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Int32:
					var y int32
					if yIsBasic {
						y = vy.(int32)
					} else {
						y = basic.Int32(vy)
					}
					v := basic.Make(t, x<<y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Int64:
					var y int64
					if yIsBasic {
						y = vy.(int64)
					} else {
						y = basic.Int64(vy)
					}
					v := basic.Make(t, x<<y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uint:
					var y uint
					if yIsBasic {
						y = vy.(uint)
					} else {
						y = basic.Uint(vy)
					}
					v := basic.Make(t, x<<y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uint8:
					var y uint8
					if yIsBasic {
						y = vy.(uint8)
					} else {
						y = basic.Uint8(vy)
					}
					v := basic.Make(t, x<<y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uint16:
					var y uint16
					if yIsBasic {
						y = vy.(uint16)
					} else {
						y = basic.Uint16(vy)
					}
					v := basic.Make(t, x<<y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uint32:
					var y uint32
					if yIsBasic {
						y = vy.(uint32)
					} else {
						y = basic.Uint32(vy)
					}
					v := basic.Make(t, x<<y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uint64:
					var y uint64
					if yIsBasic {
						y = vy.(uint64)
					} else {
						y = basic.Uint64(vy)
					}
					v := basic.Make(t, x<<y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uintptr:
					var y uintptr
					if yIsBasic {
						y = vy.(uintptr)
					} else {
						y = basic.Uintptr(vy)
					}
					v := basic.Make(t, x<<y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				}
			} else if kx == kindConst {
				x := basic.Int32(vx)
				switch ykind {
				case reflect.Int:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(int)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						y := basic.Int(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int8:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(int8)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						y := basic.Int8(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int16:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(int16)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						y := basic.Int16(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int32:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(int32)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						y := basic.Int32(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int64:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(int64)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						y := basic.Int64(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uint)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						y := basic.Uint(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint8:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uint8)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						y := basic.Uint8(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint16:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uint16)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						y := basic.Uint16(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint32:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uint32)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						y := basic.Uint32(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint64:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uint64)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						y := basic.Uint64(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uintptr:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uintptr)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						y := basic.Uintptr(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				}
			} else if ky == kindConst {
				switch ykind {
				case reflect.Int:
					var y int
					if yIsBasic {
						y = vy.(int)
					} else {
						y = basic.Int(vy)
					}
					return func(fr *frame) {
						x := basic.Int32(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int8:
					var y int8
					if yIsBasic {
						y = vy.(int8)
					} else {
						y = basic.Int8(vy)
					}
					return func(fr *frame) {
						x := basic.Int32(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int16:
					var y int16
					if yIsBasic {
						y = vy.(int16)
					} else {
						y = basic.Int16(vy)
					}
					return func(fr *frame) {
						x := basic.Int32(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int32:
					var y int32
					if yIsBasic {
						y = vy.(int32)
					} else {
						y = basic.Int32(vy)
					}
					return func(fr *frame) {
						x := basic.Int32(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int64:
					var y int64
					if yIsBasic {
						y = vy.(int64)
					} else {
						y = basic.Int64(vy)
					}
					return func(fr *frame) {
						x := basic.Int32(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint:
					var y uint
					if yIsBasic {
						y = vy.(uint)
					} else {
						y = basic.Uint(vy)
					}
					return func(fr *frame) {
						x := basic.Int32(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint8:
					var y uint8
					if yIsBasic {
						y = vy.(uint8)
					} else {
						y = basic.Uint8(vy)
					}
					return func(fr *frame) {
						x := basic.Int32(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint16:
					var y uint16
					if yIsBasic {
						y = vy.(uint16)
					} else {
						y = basic.Uint16(vy)
					}
					return func(fr *frame) {
						x := basic.Int32(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint32:
					var y uint32
					if yIsBasic {
						y = vy.(uint32)
					} else {
						y = basic.Uint32(vy)
					}
					return func(fr *frame) {
						x := basic.Int32(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint64:
					var y uint64
					if yIsBasic {
						y = vy.(uint64)
					} else {
						y = basic.Uint64(vy)
					}
					return func(fr *frame) {
						x := basic.Int32(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uintptr:
					var y uintptr
					if yIsBasic {
						y = vy.(uintptr)
					} else {
						y = basic.Uintptr(vy)
					}
					return func(fr *frame) {
						x := basic.Int32(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				}
			} else {
				switch ykind {
				case reflect.Int:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Int32(fr.reg(ix))
							y := fr.reg(iy).(int)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						x := basic.Int32(fr.reg(ix))
						y := basic.Int(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int8:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Int32(fr.reg(ix))
							y := fr.reg(iy).(int8)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						x := basic.Int32(fr.reg(ix))
						y := basic.Int8(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int16:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Int32(fr.reg(ix))
							y := fr.reg(iy).(int16)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						x := basic.Int32(fr.reg(ix))
						y := basic.Int16(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int32:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Int32(fr.reg(ix))
							y := fr.reg(iy).(int32)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						x := basic.Int32(fr.reg(ix))
						y := basic.Int32(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int64:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Int32(fr.reg(ix))
							y := fr.reg(iy).(int64)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						x := basic.Int32(fr.reg(ix))
						y := basic.Int64(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Int32(fr.reg(ix))
							y := fr.reg(iy).(uint)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						x := basic.Int32(fr.reg(ix))
						y := basic.Uint(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint8:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Int32(fr.reg(ix))
							y := fr.reg(iy).(uint8)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						x := basic.Int32(fr.reg(ix))
						y := basic.Uint8(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint16:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Int32(fr.reg(ix))
							y := fr.reg(iy).(uint16)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						x := basic.Int32(fr.reg(ix))
						y := basic.Uint16(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint32:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Int32(fr.reg(ix))
							y := fr.reg(iy).(uint32)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						x := basic.Int32(fr.reg(ix))
						y := basic.Uint32(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint64:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Int32(fr.reg(ix))
							y := fr.reg(iy).(uint64)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						x := basic.Int32(fr.reg(ix))
						y := basic.Uint64(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uintptr:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Int32(fr.reg(ix))
							y := fr.reg(iy).(uintptr)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						x := basic.Int32(fr.reg(ix))
						y := basic.Uintptr(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				}
			}
		case reflect.Int64:
			if kx == kindConst && ky == kindConst {
				x := basic.Int64(vx)
				switch ykind {
				case reflect.Int:
					var y int
					if yIsBasic {
						y = vy.(int)
					} else {
						y = basic.Int(vy)
					}
					v := basic.Make(t, x<<y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Int8:
					var y int8
					if yIsBasic {
						y = vy.(int8)
					} else {
						y = basic.Int8(vy)
					}
					v := basic.Make(t, x<<y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Int16:
					var y int16
					if yIsBasic {
						y = vy.(int16)
					} else {
						y = basic.Int16(vy)
					}
					v := basic.Make(t, x<<y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Int32:
					var y int32
					if yIsBasic {
						y = vy.(int32)
					} else {
						y = basic.Int32(vy)
					}
					v := basic.Make(t, x<<y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Int64:
					var y int64
					if yIsBasic {
						y = vy.(int64)
					} else {
						y = basic.Int64(vy)
					}
					v := basic.Make(t, x<<y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uint:
					var y uint
					if yIsBasic {
						y = vy.(uint)
					} else {
						y = basic.Uint(vy)
					}
					v := basic.Make(t, x<<y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uint8:
					var y uint8
					if yIsBasic {
						y = vy.(uint8)
					} else {
						y = basic.Uint8(vy)
					}
					v := basic.Make(t, x<<y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uint16:
					var y uint16
					if yIsBasic {
						y = vy.(uint16)
					} else {
						y = basic.Uint16(vy)
					}
					v := basic.Make(t, x<<y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uint32:
					var y uint32
					if yIsBasic {
						y = vy.(uint32)
					} else {
						y = basic.Uint32(vy)
					}
					v := basic.Make(t, x<<y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uint64:
					var y uint64
					if yIsBasic {
						y = vy.(uint64)
					} else {
						y = basic.Uint64(vy)
					}
					v := basic.Make(t, x<<y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uintptr:
					var y uintptr
					if yIsBasic {
						y = vy.(uintptr)
					} else {
						y = basic.Uintptr(vy)
					}
					v := basic.Make(t, x<<y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				}
			} else if kx == kindConst {
				x := basic.Int64(vx)
				switch ykind {
				case reflect.Int:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(int)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						y := basic.Int(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int8:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(int8)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						y := basic.Int8(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int16:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(int16)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						y := basic.Int16(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int32:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(int32)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						y := basic.Int32(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int64:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(int64)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						y := basic.Int64(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uint)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						y := basic.Uint(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint8:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uint8)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						y := basic.Uint8(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint16:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uint16)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						y := basic.Uint16(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint32:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uint32)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						y := basic.Uint32(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint64:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uint64)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						y := basic.Uint64(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uintptr:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uintptr)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						y := basic.Uintptr(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				}
			} else if ky == kindConst {
				switch ykind {
				case reflect.Int:
					var y int
					if yIsBasic {
						y = vy.(int)
					} else {
						y = basic.Int(vy)
					}
					return func(fr *frame) {
						x := basic.Int64(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int8:
					var y int8
					if yIsBasic {
						y = vy.(int8)
					} else {
						y = basic.Int8(vy)
					}
					return func(fr *frame) {
						x := basic.Int64(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int16:
					var y int16
					if yIsBasic {
						y = vy.(int16)
					} else {
						y = basic.Int16(vy)
					}
					return func(fr *frame) {
						x := basic.Int64(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int32:
					var y int32
					if yIsBasic {
						y = vy.(int32)
					} else {
						y = basic.Int32(vy)
					}
					return func(fr *frame) {
						x := basic.Int64(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int64:
					var y int64
					if yIsBasic {
						y = vy.(int64)
					} else {
						y = basic.Int64(vy)
					}
					return func(fr *frame) {
						x := basic.Int64(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint:
					var y uint
					if yIsBasic {
						y = vy.(uint)
					} else {
						y = basic.Uint(vy)
					}
					return func(fr *frame) {
						x := basic.Int64(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint8:
					var y uint8
					if yIsBasic {
						y = vy.(uint8)
					} else {
						y = basic.Uint8(vy)
					}
					return func(fr *frame) {
						x := basic.Int64(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint16:
					var y uint16
					if yIsBasic {
						y = vy.(uint16)
					} else {
						y = basic.Uint16(vy)
					}
					return func(fr *frame) {
						x := basic.Int64(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint32:
					var y uint32
					if yIsBasic {
						y = vy.(uint32)
					} else {
						y = basic.Uint32(vy)
					}
					return func(fr *frame) {
						x := basic.Int64(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint64:
					var y uint64
					if yIsBasic {
						y = vy.(uint64)
					} else {
						y = basic.Uint64(vy)
					}
					return func(fr *frame) {
						x := basic.Int64(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uintptr:
					var y uintptr
					if yIsBasic {
						y = vy.(uintptr)
					} else {
						y = basic.Uintptr(vy)
					}
					return func(fr *frame) {
						x := basic.Int64(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				}
			} else {
				switch ykind {
				case reflect.Int:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Int64(fr.reg(ix))
							y := fr.reg(iy).(int)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						x := basic.Int64(fr.reg(ix))
						y := basic.Int(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int8:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Int64(fr.reg(ix))
							y := fr.reg(iy).(int8)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						x := basic.Int64(fr.reg(ix))
						y := basic.Int8(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int16:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Int64(fr.reg(ix))
							y := fr.reg(iy).(int16)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						x := basic.Int64(fr.reg(ix))
						y := basic.Int16(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int32:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Int64(fr.reg(ix))
							y := fr.reg(iy).(int32)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						x := basic.Int64(fr.reg(ix))
						y := basic.Int32(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int64:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Int64(fr.reg(ix))
							y := fr.reg(iy).(int64)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						x := basic.Int64(fr.reg(ix))
						y := basic.Int64(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Int64(fr.reg(ix))
							y := fr.reg(iy).(uint)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						x := basic.Int64(fr.reg(ix))
						y := basic.Uint(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint8:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Int64(fr.reg(ix))
							y := fr.reg(iy).(uint8)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						x := basic.Int64(fr.reg(ix))
						y := basic.Uint8(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint16:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Int64(fr.reg(ix))
							y := fr.reg(iy).(uint16)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						x := basic.Int64(fr.reg(ix))
						y := basic.Uint16(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint32:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Int64(fr.reg(ix))
							y := fr.reg(iy).(uint32)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						x := basic.Int64(fr.reg(ix))
						y := basic.Uint32(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint64:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Int64(fr.reg(ix))
							y := fr.reg(iy).(uint64)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						x := basic.Int64(fr.reg(ix))
						y := basic.Uint64(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uintptr:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Int64(fr.reg(ix))
							y := fr.reg(iy).(uintptr)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						x := basic.Int64(fr.reg(ix))
						y := basic.Uintptr(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				}
			}
		case reflect.Uint:
			if kx == kindConst && ky == kindConst {
				x := basic.Uint(vx)
				switch ykind {
				case reflect.Int:
					var y int
					if yIsBasic {
						y = vy.(int)
					} else {
						y = basic.Int(vy)
					}
					v := basic.Make(t, x<<y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Int8:
					var y int8
					if yIsBasic {
						y = vy.(int8)
					} else {
						y = basic.Int8(vy)
					}
					v := basic.Make(t, x<<y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Int16:
					var y int16
					if yIsBasic {
						y = vy.(int16)
					} else {
						y = basic.Int16(vy)
					}
					v := basic.Make(t, x<<y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Int32:
					var y int32
					if yIsBasic {
						y = vy.(int32)
					} else {
						y = basic.Int32(vy)
					}
					v := basic.Make(t, x<<y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Int64:
					var y int64
					if yIsBasic {
						y = vy.(int64)
					} else {
						y = basic.Int64(vy)
					}
					v := basic.Make(t, x<<y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uint:
					var y uint
					if yIsBasic {
						y = vy.(uint)
					} else {
						y = basic.Uint(vy)
					}
					v := basic.Make(t, x<<y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uint8:
					var y uint8
					if yIsBasic {
						y = vy.(uint8)
					} else {
						y = basic.Uint8(vy)
					}
					v := basic.Make(t, x<<y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uint16:
					var y uint16
					if yIsBasic {
						y = vy.(uint16)
					} else {
						y = basic.Uint16(vy)
					}
					v := basic.Make(t, x<<y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uint32:
					var y uint32
					if yIsBasic {
						y = vy.(uint32)
					} else {
						y = basic.Uint32(vy)
					}
					v := basic.Make(t, x<<y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uint64:
					var y uint64
					if yIsBasic {
						y = vy.(uint64)
					} else {
						y = basic.Uint64(vy)
					}
					v := basic.Make(t, x<<y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uintptr:
					var y uintptr
					if yIsBasic {
						y = vy.(uintptr)
					} else {
						y = basic.Uintptr(vy)
					}
					v := basic.Make(t, x<<y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				}
			} else if kx == kindConst {
				x := basic.Uint(vx)
				switch ykind {
				case reflect.Int:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(int)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						y := basic.Int(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int8:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(int8)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						y := basic.Int8(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int16:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(int16)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						y := basic.Int16(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int32:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(int32)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						y := basic.Int32(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int64:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(int64)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						y := basic.Int64(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uint)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						y := basic.Uint(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint8:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uint8)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						y := basic.Uint8(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint16:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uint16)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						y := basic.Uint16(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint32:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uint32)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						y := basic.Uint32(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint64:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uint64)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						y := basic.Uint64(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uintptr:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uintptr)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						y := basic.Uintptr(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				}
			} else if ky == kindConst {
				switch ykind {
				case reflect.Int:
					var y int
					if yIsBasic {
						y = vy.(int)
					} else {
						y = basic.Int(vy)
					}
					return func(fr *frame) {
						x := basic.Uint(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int8:
					var y int8
					if yIsBasic {
						y = vy.(int8)
					} else {
						y = basic.Int8(vy)
					}
					return func(fr *frame) {
						x := basic.Uint(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int16:
					var y int16
					if yIsBasic {
						y = vy.(int16)
					} else {
						y = basic.Int16(vy)
					}
					return func(fr *frame) {
						x := basic.Uint(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int32:
					var y int32
					if yIsBasic {
						y = vy.(int32)
					} else {
						y = basic.Int32(vy)
					}
					return func(fr *frame) {
						x := basic.Uint(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int64:
					var y int64
					if yIsBasic {
						y = vy.(int64)
					} else {
						y = basic.Int64(vy)
					}
					return func(fr *frame) {
						x := basic.Uint(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint:
					var y uint
					if yIsBasic {
						y = vy.(uint)
					} else {
						y = basic.Uint(vy)
					}
					return func(fr *frame) {
						x := basic.Uint(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint8:
					var y uint8
					if yIsBasic {
						y = vy.(uint8)
					} else {
						y = basic.Uint8(vy)
					}
					return func(fr *frame) {
						x := basic.Uint(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint16:
					var y uint16
					if yIsBasic {
						y = vy.(uint16)
					} else {
						y = basic.Uint16(vy)
					}
					return func(fr *frame) {
						x := basic.Uint(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint32:
					var y uint32
					if yIsBasic {
						y = vy.(uint32)
					} else {
						y = basic.Uint32(vy)
					}
					return func(fr *frame) {
						x := basic.Uint(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint64:
					var y uint64
					if yIsBasic {
						y = vy.(uint64)
					} else {
						y = basic.Uint64(vy)
					}
					return func(fr *frame) {
						x := basic.Uint(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uintptr:
					var y uintptr
					if yIsBasic {
						y = vy.(uintptr)
					} else {
						y = basic.Uintptr(vy)
					}
					return func(fr *frame) {
						x := basic.Uint(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				}
			} else {
				switch ykind {
				case reflect.Int:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Uint(fr.reg(ix))
							y := fr.reg(iy).(int)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						x := basic.Uint(fr.reg(ix))
						y := basic.Int(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int8:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Uint(fr.reg(ix))
							y := fr.reg(iy).(int8)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						x := basic.Uint(fr.reg(ix))
						y := basic.Int8(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int16:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Uint(fr.reg(ix))
							y := fr.reg(iy).(int16)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						x := basic.Uint(fr.reg(ix))
						y := basic.Int16(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int32:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Uint(fr.reg(ix))
							y := fr.reg(iy).(int32)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						x := basic.Uint(fr.reg(ix))
						y := basic.Int32(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int64:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Uint(fr.reg(ix))
							y := fr.reg(iy).(int64)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						x := basic.Uint(fr.reg(ix))
						y := basic.Int64(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Uint(fr.reg(ix))
							y := fr.reg(iy).(uint)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						x := basic.Uint(fr.reg(ix))
						y := basic.Uint(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint8:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Uint(fr.reg(ix))
							y := fr.reg(iy).(uint8)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						x := basic.Uint(fr.reg(ix))
						y := basic.Uint8(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint16:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Uint(fr.reg(ix))
							y := fr.reg(iy).(uint16)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						x := basic.Uint(fr.reg(ix))
						y := basic.Uint16(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint32:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Uint(fr.reg(ix))
							y := fr.reg(iy).(uint32)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						x := basic.Uint(fr.reg(ix))
						y := basic.Uint32(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint64:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Uint(fr.reg(ix))
							y := fr.reg(iy).(uint64)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						x := basic.Uint(fr.reg(ix))
						y := basic.Uint64(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uintptr:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Uint(fr.reg(ix))
							y := fr.reg(iy).(uintptr)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						x := basic.Uint(fr.reg(ix))
						y := basic.Uintptr(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				}
			}
		case reflect.Uint8:
			if kx == kindConst && ky == kindConst {
				x := basic.Uint8(vx)
				switch ykind {
				case reflect.Int:
					var y int
					if yIsBasic {
						y = vy.(int)
					} else {
						y = basic.Int(vy)
					}
					v := basic.Make(t, x<<y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Int8:
					var y int8
					if yIsBasic {
						y = vy.(int8)
					} else {
						y = basic.Int8(vy)
					}
					v := basic.Make(t, x<<y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Int16:
					var y int16
					if yIsBasic {
						y = vy.(int16)
					} else {
						y = basic.Int16(vy)
					}
					v := basic.Make(t, x<<y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Int32:
					var y int32
					if yIsBasic {
						y = vy.(int32)
					} else {
						y = basic.Int32(vy)
					}
					v := basic.Make(t, x<<y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Int64:
					var y int64
					if yIsBasic {
						y = vy.(int64)
					} else {
						y = basic.Int64(vy)
					}
					v := basic.Make(t, x<<y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uint:
					var y uint
					if yIsBasic {
						y = vy.(uint)
					} else {
						y = basic.Uint(vy)
					}
					v := basic.Make(t, x<<y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uint8:
					var y uint8
					if yIsBasic {
						y = vy.(uint8)
					} else {
						y = basic.Uint8(vy)
					}
					v := basic.Make(t, x<<y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uint16:
					var y uint16
					if yIsBasic {
						y = vy.(uint16)
					} else {
						y = basic.Uint16(vy)
					}
					v := basic.Make(t, x<<y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uint32:
					var y uint32
					if yIsBasic {
						y = vy.(uint32)
					} else {
						y = basic.Uint32(vy)
					}
					v := basic.Make(t, x<<y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uint64:
					var y uint64
					if yIsBasic {
						y = vy.(uint64)
					} else {
						y = basic.Uint64(vy)
					}
					v := basic.Make(t, x<<y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uintptr:
					var y uintptr
					if yIsBasic {
						y = vy.(uintptr)
					} else {
						y = basic.Uintptr(vy)
					}
					v := basic.Make(t, x<<y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				}
			} else if kx == kindConst {
				x := basic.Uint8(vx)
				switch ykind {
				case reflect.Int:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(int)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						y := basic.Int(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int8:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(int8)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						y := basic.Int8(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int16:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(int16)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						y := basic.Int16(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int32:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(int32)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						y := basic.Int32(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int64:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(int64)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						y := basic.Int64(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uint)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						y := basic.Uint(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint8:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uint8)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						y := basic.Uint8(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint16:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uint16)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						y := basic.Uint16(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint32:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uint32)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						y := basic.Uint32(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint64:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uint64)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						y := basic.Uint64(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uintptr:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uintptr)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						y := basic.Uintptr(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				}
			} else if ky == kindConst {
				switch ykind {
				case reflect.Int:
					var y int
					if yIsBasic {
						y = vy.(int)
					} else {
						y = basic.Int(vy)
					}
					return func(fr *frame) {
						x := basic.Uint8(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int8:
					var y int8
					if yIsBasic {
						y = vy.(int8)
					} else {
						y = basic.Int8(vy)
					}
					return func(fr *frame) {
						x := basic.Uint8(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int16:
					var y int16
					if yIsBasic {
						y = vy.(int16)
					} else {
						y = basic.Int16(vy)
					}
					return func(fr *frame) {
						x := basic.Uint8(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int32:
					var y int32
					if yIsBasic {
						y = vy.(int32)
					} else {
						y = basic.Int32(vy)
					}
					return func(fr *frame) {
						x := basic.Uint8(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int64:
					var y int64
					if yIsBasic {
						y = vy.(int64)
					} else {
						y = basic.Int64(vy)
					}
					return func(fr *frame) {
						x := basic.Uint8(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint:
					var y uint
					if yIsBasic {
						y = vy.(uint)
					} else {
						y = basic.Uint(vy)
					}
					return func(fr *frame) {
						x := basic.Uint8(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint8:
					var y uint8
					if yIsBasic {
						y = vy.(uint8)
					} else {
						y = basic.Uint8(vy)
					}
					return func(fr *frame) {
						x := basic.Uint8(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint16:
					var y uint16
					if yIsBasic {
						y = vy.(uint16)
					} else {
						y = basic.Uint16(vy)
					}
					return func(fr *frame) {
						x := basic.Uint8(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint32:
					var y uint32
					if yIsBasic {
						y = vy.(uint32)
					} else {
						y = basic.Uint32(vy)
					}
					return func(fr *frame) {
						x := basic.Uint8(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint64:
					var y uint64
					if yIsBasic {
						y = vy.(uint64)
					} else {
						y = basic.Uint64(vy)
					}
					return func(fr *frame) {
						x := basic.Uint8(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uintptr:
					var y uintptr
					if yIsBasic {
						y = vy.(uintptr)
					} else {
						y = basic.Uintptr(vy)
					}
					return func(fr *frame) {
						x := basic.Uint8(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				}
			} else {
				switch ykind {
				case reflect.Int:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Uint8(fr.reg(ix))
							y := fr.reg(iy).(int)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						x := basic.Uint8(fr.reg(ix))
						y := basic.Int(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int8:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Uint8(fr.reg(ix))
							y := fr.reg(iy).(int8)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						x := basic.Uint8(fr.reg(ix))
						y := basic.Int8(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int16:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Uint8(fr.reg(ix))
							y := fr.reg(iy).(int16)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						x := basic.Uint8(fr.reg(ix))
						y := basic.Int16(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int32:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Uint8(fr.reg(ix))
							y := fr.reg(iy).(int32)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						x := basic.Uint8(fr.reg(ix))
						y := basic.Int32(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int64:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Uint8(fr.reg(ix))
							y := fr.reg(iy).(int64)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						x := basic.Uint8(fr.reg(ix))
						y := basic.Int64(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Uint8(fr.reg(ix))
							y := fr.reg(iy).(uint)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						x := basic.Uint8(fr.reg(ix))
						y := basic.Uint(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint8:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Uint8(fr.reg(ix))
							y := fr.reg(iy).(uint8)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						x := basic.Uint8(fr.reg(ix))
						y := basic.Uint8(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint16:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Uint8(fr.reg(ix))
							y := fr.reg(iy).(uint16)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						x := basic.Uint8(fr.reg(ix))
						y := basic.Uint16(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint32:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Uint8(fr.reg(ix))
							y := fr.reg(iy).(uint32)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						x := basic.Uint8(fr.reg(ix))
						y := basic.Uint32(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint64:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Uint8(fr.reg(ix))
							y := fr.reg(iy).(uint64)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						x := basic.Uint8(fr.reg(ix))
						y := basic.Uint64(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uintptr:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Uint8(fr.reg(ix))
							y := fr.reg(iy).(uintptr)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						x := basic.Uint8(fr.reg(ix))
						y := basic.Uintptr(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				}
			}
		case reflect.Uint16:
			if kx == kindConst && ky == kindConst {
				x := basic.Uint16(vx)
				switch ykind {
				case reflect.Int:
					var y int
					if yIsBasic {
						y = vy.(int)
					} else {
						y = basic.Int(vy)
					}
					v := basic.Make(t, x<<y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Int8:
					var y int8
					if yIsBasic {
						y = vy.(int8)
					} else {
						y = basic.Int8(vy)
					}
					v := basic.Make(t, x<<y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Int16:
					var y int16
					if yIsBasic {
						y = vy.(int16)
					} else {
						y = basic.Int16(vy)
					}
					v := basic.Make(t, x<<y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Int32:
					var y int32
					if yIsBasic {
						y = vy.(int32)
					} else {
						y = basic.Int32(vy)
					}
					v := basic.Make(t, x<<y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Int64:
					var y int64
					if yIsBasic {
						y = vy.(int64)
					} else {
						y = basic.Int64(vy)
					}
					v := basic.Make(t, x<<y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uint:
					var y uint
					if yIsBasic {
						y = vy.(uint)
					} else {
						y = basic.Uint(vy)
					}
					v := basic.Make(t, x<<y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uint8:
					var y uint8
					if yIsBasic {
						y = vy.(uint8)
					} else {
						y = basic.Uint8(vy)
					}
					v := basic.Make(t, x<<y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uint16:
					var y uint16
					if yIsBasic {
						y = vy.(uint16)
					} else {
						y = basic.Uint16(vy)
					}
					v := basic.Make(t, x<<y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uint32:
					var y uint32
					if yIsBasic {
						y = vy.(uint32)
					} else {
						y = basic.Uint32(vy)
					}
					v := basic.Make(t, x<<y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uint64:
					var y uint64
					if yIsBasic {
						y = vy.(uint64)
					} else {
						y = basic.Uint64(vy)
					}
					v := basic.Make(t, x<<y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uintptr:
					var y uintptr
					if yIsBasic {
						y = vy.(uintptr)
					} else {
						y = basic.Uintptr(vy)
					}
					v := basic.Make(t, x<<y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				}
			} else if kx == kindConst {
				x := basic.Uint16(vx)
				switch ykind {
				case reflect.Int:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(int)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						y := basic.Int(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int8:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(int8)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						y := basic.Int8(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int16:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(int16)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						y := basic.Int16(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int32:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(int32)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						y := basic.Int32(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int64:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(int64)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						y := basic.Int64(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uint)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						y := basic.Uint(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint8:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uint8)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						y := basic.Uint8(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint16:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uint16)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						y := basic.Uint16(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint32:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uint32)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						y := basic.Uint32(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint64:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uint64)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						y := basic.Uint64(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uintptr:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uintptr)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						y := basic.Uintptr(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				}
			} else if ky == kindConst {
				switch ykind {
				case reflect.Int:
					var y int
					if yIsBasic {
						y = vy.(int)
					} else {
						y = basic.Int(vy)
					}
					return func(fr *frame) {
						x := basic.Uint16(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int8:
					var y int8
					if yIsBasic {
						y = vy.(int8)
					} else {
						y = basic.Int8(vy)
					}
					return func(fr *frame) {
						x := basic.Uint16(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int16:
					var y int16
					if yIsBasic {
						y = vy.(int16)
					} else {
						y = basic.Int16(vy)
					}
					return func(fr *frame) {
						x := basic.Uint16(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int32:
					var y int32
					if yIsBasic {
						y = vy.(int32)
					} else {
						y = basic.Int32(vy)
					}
					return func(fr *frame) {
						x := basic.Uint16(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int64:
					var y int64
					if yIsBasic {
						y = vy.(int64)
					} else {
						y = basic.Int64(vy)
					}
					return func(fr *frame) {
						x := basic.Uint16(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint:
					var y uint
					if yIsBasic {
						y = vy.(uint)
					} else {
						y = basic.Uint(vy)
					}
					return func(fr *frame) {
						x := basic.Uint16(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint8:
					var y uint8
					if yIsBasic {
						y = vy.(uint8)
					} else {
						y = basic.Uint8(vy)
					}
					return func(fr *frame) {
						x := basic.Uint16(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint16:
					var y uint16
					if yIsBasic {
						y = vy.(uint16)
					} else {
						y = basic.Uint16(vy)
					}
					return func(fr *frame) {
						x := basic.Uint16(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint32:
					var y uint32
					if yIsBasic {
						y = vy.(uint32)
					} else {
						y = basic.Uint32(vy)
					}
					return func(fr *frame) {
						x := basic.Uint16(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint64:
					var y uint64
					if yIsBasic {
						y = vy.(uint64)
					} else {
						y = basic.Uint64(vy)
					}
					return func(fr *frame) {
						x := basic.Uint16(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uintptr:
					var y uintptr
					if yIsBasic {
						y = vy.(uintptr)
					} else {
						y = basic.Uintptr(vy)
					}
					return func(fr *frame) {
						x := basic.Uint16(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				}
			} else {
				switch ykind {
				case reflect.Int:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Uint16(fr.reg(ix))
							y := fr.reg(iy).(int)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						x := basic.Uint16(fr.reg(ix))
						y := basic.Int(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int8:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Uint16(fr.reg(ix))
							y := fr.reg(iy).(int8)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						x := basic.Uint16(fr.reg(ix))
						y := basic.Int8(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int16:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Uint16(fr.reg(ix))
							y := fr.reg(iy).(int16)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						x := basic.Uint16(fr.reg(ix))
						y := basic.Int16(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int32:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Uint16(fr.reg(ix))
							y := fr.reg(iy).(int32)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						x := basic.Uint16(fr.reg(ix))
						y := basic.Int32(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int64:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Uint16(fr.reg(ix))
							y := fr.reg(iy).(int64)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						x := basic.Uint16(fr.reg(ix))
						y := basic.Int64(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Uint16(fr.reg(ix))
							y := fr.reg(iy).(uint)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						x := basic.Uint16(fr.reg(ix))
						y := basic.Uint(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint8:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Uint16(fr.reg(ix))
							y := fr.reg(iy).(uint8)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						x := basic.Uint16(fr.reg(ix))
						y := basic.Uint8(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint16:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Uint16(fr.reg(ix))
							y := fr.reg(iy).(uint16)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						x := basic.Uint16(fr.reg(ix))
						y := basic.Uint16(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint32:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Uint16(fr.reg(ix))
							y := fr.reg(iy).(uint32)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						x := basic.Uint16(fr.reg(ix))
						y := basic.Uint32(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint64:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Uint16(fr.reg(ix))
							y := fr.reg(iy).(uint64)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						x := basic.Uint16(fr.reg(ix))
						y := basic.Uint64(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uintptr:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Uint16(fr.reg(ix))
							y := fr.reg(iy).(uintptr)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						x := basic.Uint16(fr.reg(ix))
						y := basic.Uintptr(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				}
			}
		case reflect.Uint32:
			if kx == kindConst && ky == kindConst {
				x := basic.Uint32(vx)
				switch ykind {
				case reflect.Int:
					var y int
					if yIsBasic {
						y = vy.(int)
					} else {
						y = basic.Int(vy)
					}
					v := basic.Make(t, x<<y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Int8:
					var y int8
					if yIsBasic {
						y = vy.(int8)
					} else {
						y = basic.Int8(vy)
					}
					v := basic.Make(t, x<<y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Int16:
					var y int16
					if yIsBasic {
						y = vy.(int16)
					} else {
						y = basic.Int16(vy)
					}
					v := basic.Make(t, x<<y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Int32:
					var y int32
					if yIsBasic {
						y = vy.(int32)
					} else {
						y = basic.Int32(vy)
					}
					v := basic.Make(t, x<<y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Int64:
					var y int64
					if yIsBasic {
						y = vy.(int64)
					} else {
						y = basic.Int64(vy)
					}
					v := basic.Make(t, x<<y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uint:
					var y uint
					if yIsBasic {
						y = vy.(uint)
					} else {
						y = basic.Uint(vy)
					}
					v := basic.Make(t, x<<y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uint8:
					var y uint8
					if yIsBasic {
						y = vy.(uint8)
					} else {
						y = basic.Uint8(vy)
					}
					v := basic.Make(t, x<<y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uint16:
					var y uint16
					if yIsBasic {
						y = vy.(uint16)
					} else {
						y = basic.Uint16(vy)
					}
					v := basic.Make(t, x<<y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uint32:
					var y uint32
					if yIsBasic {
						y = vy.(uint32)
					} else {
						y = basic.Uint32(vy)
					}
					v := basic.Make(t, x<<y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uint64:
					var y uint64
					if yIsBasic {
						y = vy.(uint64)
					} else {
						y = basic.Uint64(vy)
					}
					v := basic.Make(t, x<<y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uintptr:
					var y uintptr
					if yIsBasic {
						y = vy.(uintptr)
					} else {
						y = basic.Uintptr(vy)
					}
					v := basic.Make(t, x<<y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				}
			} else if kx == kindConst {
				x := basic.Uint32(vx)
				switch ykind {
				case reflect.Int:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(int)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						y := basic.Int(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int8:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(int8)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						y := basic.Int8(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int16:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(int16)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						y := basic.Int16(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int32:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(int32)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						y := basic.Int32(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int64:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(int64)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						y := basic.Int64(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uint)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						y := basic.Uint(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint8:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uint8)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						y := basic.Uint8(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint16:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uint16)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						y := basic.Uint16(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint32:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uint32)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						y := basic.Uint32(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint64:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uint64)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						y := basic.Uint64(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uintptr:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uintptr)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						y := basic.Uintptr(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				}
			} else if ky == kindConst {
				switch ykind {
				case reflect.Int:
					var y int
					if yIsBasic {
						y = vy.(int)
					} else {
						y = basic.Int(vy)
					}
					return func(fr *frame) {
						x := basic.Uint32(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int8:
					var y int8
					if yIsBasic {
						y = vy.(int8)
					} else {
						y = basic.Int8(vy)
					}
					return func(fr *frame) {
						x := basic.Uint32(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int16:
					var y int16
					if yIsBasic {
						y = vy.(int16)
					} else {
						y = basic.Int16(vy)
					}
					return func(fr *frame) {
						x := basic.Uint32(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int32:
					var y int32
					if yIsBasic {
						y = vy.(int32)
					} else {
						y = basic.Int32(vy)
					}
					return func(fr *frame) {
						x := basic.Uint32(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int64:
					var y int64
					if yIsBasic {
						y = vy.(int64)
					} else {
						y = basic.Int64(vy)
					}
					return func(fr *frame) {
						x := basic.Uint32(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint:
					var y uint
					if yIsBasic {
						y = vy.(uint)
					} else {
						y = basic.Uint(vy)
					}
					return func(fr *frame) {
						x := basic.Uint32(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint8:
					var y uint8
					if yIsBasic {
						y = vy.(uint8)
					} else {
						y = basic.Uint8(vy)
					}
					return func(fr *frame) {
						x := basic.Uint32(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint16:
					var y uint16
					if yIsBasic {
						y = vy.(uint16)
					} else {
						y = basic.Uint16(vy)
					}
					return func(fr *frame) {
						x := basic.Uint32(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint32:
					var y uint32
					if yIsBasic {
						y = vy.(uint32)
					} else {
						y = basic.Uint32(vy)
					}
					return func(fr *frame) {
						x := basic.Uint32(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint64:
					var y uint64
					if yIsBasic {
						y = vy.(uint64)
					} else {
						y = basic.Uint64(vy)
					}
					return func(fr *frame) {
						x := basic.Uint32(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uintptr:
					var y uintptr
					if yIsBasic {
						y = vy.(uintptr)
					} else {
						y = basic.Uintptr(vy)
					}
					return func(fr *frame) {
						x := basic.Uint32(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				}
			} else {
				switch ykind {
				case reflect.Int:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Uint32(fr.reg(ix))
							y := fr.reg(iy).(int)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						x := basic.Uint32(fr.reg(ix))
						y := basic.Int(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int8:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Uint32(fr.reg(ix))
							y := fr.reg(iy).(int8)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						x := basic.Uint32(fr.reg(ix))
						y := basic.Int8(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int16:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Uint32(fr.reg(ix))
							y := fr.reg(iy).(int16)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						x := basic.Uint32(fr.reg(ix))
						y := basic.Int16(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int32:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Uint32(fr.reg(ix))
							y := fr.reg(iy).(int32)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						x := basic.Uint32(fr.reg(ix))
						y := basic.Int32(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int64:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Uint32(fr.reg(ix))
							y := fr.reg(iy).(int64)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						x := basic.Uint32(fr.reg(ix))
						y := basic.Int64(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Uint32(fr.reg(ix))
							y := fr.reg(iy).(uint)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						x := basic.Uint32(fr.reg(ix))
						y := basic.Uint(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint8:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Uint32(fr.reg(ix))
							y := fr.reg(iy).(uint8)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						x := basic.Uint32(fr.reg(ix))
						y := basic.Uint8(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint16:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Uint32(fr.reg(ix))
							y := fr.reg(iy).(uint16)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						x := basic.Uint32(fr.reg(ix))
						y := basic.Uint16(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint32:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Uint32(fr.reg(ix))
							y := fr.reg(iy).(uint32)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						x := basic.Uint32(fr.reg(ix))
						y := basic.Uint32(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint64:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Uint32(fr.reg(ix))
							y := fr.reg(iy).(uint64)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						x := basic.Uint32(fr.reg(ix))
						y := basic.Uint64(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uintptr:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Uint32(fr.reg(ix))
							y := fr.reg(iy).(uintptr)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						x := basic.Uint32(fr.reg(ix))
						y := basic.Uintptr(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				}
			}
		case reflect.Uint64:
			if kx == kindConst && ky == kindConst {
				x := basic.Uint64(vx)
				switch ykind {
				case reflect.Int:
					var y int
					if yIsBasic {
						y = vy.(int)
					} else {
						y = basic.Int(vy)
					}
					v := basic.Make(t, x<<y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Int8:
					var y int8
					if yIsBasic {
						y = vy.(int8)
					} else {
						y = basic.Int8(vy)
					}
					v := basic.Make(t, x<<y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Int16:
					var y int16
					if yIsBasic {
						y = vy.(int16)
					} else {
						y = basic.Int16(vy)
					}
					v := basic.Make(t, x<<y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Int32:
					var y int32
					if yIsBasic {
						y = vy.(int32)
					} else {
						y = basic.Int32(vy)
					}
					v := basic.Make(t, x<<y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Int64:
					var y int64
					if yIsBasic {
						y = vy.(int64)
					} else {
						y = basic.Int64(vy)
					}
					v := basic.Make(t, x<<y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uint:
					var y uint
					if yIsBasic {
						y = vy.(uint)
					} else {
						y = basic.Uint(vy)
					}
					v := basic.Make(t, x<<y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uint8:
					var y uint8
					if yIsBasic {
						y = vy.(uint8)
					} else {
						y = basic.Uint8(vy)
					}
					v := basic.Make(t, x<<y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uint16:
					var y uint16
					if yIsBasic {
						y = vy.(uint16)
					} else {
						y = basic.Uint16(vy)
					}
					v := basic.Make(t, x<<y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uint32:
					var y uint32
					if yIsBasic {
						y = vy.(uint32)
					} else {
						y = basic.Uint32(vy)
					}
					v := basic.Make(t, x<<y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uint64:
					var y uint64
					if yIsBasic {
						y = vy.(uint64)
					} else {
						y = basic.Uint64(vy)
					}
					v := basic.Make(t, x<<y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uintptr:
					var y uintptr
					if yIsBasic {
						y = vy.(uintptr)
					} else {
						y = basic.Uintptr(vy)
					}
					v := basic.Make(t, x<<y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				}
			} else if kx == kindConst {
				x := basic.Uint64(vx)
				switch ykind {
				case reflect.Int:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(int)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						y := basic.Int(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int8:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(int8)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						y := basic.Int8(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int16:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(int16)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						y := basic.Int16(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int32:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(int32)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						y := basic.Int32(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int64:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(int64)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						y := basic.Int64(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uint)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						y := basic.Uint(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint8:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uint8)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						y := basic.Uint8(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint16:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uint16)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						y := basic.Uint16(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint32:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uint32)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						y := basic.Uint32(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint64:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uint64)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						y := basic.Uint64(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uintptr:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uintptr)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						y := basic.Uintptr(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				}
			} else if ky == kindConst {
				switch ykind {
				case reflect.Int:
					var y int
					if yIsBasic {
						y = vy.(int)
					} else {
						y = basic.Int(vy)
					}
					return func(fr *frame) {
						x := basic.Uint64(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int8:
					var y int8
					if yIsBasic {
						y = vy.(int8)
					} else {
						y = basic.Int8(vy)
					}
					return func(fr *frame) {
						x := basic.Uint64(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int16:
					var y int16
					if yIsBasic {
						y = vy.(int16)
					} else {
						y = basic.Int16(vy)
					}
					return func(fr *frame) {
						x := basic.Uint64(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int32:
					var y int32
					if yIsBasic {
						y = vy.(int32)
					} else {
						y = basic.Int32(vy)
					}
					return func(fr *frame) {
						x := basic.Uint64(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int64:
					var y int64
					if yIsBasic {
						y = vy.(int64)
					} else {
						y = basic.Int64(vy)
					}
					return func(fr *frame) {
						x := basic.Uint64(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint:
					var y uint
					if yIsBasic {
						y = vy.(uint)
					} else {
						y = basic.Uint(vy)
					}
					return func(fr *frame) {
						x := basic.Uint64(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint8:
					var y uint8
					if yIsBasic {
						y = vy.(uint8)
					} else {
						y = basic.Uint8(vy)
					}
					return func(fr *frame) {
						x := basic.Uint64(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint16:
					var y uint16
					if yIsBasic {
						y = vy.(uint16)
					} else {
						y = basic.Uint16(vy)
					}
					return func(fr *frame) {
						x := basic.Uint64(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint32:
					var y uint32
					if yIsBasic {
						y = vy.(uint32)
					} else {
						y = basic.Uint32(vy)
					}
					return func(fr *frame) {
						x := basic.Uint64(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint64:
					var y uint64
					if yIsBasic {
						y = vy.(uint64)
					} else {
						y = basic.Uint64(vy)
					}
					return func(fr *frame) {
						x := basic.Uint64(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uintptr:
					var y uintptr
					if yIsBasic {
						y = vy.(uintptr)
					} else {
						y = basic.Uintptr(vy)
					}
					return func(fr *frame) {
						x := basic.Uint64(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				}
			} else {
				switch ykind {
				case reflect.Int:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Uint64(fr.reg(ix))
							y := fr.reg(iy).(int)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						x := basic.Uint64(fr.reg(ix))
						y := basic.Int(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int8:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Uint64(fr.reg(ix))
							y := fr.reg(iy).(int8)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						x := basic.Uint64(fr.reg(ix))
						y := basic.Int8(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int16:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Uint64(fr.reg(ix))
							y := fr.reg(iy).(int16)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						x := basic.Uint64(fr.reg(ix))
						y := basic.Int16(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int32:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Uint64(fr.reg(ix))
							y := fr.reg(iy).(int32)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						x := basic.Uint64(fr.reg(ix))
						y := basic.Int32(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int64:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Uint64(fr.reg(ix))
							y := fr.reg(iy).(int64)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						x := basic.Uint64(fr.reg(ix))
						y := basic.Int64(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Uint64(fr.reg(ix))
							y := fr.reg(iy).(uint)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						x := basic.Uint64(fr.reg(ix))
						y := basic.Uint(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint8:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Uint64(fr.reg(ix))
							y := fr.reg(iy).(uint8)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						x := basic.Uint64(fr.reg(ix))
						y := basic.Uint8(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint16:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Uint64(fr.reg(ix))
							y := fr.reg(iy).(uint16)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						x := basic.Uint64(fr.reg(ix))
						y := basic.Uint16(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint32:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Uint64(fr.reg(ix))
							y := fr.reg(iy).(uint32)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						x := basic.Uint64(fr.reg(ix))
						y := basic.Uint32(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint64:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Uint64(fr.reg(ix))
							y := fr.reg(iy).(uint64)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						x := basic.Uint64(fr.reg(ix))
						y := basic.Uint64(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uintptr:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Uint64(fr.reg(ix))
							y := fr.reg(iy).(uintptr)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						x := basic.Uint64(fr.reg(ix))
						y := basic.Uintptr(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				}
			}
		case reflect.Uintptr:
			if kx == kindConst && ky == kindConst {
				x := basic.Uintptr(vx)
				switch ykind {
				case reflect.Int:
					var y int
					if yIsBasic {
						y = vy.(int)
					} else {
						y = basic.Int(vy)
					}
					v := basic.Make(t, x<<y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Int8:
					var y int8
					if yIsBasic {
						y = vy.(int8)
					} else {
						y = basic.Int8(vy)
					}
					v := basic.Make(t, x<<y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Int16:
					var y int16
					if yIsBasic {
						y = vy.(int16)
					} else {
						y = basic.Int16(vy)
					}
					v := basic.Make(t, x<<y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Int32:
					var y int32
					if yIsBasic {
						y = vy.(int32)
					} else {
						y = basic.Int32(vy)
					}
					v := basic.Make(t, x<<y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Int64:
					var y int64
					if yIsBasic {
						y = vy.(int64)
					} else {
						y = basic.Int64(vy)
					}
					v := basic.Make(t, x<<y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uint:
					var y uint
					if yIsBasic {
						y = vy.(uint)
					} else {
						y = basic.Uint(vy)
					}
					v := basic.Make(t, x<<y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uint8:
					var y uint8
					if yIsBasic {
						y = vy.(uint8)
					} else {
						y = basic.Uint8(vy)
					}
					v := basic.Make(t, x<<y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uint16:
					var y uint16
					if yIsBasic {
						y = vy.(uint16)
					} else {
						y = basic.Uint16(vy)
					}
					v := basic.Make(t, x<<y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uint32:
					var y uint32
					if yIsBasic {
						y = vy.(uint32)
					} else {
						y = basic.Uint32(vy)
					}
					v := basic.Make(t, x<<y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uint64:
					var y uint64
					if yIsBasic {
						y = vy.(uint64)
					} else {
						y = basic.Uint64(vy)
					}
					v := basic.Make(t, x<<y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uintptr:
					var y uintptr
					if yIsBasic {
						y = vy.(uintptr)
					} else {
						y = basic.Uintptr(vy)
					}
					v := basic.Make(t, x<<y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				}
			} else if kx == kindConst {
				x := basic.Uintptr(vx)
				switch ykind {
				case reflect.Int:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(int)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						y := basic.Int(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int8:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(int8)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						y := basic.Int8(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int16:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(int16)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						y := basic.Int16(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int32:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(int32)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						y := basic.Int32(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int64:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(int64)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						y := basic.Int64(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uint)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						y := basic.Uint(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint8:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uint8)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						y := basic.Uint8(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint16:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uint16)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						y := basic.Uint16(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint32:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uint32)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						y := basic.Uint32(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint64:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uint64)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						y := basic.Uint64(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uintptr:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uintptr)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						y := basic.Uintptr(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				}
			} else if ky == kindConst {
				switch ykind {
				case reflect.Int:
					var y int
					if yIsBasic {
						y = vy.(int)
					} else {
						y = basic.Int(vy)
					}
					return func(fr *frame) {
						x := basic.Uintptr(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int8:
					var y int8
					if yIsBasic {
						y = vy.(int8)
					} else {
						y = basic.Int8(vy)
					}
					return func(fr *frame) {
						x := basic.Uintptr(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int16:
					var y int16
					if yIsBasic {
						y = vy.(int16)
					} else {
						y = basic.Int16(vy)
					}
					return func(fr *frame) {
						x := basic.Uintptr(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int32:
					var y int32
					if yIsBasic {
						y = vy.(int32)
					} else {
						y = basic.Int32(vy)
					}
					return func(fr *frame) {
						x := basic.Uintptr(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int64:
					var y int64
					if yIsBasic {
						y = vy.(int64)
					} else {
						y = basic.Int64(vy)
					}
					return func(fr *frame) {
						x := basic.Uintptr(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint:
					var y uint
					if yIsBasic {
						y = vy.(uint)
					} else {
						y = basic.Uint(vy)
					}
					return func(fr *frame) {
						x := basic.Uintptr(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint8:
					var y uint8
					if yIsBasic {
						y = vy.(uint8)
					} else {
						y = basic.Uint8(vy)
					}
					return func(fr *frame) {
						x := basic.Uintptr(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint16:
					var y uint16
					if yIsBasic {
						y = vy.(uint16)
					} else {
						y = basic.Uint16(vy)
					}
					return func(fr *frame) {
						x := basic.Uintptr(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint32:
					var y uint32
					if yIsBasic {
						y = vy.(uint32)
					} else {
						y = basic.Uint32(vy)
					}
					return func(fr *frame) {
						x := basic.Uintptr(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint64:
					var y uint64
					if yIsBasic {
						y = vy.(uint64)
					} else {
						y = basic.Uint64(vy)
					}
					return func(fr *frame) {
						x := basic.Uintptr(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uintptr:
					var y uintptr
					if yIsBasic {
						y = vy.(uintptr)
					} else {
						y = basic.Uintptr(vy)
					}
					return func(fr *frame) {
						x := basic.Uintptr(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				}
			} else {
				switch ykind {
				case reflect.Int:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Uintptr(fr.reg(ix))
							y := fr.reg(iy).(int)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						x := basic.Uintptr(fr.reg(ix))
						y := basic.Int(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int8:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Uintptr(fr.reg(ix))
							y := fr.reg(iy).(int8)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						x := basic.Uintptr(fr.reg(ix))
						y := basic.Int8(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int16:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Uintptr(fr.reg(ix))
							y := fr.reg(iy).(int16)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						x := basic.Uintptr(fr.reg(ix))
						y := basic.Int16(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int32:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Uintptr(fr.reg(ix))
							y := fr.reg(iy).(int32)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						x := basic.Uintptr(fr.reg(ix))
						y := basic.Int32(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Int64:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Uintptr(fr.reg(ix))
							y := fr.reg(iy).(int64)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						x := basic.Uintptr(fr.reg(ix))
						y := basic.Int64(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Uintptr(fr.reg(ix))
							y := fr.reg(iy).(uint)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						x := basic.Uintptr(fr.reg(ix))
						y := basic.Uint(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint8:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Uintptr(fr.reg(ix))
							y := fr.reg(iy).(uint8)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						x := basic.Uintptr(fr.reg(ix))
						y := basic.Uint8(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint16:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Uintptr(fr.reg(ix))
							y := fr.reg(iy).(uint16)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						x := basic.Uintptr(fr.reg(ix))
						y := basic.Uint16(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint32:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Uintptr(fr.reg(ix))
							y := fr.reg(iy).(uint32)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						x := basic.Uintptr(fr.reg(ix))
						y := basic.Uint32(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uint64:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Uintptr(fr.reg(ix))
							y := fr.reg(iy).(uint64)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						x := basic.Uintptr(fr.reg(ix))
						y := basic.Uint64(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x<<y))
					}
				case reflect.Uintptr:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Uintptr(fr.reg(ix))
							y := fr.reg(iy).(uintptr)
							fr.setReg(ir, basic.Make(t, x<<y))
						}
					}
					return func(fr *frame) {
						x := basic.Uintptr(fr.reg(ix))
						y := basic.Uintptr(fr.reg(iy))
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
	xIsBasic := xtyp.PkgPath() == ""
	yIsBasic := ytyp.PkgPath() == ""
	if xIsBasic {
		switch xkind {
		case reflect.Int:
			if kx == kindConst && ky == kindConst {
				x := vx.(int)
				switch ykind {
				case reflect.Int:
					var y int
					if yIsBasic {
						y = vy.(int)
					} else {
						y = basic.Int(vy)
					}
					v := x >> y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Int8:
					var y int8
					if yIsBasic {
						y = vy.(int8)
					} else {
						y = basic.Int8(vy)
					}
					v := x >> y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Int16:
					var y int16
					if yIsBasic {
						y = vy.(int16)
					} else {
						y = basic.Int16(vy)
					}
					v := x >> y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Int32:
					var y int32
					if yIsBasic {
						y = vy.(int32)
					} else {
						y = basic.Int32(vy)
					}
					v := x >> y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Int64:
					var y int64
					if yIsBasic {
						y = vy.(int64)
					} else {
						y = basic.Int64(vy)
					}
					v := x >> y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uint:
					var y uint
					if yIsBasic {
						y = vy.(uint)
					} else {
						y = basic.Uint(vy)
					}
					v := x >> y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uint8:
					var y uint8
					if yIsBasic {
						y = vy.(uint8)
					} else {
						y = basic.Uint8(vy)
					}
					v := x >> y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uint16:
					var y uint16
					if yIsBasic {
						y = vy.(uint16)
					} else {
						y = basic.Uint16(vy)
					}
					v := x >> y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uint32:
					var y uint32
					if yIsBasic {
						y = vy.(uint32)
					} else {
						y = basic.Uint32(vy)
					}
					v := x >> y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uint64:
					var y uint64
					if yIsBasic {
						y = vy.(uint64)
					} else {
						y = basic.Uint64(vy)
					}
					v := x >> y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uintptr:
					var y uintptr
					if yIsBasic {
						y = vy.(uintptr)
					} else {
						y = basic.Uintptr(vy)
					}
					v := x >> y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				}
			} else if kx == kindConst {
				x := vx.(int)
				switch ykind {
				case reflect.Int:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(int)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						y := basic.Int(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				case reflect.Int8:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(int8)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						y := basic.Int8(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				case reflect.Int16:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(int16)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						y := basic.Int16(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				case reflect.Int32:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(int32)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						y := basic.Int32(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				case reflect.Int64:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(int64)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						y := basic.Int64(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uint)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						y := basic.Uint(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint8:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uint8)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						y := basic.Uint8(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint16:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uint16)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						y := basic.Uint16(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint32:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uint32)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						y := basic.Uint32(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint64:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uint64)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						y := basic.Uint64(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				case reflect.Uintptr:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uintptr)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						y := basic.Uintptr(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				}
			} else if ky == kindConst {
				switch ykind {
				case reflect.Int:
					var y int
					if yIsBasic {
						y = vy.(int)
					} else {
						y = basic.Int(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(int)
						fr.setReg(ir, x>>y)
					}
				case reflect.Int8:
					var y int8
					if yIsBasic {
						y = vy.(int8)
					} else {
						y = basic.Int8(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(int)
						fr.setReg(ir, x>>y)
					}
				case reflect.Int16:
					var y int16
					if yIsBasic {
						y = vy.(int16)
					} else {
						y = basic.Int16(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(int)
						fr.setReg(ir, x>>y)
					}
				case reflect.Int32:
					var y int32
					if yIsBasic {
						y = vy.(int32)
					} else {
						y = basic.Int32(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(int)
						fr.setReg(ir, x>>y)
					}
				case reflect.Int64:
					var y int64
					if yIsBasic {
						y = vy.(int64)
					} else {
						y = basic.Int64(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(int)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint:
					var y uint
					if yIsBasic {
						y = vy.(uint)
					} else {
						y = basic.Uint(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(int)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint8:
					var y uint8
					if yIsBasic {
						y = vy.(uint8)
					} else {
						y = basic.Uint8(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(int)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint16:
					var y uint16
					if yIsBasic {
						y = vy.(uint16)
					} else {
						y = basic.Uint16(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(int)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint32:
					var y uint32
					if yIsBasic {
						y = vy.(uint32)
					} else {
						y = basic.Uint32(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(int)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint64:
					var y uint64
					if yIsBasic {
						y = vy.(uint64)
					} else {
						y = basic.Uint64(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(int)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uintptr:
					var y uintptr
					if yIsBasic {
						y = vy.(uintptr)
					} else {
						y = basic.Uintptr(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(int)
						fr.setReg(ir, x>>y)
					}
				}
			} else {
				switch ykind {
				case reflect.Int:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(int)
							y := fr.reg(iy).(int)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(int)
						y := basic.Int(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				case reflect.Int8:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(int)
							y := fr.reg(iy).(int8)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(int)
						y := basic.Int8(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				case reflect.Int16:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(int)
							y := fr.reg(iy).(int16)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(int)
						y := basic.Int16(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				case reflect.Int32:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(int)
							y := fr.reg(iy).(int32)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(int)
						y := basic.Int32(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				case reflect.Int64:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(int)
							y := fr.reg(iy).(int64)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(int)
						y := basic.Int64(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(int)
							y := fr.reg(iy).(uint)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(int)
						y := basic.Uint(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint8:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(int)
							y := fr.reg(iy).(uint8)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(int)
						y := basic.Uint8(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint16:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(int)
							y := fr.reg(iy).(uint16)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(int)
						y := basic.Uint16(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint32:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(int)
							y := fr.reg(iy).(uint32)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(int)
						y := basic.Uint32(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint64:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(int)
							y := fr.reg(iy).(uint64)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(int)
						y := basic.Uint64(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				case reflect.Uintptr:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(int)
							y := fr.reg(iy).(uintptr)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(int)
						y := basic.Uintptr(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				}
			}
		case reflect.Int8:
			if kx == kindConst && ky == kindConst {
				x := vx.(int8)
				switch ykind {
				case reflect.Int:
					var y int
					if yIsBasic {
						y = vy.(int)
					} else {
						y = basic.Int(vy)
					}
					v := x >> y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Int8:
					var y int8
					if yIsBasic {
						y = vy.(int8)
					} else {
						y = basic.Int8(vy)
					}
					v := x >> y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Int16:
					var y int16
					if yIsBasic {
						y = vy.(int16)
					} else {
						y = basic.Int16(vy)
					}
					v := x >> y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Int32:
					var y int32
					if yIsBasic {
						y = vy.(int32)
					} else {
						y = basic.Int32(vy)
					}
					v := x >> y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Int64:
					var y int64
					if yIsBasic {
						y = vy.(int64)
					} else {
						y = basic.Int64(vy)
					}
					v := x >> y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uint:
					var y uint
					if yIsBasic {
						y = vy.(uint)
					} else {
						y = basic.Uint(vy)
					}
					v := x >> y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uint8:
					var y uint8
					if yIsBasic {
						y = vy.(uint8)
					} else {
						y = basic.Uint8(vy)
					}
					v := x >> y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uint16:
					var y uint16
					if yIsBasic {
						y = vy.(uint16)
					} else {
						y = basic.Uint16(vy)
					}
					v := x >> y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uint32:
					var y uint32
					if yIsBasic {
						y = vy.(uint32)
					} else {
						y = basic.Uint32(vy)
					}
					v := x >> y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uint64:
					var y uint64
					if yIsBasic {
						y = vy.(uint64)
					} else {
						y = basic.Uint64(vy)
					}
					v := x >> y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uintptr:
					var y uintptr
					if yIsBasic {
						y = vy.(uintptr)
					} else {
						y = basic.Uintptr(vy)
					}
					v := x >> y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				}
			} else if kx == kindConst {
				x := vx.(int8)
				switch ykind {
				case reflect.Int:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(int)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						y := basic.Int(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				case reflect.Int8:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(int8)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						y := basic.Int8(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				case reflect.Int16:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(int16)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						y := basic.Int16(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				case reflect.Int32:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(int32)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						y := basic.Int32(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				case reflect.Int64:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(int64)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						y := basic.Int64(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uint)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						y := basic.Uint(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint8:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uint8)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						y := basic.Uint8(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint16:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uint16)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						y := basic.Uint16(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint32:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uint32)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						y := basic.Uint32(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint64:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uint64)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						y := basic.Uint64(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				case reflect.Uintptr:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uintptr)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						y := basic.Uintptr(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				}
			} else if ky == kindConst {
				switch ykind {
				case reflect.Int:
					var y int
					if yIsBasic {
						y = vy.(int)
					} else {
						y = basic.Int(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(int8)
						fr.setReg(ir, x>>y)
					}
				case reflect.Int8:
					var y int8
					if yIsBasic {
						y = vy.(int8)
					} else {
						y = basic.Int8(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(int8)
						fr.setReg(ir, x>>y)
					}
				case reflect.Int16:
					var y int16
					if yIsBasic {
						y = vy.(int16)
					} else {
						y = basic.Int16(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(int8)
						fr.setReg(ir, x>>y)
					}
				case reflect.Int32:
					var y int32
					if yIsBasic {
						y = vy.(int32)
					} else {
						y = basic.Int32(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(int8)
						fr.setReg(ir, x>>y)
					}
				case reflect.Int64:
					var y int64
					if yIsBasic {
						y = vy.(int64)
					} else {
						y = basic.Int64(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(int8)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint:
					var y uint
					if yIsBasic {
						y = vy.(uint)
					} else {
						y = basic.Uint(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(int8)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint8:
					var y uint8
					if yIsBasic {
						y = vy.(uint8)
					} else {
						y = basic.Uint8(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(int8)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint16:
					var y uint16
					if yIsBasic {
						y = vy.(uint16)
					} else {
						y = basic.Uint16(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(int8)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint32:
					var y uint32
					if yIsBasic {
						y = vy.(uint32)
					} else {
						y = basic.Uint32(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(int8)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint64:
					var y uint64
					if yIsBasic {
						y = vy.(uint64)
					} else {
						y = basic.Uint64(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(int8)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uintptr:
					var y uintptr
					if yIsBasic {
						y = vy.(uintptr)
					} else {
						y = basic.Uintptr(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(int8)
						fr.setReg(ir, x>>y)
					}
				}
			} else {
				switch ykind {
				case reflect.Int:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(int8)
							y := fr.reg(iy).(int)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(int8)
						y := basic.Int(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				case reflect.Int8:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(int8)
							y := fr.reg(iy).(int8)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(int8)
						y := basic.Int8(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				case reflect.Int16:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(int8)
							y := fr.reg(iy).(int16)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(int8)
						y := basic.Int16(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				case reflect.Int32:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(int8)
							y := fr.reg(iy).(int32)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(int8)
						y := basic.Int32(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				case reflect.Int64:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(int8)
							y := fr.reg(iy).(int64)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(int8)
						y := basic.Int64(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(int8)
							y := fr.reg(iy).(uint)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(int8)
						y := basic.Uint(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint8:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(int8)
							y := fr.reg(iy).(uint8)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(int8)
						y := basic.Uint8(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint16:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(int8)
							y := fr.reg(iy).(uint16)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(int8)
						y := basic.Uint16(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint32:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(int8)
							y := fr.reg(iy).(uint32)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(int8)
						y := basic.Uint32(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint64:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(int8)
							y := fr.reg(iy).(uint64)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(int8)
						y := basic.Uint64(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				case reflect.Uintptr:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(int8)
							y := fr.reg(iy).(uintptr)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(int8)
						y := basic.Uintptr(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				}
			}
		case reflect.Int16:
			if kx == kindConst && ky == kindConst {
				x := vx.(int16)
				switch ykind {
				case reflect.Int:
					var y int
					if yIsBasic {
						y = vy.(int)
					} else {
						y = basic.Int(vy)
					}
					v := x >> y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Int8:
					var y int8
					if yIsBasic {
						y = vy.(int8)
					} else {
						y = basic.Int8(vy)
					}
					v := x >> y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Int16:
					var y int16
					if yIsBasic {
						y = vy.(int16)
					} else {
						y = basic.Int16(vy)
					}
					v := x >> y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Int32:
					var y int32
					if yIsBasic {
						y = vy.(int32)
					} else {
						y = basic.Int32(vy)
					}
					v := x >> y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Int64:
					var y int64
					if yIsBasic {
						y = vy.(int64)
					} else {
						y = basic.Int64(vy)
					}
					v := x >> y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uint:
					var y uint
					if yIsBasic {
						y = vy.(uint)
					} else {
						y = basic.Uint(vy)
					}
					v := x >> y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uint8:
					var y uint8
					if yIsBasic {
						y = vy.(uint8)
					} else {
						y = basic.Uint8(vy)
					}
					v := x >> y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uint16:
					var y uint16
					if yIsBasic {
						y = vy.(uint16)
					} else {
						y = basic.Uint16(vy)
					}
					v := x >> y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uint32:
					var y uint32
					if yIsBasic {
						y = vy.(uint32)
					} else {
						y = basic.Uint32(vy)
					}
					v := x >> y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uint64:
					var y uint64
					if yIsBasic {
						y = vy.(uint64)
					} else {
						y = basic.Uint64(vy)
					}
					v := x >> y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uintptr:
					var y uintptr
					if yIsBasic {
						y = vy.(uintptr)
					} else {
						y = basic.Uintptr(vy)
					}
					v := x >> y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				}
			} else if kx == kindConst {
				x := vx.(int16)
				switch ykind {
				case reflect.Int:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(int)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						y := basic.Int(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				case reflect.Int8:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(int8)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						y := basic.Int8(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				case reflect.Int16:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(int16)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						y := basic.Int16(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				case reflect.Int32:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(int32)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						y := basic.Int32(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				case reflect.Int64:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(int64)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						y := basic.Int64(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uint)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						y := basic.Uint(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint8:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uint8)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						y := basic.Uint8(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint16:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uint16)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						y := basic.Uint16(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint32:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uint32)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						y := basic.Uint32(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint64:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uint64)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						y := basic.Uint64(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				case reflect.Uintptr:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uintptr)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						y := basic.Uintptr(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				}
			} else if ky == kindConst {
				switch ykind {
				case reflect.Int:
					var y int
					if yIsBasic {
						y = vy.(int)
					} else {
						y = basic.Int(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(int16)
						fr.setReg(ir, x>>y)
					}
				case reflect.Int8:
					var y int8
					if yIsBasic {
						y = vy.(int8)
					} else {
						y = basic.Int8(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(int16)
						fr.setReg(ir, x>>y)
					}
				case reflect.Int16:
					var y int16
					if yIsBasic {
						y = vy.(int16)
					} else {
						y = basic.Int16(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(int16)
						fr.setReg(ir, x>>y)
					}
				case reflect.Int32:
					var y int32
					if yIsBasic {
						y = vy.(int32)
					} else {
						y = basic.Int32(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(int16)
						fr.setReg(ir, x>>y)
					}
				case reflect.Int64:
					var y int64
					if yIsBasic {
						y = vy.(int64)
					} else {
						y = basic.Int64(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(int16)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint:
					var y uint
					if yIsBasic {
						y = vy.(uint)
					} else {
						y = basic.Uint(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(int16)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint8:
					var y uint8
					if yIsBasic {
						y = vy.(uint8)
					} else {
						y = basic.Uint8(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(int16)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint16:
					var y uint16
					if yIsBasic {
						y = vy.(uint16)
					} else {
						y = basic.Uint16(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(int16)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint32:
					var y uint32
					if yIsBasic {
						y = vy.(uint32)
					} else {
						y = basic.Uint32(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(int16)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint64:
					var y uint64
					if yIsBasic {
						y = vy.(uint64)
					} else {
						y = basic.Uint64(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(int16)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uintptr:
					var y uintptr
					if yIsBasic {
						y = vy.(uintptr)
					} else {
						y = basic.Uintptr(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(int16)
						fr.setReg(ir, x>>y)
					}
				}
			} else {
				switch ykind {
				case reflect.Int:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(int16)
							y := fr.reg(iy).(int)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(int16)
						y := basic.Int(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				case reflect.Int8:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(int16)
							y := fr.reg(iy).(int8)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(int16)
						y := basic.Int8(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				case reflect.Int16:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(int16)
							y := fr.reg(iy).(int16)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(int16)
						y := basic.Int16(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				case reflect.Int32:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(int16)
							y := fr.reg(iy).(int32)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(int16)
						y := basic.Int32(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				case reflect.Int64:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(int16)
							y := fr.reg(iy).(int64)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(int16)
						y := basic.Int64(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(int16)
							y := fr.reg(iy).(uint)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(int16)
						y := basic.Uint(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint8:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(int16)
							y := fr.reg(iy).(uint8)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(int16)
						y := basic.Uint8(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint16:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(int16)
							y := fr.reg(iy).(uint16)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(int16)
						y := basic.Uint16(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint32:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(int16)
							y := fr.reg(iy).(uint32)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(int16)
						y := basic.Uint32(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint64:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(int16)
							y := fr.reg(iy).(uint64)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(int16)
						y := basic.Uint64(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				case reflect.Uintptr:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(int16)
							y := fr.reg(iy).(uintptr)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(int16)
						y := basic.Uintptr(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				}
			}
		case reflect.Int32:
			if kx == kindConst && ky == kindConst {
				x := vx.(int32)
				switch ykind {
				case reflect.Int:
					var y int
					if yIsBasic {
						y = vy.(int)
					} else {
						y = basic.Int(vy)
					}
					v := x >> y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Int8:
					var y int8
					if yIsBasic {
						y = vy.(int8)
					} else {
						y = basic.Int8(vy)
					}
					v := x >> y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Int16:
					var y int16
					if yIsBasic {
						y = vy.(int16)
					} else {
						y = basic.Int16(vy)
					}
					v := x >> y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Int32:
					var y int32
					if yIsBasic {
						y = vy.(int32)
					} else {
						y = basic.Int32(vy)
					}
					v := x >> y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Int64:
					var y int64
					if yIsBasic {
						y = vy.(int64)
					} else {
						y = basic.Int64(vy)
					}
					v := x >> y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uint:
					var y uint
					if yIsBasic {
						y = vy.(uint)
					} else {
						y = basic.Uint(vy)
					}
					v := x >> y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uint8:
					var y uint8
					if yIsBasic {
						y = vy.(uint8)
					} else {
						y = basic.Uint8(vy)
					}
					v := x >> y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uint16:
					var y uint16
					if yIsBasic {
						y = vy.(uint16)
					} else {
						y = basic.Uint16(vy)
					}
					v := x >> y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uint32:
					var y uint32
					if yIsBasic {
						y = vy.(uint32)
					} else {
						y = basic.Uint32(vy)
					}
					v := x >> y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uint64:
					var y uint64
					if yIsBasic {
						y = vy.(uint64)
					} else {
						y = basic.Uint64(vy)
					}
					v := x >> y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uintptr:
					var y uintptr
					if yIsBasic {
						y = vy.(uintptr)
					} else {
						y = basic.Uintptr(vy)
					}
					v := x >> y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				}
			} else if kx == kindConst {
				x := vx.(int32)
				switch ykind {
				case reflect.Int:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(int)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						y := basic.Int(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				case reflect.Int8:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(int8)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						y := basic.Int8(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				case reflect.Int16:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(int16)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						y := basic.Int16(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				case reflect.Int32:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(int32)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						y := basic.Int32(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				case reflect.Int64:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(int64)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						y := basic.Int64(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uint)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						y := basic.Uint(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint8:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uint8)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						y := basic.Uint8(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint16:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uint16)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						y := basic.Uint16(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint32:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uint32)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						y := basic.Uint32(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint64:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uint64)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						y := basic.Uint64(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				case reflect.Uintptr:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uintptr)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						y := basic.Uintptr(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				}
			} else if ky == kindConst {
				switch ykind {
				case reflect.Int:
					var y int
					if yIsBasic {
						y = vy.(int)
					} else {
						y = basic.Int(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(int32)
						fr.setReg(ir, x>>y)
					}
				case reflect.Int8:
					var y int8
					if yIsBasic {
						y = vy.(int8)
					} else {
						y = basic.Int8(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(int32)
						fr.setReg(ir, x>>y)
					}
				case reflect.Int16:
					var y int16
					if yIsBasic {
						y = vy.(int16)
					} else {
						y = basic.Int16(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(int32)
						fr.setReg(ir, x>>y)
					}
				case reflect.Int32:
					var y int32
					if yIsBasic {
						y = vy.(int32)
					} else {
						y = basic.Int32(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(int32)
						fr.setReg(ir, x>>y)
					}
				case reflect.Int64:
					var y int64
					if yIsBasic {
						y = vy.(int64)
					} else {
						y = basic.Int64(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(int32)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint:
					var y uint
					if yIsBasic {
						y = vy.(uint)
					} else {
						y = basic.Uint(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(int32)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint8:
					var y uint8
					if yIsBasic {
						y = vy.(uint8)
					} else {
						y = basic.Uint8(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(int32)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint16:
					var y uint16
					if yIsBasic {
						y = vy.(uint16)
					} else {
						y = basic.Uint16(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(int32)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint32:
					var y uint32
					if yIsBasic {
						y = vy.(uint32)
					} else {
						y = basic.Uint32(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(int32)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint64:
					var y uint64
					if yIsBasic {
						y = vy.(uint64)
					} else {
						y = basic.Uint64(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(int32)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uintptr:
					var y uintptr
					if yIsBasic {
						y = vy.(uintptr)
					} else {
						y = basic.Uintptr(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(int32)
						fr.setReg(ir, x>>y)
					}
				}
			} else {
				switch ykind {
				case reflect.Int:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(int32)
							y := fr.reg(iy).(int)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(int32)
						y := basic.Int(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				case reflect.Int8:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(int32)
							y := fr.reg(iy).(int8)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(int32)
						y := basic.Int8(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				case reflect.Int16:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(int32)
							y := fr.reg(iy).(int16)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(int32)
						y := basic.Int16(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				case reflect.Int32:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(int32)
							y := fr.reg(iy).(int32)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(int32)
						y := basic.Int32(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				case reflect.Int64:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(int32)
							y := fr.reg(iy).(int64)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(int32)
						y := basic.Int64(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(int32)
							y := fr.reg(iy).(uint)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(int32)
						y := basic.Uint(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint8:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(int32)
							y := fr.reg(iy).(uint8)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(int32)
						y := basic.Uint8(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint16:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(int32)
							y := fr.reg(iy).(uint16)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(int32)
						y := basic.Uint16(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint32:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(int32)
							y := fr.reg(iy).(uint32)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(int32)
						y := basic.Uint32(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint64:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(int32)
							y := fr.reg(iy).(uint64)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(int32)
						y := basic.Uint64(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				case reflect.Uintptr:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(int32)
							y := fr.reg(iy).(uintptr)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(int32)
						y := basic.Uintptr(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				}
			}
		case reflect.Int64:
			if kx == kindConst && ky == kindConst {
				x := vx.(int64)
				switch ykind {
				case reflect.Int:
					var y int
					if yIsBasic {
						y = vy.(int)
					} else {
						y = basic.Int(vy)
					}
					v := x >> y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Int8:
					var y int8
					if yIsBasic {
						y = vy.(int8)
					} else {
						y = basic.Int8(vy)
					}
					v := x >> y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Int16:
					var y int16
					if yIsBasic {
						y = vy.(int16)
					} else {
						y = basic.Int16(vy)
					}
					v := x >> y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Int32:
					var y int32
					if yIsBasic {
						y = vy.(int32)
					} else {
						y = basic.Int32(vy)
					}
					v := x >> y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Int64:
					var y int64
					if yIsBasic {
						y = vy.(int64)
					} else {
						y = basic.Int64(vy)
					}
					v := x >> y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uint:
					var y uint
					if yIsBasic {
						y = vy.(uint)
					} else {
						y = basic.Uint(vy)
					}
					v := x >> y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uint8:
					var y uint8
					if yIsBasic {
						y = vy.(uint8)
					} else {
						y = basic.Uint8(vy)
					}
					v := x >> y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uint16:
					var y uint16
					if yIsBasic {
						y = vy.(uint16)
					} else {
						y = basic.Uint16(vy)
					}
					v := x >> y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uint32:
					var y uint32
					if yIsBasic {
						y = vy.(uint32)
					} else {
						y = basic.Uint32(vy)
					}
					v := x >> y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uint64:
					var y uint64
					if yIsBasic {
						y = vy.(uint64)
					} else {
						y = basic.Uint64(vy)
					}
					v := x >> y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uintptr:
					var y uintptr
					if yIsBasic {
						y = vy.(uintptr)
					} else {
						y = basic.Uintptr(vy)
					}
					v := x >> y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				}
			} else if kx == kindConst {
				x := vx.(int64)
				switch ykind {
				case reflect.Int:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(int)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						y := basic.Int(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				case reflect.Int8:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(int8)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						y := basic.Int8(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				case reflect.Int16:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(int16)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						y := basic.Int16(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				case reflect.Int32:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(int32)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						y := basic.Int32(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				case reflect.Int64:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(int64)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						y := basic.Int64(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uint)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						y := basic.Uint(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint8:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uint8)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						y := basic.Uint8(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint16:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uint16)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						y := basic.Uint16(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint32:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uint32)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						y := basic.Uint32(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint64:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uint64)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						y := basic.Uint64(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				case reflect.Uintptr:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uintptr)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						y := basic.Uintptr(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				}
			} else if ky == kindConst {
				switch ykind {
				case reflect.Int:
					var y int
					if yIsBasic {
						y = vy.(int)
					} else {
						y = basic.Int(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(int64)
						fr.setReg(ir, x>>y)
					}
				case reflect.Int8:
					var y int8
					if yIsBasic {
						y = vy.(int8)
					} else {
						y = basic.Int8(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(int64)
						fr.setReg(ir, x>>y)
					}
				case reflect.Int16:
					var y int16
					if yIsBasic {
						y = vy.(int16)
					} else {
						y = basic.Int16(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(int64)
						fr.setReg(ir, x>>y)
					}
				case reflect.Int32:
					var y int32
					if yIsBasic {
						y = vy.(int32)
					} else {
						y = basic.Int32(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(int64)
						fr.setReg(ir, x>>y)
					}
				case reflect.Int64:
					var y int64
					if yIsBasic {
						y = vy.(int64)
					} else {
						y = basic.Int64(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(int64)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint:
					var y uint
					if yIsBasic {
						y = vy.(uint)
					} else {
						y = basic.Uint(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(int64)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint8:
					var y uint8
					if yIsBasic {
						y = vy.(uint8)
					} else {
						y = basic.Uint8(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(int64)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint16:
					var y uint16
					if yIsBasic {
						y = vy.(uint16)
					} else {
						y = basic.Uint16(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(int64)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint32:
					var y uint32
					if yIsBasic {
						y = vy.(uint32)
					} else {
						y = basic.Uint32(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(int64)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint64:
					var y uint64
					if yIsBasic {
						y = vy.(uint64)
					} else {
						y = basic.Uint64(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(int64)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uintptr:
					var y uintptr
					if yIsBasic {
						y = vy.(uintptr)
					} else {
						y = basic.Uintptr(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(int64)
						fr.setReg(ir, x>>y)
					}
				}
			} else {
				switch ykind {
				case reflect.Int:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(int64)
							y := fr.reg(iy).(int)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(int64)
						y := basic.Int(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				case reflect.Int8:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(int64)
							y := fr.reg(iy).(int8)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(int64)
						y := basic.Int8(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				case reflect.Int16:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(int64)
							y := fr.reg(iy).(int16)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(int64)
						y := basic.Int16(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				case reflect.Int32:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(int64)
							y := fr.reg(iy).(int32)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(int64)
						y := basic.Int32(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				case reflect.Int64:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(int64)
							y := fr.reg(iy).(int64)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(int64)
						y := basic.Int64(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(int64)
							y := fr.reg(iy).(uint)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(int64)
						y := basic.Uint(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint8:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(int64)
							y := fr.reg(iy).(uint8)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(int64)
						y := basic.Uint8(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint16:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(int64)
							y := fr.reg(iy).(uint16)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(int64)
						y := basic.Uint16(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint32:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(int64)
							y := fr.reg(iy).(uint32)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(int64)
						y := basic.Uint32(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint64:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(int64)
							y := fr.reg(iy).(uint64)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(int64)
						y := basic.Uint64(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				case reflect.Uintptr:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(int64)
							y := fr.reg(iy).(uintptr)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(int64)
						y := basic.Uintptr(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				}
			}
		case reflect.Uint:
			if kx == kindConst && ky == kindConst {
				x := vx.(uint)
				switch ykind {
				case reflect.Int:
					var y int
					if yIsBasic {
						y = vy.(int)
					} else {
						y = basic.Int(vy)
					}
					v := x >> y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Int8:
					var y int8
					if yIsBasic {
						y = vy.(int8)
					} else {
						y = basic.Int8(vy)
					}
					v := x >> y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Int16:
					var y int16
					if yIsBasic {
						y = vy.(int16)
					} else {
						y = basic.Int16(vy)
					}
					v := x >> y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Int32:
					var y int32
					if yIsBasic {
						y = vy.(int32)
					} else {
						y = basic.Int32(vy)
					}
					v := x >> y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Int64:
					var y int64
					if yIsBasic {
						y = vy.(int64)
					} else {
						y = basic.Int64(vy)
					}
					v := x >> y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uint:
					var y uint
					if yIsBasic {
						y = vy.(uint)
					} else {
						y = basic.Uint(vy)
					}
					v := x >> y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uint8:
					var y uint8
					if yIsBasic {
						y = vy.(uint8)
					} else {
						y = basic.Uint8(vy)
					}
					v := x >> y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uint16:
					var y uint16
					if yIsBasic {
						y = vy.(uint16)
					} else {
						y = basic.Uint16(vy)
					}
					v := x >> y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uint32:
					var y uint32
					if yIsBasic {
						y = vy.(uint32)
					} else {
						y = basic.Uint32(vy)
					}
					v := x >> y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uint64:
					var y uint64
					if yIsBasic {
						y = vy.(uint64)
					} else {
						y = basic.Uint64(vy)
					}
					v := x >> y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uintptr:
					var y uintptr
					if yIsBasic {
						y = vy.(uintptr)
					} else {
						y = basic.Uintptr(vy)
					}
					v := x >> y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				}
			} else if kx == kindConst {
				x := vx.(uint)
				switch ykind {
				case reflect.Int:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(int)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						y := basic.Int(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				case reflect.Int8:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(int8)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						y := basic.Int8(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				case reflect.Int16:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(int16)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						y := basic.Int16(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				case reflect.Int32:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(int32)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						y := basic.Int32(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				case reflect.Int64:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(int64)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						y := basic.Int64(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uint)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						y := basic.Uint(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint8:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uint8)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						y := basic.Uint8(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint16:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uint16)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						y := basic.Uint16(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint32:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uint32)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						y := basic.Uint32(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint64:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uint64)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						y := basic.Uint64(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				case reflect.Uintptr:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uintptr)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						y := basic.Uintptr(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				}
			} else if ky == kindConst {
				switch ykind {
				case reflect.Int:
					var y int
					if yIsBasic {
						y = vy.(int)
					} else {
						y = basic.Int(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uint)
						fr.setReg(ir, x>>y)
					}
				case reflect.Int8:
					var y int8
					if yIsBasic {
						y = vy.(int8)
					} else {
						y = basic.Int8(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uint)
						fr.setReg(ir, x>>y)
					}
				case reflect.Int16:
					var y int16
					if yIsBasic {
						y = vy.(int16)
					} else {
						y = basic.Int16(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uint)
						fr.setReg(ir, x>>y)
					}
				case reflect.Int32:
					var y int32
					if yIsBasic {
						y = vy.(int32)
					} else {
						y = basic.Int32(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uint)
						fr.setReg(ir, x>>y)
					}
				case reflect.Int64:
					var y int64
					if yIsBasic {
						y = vy.(int64)
					} else {
						y = basic.Int64(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uint)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint:
					var y uint
					if yIsBasic {
						y = vy.(uint)
					} else {
						y = basic.Uint(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uint)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint8:
					var y uint8
					if yIsBasic {
						y = vy.(uint8)
					} else {
						y = basic.Uint8(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uint)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint16:
					var y uint16
					if yIsBasic {
						y = vy.(uint16)
					} else {
						y = basic.Uint16(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uint)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint32:
					var y uint32
					if yIsBasic {
						y = vy.(uint32)
					} else {
						y = basic.Uint32(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uint)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint64:
					var y uint64
					if yIsBasic {
						y = vy.(uint64)
					} else {
						y = basic.Uint64(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uint)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uintptr:
					var y uintptr
					if yIsBasic {
						y = vy.(uintptr)
					} else {
						y = basic.Uintptr(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uint)
						fr.setReg(ir, x>>y)
					}
				}
			} else {
				switch ykind {
				case reflect.Int:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(uint)
							y := fr.reg(iy).(int)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uint)
						y := basic.Int(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				case reflect.Int8:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(uint)
							y := fr.reg(iy).(int8)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uint)
						y := basic.Int8(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				case reflect.Int16:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(uint)
							y := fr.reg(iy).(int16)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uint)
						y := basic.Int16(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				case reflect.Int32:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(uint)
							y := fr.reg(iy).(int32)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uint)
						y := basic.Int32(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				case reflect.Int64:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(uint)
							y := fr.reg(iy).(int64)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uint)
						y := basic.Int64(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(uint)
							y := fr.reg(iy).(uint)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uint)
						y := basic.Uint(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint8:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(uint)
							y := fr.reg(iy).(uint8)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uint)
						y := basic.Uint8(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint16:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(uint)
							y := fr.reg(iy).(uint16)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uint)
						y := basic.Uint16(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint32:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(uint)
							y := fr.reg(iy).(uint32)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uint)
						y := basic.Uint32(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint64:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(uint)
							y := fr.reg(iy).(uint64)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uint)
						y := basic.Uint64(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				case reflect.Uintptr:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(uint)
							y := fr.reg(iy).(uintptr)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uint)
						y := basic.Uintptr(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				}
			}
		case reflect.Uint8:
			if kx == kindConst && ky == kindConst {
				x := vx.(uint8)
				switch ykind {
				case reflect.Int:
					var y int
					if yIsBasic {
						y = vy.(int)
					} else {
						y = basic.Int(vy)
					}
					v := x >> y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Int8:
					var y int8
					if yIsBasic {
						y = vy.(int8)
					} else {
						y = basic.Int8(vy)
					}
					v := x >> y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Int16:
					var y int16
					if yIsBasic {
						y = vy.(int16)
					} else {
						y = basic.Int16(vy)
					}
					v := x >> y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Int32:
					var y int32
					if yIsBasic {
						y = vy.(int32)
					} else {
						y = basic.Int32(vy)
					}
					v := x >> y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Int64:
					var y int64
					if yIsBasic {
						y = vy.(int64)
					} else {
						y = basic.Int64(vy)
					}
					v := x >> y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uint:
					var y uint
					if yIsBasic {
						y = vy.(uint)
					} else {
						y = basic.Uint(vy)
					}
					v := x >> y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uint8:
					var y uint8
					if yIsBasic {
						y = vy.(uint8)
					} else {
						y = basic.Uint8(vy)
					}
					v := x >> y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uint16:
					var y uint16
					if yIsBasic {
						y = vy.(uint16)
					} else {
						y = basic.Uint16(vy)
					}
					v := x >> y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uint32:
					var y uint32
					if yIsBasic {
						y = vy.(uint32)
					} else {
						y = basic.Uint32(vy)
					}
					v := x >> y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uint64:
					var y uint64
					if yIsBasic {
						y = vy.(uint64)
					} else {
						y = basic.Uint64(vy)
					}
					v := x >> y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uintptr:
					var y uintptr
					if yIsBasic {
						y = vy.(uintptr)
					} else {
						y = basic.Uintptr(vy)
					}
					v := x >> y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				}
			} else if kx == kindConst {
				x := vx.(uint8)
				switch ykind {
				case reflect.Int:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(int)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						y := basic.Int(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				case reflect.Int8:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(int8)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						y := basic.Int8(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				case reflect.Int16:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(int16)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						y := basic.Int16(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				case reflect.Int32:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(int32)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						y := basic.Int32(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				case reflect.Int64:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(int64)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						y := basic.Int64(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uint)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						y := basic.Uint(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint8:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uint8)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						y := basic.Uint8(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint16:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uint16)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						y := basic.Uint16(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint32:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uint32)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						y := basic.Uint32(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint64:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uint64)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						y := basic.Uint64(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				case reflect.Uintptr:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uintptr)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						y := basic.Uintptr(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				}
			} else if ky == kindConst {
				switch ykind {
				case reflect.Int:
					var y int
					if yIsBasic {
						y = vy.(int)
					} else {
						y = basic.Int(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uint8)
						fr.setReg(ir, x>>y)
					}
				case reflect.Int8:
					var y int8
					if yIsBasic {
						y = vy.(int8)
					} else {
						y = basic.Int8(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uint8)
						fr.setReg(ir, x>>y)
					}
				case reflect.Int16:
					var y int16
					if yIsBasic {
						y = vy.(int16)
					} else {
						y = basic.Int16(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uint8)
						fr.setReg(ir, x>>y)
					}
				case reflect.Int32:
					var y int32
					if yIsBasic {
						y = vy.(int32)
					} else {
						y = basic.Int32(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uint8)
						fr.setReg(ir, x>>y)
					}
				case reflect.Int64:
					var y int64
					if yIsBasic {
						y = vy.(int64)
					} else {
						y = basic.Int64(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uint8)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint:
					var y uint
					if yIsBasic {
						y = vy.(uint)
					} else {
						y = basic.Uint(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uint8)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint8:
					var y uint8
					if yIsBasic {
						y = vy.(uint8)
					} else {
						y = basic.Uint8(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uint8)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint16:
					var y uint16
					if yIsBasic {
						y = vy.(uint16)
					} else {
						y = basic.Uint16(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uint8)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint32:
					var y uint32
					if yIsBasic {
						y = vy.(uint32)
					} else {
						y = basic.Uint32(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uint8)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint64:
					var y uint64
					if yIsBasic {
						y = vy.(uint64)
					} else {
						y = basic.Uint64(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uint8)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uintptr:
					var y uintptr
					if yIsBasic {
						y = vy.(uintptr)
					} else {
						y = basic.Uintptr(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uint8)
						fr.setReg(ir, x>>y)
					}
				}
			} else {
				switch ykind {
				case reflect.Int:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(uint8)
							y := fr.reg(iy).(int)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uint8)
						y := basic.Int(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				case reflect.Int8:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(uint8)
							y := fr.reg(iy).(int8)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uint8)
						y := basic.Int8(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				case reflect.Int16:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(uint8)
							y := fr.reg(iy).(int16)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uint8)
						y := basic.Int16(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				case reflect.Int32:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(uint8)
							y := fr.reg(iy).(int32)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uint8)
						y := basic.Int32(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				case reflect.Int64:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(uint8)
							y := fr.reg(iy).(int64)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uint8)
						y := basic.Int64(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(uint8)
							y := fr.reg(iy).(uint)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uint8)
						y := basic.Uint(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint8:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(uint8)
							y := fr.reg(iy).(uint8)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uint8)
						y := basic.Uint8(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint16:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(uint8)
							y := fr.reg(iy).(uint16)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uint8)
						y := basic.Uint16(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint32:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(uint8)
							y := fr.reg(iy).(uint32)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uint8)
						y := basic.Uint32(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint64:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(uint8)
							y := fr.reg(iy).(uint64)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uint8)
						y := basic.Uint64(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				case reflect.Uintptr:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(uint8)
							y := fr.reg(iy).(uintptr)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uint8)
						y := basic.Uintptr(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				}
			}
		case reflect.Uint16:
			if kx == kindConst && ky == kindConst {
				x := vx.(uint16)
				switch ykind {
				case reflect.Int:
					var y int
					if yIsBasic {
						y = vy.(int)
					} else {
						y = basic.Int(vy)
					}
					v := x >> y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Int8:
					var y int8
					if yIsBasic {
						y = vy.(int8)
					} else {
						y = basic.Int8(vy)
					}
					v := x >> y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Int16:
					var y int16
					if yIsBasic {
						y = vy.(int16)
					} else {
						y = basic.Int16(vy)
					}
					v := x >> y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Int32:
					var y int32
					if yIsBasic {
						y = vy.(int32)
					} else {
						y = basic.Int32(vy)
					}
					v := x >> y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Int64:
					var y int64
					if yIsBasic {
						y = vy.(int64)
					} else {
						y = basic.Int64(vy)
					}
					v := x >> y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uint:
					var y uint
					if yIsBasic {
						y = vy.(uint)
					} else {
						y = basic.Uint(vy)
					}
					v := x >> y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uint8:
					var y uint8
					if yIsBasic {
						y = vy.(uint8)
					} else {
						y = basic.Uint8(vy)
					}
					v := x >> y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uint16:
					var y uint16
					if yIsBasic {
						y = vy.(uint16)
					} else {
						y = basic.Uint16(vy)
					}
					v := x >> y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uint32:
					var y uint32
					if yIsBasic {
						y = vy.(uint32)
					} else {
						y = basic.Uint32(vy)
					}
					v := x >> y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uint64:
					var y uint64
					if yIsBasic {
						y = vy.(uint64)
					} else {
						y = basic.Uint64(vy)
					}
					v := x >> y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uintptr:
					var y uintptr
					if yIsBasic {
						y = vy.(uintptr)
					} else {
						y = basic.Uintptr(vy)
					}
					v := x >> y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				}
			} else if kx == kindConst {
				x := vx.(uint16)
				switch ykind {
				case reflect.Int:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(int)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						y := basic.Int(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				case reflect.Int8:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(int8)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						y := basic.Int8(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				case reflect.Int16:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(int16)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						y := basic.Int16(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				case reflect.Int32:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(int32)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						y := basic.Int32(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				case reflect.Int64:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(int64)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						y := basic.Int64(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uint)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						y := basic.Uint(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint8:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uint8)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						y := basic.Uint8(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint16:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uint16)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						y := basic.Uint16(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint32:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uint32)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						y := basic.Uint32(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint64:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uint64)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						y := basic.Uint64(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				case reflect.Uintptr:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uintptr)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						y := basic.Uintptr(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				}
			} else if ky == kindConst {
				switch ykind {
				case reflect.Int:
					var y int
					if yIsBasic {
						y = vy.(int)
					} else {
						y = basic.Int(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uint16)
						fr.setReg(ir, x>>y)
					}
				case reflect.Int8:
					var y int8
					if yIsBasic {
						y = vy.(int8)
					} else {
						y = basic.Int8(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uint16)
						fr.setReg(ir, x>>y)
					}
				case reflect.Int16:
					var y int16
					if yIsBasic {
						y = vy.(int16)
					} else {
						y = basic.Int16(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uint16)
						fr.setReg(ir, x>>y)
					}
				case reflect.Int32:
					var y int32
					if yIsBasic {
						y = vy.(int32)
					} else {
						y = basic.Int32(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uint16)
						fr.setReg(ir, x>>y)
					}
				case reflect.Int64:
					var y int64
					if yIsBasic {
						y = vy.(int64)
					} else {
						y = basic.Int64(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uint16)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint:
					var y uint
					if yIsBasic {
						y = vy.(uint)
					} else {
						y = basic.Uint(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uint16)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint8:
					var y uint8
					if yIsBasic {
						y = vy.(uint8)
					} else {
						y = basic.Uint8(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uint16)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint16:
					var y uint16
					if yIsBasic {
						y = vy.(uint16)
					} else {
						y = basic.Uint16(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uint16)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint32:
					var y uint32
					if yIsBasic {
						y = vy.(uint32)
					} else {
						y = basic.Uint32(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uint16)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint64:
					var y uint64
					if yIsBasic {
						y = vy.(uint64)
					} else {
						y = basic.Uint64(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uint16)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uintptr:
					var y uintptr
					if yIsBasic {
						y = vy.(uintptr)
					} else {
						y = basic.Uintptr(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uint16)
						fr.setReg(ir, x>>y)
					}
				}
			} else {
				switch ykind {
				case reflect.Int:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(uint16)
							y := fr.reg(iy).(int)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uint16)
						y := basic.Int(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				case reflect.Int8:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(uint16)
							y := fr.reg(iy).(int8)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uint16)
						y := basic.Int8(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				case reflect.Int16:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(uint16)
							y := fr.reg(iy).(int16)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uint16)
						y := basic.Int16(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				case reflect.Int32:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(uint16)
							y := fr.reg(iy).(int32)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uint16)
						y := basic.Int32(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				case reflect.Int64:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(uint16)
							y := fr.reg(iy).(int64)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uint16)
						y := basic.Int64(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(uint16)
							y := fr.reg(iy).(uint)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uint16)
						y := basic.Uint(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint8:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(uint16)
							y := fr.reg(iy).(uint8)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uint16)
						y := basic.Uint8(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint16:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(uint16)
							y := fr.reg(iy).(uint16)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uint16)
						y := basic.Uint16(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint32:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(uint16)
							y := fr.reg(iy).(uint32)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uint16)
						y := basic.Uint32(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint64:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(uint16)
							y := fr.reg(iy).(uint64)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uint16)
						y := basic.Uint64(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				case reflect.Uintptr:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(uint16)
							y := fr.reg(iy).(uintptr)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uint16)
						y := basic.Uintptr(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				}
			}
		case reflect.Uint32:
			if kx == kindConst && ky == kindConst {
				x := vx.(uint32)
				switch ykind {
				case reflect.Int:
					var y int
					if yIsBasic {
						y = vy.(int)
					} else {
						y = basic.Int(vy)
					}
					v := x >> y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Int8:
					var y int8
					if yIsBasic {
						y = vy.(int8)
					} else {
						y = basic.Int8(vy)
					}
					v := x >> y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Int16:
					var y int16
					if yIsBasic {
						y = vy.(int16)
					} else {
						y = basic.Int16(vy)
					}
					v := x >> y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Int32:
					var y int32
					if yIsBasic {
						y = vy.(int32)
					} else {
						y = basic.Int32(vy)
					}
					v := x >> y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Int64:
					var y int64
					if yIsBasic {
						y = vy.(int64)
					} else {
						y = basic.Int64(vy)
					}
					v := x >> y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uint:
					var y uint
					if yIsBasic {
						y = vy.(uint)
					} else {
						y = basic.Uint(vy)
					}
					v := x >> y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uint8:
					var y uint8
					if yIsBasic {
						y = vy.(uint8)
					} else {
						y = basic.Uint8(vy)
					}
					v := x >> y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uint16:
					var y uint16
					if yIsBasic {
						y = vy.(uint16)
					} else {
						y = basic.Uint16(vy)
					}
					v := x >> y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uint32:
					var y uint32
					if yIsBasic {
						y = vy.(uint32)
					} else {
						y = basic.Uint32(vy)
					}
					v := x >> y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uint64:
					var y uint64
					if yIsBasic {
						y = vy.(uint64)
					} else {
						y = basic.Uint64(vy)
					}
					v := x >> y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uintptr:
					var y uintptr
					if yIsBasic {
						y = vy.(uintptr)
					} else {
						y = basic.Uintptr(vy)
					}
					v := x >> y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				}
			} else if kx == kindConst {
				x := vx.(uint32)
				switch ykind {
				case reflect.Int:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(int)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						y := basic.Int(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				case reflect.Int8:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(int8)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						y := basic.Int8(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				case reflect.Int16:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(int16)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						y := basic.Int16(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				case reflect.Int32:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(int32)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						y := basic.Int32(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				case reflect.Int64:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(int64)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						y := basic.Int64(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uint)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						y := basic.Uint(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint8:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uint8)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						y := basic.Uint8(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint16:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uint16)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						y := basic.Uint16(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint32:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uint32)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						y := basic.Uint32(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint64:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uint64)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						y := basic.Uint64(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				case reflect.Uintptr:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uintptr)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						y := basic.Uintptr(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				}
			} else if ky == kindConst {
				switch ykind {
				case reflect.Int:
					var y int
					if yIsBasic {
						y = vy.(int)
					} else {
						y = basic.Int(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uint32)
						fr.setReg(ir, x>>y)
					}
				case reflect.Int8:
					var y int8
					if yIsBasic {
						y = vy.(int8)
					} else {
						y = basic.Int8(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uint32)
						fr.setReg(ir, x>>y)
					}
				case reflect.Int16:
					var y int16
					if yIsBasic {
						y = vy.(int16)
					} else {
						y = basic.Int16(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uint32)
						fr.setReg(ir, x>>y)
					}
				case reflect.Int32:
					var y int32
					if yIsBasic {
						y = vy.(int32)
					} else {
						y = basic.Int32(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uint32)
						fr.setReg(ir, x>>y)
					}
				case reflect.Int64:
					var y int64
					if yIsBasic {
						y = vy.(int64)
					} else {
						y = basic.Int64(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uint32)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint:
					var y uint
					if yIsBasic {
						y = vy.(uint)
					} else {
						y = basic.Uint(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uint32)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint8:
					var y uint8
					if yIsBasic {
						y = vy.(uint8)
					} else {
						y = basic.Uint8(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uint32)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint16:
					var y uint16
					if yIsBasic {
						y = vy.(uint16)
					} else {
						y = basic.Uint16(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uint32)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint32:
					var y uint32
					if yIsBasic {
						y = vy.(uint32)
					} else {
						y = basic.Uint32(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uint32)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint64:
					var y uint64
					if yIsBasic {
						y = vy.(uint64)
					} else {
						y = basic.Uint64(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uint32)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uintptr:
					var y uintptr
					if yIsBasic {
						y = vy.(uintptr)
					} else {
						y = basic.Uintptr(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uint32)
						fr.setReg(ir, x>>y)
					}
				}
			} else {
				switch ykind {
				case reflect.Int:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(uint32)
							y := fr.reg(iy).(int)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uint32)
						y := basic.Int(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				case reflect.Int8:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(uint32)
							y := fr.reg(iy).(int8)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uint32)
						y := basic.Int8(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				case reflect.Int16:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(uint32)
							y := fr.reg(iy).(int16)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uint32)
						y := basic.Int16(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				case reflect.Int32:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(uint32)
							y := fr.reg(iy).(int32)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uint32)
						y := basic.Int32(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				case reflect.Int64:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(uint32)
							y := fr.reg(iy).(int64)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uint32)
						y := basic.Int64(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(uint32)
							y := fr.reg(iy).(uint)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uint32)
						y := basic.Uint(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint8:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(uint32)
							y := fr.reg(iy).(uint8)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uint32)
						y := basic.Uint8(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint16:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(uint32)
							y := fr.reg(iy).(uint16)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uint32)
						y := basic.Uint16(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint32:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(uint32)
							y := fr.reg(iy).(uint32)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uint32)
						y := basic.Uint32(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint64:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(uint32)
							y := fr.reg(iy).(uint64)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uint32)
						y := basic.Uint64(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				case reflect.Uintptr:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(uint32)
							y := fr.reg(iy).(uintptr)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uint32)
						y := basic.Uintptr(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				}
			}
		case reflect.Uint64:
			if kx == kindConst && ky == kindConst {
				x := vx.(uint64)
				switch ykind {
				case reflect.Int:
					var y int
					if yIsBasic {
						y = vy.(int)
					} else {
						y = basic.Int(vy)
					}
					v := x >> y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Int8:
					var y int8
					if yIsBasic {
						y = vy.(int8)
					} else {
						y = basic.Int8(vy)
					}
					v := x >> y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Int16:
					var y int16
					if yIsBasic {
						y = vy.(int16)
					} else {
						y = basic.Int16(vy)
					}
					v := x >> y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Int32:
					var y int32
					if yIsBasic {
						y = vy.(int32)
					} else {
						y = basic.Int32(vy)
					}
					v := x >> y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Int64:
					var y int64
					if yIsBasic {
						y = vy.(int64)
					} else {
						y = basic.Int64(vy)
					}
					v := x >> y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uint:
					var y uint
					if yIsBasic {
						y = vy.(uint)
					} else {
						y = basic.Uint(vy)
					}
					v := x >> y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uint8:
					var y uint8
					if yIsBasic {
						y = vy.(uint8)
					} else {
						y = basic.Uint8(vy)
					}
					v := x >> y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uint16:
					var y uint16
					if yIsBasic {
						y = vy.(uint16)
					} else {
						y = basic.Uint16(vy)
					}
					v := x >> y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uint32:
					var y uint32
					if yIsBasic {
						y = vy.(uint32)
					} else {
						y = basic.Uint32(vy)
					}
					v := x >> y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uint64:
					var y uint64
					if yIsBasic {
						y = vy.(uint64)
					} else {
						y = basic.Uint64(vy)
					}
					v := x >> y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uintptr:
					var y uintptr
					if yIsBasic {
						y = vy.(uintptr)
					} else {
						y = basic.Uintptr(vy)
					}
					v := x >> y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				}
			} else if kx == kindConst {
				x := vx.(uint64)
				switch ykind {
				case reflect.Int:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(int)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						y := basic.Int(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				case reflect.Int8:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(int8)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						y := basic.Int8(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				case reflect.Int16:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(int16)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						y := basic.Int16(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				case reflect.Int32:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(int32)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						y := basic.Int32(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				case reflect.Int64:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(int64)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						y := basic.Int64(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uint)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						y := basic.Uint(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint8:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uint8)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						y := basic.Uint8(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint16:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uint16)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						y := basic.Uint16(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint32:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uint32)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						y := basic.Uint32(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint64:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uint64)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						y := basic.Uint64(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				case reflect.Uintptr:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uintptr)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						y := basic.Uintptr(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				}
			} else if ky == kindConst {
				switch ykind {
				case reflect.Int:
					var y int
					if yIsBasic {
						y = vy.(int)
					} else {
						y = basic.Int(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uint64)
						fr.setReg(ir, x>>y)
					}
				case reflect.Int8:
					var y int8
					if yIsBasic {
						y = vy.(int8)
					} else {
						y = basic.Int8(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uint64)
						fr.setReg(ir, x>>y)
					}
				case reflect.Int16:
					var y int16
					if yIsBasic {
						y = vy.(int16)
					} else {
						y = basic.Int16(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uint64)
						fr.setReg(ir, x>>y)
					}
				case reflect.Int32:
					var y int32
					if yIsBasic {
						y = vy.(int32)
					} else {
						y = basic.Int32(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uint64)
						fr.setReg(ir, x>>y)
					}
				case reflect.Int64:
					var y int64
					if yIsBasic {
						y = vy.(int64)
					} else {
						y = basic.Int64(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uint64)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint:
					var y uint
					if yIsBasic {
						y = vy.(uint)
					} else {
						y = basic.Uint(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uint64)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint8:
					var y uint8
					if yIsBasic {
						y = vy.(uint8)
					} else {
						y = basic.Uint8(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uint64)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint16:
					var y uint16
					if yIsBasic {
						y = vy.(uint16)
					} else {
						y = basic.Uint16(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uint64)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint32:
					var y uint32
					if yIsBasic {
						y = vy.(uint32)
					} else {
						y = basic.Uint32(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uint64)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint64:
					var y uint64
					if yIsBasic {
						y = vy.(uint64)
					} else {
						y = basic.Uint64(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uint64)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uintptr:
					var y uintptr
					if yIsBasic {
						y = vy.(uintptr)
					} else {
						y = basic.Uintptr(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uint64)
						fr.setReg(ir, x>>y)
					}
				}
			} else {
				switch ykind {
				case reflect.Int:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(uint64)
							y := fr.reg(iy).(int)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uint64)
						y := basic.Int(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				case reflect.Int8:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(uint64)
							y := fr.reg(iy).(int8)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uint64)
						y := basic.Int8(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				case reflect.Int16:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(uint64)
							y := fr.reg(iy).(int16)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uint64)
						y := basic.Int16(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				case reflect.Int32:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(uint64)
							y := fr.reg(iy).(int32)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uint64)
						y := basic.Int32(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				case reflect.Int64:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(uint64)
							y := fr.reg(iy).(int64)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uint64)
						y := basic.Int64(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(uint64)
							y := fr.reg(iy).(uint)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uint64)
						y := basic.Uint(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint8:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(uint64)
							y := fr.reg(iy).(uint8)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uint64)
						y := basic.Uint8(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint16:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(uint64)
							y := fr.reg(iy).(uint16)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uint64)
						y := basic.Uint16(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint32:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(uint64)
							y := fr.reg(iy).(uint32)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uint64)
						y := basic.Uint32(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint64:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(uint64)
							y := fr.reg(iy).(uint64)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uint64)
						y := basic.Uint64(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				case reflect.Uintptr:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(uint64)
							y := fr.reg(iy).(uintptr)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uint64)
						y := basic.Uintptr(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				}
			}
		case reflect.Uintptr:
			if kx == kindConst && ky == kindConst {
				x := vx.(uintptr)
				switch ykind {
				case reflect.Int:
					var y int
					if yIsBasic {
						y = vy.(int)
					} else {
						y = basic.Int(vy)
					}
					v := x >> y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Int8:
					var y int8
					if yIsBasic {
						y = vy.(int8)
					} else {
						y = basic.Int8(vy)
					}
					v := x >> y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Int16:
					var y int16
					if yIsBasic {
						y = vy.(int16)
					} else {
						y = basic.Int16(vy)
					}
					v := x >> y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Int32:
					var y int32
					if yIsBasic {
						y = vy.(int32)
					} else {
						y = basic.Int32(vy)
					}
					v := x >> y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Int64:
					var y int64
					if yIsBasic {
						y = vy.(int64)
					} else {
						y = basic.Int64(vy)
					}
					v := x >> y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uint:
					var y uint
					if yIsBasic {
						y = vy.(uint)
					} else {
						y = basic.Uint(vy)
					}
					v := x >> y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uint8:
					var y uint8
					if yIsBasic {
						y = vy.(uint8)
					} else {
						y = basic.Uint8(vy)
					}
					v := x >> y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uint16:
					var y uint16
					if yIsBasic {
						y = vy.(uint16)
					} else {
						y = basic.Uint16(vy)
					}
					v := x >> y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uint32:
					var y uint32
					if yIsBasic {
						y = vy.(uint32)
					} else {
						y = basic.Uint32(vy)
					}
					v := x >> y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uint64:
					var y uint64
					if yIsBasic {
						y = vy.(uint64)
					} else {
						y = basic.Uint64(vy)
					}
					v := x >> y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uintptr:
					var y uintptr
					if yIsBasic {
						y = vy.(uintptr)
					} else {
						y = basic.Uintptr(vy)
					}
					v := x >> y
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				}
			} else if kx == kindConst {
				x := vx.(uintptr)
				switch ykind {
				case reflect.Int:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(int)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						y := basic.Int(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				case reflect.Int8:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(int8)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						y := basic.Int8(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				case reflect.Int16:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(int16)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						y := basic.Int16(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				case reflect.Int32:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(int32)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						y := basic.Int32(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				case reflect.Int64:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(int64)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						y := basic.Int64(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uint)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						y := basic.Uint(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint8:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uint8)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						y := basic.Uint8(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint16:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uint16)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						y := basic.Uint16(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint32:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uint32)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						y := basic.Uint32(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint64:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uint64)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						y := basic.Uint64(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				case reflect.Uintptr:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uintptr)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						y := basic.Uintptr(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				}
			} else if ky == kindConst {
				switch ykind {
				case reflect.Int:
					var y int
					if yIsBasic {
						y = vy.(int)
					} else {
						y = basic.Int(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uintptr)
						fr.setReg(ir, x>>y)
					}
				case reflect.Int8:
					var y int8
					if yIsBasic {
						y = vy.(int8)
					} else {
						y = basic.Int8(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uintptr)
						fr.setReg(ir, x>>y)
					}
				case reflect.Int16:
					var y int16
					if yIsBasic {
						y = vy.(int16)
					} else {
						y = basic.Int16(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uintptr)
						fr.setReg(ir, x>>y)
					}
				case reflect.Int32:
					var y int32
					if yIsBasic {
						y = vy.(int32)
					} else {
						y = basic.Int32(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uintptr)
						fr.setReg(ir, x>>y)
					}
				case reflect.Int64:
					var y int64
					if yIsBasic {
						y = vy.(int64)
					} else {
						y = basic.Int64(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uintptr)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint:
					var y uint
					if yIsBasic {
						y = vy.(uint)
					} else {
						y = basic.Uint(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uintptr)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint8:
					var y uint8
					if yIsBasic {
						y = vy.(uint8)
					} else {
						y = basic.Uint8(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uintptr)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint16:
					var y uint16
					if yIsBasic {
						y = vy.(uint16)
					} else {
						y = basic.Uint16(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uintptr)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint32:
					var y uint32
					if yIsBasic {
						y = vy.(uint32)
					} else {
						y = basic.Uint32(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uintptr)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint64:
					var y uint64
					if yIsBasic {
						y = vy.(uint64)
					} else {
						y = basic.Uint64(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uintptr)
						fr.setReg(ir, x>>y)
					}
				case reflect.Uintptr:
					var y uintptr
					if yIsBasic {
						y = vy.(uintptr)
					} else {
						y = basic.Uintptr(vy)
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uintptr)
						fr.setReg(ir, x>>y)
					}
				}
			} else {
				switch ykind {
				case reflect.Int:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(uintptr)
							y := fr.reg(iy).(int)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uintptr)
						y := basic.Int(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				case reflect.Int8:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(uintptr)
							y := fr.reg(iy).(int8)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uintptr)
						y := basic.Int8(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				case reflect.Int16:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(uintptr)
							y := fr.reg(iy).(int16)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uintptr)
						y := basic.Int16(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				case reflect.Int32:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(uintptr)
							y := fr.reg(iy).(int32)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uintptr)
						y := basic.Int32(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				case reflect.Int64:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(uintptr)
							y := fr.reg(iy).(int64)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uintptr)
						y := basic.Int64(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(uintptr)
							y := fr.reg(iy).(uint)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uintptr)
						y := basic.Uint(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint8:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(uintptr)
							y := fr.reg(iy).(uint8)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uintptr)
						y := basic.Uint8(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint16:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(uintptr)
							y := fr.reg(iy).(uint16)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uintptr)
						y := basic.Uint16(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint32:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(uintptr)
							y := fr.reg(iy).(uint32)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uintptr)
						y := basic.Uint32(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				case reflect.Uint64:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(uintptr)
							y := fr.reg(iy).(uint64)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uintptr)
						y := basic.Uint64(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				case reflect.Uintptr:
					if yIsBasic {
						return func(fr *frame) {
							x := fr.reg(ix).(uintptr)
							y := fr.reg(iy).(uintptr)
							fr.setReg(ir, x>>y)
						}
					}
					return func(fr *frame) {
						x := fr.reg(ix).(uintptr)
						y := basic.Uintptr(fr.reg(iy))
						fr.setReg(ir, x>>y)
					}
				}
			}
		}
	} else {
		t := basic.TypeOfType(xtyp)
		switch xkind {
		case reflect.Int:
			if kx == kindConst && ky == kindConst {
				x := basic.Int(vx)
				switch ykind {
				case reflect.Int:
					var y int
					if yIsBasic {
						y = vy.(int)
					} else {
						y = basic.Int(vy)
					}
					v := basic.Make(t, x>>y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Int8:
					var y int8
					if yIsBasic {
						y = vy.(int8)
					} else {
						y = basic.Int8(vy)
					}
					v := basic.Make(t, x>>y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Int16:
					var y int16
					if yIsBasic {
						y = vy.(int16)
					} else {
						y = basic.Int16(vy)
					}
					v := basic.Make(t, x>>y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Int32:
					var y int32
					if yIsBasic {
						y = vy.(int32)
					} else {
						y = basic.Int32(vy)
					}
					v := basic.Make(t, x>>y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Int64:
					var y int64
					if yIsBasic {
						y = vy.(int64)
					} else {
						y = basic.Int64(vy)
					}
					v := basic.Make(t, x>>y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uint:
					var y uint
					if yIsBasic {
						y = vy.(uint)
					} else {
						y = basic.Uint(vy)
					}
					v := basic.Make(t, x>>y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uint8:
					var y uint8
					if yIsBasic {
						y = vy.(uint8)
					} else {
						y = basic.Uint8(vy)
					}
					v := basic.Make(t, x>>y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uint16:
					var y uint16
					if yIsBasic {
						y = vy.(uint16)
					} else {
						y = basic.Uint16(vy)
					}
					v := basic.Make(t, x>>y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uint32:
					var y uint32
					if yIsBasic {
						y = vy.(uint32)
					} else {
						y = basic.Uint32(vy)
					}
					v := basic.Make(t, x>>y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uint64:
					var y uint64
					if yIsBasic {
						y = vy.(uint64)
					} else {
						y = basic.Uint64(vy)
					}
					v := basic.Make(t, x>>y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uintptr:
					var y uintptr
					if yIsBasic {
						y = vy.(uintptr)
					} else {
						y = basic.Uintptr(vy)
					}
					v := basic.Make(t, x>>y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				}
			} else if kx == kindConst {
				x := basic.Int(vx)
				switch ykind {
				case reflect.Int:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(int)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						y := basic.Int(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int8:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(int8)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						y := basic.Int8(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int16:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(int16)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						y := basic.Int16(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int32:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(int32)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						y := basic.Int32(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int64:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(int64)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						y := basic.Int64(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uint)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						y := basic.Uint(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint8:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uint8)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						y := basic.Uint8(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint16:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uint16)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						y := basic.Uint16(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint32:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uint32)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						y := basic.Uint32(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint64:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uint64)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						y := basic.Uint64(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uintptr:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uintptr)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						y := basic.Uintptr(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				}
			} else if ky == kindConst {
				switch ykind {
				case reflect.Int:
					var y int
					if yIsBasic {
						y = vy.(int)
					} else {
						y = basic.Int(vy)
					}
					return func(fr *frame) {
						x := basic.Int(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int8:
					var y int8
					if yIsBasic {
						y = vy.(int8)
					} else {
						y = basic.Int8(vy)
					}
					return func(fr *frame) {
						x := basic.Int(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int16:
					var y int16
					if yIsBasic {
						y = vy.(int16)
					} else {
						y = basic.Int16(vy)
					}
					return func(fr *frame) {
						x := basic.Int(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int32:
					var y int32
					if yIsBasic {
						y = vy.(int32)
					} else {
						y = basic.Int32(vy)
					}
					return func(fr *frame) {
						x := basic.Int(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int64:
					var y int64
					if yIsBasic {
						y = vy.(int64)
					} else {
						y = basic.Int64(vy)
					}
					return func(fr *frame) {
						x := basic.Int(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint:
					var y uint
					if yIsBasic {
						y = vy.(uint)
					} else {
						y = basic.Uint(vy)
					}
					return func(fr *frame) {
						x := basic.Int(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint8:
					var y uint8
					if yIsBasic {
						y = vy.(uint8)
					} else {
						y = basic.Uint8(vy)
					}
					return func(fr *frame) {
						x := basic.Int(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint16:
					var y uint16
					if yIsBasic {
						y = vy.(uint16)
					} else {
						y = basic.Uint16(vy)
					}
					return func(fr *frame) {
						x := basic.Int(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint32:
					var y uint32
					if yIsBasic {
						y = vy.(uint32)
					} else {
						y = basic.Uint32(vy)
					}
					return func(fr *frame) {
						x := basic.Int(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint64:
					var y uint64
					if yIsBasic {
						y = vy.(uint64)
					} else {
						y = basic.Uint64(vy)
					}
					return func(fr *frame) {
						x := basic.Int(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uintptr:
					var y uintptr
					if yIsBasic {
						y = vy.(uintptr)
					} else {
						y = basic.Uintptr(vy)
					}
					return func(fr *frame) {
						x := basic.Int(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				}
			} else {
				switch ykind {
				case reflect.Int:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Int(fr.reg(ix))
							y := fr.reg(iy).(int)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						x := basic.Int(fr.reg(ix))
						y := basic.Int(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int8:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Int(fr.reg(ix))
							y := fr.reg(iy).(int8)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						x := basic.Int(fr.reg(ix))
						y := basic.Int8(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int16:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Int(fr.reg(ix))
							y := fr.reg(iy).(int16)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						x := basic.Int(fr.reg(ix))
						y := basic.Int16(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int32:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Int(fr.reg(ix))
							y := fr.reg(iy).(int32)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						x := basic.Int(fr.reg(ix))
						y := basic.Int32(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int64:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Int(fr.reg(ix))
							y := fr.reg(iy).(int64)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						x := basic.Int(fr.reg(ix))
						y := basic.Int64(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Int(fr.reg(ix))
							y := fr.reg(iy).(uint)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						x := basic.Int(fr.reg(ix))
						y := basic.Uint(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint8:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Int(fr.reg(ix))
							y := fr.reg(iy).(uint8)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						x := basic.Int(fr.reg(ix))
						y := basic.Uint8(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint16:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Int(fr.reg(ix))
							y := fr.reg(iy).(uint16)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						x := basic.Int(fr.reg(ix))
						y := basic.Uint16(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint32:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Int(fr.reg(ix))
							y := fr.reg(iy).(uint32)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						x := basic.Int(fr.reg(ix))
						y := basic.Uint32(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint64:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Int(fr.reg(ix))
							y := fr.reg(iy).(uint64)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						x := basic.Int(fr.reg(ix))
						y := basic.Uint64(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uintptr:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Int(fr.reg(ix))
							y := fr.reg(iy).(uintptr)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						x := basic.Int(fr.reg(ix))
						y := basic.Uintptr(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				}
			}
		case reflect.Int8:
			if kx == kindConst && ky == kindConst {
				x := basic.Int8(vx)
				switch ykind {
				case reflect.Int:
					var y int
					if yIsBasic {
						y = vy.(int)
					} else {
						y = basic.Int(vy)
					}
					v := basic.Make(t, x>>y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Int8:
					var y int8
					if yIsBasic {
						y = vy.(int8)
					} else {
						y = basic.Int8(vy)
					}
					v := basic.Make(t, x>>y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Int16:
					var y int16
					if yIsBasic {
						y = vy.(int16)
					} else {
						y = basic.Int16(vy)
					}
					v := basic.Make(t, x>>y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Int32:
					var y int32
					if yIsBasic {
						y = vy.(int32)
					} else {
						y = basic.Int32(vy)
					}
					v := basic.Make(t, x>>y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Int64:
					var y int64
					if yIsBasic {
						y = vy.(int64)
					} else {
						y = basic.Int64(vy)
					}
					v := basic.Make(t, x>>y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uint:
					var y uint
					if yIsBasic {
						y = vy.(uint)
					} else {
						y = basic.Uint(vy)
					}
					v := basic.Make(t, x>>y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uint8:
					var y uint8
					if yIsBasic {
						y = vy.(uint8)
					} else {
						y = basic.Uint8(vy)
					}
					v := basic.Make(t, x>>y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uint16:
					var y uint16
					if yIsBasic {
						y = vy.(uint16)
					} else {
						y = basic.Uint16(vy)
					}
					v := basic.Make(t, x>>y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uint32:
					var y uint32
					if yIsBasic {
						y = vy.(uint32)
					} else {
						y = basic.Uint32(vy)
					}
					v := basic.Make(t, x>>y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uint64:
					var y uint64
					if yIsBasic {
						y = vy.(uint64)
					} else {
						y = basic.Uint64(vy)
					}
					v := basic.Make(t, x>>y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uintptr:
					var y uintptr
					if yIsBasic {
						y = vy.(uintptr)
					} else {
						y = basic.Uintptr(vy)
					}
					v := basic.Make(t, x>>y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				}
			} else if kx == kindConst {
				x := basic.Int8(vx)
				switch ykind {
				case reflect.Int:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(int)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						y := basic.Int(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int8:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(int8)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						y := basic.Int8(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int16:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(int16)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						y := basic.Int16(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int32:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(int32)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						y := basic.Int32(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int64:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(int64)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						y := basic.Int64(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uint)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						y := basic.Uint(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint8:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uint8)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						y := basic.Uint8(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint16:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uint16)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						y := basic.Uint16(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint32:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uint32)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						y := basic.Uint32(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint64:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uint64)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						y := basic.Uint64(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uintptr:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uintptr)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						y := basic.Uintptr(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				}
			} else if ky == kindConst {
				switch ykind {
				case reflect.Int:
					var y int
					if yIsBasic {
						y = vy.(int)
					} else {
						y = basic.Int(vy)
					}
					return func(fr *frame) {
						x := basic.Int8(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int8:
					var y int8
					if yIsBasic {
						y = vy.(int8)
					} else {
						y = basic.Int8(vy)
					}
					return func(fr *frame) {
						x := basic.Int8(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int16:
					var y int16
					if yIsBasic {
						y = vy.(int16)
					} else {
						y = basic.Int16(vy)
					}
					return func(fr *frame) {
						x := basic.Int8(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int32:
					var y int32
					if yIsBasic {
						y = vy.(int32)
					} else {
						y = basic.Int32(vy)
					}
					return func(fr *frame) {
						x := basic.Int8(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int64:
					var y int64
					if yIsBasic {
						y = vy.(int64)
					} else {
						y = basic.Int64(vy)
					}
					return func(fr *frame) {
						x := basic.Int8(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint:
					var y uint
					if yIsBasic {
						y = vy.(uint)
					} else {
						y = basic.Uint(vy)
					}
					return func(fr *frame) {
						x := basic.Int8(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint8:
					var y uint8
					if yIsBasic {
						y = vy.(uint8)
					} else {
						y = basic.Uint8(vy)
					}
					return func(fr *frame) {
						x := basic.Int8(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint16:
					var y uint16
					if yIsBasic {
						y = vy.(uint16)
					} else {
						y = basic.Uint16(vy)
					}
					return func(fr *frame) {
						x := basic.Int8(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint32:
					var y uint32
					if yIsBasic {
						y = vy.(uint32)
					} else {
						y = basic.Uint32(vy)
					}
					return func(fr *frame) {
						x := basic.Int8(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint64:
					var y uint64
					if yIsBasic {
						y = vy.(uint64)
					} else {
						y = basic.Uint64(vy)
					}
					return func(fr *frame) {
						x := basic.Int8(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uintptr:
					var y uintptr
					if yIsBasic {
						y = vy.(uintptr)
					} else {
						y = basic.Uintptr(vy)
					}
					return func(fr *frame) {
						x := basic.Int8(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				}
			} else {
				switch ykind {
				case reflect.Int:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Int8(fr.reg(ix))
							y := fr.reg(iy).(int)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						x := basic.Int8(fr.reg(ix))
						y := basic.Int(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int8:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Int8(fr.reg(ix))
							y := fr.reg(iy).(int8)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						x := basic.Int8(fr.reg(ix))
						y := basic.Int8(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int16:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Int8(fr.reg(ix))
							y := fr.reg(iy).(int16)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						x := basic.Int8(fr.reg(ix))
						y := basic.Int16(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int32:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Int8(fr.reg(ix))
							y := fr.reg(iy).(int32)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						x := basic.Int8(fr.reg(ix))
						y := basic.Int32(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int64:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Int8(fr.reg(ix))
							y := fr.reg(iy).(int64)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						x := basic.Int8(fr.reg(ix))
						y := basic.Int64(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Int8(fr.reg(ix))
							y := fr.reg(iy).(uint)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						x := basic.Int8(fr.reg(ix))
						y := basic.Uint(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint8:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Int8(fr.reg(ix))
							y := fr.reg(iy).(uint8)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						x := basic.Int8(fr.reg(ix))
						y := basic.Uint8(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint16:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Int8(fr.reg(ix))
							y := fr.reg(iy).(uint16)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						x := basic.Int8(fr.reg(ix))
						y := basic.Uint16(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint32:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Int8(fr.reg(ix))
							y := fr.reg(iy).(uint32)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						x := basic.Int8(fr.reg(ix))
						y := basic.Uint32(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint64:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Int8(fr.reg(ix))
							y := fr.reg(iy).(uint64)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						x := basic.Int8(fr.reg(ix))
						y := basic.Uint64(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uintptr:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Int8(fr.reg(ix))
							y := fr.reg(iy).(uintptr)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						x := basic.Int8(fr.reg(ix))
						y := basic.Uintptr(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				}
			}
		case reflect.Int16:
			if kx == kindConst && ky == kindConst {
				x := basic.Int16(vx)
				switch ykind {
				case reflect.Int:
					var y int
					if yIsBasic {
						y = vy.(int)
					} else {
						y = basic.Int(vy)
					}
					v := basic.Make(t, x>>y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Int8:
					var y int8
					if yIsBasic {
						y = vy.(int8)
					} else {
						y = basic.Int8(vy)
					}
					v := basic.Make(t, x>>y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Int16:
					var y int16
					if yIsBasic {
						y = vy.(int16)
					} else {
						y = basic.Int16(vy)
					}
					v := basic.Make(t, x>>y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Int32:
					var y int32
					if yIsBasic {
						y = vy.(int32)
					} else {
						y = basic.Int32(vy)
					}
					v := basic.Make(t, x>>y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Int64:
					var y int64
					if yIsBasic {
						y = vy.(int64)
					} else {
						y = basic.Int64(vy)
					}
					v := basic.Make(t, x>>y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uint:
					var y uint
					if yIsBasic {
						y = vy.(uint)
					} else {
						y = basic.Uint(vy)
					}
					v := basic.Make(t, x>>y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uint8:
					var y uint8
					if yIsBasic {
						y = vy.(uint8)
					} else {
						y = basic.Uint8(vy)
					}
					v := basic.Make(t, x>>y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uint16:
					var y uint16
					if yIsBasic {
						y = vy.(uint16)
					} else {
						y = basic.Uint16(vy)
					}
					v := basic.Make(t, x>>y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uint32:
					var y uint32
					if yIsBasic {
						y = vy.(uint32)
					} else {
						y = basic.Uint32(vy)
					}
					v := basic.Make(t, x>>y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uint64:
					var y uint64
					if yIsBasic {
						y = vy.(uint64)
					} else {
						y = basic.Uint64(vy)
					}
					v := basic.Make(t, x>>y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uintptr:
					var y uintptr
					if yIsBasic {
						y = vy.(uintptr)
					} else {
						y = basic.Uintptr(vy)
					}
					v := basic.Make(t, x>>y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				}
			} else if kx == kindConst {
				x := basic.Int16(vx)
				switch ykind {
				case reflect.Int:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(int)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						y := basic.Int(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int8:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(int8)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						y := basic.Int8(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int16:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(int16)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						y := basic.Int16(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int32:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(int32)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						y := basic.Int32(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int64:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(int64)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						y := basic.Int64(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uint)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						y := basic.Uint(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint8:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uint8)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						y := basic.Uint8(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint16:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uint16)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						y := basic.Uint16(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint32:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uint32)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						y := basic.Uint32(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint64:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uint64)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						y := basic.Uint64(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uintptr:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uintptr)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						y := basic.Uintptr(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				}
			} else if ky == kindConst {
				switch ykind {
				case reflect.Int:
					var y int
					if yIsBasic {
						y = vy.(int)
					} else {
						y = basic.Int(vy)
					}
					return func(fr *frame) {
						x := basic.Int16(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int8:
					var y int8
					if yIsBasic {
						y = vy.(int8)
					} else {
						y = basic.Int8(vy)
					}
					return func(fr *frame) {
						x := basic.Int16(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int16:
					var y int16
					if yIsBasic {
						y = vy.(int16)
					} else {
						y = basic.Int16(vy)
					}
					return func(fr *frame) {
						x := basic.Int16(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int32:
					var y int32
					if yIsBasic {
						y = vy.(int32)
					} else {
						y = basic.Int32(vy)
					}
					return func(fr *frame) {
						x := basic.Int16(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int64:
					var y int64
					if yIsBasic {
						y = vy.(int64)
					} else {
						y = basic.Int64(vy)
					}
					return func(fr *frame) {
						x := basic.Int16(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint:
					var y uint
					if yIsBasic {
						y = vy.(uint)
					} else {
						y = basic.Uint(vy)
					}
					return func(fr *frame) {
						x := basic.Int16(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint8:
					var y uint8
					if yIsBasic {
						y = vy.(uint8)
					} else {
						y = basic.Uint8(vy)
					}
					return func(fr *frame) {
						x := basic.Int16(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint16:
					var y uint16
					if yIsBasic {
						y = vy.(uint16)
					} else {
						y = basic.Uint16(vy)
					}
					return func(fr *frame) {
						x := basic.Int16(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint32:
					var y uint32
					if yIsBasic {
						y = vy.(uint32)
					} else {
						y = basic.Uint32(vy)
					}
					return func(fr *frame) {
						x := basic.Int16(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint64:
					var y uint64
					if yIsBasic {
						y = vy.(uint64)
					} else {
						y = basic.Uint64(vy)
					}
					return func(fr *frame) {
						x := basic.Int16(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uintptr:
					var y uintptr
					if yIsBasic {
						y = vy.(uintptr)
					} else {
						y = basic.Uintptr(vy)
					}
					return func(fr *frame) {
						x := basic.Int16(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				}
			} else {
				switch ykind {
				case reflect.Int:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Int16(fr.reg(ix))
							y := fr.reg(iy).(int)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						x := basic.Int16(fr.reg(ix))
						y := basic.Int(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int8:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Int16(fr.reg(ix))
							y := fr.reg(iy).(int8)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						x := basic.Int16(fr.reg(ix))
						y := basic.Int8(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int16:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Int16(fr.reg(ix))
							y := fr.reg(iy).(int16)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						x := basic.Int16(fr.reg(ix))
						y := basic.Int16(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int32:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Int16(fr.reg(ix))
							y := fr.reg(iy).(int32)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						x := basic.Int16(fr.reg(ix))
						y := basic.Int32(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int64:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Int16(fr.reg(ix))
							y := fr.reg(iy).(int64)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						x := basic.Int16(fr.reg(ix))
						y := basic.Int64(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Int16(fr.reg(ix))
							y := fr.reg(iy).(uint)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						x := basic.Int16(fr.reg(ix))
						y := basic.Uint(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint8:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Int16(fr.reg(ix))
							y := fr.reg(iy).(uint8)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						x := basic.Int16(fr.reg(ix))
						y := basic.Uint8(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint16:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Int16(fr.reg(ix))
							y := fr.reg(iy).(uint16)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						x := basic.Int16(fr.reg(ix))
						y := basic.Uint16(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint32:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Int16(fr.reg(ix))
							y := fr.reg(iy).(uint32)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						x := basic.Int16(fr.reg(ix))
						y := basic.Uint32(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint64:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Int16(fr.reg(ix))
							y := fr.reg(iy).(uint64)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						x := basic.Int16(fr.reg(ix))
						y := basic.Uint64(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uintptr:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Int16(fr.reg(ix))
							y := fr.reg(iy).(uintptr)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						x := basic.Int16(fr.reg(ix))
						y := basic.Uintptr(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				}
			}
		case reflect.Int32:
			if kx == kindConst && ky == kindConst {
				x := basic.Int32(vx)
				switch ykind {
				case reflect.Int:
					var y int
					if yIsBasic {
						y = vy.(int)
					} else {
						y = basic.Int(vy)
					}
					v := basic.Make(t, x>>y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Int8:
					var y int8
					if yIsBasic {
						y = vy.(int8)
					} else {
						y = basic.Int8(vy)
					}
					v := basic.Make(t, x>>y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Int16:
					var y int16
					if yIsBasic {
						y = vy.(int16)
					} else {
						y = basic.Int16(vy)
					}
					v := basic.Make(t, x>>y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Int32:
					var y int32
					if yIsBasic {
						y = vy.(int32)
					} else {
						y = basic.Int32(vy)
					}
					v := basic.Make(t, x>>y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Int64:
					var y int64
					if yIsBasic {
						y = vy.(int64)
					} else {
						y = basic.Int64(vy)
					}
					v := basic.Make(t, x>>y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uint:
					var y uint
					if yIsBasic {
						y = vy.(uint)
					} else {
						y = basic.Uint(vy)
					}
					v := basic.Make(t, x>>y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uint8:
					var y uint8
					if yIsBasic {
						y = vy.(uint8)
					} else {
						y = basic.Uint8(vy)
					}
					v := basic.Make(t, x>>y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uint16:
					var y uint16
					if yIsBasic {
						y = vy.(uint16)
					} else {
						y = basic.Uint16(vy)
					}
					v := basic.Make(t, x>>y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uint32:
					var y uint32
					if yIsBasic {
						y = vy.(uint32)
					} else {
						y = basic.Uint32(vy)
					}
					v := basic.Make(t, x>>y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uint64:
					var y uint64
					if yIsBasic {
						y = vy.(uint64)
					} else {
						y = basic.Uint64(vy)
					}
					v := basic.Make(t, x>>y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uintptr:
					var y uintptr
					if yIsBasic {
						y = vy.(uintptr)
					} else {
						y = basic.Uintptr(vy)
					}
					v := basic.Make(t, x>>y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				}
			} else if kx == kindConst {
				x := basic.Int32(vx)
				switch ykind {
				case reflect.Int:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(int)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						y := basic.Int(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int8:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(int8)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						y := basic.Int8(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int16:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(int16)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						y := basic.Int16(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int32:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(int32)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						y := basic.Int32(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int64:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(int64)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						y := basic.Int64(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uint)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						y := basic.Uint(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint8:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uint8)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						y := basic.Uint8(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint16:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uint16)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						y := basic.Uint16(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint32:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uint32)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						y := basic.Uint32(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint64:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uint64)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						y := basic.Uint64(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uintptr:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uintptr)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						y := basic.Uintptr(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				}
			} else if ky == kindConst {
				switch ykind {
				case reflect.Int:
					var y int
					if yIsBasic {
						y = vy.(int)
					} else {
						y = basic.Int(vy)
					}
					return func(fr *frame) {
						x := basic.Int32(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int8:
					var y int8
					if yIsBasic {
						y = vy.(int8)
					} else {
						y = basic.Int8(vy)
					}
					return func(fr *frame) {
						x := basic.Int32(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int16:
					var y int16
					if yIsBasic {
						y = vy.(int16)
					} else {
						y = basic.Int16(vy)
					}
					return func(fr *frame) {
						x := basic.Int32(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int32:
					var y int32
					if yIsBasic {
						y = vy.(int32)
					} else {
						y = basic.Int32(vy)
					}
					return func(fr *frame) {
						x := basic.Int32(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int64:
					var y int64
					if yIsBasic {
						y = vy.(int64)
					} else {
						y = basic.Int64(vy)
					}
					return func(fr *frame) {
						x := basic.Int32(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint:
					var y uint
					if yIsBasic {
						y = vy.(uint)
					} else {
						y = basic.Uint(vy)
					}
					return func(fr *frame) {
						x := basic.Int32(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint8:
					var y uint8
					if yIsBasic {
						y = vy.(uint8)
					} else {
						y = basic.Uint8(vy)
					}
					return func(fr *frame) {
						x := basic.Int32(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint16:
					var y uint16
					if yIsBasic {
						y = vy.(uint16)
					} else {
						y = basic.Uint16(vy)
					}
					return func(fr *frame) {
						x := basic.Int32(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint32:
					var y uint32
					if yIsBasic {
						y = vy.(uint32)
					} else {
						y = basic.Uint32(vy)
					}
					return func(fr *frame) {
						x := basic.Int32(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint64:
					var y uint64
					if yIsBasic {
						y = vy.(uint64)
					} else {
						y = basic.Uint64(vy)
					}
					return func(fr *frame) {
						x := basic.Int32(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uintptr:
					var y uintptr
					if yIsBasic {
						y = vy.(uintptr)
					} else {
						y = basic.Uintptr(vy)
					}
					return func(fr *frame) {
						x := basic.Int32(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				}
			} else {
				switch ykind {
				case reflect.Int:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Int32(fr.reg(ix))
							y := fr.reg(iy).(int)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						x := basic.Int32(fr.reg(ix))
						y := basic.Int(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int8:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Int32(fr.reg(ix))
							y := fr.reg(iy).(int8)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						x := basic.Int32(fr.reg(ix))
						y := basic.Int8(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int16:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Int32(fr.reg(ix))
							y := fr.reg(iy).(int16)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						x := basic.Int32(fr.reg(ix))
						y := basic.Int16(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int32:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Int32(fr.reg(ix))
							y := fr.reg(iy).(int32)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						x := basic.Int32(fr.reg(ix))
						y := basic.Int32(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int64:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Int32(fr.reg(ix))
							y := fr.reg(iy).(int64)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						x := basic.Int32(fr.reg(ix))
						y := basic.Int64(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Int32(fr.reg(ix))
							y := fr.reg(iy).(uint)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						x := basic.Int32(fr.reg(ix))
						y := basic.Uint(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint8:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Int32(fr.reg(ix))
							y := fr.reg(iy).(uint8)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						x := basic.Int32(fr.reg(ix))
						y := basic.Uint8(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint16:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Int32(fr.reg(ix))
							y := fr.reg(iy).(uint16)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						x := basic.Int32(fr.reg(ix))
						y := basic.Uint16(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint32:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Int32(fr.reg(ix))
							y := fr.reg(iy).(uint32)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						x := basic.Int32(fr.reg(ix))
						y := basic.Uint32(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint64:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Int32(fr.reg(ix))
							y := fr.reg(iy).(uint64)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						x := basic.Int32(fr.reg(ix))
						y := basic.Uint64(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uintptr:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Int32(fr.reg(ix))
							y := fr.reg(iy).(uintptr)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						x := basic.Int32(fr.reg(ix))
						y := basic.Uintptr(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				}
			}
		case reflect.Int64:
			if kx == kindConst && ky == kindConst {
				x := basic.Int64(vx)
				switch ykind {
				case reflect.Int:
					var y int
					if yIsBasic {
						y = vy.(int)
					} else {
						y = basic.Int(vy)
					}
					v := basic.Make(t, x>>y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Int8:
					var y int8
					if yIsBasic {
						y = vy.(int8)
					} else {
						y = basic.Int8(vy)
					}
					v := basic.Make(t, x>>y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Int16:
					var y int16
					if yIsBasic {
						y = vy.(int16)
					} else {
						y = basic.Int16(vy)
					}
					v := basic.Make(t, x>>y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Int32:
					var y int32
					if yIsBasic {
						y = vy.(int32)
					} else {
						y = basic.Int32(vy)
					}
					v := basic.Make(t, x>>y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Int64:
					var y int64
					if yIsBasic {
						y = vy.(int64)
					} else {
						y = basic.Int64(vy)
					}
					v := basic.Make(t, x>>y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uint:
					var y uint
					if yIsBasic {
						y = vy.(uint)
					} else {
						y = basic.Uint(vy)
					}
					v := basic.Make(t, x>>y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uint8:
					var y uint8
					if yIsBasic {
						y = vy.(uint8)
					} else {
						y = basic.Uint8(vy)
					}
					v := basic.Make(t, x>>y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uint16:
					var y uint16
					if yIsBasic {
						y = vy.(uint16)
					} else {
						y = basic.Uint16(vy)
					}
					v := basic.Make(t, x>>y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uint32:
					var y uint32
					if yIsBasic {
						y = vy.(uint32)
					} else {
						y = basic.Uint32(vy)
					}
					v := basic.Make(t, x>>y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uint64:
					var y uint64
					if yIsBasic {
						y = vy.(uint64)
					} else {
						y = basic.Uint64(vy)
					}
					v := basic.Make(t, x>>y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uintptr:
					var y uintptr
					if yIsBasic {
						y = vy.(uintptr)
					} else {
						y = basic.Uintptr(vy)
					}
					v := basic.Make(t, x>>y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				}
			} else if kx == kindConst {
				x := basic.Int64(vx)
				switch ykind {
				case reflect.Int:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(int)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						y := basic.Int(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int8:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(int8)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						y := basic.Int8(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int16:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(int16)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						y := basic.Int16(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int32:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(int32)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						y := basic.Int32(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int64:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(int64)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						y := basic.Int64(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uint)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						y := basic.Uint(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint8:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uint8)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						y := basic.Uint8(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint16:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uint16)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						y := basic.Uint16(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint32:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uint32)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						y := basic.Uint32(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint64:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uint64)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						y := basic.Uint64(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uintptr:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uintptr)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						y := basic.Uintptr(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				}
			} else if ky == kindConst {
				switch ykind {
				case reflect.Int:
					var y int
					if yIsBasic {
						y = vy.(int)
					} else {
						y = basic.Int(vy)
					}
					return func(fr *frame) {
						x := basic.Int64(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int8:
					var y int8
					if yIsBasic {
						y = vy.(int8)
					} else {
						y = basic.Int8(vy)
					}
					return func(fr *frame) {
						x := basic.Int64(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int16:
					var y int16
					if yIsBasic {
						y = vy.(int16)
					} else {
						y = basic.Int16(vy)
					}
					return func(fr *frame) {
						x := basic.Int64(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int32:
					var y int32
					if yIsBasic {
						y = vy.(int32)
					} else {
						y = basic.Int32(vy)
					}
					return func(fr *frame) {
						x := basic.Int64(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int64:
					var y int64
					if yIsBasic {
						y = vy.(int64)
					} else {
						y = basic.Int64(vy)
					}
					return func(fr *frame) {
						x := basic.Int64(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint:
					var y uint
					if yIsBasic {
						y = vy.(uint)
					} else {
						y = basic.Uint(vy)
					}
					return func(fr *frame) {
						x := basic.Int64(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint8:
					var y uint8
					if yIsBasic {
						y = vy.(uint8)
					} else {
						y = basic.Uint8(vy)
					}
					return func(fr *frame) {
						x := basic.Int64(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint16:
					var y uint16
					if yIsBasic {
						y = vy.(uint16)
					} else {
						y = basic.Uint16(vy)
					}
					return func(fr *frame) {
						x := basic.Int64(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint32:
					var y uint32
					if yIsBasic {
						y = vy.(uint32)
					} else {
						y = basic.Uint32(vy)
					}
					return func(fr *frame) {
						x := basic.Int64(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint64:
					var y uint64
					if yIsBasic {
						y = vy.(uint64)
					} else {
						y = basic.Uint64(vy)
					}
					return func(fr *frame) {
						x := basic.Int64(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uintptr:
					var y uintptr
					if yIsBasic {
						y = vy.(uintptr)
					} else {
						y = basic.Uintptr(vy)
					}
					return func(fr *frame) {
						x := basic.Int64(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				}
			} else {
				switch ykind {
				case reflect.Int:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Int64(fr.reg(ix))
							y := fr.reg(iy).(int)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						x := basic.Int64(fr.reg(ix))
						y := basic.Int(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int8:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Int64(fr.reg(ix))
							y := fr.reg(iy).(int8)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						x := basic.Int64(fr.reg(ix))
						y := basic.Int8(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int16:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Int64(fr.reg(ix))
							y := fr.reg(iy).(int16)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						x := basic.Int64(fr.reg(ix))
						y := basic.Int16(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int32:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Int64(fr.reg(ix))
							y := fr.reg(iy).(int32)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						x := basic.Int64(fr.reg(ix))
						y := basic.Int32(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int64:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Int64(fr.reg(ix))
							y := fr.reg(iy).(int64)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						x := basic.Int64(fr.reg(ix))
						y := basic.Int64(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Int64(fr.reg(ix))
							y := fr.reg(iy).(uint)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						x := basic.Int64(fr.reg(ix))
						y := basic.Uint(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint8:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Int64(fr.reg(ix))
							y := fr.reg(iy).(uint8)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						x := basic.Int64(fr.reg(ix))
						y := basic.Uint8(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint16:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Int64(fr.reg(ix))
							y := fr.reg(iy).(uint16)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						x := basic.Int64(fr.reg(ix))
						y := basic.Uint16(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint32:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Int64(fr.reg(ix))
							y := fr.reg(iy).(uint32)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						x := basic.Int64(fr.reg(ix))
						y := basic.Uint32(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint64:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Int64(fr.reg(ix))
							y := fr.reg(iy).(uint64)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						x := basic.Int64(fr.reg(ix))
						y := basic.Uint64(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uintptr:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Int64(fr.reg(ix))
							y := fr.reg(iy).(uintptr)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						x := basic.Int64(fr.reg(ix))
						y := basic.Uintptr(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				}
			}
		case reflect.Uint:
			if kx == kindConst && ky == kindConst {
				x := basic.Uint(vx)
				switch ykind {
				case reflect.Int:
					var y int
					if yIsBasic {
						y = vy.(int)
					} else {
						y = basic.Int(vy)
					}
					v := basic.Make(t, x>>y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Int8:
					var y int8
					if yIsBasic {
						y = vy.(int8)
					} else {
						y = basic.Int8(vy)
					}
					v := basic.Make(t, x>>y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Int16:
					var y int16
					if yIsBasic {
						y = vy.(int16)
					} else {
						y = basic.Int16(vy)
					}
					v := basic.Make(t, x>>y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Int32:
					var y int32
					if yIsBasic {
						y = vy.(int32)
					} else {
						y = basic.Int32(vy)
					}
					v := basic.Make(t, x>>y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Int64:
					var y int64
					if yIsBasic {
						y = vy.(int64)
					} else {
						y = basic.Int64(vy)
					}
					v := basic.Make(t, x>>y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uint:
					var y uint
					if yIsBasic {
						y = vy.(uint)
					} else {
						y = basic.Uint(vy)
					}
					v := basic.Make(t, x>>y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uint8:
					var y uint8
					if yIsBasic {
						y = vy.(uint8)
					} else {
						y = basic.Uint8(vy)
					}
					v := basic.Make(t, x>>y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uint16:
					var y uint16
					if yIsBasic {
						y = vy.(uint16)
					} else {
						y = basic.Uint16(vy)
					}
					v := basic.Make(t, x>>y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uint32:
					var y uint32
					if yIsBasic {
						y = vy.(uint32)
					} else {
						y = basic.Uint32(vy)
					}
					v := basic.Make(t, x>>y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uint64:
					var y uint64
					if yIsBasic {
						y = vy.(uint64)
					} else {
						y = basic.Uint64(vy)
					}
					v := basic.Make(t, x>>y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uintptr:
					var y uintptr
					if yIsBasic {
						y = vy.(uintptr)
					} else {
						y = basic.Uintptr(vy)
					}
					v := basic.Make(t, x>>y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				}
			} else if kx == kindConst {
				x := basic.Uint(vx)
				switch ykind {
				case reflect.Int:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(int)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						y := basic.Int(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int8:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(int8)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						y := basic.Int8(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int16:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(int16)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						y := basic.Int16(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int32:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(int32)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						y := basic.Int32(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int64:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(int64)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						y := basic.Int64(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uint)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						y := basic.Uint(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint8:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uint8)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						y := basic.Uint8(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint16:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uint16)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						y := basic.Uint16(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint32:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uint32)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						y := basic.Uint32(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint64:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uint64)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						y := basic.Uint64(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uintptr:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uintptr)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						y := basic.Uintptr(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				}
			} else if ky == kindConst {
				switch ykind {
				case reflect.Int:
					var y int
					if yIsBasic {
						y = vy.(int)
					} else {
						y = basic.Int(vy)
					}
					return func(fr *frame) {
						x := basic.Uint(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int8:
					var y int8
					if yIsBasic {
						y = vy.(int8)
					} else {
						y = basic.Int8(vy)
					}
					return func(fr *frame) {
						x := basic.Uint(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int16:
					var y int16
					if yIsBasic {
						y = vy.(int16)
					} else {
						y = basic.Int16(vy)
					}
					return func(fr *frame) {
						x := basic.Uint(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int32:
					var y int32
					if yIsBasic {
						y = vy.(int32)
					} else {
						y = basic.Int32(vy)
					}
					return func(fr *frame) {
						x := basic.Uint(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int64:
					var y int64
					if yIsBasic {
						y = vy.(int64)
					} else {
						y = basic.Int64(vy)
					}
					return func(fr *frame) {
						x := basic.Uint(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint:
					var y uint
					if yIsBasic {
						y = vy.(uint)
					} else {
						y = basic.Uint(vy)
					}
					return func(fr *frame) {
						x := basic.Uint(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint8:
					var y uint8
					if yIsBasic {
						y = vy.(uint8)
					} else {
						y = basic.Uint8(vy)
					}
					return func(fr *frame) {
						x := basic.Uint(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint16:
					var y uint16
					if yIsBasic {
						y = vy.(uint16)
					} else {
						y = basic.Uint16(vy)
					}
					return func(fr *frame) {
						x := basic.Uint(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint32:
					var y uint32
					if yIsBasic {
						y = vy.(uint32)
					} else {
						y = basic.Uint32(vy)
					}
					return func(fr *frame) {
						x := basic.Uint(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint64:
					var y uint64
					if yIsBasic {
						y = vy.(uint64)
					} else {
						y = basic.Uint64(vy)
					}
					return func(fr *frame) {
						x := basic.Uint(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uintptr:
					var y uintptr
					if yIsBasic {
						y = vy.(uintptr)
					} else {
						y = basic.Uintptr(vy)
					}
					return func(fr *frame) {
						x := basic.Uint(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				}
			} else {
				switch ykind {
				case reflect.Int:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Uint(fr.reg(ix))
							y := fr.reg(iy).(int)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						x := basic.Uint(fr.reg(ix))
						y := basic.Int(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int8:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Uint(fr.reg(ix))
							y := fr.reg(iy).(int8)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						x := basic.Uint(fr.reg(ix))
						y := basic.Int8(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int16:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Uint(fr.reg(ix))
							y := fr.reg(iy).(int16)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						x := basic.Uint(fr.reg(ix))
						y := basic.Int16(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int32:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Uint(fr.reg(ix))
							y := fr.reg(iy).(int32)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						x := basic.Uint(fr.reg(ix))
						y := basic.Int32(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int64:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Uint(fr.reg(ix))
							y := fr.reg(iy).(int64)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						x := basic.Uint(fr.reg(ix))
						y := basic.Int64(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Uint(fr.reg(ix))
							y := fr.reg(iy).(uint)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						x := basic.Uint(fr.reg(ix))
						y := basic.Uint(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint8:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Uint(fr.reg(ix))
							y := fr.reg(iy).(uint8)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						x := basic.Uint(fr.reg(ix))
						y := basic.Uint8(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint16:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Uint(fr.reg(ix))
							y := fr.reg(iy).(uint16)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						x := basic.Uint(fr.reg(ix))
						y := basic.Uint16(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint32:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Uint(fr.reg(ix))
							y := fr.reg(iy).(uint32)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						x := basic.Uint(fr.reg(ix))
						y := basic.Uint32(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint64:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Uint(fr.reg(ix))
							y := fr.reg(iy).(uint64)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						x := basic.Uint(fr.reg(ix))
						y := basic.Uint64(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uintptr:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Uint(fr.reg(ix))
							y := fr.reg(iy).(uintptr)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						x := basic.Uint(fr.reg(ix))
						y := basic.Uintptr(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				}
			}
		case reflect.Uint8:
			if kx == kindConst && ky == kindConst {
				x := basic.Uint8(vx)
				switch ykind {
				case reflect.Int:
					var y int
					if yIsBasic {
						y = vy.(int)
					} else {
						y = basic.Int(vy)
					}
					v := basic.Make(t, x>>y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Int8:
					var y int8
					if yIsBasic {
						y = vy.(int8)
					} else {
						y = basic.Int8(vy)
					}
					v := basic.Make(t, x>>y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Int16:
					var y int16
					if yIsBasic {
						y = vy.(int16)
					} else {
						y = basic.Int16(vy)
					}
					v := basic.Make(t, x>>y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Int32:
					var y int32
					if yIsBasic {
						y = vy.(int32)
					} else {
						y = basic.Int32(vy)
					}
					v := basic.Make(t, x>>y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Int64:
					var y int64
					if yIsBasic {
						y = vy.(int64)
					} else {
						y = basic.Int64(vy)
					}
					v := basic.Make(t, x>>y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uint:
					var y uint
					if yIsBasic {
						y = vy.(uint)
					} else {
						y = basic.Uint(vy)
					}
					v := basic.Make(t, x>>y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uint8:
					var y uint8
					if yIsBasic {
						y = vy.(uint8)
					} else {
						y = basic.Uint8(vy)
					}
					v := basic.Make(t, x>>y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uint16:
					var y uint16
					if yIsBasic {
						y = vy.(uint16)
					} else {
						y = basic.Uint16(vy)
					}
					v := basic.Make(t, x>>y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uint32:
					var y uint32
					if yIsBasic {
						y = vy.(uint32)
					} else {
						y = basic.Uint32(vy)
					}
					v := basic.Make(t, x>>y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uint64:
					var y uint64
					if yIsBasic {
						y = vy.(uint64)
					} else {
						y = basic.Uint64(vy)
					}
					v := basic.Make(t, x>>y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uintptr:
					var y uintptr
					if yIsBasic {
						y = vy.(uintptr)
					} else {
						y = basic.Uintptr(vy)
					}
					v := basic.Make(t, x>>y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				}
			} else if kx == kindConst {
				x := basic.Uint8(vx)
				switch ykind {
				case reflect.Int:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(int)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						y := basic.Int(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int8:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(int8)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						y := basic.Int8(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int16:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(int16)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						y := basic.Int16(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int32:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(int32)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						y := basic.Int32(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int64:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(int64)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						y := basic.Int64(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uint)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						y := basic.Uint(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint8:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uint8)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						y := basic.Uint8(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint16:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uint16)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						y := basic.Uint16(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint32:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uint32)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						y := basic.Uint32(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint64:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uint64)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						y := basic.Uint64(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uintptr:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uintptr)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						y := basic.Uintptr(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				}
			} else if ky == kindConst {
				switch ykind {
				case reflect.Int:
					var y int
					if yIsBasic {
						y = vy.(int)
					} else {
						y = basic.Int(vy)
					}
					return func(fr *frame) {
						x := basic.Uint8(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int8:
					var y int8
					if yIsBasic {
						y = vy.(int8)
					} else {
						y = basic.Int8(vy)
					}
					return func(fr *frame) {
						x := basic.Uint8(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int16:
					var y int16
					if yIsBasic {
						y = vy.(int16)
					} else {
						y = basic.Int16(vy)
					}
					return func(fr *frame) {
						x := basic.Uint8(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int32:
					var y int32
					if yIsBasic {
						y = vy.(int32)
					} else {
						y = basic.Int32(vy)
					}
					return func(fr *frame) {
						x := basic.Uint8(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int64:
					var y int64
					if yIsBasic {
						y = vy.(int64)
					} else {
						y = basic.Int64(vy)
					}
					return func(fr *frame) {
						x := basic.Uint8(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint:
					var y uint
					if yIsBasic {
						y = vy.(uint)
					} else {
						y = basic.Uint(vy)
					}
					return func(fr *frame) {
						x := basic.Uint8(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint8:
					var y uint8
					if yIsBasic {
						y = vy.(uint8)
					} else {
						y = basic.Uint8(vy)
					}
					return func(fr *frame) {
						x := basic.Uint8(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint16:
					var y uint16
					if yIsBasic {
						y = vy.(uint16)
					} else {
						y = basic.Uint16(vy)
					}
					return func(fr *frame) {
						x := basic.Uint8(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint32:
					var y uint32
					if yIsBasic {
						y = vy.(uint32)
					} else {
						y = basic.Uint32(vy)
					}
					return func(fr *frame) {
						x := basic.Uint8(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint64:
					var y uint64
					if yIsBasic {
						y = vy.(uint64)
					} else {
						y = basic.Uint64(vy)
					}
					return func(fr *frame) {
						x := basic.Uint8(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uintptr:
					var y uintptr
					if yIsBasic {
						y = vy.(uintptr)
					} else {
						y = basic.Uintptr(vy)
					}
					return func(fr *frame) {
						x := basic.Uint8(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				}
			} else {
				switch ykind {
				case reflect.Int:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Uint8(fr.reg(ix))
							y := fr.reg(iy).(int)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						x := basic.Uint8(fr.reg(ix))
						y := basic.Int(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int8:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Uint8(fr.reg(ix))
							y := fr.reg(iy).(int8)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						x := basic.Uint8(fr.reg(ix))
						y := basic.Int8(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int16:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Uint8(fr.reg(ix))
							y := fr.reg(iy).(int16)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						x := basic.Uint8(fr.reg(ix))
						y := basic.Int16(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int32:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Uint8(fr.reg(ix))
							y := fr.reg(iy).(int32)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						x := basic.Uint8(fr.reg(ix))
						y := basic.Int32(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int64:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Uint8(fr.reg(ix))
							y := fr.reg(iy).(int64)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						x := basic.Uint8(fr.reg(ix))
						y := basic.Int64(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Uint8(fr.reg(ix))
							y := fr.reg(iy).(uint)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						x := basic.Uint8(fr.reg(ix))
						y := basic.Uint(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint8:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Uint8(fr.reg(ix))
							y := fr.reg(iy).(uint8)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						x := basic.Uint8(fr.reg(ix))
						y := basic.Uint8(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint16:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Uint8(fr.reg(ix))
							y := fr.reg(iy).(uint16)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						x := basic.Uint8(fr.reg(ix))
						y := basic.Uint16(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint32:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Uint8(fr.reg(ix))
							y := fr.reg(iy).(uint32)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						x := basic.Uint8(fr.reg(ix))
						y := basic.Uint32(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint64:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Uint8(fr.reg(ix))
							y := fr.reg(iy).(uint64)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						x := basic.Uint8(fr.reg(ix))
						y := basic.Uint64(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uintptr:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Uint8(fr.reg(ix))
							y := fr.reg(iy).(uintptr)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						x := basic.Uint8(fr.reg(ix))
						y := basic.Uintptr(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				}
			}
		case reflect.Uint16:
			if kx == kindConst && ky == kindConst {
				x := basic.Uint16(vx)
				switch ykind {
				case reflect.Int:
					var y int
					if yIsBasic {
						y = vy.(int)
					} else {
						y = basic.Int(vy)
					}
					v := basic.Make(t, x>>y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Int8:
					var y int8
					if yIsBasic {
						y = vy.(int8)
					} else {
						y = basic.Int8(vy)
					}
					v := basic.Make(t, x>>y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Int16:
					var y int16
					if yIsBasic {
						y = vy.(int16)
					} else {
						y = basic.Int16(vy)
					}
					v := basic.Make(t, x>>y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Int32:
					var y int32
					if yIsBasic {
						y = vy.(int32)
					} else {
						y = basic.Int32(vy)
					}
					v := basic.Make(t, x>>y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Int64:
					var y int64
					if yIsBasic {
						y = vy.(int64)
					} else {
						y = basic.Int64(vy)
					}
					v := basic.Make(t, x>>y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uint:
					var y uint
					if yIsBasic {
						y = vy.(uint)
					} else {
						y = basic.Uint(vy)
					}
					v := basic.Make(t, x>>y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uint8:
					var y uint8
					if yIsBasic {
						y = vy.(uint8)
					} else {
						y = basic.Uint8(vy)
					}
					v := basic.Make(t, x>>y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uint16:
					var y uint16
					if yIsBasic {
						y = vy.(uint16)
					} else {
						y = basic.Uint16(vy)
					}
					v := basic.Make(t, x>>y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uint32:
					var y uint32
					if yIsBasic {
						y = vy.(uint32)
					} else {
						y = basic.Uint32(vy)
					}
					v := basic.Make(t, x>>y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uint64:
					var y uint64
					if yIsBasic {
						y = vy.(uint64)
					} else {
						y = basic.Uint64(vy)
					}
					v := basic.Make(t, x>>y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uintptr:
					var y uintptr
					if yIsBasic {
						y = vy.(uintptr)
					} else {
						y = basic.Uintptr(vy)
					}
					v := basic.Make(t, x>>y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				}
			} else if kx == kindConst {
				x := basic.Uint16(vx)
				switch ykind {
				case reflect.Int:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(int)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						y := basic.Int(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int8:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(int8)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						y := basic.Int8(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int16:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(int16)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						y := basic.Int16(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int32:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(int32)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						y := basic.Int32(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int64:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(int64)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						y := basic.Int64(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uint)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						y := basic.Uint(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint8:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uint8)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						y := basic.Uint8(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint16:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uint16)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						y := basic.Uint16(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint32:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uint32)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						y := basic.Uint32(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint64:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uint64)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						y := basic.Uint64(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uintptr:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uintptr)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						y := basic.Uintptr(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				}
			} else if ky == kindConst {
				switch ykind {
				case reflect.Int:
					var y int
					if yIsBasic {
						y = vy.(int)
					} else {
						y = basic.Int(vy)
					}
					return func(fr *frame) {
						x := basic.Uint16(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int8:
					var y int8
					if yIsBasic {
						y = vy.(int8)
					} else {
						y = basic.Int8(vy)
					}
					return func(fr *frame) {
						x := basic.Uint16(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int16:
					var y int16
					if yIsBasic {
						y = vy.(int16)
					} else {
						y = basic.Int16(vy)
					}
					return func(fr *frame) {
						x := basic.Uint16(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int32:
					var y int32
					if yIsBasic {
						y = vy.(int32)
					} else {
						y = basic.Int32(vy)
					}
					return func(fr *frame) {
						x := basic.Uint16(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int64:
					var y int64
					if yIsBasic {
						y = vy.(int64)
					} else {
						y = basic.Int64(vy)
					}
					return func(fr *frame) {
						x := basic.Uint16(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint:
					var y uint
					if yIsBasic {
						y = vy.(uint)
					} else {
						y = basic.Uint(vy)
					}
					return func(fr *frame) {
						x := basic.Uint16(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint8:
					var y uint8
					if yIsBasic {
						y = vy.(uint8)
					} else {
						y = basic.Uint8(vy)
					}
					return func(fr *frame) {
						x := basic.Uint16(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint16:
					var y uint16
					if yIsBasic {
						y = vy.(uint16)
					} else {
						y = basic.Uint16(vy)
					}
					return func(fr *frame) {
						x := basic.Uint16(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint32:
					var y uint32
					if yIsBasic {
						y = vy.(uint32)
					} else {
						y = basic.Uint32(vy)
					}
					return func(fr *frame) {
						x := basic.Uint16(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint64:
					var y uint64
					if yIsBasic {
						y = vy.(uint64)
					} else {
						y = basic.Uint64(vy)
					}
					return func(fr *frame) {
						x := basic.Uint16(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uintptr:
					var y uintptr
					if yIsBasic {
						y = vy.(uintptr)
					} else {
						y = basic.Uintptr(vy)
					}
					return func(fr *frame) {
						x := basic.Uint16(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				}
			} else {
				switch ykind {
				case reflect.Int:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Uint16(fr.reg(ix))
							y := fr.reg(iy).(int)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						x := basic.Uint16(fr.reg(ix))
						y := basic.Int(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int8:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Uint16(fr.reg(ix))
							y := fr.reg(iy).(int8)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						x := basic.Uint16(fr.reg(ix))
						y := basic.Int8(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int16:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Uint16(fr.reg(ix))
							y := fr.reg(iy).(int16)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						x := basic.Uint16(fr.reg(ix))
						y := basic.Int16(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int32:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Uint16(fr.reg(ix))
							y := fr.reg(iy).(int32)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						x := basic.Uint16(fr.reg(ix))
						y := basic.Int32(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int64:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Uint16(fr.reg(ix))
							y := fr.reg(iy).(int64)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						x := basic.Uint16(fr.reg(ix))
						y := basic.Int64(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Uint16(fr.reg(ix))
							y := fr.reg(iy).(uint)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						x := basic.Uint16(fr.reg(ix))
						y := basic.Uint(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint8:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Uint16(fr.reg(ix))
							y := fr.reg(iy).(uint8)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						x := basic.Uint16(fr.reg(ix))
						y := basic.Uint8(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint16:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Uint16(fr.reg(ix))
							y := fr.reg(iy).(uint16)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						x := basic.Uint16(fr.reg(ix))
						y := basic.Uint16(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint32:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Uint16(fr.reg(ix))
							y := fr.reg(iy).(uint32)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						x := basic.Uint16(fr.reg(ix))
						y := basic.Uint32(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint64:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Uint16(fr.reg(ix))
							y := fr.reg(iy).(uint64)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						x := basic.Uint16(fr.reg(ix))
						y := basic.Uint64(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uintptr:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Uint16(fr.reg(ix))
							y := fr.reg(iy).(uintptr)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						x := basic.Uint16(fr.reg(ix))
						y := basic.Uintptr(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				}
			}
		case reflect.Uint32:
			if kx == kindConst && ky == kindConst {
				x := basic.Uint32(vx)
				switch ykind {
				case reflect.Int:
					var y int
					if yIsBasic {
						y = vy.(int)
					} else {
						y = basic.Int(vy)
					}
					v := basic.Make(t, x>>y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Int8:
					var y int8
					if yIsBasic {
						y = vy.(int8)
					} else {
						y = basic.Int8(vy)
					}
					v := basic.Make(t, x>>y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Int16:
					var y int16
					if yIsBasic {
						y = vy.(int16)
					} else {
						y = basic.Int16(vy)
					}
					v := basic.Make(t, x>>y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Int32:
					var y int32
					if yIsBasic {
						y = vy.(int32)
					} else {
						y = basic.Int32(vy)
					}
					v := basic.Make(t, x>>y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Int64:
					var y int64
					if yIsBasic {
						y = vy.(int64)
					} else {
						y = basic.Int64(vy)
					}
					v := basic.Make(t, x>>y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uint:
					var y uint
					if yIsBasic {
						y = vy.(uint)
					} else {
						y = basic.Uint(vy)
					}
					v := basic.Make(t, x>>y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uint8:
					var y uint8
					if yIsBasic {
						y = vy.(uint8)
					} else {
						y = basic.Uint8(vy)
					}
					v := basic.Make(t, x>>y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uint16:
					var y uint16
					if yIsBasic {
						y = vy.(uint16)
					} else {
						y = basic.Uint16(vy)
					}
					v := basic.Make(t, x>>y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uint32:
					var y uint32
					if yIsBasic {
						y = vy.(uint32)
					} else {
						y = basic.Uint32(vy)
					}
					v := basic.Make(t, x>>y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uint64:
					var y uint64
					if yIsBasic {
						y = vy.(uint64)
					} else {
						y = basic.Uint64(vy)
					}
					v := basic.Make(t, x>>y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uintptr:
					var y uintptr
					if yIsBasic {
						y = vy.(uintptr)
					} else {
						y = basic.Uintptr(vy)
					}
					v := basic.Make(t, x>>y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				}
			} else if kx == kindConst {
				x := basic.Uint32(vx)
				switch ykind {
				case reflect.Int:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(int)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						y := basic.Int(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int8:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(int8)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						y := basic.Int8(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int16:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(int16)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						y := basic.Int16(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int32:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(int32)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						y := basic.Int32(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int64:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(int64)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						y := basic.Int64(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uint)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						y := basic.Uint(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint8:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uint8)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						y := basic.Uint8(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint16:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uint16)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						y := basic.Uint16(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint32:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uint32)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						y := basic.Uint32(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint64:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uint64)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						y := basic.Uint64(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uintptr:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uintptr)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						y := basic.Uintptr(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				}
			} else if ky == kindConst {
				switch ykind {
				case reflect.Int:
					var y int
					if yIsBasic {
						y = vy.(int)
					} else {
						y = basic.Int(vy)
					}
					return func(fr *frame) {
						x := basic.Uint32(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int8:
					var y int8
					if yIsBasic {
						y = vy.(int8)
					} else {
						y = basic.Int8(vy)
					}
					return func(fr *frame) {
						x := basic.Uint32(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int16:
					var y int16
					if yIsBasic {
						y = vy.(int16)
					} else {
						y = basic.Int16(vy)
					}
					return func(fr *frame) {
						x := basic.Uint32(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int32:
					var y int32
					if yIsBasic {
						y = vy.(int32)
					} else {
						y = basic.Int32(vy)
					}
					return func(fr *frame) {
						x := basic.Uint32(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int64:
					var y int64
					if yIsBasic {
						y = vy.(int64)
					} else {
						y = basic.Int64(vy)
					}
					return func(fr *frame) {
						x := basic.Uint32(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint:
					var y uint
					if yIsBasic {
						y = vy.(uint)
					} else {
						y = basic.Uint(vy)
					}
					return func(fr *frame) {
						x := basic.Uint32(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint8:
					var y uint8
					if yIsBasic {
						y = vy.(uint8)
					} else {
						y = basic.Uint8(vy)
					}
					return func(fr *frame) {
						x := basic.Uint32(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint16:
					var y uint16
					if yIsBasic {
						y = vy.(uint16)
					} else {
						y = basic.Uint16(vy)
					}
					return func(fr *frame) {
						x := basic.Uint32(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint32:
					var y uint32
					if yIsBasic {
						y = vy.(uint32)
					} else {
						y = basic.Uint32(vy)
					}
					return func(fr *frame) {
						x := basic.Uint32(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint64:
					var y uint64
					if yIsBasic {
						y = vy.(uint64)
					} else {
						y = basic.Uint64(vy)
					}
					return func(fr *frame) {
						x := basic.Uint32(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uintptr:
					var y uintptr
					if yIsBasic {
						y = vy.(uintptr)
					} else {
						y = basic.Uintptr(vy)
					}
					return func(fr *frame) {
						x := basic.Uint32(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				}
			} else {
				switch ykind {
				case reflect.Int:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Uint32(fr.reg(ix))
							y := fr.reg(iy).(int)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						x := basic.Uint32(fr.reg(ix))
						y := basic.Int(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int8:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Uint32(fr.reg(ix))
							y := fr.reg(iy).(int8)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						x := basic.Uint32(fr.reg(ix))
						y := basic.Int8(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int16:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Uint32(fr.reg(ix))
							y := fr.reg(iy).(int16)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						x := basic.Uint32(fr.reg(ix))
						y := basic.Int16(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int32:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Uint32(fr.reg(ix))
							y := fr.reg(iy).(int32)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						x := basic.Uint32(fr.reg(ix))
						y := basic.Int32(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int64:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Uint32(fr.reg(ix))
							y := fr.reg(iy).(int64)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						x := basic.Uint32(fr.reg(ix))
						y := basic.Int64(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Uint32(fr.reg(ix))
							y := fr.reg(iy).(uint)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						x := basic.Uint32(fr.reg(ix))
						y := basic.Uint(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint8:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Uint32(fr.reg(ix))
							y := fr.reg(iy).(uint8)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						x := basic.Uint32(fr.reg(ix))
						y := basic.Uint8(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint16:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Uint32(fr.reg(ix))
							y := fr.reg(iy).(uint16)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						x := basic.Uint32(fr.reg(ix))
						y := basic.Uint16(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint32:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Uint32(fr.reg(ix))
							y := fr.reg(iy).(uint32)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						x := basic.Uint32(fr.reg(ix))
						y := basic.Uint32(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint64:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Uint32(fr.reg(ix))
							y := fr.reg(iy).(uint64)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						x := basic.Uint32(fr.reg(ix))
						y := basic.Uint64(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uintptr:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Uint32(fr.reg(ix))
							y := fr.reg(iy).(uintptr)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						x := basic.Uint32(fr.reg(ix))
						y := basic.Uintptr(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				}
			}
		case reflect.Uint64:
			if kx == kindConst && ky == kindConst {
				x := basic.Uint64(vx)
				switch ykind {
				case reflect.Int:
					var y int
					if yIsBasic {
						y = vy.(int)
					} else {
						y = basic.Int(vy)
					}
					v := basic.Make(t, x>>y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Int8:
					var y int8
					if yIsBasic {
						y = vy.(int8)
					} else {
						y = basic.Int8(vy)
					}
					v := basic.Make(t, x>>y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Int16:
					var y int16
					if yIsBasic {
						y = vy.(int16)
					} else {
						y = basic.Int16(vy)
					}
					v := basic.Make(t, x>>y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Int32:
					var y int32
					if yIsBasic {
						y = vy.(int32)
					} else {
						y = basic.Int32(vy)
					}
					v := basic.Make(t, x>>y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Int64:
					var y int64
					if yIsBasic {
						y = vy.(int64)
					} else {
						y = basic.Int64(vy)
					}
					v := basic.Make(t, x>>y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uint:
					var y uint
					if yIsBasic {
						y = vy.(uint)
					} else {
						y = basic.Uint(vy)
					}
					v := basic.Make(t, x>>y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uint8:
					var y uint8
					if yIsBasic {
						y = vy.(uint8)
					} else {
						y = basic.Uint8(vy)
					}
					v := basic.Make(t, x>>y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uint16:
					var y uint16
					if yIsBasic {
						y = vy.(uint16)
					} else {
						y = basic.Uint16(vy)
					}
					v := basic.Make(t, x>>y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uint32:
					var y uint32
					if yIsBasic {
						y = vy.(uint32)
					} else {
						y = basic.Uint32(vy)
					}
					v := basic.Make(t, x>>y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uint64:
					var y uint64
					if yIsBasic {
						y = vy.(uint64)
					} else {
						y = basic.Uint64(vy)
					}
					v := basic.Make(t, x>>y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uintptr:
					var y uintptr
					if yIsBasic {
						y = vy.(uintptr)
					} else {
						y = basic.Uintptr(vy)
					}
					v := basic.Make(t, x>>y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				}
			} else if kx == kindConst {
				x := basic.Uint64(vx)
				switch ykind {
				case reflect.Int:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(int)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						y := basic.Int(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int8:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(int8)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						y := basic.Int8(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int16:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(int16)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						y := basic.Int16(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int32:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(int32)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						y := basic.Int32(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int64:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(int64)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						y := basic.Int64(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uint)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						y := basic.Uint(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint8:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uint8)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						y := basic.Uint8(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint16:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uint16)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						y := basic.Uint16(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint32:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uint32)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						y := basic.Uint32(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint64:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uint64)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						y := basic.Uint64(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uintptr:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uintptr)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						y := basic.Uintptr(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				}
			} else if ky == kindConst {
				switch ykind {
				case reflect.Int:
					var y int
					if yIsBasic {
						y = vy.(int)
					} else {
						y = basic.Int(vy)
					}
					return func(fr *frame) {
						x := basic.Uint64(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int8:
					var y int8
					if yIsBasic {
						y = vy.(int8)
					} else {
						y = basic.Int8(vy)
					}
					return func(fr *frame) {
						x := basic.Uint64(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int16:
					var y int16
					if yIsBasic {
						y = vy.(int16)
					} else {
						y = basic.Int16(vy)
					}
					return func(fr *frame) {
						x := basic.Uint64(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int32:
					var y int32
					if yIsBasic {
						y = vy.(int32)
					} else {
						y = basic.Int32(vy)
					}
					return func(fr *frame) {
						x := basic.Uint64(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int64:
					var y int64
					if yIsBasic {
						y = vy.(int64)
					} else {
						y = basic.Int64(vy)
					}
					return func(fr *frame) {
						x := basic.Uint64(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint:
					var y uint
					if yIsBasic {
						y = vy.(uint)
					} else {
						y = basic.Uint(vy)
					}
					return func(fr *frame) {
						x := basic.Uint64(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint8:
					var y uint8
					if yIsBasic {
						y = vy.(uint8)
					} else {
						y = basic.Uint8(vy)
					}
					return func(fr *frame) {
						x := basic.Uint64(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint16:
					var y uint16
					if yIsBasic {
						y = vy.(uint16)
					} else {
						y = basic.Uint16(vy)
					}
					return func(fr *frame) {
						x := basic.Uint64(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint32:
					var y uint32
					if yIsBasic {
						y = vy.(uint32)
					} else {
						y = basic.Uint32(vy)
					}
					return func(fr *frame) {
						x := basic.Uint64(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint64:
					var y uint64
					if yIsBasic {
						y = vy.(uint64)
					} else {
						y = basic.Uint64(vy)
					}
					return func(fr *frame) {
						x := basic.Uint64(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uintptr:
					var y uintptr
					if yIsBasic {
						y = vy.(uintptr)
					} else {
						y = basic.Uintptr(vy)
					}
					return func(fr *frame) {
						x := basic.Uint64(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				}
			} else {
				switch ykind {
				case reflect.Int:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Uint64(fr.reg(ix))
							y := fr.reg(iy).(int)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						x := basic.Uint64(fr.reg(ix))
						y := basic.Int(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int8:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Uint64(fr.reg(ix))
							y := fr.reg(iy).(int8)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						x := basic.Uint64(fr.reg(ix))
						y := basic.Int8(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int16:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Uint64(fr.reg(ix))
							y := fr.reg(iy).(int16)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						x := basic.Uint64(fr.reg(ix))
						y := basic.Int16(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int32:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Uint64(fr.reg(ix))
							y := fr.reg(iy).(int32)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						x := basic.Uint64(fr.reg(ix))
						y := basic.Int32(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int64:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Uint64(fr.reg(ix))
							y := fr.reg(iy).(int64)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						x := basic.Uint64(fr.reg(ix))
						y := basic.Int64(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Uint64(fr.reg(ix))
							y := fr.reg(iy).(uint)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						x := basic.Uint64(fr.reg(ix))
						y := basic.Uint(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint8:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Uint64(fr.reg(ix))
							y := fr.reg(iy).(uint8)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						x := basic.Uint64(fr.reg(ix))
						y := basic.Uint8(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint16:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Uint64(fr.reg(ix))
							y := fr.reg(iy).(uint16)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						x := basic.Uint64(fr.reg(ix))
						y := basic.Uint16(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint32:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Uint64(fr.reg(ix))
							y := fr.reg(iy).(uint32)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						x := basic.Uint64(fr.reg(ix))
						y := basic.Uint32(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint64:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Uint64(fr.reg(ix))
							y := fr.reg(iy).(uint64)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						x := basic.Uint64(fr.reg(ix))
						y := basic.Uint64(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uintptr:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Uint64(fr.reg(ix))
							y := fr.reg(iy).(uintptr)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						x := basic.Uint64(fr.reg(ix))
						y := basic.Uintptr(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				}
			}
		case reflect.Uintptr:
			if kx == kindConst && ky == kindConst {
				x := basic.Uintptr(vx)
				switch ykind {
				case reflect.Int:
					var y int
					if yIsBasic {
						y = vy.(int)
					} else {
						y = basic.Int(vy)
					}
					v := basic.Make(t, x>>y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Int8:
					var y int8
					if yIsBasic {
						y = vy.(int8)
					} else {
						y = basic.Int8(vy)
					}
					v := basic.Make(t, x>>y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Int16:
					var y int16
					if yIsBasic {
						y = vy.(int16)
					} else {
						y = basic.Int16(vy)
					}
					v := basic.Make(t, x>>y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Int32:
					var y int32
					if yIsBasic {
						y = vy.(int32)
					} else {
						y = basic.Int32(vy)
					}
					v := basic.Make(t, x>>y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Int64:
					var y int64
					if yIsBasic {
						y = vy.(int64)
					} else {
						y = basic.Int64(vy)
					}
					v := basic.Make(t, x>>y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uint:
					var y uint
					if yIsBasic {
						y = vy.(uint)
					} else {
						y = basic.Uint(vy)
					}
					v := basic.Make(t, x>>y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uint8:
					var y uint8
					if yIsBasic {
						y = vy.(uint8)
					} else {
						y = basic.Uint8(vy)
					}
					v := basic.Make(t, x>>y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uint16:
					var y uint16
					if yIsBasic {
						y = vy.(uint16)
					} else {
						y = basic.Uint16(vy)
					}
					v := basic.Make(t, x>>y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uint32:
					var y uint32
					if yIsBasic {
						y = vy.(uint32)
					} else {
						y = basic.Uint32(vy)
					}
					v := basic.Make(t, x>>y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uint64:
					var y uint64
					if yIsBasic {
						y = vy.(uint64)
					} else {
						y = basic.Uint64(vy)
					}
					v := basic.Make(t, x>>y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				case reflect.Uintptr:
					var y uintptr
					if yIsBasic {
						y = vy.(uintptr)
					} else {
						y = basic.Uintptr(vy)
					}
					v := basic.Make(t, x>>y)
					return func(fr *frame) {
						fr.setReg(ir, v)
					}
				}
			} else if kx == kindConst {
				x := basic.Uintptr(vx)
				switch ykind {
				case reflect.Int:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(int)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						y := basic.Int(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int8:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(int8)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						y := basic.Int8(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int16:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(int16)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						y := basic.Int16(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int32:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(int32)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						y := basic.Int32(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int64:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(int64)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						y := basic.Int64(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uint)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						y := basic.Uint(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint8:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uint8)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						y := basic.Uint8(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint16:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uint16)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						y := basic.Uint16(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint32:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uint32)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						y := basic.Uint32(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint64:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uint64)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						y := basic.Uint64(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uintptr:
					if yIsBasic {
						return func(fr *frame) {
							y := fr.reg(iy).(uintptr)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						y := basic.Uintptr(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				}
			} else if ky == kindConst {
				switch ykind {
				case reflect.Int:
					var y int
					if yIsBasic {
						y = vy.(int)
					} else {
						y = basic.Int(vy)
					}
					return func(fr *frame) {
						x := basic.Uintptr(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int8:
					var y int8
					if yIsBasic {
						y = vy.(int8)
					} else {
						y = basic.Int8(vy)
					}
					return func(fr *frame) {
						x := basic.Uintptr(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int16:
					var y int16
					if yIsBasic {
						y = vy.(int16)
					} else {
						y = basic.Int16(vy)
					}
					return func(fr *frame) {
						x := basic.Uintptr(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int32:
					var y int32
					if yIsBasic {
						y = vy.(int32)
					} else {
						y = basic.Int32(vy)
					}
					return func(fr *frame) {
						x := basic.Uintptr(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int64:
					var y int64
					if yIsBasic {
						y = vy.(int64)
					} else {
						y = basic.Int64(vy)
					}
					return func(fr *frame) {
						x := basic.Uintptr(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint:
					var y uint
					if yIsBasic {
						y = vy.(uint)
					} else {
						y = basic.Uint(vy)
					}
					return func(fr *frame) {
						x := basic.Uintptr(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint8:
					var y uint8
					if yIsBasic {
						y = vy.(uint8)
					} else {
						y = basic.Uint8(vy)
					}
					return func(fr *frame) {
						x := basic.Uintptr(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint16:
					var y uint16
					if yIsBasic {
						y = vy.(uint16)
					} else {
						y = basic.Uint16(vy)
					}
					return func(fr *frame) {
						x := basic.Uintptr(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint32:
					var y uint32
					if yIsBasic {
						y = vy.(uint32)
					} else {
						y = basic.Uint32(vy)
					}
					return func(fr *frame) {
						x := basic.Uintptr(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint64:
					var y uint64
					if yIsBasic {
						y = vy.(uint64)
					} else {
						y = basic.Uint64(vy)
					}
					return func(fr *frame) {
						x := basic.Uintptr(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uintptr:
					var y uintptr
					if yIsBasic {
						y = vy.(uintptr)
					} else {
						y = basic.Uintptr(vy)
					}
					return func(fr *frame) {
						x := basic.Uintptr(fr.reg(ix))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				}
			} else {
				switch ykind {
				case reflect.Int:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Uintptr(fr.reg(ix))
							y := fr.reg(iy).(int)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						x := basic.Uintptr(fr.reg(ix))
						y := basic.Int(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int8:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Uintptr(fr.reg(ix))
							y := fr.reg(iy).(int8)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						x := basic.Uintptr(fr.reg(ix))
						y := basic.Int8(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int16:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Uintptr(fr.reg(ix))
							y := fr.reg(iy).(int16)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						x := basic.Uintptr(fr.reg(ix))
						y := basic.Int16(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int32:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Uintptr(fr.reg(ix))
							y := fr.reg(iy).(int32)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						x := basic.Uintptr(fr.reg(ix))
						y := basic.Int32(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Int64:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Uintptr(fr.reg(ix))
							y := fr.reg(iy).(int64)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						x := basic.Uintptr(fr.reg(ix))
						y := basic.Int64(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Uintptr(fr.reg(ix))
							y := fr.reg(iy).(uint)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						x := basic.Uintptr(fr.reg(ix))
						y := basic.Uint(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint8:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Uintptr(fr.reg(ix))
							y := fr.reg(iy).(uint8)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						x := basic.Uintptr(fr.reg(ix))
						y := basic.Uint8(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint16:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Uintptr(fr.reg(ix))
							y := fr.reg(iy).(uint16)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						x := basic.Uintptr(fr.reg(ix))
						y := basic.Uint16(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint32:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Uintptr(fr.reg(ix))
							y := fr.reg(iy).(uint32)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						x := basic.Uintptr(fr.reg(ix))
						y := basic.Uint32(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uint64:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Uintptr(fr.reg(ix))
							y := fr.reg(iy).(uint64)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						x := basic.Uintptr(fr.reg(ix))
						y := basic.Uint64(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				case reflect.Uintptr:
					if yIsBasic {
						return func(fr *frame) {
							x := basic.Uintptr(fr.reg(ix))
							y := fr.reg(iy).(uintptr)
							fr.setReg(ir, basic.Make(t, x>>y))
						}
					}
					return func(fr *frame) {
						x := basic.Uintptr(fr.reg(ix))
						y := basic.Uintptr(fr.reg(iy))
						fr.setReg(ir, basic.Make(t, x>>y))
					}
				}
			}
		}
	}
	panic("unreachable")
}
