// export by github.com/goplus/interp/cmd/qexp

package bits

import (
	"math/bits"

	"github.com/goplus/interp"
)

func init() {
	interp.RegisterPackage("math/bits", extMap, typList)
}

var extMap = map[string]interface{}{
	"math/bits.Add":             bits.Add,
	"math/bits.Add32":           bits.Add32,
	"math/bits.Add64":           bits.Add64,
	"math/bits.Div":             bits.Div,
	"math/bits.Div32":           bits.Div32,
	"math/bits.Div64":           bits.Div64,
	"math/bits.LeadingZeros":    bits.LeadingZeros,
	"math/bits.LeadingZeros16":  bits.LeadingZeros16,
	"math/bits.LeadingZeros32":  bits.LeadingZeros32,
	"math/bits.LeadingZeros64":  bits.LeadingZeros64,
	"math/bits.LeadingZeros8":   bits.LeadingZeros8,
	"math/bits.Len":             bits.Len,
	"math/bits.Len16":           bits.Len16,
	"math/bits.Len32":           bits.Len32,
	"math/bits.Len64":           bits.Len64,
	"math/bits.Len8":            bits.Len8,
	"math/bits.Mul":             bits.Mul,
	"math/bits.Mul32":           bits.Mul32,
	"math/bits.Mul64":           bits.Mul64,
	"math/bits.OnesCount":       bits.OnesCount,
	"math/bits.OnesCount16":     bits.OnesCount16,
	"math/bits.OnesCount32":     bits.OnesCount32,
	"math/bits.OnesCount64":     bits.OnesCount64,
	"math/bits.OnesCount8":      bits.OnesCount8,
	"math/bits.Rem":             bits.Rem,
	"math/bits.Rem32":           bits.Rem32,
	"math/bits.Rem64":           bits.Rem64,
	"math/bits.Reverse":         bits.Reverse,
	"math/bits.Reverse16":       bits.Reverse16,
	"math/bits.Reverse32":       bits.Reverse32,
	"math/bits.Reverse64":       bits.Reverse64,
	"math/bits.Reverse8":        bits.Reverse8,
	"math/bits.ReverseBytes":    bits.ReverseBytes,
	"math/bits.ReverseBytes16":  bits.ReverseBytes16,
	"math/bits.ReverseBytes32":  bits.ReverseBytes32,
	"math/bits.ReverseBytes64":  bits.ReverseBytes64,
	"math/bits.RotateLeft":      bits.RotateLeft,
	"math/bits.RotateLeft16":    bits.RotateLeft16,
	"math/bits.RotateLeft32":    bits.RotateLeft32,
	"math/bits.RotateLeft64":    bits.RotateLeft64,
	"math/bits.RotateLeft8":     bits.RotateLeft8,
	"math/bits.Sub":             bits.Sub,
	"math/bits.Sub32":           bits.Sub32,
	"math/bits.Sub64":           bits.Sub64,
	"math/bits.TrailingZeros":   bits.TrailingZeros,
	"math/bits.TrailingZeros16": bits.TrailingZeros16,
	"math/bits.TrailingZeros32": bits.TrailingZeros32,
	"math/bits.TrailingZeros64": bits.TrailingZeros64,
	"math/bits.TrailingZeros8":  bits.TrailingZeros8,
}

var typList = []interface{}{}
