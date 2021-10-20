// export by github.com/goplus/gossa/cmd/qexp

package tls

import (
	"crypto/tls"

	"github.com/goplus/gossa"
)

func init() {
	gossa.RegisterPackage("crypto/tls", extMap, typList)
}

var extMap = map[string]interface{}{
	"(*crypto/tls.CertificateRequestInfo).SupportsCertificate": (*tls.CertificateRequestInfo).SupportsCertificate,
	"(*crypto/tls.ClientHelloInfo).SupportsCertificate":        (*tls.ClientHelloInfo).SupportsCertificate,
	"(*crypto/tls.Config).BuildNameToCertificate":              (*tls.Config).BuildNameToCertificate,
	"(*crypto/tls.Config).Clone":                               (*tls.Config).Clone,
	"(*crypto/tls.Config).SetSessionTicketKeys":                (*tls.Config).SetSessionTicketKeys,
	"(*crypto/tls.Conn).Close":                                 (*tls.Conn).Close,
	"(*crypto/tls.Conn).CloseWrite":                            (*tls.Conn).CloseWrite,
	"(*crypto/tls.Conn).ConnectionState":                       (*tls.Conn).ConnectionState,
	"(*crypto/tls.Conn).Handshake":                             (*tls.Conn).Handshake,
	"(*crypto/tls.Conn).LocalAddr":                             (*tls.Conn).LocalAddr,
	"(*crypto/tls.Conn).OCSPResponse":                          (*tls.Conn).OCSPResponse,
	"(*crypto/tls.Conn).Read":                                  (*tls.Conn).Read,
	"(*crypto/tls.Conn).RemoteAddr":                            (*tls.Conn).RemoteAddr,
	"(*crypto/tls.Conn).SetDeadline":                           (*tls.Conn).SetDeadline,
	"(*crypto/tls.Conn).SetReadDeadline":                       (*tls.Conn).SetReadDeadline,
	"(*crypto/tls.Conn).SetWriteDeadline":                      (*tls.Conn).SetWriteDeadline,
	"(*crypto/tls.Conn).VerifyHostname":                        (*tls.Conn).VerifyHostname,
	"(*crypto/tls.Conn).Write":                                 (*tls.Conn).Write,
	"(*crypto/tls.ConnectionState).ExportKeyingMaterial":       (*tls.ConnectionState).ExportKeyingMaterial,
	"(crypto/tls.ClientSessionCache).Get":                      (tls.ClientSessionCache).Get,
	"(crypto/tls.ClientSessionCache).Put":                      (tls.ClientSessionCache).Put,
	"(crypto/tls.RecordHeaderError).Error":                     (tls.RecordHeaderError).Error,
	"crypto/tls.CipherSuiteName":                               tls.CipherSuiteName,
	"crypto/tls.CipherSuites":                                  tls.CipherSuites,
	"crypto/tls.Client":                                        tls.Client,
	"crypto/tls.Dial":                                          tls.Dial,
	"crypto/tls.DialWithDialer":                                tls.DialWithDialer,
	"crypto/tls.InsecureCipherSuites":                          tls.InsecureCipherSuites,
	"crypto/tls.Listen":                                        tls.Listen,
	"crypto/tls.LoadX509KeyPair":                               tls.LoadX509KeyPair,
	"crypto/tls.NewLRUClientSessionCache":                      tls.NewLRUClientSessionCache,
	"crypto/tls.NewListener":                                   tls.NewListener,
	"crypto/tls.Server":                                        tls.Server,
	"crypto/tls.X509KeyPair":                                   tls.X509KeyPair,
}

var typList = []interface{}{
	(*tls.Certificate)(nil),
	(*tls.CertificateRequestInfo)(nil),
	(*tls.CipherSuite)(nil),
	(*tls.ClientAuthType)(nil),
	(*tls.ClientHelloInfo)(nil),
	(*tls.ClientSessionCache)(nil),
	(*tls.ClientSessionState)(nil),
	(*tls.Config)(nil),
	(*tls.Conn)(nil),
	(*tls.ConnectionState)(nil),
	(*tls.CurveID)(nil),
	(*tls.RecordHeaderError)(nil),
	(*tls.RenegotiationSupport)(nil),
	(*tls.SignatureScheme)(nil),
}
