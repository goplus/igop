// export by github.com/goplus/gossa/cmd/qexp

//+build go1.15,!go1.16

package base32

import (
	q "encoding/base32"

	"go/constant"
	"reflect"

	"github.com/goplus/gossa"
)

func init() {
	gossa.RegisterPackage(&gossa.Package{
		Name: "base32",
		Path: "encoding/base32",
		Deps: map[string]string{
			"io":      "io",
			"strconv": "strconv",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]gossa.NamedType{
			"CorruptInputError": {reflect.TypeOf((*q.CorruptInputError)(nil)).Elem(), "Error", ""},
			"Encoding":          {reflect.TypeOf((*q.Encoding)(nil)).Elem(), "WithPadding", "Decode,DecodeString,DecodedLen,Encode,EncodeToString,EncodedLen,decode"},
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
		TypedConsts: map[string]gossa.TypedConst{
			"NoPadding":  {reflect.TypeOf(q.NoPadding), constant.MakeInt64(int64(q.NoPadding))},
			"StdPadding": {reflect.TypeOf(q.StdPadding), constant.MakeInt64(int64(q.StdPadding))},
		},
		UntypedConsts: map[string]gossa.UntypedConst{},
	})
}
