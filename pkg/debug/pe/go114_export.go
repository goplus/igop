// export by github.com/goplus/igop/cmd/qexp

//+build go1.14,!go1.15

package pe

import (
	q "debug/pe"

	"go/constant"
	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "pe",
		Path: "debug/pe",
		Deps: map[string]string{
			"bytes":           "bytes",
			"compress/zlib":   "zlib",
			"debug/dwarf":     "dwarf",
			"encoding/binary": "binary",
			"fmt":             "fmt",
			"io":              "io",
			"os":              "os",
			"strconv":         "strconv",
			"strings":         "strings",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]igop.NamedType{
			"COFFSymbol":       {reflect.TypeOf((*q.COFFSymbol)(nil)).Elem(), "", "FullName"},
			"DataDirectory":    {reflect.TypeOf((*q.DataDirectory)(nil)).Elem(), "", ""},
			"File":             {reflect.TypeOf((*q.File)(nil)).Elem(), "", "Close,DWARF,ImportedLibraries,ImportedSymbols,Section"},
			"FileHeader":       {reflect.TypeOf((*q.FileHeader)(nil)).Elem(), "", ""},
			"FormatError":      {reflect.TypeOf((*q.FormatError)(nil)).Elem(), "", "Error"},
			"ImportDirectory":  {reflect.TypeOf((*q.ImportDirectory)(nil)).Elem(), "", ""},
			"OptionalHeader32": {reflect.TypeOf((*q.OptionalHeader32)(nil)).Elem(), "", ""},
			"OptionalHeader64": {reflect.TypeOf((*q.OptionalHeader64)(nil)).Elem(), "", ""},
			"Reloc":            {reflect.TypeOf((*q.Reloc)(nil)).Elem(), "", ""},
			"Section":          {reflect.TypeOf((*q.Section)(nil)).Elem(), "", "Data,Open"},
			"SectionHeader":    {reflect.TypeOf((*q.SectionHeader)(nil)).Elem(), "", ""},
			"SectionHeader32":  {reflect.TypeOf((*q.SectionHeader32)(nil)).Elem(), "", "fullName"},
			"StringTable":      {reflect.TypeOf((*q.StringTable)(nil)).Elem(), "String", ""},
			"Symbol":           {reflect.TypeOf((*q.Symbol)(nil)).Elem(), "", ""},
		},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"NewFile": reflect.ValueOf(q.NewFile),
			"Open":    reflect.ValueOf(q.Open),
		},
		TypedConsts: map[string]igop.TypedConst{},
		UntypedConsts: map[string]igop.UntypedConst{
			"COFFSymbolSize":                       {"untyped int", constant.MakeInt64(int64(q.COFFSymbolSize))},
			"IMAGE_DIRECTORY_ENTRY_ARCHITECTURE":   {"untyped int", constant.MakeInt64(int64(q.IMAGE_DIRECTORY_ENTRY_ARCHITECTURE))},
			"IMAGE_DIRECTORY_ENTRY_BASERELOC":      {"untyped int", constant.MakeInt64(int64(q.IMAGE_DIRECTORY_ENTRY_BASERELOC))},
			"IMAGE_DIRECTORY_ENTRY_BOUND_IMPORT":   {"untyped int", constant.MakeInt64(int64(q.IMAGE_DIRECTORY_ENTRY_BOUND_IMPORT))},
			"IMAGE_DIRECTORY_ENTRY_COM_DESCRIPTOR": {"untyped int", constant.MakeInt64(int64(q.IMAGE_DIRECTORY_ENTRY_COM_DESCRIPTOR))},
			"IMAGE_DIRECTORY_ENTRY_DEBUG":          {"untyped int", constant.MakeInt64(int64(q.IMAGE_DIRECTORY_ENTRY_DEBUG))},
			"IMAGE_DIRECTORY_ENTRY_DELAY_IMPORT":   {"untyped int", constant.MakeInt64(int64(q.IMAGE_DIRECTORY_ENTRY_DELAY_IMPORT))},
			"IMAGE_DIRECTORY_ENTRY_EXCEPTION":      {"untyped int", constant.MakeInt64(int64(q.IMAGE_DIRECTORY_ENTRY_EXCEPTION))},
			"IMAGE_DIRECTORY_ENTRY_EXPORT":         {"untyped int", constant.MakeInt64(int64(q.IMAGE_DIRECTORY_ENTRY_EXPORT))},
			"IMAGE_DIRECTORY_ENTRY_GLOBALPTR":      {"untyped int", constant.MakeInt64(int64(q.IMAGE_DIRECTORY_ENTRY_GLOBALPTR))},
			"IMAGE_DIRECTORY_ENTRY_IAT":            {"untyped int", constant.MakeInt64(int64(q.IMAGE_DIRECTORY_ENTRY_IAT))},
			"IMAGE_DIRECTORY_ENTRY_IMPORT":         {"untyped int", constant.MakeInt64(int64(q.IMAGE_DIRECTORY_ENTRY_IMPORT))},
			"IMAGE_DIRECTORY_ENTRY_LOAD_CONFIG":    {"untyped int", constant.MakeInt64(int64(q.IMAGE_DIRECTORY_ENTRY_LOAD_CONFIG))},
			"IMAGE_DIRECTORY_ENTRY_RESOURCE":       {"untyped int", constant.MakeInt64(int64(q.IMAGE_DIRECTORY_ENTRY_RESOURCE))},
			"IMAGE_DIRECTORY_ENTRY_SECURITY":       {"untyped int", constant.MakeInt64(int64(q.IMAGE_DIRECTORY_ENTRY_SECURITY))},
			"IMAGE_DIRECTORY_ENTRY_TLS":            {"untyped int", constant.MakeInt64(int64(q.IMAGE_DIRECTORY_ENTRY_TLS))},
			"IMAGE_FILE_MACHINE_AM33":              {"untyped int", constant.MakeInt64(int64(q.IMAGE_FILE_MACHINE_AM33))},
			"IMAGE_FILE_MACHINE_AMD64":             {"untyped int", constant.MakeInt64(int64(q.IMAGE_FILE_MACHINE_AMD64))},
			"IMAGE_FILE_MACHINE_ARM":               {"untyped int", constant.MakeInt64(int64(q.IMAGE_FILE_MACHINE_ARM))},
			"IMAGE_FILE_MACHINE_ARM64":             {"untyped int", constant.MakeInt64(int64(q.IMAGE_FILE_MACHINE_ARM64))},
			"IMAGE_FILE_MACHINE_ARMNT":             {"untyped int", constant.MakeInt64(int64(q.IMAGE_FILE_MACHINE_ARMNT))},
			"IMAGE_FILE_MACHINE_EBC":               {"untyped int", constant.MakeInt64(int64(q.IMAGE_FILE_MACHINE_EBC))},
			"IMAGE_FILE_MACHINE_I386":              {"untyped int", constant.MakeInt64(int64(q.IMAGE_FILE_MACHINE_I386))},
			"IMAGE_FILE_MACHINE_IA64":              {"untyped int", constant.MakeInt64(int64(q.IMAGE_FILE_MACHINE_IA64))},
			"IMAGE_FILE_MACHINE_M32R":              {"untyped int", constant.MakeInt64(int64(q.IMAGE_FILE_MACHINE_M32R))},
			"IMAGE_FILE_MACHINE_MIPS16":            {"untyped int", constant.MakeInt64(int64(q.IMAGE_FILE_MACHINE_MIPS16))},
			"IMAGE_FILE_MACHINE_MIPSFPU":           {"untyped int", constant.MakeInt64(int64(q.IMAGE_FILE_MACHINE_MIPSFPU))},
			"IMAGE_FILE_MACHINE_MIPSFPU16":         {"untyped int", constant.MakeInt64(int64(q.IMAGE_FILE_MACHINE_MIPSFPU16))},
			"IMAGE_FILE_MACHINE_POWERPC":           {"untyped int", constant.MakeInt64(int64(q.IMAGE_FILE_MACHINE_POWERPC))},
			"IMAGE_FILE_MACHINE_POWERPCFP":         {"untyped int", constant.MakeInt64(int64(q.IMAGE_FILE_MACHINE_POWERPCFP))},
			"IMAGE_FILE_MACHINE_R4000":             {"untyped int", constant.MakeInt64(int64(q.IMAGE_FILE_MACHINE_R4000))},
			"IMAGE_FILE_MACHINE_SH3":               {"untyped int", constant.MakeInt64(int64(q.IMAGE_FILE_MACHINE_SH3))},
			"IMAGE_FILE_MACHINE_SH3DSP":            {"untyped int", constant.MakeInt64(int64(q.IMAGE_FILE_MACHINE_SH3DSP))},
			"IMAGE_FILE_MACHINE_SH4":               {"untyped int", constant.MakeInt64(int64(q.IMAGE_FILE_MACHINE_SH4))},
			"IMAGE_FILE_MACHINE_SH5":               {"untyped int", constant.MakeInt64(int64(q.IMAGE_FILE_MACHINE_SH5))},
			"IMAGE_FILE_MACHINE_THUMB":             {"untyped int", constant.MakeInt64(int64(q.IMAGE_FILE_MACHINE_THUMB))},
			"IMAGE_FILE_MACHINE_UNKNOWN":           {"untyped int", constant.MakeInt64(int64(q.IMAGE_FILE_MACHINE_UNKNOWN))},
			"IMAGE_FILE_MACHINE_WCEMIPSV2":         {"untyped int", constant.MakeInt64(int64(q.IMAGE_FILE_MACHINE_WCEMIPSV2))},
		},
	})
}
