// export by github.com/goplus/gossa/cmd/qexp

package zlib

import (
	"compress/zlib"

	"github.com/goplus/gossa"
)

func init() {
	gossa.RegisterPackage("compress/zlib", extMap, typList)
}

var extMap = map[string]interface{}{
	"(*compress/zlib.Writer).Close":    (*zlib.Writer).Close,
	"(*compress/zlib.Writer).Flush":    (*zlib.Writer).Flush,
	"(*compress/zlib.Writer).Reset":    (*zlib.Writer).Reset,
	"(*compress/zlib.Writer).Write":    (*zlib.Writer).Write,
	"(compress/zlib.Resetter).Reset":   (zlib.Resetter).Reset,
	"compress/zlib.ErrChecksum":        &zlib.ErrChecksum,
	"compress/zlib.ErrDictionary":      &zlib.ErrDictionary,
	"compress/zlib.ErrHeader":          &zlib.ErrHeader,
	"compress/zlib.NewReader":          zlib.NewReader,
	"compress/zlib.NewReaderDict":      zlib.NewReaderDict,
	"compress/zlib.NewWriter":          zlib.NewWriter,
	"compress/zlib.NewWriterLevel":     zlib.NewWriterLevel,
	"compress/zlib.NewWriterLevelDict": zlib.NewWriterLevelDict,
}

var typList = []interface{}{
	(*zlib.Resetter)(nil),
	(*zlib.Writer)(nil),
}
