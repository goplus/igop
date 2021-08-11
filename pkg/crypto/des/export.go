// export by github.com/goplus/interp/cmd/qexp

package des

import (
	"crypto/des"

	"github.com/goplus/interp"
)

func init() {
	interp.RegisterPackage("crypto/des", extMap, typList)
}

var extMap = map[string]interface{}{
	"(crypto/des.KeySizeError).Error": (des.KeySizeError).Error,
	"crypto/des.NewCipher":            des.NewCipher,
	"crypto/des.NewTripleDESCipher":   des.NewTripleDESCipher,
}

var typList = []interface{}{
	(*des.KeySizeError)(nil),
}
