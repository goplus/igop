package igop

import (
	"errors"
	"reflect"

	"github.com/goplus/reflectx"
)

func fieldAddrX(v interface{}, index int) (interface{}, error) {
	x := reflect.ValueOf(v).Elem()
	if !x.IsValid() {
		return nil, errors.New("invalid memory address or nil pointer dereference")
	}
	return reflectx.FieldX(x, index).Addr().Interface(), nil
}

func fieldX(v interface{}, index int) (interface{}, error) {
	x := reflect.ValueOf(v)
	for x.Kind() == reflect.Ptr {
		x = x.Elem()
	}
	if !x.IsValid() {
		return nil, errors.New("invalid memory address or nil pointer dereference")
	}
	return reflectx.FieldX(x, index).Interface(), nil
}

func allMethodX(typ reflect.Type) []reflect.Method {
	n := reflectx.NumMethodX(typ)
	if n == 0 {
		return nil
	}
	ms := make([]reflect.Method, n)
	for i := 0; i < n; i++ {
		ms[i] = reflectx.MethodX(typ, i)
	}
	return ms
}
