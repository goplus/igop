// export by github.com/goplus/interp/cmd/qexp

//go:build go1.16
// +build go1.16

package zip

import (
	"archive/zip"

	"github.com/goplus/interp"
)

func init() {
	interp.RegisterPackage("archive/zip", extMap_go116, typList_go116)
}

var extMap_go116 = map[string]interface{}{
	"(*archive/zip.ReadCloser).Open": (*zip.ReadCloser).Open,
	"(*archive/zip.Reader).Open":     (*zip.Reader).Open,
}

var typList_go116 = []interface{}{}
