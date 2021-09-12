// export by github.com/goplus/interp/cmd/qexp

//go:build go1.15
// +build go1.15

package big

import (
	"math/big"

	"github.com/goplus/interp"
)

func init() {
	interp.RegisterPackage("math/big", extMap_go115, typList_go115)
}

var extMap_go115 = map[string]interface{}{
	"(*math/big.Int).FillBytes": (*big.Int).FillBytes,
}

var typList_go115 = []interface{}{}
