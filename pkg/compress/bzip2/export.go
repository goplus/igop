// export by github.com/goplus/interp/cmd/qexp

package bzip2

import (
	"compress/bzip2"

	"github.com/goplus/interp"
)

func init() {
	interp.RegisterPackage("compress/bzip2", extMap, typList)
}

var extMap = map[string]interface{}{
	"(compress/bzip2.StructuralError).Error": (bzip2.StructuralError).Error,
	"compress/bzip2.NewReader":               bzip2.NewReader,
}

var typList = []interface{}{
	(*bzip2.StructuralError)(nil),
}
