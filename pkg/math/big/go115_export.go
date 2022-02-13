// export by github.com/goplus/gossa/cmd/qexp

//+build go1.15,!go1.16

package big

import (
	q "math/big"

	"go/constant"
	"reflect"

	"github.com/goplus/gossa"
)

func init() {
	gossa.RegisterPackage(&gossa.Package{
		Name: "big",
		Path: "math/big",
		Deps: map[string]string{
			"bytes":           "bytes",
			"encoding/binary": "binary",
			"errors":          "errors",
			"fmt":             "fmt",
			"internal/cpu":    "cpu",
			"io":              "io",
			"math":            "math",
			"math/bits":       "bits",
			"math/rand":       "rand",
			"strconv":         "strconv",
			"strings":         "strings",
			"sync":            "sync",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]gossa.NamedType{
			"Accuracy":     {reflect.TypeOf((*q.Accuracy)(nil)).Elem(), "String", ""},
			"ErrNaN":       {reflect.TypeOf((*q.ErrNaN)(nil)).Elem(), "Error", ""},
			"Float":        {reflect.TypeOf((*q.Float)(nil)).Elem(), "", "Abs,Acc,Add,Append,Cmp,Copy,Float32,Float64,Format,GobDecode,GobEncode,Int,Int64,IsInf,IsInt,MantExp,MarshalText,MinPrec,Mode,Mul,Neg,Parse,Prec,Quo,Rat,Scan,Set,SetFloat64,SetInf,SetInt,SetInt64,SetMantExp,SetMode,SetPrec,SetRat,SetString,SetUint64,Sign,Signbit,Sqrt,String,Sub,Text,Uint64,UnmarshalText,fmtB,fmtP,fmtX,ord,pow5,round,scan,setBits64,setExpAndRound,sqrtInverse,uadd,ucmp,umul,uquo,usub,validate"},
			"Int":          {reflect.TypeOf((*q.Int)(nil)).Elem(), "", "Abs,Add,And,AndNot,Append,Binomial,Bit,BitLen,Bits,Bytes,Cmp,CmpAbs,Div,DivMod,Exp,FillBytes,Format,GCD,GobDecode,GobEncode,Int64,IsInt64,IsUint64,Lsh,MarshalJSON,MarshalText,Mod,ModInverse,ModSqrt,Mul,MulRange,Neg,Not,Or,ProbablyPrime,Quo,QuoRem,Rand,Rem,Rsh,Scan,Set,SetBit,SetBits,SetBytes,SetInt64,SetString,SetUint64,Sign,Sqrt,String,Sub,Text,TrailingZeroBits,Uint64,UnmarshalJSON,UnmarshalText,Xor,lehmerGCD,modSqrt3Mod4Prime,modSqrt5Mod8Prime,modSqrtTonelliShanks,scaleDenom,scan,setFromScanner"},
			"Rat":          {reflect.TypeOf((*q.Rat)(nil)).Elem(), "", "Abs,Add,Cmp,Denom,Float32,Float64,FloatString,GobDecode,GobEncode,Inv,IsInt,MarshalText,Mul,Neg,Num,Quo,RatString,Scan,Set,SetFloat64,SetFrac,SetFrac64,SetInt,SetInt64,SetString,SetUint64,Sign,String,Sub,UnmarshalText,marshal,norm"},
			"RoundingMode": {reflect.TypeOf((*q.RoundingMode)(nil)).Elem(), "String", ""},
			"Word":         {reflect.TypeOf((*q.Word)(nil)).Elem(), "", ""},
		},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"Jacobi":     reflect.ValueOf(q.Jacobi),
			"NewFloat":   reflect.ValueOf(q.NewFloat),
			"NewInt":     reflect.ValueOf(q.NewInt),
			"NewRat":     reflect.ValueOf(q.NewRat),
			"ParseFloat": reflect.ValueOf(q.ParseFloat),
		},
		TypedConsts: map[string]gossa.TypedConst{
			"Above":         {reflect.TypeOf(q.Above), constant.MakeInt64(int64(q.Above))},
			"AwayFromZero":  {reflect.TypeOf(q.AwayFromZero), constant.MakeInt64(int64(q.AwayFromZero))},
			"Below":         {reflect.TypeOf(q.Below), constant.MakeInt64(int64(q.Below))},
			"Exact":         {reflect.TypeOf(q.Exact), constant.MakeInt64(int64(q.Exact))},
			"ToNearestAway": {reflect.TypeOf(q.ToNearestAway), constant.MakeInt64(int64(q.ToNearestAway))},
			"ToNearestEven": {reflect.TypeOf(q.ToNearestEven), constant.MakeInt64(int64(q.ToNearestEven))},
			"ToNegativeInf": {reflect.TypeOf(q.ToNegativeInf), constant.MakeInt64(int64(q.ToNegativeInf))},
			"ToPositiveInf": {reflect.TypeOf(q.ToPositiveInf), constant.MakeInt64(int64(q.ToPositiveInf))},
			"ToZero":        {reflect.TypeOf(q.ToZero), constant.MakeInt64(int64(q.ToZero))},
		},
		UntypedConsts: map[string]gossa.UntypedConst{
			"MaxBase": {"untyped rune", constant.MakeInt64(int64(q.MaxBase))},
			"MaxExp":  {"untyped int", constant.MakeInt64(int64(q.MaxExp))},
			"MaxPrec": {"untyped int", constant.MakeInt64(int64(q.MaxPrec))},
			"MinExp":  {"untyped int", constant.MakeInt64(int64(q.MinExp))},
		},
	})
}
