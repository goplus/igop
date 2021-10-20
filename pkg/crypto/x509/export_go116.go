// export by github.com/goplus/gossa/cmd/qexp

//go:build go1.16
// +build go1.16

package x509

import (
	"crypto/x509"

	"github.com/goplus/gossa"
)

func init() {
	gossa.RegisterPackage("crypto/x509", extMap_go116, typList_go116)
}

var extMap_go116 = map[string]interface{}{
	"(crypto/x509.SystemRootsError).Unwrap": (x509.SystemRootsError).Unwrap,
}

var typList_go116 = []interface{}{}
