// export by github.com/goplus/igop/cmd/qexp

//+build go1.14,!go1.15

package cookiejar

import (
	q "net/http/cookiejar"

	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "cookiejar",
		Path: "net/http/cookiejar",
		Deps: map[string]string{
			"errors":       "errors",
			"fmt":          "fmt",
			"net":          "net",
			"net/http":     "http",
			"net/url":      "url",
			"sort":         "sort",
			"strings":      "strings",
			"sync":         "sync",
			"time":         "time",
			"unicode/utf8": "utf8",
		},
		Interfaces: map[string]reflect.Type{
			"PublicSuffixList": reflect.TypeOf((*q.PublicSuffixList)(nil)).Elem(),
		},
		NamedTypes: map[string]reflect.Type{
			"Jar":     reflect.TypeOf((*q.Jar)(nil)).Elem(),
			"Options": reflect.TypeOf((*q.Options)(nil)).Elem(),
		},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"New": reflect.ValueOf(q.New),
		},
		TypedConsts:   map[string]igop.TypedConst{},
		UntypedConsts: map[string]igop.UntypedConst{},
	})
}
