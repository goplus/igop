// export by github.com/goplus/gossa/cmd/qexp

//go:build go1.15
// +build go1.15

package elliptic

import (
	"crypto/elliptic"

	"github.com/goplus/gossa"
)

func init() {
	gossa.RegisterPackage("crypto/elliptic", extMap_go115, typList_go115)
}

var extMap_go115 = map[string]interface{}{
	"crypto/elliptic.MarshalCompressed":   elliptic.MarshalCompressed,
	"crypto/elliptic.UnmarshalCompressed": elliptic.UnmarshalCompressed,
}

var typList_go115 = []interface{}{}
