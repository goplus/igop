// export by github.com/goplus/gossa/cmd/qexp

//go:build go1.17
// +build go1.17

package user

import (
	q "os/user"

	"reflect"

	"github.com/goplus/gossa"
)

func init() {
	gossa.RegisterPackage(&gossa.Package{
		Name: "user",
		Path: "os/user",
		Deps: map[string]string{
			"fmt":         "fmt",
			"runtime/cgo": "cgo",
			"strconv":     "strconv",
			"strings":     "strings",
			"sync":        "sync",
			"syscall":     "syscall",
			"unsafe":      "unsafe",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]gossa.NamedType{
			"Group":               {reflect.TypeOf((*q.Group)(nil)).Elem(), "", ""},
			"UnknownGroupError":   {reflect.TypeOf((*q.UnknownGroupError)(nil)).Elem(), "Error", ""},
			"UnknownGroupIdError": {reflect.TypeOf((*q.UnknownGroupIdError)(nil)).Elem(), "Error", ""},
			"UnknownUserError":    {reflect.TypeOf((*q.UnknownUserError)(nil)).Elem(), "Error", ""},
			"UnknownUserIdError":  {reflect.TypeOf((*q.UnknownUserIdError)(nil)).Elem(), "Error", ""},
			"User":                {reflect.TypeOf((*q.User)(nil)).Elem(), "", "GroupIds"},
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
		TypedConsts:   map[string]gossa.TypedConst{},
		UntypedConsts: map[string]gossa.UntypedConst{},
	})
}
