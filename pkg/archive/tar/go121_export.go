// export by github.com/goplus/igop/cmd/qexp

//go:build go1.21 && !go1.22
// +build go1.21,!go1.22

package tar

import (
	q "archive/tar"

	"go/constant"
	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "tar",
		Path: "archive/tar",
		Deps: map[string]string{
			"bytes":            "bytes",
			"errors":           "errors",
			"fmt":              "fmt",
			"internal/godebug": "godebug",
			"io":               "io",
			"io/fs":            "fs",
			"math":             "math",
			"os/user":          "user",
			"path":             "path",
			"path/filepath":    "filepath",
			"reflect":          "reflect",
			"runtime":          "runtime",
			"sort":             "sort",
			"strconv":          "strconv",
			"strings":          "strings",
			"sync":             "sync",
			"syscall":          "syscall",
			"time":             "time",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{
			"Format": reflect.TypeOf((*q.Format)(nil)).Elem(),
			"Header": reflect.TypeOf((*q.Header)(nil)).Elem(),
			"Reader": reflect.TypeOf((*q.Reader)(nil)).Elem(),
			"Writer": reflect.TypeOf((*q.Writer)(nil)).Elem(),
		},
		AliasTypes: map[string]reflect.Type{},
		Vars: map[string]reflect.Value{
			"ErrFieldTooLong":    reflect.ValueOf(&q.ErrFieldTooLong),
			"ErrHeader":          reflect.ValueOf(&q.ErrHeader),
			"ErrInsecurePath":    reflect.ValueOf(&q.ErrInsecurePath),
			"ErrWriteAfterClose": reflect.ValueOf(&q.ErrWriteAfterClose),
			"ErrWriteTooLong":    reflect.ValueOf(&q.ErrWriteTooLong),
		},
		Funcs: map[string]reflect.Value{
			"FileInfoHeader": reflect.ValueOf(q.FileInfoHeader),
			"NewReader":      reflect.ValueOf(q.NewReader),
			"NewWriter":      reflect.ValueOf(q.NewWriter),
		},
		TypedConsts: map[string]igop.TypedConst{
			"FormatGNU":     {reflect.TypeOf(q.FormatGNU), constant.MakeInt64(int64(q.FormatGNU))},
			"FormatPAX":     {reflect.TypeOf(q.FormatPAX), constant.MakeInt64(int64(q.FormatPAX))},
			"FormatUSTAR":   {reflect.TypeOf(q.FormatUSTAR), constant.MakeInt64(int64(q.FormatUSTAR))},
			"FormatUnknown": {reflect.TypeOf(q.FormatUnknown), constant.MakeInt64(int64(q.FormatUnknown))},
		},
		UntypedConsts: map[string]igop.UntypedConst{
			"TypeBlock":         {"untyped rune", constant.MakeInt64(int64(q.TypeBlock))},
			"TypeChar":          {"untyped rune", constant.MakeInt64(int64(q.TypeChar))},
			"TypeCont":          {"untyped rune", constant.MakeInt64(int64(q.TypeCont))},
			"TypeDir":           {"untyped rune", constant.MakeInt64(int64(q.TypeDir))},
			"TypeFifo":          {"untyped rune", constant.MakeInt64(int64(q.TypeFifo))},
			"TypeGNULongLink":   {"untyped rune", constant.MakeInt64(int64(q.TypeGNULongLink))},
			"TypeGNULongName":   {"untyped rune", constant.MakeInt64(int64(q.TypeGNULongName))},
			"TypeGNUSparse":     {"untyped rune", constant.MakeInt64(int64(q.TypeGNUSparse))},
			"TypeLink":          {"untyped rune", constant.MakeInt64(int64(q.TypeLink))},
			"TypeReg":           {"untyped rune", constant.MakeInt64(int64(q.TypeReg))},
			"TypeRegA":          {"untyped rune", constant.MakeInt64(int64(q.TypeRegA))},
			"TypeSymlink":       {"untyped rune", constant.MakeInt64(int64(q.TypeSymlink))},
			"TypeXGlobalHeader": {"untyped rune", constant.MakeInt64(int64(q.TypeXGlobalHeader))},
			"TypeXHeader":       {"untyped rune", constant.MakeInt64(int64(q.TypeXHeader))},
		},
	})
}
