// export by github.com/goplus/igop/cmd/qexp

//go:build go1.20 && !go1.20
// +build go1.20,!go1.20

package base64

import (
	q "encoding/base64"

	"go/constant"
	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "base64",
		Path: "encoding/base64",
		Deps: map[string]string{
			"encoding/binary": "binary",
			"io":              "io",
			"strconv":         "strconv",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{
			"CorruptInputError": reflect.TypeOf((*q.CorruptInputError)(nil)).Elem(),
			"Encoding":          reflect.TypeOf((*q.Encoding)(nil)).Elem(),
		},
		AliasTypes: map[string]reflect.Type{},
		Vars: map[string]reflect.Value{
			"RawStdEncoding": reflect.ValueOf(&q.RawStdEncoding),
			"RawURLEncoding": reflect.ValueOf(&q.RawURLEncoding),
			"StdEncoding":    reflect.ValueOf(&q.StdEncoding),
			"URLEncoding":    reflect.ValueOf(&q.URLEncoding),
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
