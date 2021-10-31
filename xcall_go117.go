//go:build go1.17
// +build go1.17

package gossa

import "reflect"

func AllMethod(typ reflect.Type) []reflect.Method {
	n := typ.NumMethod()
	if n == 0 {
		return nil
	}
	ms := make([]reflect.Method, n, n)
	for i := 0; i < n; i++ {
		ms[i] = typ.Method(i)
	}
	return ms
}
