// export by github.com/goplus/interp/cmd/qexp

//go:build go1.15
// +build go1.15

package crypto

import (
	"crypto"

	"github.com/goplus/interp"
)

func init() {
	interp.RegisterPackage("crypto", extMap_go115, typList_go115)
}

var extMap_go115 = map[string]interface{}{
	"(crypto.Hash).String": (crypto.Hash).String,
}

var typList_go115 = []interface{}{}
