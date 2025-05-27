// export by github.com/goplus/ixgo/cmd/qexp

package stringslice

import (
	q "github.com/qiniu/x/stringslice"

	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage(&ixgo.Package{
		Name: "stringslice",
		Path: "github.com/qiniu/x/stringslice",
		Deps: map[string]string{
			"github.com/qiniu/x/stringutil": "stringutil",
			"strings":                       "strings",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"Capitalize": reflect.ValueOf(q.Capitalize),
			"Repeat":     reflect.ValueOf(q.Repeat),
			"Replace":    reflect.ValueOf(q.Replace),
			"ReplaceAll": reflect.ValueOf(q.ReplaceAll),
			"ToLower":    reflect.ValueOf(q.ToLower),
			"ToTitle":    reflect.ValueOf(q.ToTitle),
			"ToUpper":    reflect.ValueOf(q.ToUpper),
			"Trim":       reflect.ValueOf(q.Trim),
			"TrimLeft":   reflect.ValueOf(q.TrimLeft),
			"TrimPrefix": reflect.ValueOf(q.TrimPrefix),
			"TrimRight":  reflect.ValueOf(q.TrimRight),
			"TrimSpace":  reflect.ValueOf(q.TrimSpace),
			"TrimSuffix": reflect.ValueOf(q.TrimSuffix),
		},
		TypedConsts:   map[string]ixgo.TypedConst{},
		UntypedConsts: map[string]ixgo.UntypedConst{},
	})
}
