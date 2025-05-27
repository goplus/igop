// export by github.com/goplus/ixgo/cmd/qexp

//+build go1.14,!go1.15

package multipart

import (
	q "mime/multipart"

	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage(&ixgo.Package{
		Name: "multipart",
		Path: "mime/multipart",
		Deps: map[string]string{
			"bufio":                "bufio",
			"bytes":                "bytes",
			"crypto/rand":          "rand",
			"errors":               "errors",
			"fmt":                  "fmt",
			"io":                   "io",
			"io/ioutil":            "ioutil",
			"mime":                 "mime",
			"mime/quotedprintable": "quotedprintable",
			"net/textproto":        "textproto",
			"os":                   "os",
			"sort":                 "sort",
			"strings":              "strings",
		},
		Interfaces: map[string]reflect.Type{
			"File": reflect.TypeOf((*q.File)(nil)).Elem(),
		},
		NamedTypes: map[string]reflect.Type{
			"FileHeader": reflect.TypeOf((*q.FileHeader)(nil)).Elem(),
			"Form":       reflect.TypeOf((*q.Form)(nil)).Elem(),
			"Part":       reflect.TypeOf((*q.Part)(nil)).Elem(),
			"Reader":     reflect.TypeOf((*q.Reader)(nil)).Elem(),
			"Writer":     reflect.TypeOf((*q.Writer)(nil)).Elem(),
		},
		AliasTypes: map[string]reflect.Type{},
		Vars: map[string]reflect.Value{
			"ErrMessageTooLarge": reflect.ValueOf(&q.ErrMessageTooLarge),
		},
		Funcs: map[string]reflect.Value{
			"NewReader": reflect.ValueOf(q.NewReader),
			"NewWriter": reflect.ValueOf(q.NewWriter),
		},
		TypedConsts:   map[string]ixgo.TypedConst{},
		UntypedConsts: map[string]ixgo.UntypedConst{},
	})
}
