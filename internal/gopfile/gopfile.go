package gopfile

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"

	"github.com/goplus/gop/cl"
	"github.com/goplus/gop/parser"
	"github.com/goplus/gop/token"
	"github.com/goplus/gox"
)

func Build(filename string, src interface{}) ([]byte, error) {
	srcDir, _ := filepath.Split(filename)
	fset := token.NewFileSet()
	pkgs, err := parser.Parse(fset, filename, src, 0)
	if err != nil {
		return nil, err
	}

	mainPkg, ok := pkgs["main"]
	if !ok {
		return nil, fmt.Errorf("not a main package")
	}

	modDir, noCacheFile := findGoModDir(srcDir)
	conf := &cl.Config{
		Dir: modDir, TargetDir: srcDir, Fset: fset, CacheLoadPkgs: true, PersistLoadPkgs: !noCacheFile}
	out, err := cl.NewPackage("", mainPkg, conf)
	if err != nil {
		return nil, err
	}
	var buf bytes.Buffer
	if err = gox.WriteTo(&buf, out, false); err != nil {
		return nil, err
	}
	conf.PkgsLoader.Save()
	return buf.Bytes(), nil
}

func findGoModFile(dir string) (modfile string, noCacheFile bool, err error) {
	modfile, err = cl.FindGoModFile(dir)
	if err != nil {
		home := os.Getenv("HOME")
		modfile = home + "/gop/go.mod"
		if fi, e := os.Lstat(modfile); e == nil && !fi.IsDir() {
			return modfile, true, nil
		}
		modfile = home + "/goplus/go.mod"
		if fi, e := os.Lstat(modfile); e == nil && !fi.IsDir() {
			return modfile, true, nil
		}
	}
	return
}

func findGoModDir(dir string) (string, bool) {
	modfile, nocachefile, err := findGoModFile(dir)
	if err != nil {
		return dir, true
	}
	return filepath.Dir(modfile), nocachefile
}
