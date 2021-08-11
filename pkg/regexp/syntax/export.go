// export by github.com/goplus/interp/cmd/qexp

package syntax

import (
	"regexp/syntax"

	"github.com/goplus/interp"
)

func init() {
	interp.RegisterPackage("regexp/syntax", extMap, typList)
}

var extMap = map[string]interface{}{
	"(*regexp/syntax.Error).Error":          (*syntax.Error).Error,
	"(*regexp/syntax.Inst).MatchEmptyWidth": (*syntax.Inst).MatchEmptyWidth,
	"(*regexp/syntax.Inst).MatchRune":       (*syntax.Inst).MatchRune,
	"(*regexp/syntax.Inst).MatchRunePos":    (*syntax.Inst).MatchRunePos,
	"(*regexp/syntax.Inst).String":          (*syntax.Inst).String,
	"(*regexp/syntax.Prog).Prefix":          (*syntax.Prog).Prefix,
	"(*regexp/syntax.Prog).StartCond":       (*syntax.Prog).StartCond,
	"(*regexp/syntax.Prog).String":          (*syntax.Prog).String,
	"(*regexp/syntax.Regexp).CapNames":      (*syntax.Regexp).CapNames,
	"(*regexp/syntax.Regexp).Equal":         (*syntax.Regexp).Equal,
	"(*regexp/syntax.Regexp).MaxCap":        (*syntax.Regexp).MaxCap,
	"(*regexp/syntax.Regexp).Simplify":      (*syntax.Regexp).Simplify,
	"(*regexp/syntax.Regexp).String":        (*syntax.Regexp).String,
	"(regexp/syntax.ErrorCode).String":      (syntax.ErrorCode).String,
	"(regexp/syntax.InstOp).String":         (syntax.InstOp).String,
	"(regexp/syntax.Op).String":             (syntax.Op).String,
	"regexp/syntax.Compile":                 syntax.Compile,
	"regexp/syntax.EmptyOpContext":          syntax.EmptyOpContext,
	"regexp/syntax.IsWordChar":              syntax.IsWordChar,
	"regexp/syntax.Parse":                   syntax.Parse,
}

var typList = []interface{}{
	(*syntax.EmptyOp)(nil),
	(*syntax.Error)(nil),
	(*syntax.ErrorCode)(nil),
	(*syntax.Flags)(nil),
	(*syntax.Inst)(nil),
	(*syntax.InstOp)(nil),
	(*syntax.Op)(nil),
	(*syntax.Prog)(nil),
	(*syntax.Regexp)(nil),
}
