// export by github.com/goplus/igop/cmd/qexp

//go:build go1.17 && !go1.18
// +build go1.17,!go1.18

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
		NamedTypes: map[string]reflect.Type{
			"Error":        reflect.TypeOf((*q.Error)(nil)).Elem(),
			"ErrorHandler": reflect.TypeOf((*q.ErrorHandler)(nil)).Elem(),
			"ErrorList":    reflect.TypeOf((*q.ErrorList)(nil)).Elem(),
			"Mode":         reflect.TypeOf((*q.Mode)(nil)).Elem(),
			"Scanner":      reflect.TypeOf((*q.Scanner)(nil)).Elem(),
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
