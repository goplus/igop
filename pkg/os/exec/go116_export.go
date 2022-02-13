// export by github.com/goplus/gossa/cmd/qexp

//+build go1.16,!go1.17

package exec

import (
	q "os/exec"

	"reflect"

	"github.com/goplus/gossa"
)

func init() {
	gossa.RegisterPackage(&gossa.Package{
		Name: "exec",
		Path: "os/exec",
		Deps: map[string]string{
			"bytes":                    "bytes",
			"context":                  "context",
			"errors":                   "errors",
			"internal/syscall/execenv": "execenv",
			"io":                       "io",
			"io/fs":                    "fs",
			"os":                       "os",
			"path/filepath":            "filepath",
			"runtime":                  "runtime",
			"strconv":                  "strconv",
			"strings":                  "strings",
			"sync":                     "sync",
			"syscall":                  "syscall",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]gossa.NamedType{
			"Cmd":       {reflect.TypeOf((*q.Cmd)(nil)).Elem(), "", "CombinedOutput,Output,Run,Start,StderrPipe,StdinPipe,StdoutPipe,String,Wait,argv,closeDescriptors,envv,stderr,stdin,stdout,writerDescriptor"},
			"Error":     {reflect.TypeOf((*q.Error)(nil)).Elem(), "", "Error,Unwrap"},
			"ExitError": {reflect.TypeOf((*q.ExitError)(nil)).Elem(), "", "Error"},
		},
		AliasTypes: map[string]reflect.Type{},
		Vars: map[string]reflect.Value{
			"ErrNotFound": reflect.ValueOf(&q.ErrNotFound),
		},
		Funcs: map[string]reflect.Value{
			"Command":        reflect.ValueOf(q.Command),
			"CommandContext": reflect.ValueOf(q.CommandContext),
			"LookPath":       reflect.ValueOf(q.LookPath),
		},
		TypedConsts:   map[string]gossa.TypedConst{},
		UntypedConsts: map[string]gossa.UntypedConst{},
	})
}
