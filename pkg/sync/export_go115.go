// export by github.com/goplus/gossa/cmd/qexp

//go:build go1.15
// +build go1.15

package sync

import (
	"sync"

	"github.com/goplus/gossa"
)

func init() {
	gossa.RegisterPackage("sync", extMap_go115, typList_go115)
}

var extMap_go115 = map[string]interface{}{
	"(*sync.Map).LoadAndDelete": (*sync.Map).LoadAndDelete,
}

var typList_go115 = []interface{}{}
