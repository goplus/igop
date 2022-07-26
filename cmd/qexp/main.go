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
	flagUseGoApi       bool
	flagBuildContext   string
	flagCustomTags     string
	flagExportFileName string
)

func init() {
	flag.StringVar(&flagExportDir, "outdir", "", "set export pkg path")
	//flag.BoolVar(&flagUseGoApi, "api", false, "export by $GOROOT/api")
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
		args = stdList
	}
	if flagExportFileName == "" {
		flagExportFileName = "export"
	}
	ctxList := parserContextList(flagBuildContext)
	//flagUseGoApi = false
	var ac *ApiCheck
	if flagUseGoApi {
		ac = NewApiCheck()
		err := ac.LoadBase("go1", "go1.1", "go1.2", "go1.3", "go1.4", "go1.5", "go1.6", "go1.7", "go1.8", "go1.9", "go1.10", "go1.11", "go1.12", "go1.13", "go1.14")
		if err != nil {
			panic(err)
		}
		err = ac.LoadApi("go1.15", "go1.16")
		if err != nil {
			fmt.Println("warning load api error", err)
		}
		var imports []string
		for _, pkg := range args {
			fileMap, err := ac.Export(pkg)
			if err != nil {
				panic(err)
			}
			if fileMap == nil {
				fmt.Println("warning skip empty export pkg", pkg)
				continue
			}
			imports = append(imports, fmt.Sprintf(`_ "github.com/goplus/interp/pkg/%v"`, pkg))
			if _, ok := fileMap[""]; !ok {
				fileMap[""] = &File{}
			}
			for _, v := range fileMap {
				data, err := exportSource(pkg, v.Name, v.Tags, v.ExtList, v.TypList)
				if err != nil {
					panic(err)
				}
				if flagExportDir == "" {
					fmt.Println("========", "export"+v.Name+".go")
					fmt.Println(string(data))
					continue
				}
				err = writeFile(filepath.Join(flagExportDir, pkg), "export"+v.Name+".go", data)
				if err != nil {
					panic(err)
				}
			}
			fmt.Println("export", pkg)
		}
		fmt.Printf("\nimport (\n\t%v\n)\n", strings.Join(imports, "\n\t"))
	} else {
		for _, pkg := range args {
			if pkg == "unsafe" {
				continue
			}
			Export(pkg, ctxList)
		}
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
