// export by github.com/goplus/gossa/cmd/qexp

//go:build go1.18
// +build go1.18

package buildinfo

import (
	q "debug/buildinfo"

	"reflect"

	"github.com/goplus/gossa"
)

func init() {
	gossa.RegisterPackage(&gossa.Package{
		Name: "buildinfo",
		Path: "debug/buildinfo",
		Deps: map[string]string{
			"bytes":           "bytes",
			"debug/elf":       "elf",
			"debug/macho":     "macho",
			"debug/pe":        "pe",
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
		NamedTypes: map[string]gossa.NamedType{},
		AliasTypes: map[string]reflect.Type{
			"BuildInfo": reflect.TypeOf((*q.BuildInfo)(nil)).Elem(),
		},
		Vars: map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"Read":     reflect.ValueOf(q.Read),
			"ReadFile": reflect.ValueOf(q.ReadFile),
		},
		TypedConsts:   map[string]gossa.TypedConst{},
		UntypedConsts: map[string]gossa.UntypedConst{},
	})
}
