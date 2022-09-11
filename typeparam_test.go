//go:build go1.18
// +build go1.18

package igop_test

import (
	"testing"

	"github.com/goplus/igop"
	_ "github.com/goplus/igop/pkg/path/filepath"
	_ "github.com/goplus/igop/pkg/reflect"
)

func TestTypeParamNamed(t *testing.T) {
	src := `package main

import (
	"path/filepath"
	"reflect"
)

type N int
type _Nil[a any, b any] struct{}

func main() {
	var v1 _Nil[int, N]
	var v2 _Nil[filepath.WalkFunc, _Nil[int, N]]
	if s := reflect.TypeOf(v1).String(); s != "main._Nil[int,main.N]" {
		panic(s)
	}
	if s := reflect.TypeOf(v2).String(); s != "main._Nil[path/filepath.WalkFunc,main._Nil[int,main.N]]" {
		panic(s)
	}
}
`
	_, err := igop.RunFile("main.go", src, nil, 0)
	if err != nil {
		t.Fatal(err)
	}
}
