// export by github.com/goplus/interp/cmd/qexp

package aes

import (
	"crypto/aes"

	"github.com/goplus/interp"
)

func init() {
	interp.RegisterPackage("crypto/aes", extMap, typList)
}

var extMap = map[string]interface{}{
	"(crypto/aes.KeySizeError).Error": (aes.KeySizeError).Error,
	"crypto/aes.NewCipher":            aes.NewCipher,
}

var typList = []interface{}{
	(*aes.KeySizeError)(nil),
}
