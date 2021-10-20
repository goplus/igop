// export by github.com/goplus/gossa/cmd/qexp

package pe

import (
	"debug/pe"

	"github.com/goplus/gossa"
)

func init() {
	gossa.RegisterPackage("debug/pe", extMap, typList)
}

var extMap = map[string]interface{}{
	"(*debug/pe.COFFSymbol).FullName":    (*pe.COFFSymbol).FullName,
	"(*debug/pe.File).Close":             (*pe.File).Close,
	"(*debug/pe.File).DWARF":             (*pe.File).DWARF,
	"(*debug/pe.File).ImportedLibraries": (*pe.File).ImportedLibraries,
	"(*debug/pe.File).ImportedSymbols":   (*pe.File).ImportedSymbols,
	"(*debug/pe.File).Section":           (*pe.File).Section,
	"(*debug/pe.FormatError).Error":      (*pe.FormatError).Error,
	"(*debug/pe.Section).Data":           (*pe.Section).Data,
	"(*debug/pe.Section).Open":           (*pe.Section).Open,
	"(debug/pe.Section).ReadAt":          (pe.Section).ReadAt,
	"(debug/pe.StringTable).String":      (pe.StringTable).String,
	"debug/pe.NewFile":                   pe.NewFile,
	"debug/pe.Open":                      pe.Open,
}

var typList = []interface{}{
	(*pe.COFFSymbol)(nil),
	(*pe.DataDirectory)(nil),
	(*pe.File)(nil),
	(*pe.FileHeader)(nil),
	(*pe.FormatError)(nil),
	(*pe.ImportDirectory)(nil),
	(*pe.OptionalHeader32)(nil),
	(*pe.OptionalHeader64)(nil),
	(*pe.Reloc)(nil),
	(*pe.Section)(nil),
	(*pe.SectionHeader)(nil),
	(*pe.SectionHeader32)(nil),
	(*pe.StringTable)(nil),
	(*pe.Symbol)(nil),
}
