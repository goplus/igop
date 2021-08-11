// export by github.com/goplus/interp/cmd/qexp

package lzw

import (
	"compress/lzw"

	"github.com/goplus/interp"
)

func init() {
	interp.RegisterPackage("compress/lzw", extMap, typList)
}

var extMap = map[string]interface{}{
	"compress/lzw.NewReader": lzw.NewReader,
	"compress/lzw.NewWriter": lzw.NewWriter,
}

var typList = []interface{}{
	(*lzw.Order)(nil),
}
