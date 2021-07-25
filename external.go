// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package interp

import (
	"reflect"
)

var externals = make(map[string]reflect.Value)

// register external interface
func RegisterExternal(key string, fn interface{}) {
	externals[key] = reflect.ValueOf(fn)
}

func init() {
	RegisterExternal("runtime.init", func() {})
	RegisterExternal("os.Exit", func(code int) {
		panic(exitPanic(code))
	})
}
