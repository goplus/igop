// export by github.com/goplus/gossa/cmd/qexp

package plan9obj

import (
	"debug/plan9obj"

	"github.com/goplus/gossa"
)

func init() {
	gossa.RegisterPackage("debug/plan9obj", extMap, typList)
}

var extMap = map[string]interface{}{
	"(*debug/plan9obj.File).Close":    (*plan9obj.File).Close,
	"(*debug/plan9obj.File).Section":  (*plan9obj.File).Section,
	"(*debug/plan9obj.File).Symbols":  (*plan9obj.File).Symbols,
	"(*debug/plan9obj.Section).Data":  (*plan9obj.Section).Data,
	"(*debug/plan9obj.Section).Open":  (*plan9obj.Section).Open,
	"(debug/plan9obj.Section).ReadAt": (plan9obj.Section).ReadAt,
	"debug/plan9obj.NewFile":          plan9obj.NewFile,
	"debug/plan9obj.Open":             plan9obj.Open,
}

var typList = []interface{}{
	(*plan9obj.File)(nil),
	(*plan9obj.FileHeader)(nil),
	(*plan9obj.Section)(nil),
	(*plan9obj.SectionHeader)(nil),
	(*plan9obj.Sym)(nil),
}
