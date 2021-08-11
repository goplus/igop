// export by github.com/goplus/interp/cmd/qexp

package hmac

import (
	"crypto/hmac"

	"github.com/goplus/interp"
)

func init() {
	interp.RegisterPackage("crypto/hmac", extMap, typList)
}

var extMap = map[string]interface{}{
	"crypto/hmac.Equal": hmac.Equal,
	"crypto/hmac.New":   hmac.New,
}

var typList = []interface{}{}
