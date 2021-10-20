// export by github.com/goplus/gossa/cmd/qexp

package httptest

import (
	"net/http/httptest"

	"github.com/goplus/gossa"
)

func init() {
	gossa.RegisterPackage("net/http/httptest", extMap, typList)
}

var extMap = map[string]interface{}{
	"(*net/http/httptest.ResponseRecorder).Flush":        (*httptest.ResponseRecorder).Flush,
	"(*net/http/httptest.ResponseRecorder).Header":       (*httptest.ResponseRecorder).Header,
	"(*net/http/httptest.ResponseRecorder).Result":       (*httptest.ResponseRecorder).Result,
	"(*net/http/httptest.ResponseRecorder).Write":        (*httptest.ResponseRecorder).Write,
	"(*net/http/httptest.ResponseRecorder).WriteHeader":  (*httptest.ResponseRecorder).WriteHeader,
	"(*net/http/httptest.ResponseRecorder).WriteString":  (*httptest.ResponseRecorder).WriteString,
	"(*net/http/httptest.Server).Certificate":            (*httptest.Server).Certificate,
	"(*net/http/httptest.Server).Client":                 (*httptest.Server).Client,
	"(*net/http/httptest.Server).Close":                  (*httptest.Server).Close,
	"(*net/http/httptest.Server).CloseClientConnections": (*httptest.Server).CloseClientConnections,
	"(*net/http/httptest.Server).Start":                  (*httptest.Server).Start,
	"(*net/http/httptest.Server).StartTLS":               (*httptest.Server).StartTLS,
	"net/http/httptest.NewRecorder":                      httptest.NewRecorder,
	"net/http/httptest.NewRequest":                       httptest.NewRequest,
	"net/http/httptest.NewServer":                        httptest.NewServer,
	"net/http/httptest.NewTLSServer":                     httptest.NewTLSServer,
	"net/http/httptest.NewUnstartedServer":               httptest.NewUnstartedServer,
}

var typList = []interface{}{
	(*httptest.ResponseRecorder)(nil),
	(*httptest.Server)(nil),
}
