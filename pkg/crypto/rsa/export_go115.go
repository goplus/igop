// export by github.com/goplus/interp/cmd/qexp

//go:build go1.15
// +build go1.15

package rsa

import (
	"crypto/rsa"

	"github.com/goplus/interp"
)

func init() {
	interp.RegisterPackage("crypto/rsa", extMap_go115, typList_go115)
}

var extMap_go115 = map[string]interface{}{
	"(*crypto/rsa.PrivateKey).Equal": (*rsa.PrivateKey).Equal,
	"(*crypto/rsa.PublicKey).Equal":  (*rsa.PublicKey).Equal,
}

var typList_go115 = []interface{}{}
