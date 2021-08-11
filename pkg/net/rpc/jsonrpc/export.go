// export by github.com/goplus/interp/cmd/qexp

package jsonrpc

import (
	"net/rpc/jsonrpc"

	"github.com/goplus/interp"
)

func init() {
	interp.RegisterPackage("net/rpc/jsonrpc", extMap, typList)
}

var extMap = map[string]interface{}{
	"net/rpc/jsonrpc.Dial":           jsonrpc.Dial,
	"net/rpc/jsonrpc.NewClient":      jsonrpc.NewClient,
	"net/rpc/jsonrpc.NewClientCodec": jsonrpc.NewClientCodec,
	"net/rpc/jsonrpc.NewServerCodec": jsonrpc.NewServerCodec,
	"net/rpc/jsonrpc.ServeConn":      jsonrpc.ServeConn,
}

var typList = []interface{}{}
