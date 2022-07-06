package constant

import (
	"go/constant"
	"go/token"
	"testing"
)

func TestExactConstant1(t *testing.T) {
	const c = "3.14159265358979323846264338327950288419716939937510582097494459"
	x := constant.MakeFromLiteral(c, token.FLOAT, 0)
	if v := ExactConstant(x); v != c {
		t.Fatalf(v)
	}
}

func TestExactConstant2(t *testing.T) {
	const c = "1/713623846352979940529142984724747568191373312"
	x := constant.BinaryOp(constant.MakeFromLiteral("1", token.INT, 0), token.QUO, constant.MakeFromLiteral("713623846352979940529142984724747568191373312", token.INT, 0))
	if v := ExactConstant(x); v != c {
		t.Fatalf(v)
	}
}

func TestExactConstant3(t *testing.T) {
	const c = "1.01"
	x := constant.BinaryOp(constant.MakeFromLiteral("101", token.INT, 0), token.QUO, constant.MakeFromLiteral("100", token.INT, 0))
	if v := ExactConstant(x); v != c {
		t.Fatalf(v)
	}
}

func TestExactConstant4(t *testing.T) {
	const c = "3.14e+100"
	x := constant.MakeFromLiteral("3.14e100", token.FLOAT, 0)
	if v := ExactConstant(x); v != c {
		t.Fatalf(v)
	}
}

func TestExactConstant5(t *testing.T) {
	const c = "1.1+5i"
	x := constant.BinaryOp(constant.MakeFromLiteral("1.1", token.FLOAT, 0), token.ADD, constant.MakeImag(constant.MakeInt64(5)))
	if v := ExactConstant(x); v != c {
		t.Fatalf(v)
	}
}
