// export by github.com/goplus/interp/cmd/qexp

package subtle

import (
	"crypto/subtle"

	"github.com/goplus/interp"
)

func init() {
	interp.RegisterPackage("crypto/subtle", extMap, typList)
}

var extMap = map[string]interface{}{
	"crypto/subtle.ConstantTimeByteEq":   subtle.ConstantTimeByteEq,
	"crypto/subtle.ConstantTimeCompare":  subtle.ConstantTimeCompare,
	"crypto/subtle.ConstantTimeCopy":     subtle.ConstantTimeCopy,
	"crypto/subtle.ConstantTimeEq":       subtle.ConstantTimeEq,
	"crypto/subtle.ConstantTimeLessOrEq": subtle.ConstantTimeLessOrEq,
	"crypto/subtle.ConstantTimeSelect":   subtle.ConstantTimeSelect,
}

var typList = []interface{}{}
