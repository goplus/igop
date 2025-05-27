// export by github.com/goplus/ixgo/cmd/qexp

//+build go1.14,!go1.15

package big

import (
	q "math/big"

	"go/constant"
	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage(&ixgo.Package{
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
		NamedTypes: map[string]reflect.Type{
			"Accuracy":     reflect.TypeOf((*q.Accuracy)(nil)).Elem(),
			"ErrNaN":       reflect.TypeOf((*q.ErrNaN)(nil)).Elem(),
			"Float":        reflect.TypeOf((*q.Float)(nil)).Elem(),
			"Int":          reflect.TypeOf((*q.Int)(nil)).Elem(),
			"Rat":          reflect.TypeOf((*q.Rat)(nil)).Elem(),
			"RoundingMode": reflect.TypeOf((*q.RoundingMode)(nil)).Elem(),
			"Word":         reflect.TypeOf((*q.Word)(nil)).Elem(),
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
		TypedConsts: map[string]ixgo.TypedConst{
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
		UntypedConsts: map[string]ixgo.UntypedConst{
			"MaxBase": {"untyped rune", constant.MakeInt64(int64(q.MaxBase))},
			"MaxExp":  {"untyped int", constant.MakeInt64(int64(q.MaxExp))},
			"MaxPrec": {"untyped int", constant.MakeInt64(int64(q.MaxPrec))},
			"MinExp":  {"untyped int", constant.MakeInt64(int64(q.MinExp))},
		},
	})
}
