// export by github.com/goplus/gossa/cmd/qexp

//go:build go1.15
// +build go1.15

package os

import (
	"os"

	"github.com/goplus/gossa"
)

func init() {
	gossa.RegisterPackage("os", extMap_go115, typList_go115)
}

var extMap_go115 = map[string]interface{}{
	"(*os.File).ReadFrom":    (*os.File).ReadFrom,
	"os.ErrDeadlineExceeded": &os.ErrDeadlineExceeded,
}

var typList_go115 = []interface{}{}
