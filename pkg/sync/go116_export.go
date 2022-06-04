// export by github.com/goplus/igop/cmd/qexp

//+build go1.16,!go1.17

package sync

import (
	q "sync"

	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "sync",
		Path: "sync",
		Deps: map[string]string{
			"internal/race": "race",
			"runtime":       "runtime",
			"sync/atomic":   "atomic",
			"unsafe":        "unsafe",
		},
		Interfaces: map[string]reflect.Type{
			"Locker": reflect.TypeOf((*q.Locker)(nil)).Elem(),
		},
		NamedTypes: map[string]igop.NamedType{
			"Cond":      {reflect.TypeOf((*q.Cond)(nil)).Elem(), "", "Broadcast,Signal,Wait"},
			"Map":       {reflect.TypeOf((*q.Map)(nil)).Elem(), "", "Delete,Load,LoadAndDelete,LoadOrStore,Range,Store,dirtyLocked,missLocked"},
			"Mutex":     {reflect.TypeOf((*q.Mutex)(nil)).Elem(), "", "Lock,Unlock,lockSlow,unlockSlow"},
			"Once":      {reflect.TypeOf((*q.Once)(nil)).Elem(), "", "Do,doSlow"},
			"Pool":      {reflect.TypeOf((*q.Pool)(nil)).Elem(), "", "Get,Put,getSlow,pin,pinSlow"},
			"RWMutex":   {reflect.TypeOf((*q.RWMutex)(nil)).Elem(), "", "Lock,RLock,RLocker,RUnlock,Unlock,rUnlockSlow"},
			"WaitGroup": {reflect.TypeOf((*q.WaitGroup)(nil)).Elem(), "", "Add,Done,Wait,state"},
		},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"NewCond": reflect.ValueOf(q.NewCond),
		},
		TypedConsts:   map[string]igop.TypedConst{},
		UntypedConsts: map[string]igop.UntypedConst{},
	})
}
