//go:build !windows && !plan9
// +build !windows,!plan9

package pkg

import (
	_ "github.com/goplus/gossa/pkg/log/syslog"
)
