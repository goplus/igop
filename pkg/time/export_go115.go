// export by github.com/goplus/gossa/cmd/qexp

//go:build go1.15
// +build go1.15

package time

import (
	"time"

	"github.com/goplus/gossa"
)

func init() {
	gossa.RegisterPackage("time", extMap_go115, typList_go115)
}

var extMap_go115 = map[string]interface{}{
	"(*time.Ticker).Reset": (*time.Ticker).Reset,
}

var typList_go115 = []interface{}{}
