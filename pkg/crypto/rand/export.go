// export by github.com/goplus/interp/cmd/qexp

package rand

import (
	"crypto/rand"

	"github.com/goplus/interp"
)

func init() {
	interp.RegisterPackage("crypto/rand", extMap, typList)
}

var extMap = map[string]interface{}{
	"crypto/rand.Int":    rand.Int,
	"crypto/rand.Prime":  rand.Prime,
	"crypto/rand.Read":   rand.Read,
	"crypto/rand.Reader": &rand.Reader,
}

var typList = []interface{}{}
