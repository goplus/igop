// export by github.com/goplus/gossa/cmd/qexp

package des

import (
	"crypto/des"

	"github.com/goplus/gossa"
)

func init() {
	gossa.RegisterPackage("crypto/des", extMap, typList)
}

var extMap = map[string]interface{}{
	"(crypto/des.KeySizeError).Error": (des.KeySizeError).Error,
	"crypto/des.NewCipher":            des.NewCipher,
	"crypto/des.NewTripleDESCipher":   des.NewTripleDESCipher,
}

var typList = []interface{}{
	(*des.KeySizeError)(nil),
}
