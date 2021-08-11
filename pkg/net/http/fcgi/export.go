// export by github.com/goplus/interp/cmd/qexp

package fcgi

import (
	"net/http/fcgi"

	"github.com/goplus/interp"
)

func init() {
	interp.RegisterPackage("net/http/fcgi", extMap, typList)
}

var extMap = map[string]interface{}{
	"net/http/fcgi.ErrConnClosed":     &fcgi.ErrConnClosed,
	"net/http/fcgi.ErrRequestAborted": &fcgi.ErrRequestAborted,
	"net/http/fcgi.ProcessEnv":        fcgi.ProcessEnv,
	"net/http/fcgi.Serve":             fcgi.Serve,
}

var typList = []interface{}{}
