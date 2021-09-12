// export by github.com/goplus/interp/cmd/qexp

//go:build go1.16
// +build go1.16

package io

import (
	"io"

	"github.com/goplus/interp"
)

func init() {
	interp.RegisterPackage("io", extMap_go116, typList_go116)
}

var extMap_go116 = map[string]interface{}{
	"(io.ReadSeekCloser).Close": (io.ReadSeekCloser).Close,
	"(io.ReadSeekCloser).Read":  (io.ReadSeekCloser).Read,
	"(io.ReadSeekCloser).Seek":  (io.ReadSeekCloser).Seek,
	"io.Discard":                &io.Discard,
	"io.NopCloser":              io.NopCloser,
	"io.ReadAll":                io.ReadAll,
}

var typList_go116 = []interface{}{
	(*io.ReadSeekCloser)(nil),
}
