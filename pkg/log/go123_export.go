// export by github.com/goplus/igop/cmd/qexp

//go:build go1.23
// +build go1.23

package log

import (
	q "log"

	"go/constant"
	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "log",
		Path: "log",
		Deps: map[string]string{
			"fmt":          "fmt",
			"io":           "io",
			"log/internal": "internal",
			"os":           "os",
			"runtime":      "runtime",
			"sync":         "sync",
			"sync/atomic":  "atomic",
			"time":         "time",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{
			"Logger": reflect.TypeOf((*q.Logger)(nil)).Elem(),
		},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"Default":   reflect.ValueOf(q.Default),
			"Fatal":     reflect.ValueOf(q.Fatal),
			"Fatalf":    reflect.ValueOf(q.Fatalf),
			"Fatalln":   reflect.ValueOf(q.Fatalln),
			"Flags":     reflect.ValueOf(q.Flags),
			"New":       reflect.ValueOf(q.New),
			"Output":    reflect.ValueOf(q.Output),
			"Panic":     reflect.ValueOf(q.Panic),
			"Panicf":    reflect.ValueOf(q.Panicf),
			"Panicln":   reflect.ValueOf(q.Panicln),
			"Prefix":    reflect.ValueOf(q.Prefix),
			"Print":     reflect.ValueOf(q.Print),
			"Printf":    reflect.ValueOf(q.Printf),
			"Println":   reflect.ValueOf(q.Println),
			"SetFlags":  reflect.ValueOf(q.SetFlags),
			"SetOutput": reflect.ValueOf(q.SetOutput),
			"SetPrefix": reflect.ValueOf(q.SetPrefix),
			"Writer":    reflect.ValueOf(q.Writer),
		},
		TypedConsts: map[string]igop.TypedConst{},
		UntypedConsts: map[string]igop.UntypedConst{
			"LUTC":          {"untyped int", constant.MakeInt64(int64(q.LUTC))},
			"Ldate":         {"untyped int", constant.MakeInt64(int64(q.Ldate))},
			"Llongfile":     {"untyped int", constant.MakeInt64(int64(q.Llongfile))},
			"Lmicroseconds": {"untyped int", constant.MakeInt64(int64(q.Lmicroseconds))},
			"Lmsgprefix":    {"untyped int", constant.MakeInt64(int64(q.Lmsgprefix))},
			"Lshortfile":    {"untyped int", constant.MakeInt64(int64(q.Lshortfile))},
			"LstdFlags":     {"untyped int", constant.MakeInt64(int64(q.LstdFlags))},
			"Ltime":         {"untyped int", constant.MakeInt64(int64(q.Ltime))},
		},
	})
}
