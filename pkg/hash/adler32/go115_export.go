// export by github.com/goplus/igop/cmd/qexp

//+build go1.15,!go1.16

package adler32

import (
	q "hash/adler32"

	"go/constant"
	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "adler32",
		Path: "hash/adler32",
		Deps: map[string]string{
			"errors": "errors",
			"hash":   "hash",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"Checksum": reflect.ValueOf(q.Checksum),
			"New":      reflect.ValueOf(q.New),
		},
		TypedConsts: map[string]igop.TypedConst{},
		UntypedConsts: map[string]igop.UntypedConst{
			"Size": {"untyped int", constant.MakeInt64(int64(q.Size))},
		},
	})
}
