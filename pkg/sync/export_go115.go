// export by github.com/goplus/interp/cmd/qexp

//go:build go1.15
// +build go1.15

package sync

import (
	"sync"

	"github.com/goplus/interp"
)

func init() {
	interp.RegisterPackage("sync", extMap_go115, typList_go115)
}

var extMap_go115 = map[string]interface{}{
	"(*sync.Map).LoadAndDelete": (*sync.Map).LoadAndDelete,
}

var typList_go115 = []interface{}{}
