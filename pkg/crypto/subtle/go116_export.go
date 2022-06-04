// export by github.com/goplus/igop/cmd/qexp

//+build go1.16,!go1.17

package subtle

import (
	q "crypto/subtle"

	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name:       "subtle",
		Path:       "crypto/subtle",
		Deps:       map[string]string{},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]igop.NamedType{},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"ConstantTimeByteEq":   reflect.ValueOf(q.ConstantTimeByteEq),
			"ConstantTimeCompare":  reflect.ValueOf(q.ConstantTimeCompare),
			"ConstantTimeCopy":     reflect.ValueOf(q.ConstantTimeCopy),
			"ConstantTimeEq":       reflect.ValueOf(q.ConstantTimeEq),
			"ConstantTimeLessOrEq": reflect.ValueOf(q.ConstantTimeLessOrEq),
			"ConstantTimeSelect":   reflect.ValueOf(q.ConstantTimeSelect),
		},
		TypedConsts:   map[string]igop.TypedConst{},
		UntypedConsts: map[string]igop.UntypedConst{},
	})
}
