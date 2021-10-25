// export by github.com/goplus/gossa/cmd/qexp

//go:build go1.16
// +build go1.16

package http

import (
	"net/http"

	"github.com/goplus/gossa"
)

func init() {
	gossa.RegisterPackage("net/http", extMap_go116, typList_go116)
}

var extMap_go116 = map[string]interface{}{
	"net/http.FS": http.FS,
}

var typList_go116 = []interface{}{}
