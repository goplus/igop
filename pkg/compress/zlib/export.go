// export by github.com/goplus/gossa/cmd/qexp

package zlib

import (
	q "compress/zlib"

	"go/constant"
	"reflect"

	"github.com/goplus/gossa"
)

func init() {
	gossa.RegisterPackage(&gossa.Package{
		Name: "zlib",
		Path: "compress/zlib",
		Deps: map[string]string{
			"bufio":           "bufio",
			"compress/flate":  "flate",
			"encoding/binary": "binary",
			"errors":          "errors",
			"fmt":             "fmt",
			"hash":            "hash",
			"hash/adler32":    "adler32",
			"io":              "io",
		},
		Interfaces: map[string]reflect.Type{
			"Resetter": reflect.TypeOf((*q.Resetter)(nil)).Elem(),
		},
		NamedTypes: map[string]gossa.NamedType{
			"Writer": {reflect.TypeOf((*q.Writer)(nil)).Elem(), "", "Close,Flush,Reset,Write,writeHeader"},
		},
		AliasTypes: map[string]reflect.Type{},
		Vars: map[string]reflect.Value{
			"ErrChecksum":   reflect.ValueOf(&q.ErrChecksum),
			"ErrDictionary": reflect.ValueOf(&q.ErrDictionary),
			"ErrHeader":     reflect.ValueOf(&q.ErrHeader),
		},
		Funcs: map[string]reflect.Value{
			"NewReader":          reflect.ValueOf(q.NewReader),
			"NewReaderDict":      reflect.ValueOf(q.NewReaderDict),
			"NewWriter":          reflect.ValueOf(q.NewWriter),
			"NewWriterLevel":     reflect.ValueOf(q.NewWriterLevel),
			"NewWriterLevelDict": reflect.ValueOf(q.NewWriterLevelDict),
		},
		TypedConsts: map[string]gossa.TypedConst{},
		UntypedConsts: map[string]gossa.UntypedConst{
			"BestCompression":    {"untyped int", constant.MakeInt64(9)},
			"BestSpeed":          {"untyped int", constant.MakeInt64(1)},
			"DefaultCompression": {"untyped int", constant.MakeInt64(-1)},
			"HuffmanOnly":        {"untyped int", constant.MakeInt64(-2)},
			"NoCompression":      {"untyped int", constant.MakeInt64(0)},
		},
	})
}
