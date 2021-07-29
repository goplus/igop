// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package interp

import (
	"reflect"
)

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
	}
}

func init() {
	RegisterExternal("runtime.init", func() {})
	RegisterExternal("os.Exit", func(code int) {
		panic(exitPanic(code))
	})
}
