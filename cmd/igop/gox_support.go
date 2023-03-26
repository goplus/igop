//go:build go1.18
// +build go1.18

package main

import (
	"go/ast"
	"go/constant"
	"go/token"
	"go/types"
	_ "unsafe"

	"github.com/goplus/igop"
)

type operandMode byte
type builtinId int

type operand struct {
	mode operandMode
	expr ast.Expr
	typ  types.Type
	val  constant.Value
	id   builtinId
}

type positioner interface {
	Pos() token.Pos
}

//go:linkname checker_infer go/types.(*Checker).infer
func checker_infer(check *types.Checker, posn positioner, tparams []*types.TypeParam, targs []types.Type, params *types.Tuple, args []*operand) (result []types.Type)

func init() {
	// support github.com/goplus/gox.checker_infer
	igop.RegisterExternal("go/types.(*Checker).infer", checker_infer)
}
