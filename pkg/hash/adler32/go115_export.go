// export by github.com/goplus/gossa/cmd/qexp

//+build go1.15,!go1.16

package adler32

import (
	q "hash/adler32"

	"go/constant"
	"reflect"

	"github.com/goplus/gossa"
)

func init() {
	gossa.RegisterPackage(&gossa.Package{
		Name: "adler32",
		Path: "hash/adler32",
		Deps: map[string]string{
			"errors": "errors",
			"hash":   "hash",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]gossa.NamedType{},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"Checksum": reflect.ValueOf(q.Checksum),
			"New":      reflect.ValueOf(q.New),
		},
		TypedConsts: map[string]gossa.TypedConst{},
		UntypedConsts: map[string]gossa.UntypedConst{
			"Size": {"untyped int", constant.MakeInt64(int64(q.Size))},
		},
	})
}
