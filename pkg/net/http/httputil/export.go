// export by github.com/goplus/gossa/cmd/qexp

package httputil

import (
	q "net/http/httputil"

	"reflect"

	"github.com/goplus/gossa"
)

func init() {
	gossa.RegisterPackage(&gossa.Package{
		Name: "httputil",
		Path: "net/http/httputil",
		Deps: map[string]string{
			"bufio":                                 "bufio",
			"bytes":                                 "bytes",
			"context":                               "context",
			"errors":                                "errors",
			"fmt":                                   "fmt",
			"io":                                    "io",
			"log":                                   "log",
			"net":                                   "net",
			"net/http":                              "http",
			"net/http/internal":                     "internal",
			"net/textproto":                         "textproto",
			"net/url":                               "url",
			"strings":                               "strings",
			"sync":                                  "sync",
			"time":                                  "time",
			"vendor/golang.org/x/net/http/httpguts": "httpguts",
		},
		Interfaces: map[string]reflect.Type{
			"BufferPool": reflect.TypeOf((*q.BufferPool)(nil)).Elem(),
		},
		NamedTypes: map[string]gossa.NamedType{
			"ClientConn":   {reflect.TypeOf((*q.ClientConn)(nil)).Elem(), "", "Close,Do,Hijack,Pending,Read,Write"},
			"ReverseProxy": {reflect.TypeOf((*q.ReverseProxy)(nil)).Elem(), "", "ServeHTTP,copyBuffer,copyResponse,defaultErrorHandler,flushInterval,getErrorHandler,handleUpgradeResponse,logf,modifyResponse"},
			"ServerConn":   {reflect.TypeOf((*q.ServerConn)(nil)).Elem(), "", "Close,Hijack,Pending,Read,Write"},
		},
		AliasTypes: map[string]reflect.Type{},
		Vars: map[string]reflect.Value{
			"ErrClosed":      reflect.ValueOf(&q.ErrClosed),
			"ErrLineTooLong": reflect.ValueOf(&q.ErrLineTooLong),
			"ErrPersistEOF":  reflect.ValueOf(&q.ErrPersistEOF),
			"ErrPipeline":    reflect.ValueOf(&q.ErrPipeline),
		},
		Funcs: map[string]reflect.Value{
			"DumpRequest":               reflect.ValueOf(q.DumpRequest),
			"DumpRequestOut":            reflect.ValueOf(q.DumpRequestOut),
			"DumpResponse":              reflect.ValueOf(q.DumpResponse),
			"NewChunkedReader":          reflect.ValueOf(q.NewChunkedReader),
			"NewChunkedWriter":          reflect.ValueOf(q.NewChunkedWriter),
			"NewClientConn":             reflect.ValueOf(q.NewClientConn),
			"NewProxyClientConn":        reflect.ValueOf(q.NewProxyClientConn),
			"NewServerConn":             reflect.ValueOf(q.NewServerConn),
			"NewSingleHostReverseProxy": reflect.ValueOf(q.NewSingleHostReverseProxy),
		},
		TypedConsts:   map[string]gossa.TypedConst{},
		UntypedConsts: map[string]gossa.UntypedConst{},
	})
}
