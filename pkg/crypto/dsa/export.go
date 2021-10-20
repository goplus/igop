// export by github.com/goplus/gossa/cmd/qexp

package dsa

import (
	"crypto/dsa"

	"github.com/goplus/gossa"
)

func init() {
	gossa.RegisterPackage("crypto/dsa", extMap, typList)
}

var extMap = map[string]interface{}{
	"crypto/dsa.ErrInvalidPublicKey": &dsa.ErrInvalidPublicKey,
	"crypto/dsa.GenerateKey":         dsa.GenerateKey,
	"crypto/dsa.GenerateParameters":  dsa.GenerateParameters,
	"crypto/dsa.Sign":                dsa.Sign,
	"crypto/dsa.Verify":              dsa.Verify,
}

var typList = []interface{}{
	(*dsa.ParameterSizes)(nil),
	(*dsa.Parameters)(nil),
	(*dsa.PrivateKey)(nil),
	(*dsa.PublicKey)(nil),
}
