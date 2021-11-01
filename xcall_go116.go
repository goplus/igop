//go:build !go1.17
// +build !go1.17

package gossa

import (
	"reflect"

	"github.com/goplus/reflectx"
)

func AllMethod(typ reflect.Type) []reflect.Method {
	n := reflectx.AllNumMethod(typ)
	if n == 0 {
		return nil
	}
	ms := make([]reflect.Method, n, n)
	for i := 0; i < n; i++ {
		ms[i] = reflectx.AllMethod(typ, i)
	}
	return ms
}
