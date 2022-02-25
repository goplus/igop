package gossa

import (
	"reflect"

	"golang.org/x/tools/go/ssa"
)

/*
type Op int

const (
	OpInvalid = iota
	// Value-defining instructions
	OpAlloc
	OpPhi
	OpCall
	OpBinOp
	OpUnOp
	OpChangeType
	OpConvert
	OpChangeInterface
	OpSliceToArrayPointer
	OpMakeInterface
	OpMakeClosure
	OpMakeMap
	OpMakeChan
	OpMakeSlice
	OpSlice
	OpFieldAddr
	OpField
	OpIndexAddr
	OpIndex
	OpLookup
	OpSelect
	OpRange
	OpNext
	OpTypeAssert
	OpExtract
	// Instructions executed for effect
	OpJump
	OpIf
	OpReturn
	OpRunDefers
	OpPanic
	OpGo
	OpDefer
	OpSend
	OpStore
	OpMapUpdate
	OpDebugRef
)
*/

type Instr interface {
	ssa.Instruction
	exec(fr *callFrame) int
}

type callFrame struct {
	env    map[Instr]value
	locals map[Instr]reflect.Value
}

type opAlloc struct {
	*ssa.Alloc
	Type reflect.Type
}

func (instr *opAlloc) exec(fr *callFrame) int {
	if instr.Heap {
		fr.env[instr] = reflect.New(instr.Type).Interface()
	} else {
		v := reflect.ValueOf(fr.env[instr])
		SetValue(v.Elem(), fr.locals[instr])
	}
	return kNext
}
