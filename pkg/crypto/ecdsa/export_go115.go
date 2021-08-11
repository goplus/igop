// export by github.com/goplus/interp/cmd/qexp

// +build go1.15

package ecdsa

import (
	"crypto/ecdsa"

	"github.com/goplus/interp"
)

func init() {
	interp.RegisterPackage("crypto/ecdsa", extMap_go115, typList_go115)
}

var extMap_go115 = map[string]interface{}{
	"(*crypto/ecdsa.PrivateKey).Equal": (*ecdsa.PrivateKey).Equal,
	"(*crypto/ecdsa.PublicKey).Equal":  (*ecdsa.PublicKey).Equal,
	"crypto/ecdsa.SignASN1":            ecdsa.SignASN1,
	"crypto/ecdsa.VerifyASN1":          ecdsa.VerifyASN1,
}

var typList_go115 = []interface{}{}
