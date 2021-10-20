// export by github.com/goplus/gossa/cmd/qexp

package template

import (
	"html/template"

	"github.com/goplus/gossa"
)

func init() {
	gossa.RegisterPackage("html/template", extMap, typList)
}

var extMap = map[string]interface{}{
	"(*html/template.Error).Error":               (*template.Error).Error,
	"(*html/template.Template).AddParseTree":     (*template.Template).AddParseTree,
	"(*html/template.Template).Clone":            (*template.Template).Clone,
	"(*html/template.Template).DefinedTemplates": (*template.Template).DefinedTemplates,
	"(*html/template.Template).Delims":           (*template.Template).Delims,
	"(*html/template.Template).Execute":          (*template.Template).Execute,
	"(*html/template.Template).ExecuteTemplate":  (*template.Template).ExecuteTemplate,
	"(*html/template.Template).Funcs":            (*template.Template).Funcs,
	"(*html/template.Template).Lookup":           (*template.Template).Lookup,
	"(*html/template.Template).Name":             (*template.Template).Name,
	"(*html/template.Template).New":              (*template.Template).New,
	"(*html/template.Template).Option":           (*template.Template).Option,
	"(*html/template.Template).Parse":            (*template.Template).Parse,
	"(*html/template.Template).ParseFiles":       (*template.Template).ParseFiles,
	"(*html/template.Template).ParseGlob":        (*template.Template).ParseGlob,
	"(*html/template.Template).Templates":        (*template.Template).Templates,
	"html/template.HTMLEscape":                   template.HTMLEscape,
	"html/template.HTMLEscapeString":             template.HTMLEscapeString,
	"html/template.HTMLEscaper":                  template.HTMLEscaper,
	"html/template.IsTrue":                       template.IsTrue,
	"html/template.JSEscape":                     template.JSEscape,
	"html/template.JSEscapeString":               template.JSEscapeString,
	"html/template.JSEscaper":                    template.JSEscaper,
	"html/template.Must":                         template.Must,
	"html/template.New":                          template.New,
	"html/template.ParseFiles":                   template.ParseFiles,
	"html/template.ParseGlob":                    template.ParseGlob,
	"html/template.URLQueryEscaper":              template.URLQueryEscaper,
}

var typList = []interface{}{
	(*template.CSS)(nil),
	(*template.Error)(nil),
	(*template.ErrorCode)(nil),
	(*template.FuncMap)(nil),
	(*template.HTML)(nil),
	(*template.HTMLAttr)(nil),
	(*template.JS)(nil),
	(*template.JSStr)(nil),
	(*template.Srcset)(nil),
	(*template.Template)(nil),
	(*template.URL)(nil),
}
