// export by github.com/goplus/gossa/cmd/qexp

//go:build go1.16
// +build go1.16

package log

import (
	"log"

	"github.com/goplus/gossa"
)

func init() {
	gossa.RegisterPackage("log", extMap_go116, typList_go116)
}

var extMap_go116 = map[string]interface{}{
	"log.Default": log.Default,
}

var typList_go116 = []interface{}{}
