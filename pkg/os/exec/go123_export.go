// export by github.com/goplus/igop/cmd/qexp

//go:build go1.23 && !go1.24
// +build go1.23,!go1.24

package exec

import (
	q "os/exec"

	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "exec",
		Path: "os/exec",
		Deps: map[string]string{
			"bytes":                    "bytes",
			"context":                  "context",
			"errors":                   "errors",
			"internal/godebug":         "godebug",
			"internal/syscall/execenv": "execenv",
			"internal/syscall/unix":    "unix",
			"io":                       "io",
			"io/fs":                    "fs",
			"os":                       "os",
			"path/filepath":            "filepath",
			"runtime":                  "runtime",
			"strconv":                  "strconv",
			"strings":                  "strings",
			"syscall":                  "syscall",
			"time":                     "time",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{
			"Cmd":       reflect.TypeOf((*q.Cmd)(nil)).Elem(),
			"Error":     reflect.TypeOf((*q.Error)(nil)).Elem(),
			"ExitError": reflect.TypeOf((*q.ExitError)(nil)).Elem(),
		},
		AliasTypes: map[string]reflect.Type{},
		Vars: map[string]reflect.Value{
			"ErrDot":       reflect.ValueOf(&q.ErrDot),
			"ErrNotFound":  reflect.ValueOf(&q.ErrNotFound),
			"ErrWaitDelay": reflect.ValueOf(&q.ErrWaitDelay),
		},
		Funcs: map[string]reflect.Value{
			"Command":        reflect.ValueOf(q.Command),
			"CommandContext": reflect.ValueOf(q.CommandContext),
			"LookPath":       reflect.ValueOf(q.LookPath),
		},
		TypedConsts:   map[string]igop.TypedConst{},
		UntypedConsts: map[string]igop.UntypedConst{},
	})
}
