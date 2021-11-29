package gossa

import (
	"go/types"
)

type Importer struct {
	loader Loader
	pkgs   map[string]*types.Package
	impl   types.Importer
}

func NewImporter(loader Loader, importer types.Importer) *Importer {
	return &Importer{
		loader: loader,
		pkgs:   make(map[string]*types.Package),
		impl:   importer,
	}
}

func (i *Importer) Import(path string) (*types.Package, error) {
	if path == "unsafe" {
		return types.Unsafe, nil
	}
	if pkg, ok := i.pkgs[path]; ok {
		return pkg, nil
	}
	if pkg, err := i.loader.Import(path); err == nil {
		i.pkgs[path] = pkg
		return pkg, nil
	}
	if i.impl == nil {
		return nil, ErrNotFoundImporter
	}
	pkg, err := i.impl.Import(path)
	if err == nil {
		i.pkgs[path] = pkg
	}
	return pkg, err
}
