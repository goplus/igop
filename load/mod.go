package load

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/mod/modfile"
)

func GetModulePath(dir string) (string, error) {
	mod, err := GOMOD(dir)
	if err != nil {
		return "", err
	}
	f, err := LoadModFile(mod)
	if err != nil {
		return "", err
	}
	if f.Module != nil {
		return f.Module.Mod.Path, nil
	}
	return "", errors.New("not found mod path")
}

func GOMOD(dir string) (file string, err error) {
	if dir == "" {
		dir = "."
	}
	if dir, err = filepath.Abs(dir); err != nil {
		return
	}
	for dir != "" {
		file = filepath.Join(dir, "go.mod")
		if fi, e := os.Lstat(file); e == nil && !fi.IsDir() {
			return
		}
		if dir, file = filepath.Split(strings.TrimRight(dir, "/\\")); file == "" {
			break
		}
	}
	return "", errors.New("not found go.mod")
}

func LoadModFile(file string) (*modfile.File, error) {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	fix := func(path, vers string) (resolved string, err error) {
		// do nothing
		return vers, nil
	}
	f, err := modfile.Parse(file, data, fix)
	if err != nil {
		return nil, fmt.Errorf("parse go.mod error %w", err)
	}
	return f, nil
}
