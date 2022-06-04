// export by github.com/goplus/igop/cmd/qexp

//go:build go1.17 && !go1.18
// +build go1.17,!go1.18

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
			"bytes":   "bytes",
			"errors":  "errors",
			"fmt":     "fmt",
			"io":      "io",
			"io/fs":   "fs",
			"math":    "math",
			"os/user": "user",
			"path":    "path",
			"reflect": "reflect",
			"runtime": "runtime",
			"sort":    "sort",
			"strconv": "strconv",
			"strings": "strings",
			"sync":    "sync",
			"syscall": "syscall",
			"time":    "time",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]igop.NamedType{
			"Format": {reflect.TypeOf((*q.Format)(nil)).Elem(), "String,has", "mayBe,mayOnlyBe,mustNotBe"},
			"Header": {reflect.TypeOf((*q.Header)(nil)).Elem(), "allowedFormats", "FileInfo"},
			"Reader": {reflect.TypeOf((*q.Reader)(nil)).Elem(), "", "Next,Read,handleRegularFile,handleSparseFile,next,readGNUSparsePAXHeaders,readHeader,readOldGNUSparseMap,writeTo"},
			"Writer": {reflect.TypeOf((*q.Writer)(nil)).Elem(), "", "Close,Flush,Write,WriteHeader,readFrom,templateV7Plus,writeGNUHeader,writePAXHeader,writeRawFile,writeRawHeader,writeUSTARHeader"},
		},
		AliasTypes: map[string]reflect.Type{},
		Vars: map[string]reflect.Value{
			"ErrFieldTooLong":    reflect.ValueOf(&q.ErrFieldTooLong),
			"ErrHeader":          reflect.ValueOf(&q.ErrHeader),
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
