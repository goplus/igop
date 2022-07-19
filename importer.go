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
	if i.ctx.Mode&EnableDumpImports != 0 {
		fmt.Printf("import %q\n", path)
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
		if pkg.Info == nil {
			info, err := i.ctx.checkTypesInfo(pkg.Package, pkg.Files)
			if err != nil {
				return nil, err
			}
			pkg.Info = info
		}
		i.pkgs[path] = pkg.Package
		return pkg.Package, nil
	}
	if i.ctx.External != nil {
		if dir, found := i.ctx.External.Lookup(path); found {
			if i.ctx.AddImportDir(path, dir) == nil {
				pkg := i.ctx.pkgs[path]
				info, err := i.ctx.checkTypesInfo(pkg.Package, pkg.Files)
				if err != nil {
					return nil, err
				}
				pkg.Info = info
				return pkg.Package, nil
			}
		}
	}
	pkg, err := i.defaultImpl.Import(path)
	if err == nil {
		i.pkgs[path] = pkg
	}
	return pkg, err
}
