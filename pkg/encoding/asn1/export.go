// export by github.com/goplus/interp/cmd/qexp

package asn1

import (
	"encoding/asn1"

	"github.com/goplus/interp"
)

func init() {
	interp.RegisterPackage("encoding/asn1", extMap, typList)
}

var extMap = map[string]interface{}{
	"(encoding/asn1.BitString).At":            (asn1.BitString).At,
	"(encoding/asn1.BitString).RightAlign":    (asn1.BitString).RightAlign,
	"(encoding/asn1.ObjectIdentifier).Equal":  (asn1.ObjectIdentifier).Equal,
	"(encoding/asn1.ObjectIdentifier).String": (asn1.ObjectIdentifier).String,
	"(encoding/asn1.StructuralError).Error":   (asn1.StructuralError).Error,
	"(encoding/asn1.SyntaxError).Error":       (asn1.SyntaxError).Error,
	"encoding/asn1.Marshal":                   asn1.Marshal,
	"encoding/asn1.MarshalWithParams":         asn1.MarshalWithParams,
	"encoding/asn1.NullBytes":                 &asn1.NullBytes,
	"encoding/asn1.NullRawValue":              &asn1.NullRawValue,
	"encoding/asn1.Unmarshal":                 asn1.Unmarshal,
	"encoding/asn1.UnmarshalWithParams":       asn1.UnmarshalWithParams,
}

var typList = []interface{}{
	(*asn1.BitString)(nil),
	(*asn1.Enumerated)(nil),
	(*asn1.Flag)(nil),
	(*asn1.ObjectIdentifier)(nil),
	(*asn1.RawContent)(nil),
	(*asn1.RawValue)(nil),
	(*asn1.StructuralError)(nil),
	(*asn1.SyntaxError)(nil),
}
