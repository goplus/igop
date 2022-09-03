package main

import (
	"bytes"
	"errors"
	"flag"
	"fmt"
	"go/build"
	"log"
	"os/exec"
	"path/filepath"
	"strings"
)

var (
	flagExportDir      string
	flagBuildContext   string
	flagCustomTags     string
	flagExportFileName string
)

func init() {
	flag.StringVar(&flagExportDir, "outdir", "", "set export pkg path")
	flag.StringVar(&flagBuildContext, "contexts", "", "set customer build contexts goos_goarch list. eg \"drawin_amd64 darwin_arm64\"")
	flag.StringVar(&flagCustomTags, "addtags", "", "add custom tags, split by ;")
	flag.StringVar(&flagExportFileName, "filename", "export", "set export file name")
}

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) == 0 {
		flag.Usage()
	}
	if len(args) == 1 {
		pkgs, err := parserPkgs(args[0])
		if err != nil {
			log.Panicln(err)
		}
		if len(pkgs) == 0 {
			log.Panicln("not found packages")
		}
		args = pkgs
	}
	if flagExportFileName == "" {
		flagExportFileName = "export"
	}
	ctxList := parserContextList(flagBuildContext)
	Export(args, ctxList)
}

// pkg/...
func parserPkgs(expr string) ([]string, error) {
	data, err := runGoCommand("list", "-f={{.ImportPath}}={{.Name}}", expr)
	if err != nil {
		return nil, err
	}
	var pkgs []string
	for _, line := range strings.Split(string(data), "\n") {
		pos := strings.Index(line, "=")
		if pos != -1 {
			pkg, name := line[:pos], line[pos+1:]
			if name == "main" || isSkipPkg(pkg) {
				continue
			}
			pkgs = append(pkgs, pkg)
		}
	}
	return pkgs, nil
}

func runGoCommand(args ...string) (ret []byte, err error) {
	var stdout, stderr bytes.Buffer
	cmd := exec.Command("go", args...)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err = cmd.Run()
	if err == nil {
		ret = stdout.Bytes()
	} else if stderr.Len() > 0 {
		err = errors.New(stderr.String())
	}
	return
}

func isSkipPkg(pkg string) bool {
	switch pkg {
	case "unsafe":
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

func Export(pkgs []string, ctxList []*build.Context) {
	log.Println("process", pkgs)
	if len(ctxList) == 0 {
		ExportPkgs(pkgs, nil)
	} else {
		for _, ctx := range ctxList {
			ExportPkgs(pkgs, ctx)
		}
	}
}

func ExportPkgs(pkgs []string, ctx *build.Context) {
	prog := NewProgram(ctx)
	err := prog.Load(pkgs)
	if err != nil {
		log.Panicln(err)
	}
	for _, pkg := range pkgs {
		if pkg == "unsafe" {
			continue
		}
		fpath, err := ExportPkg(prog, pkg, ctx)
		if err != nil {
			log.Printf("export %v failed: %v\n", pkg, err)
		} else {
			log.Printf("export %v: %v\n", pkg, fpath)
		}
	}
}

func ExportPkg(prog *Program, pkg string, ctx *build.Context) (string, error) {
	e, err := prog.ExportPkg(pkg, "q")
	if err != nil {
		return "", err
	}
	var tags []string
	if flagCustomTags != "" {
		tags = strings.Split(flagCustomTags, ";")
	}
	data, err := exportPkg(e, "q", "", tags)
	if err != nil {
		return "", err
	}
	if flagExportDir == "" {
		fmt.Println(string(data))
		return "stdout", nil
	}
	fpath := filepath.Join(flagExportDir, pkg)
	var fname string
	if ctx != nil {
		fname = flagExportFileName + "_" + ctx.GOOS + "_" + ctx.GOARCH + ".go"
	} else {
		fname = flagExportFileName + ".go"
	}
	err = writeFile(fpath, fname, data)
	return filepath.Join(fpath, fname), err
}

func parserContextList(list string) (ctxs []*build.Context) {
	for _, info := range strings.Split(list, " ") {
		info = strings.TrimSpace(info)
		ar := strings.Split(info, "_")
		if len(ar) != 2 {
			continue
		}
		ctx := build.Default
		ctx.GOOS = ar[0]
		ctx.GOARCH = ar[1]
		ctxs = append(ctxs, &ctx)
	}
	return
}
