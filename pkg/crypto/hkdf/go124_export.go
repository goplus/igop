// export by github.com/goplus/igop/cmd/qexp

//go:build go1.24
// +build go1.24

package hkdf

import (
	_ "crypto/hkdf"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "hkdf",
		Path: "crypto/hkdf",
		Deps: map[string]string{
			"crypto/internal/fips140/hkdf": "hkdf",
			"crypto/internal/fips140hash":  "fips140hash",
			"crypto/internal/fips140only":  "fips140only",
			"errors":                       "errors",
			"hash":                         "hash",
		},
	})
}
