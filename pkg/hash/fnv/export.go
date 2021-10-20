// export by github.com/goplus/gossa/cmd/qexp

package fnv

import (
	"hash/fnv"

	"github.com/goplus/gossa"
)

func init() {
	gossa.RegisterPackage("hash/fnv", extMap, typList)
}

var extMap = map[string]interface{}{
	"hash/fnv.New128":  fnv.New128,
	"hash/fnv.New128a": fnv.New128a,
	"hash/fnv.New32":   fnv.New32,
	"hash/fnv.New32a":  fnv.New32a,
	"hash/fnv.New64":   fnv.New64,
	"hash/fnv.New64a":  fnv.New64a,
}

var typList = []interface{}{}
