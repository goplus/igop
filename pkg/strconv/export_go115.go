// export by github.com/goplus/interp/cmd/qexp

// +build go1.15

package strconv

import (
	"strconv"

	"github.com/goplus/interp"
)

func init() {
	interp.RegisterPackage("strconv", extMap_go115, typList_go115)
}

var extMap_go115 = map[string]interface{}{
	"strconv.FormatComplex": strconv.FormatComplex,
	"strconv.ParseComplex":  strconv.ParseComplex,
}

var typList_go115 = []interface{}{}
