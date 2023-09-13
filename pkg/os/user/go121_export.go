// export by github.com/goplus/igop/cmd/qexp

//go:build go1.21
// +build go1.21

package user

import (
	q "os/user"

	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "user",
		Path: "os/user",
		Deps: map[string]string{
			"fmt":                   "fmt",
			"internal/syscall/unix": "unix",
			"runtime":               "runtime",
			"strconv":               "strconv",
			"strings":               "strings",
			"sync":                  "sync",
			"syscall":               "syscall",
			"unsafe":                "unsafe",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{
			"Group":               reflect.TypeOf((*q.Group)(nil)).Elem(),
			"UnknownGroupError":   reflect.TypeOf((*q.UnknownGroupError)(nil)).Elem(),
			"UnknownGroupIdError": reflect.TypeOf((*q.UnknownGroupIdError)(nil)).Elem(),
			"UnknownUserError":    reflect.TypeOf((*q.UnknownUserError)(nil)).Elem(),
			"UnknownUserIdError":  reflect.TypeOf((*q.UnknownUserIdError)(nil)).Elem(),
			"User":                reflect.TypeOf((*q.User)(nil)).Elem(),
		},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"Current":       reflect.ValueOf(q.Current),
			"Lookup":        reflect.ValueOf(q.Lookup),
			"LookupGroup":   reflect.ValueOf(q.LookupGroup),
			"LookupGroupId": reflect.ValueOf(q.LookupGroupId),
			"LookupId":      reflect.ValueOf(q.LookupId),
		},
		TypedConsts:   map[string]igop.TypedConst{},
		UntypedConsts: map[string]igop.UntypedConst{},
	})
}
