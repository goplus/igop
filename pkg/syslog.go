//go:build !windows && !plan9
// +build !windows,!plan9

package pkg

import (
	_ "github.com/goplus/ixgo/pkg/log/syslog"
)
