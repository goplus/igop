// export by github.com/goplus/interp/cmd/qexp

//go:build go1.16
// +build go1.16

package net

import (
	"net"

	"github.com/goplus/interp"
)

func init() {
	interp.RegisterPackage("net", extMap_go116, typList_go116)
}

var extMap_go116 = map[string]interface{}{
	"net.ErrClosed": &net.ErrClosed,
}

var typList_go116 = []interface{}{}
