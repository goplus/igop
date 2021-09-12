// export by github.com/goplus/interp/cmd/qexp

//go:build go1.15
// +build go1.15

package x509

import (
	"crypto/x509"

	"github.com/goplus/interp"
)

func init() {
	interp.RegisterPackage("crypto/x509", extMap_go115, typList_go115)
}

var extMap_go115 = map[string]interface{}{
	"crypto/x509.CreateRevocationList": x509.CreateRevocationList,
}

var typList_go115 = []interface{}{
	(*x509.RevocationList)(nil),
}
