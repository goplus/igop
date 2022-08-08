// export by github.com/goplus/igop/cmd/qexp

package iox

import (
	q "github.com/goplus/gop/builtin/iox"

	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "iox",
		Path: "github.com/goplus/gop/builtin/iox",
		Deps: map[string]string{
			"bufio": "bufio",
			"io":    "io",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{
			"BLineIter":   reflect.TypeOf((*q.BLineIter)(nil)).Elem(),
			"BLineReader": reflect.TypeOf((*q.BLineReader)(nil)).Elem(),
			"LineIter":    reflect.TypeOf((*q.LineIter)(nil)).Elem(),
			"LineReader":  reflect.TypeOf((*q.LineReader)(nil)).Elem(),
		},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"BLines":     reflect.ValueOf(q.BLines),
			"EnumBLines": reflect.ValueOf(q.EnumBLines),
			"EnumLines":  reflect.ValueOf(q.EnumLines),
			"Lines":      reflect.ValueOf(q.Lines),
		},
		TypedConsts:   map[string]igop.TypedConst{},
		UntypedConsts: map[string]igop.UntypedConst{},
	})
}
