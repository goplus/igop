// export by github.com/goplus/interp/cmd/qexp

// +build go1.16

package x509

import (
	"crypto/x509"

	"github.com/goplus/interp"
)

func init() {
	interp.RegisterPackage("crypto/x509", extMap_go116, typList_go116)
}

var extMap_go116 = map[string]interface{}{
	"(crypto/x509.SystemRootsError).Unwrap": (x509.SystemRootsError).Unwrap,
}

var typList_go116 = []interface{}{}
