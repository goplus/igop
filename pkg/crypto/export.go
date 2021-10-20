// export by github.com/goplus/gossa/cmd/qexp

package crypto

import (
	"crypto"

	"github.com/goplus/gossa"
)

func init() {
	gossa.RegisterPackage("crypto", extMap, typList)
}

var extMap = map[string]interface{}{
	"(crypto.Decrypter).Decrypt":   (crypto.Decrypter).Decrypt,
	"(crypto.Decrypter).Public":    (crypto.Decrypter).Public,
	"(crypto.Hash).Available":      (crypto.Hash).Available,
	"(crypto.Hash).HashFunc":       (crypto.Hash).HashFunc,
	"(crypto.Hash).New":            (crypto.Hash).New,
	"(crypto.Hash).Size":           (crypto.Hash).Size,
	"(crypto.Signer).Public":       (crypto.Signer).Public,
	"(crypto.Signer).Sign":         (crypto.Signer).Sign,
	"(crypto.SignerOpts).HashFunc": (crypto.SignerOpts).HashFunc,
	"crypto.RegisterHash":          crypto.RegisterHash,
}

var typList = []interface{}{
	(*crypto.Decrypter)(nil),
	(*crypto.DecrypterOpts)(nil),
	(*crypto.Hash)(nil),
	(*crypto.PrivateKey)(nil),
	(*crypto.PublicKey)(nil),
	(*crypto.Signer)(nil),
	(*crypto.SignerOpts)(nil),
}
