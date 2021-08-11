// export by github.com/goplus/interp/cmd/qexp

package binary

import (
	"encoding/binary"

	"github.com/goplus/interp"
)

func init() {
	interp.RegisterPackage("encoding/binary", extMap, typList)
}

var extMap = map[string]interface{}{
	"(encoding/binary.ByteOrder).PutUint16": (binary.ByteOrder).PutUint16,
	"(encoding/binary.ByteOrder).PutUint32": (binary.ByteOrder).PutUint32,
	"(encoding/binary.ByteOrder).PutUint64": (binary.ByteOrder).PutUint64,
	"(encoding/binary.ByteOrder).String":    (binary.ByteOrder).String,
	"(encoding/binary.ByteOrder).Uint16":    (binary.ByteOrder).Uint16,
	"(encoding/binary.ByteOrder).Uint32":    (binary.ByteOrder).Uint32,
	"(encoding/binary.ByteOrder).Uint64":    (binary.ByteOrder).Uint64,
	"encoding/binary.BigEndian":             &binary.BigEndian,
	"encoding/binary.LittleEndian":          &binary.LittleEndian,
	"encoding/binary.PutUvarint":            binary.PutUvarint,
	"encoding/binary.PutVarint":             binary.PutVarint,
	"encoding/binary.Read":                  binary.Read,
	"encoding/binary.ReadUvarint":           binary.ReadUvarint,
	"encoding/binary.ReadVarint":            binary.ReadVarint,
	"encoding/binary.Size":                  binary.Size,
	"encoding/binary.Uvarint":               binary.Uvarint,
	"encoding/binary.Varint":                binary.Varint,
	"encoding/binary.Write":                 binary.Write,
}

var typList = []interface{}{
	(*binary.ByteOrder)(nil),
}
