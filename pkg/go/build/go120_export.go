// export by github.com/goplus/igop/cmd/qexp

//go:build go1.20
// +build go1.20

package build

import (
	q "go/build"

	"go/constant"
	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "build",
		Path: "go/build",
		Deps: map[string]string{
			"bufio":               "bufio",
			"bytes":               "bytes",
			"errors":              "errors",
			"fmt":                 "fmt",
			"go/ast":              "ast",
			"go/build/constraint": "constraint",
			"go/doc":              "doc",
			"go/parser":           "parser",
			"go/token":            "token",
			"internal/buildcfg":   "buildcfg",
			"internal/godebug":    "godebug",
			"internal/goroot":     "goroot",
			"internal/goversion":  "goversion",
			"io":                  "io",
			"io/fs":               "fs",
			"os":                  "os",
			"os/exec":             "exec",
			"path":                "path",
			"path/filepath":       "filepath",
			"runtime":             "runtime",
			"sort":                "sort",
			"strconv":             "strconv",
			"strings":             "strings",
			"unicode":             "unicode",
			"unicode/utf8":        "utf8",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{
			"Context":              reflect.TypeOf((*q.Context)(nil)).Elem(),
			"ImportMode":           reflect.TypeOf((*q.ImportMode)(nil)).Elem(),
			"MultiplePackageError": reflect.TypeOf((*q.MultiplePackageError)(nil)).Elem(),
			"NoGoError":            reflect.TypeOf((*q.NoGoError)(nil)).Elem(),
			"Package":              reflect.TypeOf((*q.Package)(nil)).Elem(),
		},
		AliasTypes: map[string]reflect.Type{},
		Vars: map[string]reflect.Value{
			"Default": reflect.ValueOf(&q.Default),
			"ToolDir": reflect.ValueOf(&q.ToolDir),
		},
		Funcs: map[string]reflect.Value{
			"ArchChar":      reflect.ValueOf(q.ArchChar),
			"Import":        reflect.ValueOf(q.Import),
			"ImportDir":     reflect.ValueOf(q.ImportDir),
			"IsLocalImport": reflect.ValueOf(q.IsLocalImport),
		},
		TypedConsts: map[string]igop.TypedConst{
			"AllowBinary":   {reflect.TypeOf(q.AllowBinary), constant.MakeInt64(int64(q.AllowBinary))},
			"FindOnly":      {reflect.TypeOf(q.FindOnly), constant.MakeInt64(int64(q.FindOnly))},
			"IgnoreVendor":  {reflect.TypeOf(q.IgnoreVendor), constant.MakeInt64(int64(q.IgnoreVendor))},
			"ImportComment": {reflect.TypeOf(q.ImportComment), constant.MakeInt64(int64(q.ImportComment))},
		},
		UntypedConsts: map[string]igop.UntypedConst{},
	})
}
