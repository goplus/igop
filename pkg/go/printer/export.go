// export by github.com/goplus/gossa/cmd/qexp

package printer

import (
	"go/printer"

	"github.com/goplus/gossa"
)

func init() {
	gossa.RegisterPackage("go/printer", extMap, typList)
}

var extMap = map[string]interface{}{
	"(*go/printer.Config).Fprint": (*printer.Config).Fprint,
	"go/printer.Fprint":           printer.Fprint,
}

var typList = []interface{}{
	(*printer.CommentedNode)(nil),
	(*printer.Config)(nil),
	(*printer.Mode)(nil),
}
