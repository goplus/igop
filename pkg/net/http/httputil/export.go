// export by github.com/goplus/interp/cmd/qexp

package httputil

import (
	"net/http/httputil"

	"github.com/goplus/interp"
)

func init() {
	interp.RegisterPackage("net/http/httputil", extMap, typList)
}

var extMap = map[string]interface{}{
	"(*net/http/httputil.ClientConn).Close":       (*httputil.ClientConn).Close,
	"(*net/http/httputil.ClientConn).Do":          (*httputil.ClientConn).Do,
	"(*net/http/httputil.ClientConn).Hijack":      (*httputil.ClientConn).Hijack,
	"(*net/http/httputil.ClientConn).Pending":     (*httputil.ClientConn).Pending,
	"(*net/http/httputil.ClientConn).Read":        (*httputil.ClientConn).Read,
	"(*net/http/httputil.ClientConn).Write":       (*httputil.ClientConn).Write,
	"(*net/http/httputil.ReverseProxy).ServeHTTP": (*httputil.ReverseProxy).ServeHTTP,
	"(*net/http/httputil.ServerConn).Close":       (*httputil.ServerConn).Close,
	"(*net/http/httputil.ServerConn).Hijack":      (*httputil.ServerConn).Hijack,
	"(*net/http/httputil.ServerConn).Pending":     (*httputil.ServerConn).Pending,
	"(*net/http/httputil.ServerConn).Read":        (*httputil.ServerConn).Read,
	"(*net/http/httputil.ServerConn).Write":       (*httputil.ServerConn).Write,
	"(net/http/httputil.BufferPool).Get":          (httputil.BufferPool).Get,
	"(net/http/httputil.BufferPool).Put":          (httputil.BufferPool).Put,
	"net/http/httputil.DumpRequest":               httputil.DumpRequest,
	"net/http/httputil.DumpRequestOut":            httputil.DumpRequestOut,
	"net/http/httputil.DumpResponse":              httputil.DumpResponse,
	"net/http/httputil.ErrClosed":                 &httputil.ErrClosed,
	"net/http/httputil.ErrLineTooLong":            &httputil.ErrLineTooLong,
	"net/http/httputil.ErrPersistEOF":             &httputil.ErrPersistEOF,
	"net/http/httputil.ErrPipeline":               &httputil.ErrPipeline,
	"net/http/httputil.NewChunkedReader":          httputil.NewChunkedReader,
	"net/http/httputil.NewChunkedWriter":          httputil.NewChunkedWriter,
	"net/http/httputil.NewClientConn":             httputil.NewClientConn,
	"net/http/httputil.NewProxyClientConn":        httputil.NewProxyClientConn,
	"net/http/httputil.NewServerConn":             httputil.NewServerConn,
	"net/http/httputil.NewSingleHostReverseProxy": httputil.NewSingleHostReverseProxy,
}

var typList = []interface{}{
	(*httputil.BufferPool)(nil),
	(*httputil.ClientConn)(nil),
	(*httputil.ReverseProxy)(nil),
	(*httputil.ServerConn)(nil),
}
