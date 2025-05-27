// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.14 && !go1.15 && js
// +build go1.14,!go1.15,js

package js

import (
	q "syscall/js"

	"go/constant"
	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage(&ixgo.Package{
		Name: "js",
		Path: "syscall/js",
		Deps: map[string]string{
			"runtime": "runtime",
			"sync":    "sync",
			"unsafe":  "unsafe",
		},
		Interfaces: map[string]reflect.Type{
			"Wrapper": reflect.TypeOf((*q.Wrapper)(nil)).Elem(),
		},
		NamedTypes: map[string]reflect.Type{
			"Error":      reflect.TypeOf((*q.Error)(nil)).Elem(),
			"Func":       reflect.TypeOf((*q.Func)(nil)).Elem(),
			"Type":       reflect.TypeOf((*q.Type)(nil)).Elem(),
			"Value":      reflect.TypeOf((*q.Value)(nil)).Elem(),
			"ValueError": reflect.TypeOf((*q.ValueError)(nil)).Elem(),
		},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"CopyBytesToGo": reflect.ValueOf(q.CopyBytesToGo),
			"CopyBytesToJS": reflect.ValueOf(q.CopyBytesToJS),
			"FuncOf":        reflect.ValueOf(q.FuncOf),
			"Global":        reflect.ValueOf(q.Global),
			"Null":          reflect.ValueOf(q.Null),
			"Undefined":     reflect.ValueOf(q.Undefined),
			"ValueOf":       reflect.ValueOf(q.ValueOf),
		},
		TypedConsts: map[string]ixgo.TypedConst{
			"TypeBoolean":   {reflect.TypeOf(q.TypeBoolean), constant.MakeInt64(int64(q.TypeBoolean))},
			"TypeFunction":  {reflect.TypeOf(q.TypeFunction), constant.MakeInt64(int64(q.TypeFunction))},
			"TypeNull":      {reflect.TypeOf(q.TypeNull), constant.MakeInt64(int64(q.TypeNull))},
			"TypeNumber":    {reflect.TypeOf(q.TypeNumber), constant.MakeInt64(int64(q.TypeNumber))},
			"TypeObject":    {reflect.TypeOf(q.TypeObject), constant.MakeInt64(int64(q.TypeObject))},
			"TypeString":    {reflect.TypeOf(q.TypeString), constant.MakeInt64(int64(q.TypeString))},
			"TypeSymbol":    {reflect.TypeOf(q.TypeSymbol), constant.MakeInt64(int64(q.TypeSymbol))},
			"TypeUndefined": {reflect.TypeOf(q.TypeUndefined), constant.MakeInt64(int64(q.TypeUndefined))},
		},
		UntypedConsts: map[string]ixgo.UntypedConst{},
	})
}
