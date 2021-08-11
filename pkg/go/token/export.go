// export by github.com/goplus/interp/cmd/qexp

package token

import (
	"go/token"

	"github.com/goplus/interp"
)

func init() {
	interp.RegisterPackage("go/token", extMap, typList)
}

var extMap = map[string]interface{}{
	"(*go/token.File).AddLine":            (*token.File).AddLine,
	"(*go/token.File).AddLineColumnInfo":  (*token.File).AddLineColumnInfo,
	"(*go/token.File).AddLineInfo":        (*token.File).AddLineInfo,
	"(*go/token.File).Base":               (*token.File).Base,
	"(*go/token.File).Line":               (*token.File).Line,
	"(*go/token.File).LineCount":          (*token.File).LineCount,
	"(*go/token.File).LineStart":          (*token.File).LineStart,
	"(*go/token.File).MergeLine":          (*token.File).MergeLine,
	"(*go/token.File).Name":               (*token.File).Name,
	"(*go/token.File).Offset":             (*token.File).Offset,
	"(*go/token.File).Pos":                (*token.File).Pos,
	"(*go/token.File).Position":           (*token.File).Position,
	"(*go/token.File).PositionFor":        (*token.File).PositionFor,
	"(*go/token.File).SetLines":           (*token.File).SetLines,
	"(*go/token.File).SetLinesForContent": (*token.File).SetLinesForContent,
	"(*go/token.File).Size":               (*token.File).Size,
	"(*go/token.FileSet).AddFile":         (*token.FileSet).AddFile,
	"(*go/token.FileSet).Base":            (*token.FileSet).Base,
	"(*go/token.FileSet).File":            (*token.FileSet).File,
	"(*go/token.FileSet).Iterate":         (*token.FileSet).Iterate,
	"(*go/token.FileSet).Position":        (*token.FileSet).Position,
	"(*go/token.FileSet).PositionFor":     (*token.FileSet).PositionFor,
	"(*go/token.FileSet).Read":            (*token.FileSet).Read,
	"(*go/token.FileSet).Write":           (*token.FileSet).Write,
	"(*go/token.Position).IsValid":        (*token.Position).IsValid,
	"(go/token.Pos).IsValid":              (token.Pos).IsValid,
	"(go/token.Position).String":          (token.Position).String,
	"(go/token.Token).IsKeyword":          (token.Token).IsKeyword,
	"(go/token.Token).IsLiteral":          (token.Token).IsLiteral,
	"(go/token.Token).IsOperator":         (token.Token).IsOperator,
	"(go/token.Token).Precedence":         (token.Token).Precedence,
	"(go/token.Token).String":             (token.Token).String,
	"go/token.IsExported":                 token.IsExported,
	"go/token.IsIdentifier":               token.IsIdentifier,
	"go/token.IsKeyword":                  token.IsKeyword,
	"go/token.Lookup":                     token.Lookup,
	"go/token.NewFileSet":                 token.NewFileSet,
}

var typList = []interface{}{
	(*token.File)(nil),
	(*token.FileSet)(nil),
	(*token.Pos)(nil),
	(*token.Position)(nil),
	(*token.Token)(nil),
}
