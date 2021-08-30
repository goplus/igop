// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package interp

import (
	"go/token"
	"reflect"
)

var externPackages = make(map[string]bool)
var externValues = make(map[string]reflect.Value)
var externTypes = make(map[string]reflect.Type)

// register external interface
func RegisterExternal(key string, i interface{}) {
	externValues[key] = reflect.ValueOf(i)
}

// register external interface map
func RegisterExternals(m map[string]interface{}) {
	for k, v := range m {
		externValues[k] = reflect.ValueOf(v)
	}
}

// register external type
func RegisterType(key string, typ reflect.Type) {
	externTypes[key] = typ
	if typ.Kind() == reflect.Interface {
		registerInterface(typ)
	}
}

func registerType(typ reflect.Type) {
	if pkgPath, pkgName := typ.PkgPath(), typ.Name(); pkgPath != "" && pkgName != "" {
		externTypes[pkgPath+"."+pkgName] = typ
		return
	}
	switch typ.Kind() {
	case reflect.Interface:
		n := typ.NumMethod()
		for i := 0; i < n; i++ {
			registerType(typ.Method(i).Type)
		}
	case reflect.Func:
		numIn := typ.NumIn()
		for i := 0; i < numIn; i++ {
			registerType(typ.In(i))
		}
		numOut := typ.NumOut()
		for i := 0; i < numOut; i++ {
			registerType(typ.Out(i))
		}
	case reflect.Ptr:
		registerType(typ.Elem())
	case reflect.Array:
		registerType(typ.Elem())
	case reflect.Slice:
		registerType(typ.Elem())
	case reflect.Chan:
		registerType(typ.Elem())
	case reflect.Map:
		registerType(typ.Key())
		registerType(typ.Elem())
	}
}

func registerInterface(typ reflect.Type) {
	n := typ.NumMethod()
	for i := 0; i < n; i++ {
		m := typ.Method(i)
		if token.IsExported(m.Name) {
			continue
		}
		registerType(m.Type)
	}
}

// register external type list
func RegisterTypeOf(ptrs ...interface{}) {
	for _, ptr := range ptrs {
		typ := reflect.TypeOf(ptr).Elem()
		var key string
		if pkgPath := typ.PkgPath(); pkgPath != "" {
			key = pkgPath + "."
		}
		key += typ.Name()
		externTypes[key] = typ
		if typ.Kind() == reflect.Interface {
			registerInterface(typ)
		}
	}
}

// register extern package
func RegisterPackage(pkg string, ifacemap map[string]interface{}, typelist []interface{}) {
	externPackages[pkg] = true
	RegisterExternal(pkg+".init", func() {})
	RegisterExternals(ifacemap)
	RegisterTypeOf(typelist...)
}

func init() {
	RegisterExternal("runtime.init", func() {})
	RegisterPackage("unsafe", nil, nil)
}
