// export by github.com/goplus/interp/cmd/qexp

package rsa

import (
	"crypto/rsa"

	"github.com/goplus/interp"
)

func init() {
	interp.RegisterPackage("crypto/rsa", extMap, typList)
}

var extMap = map[string]interface{}{
	"(*crypto/rsa.PSSOptions).HashFunc":    (*rsa.PSSOptions).HashFunc,
	"(*crypto/rsa.PrivateKey).Decrypt":     (*rsa.PrivateKey).Decrypt,
	"(*crypto/rsa.PrivateKey).Precompute":  (*rsa.PrivateKey).Precompute,
	"(*crypto/rsa.PrivateKey).Public":      (*rsa.PrivateKey).Public,
	"(*crypto/rsa.PrivateKey).Sign":        (*rsa.PrivateKey).Sign,
	"(*crypto/rsa.PrivateKey).Size":        (*rsa.PrivateKey).Size,
	"(*crypto/rsa.PrivateKey).Validate":    (*rsa.PrivateKey).Validate,
	"(*crypto/rsa.PublicKey).Size":         (*rsa.PublicKey).Size,
	"crypto/rsa.DecryptOAEP":               rsa.DecryptOAEP,
	"crypto/rsa.DecryptPKCS1v15":           rsa.DecryptPKCS1v15,
	"crypto/rsa.DecryptPKCS1v15SessionKey": rsa.DecryptPKCS1v15SessionKey,
	"crypto/rsa.EncryptOAEP":               rsa.EncryptOAEP,
	"crypto/rsa.EncryptPKCS1v15":           rsa.EncryptPKCS1v15,
	"crypto/rsa.ErrDecryption":             &rsa.ErrDecryption,
	"crypto/rsa.ErrMessageTooLong":         &rsa.ErrMessageTooLong,
	"crypto/rsa.ErrVerification":           &rsa.ErrVerification,
	"crypto/rsa.GenerateKey":               rsa.GenerateKey,
	"crypto/rsa.GenerateMultiPrimeKey":     rsa.GenerateMultiPrimeKey,
	"crypto/rsa.SignPKCS1v15":              rsa.SignPKCS1v15,
	"crypto/rsa.SignPSS":                   rsa.SignPSS,
	"crypto/rsa.VerifyPKCS1v15":            rsa.VerifyPKCS1v15,
	"crypto/rsa.VerifyPSS":                 rsa.VerifyPSS,
}

var typList = []interface{}{
	(*rsa.CRTValue)(nil),
	(*rsa.OAEPOptions)(nil),
	(*rsa.PKCS1v15DecryptOptions)(nil),
	(*rsa.PSSOptions)(nil),
	(*rsa.PrecomputedValues)(nil),
	(*rsa.PrivateKey)(nil),
	(*rsa.PublicKey)(nil),
}
