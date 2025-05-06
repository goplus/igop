// export by github.com/goplus/igop/cmd/qexp

//go:build go1.23 && !go1.24
// +build go1.23,!go1.24

package template

import (
	q "text/template"

	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "template",
		Path: "text/template",
		Deps: map[string]string{
			"errors":              "errors",
			"fmt":                 "fmt",
			"internal/fmtsort":    "fmtsort",
			"io":                  "io",
			"io/fs":               "fs",
			"net/url":             "url",
			"os":                  "os",
			"path":                "path",
			"path/filepath":       "filepath",
			"reflect":             "reflect",
			"runtime":             "runtime",
			"strings":             "strings",
			"sync":                "sync",
			"text/template/parse": "parse",
			"unicode":             "unicode",
			"unicode/utf8":        "utf8",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{
			"ExecError": reflect.TypeOf((*q.ExecError)(nil)).Elem(),
			"FuncMap":   reflect.TypeOf((*q.FuncMap)(nil)).Elem(),
			"Template":  reflect.TypeOf((*q.Template)(nil)).Elem(),
		},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"HTMLEscape":       reflect.ValueOf(q.HTMLEscape),
			"HTMLEscapeString": reflect.ValueOf(q.HTMLEscapeString),
			"HTMLEscaper":      reflect.ValueOf(q.HTMLEscaper),
			"IsTrue":           reflect.ValueOf(q.IsTrue),
			"JSEscape":         reflect.ValueOf(q.JSEscape),
			"JSEscapeString":   reflect.ValueOf(q.JSEscapeString),
			"JSEscaper":        reflect.ValueOf(q.JSEscaper),
			"Must":             reflect.ValueOf(q.Must),
			"New":              reflect.ValueOf(q.New),
			"ParseFS":          reflect.ValueOf(q.ParseFS),
			"ParseFiles":       reflect.ValueOf(q.ParseFiles),
			"ParseGlob":        reflect.ValueOf(q.ParseGlob),
			"URLQueryEscaper":  reflect.ValueOf(q.URLQueryEscaper),
		},
		TypedConsts:   map[string]igop.TypedConst{},
		UntypedConsts: map[string]igop.UntypedConst{},
	})
}
