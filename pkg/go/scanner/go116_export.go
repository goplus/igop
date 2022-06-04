// export by github.com/goplus/igop/cmd/qexp

//+build go1.16,!go1.17

package scanner

import (
	q "go/scanner"

	"go/constant"
	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "scanner",
		Path: "go/scanner",
		Deps: map[string]string{
			"bytes":         "bytes",
			"fmt":           "fmt",
			"go/token":      "token",
			"io":            "io",
			"path/filepath": "filepath",
			"sort":          "sort",
			"strconv":       "strconv",
			"unicode":       "unicode",
			"unicode/utf8":  "utf8",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]igop.NamedType{
			"Error":        {reflect.TypeOf((*q.Error)(nil)).Elem(), "Error", ""},
			"ErrorHandler": {reflect.TypeOf((*q.ErrorHandler)(nil)).Elem(), "", ""},
			"ErrorList":    {reflect.TypeOf((*q.ErrorList)(nil)).Elem(), "Err,Error,Len,Less,Sort,Swap", "Add,RemoveMultiples,Reset"},
			"Mode":         {reflect.TypeOf((*q.Mode)(nil)).Elem(), "", ""},
			"Scanner":      {reflect.TypeOf((*q.Scanner)(nil)).Elem(), "", "Init,Scan,digits,error,errorf,findLineEnd,next,peek,scanComment,scanEscape,scanIdentifier,scanNumber,scanRawString,scanRune,scanString,skipWhitespace,switch2,switch3,switch4,updateLineInfo"},
		},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"PrintError": reflect.ValueOf(q.PrintError),
		},
		TypedConsts: map[string]igop.TypedConst{
			"ScanComments": {reflect.TypeOf(q.ScanComments), constant.MakeInt64(int64(q.ScanComments))},
		},
		UntypedConsts: map[string]igop.UntypedConst{},
	})
}
