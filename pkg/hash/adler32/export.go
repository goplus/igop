// export by github.com/goplus/gossa/cmd/qexp

package adler32

import (
	"hash/adler32"

	"github.com/goplus/gossa"
)

func init() {
	gossa.RegisterPackage("hash/adler32", extMap, typList)
}

var extMap = map[string]interface{}{
	"hash/adler32.Checksum": adler32.Checksum,
	"hash/adler32.New":      adler32.New,
}

var typList = []interface{}{}
