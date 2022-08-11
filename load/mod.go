package load

import (
	"github.com/goplus/mod"
	"github.com/goplus/mod/modload"
)

func GetModulePath(dir string) (string, error) {
	m, err := modload.Load(dir, mod.GoModOnly)
	if err != nil {
		return "", err
	}
	return m.Path(), nil
}
