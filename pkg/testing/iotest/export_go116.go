// export by github.com/goplus/gossa/cmd/qexp

//go:build go1.16
// +build go1.16

package iotest

import (
	"testing/iotest"

	"github.com/goplus/gossa"
)

func init() {
	gossa.RegisterPackage("testing/iotest", extMap_go116, typList_go116)
}

var extMap_go116 = map[string]interface{}{
	"testing/iotest.ErrReader":  iotest.ErrReader,
	"testing/iotest.TestReader": iotest.TestReader,
}

var typList_go116 = []interface{}{}
