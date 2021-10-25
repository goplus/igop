// export by github.com/goplus/gossa/cmd/qexp

//go:build go1.15
// +build go1.15

package ed25519

import (
	"crypto/ed25519"

	"github.com/goplus/gossa"
)

func init() {
	gossa.RegisterPackage("crypto/ed25519", extMap_go115, typList_go115)
}

var extMap_go115 = map[string]interface{}{
	"(crypto/ed25519.PrivateKey).Equal": (ed25519.PrivateKey).Equal,
	"(crypto/ed25519.PublicKey).Equal":  (ed25519.PublicKey).Equal,
}

var typList_go115 = []interface{}{}
