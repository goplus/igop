// export by github.com/goplus/interp/cmd/qexp

// +build go1.15

package os

import (
	"os"

	"github.com/goplus/interp"
)

func init() {
	interp.RegisterPackage("os", extMap_go115, typList_go115)
}

var extMap_go115 = map[string]interface{}{
	"(*os.File).ReadFrom":    (*os.File).ReadFrom,
	"os.ErrDeadlineExceeded": &os.ErrDeadlineExceeded,
}

var typList_go115 = []interface{}{}
