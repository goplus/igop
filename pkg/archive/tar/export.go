// export by github.com/goplus/gossa/cmd/qexp

package tar

import (
	q "archive/tar"

	"go/constant"
	"reflect"

	"github.com/goplus/gossa"
)

func init() {
	gossa.RegisterPackage(&gossa.Package{
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
		NamedTypes: map[string]gossa.NamedType{
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
		TypedConsts: map[string]gossa.TypedConst{
			"FormatGNU":     {reflect.TypeOf(q.FormatGNU), constant.MakeInt64(8)},
			"FormatPAX":     {reflect.TypeOf(q.FormatPAX), constant.MakeInt64(4)},
			"FormatUSTAR":   {reflect.TypeOf(q.FormatUSTAR), constant.MakeInt64(2)},
			"FormatUnknown": {reflect.TypeOf(q.FormatUnknown), constant.MakeInt64(0)},
		},
		UntypedConsts: map[string]gossa.UntypedConst{
			"TypeBlock":         {"untyped rune", constant.MakeInt64(52)},
			"TypeChar":          {"untyped rune", constant.MakeInt64(51)},
			"TypeCont":          {"untyped rune", constant.MakeInt64(55)},
			"TypeDir":           {"untyped rune", constant.MakeInt64(53)},
			"TypeFifo":          {"untyped rune", constant.MakeInt64(54)},
			"TypeGNULongLink":   {"untyped rune", constant.MakeInt64(75)},
			"TypeGNULongName":   {"untyped rune", constant.MakeInt64(76)},
			"TypeGNUSparse":     {"untyped rune", constant.MakeInt64(83)},
			"TypeLink":          {"untyped rune", constant.MakeInt64(49)},
			"TypeReg":           {"untyped rune", constant.MakeInt64(48)},
			"TypeRegA":          {"untyped rune", constant.MakeInt64(0)},
			"TypeSymlink":       {"untyped rune", constant.MakeInt64(50)},
			"TypeXGlobalHeader": {"untyped rune", constant.MakeInt64(103)},
			"TypeXHeader":       {"untyped rune", constant.MakeInt64(120)},
		},
	})
}
