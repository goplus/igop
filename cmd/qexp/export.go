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
	src := r.Replace(template_tags)
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
