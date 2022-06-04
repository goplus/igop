// export by github.com/goplus/igop/cmd/qexp

//+build go1.15,!go1.16

package expvar

import (
	q "expvar"

	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "expvar",
		Path: "expvar",
		Deps: map[string]string{
			"encoding/json": "json",
			"fmt":           "fmt",
			"log":           "log",
			"math":          "math",
			"net/http":      "http",
			"os":            "os",
			"runtime":       "runtime",
			"sort":          "sort",
			"strconv":       "strconv",
			"strings":       "strings",
			"sync":          "sync",
			"sync/atomic":   "atomic",
		},
		Interfaces: map[string]reflect.Type{
			"Var": reflect.TypeOf((*q.Var)(nil)).Elem(),
		},
		NamedTypes: map[string]igop.NamedType{
			"Float":    {reflect.TypeOf((*q.Float)(nil)).Elem(), "", "Add,Set,String,Value"},
			"Func":     {reflect.TypeOf((*q.Func)(nil)).Elem(), "String,Value", ""},
			"Int":      {reflect.TypeOf((*q.Int)(nil)).Elem(), "", "Add,Set,String,Value"},
			"KeyValue": {reflect.TypeOf((*q.KeyValue)(nil)).Elem(), "", ""},
			"Map":      {reflect.TypeOf((*q.Map)(nil)).Elem(), "", "Add,AddFloat,Delete,Do,Get,Init,Set,String,addKey"},
			"String":   {reflect.TypeOf((*q.String)(nil)).Elem(), "", "Set,String,Value"},
		},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"Do":        reflect.ValueOf(q.Do),
			"Get":       reflect.ValueOf(q.Get),
			"Handler":   reflect.ValueOf(q.Handler),
			"NewFloat":  reflect.ValueOf(q.NewFloat),
			"NewInt":    reflect.ValueOf(q.NewInt),
			"NewMap":    reflect.ValueOf(q.NewMap),
			"NewString": reflect.ValueOf(q.NewString),
			"Publish":   reflect.ValueOf(q.Publish),
		},
		TypedConsts:   map[string]igop.TypedConst{},
		UntypedConsts: map[string]igop.UntypedConst{},
	})
}
