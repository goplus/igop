//go:build go1.16
// +build go1.16

package load

import (
	"io/ioutil"
	"path/filepath"

	"github.com/goplus/igop"
	"github.com/goplus/igop/gopbuild"
)

const (
	SupportGop = true
)

func BuildGopDir(ctx *igop.Context, path string) error {
	data, err := gopbuild.BuildDir(ctx, path)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filepath.Join(path, "gop_autogen.go"), data, 0666)
}

func IsGopProject(dir string) bool {
	return ContainsExt(dir, ".gop")
}
