// export by github.com/goplus/interp/cmd/qexp

//go:build go1.16
// +build go1.16

package log

import (
	"log"

	"github.com/goplus/interp"
)

func init() {
	interp.RegisterPackage("log", extMap_go116, typList_go116)
}

var extMap_go116 = map[string]interface{}{
	"log.Default": log.Default,
}

var typList_go116 = []interface{}{}
