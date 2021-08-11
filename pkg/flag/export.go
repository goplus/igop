// export by github.com/goplus/interp/cmd/qexp

package flag

import (
	"flag"

	"github.com/goplus/interp"
)

func init() {
	interp.RegisterPackage("flag", extMap, typList)
}

var extMap = map[string]interface{}{
	"(*flag.FlagSet).Arg":           (*flag.FlagSet).Arg,
	"(*flag.FlagSet).Args":          (*flag.FlagSet).Args,
	"(*flag.FlagSet).Bool":          (*flag.FlagSet).Bool,
	"(*flag.FlagSet).BoolVar":       (*flag.FlagSet).BoolVar,
	"(*flag.FlagSet).Duration":      (*flag.FlagSet).Duration,
	"(*flag.FlagSet).DurationVar":   (*flag.FlagSet).DurationVar,
	"(*flag.FlagSet).ErrorHandling": (*flag.FlagSet).ErrorHandling,
	"(*flag.FlagSet).Float64":       (*flag.FlagSet).Float64,
	"(*flag.FlagSet).Float64Var":    (*flag.FlagSet).Float64Var,
	"(*flag.FlagSet).Init":          (*flag.FlagSet).Init,
	"(*flag.FlagSet).Int":           (*flag.FlagSet).Int,
	"(*flag.FlagSet).Int64":         (*flag.FlagSet).Int64,
	"(*flag.FlagSet).Int64Var":      (*flag.FlagSet).Int64Var,
	"(*flag.FlagSet).IntVar":        (*flag.FlagSet).IntVar,
	"(*flag.FlagSet).Lookup":        (*flag.FlagSet).Lookup,
	"(*flag.FlagSet).NArg":          (*flag.FlagSet).NArg,
	"(*flag.FlagSet).NFlag":         (*flag.FlagSet).NFlag,
	"(*flag.FlagSet).Name":          (*flag.FlagSet).Name,
	"(*flag.FlagSet).Output":        (*flag.FlagSet).Output,
	"(*flag.FlagSet).Parse":         (*flag.FlagSet).Parse,
	"(*flag.FlagSet).Parsed":        (*flag.FlagSet).Parsed,
	"(*flag.FlagSet).PrintDefaults": (*flag.FlagSet).PrintDefaults,
	"(*flag.FlagSet).Set":           (*flag.FlagSet).Set,
	"(*flag.FlagSet).SetOutput":     (*flag.FlagSet).SetOutput,
	"(*flag.FlagSet).String":        (*flag.FlagSet).String,
	"(*flag.FlagSet).StringVar":     (*flag.FlagSet).StringVar,
	"(*flag.FlagSet).Uint":          (*flag.FlagSet).Uint,
	"(*flag.FlagSet).Uint64":        (*flag.FlagSet).Uint64,
	"(*flag.FlagSet).Uint64Var":     (*flag.FlagSet).Uint64Var,
	"(*flag.FlagSet).UintVar":       (*flag.FlagSet).UintVar,
	"(*flag.FlagSet).Var":           (*flag.FlagSet).Var,
	"(*flag.FlagSet).Visit":         (*flag.FlagSet).Visit,
	"(*flag.FlagSet).VisitAll":      (*flag.FlagSet).VisitAll,
	"(flag.Getter).Get":             (flag.Getter).Get,
	"(flag.Getter).Set":             (flag.Getter).Set,
	"(flag.Getter).String":          (flag.Getter).String,
	"(flag.Value).Set":              (flag.Value).Set,
	"(flag.Value).String":           (flag.Value).String,
	"flag.Arg":                      flag.Arg,
	"flag.Args":                     flag.Args,
	"flag.Bool":                     flag.Bool,
	"flag.BoolVar":                  flag.BoolVar,
	"flag.CommandLine":              &flag.CommandLine,
	"flag.Duration":                 flag.Duration,
	"flag.DurationVar":              flag.DurationVar,
	"flag.ErrHelp":                  &flag.ErrHelp,
	"flag.Float64":                  flag.Float64,
	"flag.Float64Var":               flag.Float64Var,
	"flag.Int":                      flag.Int,
	"flag.Int64":                    flag.Int64,
	"flag.Int64Var":                 flag.Int64Var,
	"flag.IntVar":                   flag.IntVar,
	"flag.Lookup":                   flag.Lookup,
	"flag.NArg":                     flag.NArg,
	"flag.NFlag":                    flag.NFlag,
	"flag.NewFlagSet":               flag.NewFlagSet,
	"flag.Parse":                    flag.Parse,
	"flag.Parsed":                   flag.Parsed,
	"flag.PrintDefaults":            flag.PrintDefaults,
	"flag.Set":                      flag.Set,
	"flag.String":                   flag.String,
	"flag.StringVar":                flag.StringVar,
	"flag.Uint":                     flag.Uint,
	"flag.Uint64":                   flag.Uint64,
	"flag.Uint64Var":                flag.Uint64Var,
	"flag.UintVar":                  flag.UintVar,
	"flag.UnquoteUsage":             flag.UnquoteUsage,
	"flag.Usage":                    &flag.Usage,
	"flag.Var":                      flag.Var,
	"flag.Visit":                    flag.Visit,
	"flag.VisitAll":                 flag.VisitAll,
}

var typList = []interface{}{
	(*flag.ErrorHandling)(nil),
	(*flag.Flag)(nil),
	(*flag.FlagSet)(nil),
	(*flag.Getter)(nil),
	(*flag.Value)(nil),
}
