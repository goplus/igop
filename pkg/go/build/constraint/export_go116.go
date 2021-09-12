// export by github.com/goplus/interp/cmd/qexp

//go:build go1.16
// +build go1.16

package constraint

import (
	"go/build/constraint"

	"github.com/goplus/interp"
)

func init() {
	interp.RegisterPackage("go/build/constraint", extMap_go116, typList_go116)
}

var extMap_go116 = map[string]interface{}{
	"(*go/build/constraint.AndExpr).Eval":      (*constraint.AndExpr).Eval,
	"(*go/build/constraint.AndExpr).String":    (*constraint.AndExpr).String,
	"(*go/build/constraint.NotExpr).Eval":      (*constraint.NotExpr).Eval,
	"(*go/build/constraint.NotExpr).String":    (*constraint.NotExpr).String,
	"(*go/build/constraint.OrExpr).Eval":       (*constraint.OrExpr).Eval,
	"(*go/build/constraint.OrExpr).String":     (*constraint.OrExpr).String,
	"(*go/build/constraint.SyntaxError).Error": (*constraint.SyntaxError).Error,
	"(*go/build/constraint.TagExpr).Eval":      (*constraint.TagExpr).Eval,
	"(*go/build/constraint.TagExpr).String":    (*constraint.TagExpr).String,
	"(go/build/constraint.Expr).Eval":          (constraint.Expr).Eval,
	"(go/build/constraint.Expr).String":        (constraint.Expr).String,
	"go/build/constraint.IsGoBuild":            constraint.IsGoBuild,
	"go/build/constraint.IsPlusBuild":          constraint.IsPlusBuild,
	"go/build/constraint.Parse":                constraint.Parse,
	"go/build/constraint.PlusBuildLines":       constraint.PlusBuildLines,
}

var typList_go116 = []interface{}{
	(*constraint.AndExpr)(nil),
	(*constraint.Expr)(nil),
	(*constraint.NotExpr)(nil),
	(*constraint.OrExpr)(nil),
	(*constraint.SyntaxError)(nil),
	(*constraint.TagExpr)(nil),
}
