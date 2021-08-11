// export by github.com/goplus/interp/cmd/qexp

package md5

import (
	"crypto/md5"

	"github.com/goplus/interp"
)

func init() {
	interp.RegisterPackage("crypto/md5", extMap, typList)
}

var extMap = map[string]interface{}{
	"crypto/md5.New": md5.New,
	"crypto/md5.Sum": md5.Sum,
}

var typList = []interface{}{}
