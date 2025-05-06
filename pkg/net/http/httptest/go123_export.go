// export by github.com/goplus/igop/cmd/qexp

//go:build go1.23 && !go1.24
// +build go1.23,!go1.24

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
			"context":                               "context",
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
		NamedTypes: map[string]reflect.Type{
			"ResponseRecorder": reflect.TypeOf((*q.ResponseRecorder)(nil)).Elem(),
			"Server":           reflect.TypeOf((*q.Server)(nil)).Elem(),
		},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"NewRecorder":           reflect.ValueOf(q.NewRecorder),
			"NewRequest":            reflect.ValueOf(q.NewRequest),
			"NewRequestWithContext": reflect.ValueOf(q.NewRequestWithContext),
			"NewServer":             reflect.ValueOf(q.NewServer),
			"NewTLSServer":          reflect.ValueOf(q.NewTLSServer),
			"NewUnstartedServer":    reflect.ValueOf(q.NewUnstartedServer),
		},
		TypedConsts: map[string]igop.TypedConst{},
		UntypedConsts: map[string]igop.UntypedConst{
			"DefaultRemoteAddr": {"untyped string", constant.MakeString(string(q.DefaultRemoteAddr))},
		},
	})
}
