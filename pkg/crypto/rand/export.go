// export by github.com/goplus/gossa/cmd/qexp

package rand

import (
	"crypto/rand"

	"github.com/goplus/gossa"
)

func init() {
	gossa.RegisterPackage("crypto/rand", extMap, typList)
}

var extMap = map[string]interface{}{
	"crypto/rand.Int":    rand.Int,
	"crypto/rand.Prime":  rand.Prime,
	"crypto/rand.Read":   rand.Read,
	"crypto/rand.Reader": &rand.Reader,
}

var typList = []interface{}{}
