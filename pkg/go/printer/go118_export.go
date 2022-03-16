// export by github.com/goplus/gossa/cmd/qexp

//go:build go1.18
// +build go1.18

package printer

import (
	q "go/printer"

	"go/constant"
	"reflect"

	"github.com/goplus/gossa"
)

func init() {
	gossa.RegisterPackage(&gossa.Package{
		Name: "printer",
		Path: "go/printer",
		Deps: map[string]string{
			"bytes":               "bytes",
			"fmt":                 "fmt",
			"go/ast":              "ast",
			"go/build/constraint": "constraint",
			"go/token":            "token",
			"io":                  "io",
			"math":                "math",
			"os":                  "os",
			"sort":                "sort",
			"strconv":             "strconv",
			"strings":             "strings",
			"text/tabwriter":      "tabwriter",
			"unicode":             "unicode",
			"unicode/utf8":        "utf8",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]gossa.NamedType{
			"CommentedNode": {reflect.TypeOf((*q.CommentedNode)(nil)).Elem(), "", ""},
			"Config":        {reflect.TypeOf((*q.Config)(nil)).Elem(), "", "Fprint,fprint"},
			"Mode":          {reflect.TypeOf((*q.Mode)(nil)).Elem(), "", ""},
		},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"Fprint": reflect.ValueOf(q.Fprint),
		},
		TypedConsts: map[string]gossa.TypedConst{
			"RawFormat": {reflect.TypeOf(q.RawFormat), constant.MakeInt64(int64(q.RawFormat))},
			"SourcePos": {reflect.TypeOf(q.SourcePos), constant.MakeInt64(int64(q.SourcePos))},
			"TabIndent": {reflect.TypeOf(q.TabIndent), constant.MakeInt64(int64(q.TabIndent))},
			"UseSpaces": {reflect.TypeOf(q.UseSpaces), constant.MakeInt64(int64(q.UseSpaces))},
		},
		UntypedConsts: map[string]gossa.UntypedConst{},
	})
}
