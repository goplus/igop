// export by github.com/goplus/igop/cmd/qexp

package types

import (
	q "github.com/goplus/gop/tpl/types"

	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "types",
		Path: "github.com/goplus/gop/tpl/types",
		Deps: map[string]string{
			"github.com/goplus/gop/tpl/token": "token",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{
			"Token": reflect.TypeOf((*q.Token)(nil)).Elem(),
		},
		AliasTypes:    map[string]reflect.Type{},
		Vars:          map[string]reflect.Value{},
		Funcs:         map[string]reflect.Value{},
		TypedConsts:   map[string]igop.TypedConst{},
		UntypedConsts: map[string]igop.UntypedConst{},
	})
}
