// export by github.com/goplus/interp/cmd/qexp

package syslog

import (
	"github.com/goplus/interp"
)

func init() {
	interp.RegisterPackage("log/syslog", nil, nil)
}
