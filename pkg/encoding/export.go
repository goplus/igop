// export by github.com/goplus/interp/cmd/qexp

package encoding

import (
	"encoding"

	"github.com/goplus/interp"
)

func init() {
	interp.RegisterPackage("encoding", extMap, typList)
}

var extMap = map[string]interface{}{
	"(encoding.BinaryMarshaler).MarshalBinary":     (encoding.BinaryMarshaler).MarshalBinary,
	"(encoding.BinaryUnmarshaler).UnmarshalBinary": (encoding.BinaryUnmarshaler).UnmarshalBinary,
	"(encoding.TextMarshaler).MarshalText":         (encoding.TextMarshaler).MarshalText,
	"(encoding.TextUnmarshaler).UnmarshalText":     (encoding.TextUnmarshaler).UnmarshalText,
}

var typList = []interface{}{
	(*encoding.BinaryMarshaler)(nil),
	(*encoding.BinaryUnmarshaler)(nil),
	(*encoding.TextMarshaler)(nil),
	(*encoding.TextUnmarshaler)(nil),
}
