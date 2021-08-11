// export by github.com/goplus/interp/cmd/qexp

// +build darwin freebsd linux netbsd openbsd

package syslog

import (
	"log/syslog"

	"github.com/goplus/interp"
)

func init() {
	interp.RegisterPackage("log/syslog", extMap_3927522974, typList_3927522974)
}

var extMap_3927522974 = map[string]interface{}{
	"(*log/syslog.Writer).Alert":   (*syslog.Writer).Alert,
	"(*log/syslog.Writer).Close":   (*syslog.Writer).Close,
	"(*log/syslog.Writer).Crit":    (*syslog.Writer).Crit,
	"(*log/syslog.Writer).Debug":   (*syslog.Writer).Debug,
	"(*log/syslog.Writer).Emerg":   (*syslog.Writer).Emerg,
	"(*log/syslog.Writer).Err":     (*syslog.Writer).Err,
	"(*log/syslog.Writer).Info":    (*syslog.Writer).Info,
	"(*log/syslog.Writer).Notice":  (*syslog.Writer).Notice,
	"(*log/syslog.Writer).Warning": (*syslog.Writer).Warning,
	"(*log/syslog.Writer).Write":   (*syslog.Writer).Write,
	"log/syslog.Dial":              syslog.Dial,
	"log/syslog.New":               syslog.New,
	"log/syslog.NewLogger":         syslog.NewLogger,
}

var typList_3927522974 = []interface{}{
	(*syslog.Priority)(nil),
	(*syslog.Writer)(nil),
}
