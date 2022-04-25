package gossa

import (
	"github.com/goplus/gossa/basic"
	"golang.org/x/tools/go/ssa"
)

func makeUnOpNOT(pfn *function, instr *ssa.UnOp) func(fr *frame) {
	ir := pfn.regIndex(instr)
	ix, kx, vx := pfn.regIndex3(instr.X)
	typ := pfn.Interp.preToType(instr.Type())
	if typ.PkgPath() == "" {
		if kx == kindConst {
			v := !vx.(bool)
			return func(fr *frame) {
				fr.setReg(ir, v)
			}
		}
		return func(fr *frame) {
			v := !fr.reg(ix).(bool)
			fr.setReg(ir, v)
		}
	} else {
		if kx == kindConst {
			v := basic.Not(vx)
			return func(fr *frame) {
				fr.setReg(ir, v)
			}
		}
		return func(fr *frame) {
			v := basic.Not(fr.reg(ix))
			fr.setReg(ir, v)
		}
	}
}
