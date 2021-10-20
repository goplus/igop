// export by github.com/goplus/gossa/cmd/qexp

//go:build go1.16
// +build go1.16

package template

import (
	"html/template"

	"github.com/goplus/gossa"
)

func init() {
	gossa.RegisterPackage("html/template", extMap_go116, typList_go116)
}

var extMap_go116 = map[string]interface{}{
	"(*html/template.Template).ParseFS": (*template.Template).ParseFS,
	"html/template.ParseFS":             template.ParseFS,
}

var typList_go116 = []interface{}{}
