// export by github.com/goplus/interp/cmd/qexp

package iotest

import (
	"testing/iotest"

	"github.com/goplus/interp"
)

func init() {
	interp.RegisterPackage("testing/iotest", extMap, typList)
}

var extMap = map[string]interface{}{
	"testing/iotest.DataErrReader":  iotest.DataErrReader,
	"testing/iotest.ErrTimeout":     &iotest.ErrTimeout,
	"testing/iotest.HalfReader":     iotest.HalfReader,
	"testing/iotest.NewReadLogger":  iotest.NewReadLogger,
	"testing/iotest.NewWriteLogger": iotest.NewWriteLogger,
	"testing/iotest.OneByteReader":  iotest.OneByteReader,
	"testing/iotest.TimeoutReader":  iotest.TimeoutReader,
	"testing/iotest.TruncateWriter": iotest.TruncateWriter,
}

var typList = []interface{}{}
