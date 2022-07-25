// export by github.com/goplus/igop/cmd/qexp

//+build go1.16,!go1.17

package dsa

import (
	q "crypto/dsa"

	"go/constant"
	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "dsa",
		Path: "crypto/dsa",
		Deps: map[string]string{
			"crypto/internal/randutil": "randutil",
			"errors":                   "errors",
			"io":                       "io",
			"math/big":                 "big",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{
			"ParameterSizes": reflect.TypeOf((*q.ParameterSizes)(nil)).Elem(),
			"Parameters":     reflect.TypeOf((*q.Parameters)(nil)).Elem(),
			"PrivateKey":     reflect.TypeOf((*q.PrivateKey)(nil)).Elem(),
			"PublicKey":      reflect.TypeOf((*q.PublicKey)(nil)).Elem(),
		},
		AliasTypes: map[string]reflect.Type{},
		Vars: map[string]reflect.Value{
			"ErrInvalidPublicKey": reflect.ValueOf(&q.ErrInvalidPublicKey),
		},
		Funcs: map[string]reflect.Value{
			"GenerateKey":        reflect.ValueOf(q.GenerateKey),
			"GenerateParameters": reflect.ValueOf(q.GenerateParameters),
			"Sign":               reflect.ValueOf(q.Sign),
			"Verify":             reflect.ValueOf(q.Verify),
		},
		TypedConsts: map[string]igop.TypedConst{
			"L1024N160": {reflect.TypeOf(q.L1024N160), constant.MakeInt64(int64(q.L1024N160))},
			"L2048N224": {reflect.TypeOf(q.L2048N224), constant.MakeInt64(int64(q.L2048N224))},
			"L2048N256": {reflect.TypeOf(q.L2048N256), constant.MakeInt64(int64(q.L2048N256))},
			"L3072N256": {reflect.TypeOf(q.L3072N256), constant.MakeInt64(int64(q.L3072N256))},
		},
		UntypedConsts: map[string]igop.UntypedConst{},
	})
}
