// export by github.com/goplus/gossa/cmd/qexp

package flag

import (
	q "flag"

	"go/constant"
	"reflect"

	"github.com/goplus/gossa"
)

func init() {
	gossa.RegisterPackage(&gossa.Package{
		Name: "flag",
		Path: "flag",
		Deps: map[string]string{
			"errors":  "errors",
			"fmt":     "fmt",
			"io":      "io",
			"os":      "os",
			"reflect": "reflect",
			"sort":    "sort",
			"strconv": "strconv",
			"strings": "strings",
			"time":    "time",
		},
		Interfaces: map[string]reflect.Type{
			"Getter": reflect.TypeOf((*q.Getter)(nil)).Elem(),
			"Value":  reflect.TypeOf((*q.Value)(nil)).Elem(),
		},
		NamedTypes: map[string]gossa.NamedType{
			"ErrorHandling": {reflect.TypeOf((*q.ErrorHandling)(nil)).Elem(), "", ""},
			"Flag":          {reflect.TypeOf((*q.Flag)(nil)).Elem(), "", ""},
			"FlagSet":       {reflect.TypeOf((*q.FlagSet)(nil)).Elem(), "", "Arg,Args,Bool,BoolVar,Duration,DurationVar,ErrorHandling,Float64,Float64Var,Func,Init,Int,Int64,Int64Var,IntVar,Lookup,NArg,NFlag,Name,Output,Parse,Parsed,PrintDefaults,Set,SetOutput,String,StringVar,Uint,Uint64,Uint64Var,UintVar,Var,Visit,VisitAll,defaultUsage,failf,parseOne,usage"},
		},
		AliasTypes: map[string]reflect.Type{},
		Vars: map[string]reflect.Value{
			"CommandLine": reflect.ValueOf(&q.CommandLine),
			"ErrHelp":     reflect.ValueOf(&q.ErrHelp),
			"Usage":       reflect.ValueOf(&q.Usage),
		},
		Funcs: map[string]reflect.Value{
			"Arg":           reflect.ValueOf(q.Arg),
			"Args":          reflect.ValueOf(q.Args),
			"Bool":          reflect.ValueOf(q.Bool),
			"BoolVar":       reflect.ValueOf(q.BoolVar),
			"Duration":      reflect.ValueOf(q.Duration),
			"DurationVar":   reflect.ValueOf(q.DurationVar),
			"Float64":       reflect.ValueOf(q.Float64),
			"Float64Var":    reflect.ValueOf(q.Float64Var),
			"Func":          reflect.ValueOf(q.Func),
			"Int":           reflect.ValueOf(q.Int),
			"Int64":         reflect.ValueOf(q.Int64),
			"Int64Var":      reflect.ValueOf(q.Int64Var),
			"IntVar":        reflect.ValueOf(q.IntVar),
			"Lookup":        reflect.ValueOf(q.Lookup),
			"NArg":          reflect.ValueOf(q.NArg),
			"NFlag":         reflect.ValueOf(q.NFlag),
			"NewFlagSet":    reflect.ValueOf(q.NewFlagSet),
			"Parse":         reflect.ValueOf(q.Parse),
			"Parsed":        reflect.ValueOf(q.Parsed),
			"PrintDefaults": reflect.ValueOf(q.PrintDefaults),
			"Set":           reflect.ValueOf(q.Set),
			"String":        reflect.ValueOf(q.String),
			"StringVar":     reflect.ValueOf(q.StringVar),
			"Uint":          reflect.ValueOf(q.Uint),
			"Uint64":        reflect.ValueOf(q.Uint64),
			"Uint64Var":     reflect.ValueOf(q.Uint64Var),
			"UintVar":       reflect.ValueOf(q.UintVar),
			"UnquoteUsage":  reflect.ValueOf(q.UnquoteUsage),
			"Var":           reflect.ValueOf(q.Var),
			"Visit":         reflect.ValueOf(q.Visit),
			"VisitAll":      reflect.ValueOf(q.VisitAll),
		},
		TypedConsts: map[string]gossa.TypedConst{
			"ContinueOnError": {reflect.TypeOf(q.ContinueOnError), constant.MakeInt64(0)},
			"ExitOnError":     {reflect.TypeOf(q.ExitOnError), constant.MakeInt64(1)},
			"PanicOnError":    {reflect.TypeOf(q.PanicOnError), constant.MakeInt64(2)},
		},
		UntypedConsts: map[string]gossa.UntypedConst{},
	})
}
