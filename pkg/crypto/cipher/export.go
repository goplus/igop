// export by github.com/goplus/gossa/cmd/qexp

package cipher

import (
	"crypto/cipher"

	"github.com/goplus/gossa"
)

func init() {
	gossa.RegisterPackage("crypto/cipher", extMap, typList)
}

var extMap = map[string]interface{}{
	"(crypto/cipher.AEAD).NonceSize":        (cipher.AEAD).NonceSize,
	"(crypto/cipher.AEAD).Open":             (cipher.AEAD).Open,
	"(crypto/cipher.AEAD).Overhead":         (cipher.AEAD).Overhead,
	"(crypto/cipher.AEAD).Seal":             (cipher.AEAD).Seal,
	"(crypto/cipher.Block).BlockSize":       (cipher.Block).BlockSize,
	"(crypto/cipher.Block).Decrypt":         (cipher.Block).Decrypt,
	"(crypto/cipher.Block).Encrypt":         (cipher.Block).Encrypt,
	"(crypto/cipher.BlockMode).BlockSize":   (cipher.BlockMode).BlockSize,
	"(crypto/cipher.BlockMode).CryptBlocks": (cipher.BlockMode).CryptBlocks,
	"(crypto/cipher.Stream).XORKeyStream":   (cipher.Stream).XORKeyStream,
	"(crypto/cipher.StreamReader).Read":     (cipher.StreamReader).Read,
	"(crypto/cipher.StreamWriter).Close":    (cipher.StreamWriter).Close,
	"(crypto/cipher.StreamWriter).Write":    (cipher.StreamWriter).Write,
	"crypto/cipher.NewCBCDecrypter":         cipher.NewCBCDecrypter,
	"crypto/cipher.NewCBCEncrypter":         cipher.NewCBCEncrypter,
	"crypto/cipher.NewCFBDecrypter":         cipher.NewCFBDecrypter,
	"crypto/cipher.NewCFBEncrypter":         cipher.NewCFBEncrypter,
	"crypto/cipher.NewCTR":                  cipher.NewCTR,
	"crypto/cipher.NewGCM":                  cipher.NewGCM,
	"crypto/cipher.NewGCMWithNonceSize":     cipher.NewGCMWithNonceSize,
	"crypto/cipher.NewGCMWithTagSize":       cipher.NewGCMWithTagSize,
	"crypto/cipher.NewOFB":                  cipher.NewOFB,
}

var typList = []interface{}{
	(*cipher.AEAD)(nil),
	(*cipher.Block)(nil),
	(*cipher.BlockMode)(nil),
	(*cipher.Stream)(nil),
	(*cipher.StreamReader)(nil),
	(*cipher.StreamWriter)(nil),
}
