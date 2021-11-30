// export by github.com/goplus/gossa/cmd/qexp

package scanner

import (
	q "go/scanner"

	"go/constant"
	"reflect"

	"github.com/goplus/gossa"
)

func init() {
	gossa.RegisterPackage(&gossa.Package{
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
		NamedTypes: map[string]gossa.NamedType{
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
		TypedConsts: map[string]gossa.TypedConst{
			"ScanComments": {reflect.TypeOf(q.ScanComments), constant.MakeInt64(1)},
		},
		UntypedConsts: map[string]gossa.UntypedConst{},
	})
}
