// export by github.com/goplus/interp/cmd/qexp

package exec

import (
	"os/exec"

	"github.com/goplus/interp"
)

func init() {
	interp.RegisterPackage("os/exec", extMap, typList)
}

var extMap = map[string]interface{}{
	"(*os/exec.Cmd).CombinedOutput":  (*exec.Cmd).CombinedOutput,
	"(*os/exec.Cmd).Output":          (*exec.Cmd).Output,
	"(*os/exec.Cmd).Run":             (*exec.Cmd).Run,
	"(*os/exec.Cmd).Start":           (*exec.Cmd).Start,
	"(*os/exec.Cmd).StderrPipe":      (*exec.Cmd).StderrPipe,
	"(*os/exec.Cmd).StdinPipe":       (*exec.Cmd).StdinPipe,
	"(*os/exec.Cmd).StdoutPipe":      (*exec.Cmd).StdoutPipe,
	"(*os/exec.Cmd).String":          (*exec.Cmd).String,
	"(*os/exec.Cmd).Wait":            (*exec.Cmd).Wait,
	"(*os/exec.Error).Error":         (*exec.Error).Error,
	"(*os/exec.Error).Unwrap":        (*exec.Error).Unwrap,
	"(*os/exec.ExitError).Error":     (*exec.ExitError).Error,
	"(os/exec.ExitError).ExitCode":   (exec.ExitError).ExitCode,
	"(os/exec.ExitError).Exited":     (exec.ExitError).Exited,
	"(os/exec.ExitError).Pid":        (exec.ExitError).Pid,
	"(os/exec.ExitError).String":     (exec.ExitError).String,
	"(os/exec.ExitError).Success":    (exec.ExitError).Success,
	"(os/exec.ExitError).Sys":        (exec.ExitError).Sys,
	"(os/exec.ExitError).SysUsage":   (exec.ExitError).SysUsage,
	"(os/exec.ExitError).SystemTime": (exec.ExitError).SystemTime,
	"(os/exec.ExitError).UserTime":   (exec.ExitError).UserTime,
	"os/exec.Command":                exec.Command,
	"os/exec.CommandContext":         exec.CommandContext,
	"os/exec.ErrNotFound":            &exec.ErrNotFound,
	"os/exec.LookPath":               exec.LookPath,
}

var typList = []interface{}{
	(*exec.Cmd)(nil),
	(*exec.Error)(nil),
	(*exec.ExitError)(nil),
}
