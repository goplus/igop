// export by github.com/goplus/interp/cmd/qexp

package elliptic

import (
	"crypto/elliptic"

	"github.com/goplus/interp"
)

func init() {
	interp.RegisterPackage("crypto/elliptic", extMap, typList)
}

var extMap = map[string]interface{}{
	"(*crypto/elliptic.CurveParams).Add":            (*elliptic.CurveParams).Add,
	"(*crypto/elliptic.CurveParams).Double":         (*elliptic.CurveParams).Double,
	"(*crypto/elliptic.CurveParams).IsOnCurve":      (*elliptic.CurveParams).IsOnCurve,
	"(*crypto/elliptic.CurveParams).Params":         (*elliptic.CurveParams).Params,
	"(*crypto/elliptic.CurveParams).ScalarBaseMult": (*elliptic.CurveParams).ScalarBaseMult,
	"(*crypto/elliptic.CurveParams).ScalarMult":     (*elliptic.CurveParams).ScalarMult,
	"(crypto/elliptic.Curve).Add":                   (elliptic.Curve).Add,
	"(crypto/elliptic.Curve).Double":                (elliptic.Curve).Double,
	"(crypto/elliptic.Curve).IsOnCurve":             (elliptic.Curve).IsOnCurve,
	"(crypto/elliptic.Curve).Params":                (elliptic.Curve).Params,
	"(crypto/elliptic.Curve).ScalarBaseMult":        (elliptic.Curve).ScalarBaseMult,
	"(crypto/elliptic.Curve).ScalarMult":            (elliptic.Curve).ScalarMult,
	"crypto/elliptic.GenerateKey":                   elliptic.GenerateKey,
	"crypto/elliptic.Marshal":                       elliptic.Marshal,
	"crypto/elliptic.P224":                          elliptic.P224,
	"crypto/elliptic.P256":                          elliptic.P256,
	"crypto/elliptic.P384":                          elliptic.P384,
	"crypto/elliptic.P521":                          elliptic.P521,
	"crypto/elliptic.Unmarshal":                     elliptic.Unmarshal,
}

var typList = []interface{}{
	(*elliptic.Curve)(nil),
	(*elliptic.CurveParams)(nil),
}
