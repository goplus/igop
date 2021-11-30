// export by github.com/goplus/gossa/cmd/qexp

package internal

import (
	q "net/http/internal"

	"reflect"

	"github.com/goplus/gossa"
)

func init() {
	gossa.RegisterPackage(&gossa.Package{
		Name: "internal",
		Path: "net/http/internal",
		Deps: map[string]string{
			"bufio":   "bufio",
			"bytes":   "bytes",
			"errors":  "errors",
			"fmt":     "fmt",
			"io":      "io",
			"strings": "strings",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]gossa.NamedType{
			"FlushAfterChunkWriter": {reflect.TypeOf((*q.FlushAfterChunkWriter)(nil)).Elem(), "", ""},
		},
		AliasTypes: map[string]reflect.Type{},
		Vars: map[string]reflect.Value{
			"ErrLineTooLong": reflect.ValueOf(&q.ErrLineTooLong),
			"LocalhostCert":  reflect.ValueOf(&q.LocalhostCert),
			"LocalhostKey":   reflect.ValueOf(&q.LocalhostKey),
		},
		Funcs: map[string]reflect.Value{
			"NewChunkedReader": reflect.ValueOf(q.NewChunkedReader),
			"NewChunkedWriter": reflect.ValueOf(q.NewChunkedWriter),
		},
		TypedConsts:   map[string]gossa.TypedConst{},
		UntypedConsts: map[string]gossa.UntypedConst{},
	})
}
