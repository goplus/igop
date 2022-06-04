// export by github.com/goplus/igop/cmd/qexp

//+build go1.16,!go1.17

package jsonrpc

import (
	q "net/rpc/jsonrpc"

	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "jsonrpc",
		Path: "net/rpc/jsonrpc",
		Deps: map[string]string{
			"encoding/json": "json",
			"errors":        "errors",
			"fmt":           "fmt",
			"io":            "io",
			"net":           "net",
			"net/rpc":       "rpc",
			"sync":          "sync",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]igop.NamedType{},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"Dial":           reflect.ValueOf(q.Dial),
			"NewClient":      reflect.ValueOf(q.NewClient),
			"NewClientCodec": reflect.ValueOf(q.NewClientCodec),
			"NewServerCodec": reflect.ValueOf(q.NewServerCodec),
			"ServeConn":      reflect.ValueOf(q.ServeConn),
		},
		TypedConsts:   map[string]igop.TypedConst{},
		UntypedConsts: map[string]igop.UntypedConst{},
	})
}
