// export by github.com/goplus/gossa/cmd/qexp

package crc64

import (
	"hash/crc64"

	"github.com/goplus/gossa"
)

func init() {
	gossa.RegisterPackage("hash/crc64", extMap, typList)
}

var extMap = map[string]interface{}{
	"hash/crc64.Checksum":  crc64.Checksum,
	"hash/crc64.MakeTable": crc64.MakeTable,
	"hash/crc64.New":       crc64.New,
	"hash/crc64.Update":    crc64.Update,
}

var typList = []interface{}{
	(*crc64.Table)(nil),
}
