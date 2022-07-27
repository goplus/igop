package igop

import (
	"fmt"
	"go/importer"
	"go/types"
)

type Importer struct {
	ctx         *Context
	pkgs        map[string]*types.Package
	importing   map[string]bool
	defaultImpl types.Importer
}

func NewImporter(ctx *Context) *Importer {
	return &Importer{
		ctx:         ctx,
		pkgs:        make(map[string]*types.Package),
		importing:   make(map[string]bool),
		defaultImpl: importer.Default(),
	}
}

func (i *Importer) Import(path string) (*types.Package, error) {
	if pkg, ok := i.pkgs[path]; ok {
		return pkg, nil
	}
	if i.importing[path] {
		return nil, fmt.Errorf("cycle importing package %q", path)
	}
	i.importing[path] = true
	defer func() {
		i.importing[path] = false
	}()
	if pkg, err := i.ctx.Loader.Import(path); err == nil {
		i.pkgs[path] = pkg
		return pkg, nil
	}
	if pkg, ok := i.ctx.pkgs[path]; ok {
		if !pkg.Package.Complete() {
			if err := pkg.Load(); err != nil {
				return nil, err
			}
		}
		i.pkgs[path] = pkg.Package
		return pkg.Package, nil
	}
	if dir, found := i.ctx.lookupPath(path); found {
		if err := i.ctx.AddImportDir(path, dir); err != nil {
			return nil, err
		}
		pkg := i.ctx.pkgs[path]
		if err := pkg.Load(); err != nil {
			return nil, err
		}
		return pkg.Package, nil
	}
	return nil, ErrNotFoundPackage
}
