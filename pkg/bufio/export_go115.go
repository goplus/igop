// export by github.com/goplus/gossa/cmd/qexp

//go:build go1.15
// +build go1.15

package bufio

import (
	"bufio"

	"github.com/goplus/gossa"
)

func init() {
	gossa.RegisterPackage("bufio", extMap_go115, typList_go115)
}

var extMap_go115 = map[string]interface{}{
	"bufio.ErrBadReadCount": &bufio.ErrBadReadCount,
}

var typList_go115 = []interface{}{}
