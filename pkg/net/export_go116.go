// export by github.com/goplus/gossa/cmd/qexp

//go:build go1.16
// +build go1.16

package net

import (
	"net"

	"github.com/goplus/gossa"
)

func init() {
	gossa.RegisterPackage("net", extMap_go116, typList_go116)
}

var extMap_go116 = map[string]interface{}{
	"net.ErrClosed": &net.ErrClosed,
}

var typList_go116 = []interface{}{}
