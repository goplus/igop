// export by github.com/goplus/gossa/cmd/qexp

package gob

import (
	"encoding/gob"

	"github.com/goplus/gossa"
)

func init() {
	gossa.RegisterPackage("encoding/gob", extMap, typList)
}

var extMap = map[string]interface{}{
	"(*encoding/gob.Decoder).Decode":      (*gob.Decoder).Decode,
	"(*encoding/gob.Decoder).DecodeValue": (*gob.Decoder).DecodeValue,
	"(*encoding/gob.Encoder).Encode":      (*gob.Encoder).Encode,
	"(*encoding/gob.Encoder).EncodeValue": (*gob.Encoder).EncodeValue,
	"(encoding/gob.GobDecoder).GobDecode": (gob.GobDecoder).GobDecode,
	"(encoding/gob.GobEncoder).GobEncode": (gob.GobEncoder).GobEncode,
	"encoding/gob.NewDecoder":             gob.NewDecoder,
	"encoding/gob.NewEncoder":             gob.NewEncoder,
	"encoding/gob.Register":               gob.Register,
	"encoding/gob.RegisterName":           gob.RegisterName,
}

var typList = []interface{}{
	(*gob.CommonType)(nil),
	(*gob.Decoder)(nil),
	(*gob.Encoder)(nil),
	(*gob.GobDecoder)(nil),
	(*gob.GobEncoder)(nil),
}
