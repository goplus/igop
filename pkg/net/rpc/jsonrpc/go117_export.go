// export by github.com/goplus/gossa/cmd/qexp

//go:build go1.17
// +build go1.17

package jsonrpc

import (
	q "net/rpc/jsonrpc"

	"reflect"

	"github.com/goplus/gossa"
)

func init() {
	gossa.RegisterPackage(&gossa.Package{
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
		NamedTypes: map[string]gossa.NamedType{},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"Dial":           reflect.ValueOf(q.Dial),
			"NewClient":      reflect.ValueOf(q.NewClient),
			"NewClientCodec": reflect.ValueOf(q.NewClientCodec),
			"NewServerCodec": reflect.ValueOf(q.NewServerCodec),
			"ServeConn":      reflect.ValueOf(q.ServeConn),
		},
		TypedConsts:   map[string]gossa.TypedConst{},
		UntypedConsts: map[string]gossa.UntypedConst{},
	})
}
