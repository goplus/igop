// export by github.com/goplus/gossa/cmd/qexp

package sha512

import (
	"crypto/sha512"

	"github.com/goplus/gossa"
)

func init() {
	gossa.RegisterPackage("crypto/sha512", extMap, typList)
}

var extMap = map[string]interface{}{
	"crypto/sha512.New":        sha512.New,
	"crypto/sha512.New384":     sha512.New384,
	"crypto/sha512.New512_224": sha512.New512_224,
	"crypto/sha512.New512_256": sha512.New512_256,
	"crypto/sha512.Sum384":     sha512.Sum384,
	"crypto/sha512.Sum512":     sha512.Sum512,
	"crypto/sha512.Sum512_224": sha512.Sum512_224,
	"crypto/sha512.Sum512_256": sha512.Sum512_256,
}

var typList = []interface{}{}
