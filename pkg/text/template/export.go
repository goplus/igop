// export by github.com/goplus/interp/cmd/qexp

package template

import (
	"text/template"

	"github.com/goplus/interp"
)

func init() {
	interp.RegisterPackage("text/template", extMap, typList)
}

var extMap = map[string]interface{}{
	"(*text/template.Template).AddParseTree":     (*template.Template).AddParseTree,
	"(*text/template.Template).Clone":            (*template.Template).Clone,
	"(*text/template.Template).DefinedTemplates": (*template.Template).DefinedTemplates,
	"(*text/template.Template).Delims":           (*template.Template).Delims,
	"(*text/template.Template).Execute":          (*template.Template).Execute,
	"(*text/template.Template).ExecuteTemplate":  (*template.Template).ExecuteTemplate,
	"(*text/template.Template).Funcs":            (*template.Template).Funcs,
	"(*text/template.Template).Lookup":           (*template.Template).Lookup,
	"(*text/template.Template).Name":             (*template.Template).Name,
	"(*text/template.Template).New":              (*template.Template).New,
	"(*text/template.Template).Option":           (*template.Template).Option,
	"(*text/template.Template).Parse":            (*template.Template).Parse,
	"(*text/template.Template).ParseFiles":       (*template.Template).ParseFiles,
	"(*text/template.Template).ParseGlob":        (*template.Template).ParseGlob,
	"(*text/template.Template).Templates":        (*template.Template).Templates,
	"(text/template.ExecError).Error":            (template.ExecError).Error,
	"(text/template.ExecError).Unwrap":           (template.ExecError).Unwrap,
	"(text/template.Template).Copy":              (template.Template).Copy,
	"(text/template.Template).ErrorContext":      (template.Template).ErrorContext,
	"text/template.HTMLEscape":                   template.HTMLEscape,
	"text/template.HTMLEscapeString":             template.HTMLEscapeString,
	"text/template.HTMLEscaper":                  template.HTMLEscaper,
	"text/template.IsTrue":                       template.IsTrue,
	"text/template.JSEscape":                     template.JSEscape,
	"text/template.JSEscapeString":               template.JSEscapeString,
	"text/template.JSEscaper":                    template.JSEscaper,
	"text/template.Must":                         template.Must,
	"text/template.New":                          template.New,
	"text/template.ParseFiles":                   template.ParseFiles,
	"text/template.ParseGlob":                    template.ParseGlob,
	"text/template.URLQueryEscaper":              template.URLQueryEscaper,
}

var typList = []interface{}{
	(*template.ExecError)(nil),
	(*template.FuncMap)(nil),
	(*template.Template)(nil),
}
