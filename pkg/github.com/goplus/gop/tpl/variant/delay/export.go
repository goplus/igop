// export by github.com/goplus/igop/cmd/qexp

package delay

import (
	q "github.com/goplus/gop/tpl/variant/delay"

	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "delay",
		Path: "github.com/goplus/gop/tpl/variant/delay",
		Deps: map[string]string{
			"github.com/goplus/gop/tpl/token":   "token",
			"github.com/goplus/gop/tpl/variant": "variant",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{},
		AliasTypes: map[string]reflect.Type{
			"Value": reflect.TypeOf((*q.Value)(nil)).Elem(),
		},
		Vars: map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"Call":         reflect.ValueOf(q.Call),
			"ChgValue":     reflect.ValueOf(q.ChgValue),
			"Compare":      reflect.ValueOf(q.Compare),
			"Eval":         reflect.ValueOf(q.Eval),
			"EvalOp":       reflect.ValueOf(q.EvalOp),
			"IfElse":       reflect.ValueOf(q.IfElse),
			"InitUniverse": reflect.ValueOf(q.InitUniverse),
			"ListOp":       reflect.ValueOf(q.ListOp),
			"LogicOp":      reflect.ValueOf(q.LogicOp),
			"MathOp":       reflect.ValueOf(q.MathOp),
			"RangeOp":      reflect.ValueOf(q.RangeOp),
			"RepeatUntil":  reflect.ValueOf(q.RepeatUntil),
			"SetValue":     reflect.ValueOf(q.SetValue),
			"StmtList":     reflect.ValueOf(q.StmtList),
			"UnaryOp":      reflect.ValueOf(q.UnaryOp),
			"ValueOf":      reflect.ValueOf(q.ValueOf),
			"While":        reflect.ValueOf(q.While),
		},
		TypedConsts:   map[string]igop.TypedConst{},
		UntypedConsts: map[string]igop.UntypedConst{},
	})
}
