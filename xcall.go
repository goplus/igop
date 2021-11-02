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
