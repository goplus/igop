// export by github.com/goplus/gossa/cmd/qexp

package ascii85

import (
	"encoding/ascii85"

	"github.com/goplus/gossa"
)

func init() {
	gossa.RegisterPackage("encoding/ascii85", extMap, typList)
}

var extMap = map[string]interface{}{
	"(encoding/ascii85.CorruptInputError).Error": (ascii85.CorruptInputError).Error,
	"encoding/ascii85.Decode":                    ascii85.Decode,
	"encoding/ascii85.Encode":                    ascii85.Encode,
	"encoding/ascii85.MaxEncodedLen":             ascii85.MaxEncodedLen,
	"encoding/ascii85.NewDecoder":                ascii85.NewDecoder,
	"encoding/ascii85.NewEncoder":                ascii85.NewEncoder,
}

var typList = []interface{}{
	(*ascii85.CorruptInputError)(nil),
}
