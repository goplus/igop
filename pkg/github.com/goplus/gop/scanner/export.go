// export by github.com/goplus/igop/cmd/qexp

package scanner

import (
	q "github.com/goplus/gop/scanner"

	"go/constant"
	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "scanner",
		Path: "github.com/goplus/gop/scanner",
		Deps: map[string]string{
			"bytes":                       "bytes",
			"fmt":                         "fmt",
			"github.com/goplus/gop/token": "token",
			"github.com/qiniu/x/byteutil": "byteutil",
			"go/scanner":                  "scanner",
			"io":                          "io",
			"path/filepath":               "filepath",
			"strconv":                     "strconv",
			"unicode":                     "unicode",
			"unicode/utf8":                "utf8",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{
			"Mode":    reflect.TypeOf((*q.Mode)(nil)).Elem(),
			"Scanner": reflect.TypeOf((*q.Scanner)(nil)).Elem(),
		},
		AliasTypes: map[string]reflect.Type{
			"Error":        reflect.TypeOf((*q.Error)(nil)).Elem(),
			"ErrorHandler": reflect.TypeOf((*q.ErrorHandler)(nil)).Elem(),
			"ErrorList":    reflect.TypeOf((*q.ErrorList)(nil)).Elem(),
		},
		Vars: map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"New":        reflect.ValueOf(q.New),
			"PrintError": reflect.ValueOf(q.PrintError),
		},
		TypedConsts: map[string]igop.TypedConst{
			"ScanComments": {reflect.TypeOf(q.ScanComments), constant.MakeInt64(int64(q.ScanComments))},
		},
		UntypedConsts: map[string]igop.UntypedConst{},
	})
}
