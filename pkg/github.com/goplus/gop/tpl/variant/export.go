// export by github.com/goplus/igop/cmd/qexp

package variant

import (
	q "github.com/goplus/gop/tpl/variant"

	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "variant",
		Path: "github.com/goplus/gop/tpl/variant",
		Deps: map[string]string{
			"github.com/goplus/gop/tpl":       "tpl",
			"github.com/goplus/gop/tpl/token": "token",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{
			"Module": reflect.TypeOf((*q.Module)(nil)).Elem(),
		},
		AliasTypes: map[string]reflect.Type{
			"DelayValue": reflect.TypeOf((*q.DelayValue)(nil)).Elem(),
		},
		Vars: map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"Call":         reflect.ValueOf(q.Call),
			"CallObject":   reflect.ValueOf(q.CallObject),
			"Compare":      reflect.ValueOf(q.Compare),
			"Eval":         reflect.ValueOf(q.Eval),
			"Float":        reflect.ValueOf(q.Float),
			"InitUniverse": reflect.ValueOf(q.InitUniverse),
			"Int":          reflect.ValueOf(q.Int),
			"List":         reflect.ValueOf(q.List),
			"LogicOp":      reflect.ValueOf(q.LogicOp),
			"MathOp":       reflect.ValueOf(q.MathOp),
			"NewModule":    reflect.ValueOf(q.NewModule),
			"RangeOp":      reflect.ValueOf(q.RangeOp),
			"UnaryOp":      reflect.ValueOf(q.UnaryOp),
		},
		TypedConsts:   map[string]igop.TypedConst{},
		UntypedConsts: map[string]igop.UntypedConst{},
	})
}
