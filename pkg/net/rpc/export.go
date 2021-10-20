// export by github.com/goplus/gossa/cmd/qexp

package rpc

import (
	"net/rpc"

	"github.com/goplus/gossa"
)

func init() {
	gossa.RegisterPackage("net/rpc", extMap, typList)
}

var extMap = map[string]interface{}{
	"(*net/rpc.Client).Call":                   (*rpc.Client).Call,
	"(*net/rpc.Client).Close":                  (*rpc.Client).Close,
	"(*net/rpc.Client).Go":                     (*rpc.Client).Go,
	"(*net/rpc.Server).Accept":                 (*rpc.Server).Accept,
	"(*net/rpc.Server).HandleHTTP":             (*rpc.Server).HandleHTTP,
	"(*net/rpc.Server).Register":               (*rpc.Server).Register,
	"(*net/rpc.Server).RegisterName":           (*rpc.Server).RegisterName,
	"(*net/rpc.Server).ServeCodec":             (*rpc.Server).ServeCodec,
	"(*net/rpc.Server).ServeConn":              (*rpc.Server).ServeConn,
	"(*net/rpc.Server).ServeHTTP":              (*rpc.Server).ServeHTTP,
	"(*net/rpc.Server).ServeRequest":           (*rpc.Server).ServeRequest,
	"(net/rpc.ClientCodec).Close":              (rpc.ClientCodec).Close,
	"(net/rpc.ClientCodec).ReadResponseBody":   (rpc.ClientCodec).ReadResponseBody,
	"(net/rpc.ClientCodec).ReadResponseHeader": (rpc.ClientCodec).ReadResponseHeader,
	"(net/rpc.ClientCodec).WriteRequest":       (rpc.ClientCodec).WriteRequest,
	"(net/rpc.ServerCodec).Close":              (rpc.ServerCodec).Close,
	"(net/rpc.ServerCodec).ReadRequestBody":    (rpc.ServerCodec).ReadRequestBody,
	"(net/rpc.ServerCodec).ReadRequestHeader":  (rpc.ServerCodec).ReadRequestHeader,
	"(net/rpc.ServerCodec).WriteResponse":      (rpc.ServerCodec).WriteResponse,
	"(net/rpc.ServerError).Error":              (rpc.ServerError).Error,
	"net/rpc.Accept":                           rpc.Accept,
	"net/rpc.DefaultServer":                    &rpc.DefaultServer,
	"net/rpc.Dial":                             rpc.Dial,
	"net/rpc.DialHTTP":                         rpc.DialHTTP,
	"net/rpc.DialHTTPPath":                     rpc.DialHTTPPath,
	"net/rpc.ErrShutdown":                      &rpc.ErrShutdown,
	"net/rpc.HandleHTTP":                       rpc.HandleHTTP,
	"net/rpc.NewClient":                        rpc.NewClient,
	"net/rpc.NewClientWithCodec":               rpc.NewClientWithCodec,
	"net/rpc.NewServer":                        rpc.NewServer,
	"net/rpc.Register":                         rpc.Register,
	"net/rpc.RegisterName":                     rpc.RegisterName,
	"net/rpc.ServeCodec":                       rpc.ServeCodec,
	"net/rpc.ServeConn":                        rpc.ServeConn,
	"net/rpc.ServeRequest":                     rpc.ServeRequest,
}

var typList = []interface{}{
	(*rpc.Call)(nil),
	(*rpc.Client)(nil),
	(*rpc.ClientCodec)(nil),
	(*rpc.Request)(nil),
	(*rpc.Response)(nil),
	(*rpc.Server)(nil),
	(*rpc.ServerCodec)(nil),
	(*rpc.ServerError)(nil),
}
