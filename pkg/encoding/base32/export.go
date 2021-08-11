// export by github.com/goplus/interp/cmd/qexp

package base32

import (
	"encoding/base32"

	"github.com/goplus/interp"
)

func init() {
	interp.RegisterPackage("encoding/base32", extMap, typList)
}

var extMap = map[string]interface{}{
	"(*encoding/base32.Encoding).Decode":         (*base32.Encoding).Decode,
	"(*encoding/base32.Encoding).DecodeString":   (*base32.Encoding).DecodeString,
	"(*encoding/base32.Encoding).DecodedLen":     (*base32.Encoding).DecodedLen,
	"(*encoding/base32.Encoding).Encode":         (*base32.Encoding).Encode,
	"(*encoding/base32.Encoding).EncodeToString": (*base32.Encoding).EncodeToString,
	"(*encoding/base32.Encoding).EncodedLen":     (*base32.Encoding).EncodedLen,
	"(encoding/base32.CorruptInputError).Error":  (base32.CorruptInputError).Error,
	"(encoding/base32.Encoding).WithPadding":     (base32.Encoding).WithPadding,
	"encoding/base32.HexEncoding":                &base32.HexEncoding,
	"encoding/base32.NewDecoder":                 base32.NewDecoder,
	"encoding/base32.NewEncoder":                 base32.NewEncoder,
	"encoding/base32.NewEncoding":                base32.NewEncoding,
	"encoding/base32.StdEncoding":                &base32.StdEncoding,
}

var typList = []interface{}{
	(*base32.CorruptInputError)(nil),
	(*base32.Encoding)(nil),
}
