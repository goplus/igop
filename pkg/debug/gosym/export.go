// export by github.com/goplus/gossa/cmd/qexp

package gosym

import (
	"debug/gosym"

	"github.com/goplus/gossa"
)

func init() {
	gossa.RegisterPackage("debug/gosym", extMap, typList)
}

var extMap = map[string]interface{}{
	"(*debug/gosym.DecodingError).Error":    (*gosym.DecodingError).Error,
	"(*debug/gosym.LineTable).LineToPC":     (*gosym.LineTable).LineToPC,
	"(*debug/gosym.LineTable).PCToLine":     (*gosym.LineTable).PCToLine,
	"(*debug/gosym.Sym).BaseName":           (*gosym.Sym).BaseName,
	"(*debug/gosym.Sym).PackageName":        (*gosym.Sym).PackageName,
	"(*debug/gosym.Sym).ReceiverName":       (*gosym.Sym).ReceiverName,
	"(*debug/gosym.Sym).Static":             (*gosym.Sym).Static,
	"(*debug/gosym.Table).LineToPC":         (*gosym.Table).LineToPC,
	"(*debug/gosym.Table).LookupFunc":       (*gosym.Table).LookupFunc,
	"(*debug/gosym.Table).LookupSym":        (*gosym.Table).LookupSym,
	"(*debug/gosym.Table).PCToFunc":         (*gosym.Table).PCToFunc,
	"(*debug/gosym.Table).PCToLine":         (*gosym.Table).PCToLine,
	"(*debug/gosym.Table).SymByAddr":        (*gosym.Table).SymByAddr,
	"(*debug/gosym.UnknownLineError).Error": (*gosym.UnknownLineError).Error,
	"(debug/gosym.Func).BaseName":           (gosym.Func).BaseName,
	"(debug/gosym.Func).PackageName":        (gosym.Func).PackageName,
	"(debug/gosym.Func).ReceiverName":       (gosym.Func).ReceiverName,
	"(debug/gosym.Func).Static":             (gosym.Func).Static,
	"(debug/gosym.UnknownFileError).Error":  (gosym.UnknownFileError).Error,
	"debug/gosym.NewLineTable":              gosym.NewLineTable,
	"debug/gosym.NewTable":                  gosym.NewTable,
}

var typList = []interface{}{
	(*gosym.DecodingError)(nil),
	(*gosym.Func)(nil),
	(*gosym.LineTable)(nil),
	(*gosym.Obj)(nil),
	(*gosym.Sym)(nil),
	(*gosym.Table)(nil),
	(*gosym.UnknownFileError)(nil),
	(*gosym.UnknownLineError)(nil),
}
