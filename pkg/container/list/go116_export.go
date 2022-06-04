// export by github.com/goplus/igop/cmd/qexp

//+build go1.16,!go1.17

package list

import (
	q "container/list"

	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name:       "list",
		Path:       "container/list",
		Deps:       map[string]string{},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]igop.NamedType{
			"Element": {reflect.TypeOf((*q.Element)(nil)).Elem(), "", "Next,Prev"},
			"List":    {reflect.TypeOf((*q.List)(nil)).Elem(), "", "Back,Front,Init,InsertAfter,InsertBefore,Len,MoveAfter,MoveBefore,MoveToBack,MoveToFront,PushBack,PushBackList,PushFront,PushFrontList,Remove,insert,insertValue,lazyInit,move,remove"},
		},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"New": reflect.ValueOf(q.New),
		},
		TypedConsts:   map[string]igop.TypedConst{},
		UntypedConsts: map[string]igop.UntypedConst{},
	})
}
