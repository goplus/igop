// export by github.com/goplus/igop/cmd/qexp

//+build go1.15,!go1.16

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
		NamedTypes: map[string]igop.NamedType{
			"FileHeader": {reflect.TypeOf((*q.FileHeader)(nil)).Elem(), "", "Open"},
			"Form":       {reflect.TypeOf((*q.Form)(nil)).Elem(), "", "RemoveAll"},
			"Part":       {reflect.TypeOf((*q.Part)(nil)).Elem(), "", "Close,FileName,FormName,Read,parseContentDisposition,populateHeaders"},
			"Reader":     {reflect.TypeOf((*q.Reader)(nil)).Elem(), "", "NextPart,NextRawPart,ReadForm,isBoundaryDelimiterLine,isFinalBoundary,nextPart,readForm"},
			"Writer":     {reflect.TypeOf((*q.Writer)(nil)).Elem(), "", "Boundary,Close,CreateFormField,CreateFormFile,CreatePart,FormDataContentType,SetBoundary,WriteField"},
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
