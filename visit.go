package gossa

import (
	"fmt"
	"log"

	"golang.org/x/tools/go/ssa"
)

func checkPackages(intp *Interp, pkgs []*ssa.Package) (err error) {
	visit := visitor{
		intp: intp,
		prog: intp.prog,
		pkgs: make(map[*ssa.Package]bool),
		seen: make(map[*ssa.Function]bool),
	}
	for _, pkg := range pkgs {
		visit.pkgs[pkg] = true
	}
	return visit.program()
}

type visitor struct {
	intp *Interp
	prog *ssa.Program
	pkgs map[*ssa.Package]bool
	seen map[*ssa.Function]bool
}

func (visit *visitor) program() error {
	for pkg := range visit.pkgs {
		for _, mem := range pkg.Members {
			if fn, ok := mem.(*ssa.Function); ok {
				if err := visit.function(fn); err != nil {
					return err
				}
			}
		}
	}
	for _, T := range visit.prog.RuntimeTypes() {
		mset := visit.prog.MethodSets.MethodSet(T)
		for i, n := 0, mset.Len(); i < n; i++ {
			if err := visit.function(visit.prog.MethodValue(mset.At(i))); err != nil {
				return err
			}
		}
	}
	return nil
}

func (visit *visitor) function(fn *ssa.Function) error {
	if !visit.seen[fn] {
		visit.seen[fn] = true
		if fn.Blocks == nil && fn.Pkg != nil {
			if _, ok := visit.pkgs[fn.Pkg]; ok {
				return fmt.Errorf("%v: missing function body", visit.intp.fset.Position(fn.Pos()))
			}
		}
		visit.intp.loadType(fn.Type())
		for _, alloc := range fn.Locals {
			visit.intp.loadType(alloc.Type())
			visit.intp.loadType(deref(alloc.Type()))
		}
		var buf [32]*ssa.Value // avoid alloc in common case
		for _, b := range fn.Blocks {
			block := &FuncBlock{}
			block.Instrs = make([]func(*frame, *int), len(b.Instrs), len(b.Instrs))
			visit.intp.blocks[b] = block
			var index int
			for i := 0; i < len(b.Instrs); i++ {
				instr := b.Instrs[i]
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
				fn := makeInstr(visit.intp, instr)
				if fn == nil {
					continue
				}
				if visit.intp.mode&EnableDumpInstr != 0 {
					block.Instrs[index] = func(fr *frame, k *int) {
						if v, ok := instr.(ssa.Value); ok {
							log.Printf("\t%-20T %v = %v\n", instr, v.Name(), instr)
						} else {
							log.Printf("\t%-20T %v\n", instr, instr)
						}
						fn(fr, k)
					}
				} else {
					block.Instrs[index] = fn
				}
				index++
			}
			block.Instrs = block.Instrs[:index]
		}
	}
	return nil
}

type constValue struct {
	*ssa.Const
	Value value
}
