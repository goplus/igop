// export by github.com/goplus/gossa/cmd/qexp

package crc32

import (
	q "hash/crc32"

	"go/constant"
	"reflect"

	"github.com/goplus/gossa"
)

func init() {
	gossa.RegisterPackage(&gossa.Package{
		Name: "crc32",
		Path: "hash/crc32",
		Deps: map[string]string{
			"errors":       "errors",
			"hash":         "hash",
			"internal/cpu": "cpu",
			"sync":         "sync",
			"sync/atomic":  "atomic",
			"unsafe":       "unsafe",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]gossa.NamedType{
			"Table": {reflect.TypeOf((*q.Table)(nil)).Elem(), "", ""},
		},
		AliasTypes: map[string]reflect.Type{},
		Vars: map[string]reflect.Value{
			"IEEETable": reflect.ValueOf(&q.IEEETable),
		},
		Funcs: map[string]reflect.Value{
			"Checksum":     reflect.ValueOf(q.Checksum),
			"ChecksumIEEE": reflect.ValueOf(q.ChecksumIEEE),
			"MakeTable":    reflect.ValueOf(q.MakeTable),
			"New":          reflect.ValueOf(q.New),
			"NewIEEE":      reflect.ValueOf(q.NewIEEE),
			"Update":       reflect.ValueOf(q.Update),
		},
		TypedConsts: map[string]gossa.TypedConst{},
		UntypedConsts: map[string]gossa.UntypedConst{
			"Castagnoli": {"untyped int", constant.MakeInt64(2197175160)},
			"IEEE":       {"untyped int", constant.MakeInt64(3988292384)},
			"Koopman":    {"untyped int", constant.MakeInt64(3945912366)},
			"Size":       {"untyped int", constant.MakeInt64(4)},
		},
	})
}
