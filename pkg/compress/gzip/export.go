// export by github.com/goplus/interp/cmd/qexp

package gzip

import (
	"compress/gzip"

	"github.com/goplus/interp"
)

func init() {
	interp.RegisterPackage("compress/gzip", extMap, typList)
}

var extMap = map[string]interface{}{
	"(*compress/gzip.Reader).Close":       (*gzip.Reader).Close,
	"(*compress/gzip.Reader).Multistream": (*gzip.Reader).Multistream,
	"(*compress/gzip.Reader).Read":        (*gzip.Reader).Read,
	"(*compress/gzip.Reader).Reset":       (*gzip.Reader).Reset,
	"(*compress/gzip.Writer).Close":       (*gzip.Writer).Close,
	"(*compress/gzip.Writer).Flush":       (*gzip.Writer).Flush,
	"(*compress/gzip.Writer).Reset":       (*gzip.Writer).Reset,
	"(*compress/gzip.Writer).Write":       (*gzip.Writer).Write,
	"compress/gzip.ErrChecksum":           &gzip.ErrChecksum,
	"compress/gzip.ErrHeader":             &gzip.ErrHeader,
	"compress/gzip.NewReader":             gzip.NewReader,
	"compress/gzip.NewWriter":             gzip.NewWriter,
	"compress/gzip.NewWriterLevel":        gzip.NewWriterLevel,
}

var typList = []interface{}{
	(*gzip.Header)(nil),
	(*gzip.Reader)(nil),
	(*gzip.Writer)(nil),
}
