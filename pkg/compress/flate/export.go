// export by github.com/goplus/gossa/cmd/qexp

package flate

import (
	"compress/flate"

	"github.com/goplus/gossa"
)

func init() {
	gossa.RegisterPackage("compress/flate", extMap, typList)
}

var extMap = map[string]interface{}{
	"(*compress/flate.ReadError).Error":        (*flate.ReadError).Error,
	"(*compress/flate.WriteError).Error":       (*flate.WriteError).Error,
	"(*compress/flate.Writer).Close":           (*flate.Writer).Close,
	"(*compress/flate.Writer).Flush":           (*flate.Writer).Flush,
	"(*compress/flate.Writer).Reset":           (*flate.Writer).Reset,
	"(*compress/flate.Writer).Write":           (*flate.Writer).Write,
	"(compress/flate.CorruptInputError).Error": (flate.CorruptInputError).Error,
	"(compress/flate.InternalError).Error":     (flate.InternalError).Error,
	"(compress/flate.Reader).Read":             (flate.Reader).Read,
	"(compress/flate.Reader).ReadByte":         (flate.Reader).ReadByte,
	"(compress/flate.Resetter).Reset":          (flate.Resetter).Reset,
	"compress/flate.NewReader":                 flate.NewReader,
	"compress/flate.NewReaderDict":             flate.NewReaderDict,
	"compress/flate.NewWriter":                 flate.NewWriter,
	"compress/flate.NewWriterDict":             flate.NewWriterDict,
}

var typList = []interface{}{
	(*flate.CorruptInputError)(nil),
	(*flate.InternalError)(nil),
	(*flate.ReadError)(nil),
	(*flate.Reader)(nil),
	(*flate.Resetter)(nil),
	(*flate.WriteError)(nil),
	(*flate.Writer)(nil),
}
