// export by github.com/goplus/igop/cmd/qexp

package matcher

import (
	q "github.com/goplus/gop/tpl/matcher"

	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "matcher",
		Path: "github.com/goplus/gop/tpl/matcher",
		Deps: map[string]string{
			"errors":                          "errors",
			"fmt":                             "fmt",
			"github.com/goplus/gop/tpl/token": "token",
			"github.com/goplus/gop/tpl/types": "types",
		},
		Interfaces: map[string]reflect.Type{
			"Matcher": reflect.TypeOf((*q.Matcher)(nil)).Elem(),
		},
		NamedTypes: map[string]reflect.Type{
			"Context": reflect.TypeOf((*q.Context)(nil)).Elem(),
			"Error":   reflect.TypeOf((*q.Error)(nil)).Elem(),
			"Var":     reflect.TypeOf((*q.Var)(nil)).Elem(),
		},
		AliasTypes: map[string]reflect.Type{},
		Vars: map[string]reflect.Value{
			"ErrVarAssigned": reflect.ValueOf(&q.ErrVarAssigned),
		},
		Funcs: map[string]reflect.Value{
			"Choice":   reflect.ValueOf(q.Choice),
			"List":     reflect.ValueOf(q.List),
			"Literal":  reflect.ValueOf(q.Literal),
			"NewVar":   reflect.ValueOf(q.NewVar),
			"Repeat0":  reflect.ValueOf(q.Repeat0),
			"Repeat01": reflect.ValueOf(q.Repeat01),
			"Repeat1":  reflect.ValueOf(q.Repeat1),
			"Sequence": reflect.ValueOf(q.Sequence),
			"Token":    reflect.ValueOf(q.Token),
		},
		TypedConsts:   map[string]igop.TypedConst{},
		UntypedConsts: map[string]igop.UntypedConst{},
	})
}
