// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.24
// +build go1.24

package subtle

import (
	q "crypto/subtle"

	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage(&ixgo.Package{
		Name: "subtle",
		Path: "crypto/subtle",
		Deps: map[string]string{
			"crypto/internal/fips140/subtle": "subtle",
			"internal/runtime/sys":           "sys",
			"runtime":                        "runtime",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"ConstantTimeByteEq":        reflect.ValueOf(q.ConstantTimeByteEq),
			"ConstantTimeCompare":       reflect.ValueOf(q.ConstantTimeCompare),
			"ConstantTimeCopy":          reflect.ValueOf(q.ConstantTimeCopy),
			"ConstantTimeEq":            reflect.ValueOf(q.ConstantTimeEq),
			"ConstantTimeLessOrEq":      reflect.ValueOf(q.ConstantTimeLessOrEq),
			"ConstantTimeSelect":        reflect.ValueOf(q.ConstantTimeSelect),
			"WithDataIndependentTiming": reflect.ValueOf(q.WithDataIndependentTiming),
			"XORBytes":                  reflect.ValueOf(q.XORBytes),
		},
		TypedConsts:   map[string]ixgo.TypedConst{},
		UntypedConsts: map[string]ixgo.UntypedConst{},
	})
}
