// export by github.com/goplus/igop/cmd/qexp

//go:build go1.17 && !go1.18
// +build go1.17,!go1.18

package httptest

import (
	q "net/http/httptest"

	"go/constant"
	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "httptest",
		Path: "net/http/httptest",
		Deps: map[string]string{
			"bufio":                                 "bufio",
			"bytes":                                 "bytes",
			"crypto/tls":                            "tls",
			"crypto/x509":                           "x509",
			"flag":                                  "flag",
			"fmt":                                   "fmt",
			"io":                                    "io",
			"log":                                   "log",
			"net":                                   "net",
			"net/http":                              "http",
			"net/http/internal/testcert":            "testcert",
			"net/textproto":                         "textproto",
			"os":                                    "os",
			"strconv":                               "strconv",
			"strings":                               "strings",
			"sync":                                  "sync",
			"time":                                  "time",
			"vendor/golang.org/x/net/http/httpguts": "httpguts",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]igop.NamedType{
			"ResponseRecorder": {reflect.TypeOf((*q.ResponseRecorder)(nil)).Elem(), "", "Flush,Header,Result,Write,WriteHeader,WriteString,writeHeader"},
			"Server":           {reflect.TypeOf((*q.Server)(nil)).Elem(), "", "Certificate,Client,Close,CloseClientConnections,Start,StartTLS,closeConn,closeConnChan,forgetConn,goServe,logCloseHangDebugInfo,wrap"},
		},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"NewRecorder":        reflect.ValueOf(q.NewRecorder),
			"NewRequest":         reflect.ValueOf(q.NewRequest),
			"NewServer":          reflect.ValueOf(q.NewServer),
			"NewTLSServer":       reflect.ValueOf(q.NewTLSServer),
			"NewUnstartedServer": reflect.ValueOf(q.NewUnstartedServer),
		},
		TypedConsts: map[string]igop.TypedConst{},
		UntypedConsts: map[string]igop.UntypedConst{
			"DefaultRemoteAddr": {"untyped string", constant.MakeString(string(q.DefaultRemoteAddr))},
		},
	})
}
