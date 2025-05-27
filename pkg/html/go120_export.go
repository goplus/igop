// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.20 && !go1.21
// +build go1.20,!go1.21

package html

import (
	q "html"

	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage(&ixgo.Package{
		Name: "html",
		Path: "html",
		Deps: map[string]string{
			"strings":      "strings",
			"sync":         "sync",
			"unicode/utf8": "utf8",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"EscapeString":   reflect.ValueOf(q.EscapeString),
			"UnescapeString": reflect.ValueOf(q.UnescapeString),
		},
		TypedConsts:   map[string]ixgo.TypedConst{},
		UntypedConsts: map[string]ixgo.UntypedConst{},
	})
}
