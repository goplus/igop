// export by github.com/goplus/interp/cmd/qexp

// +build go1.15

package tls

import (
	"crypto/tls"

	"github.com/goplus/interp"
)

func init() {
	interp.RegisterPackage("crypto/tls", extMap_go115, typList_go115)
}

var extMap_go115 = map[string]interface{}{
	"(*crypto/tls.Dialer).Dial":           (*tls.Dialer).Dial,
	"(*crypto/tls.Dialer).DialContext":    (*tls.Dialer).DialContext,
	"(crypto/tls.ClientAuthType).String":  (tls.ClientAuthType).String,
	"(crypto/tls.CurveID).String":         (tls.CurveID).String,
	"(crypto/tls.SignatureScheme).String": (tls.SignatureScheme).String,
}

var typList_go115 = []interface{}{
	(*tls.Dialer)(nil),
}
