// export by github.com/goplus/interp/cmd/qexp

package expvar

import (
	"expvar"

	"github.com/goplus/interp"
)

func init() {
	interp.RegisterPackage("expvar", extMap, typList)
}

var extMap = map[string]interface{}{
	"(*expvar.Float).Add":     (*expvar.Float).Add,
	"(*expvar.Float).Set":     (*expvar.Float).Set,
	"(*expvar.Float).String":  (*expvar.Float).String,
	"(*expvar.Float).Value":   (*expvar.Float).Value,
	"(*expvar.Int).Add":       (*expvar.Int).Add,
	"(*expvar.Int).Set":       (*expvar.Int).Set,
	"(*expvar.Int).String":    (*expvar.Int).String,
	"(*expvar.Int).Value":     (*expvar.Int).Value,
	"(*expvar.Map).Add":       (*expvar.Map).Add,
	"(*expvar.Map).AddFloat":  (*expvar.Map).AddFloat,
	"(*expvar.Map).Delete":    (*expvar.Map).Delete,
	"(*expvar.Map).Do":        (*expvar.Map).Do,
	"(*expvar.Map).Get":       (*expvar.Map).Get,
	"(*expvar.Map).Init":      (*expvar.Map).Init,
	"(*expvar.Map).Set":       (*expvar.Map).Set,
	"(*expvar.Map).String":    (*expvar.Map).String,
	"(*expvar.String).Set":    (*expvar.String).Set,
	"(*expvar.String).String": (*expvar.String).String,
	"(*expvar.String).Value":  (*expvar.String).Value,
	"(expvar.Func).String":    (expvar.Func).String,
	"(expvar.Func).Value":     (expvar.Func).Value,
	"(expvar.Var).String":     (expvar.Var).String,
	"expvar.Do":               expvar.Do,
	"expvar.Get":              expvar.Get,
	"expvar.Handler":          expvar.Handler,
	"expvar.NewFloat":         expvar.NewFloat,
	"expvar.NewInt":           expvar.NewInt,
	"expvar.NewMap":           expvar.NewMap,
	"expvar.NewString":        expvar.NewString,
	"expvar.Publish":          expvar.Publish,
}

var typList = []interface{}{
	(*expvar.Float)(nil),
	(*expvar.Func)(nil),
	(*expvar.Int)(nil),
	(*expvar.KeyValue)(nil),
	(*expvar.Map)(nil),
	(*expvar.String)(nil),
	(*expvar.Var)(nil),
}
