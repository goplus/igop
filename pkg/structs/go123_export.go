// export by github.com/goplus/igop/cmd/qexp

//go:build go1.23 && !go1.24
// +build go1.23,!go1.24

package structs

import (
	q "structs"

	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name:       "structs",
		Path:       "structs",
		Deps:       map[string]string{},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{
			"HostLayout": reflect.TypeOf((*q.HostLayout)(nil)).Elem(),
		},
		AliasTypes:    map[string]reflect.Type{},
		Vars:          map[string]reflect.Value{},
		Funcs:         map[string]reflect.Value{},
		TypedConsts:   map[string]igop.TypedConst{},
		UntypedConsts: map[string]igop.UntypedConst{},
	})
}
