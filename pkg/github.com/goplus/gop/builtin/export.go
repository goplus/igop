// export by github.com/goplus/gossa/cmd/qexp

package builtin

import (
	q "github.com/goplus/gop/builtin"

	"go/constant"
	"reflect"

	"github.com/goplus/gossa"
)

func init() {
	gossa.RegisterPackage(&gossa.Package{
		Name:       "builtin",
		Path:       "github.com/goplus/gop/builtin",
		Deps:       map[string]string{},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]gossa.NamedType{
			"IntRange": {reflect.TypeOf((*q.IntRange)(nil)).Elem(), "", "Gop_Enum"},
		},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"NewRange__0": reflect.ValueOf(q.NewRange__0),
		},
		TypedConsts: map[string]gossa.TypedConst{},
		UntypedConsts: map[string]gossa.UntypedConst{
			"GopPackage": {"untyped bool", constant.MakeBool(bool(q.GopPackage))},
		},
	})
}
