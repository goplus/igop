//go:build go1.16
// +build go1.16

package load

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/build"
	"go/token"
	"strings"

	"github.com/visualfc/goembed"
)

func buildIdent(name string) string {
	return fmt.Sprintf("__igop_embed_%x__", name)
}

var embed_head = `package $pkg

import (
	"embed"
	"unsafe"
)

func __igop_embed_buildFS__(list []struct {
	name string
	data string
	hash [16]byte
}) (f embed.FS) {
	fs := struct {
		files *[]struct {
			name string
			data string
			hash [16]byte
		}
	}{&list}
	return *(*embed.FS)(unsafe.Pointer(&fs))
}
`

// Embed check package embed data
func Embed(bp *build.Package, fset *token.FileSet, files []*ast.File, test bool, xtest bool) ([]byte, bool) {
	var pkgName string
	var ems []*goembed.Embed
	if xtest {
		pkgName = bp.Name + "_test"
		ems = goembed.CheckEmbed(bp.XTestEmbedPatternPos, fset, files)
	} else {
		pkgName = bp.Name
		ems = goembed.CheckEmbed(bp.EmbedPatternPos, fset, files)
		if test {
			if tems := goembed.CheckEmbed(bp.TestEmbedPatternPos, fset, files); len(tems) > 0 {
				ems = append(ems, tems...)
			}
		}
	}
	if len(ems) == 0 {
		return nil, false
	}
	r := goembed.NewResolve()
	var buf bytes.Buffer
	buf.WriteString(strings.Replace(embed_head, "$pkg", pkgName, 1))
	buf.WriteString("\nvar (\n")
	for _, v := range ems {
		v.Spec.Names[0].Name = "_"
		fs, _ := r.Load(bp.Dir, v)
		switch v.Kind {
		case goembed.EmbedBytes:
			buf.WriteString(fmt.Sprintf("\t%v = []byte(%v)\n", v.Name, buildIdent(fs[0].Name)))
		case goembed.EmbedString:
			buf.WriteString(fmt.Sprintf("\t%v = %v\n", v.Name, buildIdent(fs[0].Name)))
		case goembed.EmbedFiles:
			fs = goembed.BuildFS(fs)
			buf.WriteString(fmt.Sprintf("\t%v = ", v.Name))
			buf.WriteString(`__igop_embed_buildFS__([]struct {
	name string
	data string
	hash [16]byte
}{
`)
			for _, f := range fs {
				if len(f.Data) == 0 {
					buf.WriteString(fmt.Sprintf("\t{\"%v\",\"\",[16]byte{}},\n",
						f.Name))
				} else {
					buf.WriteString(fmt.Sprintf("\t{\"%v\",%v,[16]byte{%v}},\n",
						f.Name, buildIdent(f.Name), goembed.BytesToList(f.Hash[:])))
				}
			}
			buf.WriteString("})\n")
		}
	}
	buf.WriteString("\n)\n")
	buf.WriteString("\nvar (\n")
	for _, f := range r.Files() {
		if len(f.Data) == 0 {
			buf.WriteString(fmt.Sprintf("\t%v string\n",
				buildIdent(f.Name)))
		} else {
			buf.WriteString(fmt.Sprintf("\t%v = string(\"%v\")\n",
				buildIdent(f.Name), goembed.BytesToHex(f.Data)))
		}
	}
	buf.WriteString(")\n\n")
	return buf.Bytes(), true
}
