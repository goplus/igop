// export by github.com/goplus/igop/cmd/qexp

//go:build go1.17 && !go1.18
// +build go1.17,!go1.18

package flate

import (
	q "compress/flate"

	"go/constant"
	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "flate",
		Path: "compress/flate",
		Deps: map[string]string{
			"bufio":     "bufio",
			"fmt":       "fmt",
			"io":        "io",
			"math":      "math",
			"math/bits": "bits",
			"sort":      "sort",
			"strconv":   "strconv",
			"sync":      "sync",
		},
		Interfaces: map[string]reflect.Type{
			"Reader":   reflect.TypeOf((*q.Reader)(nil)).Elem(),
			"Resetter": reflect.TypeOf((*q.Resetter)(nil)).Elem(),
		},
		NamedTypes: map[string]reflect.Type{
			"CorruptInputError": reflect.TypeOf((*q.CorruptInputError)(nil)).Elem(),
			"InternalError":     reflect.TypeOf((*q.InternalError)(nil)).Elem(),
			"ReadError":         reflect.TypeOf((*q.ReadError)(nil)).Elem(),
			"WriteError":        reflect.TypeOf((*q.WriteError)(nil)).Elem(),
			"Writer":            reflect.TypeOf((*q.Writer)(nil)).Elem(),
		},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"NewReader":     reflect.ValueOf(q.NewReader),
			"NewReaderDict": reflect.ValueOf(q.NewReaderDict),
			"NewWriter":     reflect.ValueOf(q.NewWriter),
			"NewWriterDict": reflect.ValueOf(q.NewWriterDict),
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
