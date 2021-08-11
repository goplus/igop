// export by github.com/goplus/interp/cmd/qexp

// +build go1.15

package time

import (
	"time"

	"github.com/goplus/interp"
)

func init() {
	interp.RegisterPackage("time", extMap_go115, typList_go115)
}

var extMap_go115 = map[string]interface{}{
	"(*time.Ticker).Reset": (*time.Ticker).Reset,
}

var typList_go115 = []interface{}{}
