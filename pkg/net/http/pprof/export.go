// export by github.com/goplus/gossa/cmd/qexp

package pprof

import (
	"net/http/pprof"

	"github.com/goplus/gossa"
)

func init() {
	gossa.RegisterPackage("net/http/pprof", extMap, typList)
}

var extMap = map[string]interface{}{
	"net/http/pprof.Cmdline": pprof.Cmdline,
	"net/http/pprof.Handler": pprof.Handler,
	"net/http/pprof.Index":   pprof.Index,
	"net/http/pprof.Profile": pprof.Profile,
	"net/http/pprof.Symbol":  pprof.Symbol,
	"net/http/pprof.Trace":   pprof.Trace,
}

var typList = []interface{}{}
