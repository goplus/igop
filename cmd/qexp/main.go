package main

import (
	"flag"
	"fmt"
	"go/build"
	"log"
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
	if len(args) == 1 && args[0] == "std" {
		//args = stdList
	}
	if flagExportFileName == "" {
		flagExportFileName = "export"
	}
	ctxList := parserContextList(flagBuildContext)
	for _, pkg := range args {
		if pkg == "unsafe" {
			continue
		}
		Export(pkg, ctxList)
	}
}

func Export(pkg string, ctxList []*build.Context) {
	log.Println("process", pkg)
	if len(ctxList) == 0 {
		fpath, err := ExportPkg(pkg, nil)
		if err != nil {
			log.Println("export failed", err)
		} else {
			log.Println("export", fpath)
		}
	} else {
		for _, ctx := range ctxList {
			fpath, err := ExportPkg(pkg, ctx)
			if err != nil {
				log.Println("export failed", err)
			} else {
				log.Println("export", fpath)
			}
		}
	}
}

func ExportPkg(pkg string, ctx *build.Context) (string, error) {
	prog, err := loadProgram(pkg, ctx)
	if err != nil {
		return "", fmt.Errorf("load pkg %v error: %v", pkg, err)
	}
	e := prog.ExportPkg(pkg, "q")
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
		return "", nil
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
