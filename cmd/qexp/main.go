package main

import (
	"flag"
	"fmt"
	"path/filepath"
)

var (
	flagExportDir string
	flagUseGoApi  bool
)

func init() {
	flag.StringVar(&flagExportDir, "outdir", "", "set export lib path")
	flag.BoolVar(&flagUseGoApi, "api", false, "export by $GOROOT/api")
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
		for _, pkg := range args {
			fileMap, err := ac.Export(pkg)
			if err != nil {
				panic(err)
			}
			if fileMap == nil {
				fmt.Println("warning skip empty export pkg", pkg)
				continue
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
	} else {
		for _, pkg := range args {
			prog, err := loadProgram(pkg)
			if err != nil {
				panic(fmt.Errorf("load pkg %v error: %v", pkg, err))
			}
			extList, typList := prog.Export(pkg)
			if len(extList) == 0 && len(typList) == 0 {
				fmt.Println("skip empty export", pkg)
				continue
			}
			data, err := exportSource(pkg, "", nil, extList, typList)
			if err != nil {
				panic(err)
			}
			if flagExportDir == "" {
				fmt.Println(string(data))
				continue
			}
			err = writeFile(filepath.Join(flagExportDir, pkg), "export.go", data)
			if err != nil {
				panic(err)
			}
			fmt.Println("export", pkg)
		}
	}
}
