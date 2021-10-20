// export by github.com/goplus/gossa/cmd/qexp

package doc

import (
	"go/doc"

	"github.com/goplus/gossa"
)

func init() {
	gossa.RegisterPackage("go/doc", extMap, typList)
}

var extMap = map[string]interface{}{
	"(*go/doc.Package).Filter": (*doc.Package).Filter,
	"go/doc.Examples":          doc.Examples,
	"go/doc.IllegalPrefixes":   &doc.IllegalPrefixes,
	"go/doc.IsPredeclared":     doc.IsPredeclared,
	"go/doc.New":               doc.New,
	"go/doc.NewFromFiles":      doc.NewFromFiles,
	"go/doc.Synopsis":          doc.Synopsis,
	"go/doc.ToHTML":            doc.ToHTML,
	"go/doc.ToText":            doc.ToText,
}

var typList = []interface{}{
	(*doc.Example)(nil),
	(*doc.Filter)(nil),
	(*doc.Func)(nil),
	(*doc.Mode)(nil),
	(*doc.Note)(nil),
	(*doc.Package)(nil),
	(*doc.Type)(nil),
	(*doc.Value)(nil),
}
