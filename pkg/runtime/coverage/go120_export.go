// export by github.com/goplus/igop/cmd/qexp

//go:build go1.20 && !go1.20
// +build go1.20,!go1.20

package coverage

import (
	q "runtime/coverage"

	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "coverage",
		Path: "runtime/coverage",
		Deps: map[string]string{
			"crypto/md5":                      "md5",
			"fmt":                             "fmt",
			"internal/coverage":               "coverage",
			"internal/coverage/calloc":        "calloc",
			"internal/coverage/cformat":       "cformat",
			"internal/coverage/cmerge":        "cmerge",
			"internal/coverage/decodecounter": "decodecounter",
			"internal/coverage/decodemeta":    "decodemeta",
			"internal/coverage/encodecounter": "encodecounter",
			"internal/coverage/encodemeta":    "encodemeta",
			"internal/coverage/pods":          "pods",
			"internal/coverage/rtcov":         "rtcov",
			"io":                              "io",
			"os":                              "os",
			"path/filepath":                   "filepath",
			"reflect":                         "reflect",
			"runtime":                         "runtime",
			"strings":                         "strings",
			"sync/atomic":                     "atomic",
			"time":                            "time",
			"unsafe":                          "unsafe",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"ClearCounters":    reflect.ValueOf(q.ClearCounters),
			"WriteCounters":    reflect.ValueOf(q.WriteCounters),
			"WriteCountersDir": reflect.ValueOf(q.WriteCountersDir),
			"WriteMeta":        reflect.ValueOf(q.WriteMeta),
			"WriteMetaDir":     reflect.ValueOf(q.WriteMetaDir),
		},
		TypedConsts:   map[string]igop.TypedConst{},
		UntypedConsts: map[string]igop.UntypedConst{},
	})
}
