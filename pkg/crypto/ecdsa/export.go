// export by github.com/goplus/gossa/cmd/qexp

package ecdsa

import (
	"crypto/ecdsa"

	"github.com/goplus/gossa"
)

func init() {
	gossa.RegisterPackage("crypto/ecdsa", extMap, typList)
}

var extMap = map[string]interface{}{
	"(*crypto/ecdsa.PrivateKey).Public":        (*ecdsa.PrivateKey).Public,
	"(*crypto/ecdsa.PrivateKey).Sign":          (*ecdsa.PrivateKey).Sign,
	"(crypto/ecdsa.PrivateKey).Add":            (ecdsa.PrivateKey).Add,
	"(crypto/ecdsa.PrivateKey).Double":         (ecdsa.PrivateKey).Double,
	"(crypto/ecdsa.PrivateKey).IsOnCurve":      (ecdsa.PrivateKey).IsOnCurve,
	"(crypto/ecdsa.PrivateKey).Params":         (ecdsa.PrivateKey).Params,
	"(crypto/ecdsa.PrivateKey).ScalarBaseMult": (ecdsa.PrivateKey).ScalarBaseMult,
	"(crypto/ecdsa.PrivateKey).ScalarMult":     (ecdsa.PrivateKey).ScalarMult,
	"(crypto/ecdsa.PublicKey).Add":             (ecdsa.PublicKey).Add,
	"(crypto/ecdsa.PublicKey).Double":          (ecdsa.PublicKey).Double,
	"(crypto/ecdsa.PublicKey).IsOnCurve":       (ecdsa.PublicKey).IsOnCurve,
	"(crypto/ecdsa.PublicKey).Params":          (ecdsa.PublicKey).Params,
	"(crypto/ecdsa.PublicKey).ScalarBaseMult":  (ecdsa.PublicKey).ScalarBaseMult,
	"(crypto/ecdsa.PublicKey).ScalarMult":      (ecdsa.PublicKey).ScalarMult,
	"crypto/ecdsa.GenerateKey":                 ecdsa.GenerateKey,
	"crypto/ecdsa.Sign":                        ecdsa.Sign,
	"crypto/ecdsa.Verify":                      ecdsa.Verify,
}

var typList = []interface{}{
	(*ecdsa.PrivateKey)(nil),
	(*ecdsa.PublicKey)(nil),
}
