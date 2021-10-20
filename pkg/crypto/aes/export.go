// export by github.com/goplus/gossa/cmd/qexp

package aes

import (
	"crypto/aes"

	"github.com/goplus/gossa"
)

func init() {
	gossa.RegisterPackage("crypto/aes", extMap, typList)
}

var extMap = map[string]interface{}{
	"(crypto/aes.KeySizeError).Error": (aes.KeySizeError).Error,
	"crypto/aes.NewCipher":            aes.NewCipher,
}

var typList = []interface{}{
	(*aes.KeySizeError)(nil),
}
