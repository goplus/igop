// export by github.com/goplus/gossa/cmd/qexp

package pkix

import (
	"crypto/x509/pkix"

	"github.com/goplus/gossa"
)

func init() {
	gossa.RegisterPackage("crypto/x509/pkix", extMap, typList)
}

var extMap = map[string]interface{}{
	"(*crypto/x509/pkix.CertificateList).HasExpired": (*pkix.CertificateList).HasExpired,
	"(*crypto/x509/pkix.Name).FillFromRDNSequence":   (*pkix.Name).FillFromRDNSequence,
	"(crypto/x509/pkix.Name).String":                 (pkix.Name).String,
	"(crypto/x509/pkix.Name).ToRDNSequence":          (pkix.Name).ToRDNSequence,
	"(crypto/x509/pkix.RDNSequence).String":          (pkix.RDNSequence).String,
}

var typList = []interface{}{
	(*pkix.AlgorithmIdentifier)(nil),
	(*pkix.AttributeTypeAndValue)(nil),
	(*pkix.AttributeTypeAndValueSET)(nil),
	(*pkix.CertificateList)(nil),
	(*pkix.Extension)(nil),
	(*pkix.Name)(nil),
	(*pkix.RDNSequence)(nil),
	(*pkix.RelativeDistinguishedNameSET)(nil),
	(*pkix.RevokedCertificate)(nil),
	(*pkix.TBSCertificateList)(nil),
}
