// export by github.com/goplus/interp/cmd/qexp

package tabwriter

import (
	"text/tabwriter"

	"github.com/goplus/interp"
)

func init() {
	interp.RegisterPackage("text/tabwriter", extMap, typList)
}

var extMap = map[string]interface{}{
	"(*text/tabwriter.Writer).Flush": (*tabwriter.Writer).Flush,
	"(*text/tabwriter.Writer).Init":  (*tabwriter.Writer).Init,
	"(*text/tabwriter.Writer).Write": (*tabwriter.Writer).Write,
	"text/tabwriter.NewWriter":       tabwriter.NewWriter,
}

var typList = []interface{}{
	(*tabwriter.Writer)(nil),
}
