//go:build !windows && !plan9
// +build !windows,!plan9

package pkg

import (
	_ "github.com/goplus/igop/pkg/log/syslog"
)
