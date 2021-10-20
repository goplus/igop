// export by github.com/goplus/gossa/cmd/qexp

package zip

import (
	"archive/zip"

	"github.com/goplus/gossa"
)

func init() {
	gossa.RegisterPackage("archive/zip", extMap, typList)
}

var extMap = map[string]interface{}{
	"(*archive/zip.File).DataOffset":                 (*zip.File).DataOffset,
	"(*archive/zip.File).FileInfo":                   (*zip.File).FileInfo,
	"(*archive/zip.File).ModTime":                    (*zip.File).ModTime,
	"(*archive/zip.File).Mode":                       (*zip.File).Mode,
	"(*archive/zip.File).Open":                       (*zip.File).Open,
	"(*archive/zip.File).SetModTime":                 (*zip.File).SetModTime,
	"(*archive/zip.File).SetMode":                    (*zip.File).SetMode,
	"(*archive/zip.FileHeader).FileInfo":             (*zip.FileHeader).FileInfo,
	"(*archive/zip.FileHeader).ModTime":              (*zip.FileHeader).ModTime,
	"(*archive/zip.FileHeader).Mode":                 (*zip.FileHeader).Mode,
	"(*archive/zip.FileHeader).SetModTime":           (*zip.FileHeader).SetModTime,
	"(*archive/zip.FileHeader).SetMode":              (*zip.FileHeader).SetMode,
	"(*archive/zip.ReadCloser).Close":                (*zip.ReadCloser).Close,
	"(*archive/zip.ReadCloser).RegisterDecompressor": (*zip.ReadCloser).RegisterDecompressor,
	"(*archive/zip.Reader).RegisterDecompressor":     (*zip.Reader).RegisterDecompressor,
	"(*archive/zip.Writer).Close":                    (*zip.Writer).Close,
	"(*archive/zip.Writer).Create":                   (*zip.Writer).Create,
	"(*archive/zip.Writer).CreateHeader":             (*zip.Writer).CreateHeader,
	"(*archive/zip.Writer).Flush":                    (*zip.Writer).Flush,
	"(*archive/zip.Writer).RegisterCompressor":       (*zip.Writer).RegisterCompressor,
	"(*archive/zip.Writer).SetComment":               (*zip.Writer).SetComment,
	"(*archive/zip.Writer).SetOffset":                (*zip.Writer).SetOffset,
	"archive/zip.ErrAlgorithm":                       &zip.ErrAlgorithm,
	"archive/zip.ErrChecksum":                        &zip.ErrChecksum,
	"archive/zip.ErrFormat":                          &zip.ErrFormat,
	"archive/zip.FileInfoHeader":                     zip.FileInfoHeader,
	"archive/zip.NewReader":                          zip.NewReader,
	"archive/zip.NewWriter":                          zip.NewWriter,
	"archive/zip.OpenReader":                         zip.OpenReader,
	"archive/zip.RegisterCompressor":                 zip.RegisterCompressor,
	"archive/zip.RegisterDecompressor":               zip.RegisterDecompressor,
}

var typList = []interface{}{
	(*zip.Compressor)(nil),
	(*zip.Decompressor)(nil),
	(*zip.File)(nil),
	(*zip.FileHeader)(nil),
	(*zip.ReadCloser)(nil),
	(*zip.Reader)(nil),
	(*zip.Writer)(nil),
}
