// export by github.com/goplus/gossa/cmd/qexp

//+build go1.15,!go1.16

package subtle

import (
	q "crypto/subtle"

	"reflect"

	"github.com/goplus/gossa"
)

func init() {
	gossa.RegisterPackage(&gossa.Package{
		Name:       "subtle",
		Path:       "crypto/subtle",
		Deps:       map[string]string{},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]gossa.NamedType{},
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
		TypedConsts:   map[string]gossa.TypedConst{},
		UntypedConsts: map[string]gossa.UntypedConst{},
	})
}
