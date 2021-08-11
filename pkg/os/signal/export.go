// export by github.com/goplus/interp/cmd/qexp

package signal

import (
	"os/signal"

	"github.com/goplus/interp"
)

func init() {
	interp.RegisterPackage("os/signal", extMap, typList)
}

var extMap = map[string]interface{}{
	"os/signal.Ignore":  signal.Ignore,
	"os/signal.Ignored": signal.Ignored,
	"os/signal.Notify":  signal.Notify,
	"os/signal.Reset":   signal.Reset,
	"os/signal.Stop":    signal.Stop,
}

var typList = []interface{}{}
