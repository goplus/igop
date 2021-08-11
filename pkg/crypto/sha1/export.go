// export by github.com/goplus/interp/cmd/qexp

package sha1

import (
	"crypto/sha1"

	"github.com/goplus/interp"
)

func init() {
	interp.RegisterPackage("crypto/sha1", extMap, typList)
}

var extMap = map[string]interface{}{
	"crypto/sha1.New": sha1.New,
	"crypto/sha1.Sum": sha1.Sum,
}

var typList = []interface{}{}
