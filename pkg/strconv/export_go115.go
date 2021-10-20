// export by github.com/goplus/gossa/cmd/qexp

//go:build go1.15
// +build go1.15

package strconv

import (
	"strconv"

	"github.com/goplus/gossa"
)

func init() {
	gossa.RegisterPackage("strconv", extMap_go115, typList_go115)
}

var extMap_go115 = map[string]interface{}{
	"strconv.FormatComplex": strconv.FormatComplex,
	"strconv.ParseComplex":  strconv.ParseComplex,
}

var typList_go115 = []interface{}{}
