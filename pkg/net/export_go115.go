// export by github.com/goplus/interp/cmd/qexp

// +build go1.15

package net

import (
	"net"

	"github.com/goplus/interp"
)

func init() {
	interp.RegisterPackage("net", extMap_go115, typList_go115)
}

var extMap_go115 = map[string]interface{}{
	"(*net.Resolver).LookupIP": (*net.Resolver).LookupIP,
}

var typList_go115 = []interface{}{}
