// export by github.com/goplus/igop/cmd/qexp

//go:build go1.20
// +build go1.20

package regexp

import (
	q "regexp"

	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "regexp",
		Path: "regexp",
		Deps: map[string]string{
			"bytes":         "bytes",
			"io":            "io",
			"regexp/syntax": "syntax",
			"sort":          "sort",
			"strconv":       "strconv",
			"strings":       "strings",
			"sync":          "sync",
			"unicode":       "unicode",
			"unicode/utf8":  "utf8",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{
			"Regexp": reflect.TypeOf((*q.Regexp)(nil)).Elem(),
		},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"Compile":          reflect.ValueOf(q.Compile),
			"CompilePOSIX":     reflect.ValueOf(q.CompilePOSIX),
			"Match":            reflect.ValueOf(q.Match),
			"MatchReader":      reflect.ValueOf(q.MatchReader),
			"MatchString":      reflect.ValueOf(q.MatchString),
			"MustCompile":      reflect.ValueOf(q.MustCompile),
			"MustCompilePOSIX": reflect.ValueOf(q.MustCompilePOSIX),
			"QuoteMeta":        reflect.ValueOf(q.QuoteMeta),
		},
		TypedConsts:   map[string]igop.TypedConst{},
		UntypedConsts: map[string]igop.UntypedConst{},
	})
}
