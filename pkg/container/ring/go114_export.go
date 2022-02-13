// export by github.com/goplus/gossa/cmd/qexp

//+build go1.14,!go1.15

package ring

import (
	q "container/ring"

	"reflect"

	"github.com/goplus/gossa"
)

func init() {
	gossa.RegisterPackage(&gossa.Package{
		Name:       "ring",
		Path:       "container/ring",
		Deps:       map[string]string{},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]gossa.NamedType{
			"Ring": {reflect.TypeOf((*q.Ring)(nil)).Elem(), "", "Do,Len,Link,Move,Next,Prev,Unlink,init"},
		},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"New": reflect.ValueOf(q.New),
		},
		TypedConsts:   map[string]gossa.TypedConst{},
		UntypedConsts: map[string]gossa.UntypedConst{},
	})
}
