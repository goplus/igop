// export by github.com/goplus/gossa/cmd/qexp

package pem

import (
	"encoding/pem"

	"github.com/goplus/gossa"
)

func init() {
	gossa.RegisterPackage("encoding/pem", extMap, typList)
}

var extMap = map[string]interface{}{
	"encoding/pem.Decode":         pem.Decode,
	"encoding/pem.Encode":         pem.Encode,
	"encoding/pem.EncodeToMemory": pem.EncodeToMemory,
}

var typList = []interface{}{
	(*pem.Block)(nil),
}
