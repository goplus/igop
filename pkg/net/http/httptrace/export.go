// export by github.com/goplus/interp/cmd/qexp

package httptrace

import (
	"net/http/httptrace"

	"github.com/goplus/interp"
)

func init() {
	interp.RegisterPackage("net/http/httptrace", extMap, typList)
}

var extMap = map[string]interface{}{
	"net/http/httptrace.ContextClientTrace": httptrace.ContextClientTrace,
	"net/http/httptrace.WithClientTrace":    httptrace.WithClientTrace,
}

var typList = []interface{}{
	(*httptrace.ClientTrace)(nil),
	(*httptrace.DNSDoneInfo)(nil),
	(*httptrace.DNSStartInfo)(nil),
	(*httptrace.GotConnInfo)(nil),
	(*httptrace.WroteRequestInfo)(nil),
}
