//go:build go1.16
// +build go1.16

package load

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/build"
	"go/printer"
	"go/token"
	"strconv"

	"github.com/visualfc/goembed"
)

var (
	__igop_embed_testdata_data2_txt__ string
)

func buildIdent(name string) string {
	return fmt.Sprintf("__igop_embed_%x__", name)
}

var embed_head = `package %v

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

func embedTypeError(fset *token.FileSet, spec *ast.ValueSpec) error {
	var buf bytes.Buffer
	printer.Fprint(&buf, fset, spec.Type)
	return fmt.Errorf("%v: go:embed cannot apply to var of type %v", fset.Position(spec.Names[0].NamePos), buf.String())
}

// Embed check package embed data
func Embed(bp *build.Package, fset *token.FileSet, files []*ast.File, test bool, xtest bool) ([]byte, error) {
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
		return nil, nil
	}
	r := goembed.NewResolve()
	for _, v := range ems {
		if len(v.Spec.Values) > 0 {
			return nil, fmt.Errorf("%v: go:embed cannot apply to var with initializer", v.Pos)
		}
		fs, err := r.Load(bp.Dir, v)
		if err != nil {
			return nil, err
		}
		switch v.Kind {
		default:
			switch t := v.Spec.Type.(type) {
			case *ast.Ident:
				// value = Type(data)
				// valid alias string or []byte type used by types.check
				v.Spec.Values = []ast.Expr{
					&ast.CallExpr{
						Fun: v.Spec.Type,
						Args: []ast.Expr{
							&ast.Ident{Name: buildIdent(fs[0].Name),
								NamePos: v.Spec.Names[0].NamePos},
						},
					}}
			case *ast.ArrayType:
				// value = Type(data)
				// valid alias string or []byte type used by types.check
				if _, ok := t.Elt.(*ast.Ident); ok {
					v.Spec.Values = []ast.Expr{
						&ast.CallExpr{
							Fun: v.Spec.Type,
							Args: []ast.Expr{
								&ast.Ident{Name: buildIdent(fs[0].Name),
									NamePos: v.Spec.Names[0].NamePos},
							},
						}}
					break
				}
				return nil, embedTypeError(fset, v.Spec)
			default:
				return nil, embedTypeError(fset, v.Spec)
			}
		case goembed.EmbedBytes:
			// value = []byte(data)
			v.Spec.Values = []ast.Expr{
				&ast.CallExpr{
					Fun:  v.Spec.Type,
					Args: []ast.Expr{ast.NewIdent(buildIdent(fs[0].Name))},
				}}
		case goembed.EmbedString:
			// value = data
			v.Spec.Values = []ast.Expr{ast.NewIdent(buildIdent(fs[0].Name))}
		case goembed.EmbedFiles:
			// value = __igop_embed_buildFS__([]struct{name string; data string; hash [16]byte}{...})
			fs = goembed.BuildFS(fs)
			elts := make([]ast.Expr, len(fs), len(fs))
			for i, f := range fs {
				if len(f.Data) == 0 {
					elts[i] = &ast.CompositeLit{
						Elts: []ast.Expr{
							&ast.BasicLit{Kind: token.STRING, Value: strconv.Quote(f.Name)},
							&ast.BasicLit{Kind: token.STRING, Value: `""`},
							&ast.CompositeLit{
								Type: &ast.ArrayType{
									Len: &ast.BasicLit{Kind: token.INT, Value: "16"},
									Elt: ast.NewIdent("byte"),
								},
							},
						},
					}
				} else {
					var hash [16]ast.Expr
					for j, v := range f.Hash {
						hash[j] = &ast.BasicLit{Kind: token.INT, Value: strconv.Itoa(int(v))}
					}
					elts[i] = &ast.CompositeLit{
						Elts: []ast.Expr{
							&ast.BasicLit{Kind: token.STRING, Value: strconv.Quote(f.Name)},
							ast.NewIdent(buildIdent(f.Name)),
							&ast.CompositeLit{
								Type: &ast.ArrayType{
									Len: &ast.BasicLit{Kind: token.INT, Value: "16"},
									Elt: ast.NewIdent("byte"),
								},
								Elts: hash[:],
							},
						},
					}
				}
			}
			call := &ast.CallExpr{
				Fun: ast.NewIdent("__igop_embed_buildFS__"),
				Args: []ast.Expr{
					&ast.CompositeLit{
						Type: &ast.ArrayType{
							Elt: &ast.StructType{
								Fields: &ast.FieldList{
									List: []*ast.Field{
										&ast.Field{
											Names: []*ast.Ident{ast.NewIdent("name")},
											Type:  ast.NewIdent("string"),
										},
										&ast.Field{
											Names: []*ast.Ident{ast.NewIdent("data")},
											Type:  ast.NewIdent("string"),
										},
										&ast.Field{
											Names: []*ast.Ident{ast.NewIdent("hash")},
											Type: &ast.ArrayType{
												Len: &ast.BasicLit{Kind: token.INT, Value: "16"},
												Elt: ast.NewIdent("byte"),
											},
										},
									},
								},
							},
						},
						Elts: elts,
					},
				},
			}
			v.Spec.Values = []ast.Expr{call}
		}
	}
	var buf bytes.Buffer
	fmt.Fprintf(&buf, embed_head, pkgName)
	buf.WriteString("\nconst (\n")
	for _, f := range r.Files() {
		if len(f.Data) == 0 {
			fmt.Fprintf(&buf, "\t%v = \"\"\n", buildIdent(f.Name))
		} else {
			fmt.Fprintf(&buf, "\t%v = \"%v\"\n", buildIdent(f.Name), goembed.BytesToHex(f.Data))
		}
	}
	buf.WriteString(")\n\n")
	fmt.Println(buf.String())
	return buf.Bytes(), nil
}
