// export by github.com/goplus/igop/cmd/qexp

//+build go1.15,!go1.16

package html

import (
	q "html"

	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
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
		TypedConsts:   map[string]igop.TypedConst{},
		UntypedConsts: map[string]igop.UntypedConst{},
	})
}
