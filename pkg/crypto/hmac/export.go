// export by github.com/goplus/gossa/cmd/qexp

package hmac

import (
	"crypto/hmac"

	"github.com/goplus/gossa"
)

func init() {
	gossa.RegisterPackage("crypto/hmac", extMap, typList)
}

var extMap = map[string]interface{}{
	"crypto/hmac.Equal": hmac.Equal,
	"crypto/hmac.New":   hmac.New,
}

var typList = []interface{}{}
