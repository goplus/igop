// export by github.com/goplus/igop/cmd/qexp

//+build go1.15,!go1.16

package crc32

import (
	q "hash/crc32"

	"go/constant"
	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "crc32",
		Path: "hash/crc32",
		Deps: map[string]string{
			"errors":       "errors",
			"hash":         "hash",
			"internal/cpu": "cpu",
			"sync":         "sync",
			"unsafe":       "unsafe",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]igop.NamedType{
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
		TypedConsts: map[string]igop.TypedConst{},
		UntypedConsts: map[string]igop.UntypedConst{
			"Castagnoli": {"untyped int", constant.MakeInt64(int64(q.Castagnoli))},
			"IEEE":       {"untyped int", constant.MakeInt64(int64(q.IEEE))},
			"Koopman":    {"untyped int", constant.MakeInt64(int64(q.Koopman))},
			"Size":       {"untyped int", constant.MakeInt64(int64(q.Size))},
		},
	})
}
