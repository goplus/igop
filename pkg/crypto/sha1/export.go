// export by github.com/goplus/gossa/cmd/qexp

package sha1

import (
	"crypto/sha1"

	"github.com/goplus/gossa"
)

func init() {
	gossa.RegisterPackage("crypto/sha1", extMap, typList)
}

var extMap = map[string]interface{}{
	"crypto/sha1.New": sha1.New,
	"crypto/sha1.Sum": sha1.Sum,
}

var typList = []interface{}{}
