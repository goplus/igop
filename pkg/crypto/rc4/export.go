// export by github.com/goplus/interp/cmd/qexp

package rc4

import (
	"crypto/rc4"

	"github.com/goplus/interp"
)

func init() {
	interp.RegisterPackage("crypto/rc4", extMap, typList)
}

var extMap = map[string]interface{}{
	"(*crypto/rc4.Cipher).Reset":        (*rc4.Cipher).Reset,
	"(*crypto/rc4.Cipher).XORKeyStream": (*rc4.Cipher).XORKeyStream,
	"(crypto/rc4.KeySizeError).Error":   (rc4.KeySizeError).Error,
	"crypto/rc4.NewCipher":              rc4.NewCipher,
}

var typList = []interface{}{
	(*rc4.Cipher)(nil),
	(*rc4.KeySizeError)(nil),
}
