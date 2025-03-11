// export by github.com/goplus/igop/cmd/qexp

package scanner

import (
	q "github.com/goplus/gop/tpl/scanner"

	"go/constant"
	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "scanner",
		Path: "github.com/goplus/gop/tpl/scanner",
		Deps: map[string]string{
			"bytes":                           "bytes",
			"fmt":                             "fmt",
			"github.com/goplus/gop/tpl/token": "token",
			"github.com/goplus/gop/tpl/types": "types",
			"go/scanner":                      "scanner",
			"io":                              "io",
			"path/filepath":                   "filepath",
			"strconv":                         "strconv",
			"unicode":                         "unicode",
			"unicode/utf8":                    "utf8",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{
			"ScanErrorHandler": reflect.TypeOf((*q.ScanErrorHandler)(nil)).Elem(),
			"ScanMode":         reflect.TypeOf((*q.ScanMode)(nil)).Elem(),
			"Scanner":          reflect.TypeOf((*q.Scanner)(nil)).Elem(),
		},
		AliasTypes: map[string]reflect.Type{
			"Error":     reflect.TypeOf((*q.Error)(nil)).Elem(),
			"ErrorList": reflect.TypeOf((*q.ErrorList)(nil)).Elem(),
		},
		Vars: map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"PrintError": reflect.ValueOf(q.PrintError),
		},
		TypedConsts: map[string]igop.TypedConst{
			"NoInsertSemis": {reflect.TypeOf(q.NoInsertSemis), constant.MakeInt64(int64(q.NoInsertSemis))},
			"ScanComments":  {reflect.TypeOf(q.ScanComments), constant.MakeInt64(int64(q.ScanComments))},
		},
		UntypedConsts: map[string]igop.UntypedConst{},
	})
}
