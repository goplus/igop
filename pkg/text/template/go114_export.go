// export by github.com/goplus/igop/cmd/qexp

//+build go1.14,!go1.15

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
		NamedTypes: map[string]igop.NamedType{
			"ExecError": {reflect.TypeOf((*q.ExecError)(nil)).Elem(), "Error,Unwrap", ""},
			"FuncMap":   {reflect.TypeOf((*q.FuncMap)(nil)).Elem(), "", ""},
			"Template":  {reflect.TypeOf((*q.Template)(nil)).Elem(), "", "AddParseTree,Clone,DefinedTemplates,Delims,Execute,ExecuteTemplate,Funcs,Lookup,Name,New,Option,Parse,ParseFiles,ParseGlob,Templates,associate,copy,execute,init,setOption"},
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
		TypedConsts:   map[string]igop.TypedConst{},
		UntypedConsts: map[string]igop.UntypedConst{},
	})
}
