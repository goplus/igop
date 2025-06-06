// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.21 && !go1.22
// +build go1.21,!go1.22

package x509

import (
	q "crypto/x509"

	"go/constant"
	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage(&ixgo.Package{
		Name: "x509",
		Path: "crypto/x509",
		Deps: map[string]string{
			"bytes":                                 "bytes",
			"crypto":                                "crypto",
			"crypto/aes":                            "aes",
			"crypto/cipher":                         "cipher",
			"crypto/des":                            "des",
			"crypto/dsa":                            "dsa",
			"crypto/ecdh":                           "ecdh",
			"crypto/ecdsa":                          "ecdsa",
			"crypto/ed25519":                        "ed25519",
			"crypto/elliptic":                       "elliptic",
			"crypto/md5":                            "md5",
			"crypto/rsa":                            "rsa",
			"crypto/sha1":                           "sha1",
			"crypto/sha256":                         "sha256",
			"crypto/sha512":                         "sha512",
			"crypto/x509/internal/macos":            "macOS",
			"crypto/x509/pkix":                      "pkix",
			"encoding/asn1":                         "asn1",
			"encoding/hex":                          "hex",
			"encoding/pem":                          "pem",
			"errors":                                "errors",
			"fmt":                                   "fmt",
			"internal/godebug":                      "godebug",
			"io":                                    "io",
			"math/big":                              "big",
			"net":                                   "net",
			"net/url":                               "url",
			"reflect":                               "reflect",
			"runtime":                               "runtime",
			"strconv":                               "strconv",
			"strings":                               "strings",
			"sync":                                  "sync",
			"time":                                  "time",
			"unicode":                               "unicode",
			"unicode/utf16":                         "utf16",
			"unicode/utf8":                          "utf8",
			"vendor/golang.org/x/crypto/cryptobyte": "cryptobyte",
			"vendor/golang.org/x/crypto/cryptobyte/asn1": "asn1",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{
			"CertPool":                   reflect.TypeOf((*q.CertPool)(nil)).Elem(),
			"Certificate":                reflect.TypeOf((*q.Certificate)(nil)).Elem(),
			"CertificateInvalidError":    reflect.TypeOf((*q.CertificateInvalidError)(nil)).Elem(),
			"CertificateRequest":         reflect.TypeOf((*q.CertificateRequest)(nil)).Elem(),
			"ConstraintViolationError":   reflect.TypeOf((*q.ConstraintViolationError)(nil)).Elem(),
			"ExtKeyUsage":                reflect.TypeOf((*q.ExtKeyUsage)(nil)).Elem(),
			"HostnameError":              reflect.TypeOf((*q.HostnameError)(nil)).Elem(),
			"InsecureAlgorithmError":     reflect.TypeOf((*q.InsecureAlgorithmError)(nil)).Elem(),
			"InvalidReason":              reflect.TypeOf((*q.InvalidReason)(nil)).Elem(),
			"KeyUsage":                   reflect.TypeOf((*q.KeyUsage)(nil)).Elem(),
			"PEMCipher":                  reflect.TypeOf((*q.PEMCipher)(nil)).Elem(),
			"PublicKeyAlgorithm":         reflect.TypeOf((*q.PublicKeyAlgorithm)(nil)).Elem(),
			"RevocationList":             reflect.TypeOf((*q.RevocationList)(nil)).Elem(),
			"RevocationListEntry":        reflect.TypeOf((*q.RevocationListEntry)(nil)).Elem(),
			"SignatureAlgorithm":         reflect.TypeOf((*q.SignatureAlgorithm)(nil)).Elem(),
			"SystemRootsError":           reflect.TypeOf((*q.SystemRootsError)(nil)).Elem(),
			"UnhandledCriticalExtension": reflect.TypeOf((*q.UnhandledCriticalExtension)(nil)).Elem(),
			"UnknownAuthorityError":      reflect.TypeOf((*q.UnknownAuthorityError)(nil)).Elem(),
			"VerifyOptions":              reflect.TypeOf((*q.VerifyOptions)(nil)).Elem(),
		},
		AliasTypes: map[string]reflect.Type{},
		Vars: map[string]reflect.Value{
			"ErrUnsupportedAlgorithm": reflect.ValueOf(&q.ErrUnsupportedAlgorithm),
			"IncorrectPasswordError":  reflect.ValueOf(&q.IncorrectPasswordError),
		},
		Funcs: map[string]reflect.Value{
			"CreateCertificate":        reflect.ValueOf(q.CreateCertificate),
			"CreateCertificateRequest": reflect.ValueOf(q.CreateCertificateRequest),
			"CreateRevocationList":     reflect.ValueOf(q.CreateRevocationList),
			"DecryptPEMBlock":          reflect.ValueOf(q.DecryptPEMBlock),
			"EncryptPEMBlock":          reflect.ValueOf(q.EncryptPEMBlock),
			"IsEncryptedPEMBlock":      reflect.ValueOf(q.IsEncryptedPEMBlock),
			"MarshalECPrivateKey":      reflect.ValueOf(q.MarshalECPrivateKey),
			"MarshalPKCS1PrivateKey":   reflect.ValueOf(q.MarshalPKCS1PrivateKey),
			"MarshalPKCS1PublicKey":    reflect.ValueOf(q.MarshalPKCS1PublicKey),
			"MarshalPKCS8PrivateKey":   reflect.ValueOf(q.MarshalPKCS8PrivateKey),
			"MarshalPKIXPublicKey":     reflect.ValueOf(q.MarshalPKIXPublicKey),
			"NewCertPool":              reflect.ValueOf(q.NewCertPool),
			"ParseCRL":                 reflect.ValueOf(q.ParseCRL),
			"ParseCertificate":         reflect.ValueOf(q.ParseCertificate),
			"ParseCertificateRequest":  reflect.ValueOf(q.ParseCertificateRequest),
			"ParseCertificates":        reflect.ValueOf(q.ParseCertificates),
			"ParseDERCRL":              reflect.ValueOf(q.ParseDERCRL),
			"ParseECPrivateKey":        reflect.ValueOf(q.ParseECPrivateKey),
			"ParsePKCS1PrivateKey":     reflect.ValueOf(q.ParsePKCS1PrivateKey),
			"ParsePKCS1PublicKey":      reflect.ValueOf(q.ParsePKCS1PublicKey),
			"ParsePKCS8PrivateKey":     reflect.ValueOf(q.ParsePKCS8PrivateKey),
			"ParsePKIXPublicKey":       reflect.ValueOf(q.ParsePKIXPublicKey),
			"ParseRevocationList":      reflect.ValueOf(q.ParseRevocationList),
			"SetFallbackRoots":         reflect.ValueOf(q.SetFallbackRoots),
			"SystemCertPool":           reflect.ValueOf(q.SystemCertPool),
		},
		TypedConsts: map[string]ixgo.TypedConst{
			"CANotAuthorizedForExtKeyUsage": {reflect.TypeOf(q.CANotAuthorizedForExtKeyUsage), constant.MakeInt64(int64(q.CANotAuthorizedForExtKeyUsage))},
			"CANotAuthorizedForThisName":    {reflect.TypeOf(q.CANotAuthorizedForThisName), constant.MakeInt64(int64(q.CANotAuthorizedForThisName))},
			"DSA":                           {reflect.TypeOf(q.DSA), constant.MakeInt64(int64(q.DSA))},
			"DSAWithSHA1":                   {reflect.TypeOf(q.DSAWithSHA1), constant.MakeInt64(int64(q.DSAWithSHA1))},
			"DSAWithSHA256":                 {reflect.TypeOf(q.DSAWithSHA256), constant.MakeInt64(int64(q.DSAWithSHA256))},
			"ECDSA":                         {reflect.TypeOf(q.ECDSA), constant.MakeInt64(int64(q.ECDSA))},
			"ECDSAWithSHA1":                 {reflect.TypeOf(q.ECDSAWithSHA1), constant.MakeInt64(int64(q.ECDSAWithSHA1))},
			"ECDSAWithSHA256":               {reflect.TypeOf(q.ECDSAWithSHA256), constant.MakeInt64(int64(q.ECDSAWithSHA256))},
			"ECDSAWithSHA384":               {reflect.TypeOf(q.ECDSAWithSHA384), constant.MakeInt64(int64(q.ECDSAWithSHA384))},
			"ECDSAWithSHA512":               {reflect.TypeOf(q.ECDSAWithSHA512), constant.MakeInt64(int64(q.ECDSAWithSHA512))},
			"Ed25519":                       {reflect.TypeOf(q.Ed25519), constant.MakeInt64(int64(q.Ed25519))},
			"Expired":                       {reflect.TypeOf(q.Expired), constant.MakeInt64(int64(q.Expired))},
			"ExtKeyUsageAny":                {reflect.TypeOf(q.ExtKeyUsageAny), constant.MakeInt64(int64(q.ExtKeyUsageAny))},
			"ExtKeyUsageClientAuth":         {reflect.TypeOf(q.ExtKeyUsageClientAuth), constant.MakeInt64(int64(q.ExtKeyUsageClientAuth))},
			"ExtKeyUsageCodeSigning":        {reflect.TypeOf(q.ExtKeyUsageCodeSigning), constant.MakeInt64(int64(q.ExtKeyUsageCodeSigning))},
			"ExtKeyUsageEmailProtection":    {reflect.TypeOf(q.ExtKeyUsageEmailProtection), constant.MakeInt64(int64(q.ExtKeyUsageEmailProtection))},
			"ExtKeyUsageIPSECEndSystem":     {reflect.TypeOf(q.ExtKeyUsageIPSECEndSystem), constant.MakeInt64(int64(q.ExtKeyUsageIPSECEndSystem))},
			"ExtKeyUsageIPSECTunnel":        {reflect.TypeOf(q.ExtKeyUsageIPSECTunnel), constant.MakeInt64(int64(q.ExtKeyUsageIPSECTunnel))},
			"ExtKeyUsageIPSECUser":          {reflect.TypeOf(q.ExtKeyUsageIPSECUser), constant.MakeInt64(int64(q.ExtKeyUsageIPSECUser))},
			"ExtKeyUsageMicrosoftCommercialCodeSigning": {reflect.TypeOf(q.ExtKeyUsageMicrosoftCommercialCodeSigning), constant.MakeInt64(int64(q.ExtKeyUsageMicrosoftCommercialCodeSigning))},
			"ExtKeyUsageMicrosoftKernelCodeSigning":     {reflect.TypeOf(q.ExtKeyUsageMicrosoftKernelCodeSigning), constant.MakeInt64(int64(q.ExtKeyUsageMicrosoftKernelCodeSigning))},
			"ExtKeyUsageMicrosoftServerGatedCrypto":     {reflect.TypeOf(q.ExtKeyUsageMicrosoftServerGatedCrypto), constant.MakeInt64(int64(q.ExtKeyUsageMicrosoftServerGatedCrypto))},
			"ExtKeyUsageNetscapeServerGatedCrypto":      {reflect.TypeOf(q.ExtKeyUsageNetscapeServerGatedCrypto), constant.MakeInt64(int64(q.ExtKeyUsageNetscapeServerGatedCrypto))},
			"ExtKeyUsageOCSPSigning":                    {reflect.TypeOf(q.ExtKeyUsageOCSPSigning), constant.MakeInt64(int64(q.ExtKeyUsageOCSPSigning))},
			"ExtKeyUsageServerAuth":                     {reflect.TypeOf(q.ExtKeyUsageServerAuth), constant.MakeInt64(int64(q.ExtKeyUsageServerAuth))},
			"ExtKeyUsageTimeStamping":                   {reflect.TypeOf(q.ExtKeyUsageTimeStamping), constant.MakeInt64(int64(q.ExtKeyUsageTimeStamping))},
			"IncompatibleUsage":                         {reflect.TypeOf(q.IncompatibleUsage), constant.MakeInt64(int64(q.IncompatibleUsage))},
			"KeyUsageCRLSign":                           {reflect.TypeOf(q.KeyUsageCRLSign), constant.MakeInt64(int64(q.KeyUsageCRLSign))},
			"KeyUsageCertSign":                          {reflect.TypeOf(q.KeyUsageCertSign), constant.MakeInt64(int64(q.KeyUsageCertSign))},
			"KeyUsageContentCommitment":                 {reflect.TypeOf(q.KeyUsageContentCommitment), constant.MakeInt64(int64(q.KeyUsageContentCommitment))},
			"KeyUsageDataEncipherment":                  {reflect.TypeOf(q.KeyUsageDataEncipherment), constant.MakeInt64(int64(q.KeyUsageDataEncipherment))},
			"KeyUsageDecipherOnly":                      {reflect.TypeOf(q.KeyUsageDecipherOnly), constant.MakeInt64(int64(q.KeyUsageDecipherOnly))},
			"KeyUsageDigitalSignature":                  {reflect.TypeOf(q.KeyUsageDigitalSignature), constant.MakeInt64(int64(q.KeyUsageDigitalSignature))},
			"KeyUsageEncipherOnly":                      {reflect.TypeOf(q.KeyUsageEncipherOnly), constant.MakeInt64(int64(q.KeyUsageEncipherOnly))},
			"KeyUsageKeyAgreement":                      {reflect.TypeOf(q.KeyUsageKeyAgreement), constant.MakeInt64(int64(q.KeyUsageKeyAgreement))},
			"KeyUsageKeyEncipherment":                   {reflect.TypeOf(q.KeyUsageKeyEncipherment), constant.MakeInt64(int64(q.KeyUsageKeyEncipherment))},
			"MD2WithRSA":                                {reflect.TypeOf(q.MD2WithRSA), constant.MakeInt64(int64(q.MD2WithRSA))},
			"MD5WithRSA":                                {reflect.TypeOf(q.MD5WithRSA), constant.MakeInt64(int64(q.MD5WithRSA))},
			"NameConstraintsWithoutSANs":                {reflect.TypeOf(q.NameConstraintsWithoutSANs), constant.MakeInt64(int64(q.NameConstraintsWithoutSANs))},
			"NameMismatch":                              {reflect.TypeOf(q.NameMismatch), constant.MakeInt64(int64(q.NameMismatch))},
			"NotAuthorizedToSign":                       {reflect.TypeOf(q.NotAuthorizedToSign), constant.MakeInt64(int64(q.NotAuthorizedToSign))},
			"PEMCipher3DES":                             {reflect.TypeOf(q.PEMCipher3DES), constant.MakeInt64(int64(q.PEMCipher3DES))},
			"PEMCipherAES128":                           {reflect.TypeOf(q.PEMCipherAES128), constant.MakeInt64(int64(q.PEMCipherAES128))},
			"PEMCipherAES192":                           {reflect.TypeOf(q.PEMCipherAES192), constant.MakeInt64(int64(q.PEMCipherAES192))},
			"PEMCipherAES256":                           {reflect.TypeOf(q.PEMCipherAES256), constant.MakeInt64(int64(q.PEMCipherAES256))},
			"PEMCipherDES":                              {reflect.TypeOf(q.PEMCipherDES), constant.MakeInt64(int64(q.PEMCipherDES))},
			"PureEd25519":                               {reflect.TypeOf(q.PureEd25519), constant.MakeInt64(int64(q.PureEd25519))},
			"RSA":                                       {reflect.TypeOf(q.RSA), constant.MakeInt64(int64(q.RSA))},
			"SHA1WithRSA":                               {reflect.TypeOf(q.SHA1WithRSA), constant.MakeInt64(int64(q.SHA1WithRSA))},
			"SHA256WithRSA":                             {reflect.TypeOf(q.SHA256WithRSA), constant.MakeInt64(int64(q.SHA256WithRSA))},
			"SHA256WithRSAPSS":                          {reflect.TypeOf(q.SHA256WithRSAPSS), constant.MakeInt64(int64(q.SHA256WithRSAPSS))},
			"SHA384WithRSA":                             {reflect.TypeOf(q.SHA384WithRSA), constant.MakeInt64(int64(q.SHA384WithRSA))},
			"SHA384WithRSAPSS":                          {reflect.TypeOf(q.SHA384WithRSAPSS), constant.MakeInt64(int64(q.SHA384WithRSAPSS))},
			"SHA512WithRSA":                             {reflect.TypeOf(q.SHA512WithRSA), constant.MakeInt64(int64(q.SHA512WithRSA))},
			"SHA512WithRSAPSS":                          {reflect.TypeOf(q.SHA512WithRSAPSS), constant.MakeInt64(int64(q.SHA512WithRSAPSS))},
			"TooManyConstraints":                        {reflect.TypeOf(q.TooManyConstraints), constant.MakeInt64(int64(q.TooManyConstraints))},
			"TooManyIntermediates":                      {reflect.TypeOf(q.TooManyIntermediates), constant.MakeInt64(int64(q.TooManyIntermediates))},
			"UnconstrainedName":                         {reflect.TypeOf(q.UnconstrainedName), constant.MakeInt64(int64(q.UnconstrainedName))},
			"UnknownPublicKeyAlgorithm":                 {reflect.TypeOf(q.UnknownPublicKeyAlgorithm), constant.MakeInt64(int64(q.UnknownPublicKeyAlgorithm))},
			"UnknownSignatureAlgorithm":                 {reflect.TypeOf(q.UnknownSignatureAlgorithm), constant.MakeInt64(int64(q.UnknownSignatureAlgorithm))},
		},
		UntypedConsts: map[string]ixgo.UntypedConst{},
	})
}
