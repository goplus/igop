// export by github.com/goplus/interp/cmd/qexp

package json

import (
	"encoding/json"

	"github.com/goplus/interp"
)

func init() {
	interp.RegisterPackage("encoding/json", extMap, typList)
}

var extMap = map[string]interface{}{
	"(*encoding/json.Decoder).Buffered":              (*json.Decoder).Buffered,
	"(*encoding/json.Decoder).Decode":                (*json.Decoder).Decode,
	"(*encoding/json.Decoder).DisallowUnknownFields": (*json.Decoder).DisallowUnknownFields,
	"(*encoding/json.Decoder).InputOffset":           (*json.Decoder).InputOffset,
	"(*encoding/json.Decoder).More":                  (*json.Decoder).More,
	"(*encoding/json.Decoder).Token":                 (*json.Decoder).Token,
	"(*encoding/json.Decoder).UseNumber":             (*json.Decoder).UseNumber,
	"(*encoding/json.Encoder).Encode":                (*json.Encoder).Encode,
	"(*encoding/json.Encoder).SetEscapeHTML":         (*json.Encoder).SetEscapeHTML,
	"(*encoding/json.Encoder).SetIndent":             (*json.Encoder).SetIndent,
	"(*encoding/json.InvalidUTF8Error).Error":        (*json.InvalidUTF8Error).Error,
	"(*encoding/json.InvalidUnmarshalError).Error":   (*json.InvalidUnmarshalError).Error,
	"(*encoding/json.MarshalerError).Error":          (*json.MarshalerError).Error,
	"(*encoding/json.MarshalerError).Unwrap":         (*json.MarshalerError).Unwrap,
	"(*encoding/json.RawMessage).MarshalJSON":        (*json.RawMessage).MarshalJSON,
	"(*encoding/json.RawMessage).UnmarshalJSON":      (*json.RawMessage).UnmarshalJSON,
	"(*encoding/json.SyntaxError).Error":             (*json.SyntaxError).Error,
	"(*encoding/json.UnmarshalFieldError).Error":     (*json.UnmarshalFieldError).Error,
	"(*encoding/json.UnmarshalTypeError).Error":      (*json.UnmarshalTypeError).Error,
	"(*encoding/json.UnsupportedTypeError).Error":    (*json.UnsupportedTypeError).Error,
	"(*encoding/json.UnsupportedValueError).Error":   (*json.UnsupportedValueError).Error,
	"(encoding/json.Delim).String":                   (json.Delim).String,
	"(encoding/json.Marshaler).MarshalJSON":          (json.Marshaler).MarshalJSON,
	"(encoding/json.Number).Float64":                 (json.Number).Float64,
	"(encoding/json.Number).Int64":                   (json.Number).Int64,
	"(encoding/json.Number).String":                  (json.Number).String,
	"(encoding/json.Unmarshaler).UnmarshalJSON":      (json.Unmarshaler).UnmarshalJSON,
	"encoding/json.Compact":                          json.Compact,
	"encoding/json.HTMLEscape":                       json.HTMLEscape,
	"encoding/json.Indent":                           json.Indent,
	"encoding/json.Marshal":                          json.Marshal,
	"encoding/json.MarshalIndent":                    json.MarshalIndent,
	"encoding/json.NewDecoder":                       json.NewDecoder,
	"encoding/json.NewEncoder":                       json.NewEncoder,
	"encoding/json.Unmarshal":                        json.Unmarshal,
	"encoding/json.Valid":                            json.Valid,
}

var typList = []interface{}{
	(*json.Decoder)(nil),
	(*json.Delim)(nil),
	(*json.Encoder)(nil),
	(*json.InvalidUTF8Error)(nil),
	(*json.InvalidUnmarshalError)(nil),
	(*json.Marshaler)(nil),
	(*json.MarshalerError)(nil),
	(*json.Number)(nil),
	(*json.RawMessage)(nil),
	(*json.SyntaxError)(nil),
	(*json.Token)(nil),
	(*json.UnmarshalFieldError)(nil),
	(*json.UnmarshalTypeError)(nil),
	(*json.Unmarshaler)(nil),
	(*json.UnsupportedTypeError)(nil),
	(*json.UnsupportedValueError)(nil),
}
