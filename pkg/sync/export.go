// export by github.com/goplus/interp/cmd/qexp

package sync

import (
	"sync"

	"github.com/goplus/interp"
)

func init() {
	interp.RegisterPackage("sync", extMap, typList)
}

var extMap = map[string]interface{}{
	"(*sync.Cond).Broadcast":  (*sync.Cond).Broadcast,
	"(*sync.Cond).Signal":     (*sync.Cond).Signal,
	"(*sync.Cond).Wait":       (*sync.Cond).Wait,
	"(*sync.Map).Delete":      (*sync.Map).Delete,
	"(*sync.Map).Load":        (*sync.Map).Load,
	"(*sync.Map).LoadOrStore": (*sync.Map).LoadOrStore,
	"(*sync.Map).Range":       (*sync.Map).Range,
	"(*sync.Map).Store":       (*sync.Map).Store,
	"(*sync.Mutex).Lock":      (*sync.Mutex).Lock,
	"(*sync.Mutex).Unlock":    (*sync.Mutex).Unlock,
	"(*sync.Once).Do":         (*sync.Once).Do,
	"(*sync.Pool).Get":        (*sync.Pool).Get,
	"(*sync.Pool).Put":        (*sync.Pool).Put,
	"(*sync.RWMutex).Lock":    (*sync.RWMutex).Lock,
	"(*sync.RWMutex).RLock":   (*sync.RWMutex).RLock,
	"(*sync.RWMutex).RLocker": (*sync.RWMutex).RLocker,
	"(*sync.RWMutex).RUnlock": (*sync.RWMutex).RUnlock,
	"(*sync.RWMutex).Unlock":  (*sync.RWMutex).Unlock,
	"(*sync.WaitGroup).Add":   (*sync.WaitGroup).Add,
	"(*sync.WaitGroup).Done":  (*sync.WaitGroup).Done,
	"(*sync.WaitGroup).Wait":  (*sync.WaitGroup).Wait,
	"(sync.Locker).Lock":      (sync.Locker).Lock,
	"(sync.Locker).Unlock":    (sync.Locker).Unlock,
	"sync.NewCond":            sync.NewCond,
}

var typList = []interface{}{
	(*sync.Cond)(nil),
	(*sync.Locker)(nil),
	(*sync.Map)(nil),
	(*sync.Mutex)(nil),
	(*sync.Once)(nil),
	(*sync.Pool)(nil),
	(*sync.RWMutex)(nil),
	(*sync.WaitGroup)(nil),
}
