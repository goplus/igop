// export by github.com/goplus/igop/cmd/qexp

package printer

import (
	q "github.com/goplus/gop/printer"

	"go/constant"
	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "printer",
		Path: "github.com/goplus/gop/printer",
		Deps: map[string]string{
			"bytes":                       "bytes",
			"fmt":                         "fmt",
			"github.com/goplus/gop/ast":   "ast",
			"github.com/goplus/gop/token": "token",
			"io":                          "io",
			"log":                         "log",
			"math":                        "math",
			"os":                          "os",
			"strconv":                     "strconv",
			"strings":                     "strings",
			"text/tabwriter":              "tabwriter",
			"unicode":                     "unicode",
			"unicode/utf8":                "utf8",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{
			"CommentedNode": reflect.TypeOf((*q.CommentedNode)(nil)).Elem(),
			"Config":        reflect.TypeOf((*q.Config)(nil)).Elem(),
			"Mode":          reflect.TypeOf((*q.Mode)(nil)).Elem(),
			"NewlineStmt":   reflect.TypeOf((*q.NewlineStmt)(nil)).Elem(),
		},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"Fprint":   reflect.ValueOf(q.Fprint),
			"SetDebug": reflect.ValueOf(q.SetDebug),
		},
		TypedConsts: map[string]igop.TypedConst{
			"RawFormat": {reflect.TypeOf(q.RawFormat), constant.MakeInt64(int64(q.RawFormat))},
			"SourcePos": {reflect.TypeOf(q.SourcePos), constant.MakeInt64(int64(q.SourcePos))},
			"TabIndent": {reflect.TypeOf(q.TabIndent), constant.MakeInt64(int64(q.TabIndent))},
			"UseSpaces": {reflect.TypeOf(q.UseSpaces), constant.MakeInt64(int64(q.UseSpaces))},
		},
		UntypedConsts: map[string]igop.UntypedConst{
			"DbgFlagAll": {"untyped int", constant.MakeInt64(int64(q.DbgFlagAll))},
		},
	})
}
