// export by github.com/goplus/igop/cmd/qexp

//go:build go1.17 && !go1.18
// +build go1.17,!go1.18

package gosym

import (
	q "debug/gosym"

	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "gosym",
		Path: "debug/gosym",
		Deps: map[string]string{
			"bytes":           "bytes",
			"encoding/binary": "binary",
			"fmt":             "fmt",
			"strconv":         "strconv",
			"strings":         "strings",
			"sync":            "sync",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]igop.NamedType{
			"DecodingError":    {reflect.TypeOf((*q.DecodingError)(nil)).Elem(), "", "Error"},
			"Func":             {reflect.TypeOf((*q.Func)(nil)).Elem(), "", ""},
			"LineTable":        {reflect.TypeOf((*q.LineTable)(nil)).Elem(), "", "LineToPC,PCToLine,findFileLine,findFunc,funcName,go12Funcs,go12LineToPC,go12MapFiles,go12PCToFile,go12PCToLine,initFileMap,isGo12,parse,parsePclnTab,pcvalue,readvarint,slice,step,string,stringFrom,uintptr"},
			"Obj":              {reflect.TypeOf((*q.Obj)(nil)).Elem(), "", "alineFromLine,lineFromAline"},
			"Sym":              {reflect.TypeOf((*q.Sym)(nil)).Elem(), "", "BaseName,PackageName,ReceiverName,Static"},
			"Table":            {reflect.TypeOf((*q.Table)(nil)).Elem(), "", "LineToPC,LookupFunc,LookupSym,PCToFunc,PCToLine,SymByAddr"},
			"UnknownFileError": {reflect.TypeOf((*q.UnknownFileError)(nil)).Elem(), "Error", ""},
			"UnknownLineError": {reflect.TypeOf((*q.UnknownLineError)(nil)).Elem(), "", "Error"},
		},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"NewLineTable": reflect.ValueOf(q.NewLineTable),
			"NewTable":     reflect.ValueOf(q.NewTable),
		},
		TypedConsts:   map[string]igop.TypedConst{},
		UntypedConsts: map[string]igop.UntypedConst{},
	})
}
