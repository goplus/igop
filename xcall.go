package gossa

import (
	"errors"
	"reflect"

	"github.com/goplus/reflectx"
)

func FieldAddr(v interface{}, index int) (interface{}, error) {
	x := reflect.ValueOf(v).Elem()
	if !x.IsValid() {
		return nil, errors.New("invalid memory address or nil pointer dereference")
	}
	return reflectx.FieldX(x, index).Addr().Interface(), nil
}

func Field(v interface{}, index int) (interface{}, error) {
	x := reflect.ValueOf(v)
	for x.Kind() == reflect.Ptr {
		x = x.Elem()
	}
	if !x.IsValid() {
		return nil, errors.New("invalid memory address or nil pointer dereference")
	}
	return reflectx.FieldX(x, index).Interface(), nil
}

func AllMethod(typ reflect.Type) []reflect.Method {
	n := reflectx.NumMethodX(typ)
	if n == 0 {
		return nil
	}
	ms := make([]reflect.Method, n, n)
	for i := 0; i < n; i++ {
		ms[i] = reflectx.MethodX(typ, i)
	}
	return ms
}
