// export by github.com/goplus/igop/cmd/qexp

//go:build go1.21
// +build go1.21

package base32

import (
	q "encoding/base32"

	"go/constant"
	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "base32",
		Path: "encoding/base32",
		Deps: map[string]string{
			"io":      "io",
			"strconv": "strconv",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{
			"CorruptInputError": reflect.TypeOf((*q.CorruptInputError)(nil)).Elem(),
			"Encoding":          reflect.TypeOf((*q.Encoding)(nil)).Elem(),
		},
		AliasTypes: map[string]reflect.Type{},
		Vars: map[string]reflect.Value{
			"HexEncoding": reflect.ValueOf(&q.HexEncoding),
			"StdEncoding": reflect.ValueOf(&q.StdEncoding),
		},
		Funcs: map[string]reflect.Value{
			"NewDecoder":  reflect.ValueOf(q.NewDecoder),
			"NewEncoder":  reflect.ValueOf(q.NewEncoder),
			"NewEncoding": reflect.ValueOf(q.NewEncoding),
		},
		TypedConsts: map[string]igop.TypedConst{
			"NoPadding":  {reflect.TypeOf(q.NoPadding), constant.MakeInt64(int64(q.NoPadding))},
			"StdPadding": {reflect.TypeOf(q.StdPadding), constant.MakeInt64(int64(q.StdPadding))},
		},
		UntypedConsts: map[string]igop.UntypedConst{},
	})
}
