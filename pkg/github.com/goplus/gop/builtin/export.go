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
		Name: "builtin",
		Path: "github.com/goplus/gop/builtin",
		Deps: map[string]string{
			"math/big": "big",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]gossa.NamedType{
			"Gop_bigfloat":         {reflect.TypeOf((*q.Gop_bigfloat)(nil)).Elem(), "", ""},
			"Gop_bigint":           {reflect.TypeOf((*q.Gop_bigint)(nil)).Elem(), "Gop_Add,Gop_AddAssign,Gop_And,Gop_AndAssign,Gop_AndNot,Gop_AndNotAssign,Gop_Assign,Gop_EQ,Gop_GE,Gop_GT,Gop_LE,Gop_LT,Gop_Lsh,Gop_LshAssign,Gop_Mul,Gop_MulAssign,Gop_NE,Gop_Neg,Gop_Not,Gop_Or,Gop_OrAssign,Gop_Pos,Gop_Quo,Gop_QuoAssign,Gop_Rem,Gop_RemAssign,Gop_Rsh,Gop_RshAssign,Gop_Sub,Gop_SubAssign,Gop_Xor,Gop_XorAssign,IsNil", ""},
			"Gop_bigrat":           {reflect.TypeOf((*q.Gop_bigrat)(nil)).Elem(), "Gop_Add,Gop_AddAssign,Gop_Assign,Gop_EQ,Gop_GE,Gop_GT,Gop_Inv,Gop_LE,Gop_LT,Gop_Mul,Gop_MulAssign,Gop_NE,Gop_Neg,Gop_Pos,Gop_Quo,Gop_QuoAssign,Gop_Sub,Gop_SubAssign,IsNil", ""},
			"Gop_untyped_bigfloat": {reflect.TypeOf((*q.Gop_untyped_bigfloat)(nil)).Elem(), "", ""},
			"Gop_untyped_bigint":   {reflect.TypeOf((*q.Gop_untyped_bigint)(nil)).Elem(), "", ""},
			"Gop_untyped_bigrat":   {reflect.TypeOf((*q.Gop_untyped_bigrat)(nil)).Elem(), "", ""},
			"IntRange":             {reflect.TypeOf((*q.IntRange)(nil)).Elem(), "", "Gop_Enum"},
		},
		AliasTypes: map[string]reflect.Type{
			"Gop_ninteger":                 reflect.TypeOf((*uint)(nil)).Elem(),
			"Gop_untyped_bigfloat_Default": reflect.TypeOf((*q.Gop_untyped_bigfloat_Default)(nil)).Elem(),
			"Gop_untyped_bigint_Default":   reflect.TypeOf((*q.Gop_untyped_bigint_Default)(nil)).Elem(),
			"Gop_untyped_bigrat_Default":   reflect.TypeOf((*q.Gop_untyped_bigrat_Default)(nil)).Elem(),
		},
		Vars: map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"Gop_bigint_Cast__0":         reflect.ValueOf(q.Gop_bigint_Cast__0),
			"Gop_bigint_Cast__1":         reflect.ValueOf(q.Gop_bigint_Cast__1),
			"Gop_bigint_Cast__2":         reflect.ValueOf(q.Gop_bigint_Cast__2),
			"Gop_bigint_Cast__3":         reflect.ValueOf(q.Gop_bigint_Cast__3),
			"Gop_bigint_Cast__4":         reflect.ValueOf(q.Gop_bigint_Cast__4),
			"Gop_bigint_Cast__5":         reflect.ValueOf(q.Gop_bigint_Cast__5),
			"Gop_bigint_Cast__6":         reflect.ValueOf(q.Gop_bigint_Cast__6),
			"Gop_bigint_Init__0":         reflect.ValueOf(q.Gop_bigint_Init__0),
			"Gop_bigint_Init__1":         reflect.ValueOf(q.Gop_bigint_Init__1),
			"Gop_bigint_Init__2":         reflect.ValueOf(q.Gop_bigint_Init__2),
			"Gop_bigrat_Cast__0":         reflect.ValueOf(q.Gop_bigrat_Cast__0),
			"Gop_bigrat_Cast__1":         reflect.ValueOf(q.Gop_bigrat_Cast__1),
			"Gop_bigrat_Cast__2":         reflect.ValueOf(q.Gop_bigrat_Cast__2),
			"Gop_bigrat_Cast__3":         reflect.ValueOf(q.Gop_bigrat_Cast__3),
			"Gop_bigrat_Cast__4":         reflect.ValueOf(q.Gop_bigrat_Cast__4),
			"Gop_bigrat_Init__0":         reflect.ValueOf(q.Gop_bigrat_Init__0),
			"Gop_bigrat_Init__1":         reflect.ValueOf(q.Gop_bigrat_Init__1),
			"Gop_bigrat_Init__2":         reflect.ValueOf(q.Gop_bigrat_Init__2),
			"Gop_istmp":                  reflect.ValueOf(q.Gop_istmp),
			"Gop_untyped_bigint_Init__0": reflect.ValueOf(q.Gop_untyped_bigint_Init__0),
			"Gop_untyped_bigrat_Init__0": reflect.ValueOf(q.Gop_untyped_bigrat_Init__0),
			"Gop_untyped_bigrat_Init__1": reflect.ValueOf(q.Gop_untyped_bigrat_Init__1),
			"NewRange__0":                reflect.ValueOf(q.NewRange__0),
		},
		TypedConsts: map[string]gossa.TypedConst{},
		UntypedConsts: map[string]gossa.UntypedConst{
			"GopPackage": {"untyped bool", constant.MakeBool(bool(q.GopPackage))},
		},
	})
}
