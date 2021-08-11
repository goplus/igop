// export by github.com/goplus/interp/cmd/qexp

// +build go1.15

package bufio

import (
	"bufio"

	"github.com/goplus/interp"
)

func init() {
	interp.RegisterPackage("bufio", extMap_go115, typList_go115)
}

var extMap_go115 = map[string]interface{}{
	"bufio.ErrBadReadCount": &bufio.ErrBadReadCount,
}

var typList_go115 = []interface{}{}
