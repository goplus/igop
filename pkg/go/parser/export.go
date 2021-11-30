// export by github.com/goplus/gossa/cmd/qexp

package parser

import (
	q "go/parser"

	"go/constant"
	"reflect"

	"github.com/goplus/gossa"
)

func init() {
	gossa.RegisterPackage(&gossa.Package{
		Name: "parser",
		Path: "go/parser",
		Deps: map[string]string{
			"bytes":         "bytes",
			"errors":        "errors",
			"fmt":           "fmt",
			"go/ast":        "ast",
			"go/scanner":    "scanner",
			"go/token":      "token",
			"io":            "io",
			"io/fs":         "fs",
			"os":            "os",
			"path/filepath": "filepath",
			"strconv":       "strconv",
			"strings":       "strings",
			"unicode":       "unicode",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]gossa.NamedType{
			"Mode": {reflect.TypeOf((*q.Mode)(nil)).Elem(), "", ""},
		},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"ParseDir":      reflect.ValueOf(q.ParseDir),
			"ParseExpr":     reflect.ValueOf(q.ParseExpr),
			"ParseExprFrom": reflect.ValueOf(q.ParseExprFrom),
			"ParseFile":     reflect.ValueOf(q.ParseFile),
		},
		TypedConsts: map[string]gossa.TypedConst{
			"AllErrors":         {reflect.TypeOf(q.AllErrors), constant.MakeInt64(32)},
			"DeclarationErrors": {reflect.TypeOf(q.DeclarationErrors), constant.MakeInt64(16)},
			"ImportsOnly":       {reflect.TypeOf(q.ImportsOnly), constant.MakeInt64(2)},
			"PackageClauseOnly": {reflect.TypeOf(q.PackageClauseOnly), constant.MakeInt64(1)},
			"ParseComments":     {reflect.TypeOf(q.ParseComments), constant.MakeInt64(4)},
			"SpuriousErrors":    {reflect.TypeOf(q.SpuriousErrors), constant.MakeInt64(32)},
			"Trace":             {reflect.TypeOf(q.Trace), constant.MakeInt64(8)},
		},
		UntypedConsts: map[string]gossa.UntypedConst{},
	})
}
