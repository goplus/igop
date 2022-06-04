// export by github.com/goplus/igop/cmd/qexp

//+build go1.16,!go1.17

package metrics

import (
	q "runtime/metrics"

	"go/constant"
	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "metrics",
		Path: "runtime/metrics",
		Deps: map[string]string{
			"math":    "math",
			"runtime": "runtime",
			"unsafe":  "unsafe",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]igop.NamedType{
			"Description":      {reflect.TypeOf((*q.Description)(nil)).Elem(), "", ""},
			"Float64Histogram": {reflect.TypeOf((*q.Float64Histogram)(nil)).Elem(), "", ""},
			"Sample":           {reflect.TypeOf((*q.Sample)(nil)).Elem(), "", ""},
			"Value":            {reflect.TypeOf((*q.Value)(nil)).Elem(), "Float64,Float64Histogram,Kind,Uint64", ""},
			"ValueKind":        {reflect.TypeOf((*q.ValueKind)(nil)).Elem(), "", ""},
		},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"All":  reflect.ValueOf(q.All),
			"Read": reflect.ValueOf(q.Read),
		},
		TypedConsts: map[string]igop.TypedConst{
			"KindBad":              {reflect.TypeOf(q.KindBad), constant.MakeInt64(int64(q.KindBad))},
			"KindFloat64":          {reflect.TypeOf(q.KindFloat64), constant.MakeInt64(int64(q.KindFloat64))},
			"KindFloat64Histogram": {reflect.TypeOf(q.KindFloat64Histogram), constant.MakeInt64(int64(q.KindFloat64Histogram))},
			"KindUint64":           {reflect.TypeOf(q.KindUint64), constant.MakeInt64(int64(q.KindUint64))},
		},
		UntypedConsts: map[string]igop.UntypedConst{},
	})
}
