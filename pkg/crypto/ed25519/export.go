// export by github.com/goplus/interp/cmd/qexp

package ed25519

import (
	"crypto/ed25519"

	"github.com/goplus/interp"
)

func init() {
	interp.RegisterPackage("crypto/ed25519", extMap, typList)
}

var extMap = map[string]interface{}{
	"(crypto/ed25519.PrivateKey).Public": (ed25519.PrivateKey).Public,
	"(crypto/ed25519.PrivateKey).Seed":   (ed25519.PrivateKey).Seed,
	"(crypto/ed25519.PrivateKey).Sign":   (ed25519.PrivateKey).Sign,
	"crypto/ed25519.GenerateKey":         ed25519.GenerateKey,
	"crypto/ed25519.NewKeyFromSeed":      ed25519.NewKeyFromSeed,
	"crypto/ed25519.Sign":                ed25519.Sign,
	"crypto/ed25519.Verify":              ed25519.Verify,
}

var typList = []interface{}{
	(*ed25519.PrivateKey)(nil),
	(*ed25519.PublicKey)(nil),
}
