package gossa

import (
	"go/importer"
	"go/types"
)

type Importer struct {
	loader *TypesLoader
	pkgs   map[string]*types.Package
	impl   types.Importer
}

func NewImporter(inst *TypesLoader) types.Importer {
	return &Importer{
		loader: inst,
		pkgs:   make(map[string]*types.Package),
		impl:   importer.Default(),
	}
}

func (i *Importer) Import(path string) (*types.Package, error) {
	if pkg, ok := i.pkgs[path]; ok {
		return pkg, nil
	}
	if p, ok := registerPkgs[path]; ok {
		pkg, err := i.loader.InstallPackage(p)
		if err == nil {
			i.pkgs[path] = pkg
		}
		return pkg, nil
	}
	pkg, err := i.impl.Import(path)
	if err == nil {
		i.pkgs[path] = pkg
	}
	return pkg, err
}
