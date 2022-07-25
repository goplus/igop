// export by github.com/goplus/igop/cmd/qexp

package builtin

import (
	q "github.com/goplus/gop/builtin"

	"go/constant"
	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name:       "builtin",
		Path:       "github.com/goplus/gop/builtin",
		Deps:       map[string]string{},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{
			"IntRange": reflect.TypeOf((*q.IntRange)(nil)).Elem(),
		},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"NewRange__0": reflect.ValueOf(q.NewRange__0),
		},
		TypedConsts: map[string]igop.TypedConst{},
		UntypedConsts: map[string]igop.UntypedConst{
			"GopPackage": {"untyped bool", constant.MakeBool(bool(q.GopPackage))},
		},
	})
}
