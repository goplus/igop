// export by github.com/goplus/igop/cmd/qexp

//go:build go1.20
// +build go1.20

package buildinfo

import (
	q "debug/buildinfo"

	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "buildinfo",
		Path: "debug/buildinfo",
		Deps: map[string]string{
			"bytes":           "bytes",
			"debug/elf":       "elf",
			"debug/macho":     "macho",
			"debug/pe":        "pe",
			"debug/plan9obj":  "plan9obj",
			"encoding/binary": "binary",
			"errors":          "errors",
			"fmt":             "fmt",
			"internal/xcoff":  "xcoff",
			"io":              "io",
			"io/fs":           "fs",
			"os":              "os",
			"runtime/debug":   "debug",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{},
		AliasTypes: map[string]reflect.Type{
			"BuildInfo": reflect.TypeOf((*q.BuildInfo)(nil)).Elem(),
		},
		Vars: map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"Read":     reflect.ValueOf(q.Read),
			"ReadFile": reflect.ValueOf(q.ReadFile),
		},
		TypedConsts:   map[string]igop.TypedConst{},
		UntypedConsts: map[string]igop.UntypedConst{},
	})
}
