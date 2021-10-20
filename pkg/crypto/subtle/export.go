// export by github.com/goplus/gossa/cmd/qexp

package subtle

import (
	"crypto/subtle"

	"github.com/goplus/gossa"
)

func init() {
	gossa.RegisterPackage("crypto/subtle", extMap, typList)
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
