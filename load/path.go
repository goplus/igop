package load

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"

	"golang.org/x/mod/modfile"
)

func GetImportPath(pkgName string, dir string) (string, error) {
	if BuildMod == "mod" {
		data, err := runGoCommand(dir, "list", "-e", "-mod=mod")
		if err != nil {
			return "", err
		}
		return string(bytes.TrimSuffix(data, []byte{'\n'})), nil
	}
	dir, err := absDir(dir)
	if err != nil {
		return "", err
	}
	if pkgName == "" || pkgName == "main" {
		_, pkgName = filepath.Split(dir)
	}
	mod, found := findModule(dir)
	if !found {
		return pkgName, nil
	}
	f, err := LoadModFile(mod)
	if err != nil {
		return "", err
	}
	root := filepath.Dir(f.Syntax.Name)
	modPath := f.Module.Mod.Path
	dir = filepath.ToSlash(dir)
	root = filepath.ToSlash(root)
	if dir == root {
		return modPath, nil
	}
	_dir, _ := path.Split(dir[len(root)+1:])
	return path.Join(modPath, _dir, pkgName), nil
}

func absDir(dir string) (string, error) {
	if dir == "" {
		dir = "."
	}
	return filepath.Abs(dir)
}

func findModule(dir string) (file string, found bool) {
	for dir != "" {
		file = filepath.Join(dir, "go.mod")
		if fi, e := os.Lstat(file); e == nil && !fi.IsDir() {
			found = true
			return
		}
		if dir, file = filepath.Split(strings.TrimRight(dir, "/\\")); file == "" {
			break
		}
	}
	return
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
	if f.Module == nil {
		return nil, errors.New("no module declaration in go.mod.")
	}
	return f, nil
}
