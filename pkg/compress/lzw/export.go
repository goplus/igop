// export by github.com/goplus/gossa/cmd/qexp

package lzw

import (
	"compress/lzw"

	"github.com/goplus/gossa"
)

func init() {
	gossa.RegisterPackage("compress/lzw", extMap, typList)
}

var extMap = map[string]interface{}{
	"compress/lzw.NewReader": lzw.NewReader,
	"compress/lzw.NewWriter": lzw.NewWriter,
}

var typList = []interface{}{
	(*lzw.Order)(nil),
}
