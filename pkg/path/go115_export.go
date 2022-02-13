// export by github.com/goplus/gossa/cmd/qexp

//+build go1.15,!go1.16

package path

import (
	q "path"

	"reflect"

	"github.com/goplus/gossa"
)

func init() {
	gossa.RegisterPackage(&gossa.Package{
		Name: "path",
		Path: "path",
		Deps: map[string]string{
			"errors":       "errors",
			"strings":      "strings",
			"unicode/utf8": "utf8",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]gossa.NamedType{},
		AliasTypes: map[string]reflect.Type{},
		Vars: map[string]reflect.Value{
			"ErrBadPattern": reflect.ValueOf(&q.ErrBadPattern),
		},
		Funcs: map[string]reflect.Value{
			"Base":  reflect.ValueOf(q.Base),
			"Clean": reflect.ValueOf(q.Clean),
			"Dir":   reflect.ValueOf(q.Dir),
			"Ext":   reflect.ValueOf(q.Ext),
			"IsAbs": reflect.ValueOf(q.IsAbs),
			"Join":  reflect.ValueOf(q.Join),
			"Match": reflect.ValueOf(q.Match),
			"Split": reflect.ValueOf(q.Split),
		},
		TypedConsts:   map[string]gossa.TypedConst{},
		UntypedConsts: map[string]gossa.UntypedConst{},
	})
}
