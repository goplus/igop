// export by github.com/goplus/igop/cmd/qexp

//+build go1.14,!go1.15

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
			"bufio":              "bufio",
			"bytes":              "bytes",
			"errors":             "errors",
			"fmt":                "fmt",
			"go/ast":             "ast",
			"go/doc":             "doc",
			"go/parser":          "parser",
			"go/token":           "token",
			"internal/execabs":   "execabs",
			"internal/goroot":    "goroot",
			"internal/goversion": "goversion",
			"io":                 "io",
			"io/ioutil":          "ioutil",
			"log":                "log",
			"os":                 "os",
			"path":               "path",
			"path/filepath":      "filepath",
			"runtime":            "runtime",
			"sort":               "sort",
			"strconv":            "strconv",
			"strings":            "strings",
			"unicode":            "unicode",
			"unicode/utf8":       "utf8",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]igop.NamedType{
			"Context":              {reflect.TypeOf((*q.Context)(nil)).Elem(), "", "Import,ImportDir,MatchFile,SrcDirs,goodOSArchFile,gopath,hasSubdir,importGo,isAbsPath,isDir,isFile,joinPath,makePathsAbsolute,match,matchFile,openFile,readDir,saveCgo,shouldBuild,splitPathList"},
			"ImportMode":           {reflect.TypeOf((*q.ImportMode)(nil)).Elem(), "", ""},
			"MultiplePackageError": {reflect.TypeOf((*q.MultiplePackageError)(nil)).Elem(), "", "Error"},
			"NoGoError":            {reflect.TypeOf((*q.NoGoError)(nil)).Elem(), "", "Error"},
			"Package":              {reflect.TypeOf((*q.Package)(nil)).Elem(), "", "IsCommand"},
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
