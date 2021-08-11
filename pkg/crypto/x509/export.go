// export by github.com/goplus/interp/cmd/qexp

package x509

import (
	"crypto/x509"

	"github.com/goplus/interp"
)

func init() {
	interp.RegisterPackage("crypto/x509", extMap, typList)
}

var extMap = map[string]interface{}{
	"(*crypto/x509.CertPool).AddCert":                  (*x509.CertPool).AddCert,
	"(*crypto/x509.CertPool).AppendCertsFromPEM":       (*x509.CertPool).AppendCertsFromPEM,
	"(*crypto/x509.CertPool).Subjects":                 (*x509.CertPool).Subjects,
	"(*crypto/x509.Certificate).CheckCRLSignature":     (*x509.Certificate).CheckCRLSignature,
	"(*crypto/x509.Certificate).CheckSignature":        (*x509.Certificate).CheckSignature,
	"(*crypto/x509.Certificate).CheckSignatureFrom":    (*x509.Certificate).CheckSignatureFrom,
	"(*crypto/x509.Certificate).CreateCRL":             (*x509.Certificate).CreateCRL,
	"(*crypto/x509.Certificate).Equal":                 (*x509.Certificate).Equal,
	"(*crypto/x509.Certificate).Verify":                (*x509.Certificate).Verify,
	"(*crypto/x509.Certificate).VerifyHostname":        (*x509.Certificate).VerifyHostname,
	"(*crypto/x509.CertificateRequest).CheckSignature": (*x509.CertificateRequest).CheckSignature,
	"(crypto/x509.CertificateInvalidError).Error":      (x509.CertificateInvalidError).Error,
	"(crypto/x509.ConstraintViolationError).Error":     (x509.ConstraintViolationError).Error,
	"(crypto/x509.HostnameError).Error":                (x509.HostnameError).Error,
	"(crypto/x509.InsecureAlgorithmError).Error":       (x509.InsecureAlgorithmError).Error,
	"(crypto/x509.PublicKeyAlgorithm).String":          (x509.PublicKeyAlgorithm).String,
	"(crypto/x509.SignatureAlgorithm).String":          (x509.SignatureAlgorithm).String,
	"(crypto/x509.SystemRootsError).Error":             (x509.SystemRootsError).Error,
	"(crypto/x509.UnhandledCriticalExtension).Error":   (x509.UnhandledCriticalExtension).Error,
	"(crypto/x509.UnknownAuthorityError).Error":        (x509.UnknownAuthorityError).Error,
	"crypto/x509.CreateCertificate":                    x509.CreateCertificate,
	"crypto/x509.CreateCertificateRequest":             x509.CreateCertificateRequest,
	"crypto/x509.DecryptPEMBlock":                      x509.DecryptPEMBlock,
	"crypto/x509.EncryptPEMBlock":                      x509.EncryptPEMBlock,
	"crypto/x509.ErrUnsupportedAlgorithm":              &x509.ErrUnsupportedAlgorithm,
	"crypto/x509.IncorrectPasswordError":               &x509.IncorrectPasswordError,
	"crypto/x509.IsEncryptedPEMBlock":                  x509.IsEncryptedPEMBlock,
	"crypto/x509.MarshalECPrivateKey":                  x509.MarshalECPrivateKey,
	"crypto/x509.MarshalPKCS1PrivateKey":               x509.MarshalPKCS1PrivateKey,
	"crypto/x509.MarshalPKCS1PublicKey":                x509.MarshalPKCS1PublicKey,
	"crypto/x509.MarshalPKCS8PrivateKey":               x509.MarshalPKCS8PrivateKey,
	"crypto/x509.MarshalPKIXPublicKey":                 x509.MarshalPKIXPublicKey,
	"crypto/x509.NewCertPool":                          x509.NewCertPool,
	"crypto/x509.ParseCRL":                             x509.ParseCRL,
	"crypto/x509.ParseCertificate":                     x509.ParseCertificate,
	"crypto/x509.ParseCertificateRequest":              x509.ParseCertificateRequest,
	"crypto/x509.ParseCertificates":                    x509.ParseCertificates,
	"crypto/x509.ParseDERCRL":                          x509.ParseDERCRL,
	"crypto/x509.ParseECPrivateKey":                    x509.ParseECPrivateKey,
	"crypto/x509.ParsePKCS1PrivateKey":                 x509.ParsePKCS1PrivateKey,
	"crypto/x509.ParsePKCS1PublicKey":                  x509.ParsePKCS1PublicKey,
	"crypto/x509.ParsePKCS8PrivateKey":                 x509.ParsePKCS8PrivateKey,
	"crypto/x509.ParsePKIXPublicKey":                   x509.ParsePKIXPublicKey,
	"crypto/x509.SystemCertPool":                       x509.SystemCertPool,
}

var typList = []interface{}{
	(*x509.CertPool)(nil),
	(*x509.Certificate)(nil),
	(*x509.CertificateInvalidError)(nil),
	(*x509.CertificateRequest)(nil),
	(*x509.ConstraintViolationError)(nil),
	(*x509.ExtKeyUsage)(nil),
	(*x509.HostnameError)(nil),
	(*x509.InsecureAlgorithmError)(nil),
	(*x509.InvalidReason)(nil),
	(*x509.KeyUsage)(nil),
	(*x509.PEMCipher)(nil),
	(*x509.PublicKeyAlgorithm)(nil),
	(*x509.SignatureAlgorithm)(nil),
	(*x509.SystemRootsError)(nil),
	(*x509.UnhandledCriticalExtension)(nil),
	(*x509.UnknownAuthorityError)(nil),
	(*x509.VerifyOptions)(nil),
}
