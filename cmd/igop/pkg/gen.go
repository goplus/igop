//go:build ignore
// +build ignore

package main

import (
	"os"
	"os/exec"
	"strings"
)

var (
	pkgList = []string{
		"github.com/modern-go/reflect2",
	}
)

func main() {
	var tags string
	var name string
	name = "export"
	pkgs := pkgList
	cmd := exec.Command("qexp", "-outdir", ".", "-addtags", tags, "-filename", name)
	cmd.Args = append(cmd.Args, pkgs...)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		panic(err)
	}
}

//skip syscall
func stdList() []string {
	cmd := exec.Command("go", "list", "std")
	data, err := cmd.Output()
	if err != nil {
		panic(err)
	}
	var ar []string
	for _, v := range strings.Split(string(data), "\n") {
		if v == "" {
			continue
		}
		if isSkipPkg(v) {
			continue
		}
		ar = append(ar, v)
	}
	return ar
}

func isSkipPkg(pkg string) bool {
	switch pkg {
	case "syscall", "unsafe":
		return true
	case "runtime/cgo", "runtime/race":
		return true
	case "time/tzdata":
		return true
	default:
		if strings.HasPrefix(pkg, "vendor/") {
			return true
		}
		for _, v := range strings.Split(pkg, "/") {
			if v == "internal" {
				return true
			}
		}
	}
	return false
}

func osarchList() []string {
	//go tool dist list
	cmd := exec.Command("go", "tool", "dist", "list")
	data, err := cmd.Output()
	if err != nil {
		panic(err)
	}
	var ar []string
	for _, v := range strings.Split(string(data), "\n") {
		if v == "" {
			continue
		}
		ar = append(ar, strings.Replace(v, "/", "_", 1))
	}
	return ar
}
