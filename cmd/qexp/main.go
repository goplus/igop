package main

import (
	"flag"
	"fmt"
	"go/format"
	"go/token"
	"go/types"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"golang.org/x/tools/go/loader"
	"golang.org/x/tools/go/ssa"
	"golang.org/x/tools/go/ssa/ssautil"
)

var (
	flagExportDir string
)

func init() {
	flag.StringVar(&flagExportDir, "outdir", "", "set export lib path")
}

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) == 0 {
		flag.Usage()
	}

	for _, pkg := range args {
		prog, err := loadProgram(pkg)
		if err != nil {
			panic(fmt.Errorf("load pkg %v error: %v", pkg, err))
		}
		src := prog.Export(pkg)
		data, err := format.Source([]byte(src))
		if err != nil {
			panic(fmt.Errorf("format pkg %v error: %v", pkg, err))
		}
		if flagExportDir != "" {
			dir := filepath.Join(flagExportDir, pkg)
			err := os.MkdirAll(dir, 0777)
			if err != nil {
				panic(fmt.Errorf("make dir %v error: %v", dir, err))
			}
			file := filepath.Join(dir, "export.go")
			err = ioutil.WriteFile(file, data, 0777)
			if err != nil {
				panic(fmt.Errorf("write file %v error: %v", file, err))
			}
			fmt.Println("export", pkg, file)
		} else {
			fmt.Println(string(data))
		}
	}

	//prog.DumpExport("runtime")
}

type Program struct {
	prog *ssa.Program
}

func loadProgram(path string) (*Program, error) {
	var cfg loader.Config
	cfg.Import(path)

	iprog, err := cfg.Load()
	if err != nil {
		return nil, fmt.Errorf("conf.Load failed: %s", err)
	}

	prog := ssautil.CreateProgram(iprog, ssa.SanityCheckFunctions)
	prog.Build()
	return &Program{prog}, nil
}

func (p *Program) DumpDeps(path string) {
	pkg := p.prog.ImportedPackage(path)
	for _, im := range pkg.Pkg.Imports() {
		fmt.Println(im.Path())
	}
}

func (p *Program) dumpDeps(path string, sep string) {
	pkg := p.prog.ImportedPackage(path)
	for _, im := range pkg.Pkg.Imports() {
		fmt.Println(sep, im.Path())
		p.dumpDeps(im.Path(), sep+"  ")
	}
}

func (p *Program) DumpExport(path string) {
	pkg := p.prog.ImportedPackage(path)
	for k, v := range pkg.Members {
		if token.IsExported(k) {
			fmt.Printf("%v %v %T\n", k, v, v)
		}
	}
}

func (p *Program) Export(path string) string {
	pkg := p.prog.ImportedPackage(path)
	pkgPath := pkg.Pkg.Path()
	pkgName := pkg.Pkg.Name()
	var extList []string
	var typList []string
	for k, v := range pkg.Members {
		if token.IsExported(k) {
			switch t := v.(type) {
			case *ssa.NamedConst:
			case *ssa.Global:
				extList = append(extList, fmt.Sprintf("%q : &%v", pkgPath+"."+t.Name(), pkgName+"."+t.Name()))
			case *ssa.Function:
				extList = append(extList, fmt.Sprintf("%q : %v", pkgPath+"."+t.Name(), pkgName+"."+t.Name()))
			case *ssa.Type:
				named := t.Type().(*types.Named)
				typeName := named.Obj().Name()
				typList = append(typList, fmt.Sprintf("(*%v.%v)(nil)", pkgName, typeName))
				if named.Obj().Pkg() != pkg.Pkg {
					continue
				}
				methods := IntuitiveMethodSet(t.Type())
				for _, method := range methods {
					name := method.Obj().Name()
					if token.IsExported(name) {
						info := fmt.Sprintf("(%v).%v", method.Recv(), name)
						if pkgPath == pkgName {
							extList = append(extList, fmt.Sprintf("%q : %v", info, info))
						} else {
							var fn string
							if isPointer(method.Recv()) {
								fn = fmt.Sprintf("(*%v.%v).%v", pkgName, typeName, name)
							} else {
								fn = fmt.Sprintf("(%v.%v).%v", pkgName, typeName, name)
							}
							extList = append(extList, fmt.Sprintf("%q : %v", info, fn))
						}
					}
				}
			}
		}
	}
	sort.Strings(extList)
	extList = append([]string{fmt.Sprintf(`"%v.init" : func() {}`, path)}, extList...)
	var em string
	if len(extList) > 0 {
		em = "\n\t" + strings.Join(extList, ",\n\t") + ",\n"
	}
	var tl string
	if len(typList) > 0 {
		sort.Strings(typList)
		tl = "\n\t" + strings.Join(typList, ",\n\t") + ",\n"
	}
	r := strings.NewReplacer("$PKGNAME", pkgName, "$PKGPATH", pkgPath, "$EXTMAP", em, "$TYPLIST", tl)
	return r.Replace(template_init)
}

func isPointer(T types.Type) bool {
	_, ok := T.(*types.Pointer)
	return ok
}

// golang.org/x/tools/go/types/typeutil.IntuitiveMethodSet
func IntuitiveMethodSet(T types.Type) []*types.Selection {
	isPointerToConcrete := func(T types.Type) bool {
		ptr, ok := T.(*types.Pointer)
		return ok && !types.IsInterface(ptr.Elem())
	}

	var result []*types.Selection
	mset := types.NewMethodSet(T)
	if types.IsInterface(T) || isPointerToConcrete(T) {
		for i, n := 0, mset.Len(); i < n; i++ {
			result = append(result, mset.At(i))
		}
	} else {
		// T is some other concrete type.
		// Report methods of T and *T, preferring those of T.
		pmset := types.NewMethodSet(types.NewPointer(T))
		for i, n := 0, pmset.Len(); i < n; i++ {
			meth := pmset.At(i)
			if m := mset.Lookup(meth.Obj().Pkg(), meth.Obj().Name()); m != nil {
				meth = m
			}
			result = append(result, meth)
		}
	}
	return result
}

var template_export = `// export by github.com/goplus/interp/cmd/ssadump
package $PKGNAME

import "$PKGPATH"

func Exports() (name string, exts map[string]interface{}, types []interface{}) {
	return "$PKGPATH", extMap, typList
}

var extMap = map[string]interface{}{$EXTMAP}

var typList = []interface{}{$TYPLIST}

`

var template_init = `// export by github.com/goplus/interp/cmd/qexp
package $PKGNAME

import (
	"$PKGPATH"

	"github.com/goplus/interp"
)

func init() {
	interp.RegisterExternals(extMap)
	interp.RegisterTypeOf(typList...)
}

var extMap = map[string]interface{}{$EXTMAP}

var typList = []interface{}{$TYPLIST}

`
