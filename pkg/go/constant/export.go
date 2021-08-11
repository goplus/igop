// export by github.com/goplus/interp/cmd/qexp

package constant

import (
	"go/constant"

	"github.com/goplus/interp"
)

func init() {
	interp.RegisterPackage("go/constant", extMap, typList)
}

var extMap = map[string]interface{}{
	"(go/constant.Value).ExactString": (constant.Value).ExactString,
	"(go/constant.Value).Kind":        (constant.Value).Kind,
	"(go/constant.Value).String":      (constant.Value).String,
	"go/constant.BinaryOp":            constant.BinaryOp,
	"go/constant.BitLen":              constant.BitLen,
	"go/constant.BoolVal":             constant.BoolVal,
	"go/constant.Bytes":               constant.Bytes,
	"go/constant.Compare":             constant.Compare,
	"go/constant.Denom":               constant.Denom,
	"go/constant.Float32Val":          constant.Float32Val,
	"go/constant.Float64Val":          constant.Float64Val,
	"go/constant.Imag":                constant.Imag,
	"go/constant.Int64Val":            constant.Int64Val,
	"go/constant.Make":                constant.Make,
	"go/constant.MakeBool":            constant.MakeBool,
	"go/constant.MakeFloat64":         constant.MakeFloat64,
	"go/constant.MakeFromBytes":       constant.MakeFromBytes,
	"go/constant.MakeFromLiteral":     constant.MakeFromLiteral,
	"go/constant.MakeImag":            constant.MakeImag,
	"go/constant.MakeInt64":           constant.MakeInt64,
	"go/constant.MakeString":          constant.MakeString,
	"go/constant.MakeUint64":          constant.MakeUint64,
	"go/constant.MakeUnknown":         constant.MakeUnknown,
	"go/constant.Num":                 constant.Num,
	"go/constant.Real":                constant.Real,
	"go/constant.Shift":               constant.Shift,
	"go/constant.Sign":                constant.Sign,
	"go/constant.StringVal":           constant.StringVal,
	"go/constant.ToComplex":           constant.ToComplex,
	"go/constant.ToFloat":             constant.ToFloat,
	"go/constant.ToInt":               constant.ToInt,
	"go/constant.Uint64Val":           constant.Uint64Val,
	"go/constant.UnaryOp":             constant.UnaryOp,
	"go/constant.Val":                 constant.Val,
}

var typList = []interface{}{
	(*constant.Kind)(nil),
	(*constant.Value)(nil),
}
