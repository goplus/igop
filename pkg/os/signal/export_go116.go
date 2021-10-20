// export by github.com/goplus/gossa/cmd/qexp

//go:build go1.16
// +build go1.16

package signal

import (
	"os/signal"

	"github.com/goplus/gossa"
)

func init() {
	gossa.RegisterPackage("os/signal", extMap_go116, typList_go116)
}

var extMap_go116 = map[string]interface{}{
	"os/signal.NotifyContext": signal.NotifyContext,
}

var typList_go116 = []interface{}{}
