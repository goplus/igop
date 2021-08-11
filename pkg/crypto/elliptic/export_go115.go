// export by github.com/goplus/interp/cmd/qexp

// +build go1.15

package elliptic

import (
	"crypto/elliptic"

	"github.com/goplus/interp"
)

func init() {
	interp.RegisterPackage("crypto/elliptic", extMap_go115, typList_go115)
}

var extMap_go115 = map[string]interface{}{
	"crypto/elliptic.MarshalCompressed":   elliptic.MarshalCompressed,
	"crypto/elliptic.UnmarshalCompressed": elliptic.UnmarshalCompressed,
}

var typList_go115 = []interface{}{}
