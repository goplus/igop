package igop

import (
	"fmt"
	"go/constant"
	"go/token"
	"strings"
)

func ExactConstant(c constant.Value) string {
	switch c.Kind() {
	case constant.Bool:
		if constant.BoolVal(c) {
			return "true"
		} else {
			return "false"
		}
	case constant.String:
		return constant.StringVal(c)
	case constant.Int:
		return c.ExactString()
	case constant.Float:
		s := c.ExactString()
		if pos := strings.IndexByte(s, '/'); pos >= 0 {
			sx := s[:pos]
			sy := s[pos+1:]
			// simplify 314/100 => 3.14
			// 80901699437494742410229341718281905886015458990288143106772431
			// 50000000000000000000000000000000000000000000000000000000000000
			if strings.HasPrefix(sy, "1") && strings.Count(sy, "0") == len(sy)-1 {
				if len(sx) == len(sy) {
					return fmt.Sprintf("%v.%v", sx[:1], sx[1:])
				} else if len(sx) == len(sy)-1 {
					return fmt.Sprintf("0.%v", sx)
				} else if len(sx) < len(sy) {
					return fmt.Sprintf("%v.%ve-%v", sx[:1], sx[1:], len(sy)-len(sx))
				}
			} else if strings.HasPrefix(sy, "5") && strings.Count(sy, "0") == len(sy)-1 {
				if len(sx) == len(sy) {
					c := constant.BinaryOp(constant.MakeFromLiteral(sx, token.INT, 0), token.MUL, constant.MakeInt64(2))
					sx = c.ExactString()
					return fmt.Sprintf("%v.%v", sx[:1], sx[1:])
				}
			} else if strings.HasPrefix(sx, "1") && strings.Count(sx, "0") == len(sx)-1 {
				// skip
			}
			return fmt.Sprintf("%v/%v", sx, sy)
		}
		if pos := strings.LastIndexAny(s, "123456789"); pos != -1 {
			sx := s[:pos+1]
			return fmt.Sprintf("%v.%ve+%v", sx[:1], sx[1:], len(s)-1)
		}
		return s
	case constant.Complex:
		re := ExactConstant(constant.Real(c))
		im := ExactConstant(constant.Imag(c))
		return fmt.Sprintf("%v+%vi", re, im)
	default:
		panic("unreachable")
	}
}
