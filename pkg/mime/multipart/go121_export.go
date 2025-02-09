// export by github.com/goplus/igop/cmd/qexp

//go:build go1.21 && !go1.22
// +build go1.21,!go1.22

package multipart

import (
	q "mime/multipart"

	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "multipart",
		Path: "mime/multipart",
		Deps: map[string]string{
			"bufio":                "bufio",
			"bytes":                "bytes",
			"crypto/rand":          "rand",
			"errors":               "errors",
			"fmt":                  "fmt",
			"internal/godebug":     "godebug",
			"io":                   "io",
			"math":                 "math",
			"mime":                 "mime",
			"mime/quotedprintable": "quotedprintable",
			"net/textproto":        "textproto",
			"os":                   "os",
			"path/filepath":        "filepath",
			"sort":                 "sort",
			"strconv":              "strconv",
			"strings":              "strings",
			"unsafe":               "unsafe",
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
		TypedConsts:   map[string]igop.TypedConst{},
		UntypedConsts: map[string]igop.UntypedConst{},
	})
}
