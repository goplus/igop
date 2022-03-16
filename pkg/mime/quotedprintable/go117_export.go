// export by github.com/goplus/gossa/cmd/qexp

//go:build go1.17 && !go1.18
// +build go1.17,!go1.18

package quotedprintable

import (
	q "mime/quotedprintable"

	"reflect"

	"github.com/goplus/gossa"
)

func init() {
	gossa.RegisterPackage(&gossa.Package{
		Name: "quotedprintable",
		Path: "mime/quotedprintable",
		Deps: map[string]string{
			"bufio": "bufio",
			"bytes": "bytes",
			"fmt":   "fmt",
			"io":    "io",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]gossa.NamedType{
			"Reader": {reflect.TypeOf((*q.Reader)(nil)).Elem(), "", "Read"},
			"Writer": {reflect.TypeOf((*q.Writer)(nil)).Elem(), "", "Close,Write,checkLastByte,encode,flush,insertCRLF,insertSoftLineBreak,write"},
		},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"NewReader": reflect.ValueOf(q.NewReader),
			"NewWriter": reflect.ValueOf(q.NewWriter),
		},
		TypedConsts:   map[string]gossa.TypedConst{},
		UntypedConsts: map[string]gossa.UntypedConst{},
	})
}
