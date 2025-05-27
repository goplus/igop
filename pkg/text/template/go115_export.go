// export by github.com/goplus/ixgo/cmd/qexp

//+build go1.15,!go1.16

package template

import (
	q "text/template"

	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage(&ixgo.Package{
		Name: "template",
		Path: "text/template",
		Deps: map[string]string{
			"bytes":               "bytes",
			"errors":              "errors",
			"fmt":                 "fmt",
			"internal/fmtsort":    "fmtsort",
			"io":                  "io",
			"io/ioutil":           "ioutil",
			"net/url":             "url",
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
			"ParseFiles":       reflect.ValueOf(q.ParseFiles),
			"ParseGlob":        reflect.ValueOf(q.ParseGlob),
			"URLQueryEscaper":  reflect.ValueOf(q.URLQueryEscaper),
		},
		TypedConsts:   map[string]ixgo.TypedConst{},
		UntypedConsts: map[string]ixgo.UntypedConst{},
	})
}
