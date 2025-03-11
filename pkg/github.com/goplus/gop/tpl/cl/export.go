// export by github.com/goplus/igop/cmd/qexp

package cl

import (
	q "github.com/goplus/gop/tpl/cl"

	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "cl",
		Path: "github.com/goplus/gop/tpl/cl",
		Deps: map[string]string{
			"fmt":                               "fmt",
			"github.com/goplus/gop/tpl/ast":     "ast",
			"github.com/goplus/gop/tpl/matcher": "matcher",
			"github.com/goplus/gop/tpl/token":   "token",
			"github.com/qiniu/x/errors":         "errors",
			"strconv":                           "strconv",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{
			"Result": reflect.TypeOf((*q.Result)(nil)).Elem(),
		},
		AliasTypes: map[string]reflect.Type{},
		Vars: map[string]reflect.Value{
			"ErrNoDocFound": reflect.ValueOf(&q.ErrNoDocFound),
		},
		Funcs: map[string]reflect.Value{
			"New": reflect.ValueOf(q.New),
		},
		TypedConsts:   map[string]igop.TypedConst{},
		UntypedConsts: map[string]igop.UntypedConst{},
	})
}
