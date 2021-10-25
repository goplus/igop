// export by github.com/goplus/gossa/cmd/qexp

package cmplx

import (
	"math/cmplx"

	"github.com/goplus/gossa"
)

func init() {
	gossa.RegisterPackage("math/cmplx", extMap, typList)
}

var extMap = map[string]interface{}{
	"math/cmplx.Abs":   cmplx.Abs,
	"math/cmplx.Acos":  cmplx.Acos,
	"math/cmplx.Acosh": cmplx.Acosh,
	"math/cmplx.Asin":  cmplx.Asin,
	"math/cmplx.Asinh": cmplx.Asinh,
	"math/cmplx.Atan":  cmplx.Atan,
	"math/cmplx.Atanh": cmplx.Atanh,
	"math/cmplx.Conj":  cmplx.Conj,
	"math/cmplx.Cos":   cmplx.Cos,
	"math/cmplx.Cosh":  cmplx.Cosh,
	"math/cmplx.Cot":   cmplx.Cot,
	"math/cmplx.Exp":   cmplx.Exp,
	"math/cmplx.Inf":   cmplx.Inf,
	"math/cmplx.IsInf": cmplx.IsInf,
	"math/cmplx.IsNaN": cmplx.IsNaN,
	"math/cmplx.Log":   cmplx.Log,
	"math/cmplx.Log10": cmplx.Log10,
	"math/cmplx.NaN":   cmplx.NaN,
	"math/cmplx.Phase": cmplx.Phase,
	"math/cmplx.Polar": cmplx.Polar,
	"math/cmplx.Pow":   cmplx.Pow,
	"math/cmplx.Rect":  cmplx.Rect,
	"math/cmplx.Sin":   cmplx.Sin,
	"math/cmplx.Sinh":  cmplx.Sinh,
	"math/cmplx.Sqrt":  cmplx.Sqrt,
	"math/cmplx.Tan":   cmplx.Tan,
	"math/cmplx.Tanh":  cmplx.Tanh,
}

var typList = []interface{}{}
