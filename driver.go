package igop

import (
	"fmt"
	"go/build"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/visualfc/gomod"
)

type ListDriver struct {
	init bool
	root string
	pkgs map[string]string // path -> dir
}

func (d *ListDriver) Lookup(root string, path string) (dir string, found bool) {
	if !d.init || d.root != root {
		d.init = true
		d.root = root
		err := d.Parse(root)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
	dir, found = d.pkgs[path]
	if found {
		return
	}
	var list []string
	for k := range d.pkgs {
		if strings.HasPrefix(path, k+"/") {
			list = append(list, k)
		}
	}
	switch len(list) {
	case 0:
	case 1:
		v := list[0]
		dir, found = filepath.Join(d.pkgs[v], path[len(v+"/"):]), true
	default:
		// check path/v2
		sort.Slice(list, func(i, j int) bool {
			return list[i] > list[j]
		})
		v := list[0]
		dir, found = filepath.Join(d.pkgs[v], path[len(v+"/"):]), true
	}
	return
}

func (d *ListDriver) Parse(root string) error {
	data, err := runGoCommand(root, "list", "-deps", "-e", "-f={{.ImportPath}}={{.Dir}}", ".")
	if err != nil {
		return err
	}
	d.pkgs = make(map[string]string)
	for _, line := range strings.Split(string(data), "\n") {
		pos := strings.Index(line, "=")
		if pos != -1 {
			d.pkgs[line[:pos]] = line[pos+1:]
		}
	}
	return nil
}

type ModuleDriver struct {
	init bool
	root string
	mod  *gomod.Package
}

func (d *ModuleDriver) Lookup(root string, path string) (dir string, found bool) {
	if !d.init || d.root != root {
		d.init = true
		d.root = root
		var err error
		d.mod, err = gomod.Load(root, &build.Default)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
	_, dir, found = d.mod.Lookup(path)
	return
}
