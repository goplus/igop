// export by github.com/goplus/gossa/cmd/qexp

package flate

import (
	q "compress/flate"

	"go/constant"
	"reflect"

	"github.com/goplus/gossa"
)

func init() {
	gossa.RegisterPackage(&gossa.Package{
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
		NamedTypes: map[string]gossa.NamedType{
			"CorruptInputError": {reflect.TypeOf((*q.CorruptInputError)(nil)).Elem(), "Error", ""},
			"InternalError":     {reflect.TypeOf((*q.InternalError)(nil)).Elem(), "Error", ""},
			"ReadError":         {reflect.TypeOf((*q.ReadError)(nil)).Elem(), "", "Error"},
			"WriteError":        {reflect.TypeOf((*q.WriteError)(nil)).Elem(), "", "Error"},
			"Writer":            {reflect.TypeOf((*q.Writer)(nil)).Elem(), "", "Close,Flush,Reset,Write"},
		},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"NewReader":     reflect.ValueOf(q.NewReader),
			"NewReaderDict": reflect.ValueOf(q.NewReaderDict),
			"NewWriter":     reflect.ValueOf(q.NewWriter),
			"NewWriterDict": reflect.ValueOf(q.NewWriterDict),
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
