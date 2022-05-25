package gossa

import "errors"

var (
	ErrNoPackage        = errors.New("no package")
	ErrPackage          = errors.New("package contain errors")
	ErrNotFoundMain     = errors.New("not found main package")
	ErrTestFailed       = errors.New("test failed")
	ErrNotFoundPackage  = errors.New("not found package")
	ErrNotFoundImporter = errors.New("not found provider for types.Importer")
	ErrGoexitDeadlock   = errors.New("fatal error: no goroutines (main called runtime.Goexit) - deadlock!")
)
