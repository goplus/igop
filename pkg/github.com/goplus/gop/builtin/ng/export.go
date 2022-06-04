// export by github.com/goplus/igop/cmd/qexp

package ng

import (
	q "github.com/goplus/gop/builtin/ng"

	"go/constant"
	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "ng",
		Path: "github.com/goplus/gop/builtin/ng",
		Deps: map[string]string{
			"fmt":       "fmt",
			"log":       "log",
			"math/big":  "big",
			"math/bits": "bits",
			"strconv":   "strconv",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]igop.NamedType{
			"Bigfloat":        {reflect.TypeOf((*q.Bigfloat)(nil)).Elem(), "", ""},
			"Bigint":          {reflect.TypeOf((*q.Bigint)(nil)).Elem(), "Gop_Add,Gop_AddAssign,Gop_And,Gop_AndAssign,Gop_AndNot,Gop_AndNotAssign,Gop_Dec,Gop_Dup,Gop_EQ,Gop_GE,Gop_GT,Gop_Inc,Gop_LE,Gop_LT,Gop_Lsh,Gop_LshAssign,Gop_Mul,Gop_MulAssign,Gop_NE,Gop_Neg,Gop_Not,Gop_Or,Gop_OrAssign,Gop_Quo,Gop_QuoAssign,Gop_Rcast__0,Gop_Rcast__1,Gop_Rcast__2,Gop_Rcast__3,Gop_Rem,Gop_RemAssign,Gop_Rsh,Gop_RshAssign,Gop_Sub,Gop_SubAssign,Gop_Xor,Gop_XorAssign,IsNil", ""},
			"Bigrat":          {reflect.TypeOf((*q.Bigrat)(nil)).Elem(), "Gop_Add,Gop_AddAssign,Gop_Assign,Gop_Dup,Gop_EQ,Gop_GE,Gop_GT,Gop_Inv,Gop_LE,Gop_LT,Gop_Mul,Gop_MulAssign,Gop_NE,Gop_Neg,Gop_Quo,Gop_QuoAssign,Gop_Sub,Gop_SubAssign,IsNil", ""},
			"Int128":          {reflect.TypeOf((*q.Int128)(nil)).Elem(), "AbsU,Abs__0,Abs__1,BigInt,Cmp__0,Cmp__1,Format,Gop_Add__0,Gop_Add__1,Gop_And,Gop_AndNot,Gop_Dup,Gop_EQ__0,Gop_EQ__1,Gop_GE__0,Gop_GE__1,Gop_GT__0,Gop_GT__1,Gop_LE__0,Gop_LE__1,Gop_LT__0,Gop_LT__1,Gop_Lsh,Gop_Mul__0,Gop_Mul__1,Gop_Neg,Gop_Not,Gop_Or,Gop_Quo__0,Gop_Quo__1,Gop_Rcast__0,Gop_Rcast__1,Gop_Rcast__2,Gop_Rcast__3,Gop_Rcast__4,Gop_Rcast__5,Gop_Rem__0,Gop_Rem__1,Gop_Rsh,Gop_Sub__0,Gop_Sub__1,Gop_Xor,IsZero,QuoRem__0,QuoRem__1,Sign,String,Text,ToBigInt", "Gop_AddAssign,Gop_AndAssign,Gop_AndNotAssign,Gop_Dec,Gop_Inc,Gop_LshAssign,Gop_MulAssign,Gop_OrAssign,Gop_QuoAssign,Gop_RemAssign,Gop_RshAssign,Gop_SubAssign,Gop_XorAssign,Scan"},
			"Uint128":         {reflect.TypeOf((*q.Uint128)(nil)).Elem(), "BigInt,Bit,BitLen,Cmp__0,Cmp__1,Format,Gop_Add__0,Gop_Add__1,Gop_AndNot,Gop_And__0,Gop_And__1,Gop_Dup,Gop_EQ__0,Gop_EQ__1,Gop_GE__0,Gop_GE__1,Gop_GT__0,Gop_GT__1,Gop_LE__0,Gop_LE__1,Gop_LT__0,Gop_LT__1,Gop_Lsh,Gop_Mul__0,Gop_Mul__1,Gop_Not,Gop_Or__0,Gop_Or__1,Gop_Quo__0,Gop_Quo__1,Gop_Rcast__0,Gop_Rcast__1,Gop_Rcast__2,Gop_Rcast__3,Gop_Rcast__4,Gop_Rem__0,Gop_Rem__1,Gop_Rsh,Gop_Sub__0,Gop_Sub__1,Gop_Xor__0,Gop_Xor__1,IsZero,LeadingZeros,OnesCount,QuoRem__0,QuoRem__1,Reverse,ReverseBytes,SetBit,String,Text,ToBigInt,TrailingZeros", "Gop_AddAssign,Gop_AndAssign,Gop_AndNotAssign,Gop_Dec,Gop_Inc,Gop_LshAssign,Gop_MulAssign,Gop_OrAssign,Gop_QuoAssign,Gop_RemAssign,Gop_RshAssign,Gop_SubAssign,Gop_XorAssign,Scan"},
			"UntypedBigfloat": {reflect.TypeOf((*q.UntypedBigfloat)(nil)).Elem(), "", ""},
			"UntypedBigint":   {reflect.TypeOf((*q.UntypedBigint)(nil)).Elem(), "", ""},
			"UntypedBigrat":   {reflect.TypeOf((*q.UntypedBigrat)(nil)).Elem(), "", ""},
		},
		AliasTypes: map[string]reflect.Type{
			"Gop_ninteger":            reflect.TypeOf((*uint)(nil)).Elem(),
			"UntypedBigfloat_Default": reflect.TypeOf((*q.UntypedBigfloat_Default)(nil)).Elem(),
			"UntypedBigint_Default":   reflect.TypeOf((*q.UntypedBigint_Default)(nil)).Elem(),
			"UntypedBigrat_Default":   reflect.TypeOf((*q.UntypedBigrat_Default)(nil)).Elem(),
		},
		Vars: map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"Bigint_Cast__0":        reflect.ValueOf(q.Bigint_Cast__0),
			"Bigint_Cast__1":        reflect.ValueOf(q.Bigint_Cast__1),
			"Bigint_Cast__2":        reflect.ValueOf(q.Bigint_Cast__2),
			"Bigint_Cast__3":        reflect.ValueOf(q.Bigint_Cast__3),
			"Bigint_Cast__4":        reflect.ValueOf(q.Bigint_Cast__4),
			"Bigint_Cast__5":        reflect.ValueOf(q.Bigint_Cast__5),
			"Bigint_Cast__6":        reflect.ValueOf(q.Bigint_Cast__6),
			"Bigint_Cast__7":        reflect.ValueOf(q.Bigint_Cast__7),
			"Bigint_Init__0":        reflect.ValueOf(q.Bigint_Init__0),
			"Bigint_Init__1":        reflect.ValueOf(q.Bigint_Init__1),
			"Bigint_Init__2":        reflect.ValueOf(q.Bigint_Init__2),
			"Bigrat_Cast__0":        reflect.ValueOf(q.Bigrat_Cast__0),
			"Bigrat_Cast__1":        reflect.ValueOf(q.Bigrat_Cast__1),
			"Bigrat_Cast__2":        reflect.ValueOf(q.Bigrat_Cast__2),
			"Bigrat_Cast__3":        reflect.ValueOf(q.Bigrat_Cast__3),
			"Bigrat_Cast__4":        reflect.ValueOf(q.Bigrat_Cast__4),
			"Bigrat_Cast__5":        reflect.ValueOf(q.Bigrat_Cast__5),
			"Bigrat_Cast__6":        reflect.ValueOf(q.Bigrat_Cast__6),
			"Bigrat_Init__0":        reflect.ValueOf(q.Bigrat_Init__0),
			"Bigrat_Init__1":        reflect.ValueOf(q.Bigrat_Init__1),
			"Bigrat_Init__2":        reflect.ValueOf(q.Bigrat_Init__2),
			"FormatInt128":          reflect.ValueOf(q.FormatInt128),
			"FormatUint128":         reflect.ValueOf(q.FormatUint128),
			"Gop_istmp":             reflect.ValueOf(q.Gop_istmp),
			"Int128_Cast__0":        reflect.ValueOf(q.Int128_Cast__0),
			"Int128_Cast__1":        reflect.ValueOf(q.Int128_Cast__1),
			"Int128_Cast__2":        reflect.ValueOf(q.Int128_Cast__2),
			"Int128_Cast__3":        reflect.ValueOf(q.Int128_Cast__3),
			"Int128_Cast__4":        reflect.ValueOf(q.Int128_Cast__4),
			"Int128_Cast__5":        reflect.ValueOf(q.Int128_Cast__5),
			"Int128_Cast__6":        reflect.ValueOf(q.Int128_Cast__6),
			"Int128_Cast__7":        reflect.ValueOf(q.Int128_Cast__7),
			"Int128_Cast__8":        reflect.ValueOf(q.Int128_Cast__8),
			"Int128_Cast__9":        reflect.ValueOf(q.Int128_Cast__9),
			"Int128_Cast__a":        reflect.ValueOf(q.Int128_Cast__a),
			"Int128_Init__0":        reflect.ValueOf(q.Int128_Init__0),
			"Int128_Init__1":        reflect.ValueOf(q.Int128_Init__1),
			"ParseInt128":           reflect.ValueOf(q.ParseInt128),
			"ParseUint128":          reflect.ValueOf(q.ParseUint128),
			"Uint128_Cast__0":       reflect.ValueOf(q.Uint128_Cast__0),
			"Uint128_Cast__1":       reflect.ValueOf(q.Uint128_Cast__1),
			"Uint128_Cast__2":       reflect.ValueOf(q.Uint128_Cast__2),
			"Uint128_Cast__3":       reflect.ValueOf(q.Uint128_Cast__3),
			"Uint128_Cast__4":       reflect.ValueOf(q.Uint128_Cast__4),
			"Uint128_Cast__5":       reflect.ValueOf(q.Uint128_Cast__5),
			"Uint128_Cast__6":       reflect.ValueOf(q.Uint128_Cast__6),
			"Uint128_Cast__7":       reflect.ValueOf(q.Uint128_Cast__7),
			"Uint128_Cast__8":       reflect.ValueOf(q.Uint128_Cast__8),
			"Uint128_Cast__9":       reflect.ValueOf(q.Uint128_Cast__9),
			"Uint128_Cast__a":       reflect.ValueOf(q.Uint128_Cast__a),
			"Uint128_Cast__b":       reflect.ValueOf(q.Uint128_Cast__b),
			"Uint128_Cast__c":       reflect.ValueOf(q.Uint128_Cast__c),
			"Uint128_Init__0":       reflect.ValueOf(q.Uint128_Init__0),
			"Uint128_Init__1":       reflect.ValueOf(q.Uint128_Init__1),
			"UntypedBigint_Init__0": reflect.ValueOf(q.UntypedBigint_Init__0),
			"UntypedBigrat_Init__0": reflect.ValueOf(q.UntypedBigrat_Init__0),
			"UntypedBigrat_Init__1": reflect.ValueOf(q.UntypedBigrat_Init__1),
		},
		TypedConsts: map[string]igop.TypedConst{},
		UntypedConsts: map[string]igop.UntypedConst{
			"GopPackage": {"untyped bool", constant.MakeBool(bool(q.GopPackage))},
		},
	})
}
