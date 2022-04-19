// export by github.com/goplus/gossa/cmd/qexp

//go:build go1.18
// +build go1.18

package comment

import (
	q "go/doc/comment"

	"reflect"

	"github.com/goplus/gossa"
)

func init() {
	gossa.RegisterPackage(&gossa.Package{
		Name: "comment",
		Path: "go/doc/comment",
		Deps: map[string]string{
			"bytes":        "bytes",
			"fmt":          "fmt",
			"sort":         "sort",
			"strconv":      "strconv",
			"strings":      "strings",
			"unicode":      "unicode",
			"unicode/utf8": "utf8",
		},
		Interfaces: map[string]reflect.Type{
			"Block": reflect.TypeOf((*q.Block)(nil)).Elem(),
			"Text":  reflect.TypeOf((*q.Text)(nil)).Elem(),
		},
		NamedTypes: map[string]gossa.NamedType{
			"Code":      {reflect.TypeOf((*q.Code)(nil)).Elem(), "", "block"},
			"Doc":       {reflect.TypeOf((*q.Doc)(nil)).Elem(), "", ""},
			"DocLink":   {reflect.TypeOf((*q.DocLink)(nil)).Elem(), "", "DefaultURL,text"},
			"Heading":   {reflect.TypeOf((*q.Heading)(nil)).Elem(), "", "DefaultID,block"},
			"Italic":    {reflect.TypeOf((*q.Italic)(nil)).Elem(), "text", ""},
			"Link":      {reflect.TypeOf((*q.Link)(nil)).Elem(), "", "text"},
			"LinkDef":   {reflect.TypeOf((*q.LinkDef)(nil)).Elem(), "", ""},
			"List":      {reflect.TypeOf((*q.List)(nil)).Elem(), "", "BlankBefore,BlankBetween,block"},
			"ListItem":  {reflect.TypeOf((*q.ListItem)(nil)).Elem(), "", ""},
			"Paragraph": {reflect.TypeOf((*q.Paragraph)(nil)).Elem(), "", "block"},
			"Parser":    {reflect.TypeOf((*q.Parser)(nil)).Elem(), "", "Parse"},
			"Plain":     {reflect.TypeOf((*q.Plain)(nil)).Elem(), "text", ""},
			"Printer":   {reflect.TypeOf((*q.Printer)(nil)).Elem(), "", "Comment,HTML,Markdown,Text,docLinkURL,headingID,headingLevel"},
		},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"DefaultLookupPackage": reflect.ValueOf(q.DefaultLookupPackage),
		},
		TypedConsts:   map[string]gossa.TypedConst{},
		UntypedConsts: map[string]gossa.UntypedConst{},
	})
}
