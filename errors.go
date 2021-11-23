package gossa

import "errors"

var (
	ErrNoPackage        = errors.New("no package")
	ErrPackage          = errors.New("package contain errors")
	ErrNotFoundMain     = errors.New("not found main package")
	ErrTestFailed       = errors.New("test failed")
	ErrNotFoundImporter = errors.New("not found provider for types.Importer")
)
