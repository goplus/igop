// export by github.com/goplus/gossa/cmd/qexp

//go:build go1.15
// +build go1.15

package crypto

import (
	"crypto"

	"github.com/goplus/gossa"
)

func init() {
	gossa.RegisterPackage("crypto", extMap_go115, typList_go115)
}

var extMap_go115 = map[string]interface{}{
	"(crypto.Hash).String": (crypto.Hash).String,
}

var typList_go115 = []interface{}{}
