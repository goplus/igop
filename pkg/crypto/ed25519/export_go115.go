// export by github.com/goplus/interp/cmd/qexp

// +build go1.15

package ed25519

import (
	"crypto/ed25519"

	"github.com/goplus/interp"
)

func init() {
	interp.RegisterPackage("crypto/ed25519", extMap_go115, typList_go115)
}

var extMap_go115 = map[string]interface{}{
	"(crypto/ed25519.PrivateKey).Equal": (ed25519.PrivateKey).Equal,
	"(crypto/ed25519.PublicKey).Equal":  (ed25519.PublicKey).Equal,
}

var typList_go115 = []interface{}{}
