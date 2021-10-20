// export by github.com/goplus/gossa/cmd/qexp

package tar

import (
	"archive/tar"

	"github.com/goplus/gossa"
)

func init() {
	gossa.RegisterPackage("archive/tar", extMap, typList)
}

var extMap = map[string]interface{}{
	"(*archive/tar.Header).FileInfo":    (*tar.Header).FileInfo,
	"(*archive/tar.Reader).Next":        (*tar.Reader).Next,
	"(*archive/tar.Reader).Read":        (*tar.Reader).Read,
	"(*archive/tar.Writer).Close":       (*tar.Writer).Close,
	"(*archive/tar.Writer).Flush":       (*tar.Writer).Flush,
	"(*archive/tar.Writer).Write":       (*tar.Writer).Write,
	"(*archive/tar.Writer).WriteHeader": (*tar.Writer).WriteHeader,
	"(archive/tar.Format).String":       (tar.Format).String,
	"archive/tar.ErrFieldTooLong":       &tar.ErrFieldTooLong,
	"archive/tar.ErrHeader":             &tar.ErrHeader,
	"archive/tar.ErrWriteAfterClose":    &tar.ErrWriteAfterClose,
	"archive/tar.ErrWriteTooLong":       &tar.ErrWriteTooLong,
	"archive/tar.FileInfoHeader":        tar.FileInfoHeader,
	"archive/tar.NewReader":             tar.NewReader,
	"archive/tar.NewWriter":             tar.NewWriter,
}

var typList = []interface{}{
	(*tar.Format)(nil),
	(*tar.Header)(nil),
	(*tar.Reader)(nil),
	(*tar.Writer)(nil),
}
