// export by github.com/goplus/igop/cmd/qexp

//go:build go1.18
// +build go1.18

package gzip

import (
	q "compress/gzip"

	"go/constant"
	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "gzip",
		Path: "compress/gzip",
		Deps: map[string]string{
			"bufio":           "bufio",
			"compress/flate":  "flate",
			"encoding/binary": "binary",
			"errors":          "errors",
			"fmt":             "fmt",
			"hash/crc32":      "crc32",
			"io":              "io",
			"time":            "time",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]igop.NamedType{
			"Header": {reflect.TypeOf((*q.Header)(nil)).Elem(), "", ""},
			"Reader": {reflect.TypeOf((*q.Reader)(nil)).Elem(), "", "Close,Multistream,Read,Reset,readHeader,readString"},
			"Writer": {reflect.TypeOf((*q.Writer)(nil)).Elem(), "", "Close,Flush,Reset,Write,init,writeBytes,writeString"},
		},
		AliasTypes: map[string]reflect.Type{},
		Vars: map[string]reflect.Value{
			"ErrChecksum": reflect.ValueOf(&q.ErrChecksum),
			"ErrHeader":   reflect.ValueOf(&q.ErrHeader),
		},
		Funcs: map[string]reflect.Value{
			"NewReader":      reflect.ValueOf(q.NewReader),
			"NewWriter":      reflect.ValueOf(q.NewWriter),
			"NewWriterLevel": reflect.ValueOf(q.NewWriterLevel),
		},
		TypedConsts: map[string]igop.TypedConst{},
		UntypedConsts: map[string]igop.UntypedConst{
			"BestCompression":    {"untyped int", constant.MakeInt64(int64(q.BestCompression))},
			"BestSpeed":          {"untyped int", constant.MakeInt64(int64(q.BestSpeed))},
			"DefaultCompression": {"untyped int", constant.MakeInt64(int64(q.DefaultCompression))},
			"HuffmanOnly":        {"untyped int", constant.MakeInt64(int64(q.HuffmanOnly))},
			"NoCompression":      {"untyped int", constant.MakeInt64(int64(q.NoCompression))},
		},
	})
}
