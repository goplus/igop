// export by github.com/goplus/igop/cmd/qexp

package gsh

import (
	q "github.com/qiniu/x/gsh"

	"go/constant"
	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "gsh",
		Path: "github.com/qiniu/x/gsh",
		Deps: map[string]string{
			"bytes":   "bytes",
			"errors":  "errors",
			"io":      "io",
			"os":      "os",
			"os/exec": "exec",
			"strings": "strings",
		},
		Interfaces: map[string]reflect.Type{
			"OS": reflect.TypeOf((*q.OS)(nil)).Elem(),
		},
		NamedTypes: map[string]reflect.Type{
			"App": reflect.TypeOf((*q.App)(nil)).Elem(),
		},
		AliasTypes: map[string]reflect.Type{},
		Vars: map[string]reflect.Value{
			"Sys": reflect.ValueOf(&q.Sys),
		},
		Funcs: map[string]reflect.Value{
			"Getenv":        reflect.ValueOf(q.Getenv),
			"Gopt_App_Main": reflect.ValueOf(q.Gopt_App_Main),
			"Setenv__0":     reflect.ValueOf(q.Setenv__0),
			"Setenv__1":     reflect.ValueOf(q.Setenv__1),
			"Setenv__2":     reflect.ValueOf(q.Setenv__2),
		},
		TypedConsts: map[string]igop.TypedConst{},
		UntypedConsts: map[string]igop.UntypedConst{
			"GopPackage": {"untyped bool", constant.MakeBool(bool(q.GopPackage))},
		},
	})
}
