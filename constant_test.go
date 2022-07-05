package igop

import (
	"go/constant"
	"go/token"
	"testing"
)

func TestExactConstant1(t *testing.T) {
	const pi = "3.14159265358979323846264338327950288419716939937510582097494459"
	x := constant.MakeFromLiteral(pi, token.FLOAT, 0)
	if v := ExactConstant(x); v != pi {
		t.Fatalf(v)
	}
}

func TestExactConstant2(t *testing.T) {
	const SmallestNonzeroFloat32 = "1/713623846352979940529142984724747568191373312"
	x := constant.BinaryOp(constant.MakeFromLiteral("1", token.INT, 0), token.QUO, constant.MakeFromLiteral("713623846352979940529142984724747568191373312", token.INT, 0))
	if v := ExactConstant(x); v != SmallestNonzeroFloat32 {
		t.Fatalf(v)
	}
}
