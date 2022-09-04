package export

import (
	"errors"
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

const EmptyPackage = "empty package"

var (
	errEmptyPackage = errors.New(EmptyPackage)
)

func exportPkg(pkg *Package, sname string, id string, tagList []string) ([]byte, error) {
	imports := []string{fmt.Sprintf("%v %q\n", sname, pkg.Path)}
	imports = append(imports, `"reflect"`)
	if len(pkg.UntypedConsts) > 0 || len(pkg.TypedConsts) > 0 {
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
	tmpl := template_pkg
	if pkg.IsEmpty() {
		tmpl = template_empty_pkg
		imports = []string{fmt.Sprintf("_ %q", pkg.Path)}
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
		"$TYPEDCONSTS", joinList(pkg.TypedConsts),
		"$UNTYPEDCONSTS", joinList(pkg.UntypedConsts),
		"$TAGS", strings.Join(tagList, "\n"),
		"$ID", id)
	src := r.Replace(tmpl)
	data, err := format.Source([]byte(src))
	if err != nil {
		return nil, fmt.Errorf("format pkg %v error: %v", src, err)
	}
	return data, nil
}

var template_pkg = `// export by github.com/goplus/igop/cmd/qexp

$TAGS

package $PKGNAME

import (
	$IMPORTS

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package {
		Name: "$PKGNAME",
		Path: "$PKGPATH",
		Deps: map[string]string{$DEPS},
		Interfaces: map[string]reflect.Type{$INTERFACES},
		NamedTypes: map[string]reflect.Type{$NAMEDTYPES},
		AliasTypes: map[string]reflect.Type{$ALIASTYPES},
		Vars: map[string]reflect.Value{$VARS},
		Funcs: map[string]reflect.Value{$FUNCS},
		TypedConsts: map[string]igop.TypedConst{$TYPEDCONSTS},
		UntypedConsts: map[string]igop.UntypedConst{$UNTYPEDCONSTS},
	})
}
`

var template_empty_pkg = `// export by github.com/goplus/igop/cmd/qexp

$TAGS

package $PKGNAME

import (
	$IMPORTS

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package {
		Name: "$PKGNAME",
		Path: "$PKGPATH",
		Deps: map[string]string{$DEPS},
	})
}
`
