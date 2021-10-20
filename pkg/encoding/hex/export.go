// export by github.com/goplus/gossa/cmd/qexp

package hex

import (
	"encoding/hex"

	"github.com/goplus/gossa"
)

func init() {
	gossa.RegisterPackage("encoding/hex", extMap, typList)
}

var extMap = map[string]interface{}{
	"(encoding/hex.InvalidByteError).Error": (hex.InvalidByteError).Error,
	"encoding/hex.Decode":                   hex.Decode,
	"encoding/hex.DecodeString":             hex.DecodeString,
	"encoding/hex.DecodedLen":               hex.DecodedLen,
	"encoding/hex.Dump":                     hex.Dump,
	"encoding/hex.Dumper":                   hex.Dumper,
	"encoding/hex.Encode":                   hex.Encode,
	"encoding/hex.EncodeToString":           hex.EncodeToString,
	"encoding/hex.EncodedLen":               hex.EncodedLen,
	"encoding/hex.ErrLength":                &hex.ErrLength,
	"encoding/hex.NewDecoder":               hex.NewDecoder,
	"encoding/hex.NewEncoder":               hex.NewEncoder,
}

var typList = []interface{}{
	(*hex.InvalidByteError)(nil),
}
