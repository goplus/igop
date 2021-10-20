// export by github.com/goplus/gossa/cmd/qexp

package parser

import (
	"go/parser"

	"github.com/goplus/gossa"
)

func init() {
	gossa.RegisterPackage("go/parser", extMap, typList)
}

var extMap = map[string]interface{}{
	"go/parser.ParseDir":      parser.ParseDir,
	"go/parser.ParseExpr":     parser.ParseExpr,
	"go/parser.ParseExprFrom": parser.ParseExprFrom,
	"go/parser.ParseFile":     parser.ParseFile,
}

var typList = []interface{}{
	(*parser.Mode)(nil),
}
