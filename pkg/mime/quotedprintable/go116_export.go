// export by github.com/goplus/igop/cmd/qexp

//+build go1.16,!go1.17

package quotedprintable

import (
	q "mime/quotedprintable"

	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "quotedprintable",
		Path: "mime/quotedprintable",
		Deps: map[string]string{
			"bufio": "bufio",
			"bytes": "bytes",
			"fmt":   "fmt",
			"io":    "io",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{
			"Reader": reflect.TypeOf((*q.Reader)(nil)).Elem(),
			"Writer": reflect.TypeOf((*q.Writer)(nil)).Elem(),
		},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"NewReader": reflect.ValueOf(q.NewReader),
			"NewWriter": reflect.ValueOf(q.NewWriter),
		},
		TypedConsts:   map[string]igop.TypedConst{},
		UntypedConsts: map[string]igop.UntypedConst{},
	})
}
