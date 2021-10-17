package main

import (
	"fmt"
	"go/format"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

func writeFile(dir string, file string, data []byte) error {
	err := os.MkdirAll(dir, 0777)
	if err != nil {
		return fmt.Errorf("make dir %v error: %v", dir, err)
	}
	filename := filepath.Join(dir, file)
	err = ioutil.WriteFile(filename, data, 0777)
	if err != nil {
		return fmt.Errorf("write file %v error: %v", filename, err)
	}
	return nil
}

func joinList(list []string) string {
	if len(list) == 0 {
		return ""
	}
	sort.Strings(list)
	return "\n\t" + strings.Join(list, ",\n\t") + ",\n"
}

func exportPkg(pkg *Package, sname string, id string, tagList []string) ([]byte, error) {
	imports := []string{fmt.Sprintf("%v %q\n", sname, pkg.Path)}
	imports = append(imports, `"reflect"`)
	if len(pkg.Consts) > 0 {
		imports = append(imports, `"go/constant"`)
		var hasToken bool
		for _, c := range pkg.Consts {
			if strings.Index(c, "token.") >= 0 {
				hasToken = true
				break
			}
		}
		if hasToken {
			imports = append(imports, `"go/token"`)
		}
	}

	r := strings.NewReplacer("$PKGNAME", pkg.Name,
		"$IMPORTS", strings.Join(imports, "\n"),
		"$PKGPATH", pkg.Path,
		"$TYPES", joinList(pkg.Types),
		"$VARS", joinList(pkg.Vars),
		"$FUNCS", joinList(pkg.Funcs),
		"$METHODS", joinList(pkg.Methods),
		"$CONSTS", joinList(pkg.Consts),
		"$TAGS", strings.Join(tagList, "\n"),
		"$ID", id)
	src := r.Replace(template_pkg)
	data, err := format.Source([]byte(src))
	if err != nil {
		return nil, fmt.Errorf("format pkg %v error: %v", src, err)
	}
	return data, nil
}

func exportSource(pkgPath string, id string, tagList []string, extList []string, typList []string) ([]byte, error) {
	plist := strings.Split(pkgPath, "/")
	pkgName := plist[len(plist)-1]
	var em string
	if len(extList) > 0 {
		sort.Strings(extList)
		em = "\n\t" + strings.Join(extList, ",\n\t") + ",\n"
	}
	var tl string
	if len(typList) > 0 {
		sort.Strings(typList)
		tl = "\n\t" + strings.Join(typList, ",\n\t") + ",\n"
	}
	r := strings.NewReplacer("$PKGNAME", pkgName,
		"$PKGPATH", pkgPath,
		"$EXTMAP", em,
		"$TYPLIST", tl,
		"$TAGS", strings.Join(tagList, "\n"),
		"$ID", id)
	var template string
	if len(typList) == 0 && len(extList) == 0 {
		template = template_empty
	} else {
		template = template_tags
	}
	src := r.Replace(template)
	data, err := format.Source([]byte(src))
	if err != nil {
		return nil, fmt.Errorf("format pkg %v error: %v", pkgPath, err)
	}
	return data, nil
}

var template_export = `// export by github.com/goplus/interp/cmd/qexp
package $PKGNAME

import (
	"$PKGPATH"

	"github.com/goplus/interp"
)

func init() {
	interp.RegisterPackage("$PKGPATH",extMap,typList)
}

var extMap = map[string]interface{}{$EXTMAP}

var typList = []interface{}{$TYPLIST}
`

var template_tags = `// export by github.com/goplus/interp/cmd/qexp

$TAGS

package $PKGNAME

import (
	"$PKGPATH"

	"github.com/goplus/interp"
)

func init() {
	interp.RegisterPackage("$PKGPATH",extMap$ID,typList$ID)
}

var extMap$ID = map[string]interface{}{$EXTMAP}

var typList$ID = []interface{}{$TYPLIST}
`

var template_empty = `// export by github.com/goplus/interp/cmd/qexp

$TAGS

package $PKGNAME

import (
	"github.com/goplus/interp"
)

func init() {
	interp.RegisterPackage("$PKGPATH",nil,nil)
}
`

/*
type ConstValue struct {
	Typ   string
	Value constant.Value
}

type Package struct {
	Name    string
	Path    string
	Types   []reflect.Type
	Vars    map[string]reflect.Value
	Funcs   map[string]reflect.Value
	Methods map[string]reflect.Value
	Consts  map[string]ConstValue
}
*/

var template_pkg = `// export by github.com/goplus/interp/cmd/qexp

$TAGS

package $PKGNAME

import (
	$IMPORTS

	"github.com/goplus/interp"
)

func init() {
	interp.InstallPackage(&interp.Package {
		Name: "$PKGNAME",
		Path: "$PKGPATH",
		Types: []reflect.Type{$TYPES},
		Vars: map[string]reflect.Value{$VARS},
		Funcs: map[string]reflect.Value{$FUNCS},
		Methods: map[string]reflect.Value{$METHODS},
		Consts: map[string]interp.ConstValue{$CONSTS},
	})
}
`
