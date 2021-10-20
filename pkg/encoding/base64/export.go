// export by github.com/goplus/gossa/cmd/qexp

package base64

import (
	"encoding/base64"

	"github.com/goplus/gossa"
)

func init() {
	gossa.RegisterPackage("encoding/base64", extMap, typList)
}

var extMap = map[string]interface{}{
	"(*encoding/base64.Encoding).Decode":         (*base64.Encoding).Decode,
	"(*encoding/base64.Encoding).DecodeString":   (*base64.Encoding).DecodeString,
	"(*encoding/base64.Encoding).DecodedLen":     (*base64.Encoding).DecodedLen,
	"(*encoding/base64.Encoding).Encode":         (*base64.Encoding).Encode,
	"(*encoding/base64.Encoding).EncodeToString": (*base64.Encoding).EncodeToString,
	"(*encoding/base64.Encoding).EncodedLen":     (*base64.Encoding).EncodedLen,
	"(encoding/base64.CorruptInputError).Error":  (base64.CorruptInputError).Error,
	"(encoding/base64.Encoding).Strict":          (base64.Encoding).Strict,
	"(encoding/base64.Encoding).WithPadding":     (base64.Encoding).WithPadding,
	"encoding/base64.NewDecoder":                 base64.NewDecoder,
	"encoding/base64.NewEncoder":                 base64.NewEncoder,
	"encoding/base64.NewEncoding":                base64.NewEncoding,
	"encoding/base64.RawStdEncoding":             &base64.RawStdEncoding,
	"encoding/base64.RawURLEncoding":             &base64.RawURLEncoding,
	"encoding/base64.StdEncoding":                &base64.StdEncoding,
	"encoding/base64.URLEncoding":                &base64.URLEncoding,
}

var typList = []interface{}{
	(*base64.CorruptInputError)(nil),
	(*base64.Encoding)(nil),
}
