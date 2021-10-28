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
	if len(pkg.UntypedConsts) > 0 {
		imports = append(imports, `"go/constant"`)
		var hasToken bool
		for _, c := range pkg.UntypedConsts {
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
		"$DEPS", joinList(pkg.Deps),
		"$NAMEDTYPES", joinList(pkg.NamedTypes),
		"$INTERFACES", joinList(pkg.Interfaces),
		"$ALIASTYPES", joinList(pkg.AliasTypes),
		"$VARS", joinList(pkg.Vars),
		"$FUNCS", joinList(pkg.Funcs),
		"$METHODS", joinList(pkg.Methods),
		"$TYPEDCONSTS", joinList(pkg.TypedConsts),
		"$UNTYPEDCONSTS", joinList(pkg.UntypedConsts),
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

var template_export = `// export by github.com/goplus/gossa/cmd/qexp
package $PKGNAME

import (
	"$PKGPATH"

	"github.com/goplus/gossa"
)

func init() {
	gossa.RegisterPackage("$PKGPATH",extMap,typList)
}

var extMap = map[string]interface{}{$EXTMAP}

var typList = []interface{}{$TYPLIST}
`

var template_tags = `// export by github.com/goplus/gossa/cmd/qexp

$TAGS

package $PKGNAME

import (
	"$PKGPATH"

	"github.com/goplus/gossa"
)

func init() {
	gossa.RegisterPackage("$PKGPATH",extMap$ID,typList$ID)
}

var extMap$ID = map[string]interface{}{$EXTMAP}

var typList$ID = []interface{}{$TYPLIST}
`

var template_empty = `// export by github.com/goplus/gossa/cmd/qexp

$TAGS

package $PKGNAME

import (
	"github.com/goplus/gossa"
)

func init() {
	gossa.RegisterPackage("$PKGPATH",nil,nil)
}
`

/*
type TypedConst struct {
	Typ   reflect.Type
	Value constant.Value
}

type UntypedConst struct {
	Typ   string
	Value constant.Value
}

type Package struct {
	Name          string
	Path          string
	Types         []reflect.Type
	AliasTypes    map[string]reflect.Type
	Vars          map[string]reflect.Value
	Funcs         map[string]reflect.Value
	Methods       map[string]reflect.Value
	TypedConsts   map[string]TypedConst
	UntypedConsts map[string]UntypedConst
	Deps          map[string]string
}
*/

var template_pkg = `// export by github.com/goplus/gossa/cmd/qexp

$TAGS

package $PKGNAME

import (
	$IMPORTS

	"github.com/goplus/gossa"
)

func init() {
	gossa.InstallPackage(&gossa.Package {
		Name: "$PKGNAME",
		Path: "$PKGPATH",
		Deps: map[string]string{$DEPS},
		Interfaces: map[string]reflect.Type{$INTERFACES},
		NamedTypes: map[string]gossa.NamedType{$NAMEDTYPES},
		AliasTypes: map[string]reflect.Type{$ALIASTYPES},
		Vars: map[string]reflect.Value{$VARS},
		Funcs: map[string]reflect.Value{$FUNCS},
		Methods: map[string]reflect.Value{$METHODS},
		TypedConsts: map[string]gossa.TypedConst{$TYPEDCONSTS},
		UntypedConsts: map[string]gossa.UntypedConst{$UNTYPEDCONSTS},
	})
}
`
