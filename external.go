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
func RegisterExternal(key string, fn interface{}) {
	externValues[key] = reflect.ValueOf(fn)
}

// register external interface
func RegisterType(key string, t reflect.Type) {
	externTypes[key] = t
}

func init() {
	RegisterExternal("runtime.init", func() {})
	RegisterExternal("os.Exit", func(code int) {
		panic(exitPanic(code))
	})
	RegisterType("error", reflect.TypeOf((*error)(nil)).Elem())
}
