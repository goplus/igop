// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.22 && !go1.23
// +build go1.22,!go1.23

package coverage

import (
	q "runtime/coverage"

	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage(&ixgo.Package{
		Name: "coverage",
		Path: "runtime/coverage",
		Deps: map[string]string{
			"crypto/md5":                      "md5",
			"encoding/json":                   "json",
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
			"runtime":                         "runtime",
			"runtime/internal/atomic":         "atomic",
			"strconv":                         "strconv",
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
		TypedConsts:   map[string]ixgo.TypedConst{},
		UntypedConsts: map[string]ixgo.UntypedConst{},
	})
}
