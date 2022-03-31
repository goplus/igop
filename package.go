package gossa

import (
	"go/constant"
	"reflect"
)

var (
	registerPkgs = make(map[string]*Package)
)

// lookup register pkgs
func LookupPackage(name string) (pkg *Package, ok bool) {
	pkg, ok = registerPkgs[name]
	return
}

// register pkg
func RegisterPackage(pkg *Package) {
	if p, ok := registerPkgs[pkg.Path]; ok {
		for k, v := range pkg.Interfaces {
			p.Interfaces[k] = v
		}
		for k, v := range pkg.NamedTypes {
			p.NamedTypes[k] = v
		}
		for k, v := range pkg.Vars {
			p.Vars[k] = v
		}
		for k, v := range pkg.Funcs {
			p.Funcs[k] = v
		}
		for k, v := range pkg.UntypedConsts {
			p.UntypedConsts[k] = v
		}
		return
	}
	registerPkgs[pkg.Path] = pkg
	//	externPackages[pkg.Path] = true
}

type TypedConst struct {
	Typ   reflect.Type
	Value constant.Value
}

type UntypedConst struct {
	Typ   string
	Value constant.Value
}

type NamedType struct {
	Typ        reflect.Type
	Methods    string
	PtrMethods string
}

type Package struct {
	Name          string
	Path          string
	Interfaces    map[string]reflect.Type
	NamedTypes    map[string]NamedType
	AliasTypes    map[string]reflect.Type
	Vars          map[string]reflect.Value
	Funcs         map[string]reflect.Value
	TypedConsts   map[string]TypedConst
	UntypedConsts map[string]UntypedConst
	Deps          map[string]string
	methods       map[string]reflect.Value // methods cached
}

var (
	externValues = make(map[string]reflect.Value)
)

// register external function for no function body
func RegisterExternal(key string, i interface{}) {
	externValues[key] = reflect.ValueOf(i)
}
