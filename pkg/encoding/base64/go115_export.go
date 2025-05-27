// export by github.com/goplus/ixgo/cmd/qexp

//+build go1.15,!go1.16

package base64

import (
	q "encoding/base64"

	"go/constant"
	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage(&ixgo.Package{
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
		TypedConsts: map[string]ixgo.TypedConst{
			"NoPadding":  {reflect.TypeOf(q.NoPadding), constant.MakeInt64(int64(q.NoPadding))},
			"StdPadding": {reflect.TypeOf(q.StdPadding), constant.MakeInt64(int64(q.StdPadding))},
		},
		UntypedConsts: map[string]ixgo.UntypedConst{},
	})
}
