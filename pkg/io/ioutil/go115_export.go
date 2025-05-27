// export by github.com/goplus/ixgo/cmd/qexp

//+build go1.15,!go1.16

package ioutil

import (
	q "io/ioutil"

	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage(&ixgo.Package{
		Name: "ioutil",
		Path: "io/ioutil",
		Deps: map[string]string{
			"bytes":         "bytes",
			"errors":        "errors",
			"io":            "io",
			"os":            "os",
			"path/filepath": "filepath",
			"sort":          "sort",
			"strconv":       "strconv",
			"strings":       "strings",
			"sync":          "sync",
			"time":          "time",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{},
		AliasTypes: map[string]reflect.Type{},
		Vars: map[string]reflect.Value{
			"Discard": reflect.ValueOf(&q.Discard),
		},
		Funcs: map[string]reflect.Value{
			"NopCloser": reflect.ValueOf(q.NopCloser),
			"ReadAll":   reflect.ValueOf(q.ReadAll),
			"ReadDir":   reflect.ValueOf(q.ReadDir),
			"ReadFile":  reflect.ValueOf(q.ReadFile),
			"TempDir":   reflect.ValueOf(q.TempDir),
			"TempFile":  reflect.ValueOf(q.TempFile),
			"WriteFile": reflect.ValueOf(q.WriteFile),
		},
		TypedConsts:   map[string]ixgo.TypedConst{},
		UntypedConsts: map[string]ixgo.UntypedConst{},
	})
}
