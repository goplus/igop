//go:build ignore
// +build ignore

package main

import (
	"fmt"
	"go/format"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

func main() {
	ver := runtime.Version()[:6]
	var tags string
	var name string
	var fname string
	switch ver {
	case "go1.17":
		tags = "//+build go1.17"
		name = "go117_export"
		fname = "go117_pkgs.go"
	case "go1.16":
		tags = "//+build go1.16,!go1.17"
		name = "go116_export"
		fname = "go116_pkgs.go"
	case "go1.15":
		tags = "//+build go1.15,!go1.16"
		name = "go115_export"
		fname = "go115_pkgs.go"
	case "go1.14":
		tags = "//+build go1.14,!go1.15"
		name = "go114_export"
		fname = "go114_pkgs.go"
	}

	pkgs := stdList()
	list := osarchList()

	log.Println(ver, name, tags)
	log.Println(pkgs)
	log.Println(list)

	cmd := exec.Command("go", "run", "../cmd/qexp", "-outdir", ".", "-addtags", tags, "-filename", name)
	cmd.Args = append(cmd.Args, pkgs...)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		panic(err)
	}

	cmd = exec.Command("go", "run", "../cmd/qexp", "-outdir", ".", "-addtags", tags, "-filename", name, "-contexts", strings.Join(list, " "), "syscall")
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	if ver >= "go1.17" {
		cmd.Env = os.Environ()
		cmd.Env = append(cmd.Env, "GOEXPERIMENT=noregabi")
	}
	err = cmd.Run()
	if err != nil {
		panic(err)
	}
	err = makepkg("./"+fname, []string{tags}, pkgs)
	if err != nil {
		panic(err)
	}
}

func makepkg(fname string, tags []string, std []string) error {
	//_ github.com/goplus/gossa/pkg
	var pkgs []string
	for _, v := range std {
		if strings.HasPrefix(v, "testing") {
			continue
		}
		if v == "log/syslog" {
			continue
		}
		pkgs = append(pkgs, fmt.Sprintf(`_ "github.com/goplus/gossa/pkg/%v"`, v))
	}
	pkgs = append(pkgs, fmt.Sprintf(`_ "github.com/goplus/gossa/pkg/syscall"`))
	r := strings.NewReplacer("$TAGS", strings.Join(tags, "\n"), "$PKGS", strings.Join(pkgs, "\t\n"))
	src := r.Replace(tmpl)
	data, err := format.Source([]byte(src))
	if err != nil {
		return err
	}
	return ioutil.WriteFile(fname, []byte(data), 0644)
}

var tmpl = `package pkg

$TAGS

import (
	$PKGS
)
`

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
