package gossa

import (
	"fmt"
	"go/types"
	"log"
	"reflect"

	"golang.org/x/tools/go/ssa"
)

func checkPackages(intp *Interp, pkgs []*ssa.Package) (err error) {
	defer func() {
		if v := recover(); v != nil {
			err = v.(error)
		}
	}()
	visit := visitor{
		intp: intp,
		prog: intp.prog,
		pkgs: make(map[*ssa.Package]bool),
		seen: make(map[*ssa.Function]bool),
	}
	for _, pkg := range pkgs {
		visit.pkgs[pkg] = true
	}
	visit.program()
	return
}

type visitor struct {
	intp *Interp
	prog *ssa.Program
	pkgs map[*ssa.Package]bool
	seen map[*ssa.Function]bool
}

func (visit *visitor) program() {
	chks := make(map[string]bool)
	chks[""] = true // anonymous struct embbed named type
	for pkg := range visit.pkgs {
		chks[pkg.Pkg.Path()] = true
		for _, mem := range pkg.Members {
			if fn, ok := mem.(*ssa.Function); ok {
				visit.function(fn)
			}
		}
	}
	isExtern := func(typ reflect.Type) bool {
		if typ.Kind() == reflect.Ptr {
			typ = typ.Elem()
		}
		return !chks[typ.PkgPath()]
	}
	for _, T := range visit.prog.RuntimeTypes() {
		typ := visit.intp.preToType(T)
		// skip extern type
		if isExtern(typ) {
			continue
		}
		mmap := make(map[string]*ssa.Function)
		mset := visit.prog.MethodSets.MethodSet(T)
		for i, n := 0, mset.Len(); i < n; i++ {
			sel := mset.At(i)
			obj := sel.Obj()
			// skip embbed extern type method
			if !chks[obj.Pkg().Path()] {
				continue
			}
			fn := visit.prog.MethodValue(sel)
			mmap[obj.Name()] = fn
			visit.function(fn)
		}
		visit.intp.msets[typ] = mmap
	}
}

func (visit *visitor) function(fn *ssa.Function) {
	if !visit.seen[fn] {
		visit.seen[fn] = true
		if fn.Blocks == nil {
			if _, ok := visit.pkgs[fn.Pkg]; ok {
				panic(fmt.Errorf("%v: missing function body", visit.intp.fset.Position(fn.Pos())))
			}
			return
		}
		visit.intp.loadType(fn.Type())
		for _, alloc := range fn.Locals {
			visit.intp.loadType(alloc.Type())
			visit.intp.loadType(deref(alloc.Type()))
		}
		blocks := make(map[*ssa.BasicBlock]*FuncBlock)
		pfn := &Function{
			Interp:           visit.intp,
			Fn:               fn,
			mapUnderscoreKey: make(map[types.Type]bool),
			index:            make(map[ssa.Value]int),
		}
		for _, p := range fn.Params {
			pfn.regIndex(p)
		}
		for _, p := range fn.FreeVars {
			pfn.regIndex(p)
		}
		visit.intp.funcs[fn] = pfn
		var buf [32]*ssa.Value // avoid alloc in common case
		for _, b := range fn.Blocks {
			block := &FuncBlock{Index: b.Index}
			block.Instrs = make([]func(*frame, *int), len(b.Instrs), len(b.Instrs))
			blocks[b] = block
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
					case nil:
						// skip
					default:
						visit.intp.loadType(v.Type())
					}
				}
				ifn := makeInstr(visit.intp, pfn, instr)
				if ifn == nil {
					continue
				}
				if visit.intp.mode&EnableDumpInstr != 0 {
					block.Instrs[index] = func(fr *frame, k *int) {
						if v, ok := instr.(ssa.Value); ok {
							log.Printf("\t%-20T %v = %-40v\t%v\n", instr, v.Name(), instr, v.Type())
						} else {
							log.Printf("\t%-20T %v\n", instr, instr)
						}
						ifn(fr, k)
					}
				} else {
					block.Instrs[index] = ifn
				}
				index++
			}
			block.Instrs = block.Instrs[:index]
		}
		for b, fb := range blocks {
			for _, v := range b.Succs {
				fb.Succs = append(fb.Succs, blocks[v])
			}
		}
		pfn.MainBlock = blocks[fn.Blocks[0]]
		pfn.Recover = blocks[fn.Recover]
	}
}

type constValue struct {
	*ssa.Const
	Value value
}
