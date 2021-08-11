// export by github.com/goplus/interp/cmd/qexp

package sha256

import (
	"crypto/sha256"

	"github.com/goplus/interp"
)

func init() {
	interp.RegisterPackage("crypto/sha256", extMap, typList)
}

var extMap = map[string]interface{}{
	"crypto/sha256.New":    sha256.New,
	"crypto/sha256.New224": sha256.New224,
	"crypto/sha256.Sum224": sha256.Sum224,
	"crypto/sha256.Sum256": sha256.Sum256,
}

var typList = []interface{}{}
