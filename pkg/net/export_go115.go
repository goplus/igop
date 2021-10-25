// export by github.com/goplus/gossa/cmd/qexp

//go:build go1.15
// +build go1.15

package net

import (
	"net"

	"github.com/goplus/gossa"
)

func init() {
	gossa.RegisterPackage("net", extMap_go115, typList_go115)
}

var extMap_go115 = map[string]interface{}{
	"(*net.Resolver).LookupIP": (*net.Resolver).LookupIP,
}

var typList_go115 = []interface{}{}
