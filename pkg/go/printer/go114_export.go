// export by github.com/goplus/ixgo/cmd/qexp

//+build go1.14,!go1.15

package printer

import (
	q "go/printer"

	"go/constant"
	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage(&ixgo.Package{
		Name: "printer",
		Path: "go/printer",
		Deps: map[string]string{
			"bytes":          "bytes",
			"fmt":            "fmt",
			"go/ast":         "ast",
			"go/token":       "token",
			"io":             "io",
			"math":           "math",
			"os":             "os",
			"strconv":        "strconv",
			"strings":        "strings",
			"text/tabwriter": "tabwriter",
			"unicode":        "unicode",
			"unicode/utf8":   "utf8",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{
			"CommentedNode": reflect.TypeOf((*q.CommentedNode)(nil)).Elem(),
			"Config":        reflect.TypeOf((*q.Config)(nil)).Elem(),
			"Mode":          reflect.TypeOf((*q.Mode)(nil)).Elem(),
		},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"Fprint": reflect.ValueOf(q.Fprint),
		},
		TypedConsts: map[string]ixgo.TypedConst{
			"RawFormat": {reflect.TypeOf(q.RawFormat), constant.MakeInt64(int64(q.RawFormat))},
			"SourcePos": {reflect.TypeOf(q.SourcePos), constant.MakeInt64(int64(q.SourcePos))},
			"TabIndent": {reflect.TypeOf(q.TabIndent), constant.MakeInt64(int64(q.TabIndent))},
			"UseSpaces": {reflect.TypeOf(q.UseSpaces), constant.MakeInt64(int64(q.UseSpaces))},
		},
		UntypedConsts: map[string]ixgo.UntypedConst{},
	})
}
