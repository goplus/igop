// export by github.com/goplus/gossa/cmd/qexp

package math

import (
	"math"

	"github.com/goplus/gossa"
)

func init() {
	gossa.RegisterPackage("math", extMap, typList)
}

var extMap = map[string]interface{}{
	"math.Abs":             math.Abs,
	"math.Acos":            math.Acos,
	"math.Acosh":           math.Acosh,
	"math.Asin":            math.Asin,
	"math.Asinh":           math.Asinh,
	"math.Atan":            math.Atan,
	"math.Atan2":           math.Atan2,
	"math.Atanh":           math.Atanh,
	"math.Cbrt":            math.Cbrt,
	"math.Ceil":            math.Ceil,
	"math.Copysign":        math.Copysign,
	"math.Cos":             math.Cos,
	"math.Cosh":            math.Cosh,
	"math.Dim":             math.Dim,
	"math.Erf":             math.Erf,
	"math.Erfc":            math.Erfc,
	"math.Erfcinv":         math.Erfcinv,
	"math.Erfinv":          math.Erfinv,
	"math.Exp":             math.Exp,
	"math.Exp2":            math.Exp2,
	"math.Expm1":           math.Expm1,
	"math.FMA":             math.FMA,
	"math.Float32bits":     math.Float32bits,
	"math.Float32frombits": math.Float32frombits,
	"math.Float64bits":     math.Float64bits,
	"math.Float64frombits": math.Float64frombits,
	"math.Floor":           math.Floor,
	"math.Frexp":           math.Frexp,
	"math.Gamma":           math.Gamma,
	"math.Hypot":           math.Hypot,
	"math.Ilogb":           math.Ilogb,
	"math.Inf":             math.Inf,
	"math.IsInf":           math.IsInf,
	"math.IsNaN":           math.IsNaN,
	"math.J0":              math.J0,
	"math.J1":              math.J1,
	"math.Jn":              math.Jn,
	"math.Ldexp":           math.Ldexp,
	"math.Lgamma":          math.Lgamma,
	"math.Log":             math.Log,
	"math.Log10":           math.Log10,
	"math.Log1p":           math.Log1p,
	"math.Log2":            math.Log2,
	"math.Logb":            math.Logb,
	"math.Max":             math.Max,
	"math.Min":             math.Min,
	"math.Mod":             math.Mod,
	"math.Modf":            math.Modf,
	"math.NaN":             math.NaN,
	"math.Nextafter":       math.Nextafter,
	"math.Nextafter32":     math.Nextafter32,
	"math.Pow":             math.Pow,
	"math.Pow10":           math.Pow10,
	"math.Remainder":       math.Remainder,
	"math.Round":           math.Round,
	"math.RoundToEven":     math.RoundToEven,
	"math.Signbit":         math.Signbit,
	"math.Sin":             math.Sin,
	"math.Sincos":          math.Sincos,
	"math.Sinh":            math.Sinh,
	"math.Sqrt":            math.Sqrt,
	"math.Tan":             math.Tan,
	"math.Tanh":            math.Tanh,
	"math.Trunc":           math.Trunc,
	"math.Y0":              math.Y0,
	"math.Y1":              math.Y1,
	"math.Yn":              math.Yn,
}

var typList = []interface{}{}
