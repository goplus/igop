// export by github.com/goplus/gossa/cmd/qexp

//go:build go1.16
// +build go1.16

package parse

import (
	"text/template/parse"

	"github.com/goplus/gossa"
)

func init() {
	gossa.RegisterPackage("text/template/parse", extMap_go116, typList_go116)
}

var extMap_go116 = map[string]interface{}{
	"(*text/template/parse.CommentNode).Copy":    (*parse.CommentNode).Copy,
	"(*text/template/parse.CommentNode).String":  (*parse.CommentNode).String,
	"(text/template/parse.CommentNode).Position": (parse.CommentNode).Position,
	"(text/template/parse.CommentNode).Type":     (parse.CommentNode).Type,
}

var typList_go116 = []interface{}{
	(*parse.CommentNode)(nil),
	(*parse.Mode)(nil),
}
