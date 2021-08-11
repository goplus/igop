// export by github.com/goplus/interp/cmd/qexp

// +build go1.16

package iotest

import (
	"testing/iotest"

	"github.com/goplus/interp"
)

func init() {
	interp.RegisterPackage("testing/iotest", extMap_go116, typList_go116)
}

var extMap_go116 = map[string]interface{}{
	"testing/iotest.ErrReader":  iotest.ErrReader,
	"testing/iotest.TestReader": iotest.TestReader,
}

var typList_go116 = []interface{}{}
