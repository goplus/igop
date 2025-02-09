// export by github.com/goplus/igop/cmd/qexp

//go:build go1.21 && !go1.22
// +build go1.21,!go1.22

package cmp

import (
	_ "cmp"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name:   "cmp",
		Path:   "cmp",
		Deps:   map[string]string{},
		Source: source,
	})
}

var source = "package cmp\n\ntype Ordered interface {\n\t~int | ~int8 | ~int16 | ~int32 | ~int64 |\n\t\t~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |\n\t\t~float32 | ~float64 |\n\t\t~string\n}\n\nfunc Less[T Ordered](x, y T) bool {\n\treturn (isNaN(x) && !isNaN(y)) || x < y\n}\n\nfunc Compare[T Ordered](x, y T) int {\n\txNaN := isNaN(x)\n\tyNaN := isNaN(y)\n\tif xNaN && yNaN {\n\t\treturn 0\n\t}\n\tif xNaN || x < y {\n\t\treturn -1\n\t}\n\tif yNaN || x > y {\n\t\treturn +1\n\t}\n\treturn 0\n}\n\nfunc isNaN[T Ordered](x T) bool {\n\treturn x != x\n}\n"
