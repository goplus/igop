// export by github.com/goplus/igop/cmd/qexp

package buitin

import (
	q "github.com/goplus/gop/tpl/variant/builtin"

	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "buitin",
		Path: "github.com/goplus/gop/tpl/variant/builtin",
		Deps: map[string]string{
			"github.com/goplus/gop/tpl/variant": "variant",
			"reflect":                           "reflect",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"CastInt": reflect.ValueOf(q.CastInt),
			"Type":    reflect.ValueOf(q.Type),
		},
		TypedConsts:   map[string]igop.TypedConst{},
		UntypedConsts: map[string]igop.UntypedConst{},
	})
}
