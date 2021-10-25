// export by github.com/goplus/gossa/cmd/qexp

package crc32

import (
	"hash/crc32"

	"github.com/goplus/gossa"
)

func init() {
	gossa.RegisterPackage("hash/crc32", extMap, typList)
}

var extMap = map[string]interface{}{
	"hash/crc32.Checksum":     crc32.Checksum,
	"hash/crc32.ChecksumIEEE": crc32.ChecksumIEEE,
	"hash/crc32.IEEETable":    &crc32.IEEETable,
	"hash/crc32.MakeTable":    crc32.MakeTable,
	"hash/crc32.New":          crc32.New,
	"hash/crc32.NewIEEE":      crc32.NewIEEE,
	"hash/crc32.Update":       crc32.Update,
}

var typList = []interface{}{
	(*crc32.Table)(nil),
}
