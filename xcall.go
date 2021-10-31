//go:build !go1.17
// +build !go1.17

package gossa

import (
	"reflect"

	xcall "github.com/goplus/gossa/internal/reflect"
)

func AllMethod(typ reflect.Type) []reflect.Method {
	return xcall.AllMethod(typ)
}
