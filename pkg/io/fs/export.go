// export by github.com/goplus/gossa/cmd/qexp

package fs

import (
	q "io/fs"

	"go/constant"
	"reflect"

	"github.com/goplus/gossa"
)

func init() {
	gossa.RegisterPackage(&gossa.Package{
		Name: "fs",
		Path: "io/fs",
		Deps: map[string]string{
			"errors":           "errors",
			"internal/oserror": "oserror",
			"io":               "io",
			"path":             "path",
			"sort":             "sort",
			"time":             "time",
			"unicode/utf8":     "utf8",
		},
		Interfaces: map[string]reflect.Type{
			"DirEntry":    reflect.TypeOf((*q.DirEntry)(nil)).Elem(),
			"FS":          reflect.TypeOf((*q.FS)(nil)).Elem(),
			"File":        reflect.TypeOf((*q.File)(nil)).Elem(),
			"FileInfo":    reflect.TypeOf((*q.FileInfo)(nil)).Elem(),
			"GlobFS":      reflect.TypeOf((*q.GlobFS)(nil)).Elem(),
			"ReadDirFS":   reflect.TypeOf((*q.ReadDirFS)(nil)).Elem(),
			"ReadDirFile": reflect.TypeOf((*q.ReadDirFile)(nil)).Elem(),
			"ReadFileFS":  reflect.TypeOf((*q.ReadFileFS)(nil)).Elem(),
			"StatFS":      reflect.TypeOf((*q.StatFS)(nil)).Elem(),
			"SubFS":       reflect.TypeOf((*q.SubFS)(nil)).Elem(),
		},
		NamedTypes: map[string]gossa.NamedType{
			"FileMode":    {reflect.TypeOf((*q.FileMode)(nil)).Elem(), "IsDir,IsRegular,Perm,String,Type", ""},
			"PathError":   {reflect.TypeOf((*q.PathError)(nil)).Elem(), "", "Error,Timeout,Unwrap"},
			"WalkDirFunc": {reflect.TypeOf((*q.WalkDirFunc)(nil)).Elem(), "", ""},
		},
		AliasTypes: map[string]reflect.Type{},
		Vars: map[string]reflect.Value{
			"ErrClosed":     reflect.ValueOf(&q.ErrClosed),
			"ErrExist":      reflect.ValueOf(&q.ErrExist),
			"ErrInvalid":    reflect.ValueOf(&q.ErrInvalid),
			"ErrNotExist":   reflect.ValueOf(&q.ErrNotExist),
			"ErrPermission": reflect.ValueOf(&q.ErrPermission),
			"SkipDir":       reflect.ValueOf(&q.SkipDir),
		},
		Funcs: map[string]reflect.Value{
			"Glob":      reflect.ValueOf(q.Glob),
			"ReadDir":   reflect.ValueOf(q.ReadDir),
			"ReadFile":  reflect.ValueOf(q.ReadFile),
			"Stat":      reflect.ValueOf(q.Stat),
			"Sub":       reflect.ValueOf(q.Sub),
			"ValidPath": reflect.ValueOf(q.ValidPath),
			"WalkDir":   reflect.ValueOf(q.WalkDir),
		},
		TypedConsts: map[string]gossa.TypedConst{
			"ModeAppend":     {reflect.TypeOf(q.ModeAppend), constant.MakeInt64(1073741824)},
			"ModeCharDevice": {reflect.TypeOf(q.ModeCharDevice), constant.MakeInt64(2097152)},
			"ModeDevice":     {reflect.TypeOf(q.ModeDevice), constant.MakeInt64(67108864)},
			"ModeDir":        {reflect.TypeOf(q.ModeDir), constant.MakeInt64(2147483648)},
			"ModeExclusive":  {reflect.TypeOf(q.ModeExclusive), constant.MakeInt64(536870912)},
			"ModeIrregular":  {reflect.TypeOf(q.ModeIrregular), constant.MakeInt64(524288)},
			"ModeNamedPipe":  {reflect.TypeOf(q.ModeNamedPipe), constant.MakeInt64(33554432)},
			"ModePerm":       {reflect.TypeOf(q.ModePerm), constant.MakeInt64(511)},
			"ModeSetgid":     {reflect.TypeOf(q.ModeSetgid), constant.MakeInt64(4194304)},
			"ModeSetuid":     {reflect.TypeOf(q.ModeSetuid), constant.MakeInt64(8388608)},
			"ModeSocket":     {reflect.TypeOf(q.ModeSocket), constant.MakeInt64(16777216)},
			"ModeSticky":     {reflect.TypeOf(q.ModeSticky), constant.MakeInt64(1048576)},
			"ModeSymlink":    {reflect.TypeOf(q.ModeSymlink), constant.MakeInt64(134217728)},
			"ModeTemporary":  {reflect.TypeOf(q.ModeTemporary), constant.MakeInt64(268435456)},
			"ModeType":       {reflect.TypeOf(q.ModeType), constant.MakeInt64(2401763328)},
		},
		UntypedConsts: map[string]gossa.UntypedConst{},
	})
}
