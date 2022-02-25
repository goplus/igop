package gossa

import (
	"golang.org/x/tools/go/ssa"
)

func checkPackages(intp *Interp, pkgs []*ssa.Package) {
	visit := visitor{
		intp: intp,
		prog: intp.prog,
		pkgs: pkgs,
		seen: make(map[*ssa.Function]bool),
	}
	visit.program()
}

type visitor struct {
	intp *Interp
	prog *ssa.Program
	pkgs []*ssa.Package
	seen map[*ssa.Function]bool
}

func (visit *visitor) program() {
	for _, pkg := range visit.pkgs {
		for _, mem := range pkg.Members {
			if fn, ok := mem.(*ssa.Function); ok {
				visit.function(fn)
			}
		}
	}
	for _, T := range visit.prog.RuntimeTypes() {
		mset := visit.prog.MethodSets.MethodSet(T)
		for i, n := 0, mset.Len(); i < n; i++ {
			visit.function(visit.prog.MethodValue(mset.At(i)))
		}
	}
}

func (visit *visitor) function(fn *ssa.Function) {
	if !visit.seen[fn] {
		visit.seen[fn] = true
		visit.intp.loadType(fn.Type())
		for _, alloc := range fn.Locals {
			visit.intp.loadType(alloc.Type())
			visit.intp.loadType(deref(alloc.Type()))
		}
		var buf [32]*ssa.Value // avoid alloc in common case
		for _, b := range fn.Blocks {
			for _, instr := range b.Instrs {
				ops := instr.Operands(buf[:0])
				switch instr := instr.(type) {
				case *ssa.Alloc:
					visit.intp.loadType(instr.Type())
					visit.intp.loadType(deref(instr.Type()))
				case *ssa.Next:
					// skip *ssa.opaqueType: iter
					ops = nil
				case *ssa.Extract:
					// skip
					ops = nil
				case *ssa.TypeAssert:
					visit.intp.loadType(instr.AssertedType)
				case *ssa.MakeChan:
					visit.intp.loadType(instr.Type())
				case *ssa.MakeMap:
					visit.intp.loadType(instr.Type())
				case *ssa.MakeSlice:
					visit.intp.loadType(instr.Type())
				case *ssa.SliceToArrayPointer:
					visit.intp.loadType(instr.Type())
				case *ssa.Convert:
					visit.intp.loadType(instr.Type())
				case *ssa.ChangeType:
					visit.intp.loadType(instr.Type())
				case *ssa.MakeInterface:
					visit.intp.loadType(instr.Type())
				}
				for _, op := range ops {
					switch v := (*op).(type) {
					case *ssa.Function:
						visit.function(v)
					case *ssa.Const:
						*op = &constValue{v, constToValue(visit.intp, v)}
					case nil:
						// skip
					default:
						visit.intp.loadType(v.Type())
					}
				}
			}
		}
	}
}

type constValue struct {
	*ssa.Const
	Value value
}
