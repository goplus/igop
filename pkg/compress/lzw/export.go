// export by github.com/goplus/gossa/cmd/qexp

package lzw

import (
	q "compress/lzw"

	"go/constant"
	"reflect"

	"github.com/goplus/gossa"
)

func init() {
	gossa.RegisterPackage(&gossa.Package{
		Name: "lzw",
		Path: "compress/lzw",
		Deps: map[string]string{
			"bufio":  "bufio",
			"errors": "errors",
			"fmt":    "fmt",
			"io":     "io",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]gossa.NamedType{
			"Order": {reflect.TypeOf((*q.Order)(nil)).Elem(), "", ""},
		},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"NewReader": reflect.ValueOf(q.NewReader),
			"NewWriter": reflect.ValueOf(q.NewWriter),
		},
		TypedConsts: map[string]gossa.TypedConst{
			"LSB": {reflect.TypeOf(q.LSB), constant.MakeInt64(0)},
			"MSB": {reflect.TypeOf(q.MSB), constant.MakeInt64(1)},
		},
		UntypedConsts: map[string]gossa.UntypedConst{},
	})
}
