// export by github.com/goplus/gossa/cmd/qexp

package httptrace

import (
	"net/http/httptrace"

	"github.com/goplus/gossa"
)

func init() {
	gossa.RegisterPackage("net/http/httptrace", extMap, typList)
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
