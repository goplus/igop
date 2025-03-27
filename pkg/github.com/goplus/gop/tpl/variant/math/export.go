// export by github.com/goplus/igop/cmd/qexp

package math

import (
	q "github.com/goplus/gop/tpl/variant/math"

	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "math",
		Path: "github.com/goplus/gop/tpl/variant/math",
		Deps: map[string]string{
			"github.com/goplus/gop/tpl/token":   "token",
			"github.com/goplus/gop/tpl/variant": "variant",
			"math":                              "math",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"Max": reflect.ValueOf(q.Max),
			"Min": reflect.ValueOf(q.Min),
		},
		TypedConsts:   map[string]igop.TypedConst{},
		UntypedConsts: map[string]igop.UntypedConst{},
	})
}
