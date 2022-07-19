package igop

import (
	"fmt"
	"go/types"
)

type Importer struct {
	loader    Loader
	pkgs      map[string]*types.Package
	importing map[string]bool
	impl      types.Importer
}

func NewImporter(loader Loader, importer types.Importer) *Importer {
	return &Importer{
		loader:    loader,
		pkgs:      make(map[string]*types.Package),
		importing: make(map[string]bool),
		impl:      importer,
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
